package ParsRubika

// سازنده ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const defaultBaseURL = "https://botapi.rubika.ir/v3"

// HandlerFunc برای پردازش آپدیت‌ها
type HandlerFunc func(ctx context.Context, update *Update) error

// Client کلاینت اصلی برای کار با API روبیکا
type Client struct {
	token        string
	baseURL      string
	httpClient   *http.Client
	botID        *string
	mu           sync.Mutex
	lastSentTime time.Time
}

// ClientOption برای تنظیمات اولیه کلاینت
type ClientOption func(*Client)

// WithBaseURL برای تغییر آدرس پایه API
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient برای تنظیم کلاینت HTTP سفارشی
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

// NewClient یک نمونه جدید از کلاینت می‌سازد
func NewClient(token string, opts ...ClientOption) *Client {
	c := &Client{
		token:   token,
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// --- متدهای داخلی و کمکی ---

func (c *Client) recordMessageSent() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lastSentTime = time.Now()
}

func (c *Client) isInCooldown() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return time.Since(c.lastSentTime) < 1*time.Second
}

func (c *Client) getBotID(ctx context.Context) error {
	if c.botID != nil {
		return nil
	}
	bot, err := c.GetMe(ctx)
	if err != nil {
		return fmt.Errorf("خطا در دریافت شناسه بات: %w", err)
	}
	c.botID = &bot.BotID
	log.Printf("شناسه بات دریافت شد: %s", *c.botID)
	return nil
}

func (c *Client) makeRequest(ctx context.Context, method string, apiMethod string, reqBody, respBody any) error {
	url := fmt.Sprintf("%s/%s/%s", c.baseURL, c.token, apiMethod)

	var bodyReader io.Reader
	if reqBody != nil {
		jsonBody, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("خطا در ساخت درخواست JSON: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("خطا در ایجاد درخواست HTTP: %w", err)
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("خطا در ارسال درخواست: %w", err)
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("خطا در خواندن پاسخ: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		apiErr.StatusCode = resp.StatusCode
		// تلاش برای پارس کردن خطا، در صورت شکست کل بدنه را به عنوان پیام خطا نمایش بده
		if json.Unmarshal(respData, &apiErr) != nil || apiErr.Message == "" {
			apiErr.Message = string(respData)
		}
		return &apiErr
	}

	if respBody != nil {
		var rawResp map[string]json.RawMessage
		if err := json.Unmarshal(respData, &rawResp); err != nil {
			return fmt.Errorf("خطا در پارس کردن پاسخ: %w", err)
		}
		if data, ok := rawResp["data"]; ok {
			if err := json.Unmarshal(data, respBody); err != nil {
				return fmt.Errorf("خطا در پارس کردن داده‌های پاسخ: %w", err)
			}
		}
	}
	return nil
}

// ImageToBytes تبدیل تصویر به آرایه بایت
func (c *Client) ImageToBytes(imagePath string) ([]byte, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("خطا در باز کردن فایل تصویر: %w", err)
	}
	defer file.Close()
	return io.ReadAll(file)
}

// CheckURLContentType بررسی نوع محتوای یک URL
func (c *Client) CheckURLContentType(url string) (string, error) {
	resp, err := c.httpClient.Head(url)
	if err != nil {
		return "", fmt.Errorf("خطا در ارسال درخواست HEAD: %w", err)
	}
	defer resp.Body.Close()
	return resp.Header.Get("Content-Type"), nil
}

// --- متدهای رسمی API بات ---

func (c *Client) GetMe(ctx context.Context) (*Bot, error) {
	var bot Bot
	err := c.makeRequest(ctx, http.MethodPost, "getMe", nil, &bot)
	return &bot, err
}

func (c *Client) SendMessage(ctx context.Context, req *SendMessageRequest) (string, error) {
	var resp MessageIDResponse
	err := c.makeRequest(ctx, http.MethodPost, "sendMessage", req, &resp)
	if err == nil {
		c.recordMessageSent()
	}
	return resp.MessageID, err
}

func (c *Client) SendPoll(ctx context.Context, req *SendPollRequest) (string, error) {
	var resp MessageIDResponse
	err := c.makeRequest(ctx, http.MethodPost, "sendPoll", req, &resp)
	if err == nil {
		c.recordMessageSent()
	}
	return resp.MessageID, err
}

func (c *Client) ForwardMessage(ctx context.Context, req *ForwardMessageRequest) (string, error) {
	var resp ForwardMessageResponse
	err := c.makeRequest(ctx, http.MethodPost, "forwardMessage", req, &resp)
	if err == nil {
		c.recordMessageSent()
	}
	return resp.NewMessageID, err
}

func (c *Client) EditMessageText(ctx context.Context, req *EditMessageTextRequest) error {
	err := c.makeRequest(ctx, http.MethodPost, "editMessageText", req, nil)
	if err == nil {
		c.recordMessageSent()
	}
	return err
}

func (c *Client) EditMessageKeypad(ctx context.Context, req *EditMessageKeypadRequest) error {
	// طبق مستندات، متد مربوط به ویرایش کیبورد اینلاین editInlineKeypad است
	err := c.makeRequest(ctx, http.MethodPost, "editInlineKeypad", req, nil)
	if err == nil {
		c.recordMessageSent()
	}
	return err
}

func (c *Client) DeleteMessage(ctx context.Context, req *DeleteMessageRequest) error {
	err := c.makeRequest(ctx, http.MethodPost, "deleteMessage", req, nil)
	if err == nil {
		c.recordMessageSent()
	}
	return err
}

func (c *Client) SendLocation(ctx context.Context, req *SendLocationRequest) (string, error) {
	var resp MessageIDResponse
	err := c.makeRequest(ctx, http.MethodPost, "sendLocation", req, &resp)
	if err == nil {
		c.recordMessageSent()
	}
	return resp.MessageID, err
}

func (c *Client) SendContact(ctx context.Context, req *SendContactRequest) (string, error) {
	var resp MessageIDResponse
	err := c.makeRequest(ctx, http.MethodPost, "sendContact", req, &resp)
	if err == nil {
		c.recordMessageSent()
	}
	return resp.MessageID, err
}

func (c *Client) GetChat(ctx context.Context, chatID string) (*Chat, error) {
	var chat Chat
	err := c.makeRequest(ctx, http.MethodPost, "getChat", &GetChatRequest{ChatID: chatID}, &chat)
	return &chat, err
}

func (c *Client) SendFile(ctx context.Context, req *SendFileRequest) (string, error) {
	var resp MessageIDResponse
	err := c.makeRequest(ctx, http.MethodPost, "sendFile", req, &resp)
	if err == nil {
		c.recordMessageSent()
	}
	return resp.MessageID, err
}

func (c *Client) GetFile(ctx context.Context, fileID string) (*GetFileResponse, error) {
	var resp GetFileResponse
	err := c.makeRequest(ctx, http.MethodPost, "getFile", &GetFileRequest{FileID: fileID}, &resp)
	return &resp, err
}

func (c *Client) RequestSendFile(ctx context.Context, fileType FileTypeEnum) (*RequestSendFileResponse, error) {
	var resp RequestSendFileResponse
	err := c.makeRequest(ctx, http.MethodPost, "requestSendFile", &RequestSendFileRequest{Type: fileType}, &resp)
	return &resp, err
}

// اصلاح شده: ترتیب بستن writer و تنظیم Header اصلاح شد
func (c *Client) UploadFile(uploadURL, filePath string) (*FileUploadResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("خطا در باز کردن فایل: %w", err)
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("خطا در ساخت فرم فایل: %w", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("خطا در کپی کردن محتوای فایل: %w", err)
	}
	// ابتدا writer را ببندید تا boundary نهایی نوشته شود
	writer.Close()

	req, err := http.NewRequest("POST", uploadURL, &requestBody)
	if err != nil {
		return nil, fmt.Errorf("خطا در ایجاد درخواست آپلود: %w", err)
	}
	// سپس Content-Type را با boundary صحیح تنظیم کنید
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال درخواست آپلود: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("آپلود با خطا %d مواجه شد: %s", resp.StatusCode, string(body))
	}

	var uploadResp FileUploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResp); err != nil {
		return nil, fmt.Errorf("خطا در پارس کردن پاسخ آپلود: %w", err)
	}
	return &uploadResp, nil
}

func (c *Client) GetUpdates(ctx context.Context, offsetID *string, limit int) (*GetUpdatesResponse, error) {
	req := &GetUpdatesRequest{OffsetID: offsetID, Limit: limit}
	var resp GetUpdatesResponse
	err := c.makeRequest(ctx, http.MethodPost, "getUpdates", req, &resp)
	return &resp, err
}

func (c *Client) UpdateBotEndpoints(ctx context.Context, url string, endpointType UpdateEndpointTypeEnum) error {
	req := map[string]interface{}{"url": url, "type": endpointType}
	return c.makeRequest(ctx, http.MethodPost, "updateBotEndpoints", req, nil)
}

func (c *Client) SetCommands(ctx context.Context, req *SetCommandsRequest) error {
	return c.makeRequest(ctx, http.MethodPost, "setCommands", req, nil)
}

func (c *Client) EditChatKeypad(ctx context.Context, req *EditChatKeypadRequest) error {
	err := c.makeRequest(ctx, http.MethodPost, "editChatKeypad", req, nil)
	if err == nil {
		c.recordMessageSent()
	}
	return err
}

// --- متدهای گروه و کانال ---

func (c *Client) BanChatMember(ctx context.Context, req *BanChatMemberRequest) error {
	return c.makeRequest(ctx, http.MethodPost, "banChatMember", req, nil)
}

func (c *Client) UnbanChatMember(ctx context.Context, req *UnbanChatMemberRequest) error {
	return c.makeRequest(ctx, http.MethodPost, "unbanChatMember", req, nil)
}

func (c *Client) PromoteChatMember(ctx context.Context, req *PromoteChatMemberRequest) error {
	return c.makeRequest(ctx, http.MethodPost, "promoteChatMember", req, nil)
}

func (c *Client) GetChatMember(ctx context.Context, req *GetChatMemberRequest) (*GetChatMemberResponse, error) {
	var resp GetChatMemberResponse
	err := c.makeRequest(ctx, http.MethodPost, "getChatMember", req, &resp)
	return &resp, err
}

// نام این متد با فراخوانی در bot_logic.go مطابقت دارد
func (c *Client) GetChatAdministrators(ctx context.Context, req *GetChatAdministratorsRequest) (*GetChatAdministratorsResponse, error) {
	var resp GetChatAdministratorsResponse
	err := c.makeRequest(ctx, http.MethodPost, "getChatAdministrators", req, &resp)
	return &resp, err
}

func (c *Client) GetChatMemberCount(ctx context.Context, req *GetChatMemberCountRequest) (*GetChatMemberCountResponse, error) {
	var resp GetChatMemberCountResponse
	err := c.makeRequest(ctx, http.MethodPost, "getChatMemberCount", req, &resp)
	return &resp, err
}

func (c *Client) PinChatMessage(ctx context.Context, req *PinChatMessageRequest) error {
	return c.makeRequest(ctx, http.MethodPost, "pinChatMessage", req, nil)
}

func (c *Client) UnpinChatMessage(ctx context.Context, req *UnpinChatMessageRequest) error {
	return c.makeRequest(ctx, http.MethodPost, "unpinChatMessage", req, nil)
}

func (c *Client) UnpinAllChatMessages(ctx context.Context, chatID string) error {
	req := map[string]interface{}{"chat_id": chatID}
	return c.makeRequest(ctx, http.MethodPost, "unpinAllChatMessages", req, nil)
}

// --- متدهای غیررسمی (مدیریت صفحه و استوری) ---
// (این متدها بدون تغییر هستند چون منطق داخلی آن‌ها makeRequest را فراخوانی می‌کند که اصلاح شد)

func (c *Client) DownloadStory(ctx context.Context, storyID string) (string, error) {
	var resp struct {
		URL string `json:"url"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "downloadStory", map[string]string{"story_id": storyID}, &resp)
	return resp.URL, err
}

func (c *Client) LikePost(ctx context.Context, postID string) error {
	return c.makeRequest(ctx, http.MethodPost, "likePost", map[string]string{"post_id": postID}, nil)
}

func (c *Client) AcceptFollowRequest(ctx context.Context, userID string) error {
	return c.makeRequest(ctx, http.MethodPost, "acceptFollowRequest", map[string]string{"user_id": userID}, nil)
}

func (c *Client) AddPost(ctx context.Context, text string, fileID *string) error {
	req := map[string]interface{}{"text": text}
	if fileID != nil {
		req["file_id"] = *fileID
	}
	return c.makeRequest(ctx, http.MethodPost, "addPost", req, nil)
}

func (c *Client) AddPostViewCount(ctx context.Context, postID string) error {
	return c.makeRequest(ctx, http.MethodPost, "addPostViewCount", map[string]string{"post_id": postID}, nil)
}

func (c *Client) AddPostViewTime(ctx context.Context, postID string, duration int) error {
	return c.makeRequest(ctx, http.MethodPost, "addPostViewTime", map[string]interface{}{"post_id": postID, "duration": duration}, nil)
}

func (c *Client) AddStory(ctx context.Context, fileID string, caption string) error {
	return c.makeRequest(ctx, http.MethodPost, "addStory", map[string]interface{}{"file_id": fileID, "caption": caption}, nil)
}

func (c *Client) AddViewStory(ctx context.Context, storyID string) error {
	return c.makeRequest(ctx, http.MethodPost, "addViewStory", map[string]string{"story_id": storyID}, nil)
}

func (c *Client) BlockPage(ctx context.Context, pageID string) error {
	return c.makeRequest(ctx, http.MethodPost, "blockPage", map[string]string{"page_id": pageID}, nil)
}

func (c *Client) Comment(ctx context.Context, postID, text string) error {
	return c.makeRequest(ctx, http.MethodPost, "comment", map[string]interface{}{"post_id": postID, "text": text}, nil)
}

func (c *Client) CreateHighlight(ctx context.Context, title string, storyIDs []string) error {
	return c.makeRequest(ctx, http.MethodPost, "createHighlight", map[string]interface{}{"title": title, "story_ids": storyIDs}, nil)
}

func (c *Client) CreatePage(ctx context.Context, username, bio string) error {
	return c.makeRequest(ctx, http.MethodPost, "createPage", map[string]interface{}{"username": username, "bio": bio}, nil)
}

func (c *Client) DeclineFollowRequest(ctx context.Context, userID string) error {
	return c.makeRequest(ctx, http.MethodPost, "declineFollowRequest", map[string]string{"user_id": userID}, nil)
}

func (c *Client) DeleteComment(ctx context.Context, commentID string) error {
	return c.makeRequest(ctx, http.MethodPost, "deleteComment", map[string]string{"comment_id": commentID}, nil)
}

func (c *Client) DeletePage(ctx context.Context, pageID string) error {
	return c.makeRequest(ctx, http.MethodPost, "deletePage", map[string]string{"page_id": pageID}, nil)
}

func (c *Client) DeletePost(ctx context.Context, postID string) error {
	return c.makeRequest(ctx, http.MethodPost, "deletePost", map[string]string{"post_id": postID}, nil)
}

func (c *Client) DeleteStory(ctx context.Context, storyID string) error {
	return c.makeRequest(ctx, http.MethodPost, "deleteStory", map[string]string{"story_id": storyID}, nil)
}

func (c *Client) GetExplorePosts(ctx context.Context) ([]Post, error) {
	var resp struct {
		Posts []Post `json:"posts"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getExplorePosts", nil, &resp)
	return resp.Posts, err
}

func (c *Client) GetHashTagTrend(ctx context.Context) ([]string, error) {
	var resp struct {
		Hashtags []string `json:"hashtags"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getHashTagTrend", nil, &resp)
	return resp.Hashtags, err
}

func (c *Client) GetHighlightStories(ctx context.Context, highlightID string) ([]Story, error) {
	var resp struct {
		Stories []Story `json:"stories"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getHighlightStories", map[string]string{"highlight_id": highlightID}, &resp)
	return resp.Stories, err
}

func (c *Client) GetHighlightStoryIds(ctx context.Context, pageID string) ([]string, error) {
	var resp struct {
		StoryIDs []string `json:"story_ids"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getHighlightStoryIds", map[string]string{"page_id": pageID}, &resp)
	return resp.StoryIDs, err
}

func (c *Client) GetMyArchiveStories(ctx context.Context) ([]Story, error) {
	var resp struct {
		Stories []Story `json:"stories"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getMyArchiveStories", nil, &resp)
	return resp.Stories, err
}

func (c *Client) GetNewEvents(ctx context.Context) ([]Event, error) {
	var resp struct {
		Events []Event `json:"events"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getNewEvents", nil, &resp)
	return resp.Events, err
}

func (c *Client) GetNewFollowRequests(ctx context.Context) ([]User, error) {
	var resp struct {
		Users []User `json:"users"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getNewFollowRequests", nil, &resp)
	return resp.Users, err
}

func (c *Client) GetPageFollowers(ctx context.Context, pageID string) ([]User, error) {
	var resp struct {
		Users []User `json:"users"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getPageFollowers", map[string]string{"page_id": pageID}, &resp)
	return resp.Users, err
}

func (c *Client) GetPageFollowing(ctx context.Context, pageID string) ([]Page, error) {
	var resp struct {
		Pages []Page `json:"pages"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getPageFollowing", map[string]string{"page_id": pageID}, &resp)
	return resp.Pages, err
}

func (c *Client) GetPageHighlights(ctx context.Context, pageID string) ([]Highlight, error) {
	var resp struct {
		Highlights []Highlight `json:"highlights"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getPageHighlights", map[string]string{"page_id": pageID}, &resp)
	return resp.Highlights, err
}

func (c *Client) GetPostLikes(ctx context.Context, postID string) ([]User, error) {
	var resp struct {
		Users []User `json:"users"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getPostLikes", map[string]string{"post_id": postID}, &resp)
	return resp.Users, err
}

func (c *Client) GetProfilePosts(ctx context.Context, userID string) ([]Post, error) {
	var resp struct {
		Posts []Post `json:"posts"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getProfilePosts", map[string]string{"user_id": userID}, &resp)
	return resp.Posts, err
}

func (c *Client) GetRelatedExplorePost(ctx context.Context, postID string) ([]Post, error) {
	var resp struct {
		Posts []Post `json:"posts"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getRelatedExplorePost", map[string]string{"post_id": postID}, &resp)
	return resp.Posts, err
}

func (c *Client) GetStory(ctx context.Context, storyID string) (*Story, error) {
	var story Story
	err := c.makeRequest(ctx, http.MethodPost, "getStory", map[string]string{"story_id": storyID}, &story)
	return &story, err
}

func (c *Client) GetStoryViewers(ctx context.Context, storyID string) ([]User, error) {
	var resp struct {
		Users []User `json:"users"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getStoryViewers", map[string]string{"story_id": storyID}, &resp)
	return resp.Users, err
}

func (c *Client) GetSuggested(ctx context.Context) ([]Page, error) {
	var resp struct {
		Pages []Page `json:"pages"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getSuggested", nil, &resp)
	return resp.Pages, err
}

func (c *Client) GetTaggedPosts(ctx context.Context, userID string) ([]Post, error) {
	var resp struct {
		Posts []Post `json:"posts"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getTaggedPosts", map[string]string{"user_id": userID}, &resp)
	return resp.Posts, err
}

func (c *Client) GetAllProfiles(ctx context.Context) ([]Profile, error) {
	var resp struct {
		Profiles []Profile `json:"profiles"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getAllProfiles", nil, &resp)
	return resp.Profiles, err
}

func (c *Client) GetComments(ctx context.Context, postID string) ([]Comment, error) {
	var resp struct {
		Comments []Comment `json:"comments"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getComments", map[string]string{"post_id": postID}, &resp)
	return resp.Comments, err
}

func (c *Client) GetInfoPost(ctx context.Context, postID string) (*Post, error) {
	var post Post
	err := c.makeRequest(ctx, http.MethodPost, "getInfoPost", map[string]string{"post_id": postID}, &post)
	return &post, err
}

func (c *Client) GetLink(ctx context.Context) (string, error) {
	var resp struct {
		Link string `json:"link"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getLink", nil, &resp)
	return resp.Link, err
}

func (c *Client) GetLinkShare(ctx context.Context, postID string) (string, error) {
	var resp struct {
		Link string `json:"link"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getLinkShare", map[string]string{"post_id": postID}, &resp)
	return resp.Link, err
}

func (c *Client) GetMyInfo(ctx context.Context) (*User, error) {
	var user User
	err := c.makeRequest(ctx, http.MethodPost, "getMyInfo", nil, &user)
	return &user, err
}

func (c *Client) GetSavedPosts(ctx context.Context) ([]Post, error) {
	var resp struct {
		Posts []Post `json:"posts"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "getSavedPosts", nil, &resp)
	return resp.Posts, err
}

func (c *Client) GetVideoResolution(ctx context.Context, fileID string) (*VideoResolution, error) {
	var resolution VideoResolution
	err := c.makeRequest(ctx, http.MethodPost, "getVideoResolution", map[string]string{"file_id": fileID}, &resolution)
	return &resolution, err
}

func (c *Client) HighlightStory(ctx context.Context, storyID, highlightID string) error {
	return c.makeRequest(ctx, http.MethodPost, "highlightStory", map[string]interface{}{"story_id": storyID, "highlight_id": highlightID}, nil)
}

func (c *Client) IsExistUsername(ctx context.Context, username string) (bool, error) {
	var resp struct {
		Exists bool `json:"exists"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "isExistUsername", map[string]string{"username": username}, &resp)
	return resp.Exists, err
}

func (c *Client) LikeComment(ctx context.Context, commentID string) error {
	return c.makeRequest(ctx, http.MethodPost, "likeComment", map[string]string{"comment_id": commentID}, nil)
}

func (c *Client) RemoveStoryFromHighlight(ctx context.Context, storyID, highlightID string) error {
	return c.makeRequest(ctx, http.MethodPost, "removeStoryFromHighlight", map[string]interface{}{"story_id": storyID, "highlight_id": highlightID}, nil)
}

func (c *Client) ReportPage(ctx context.Context, pageID, reason string) error {
	return c.makeRequest(ctx, http.MethodPost, "reportPage", map[string]interface{}{"page_id": pageID, "reason": reason}, nil)
}

func (c *Client) ReportPost(ctx context.Context, postID, reason string) error {
	return c.makeRequest(ctx, http.MethodPost, "reportPost", map[string]interface{}{"post_id": postID, "reason": reason}, nil)
}

func (c *Client) RequestFollow(ctx context.Context, userID string) error {
	return c.makeRequest(ctx, http.MethodPost, "requestFollow", map[string]string{"user_id": userID}, nil)
}

func (c *Client) SavePost(ctx context.Context, postID string) error {
	return c.makeRequest(ctx, http.MethodPost, "savePost", map[string]string{"post_id": postID}, nil)
}

func (c *Client) SearchFollower(ctx context.Context, pageID, query string) ([]User, error) {
	var resp struct {
		Users []User `json:"users"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "searchFollower", map[string]interface{}{"page_id": pageID, "query": query}, &resp)
	return resp.Users, err
}

func (c *Client) SearchFollowing(ctx context.Context, pageID, query string) ([]Page, error) {
	var resp struct {
		Pages []Page `json:"pages"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "searchFollowing", map[string]interface{}{"page_id": pageID, "query": query}, &resp)
	return resp.Pages, err
}

func (c *Client) SearchHashtag(ctx context.Context, query string) ([]Hashtag, error) {
	var resp struct {
		Hashtags []Hashtag `json:"hashtags"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "searchHashtag", map[string]string{"query": query}, &resp)
	return resp.Hashtags, err
}

func (c *Client) SearchPage(ctx context.Context, query string) ([]Page, error) {
	var resp struct {
		Pages []Page `json:"pages"`
	}
	err := c.makeRequest(ctx, http.MethodPost, "searchPage", map[string]string{"query": query}, &resp)
	return resp.Pages, err
}

func (c *Client) SetPageStatus(ctx context.Context, pageID, status string) error {
	return c.makeRequest(ctx, http.MethodPost, "setPageStatus", map[string]interface{}{"page_id": pageID, "status": status}, nil)
}

func (c *Client) Unfollow(ctx context.Context, userID string) error {
	return c.makeRequest(ctx, http.MethodPost, "unfollow", map[string]string{"user_id": userID}, nil)
}

func (c *Client) UnblockPage(ctx context.Context, pageID string) error {
	return c.makeRequest(ctx, http.MethodPost, "unblockPage", map[string]string{"page_id": pageID}, nil)
}

func (c *Client) UnlikePost(ctx context.Context, postID string) error {
	return c.makeRequest(ctx, http.MethodPost, "unlikePost", map[string]string{"post_id": postID}, nil)
}

func (c *Client) UnlikeComment(ctx context.Context, commentID string) error {
	return c.makeRequest(ctx, http.MethodPost, "unlikeComment", map[string]string{"comment_id": commentID}, nil)
}

func (c *Client) UnsavePost(ctx context.Context, postID string) error {
	return c.makeRequest(ctx, http.MethodPost, "unsavePost", map[string]string{"post_id": postID}, nil)
}

// --- متدهای آپلود فایل ---

func (c *Client) UploadAvatar(ctx context.Context, filePath string) (*File, error) {
	return c.UploadFileDirectly(ctx, filePath, ImageType)
}

func (c *Client) UploadFileDirectly(ctx context.Context, filePath string, fileType FileTypeEnum) (*File, error) {
	uploadResp, err := c.RequestSendFile(ctx, fileType)
	if err != nil {
		return nil, fmt.Errorf("خطا در درخواست آدرس آپلود: %w", err)
	}

	fileResp, err := c.UploadFile(uploadResp.UploadURL, filePath)
	if err != nil {
		return nil, fmt.Errorf("خطا در آپلود فایل: %w", err)
	}

	fileInfo, err := c.GetFile(ctx, fileResp.FileID)
	if err != nil {
		return nil, fmt.Errorf("خطا در دریافت اطلاعات فایل آپلود شده: %w", err)
	}

	return &File{
		FileID:   fileResp.FileID,
		FileName: fileInfo.FileName,
		Size:     fileInfo.Size,
	}, nil
}

// --- متدهای دیگر ---

// GetUsername از اطلاعات کاربر یا چت، نام کاربری را استخراج می‌کند
func (c *Client) GetUsername(chat *Chat) string {
	if chat.Username != "" {
		return chat.Username
	}
	// اگر نام کاربری وجود نداشت، می‌توان از شناسه کاربر برای جستجو استفاده کرد
	// اما این کار یک درخواست اضافی است. در اینجا خالی برمی‌گردانیم.
	return ""
}
