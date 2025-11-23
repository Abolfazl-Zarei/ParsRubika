package ParsRubika

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// WebhookOptions تنظیمات وب‌هوک
type WebhookOptions struct {
	Port    int
	Path    string
	Handler HandlerFunc
	Secret  string // برای احراز هویت
}

// StartWebhook راه‌اندازی سرور وب‌هوک
func (c *BotClient) StartWebhook(ctx context.Context, opts WebhookOptions) error {
	if opts.Handler == nil {
		opts.Handler = c.ProcessUpdate
	}

	mux := http.NewServeMux()
	mux.HandleFunc(opts.Path, c.webhookHandler(opts))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", opts.Port),
		Handler: mux,
	}

	// مدیریت graceful shutdown
	go func() {
		<-ctx.Done()
		log.Println("در حال بستن سرور وب‌هوک...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("خطا در بستن سرور وب‌هوک: %v", err)
		}
	}()

	log.Printf("سرور وب‌هوک روی پورت %d و مسیر %s راه‌اندازی شد", opts.Port, opts.Path)

	// ابتدا بات را شروع می‌کنیم تا اطمینان حاصل شود که آماده است
	if err := c.Start(ctx); err != nil {
		return err
	}

	return server.ListenAndServe()
}

// webhookHandler هندلر اصلی وب‌هوک
func (c *BotClient) webhookHandler(opts WebhookOptions) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// بررسی secret اگر تنظیم شده باشد
		if opts.Secret != "" {
			secret := r.Header.Get("X-Rubika-Secret")
			if secret != opts.Secret {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("خطا در خواندن بدنه درخواست: %v", err)
			http.Error(w, "خطا در خواندن بدنه", http.StatusBadRequest)
			return
		}

		update, err := c.parseWebhookBody(body)
		if err != nil {
			log.Printf("خطا در پارس کردن بدنه وب‌هوک: %v", err)
			http.Error(w, "فرمت بدنه نامعتبر", http.StatusBadRequest)
			return
		}

		// پردازش آپدیت با استفاده از هندلر تنظیم شده
		if err := opts.Handler(r.Context(), update); err != nil {
			log.Printf("خطا در پردازش آپدیت وب‌هوک: %v", err)
			http.Error(w, "خطا در پردازش", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	}
}

// parseWebhookBody پارس کردن بدنه وب‌هوک
func (c *BotClient) parseWebhookBody(body []byte) (*Update, error) {
	// تلاش برای receiveUpdate
	var updateWrapper struct {
		Update *Update `json:"update"`
	}
	if err := json.Unmarshal(body, &updateWrapper); err == nil && updateWrapper.Update != nil {
		return updateWrapper.Update, nil
	}

	// تلاش برای receiveInlineMessage
	var inlineWrapper struct {
		InlineMessage *InlineMessage `json:"inline_message"`
	}
	if err := json.Unmarshal(body, &inlineWrapper); err == nil && inlineWrapper.InlineMessage != nil {
		return c.convertInlineToUpdate(inlineWrapper.InlineMessage), nil
	}

	// تلاش برای receiveQuery
	var queryWrapper struct {
		QueryID  string `json:"query_id"`
		ChatID   string `json:"chat_id"`
		SenderID string `json:"sender_id"`
		Data     string `json:"data"`
	}
	if err := json.Unmarshal(body, &queryWrapper); err == nil && queryWrapper.QueryID != "" {
		return c.convertQueryToUpdate(&queryWrapper), nil
	}

	return nil, fmt.Errorf("نوع وب‌هوک شناخته نشد")
}

// convertInlineToUpdate تبدیل InlineMessage به Update
func (c *BotClient) convertInlineToUpdate(inlineMsg *InlineMessage) *Update {
	// تبدیل MessageID از string به int64 برای مطابقت با مدل Message
	messageID, _ := strconv.ParseInt(inlineMsg.MessageID, 10, 64)
	return &Update{
		Type:   NewMessage,
		ChatID: inlineMsg.ChatID,
		NewMessage: &Message{
			MessageID: messageID,
			Text:      inlineMsg.Text,
			// تبدیل زمان به string برای مطابقت با مدل Message
			Time:     strconv.FormatInt(time.Now().Unix(), 10),
			SenderID: inlineMsg.SenderID,
			AuxData:  inlineMsg.AuxData,
			File:     inlineMsg.File,
			Location: inlineMsg.Location,
		},
	}
}

// convertQueryToUpdate تبدیل Query به Update
func (c *BotClient) convertQueryToUpdate(query *struct {
	QueryID  string `json:"query_id"`
	ChatID   string `json:"chat_id"`
	SenderID string `json:"sender_id"`
	Data     string `json:"data"`
}) *Update {
	// تبدیل QueryID از string به int64 برای مطابقت با مدل Message
	messageID, _ := strconv.ParseInt(query.QueryID, 10, 64)
	return &Update{
		Type:   NewMessage,
		ChatID: query.ChatID,
		NewMessage: &Message{
			MessageID: messageID,
			Text:      query.Data,
			// تبدیل زمان به string برای مطابقت با مدل Message
			Time:     strconv.FormatInt(time.Now().Unix(), 10),
			SenderID: query.SenderID,
			AuxData:  &AuxData{ButtonID: &query.QueryID},
		},
	}
}
