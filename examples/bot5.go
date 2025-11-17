package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Abolfazl-Zarei/ParsRubika-bot-go"
)

// GroupState وضعیت هر گروه را در حافظه نگه می‌دارد
type GroupState struct {
	Locks      map[string]bool
	MutedUsers map[string]bool
	Owners     map[string]bool
	Admins     map[string]bool
}

// نگهداری وضعیت تمام گروه‌ها به صورت thread-safe
var groupStates sync.Map

func main() {
	botToken := os.Getenv("RUBIKA_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("توکن ربات را در متغیر محیطی RUBIKA_BOT_TOKEN قرار دهید")
	}

	client := ParsRubika.NewClient(botToken)
	var messageCache sync.Map

	// برای تولید اعداد تصادفی
	rand.Seed(time.Now().UnixNano())

	handler := func(ctx context.Context, update *ParsRubika.Update) error {
		if update.NewMessage == nil {
			return nil
		}

		message := update.NewMessage
		chatID := update.ChatID
		text := message.Text
		senderID := message.SenderID

		// ذخیره پیام در حافظه موقت برای دستورات ریپلای
		if message.MessageID != "" {
			messageCache.Store(message.MessageID, senderID)
		}

		// دریافت وضعیت فعلی گروه و تعیین مالک اولیه در صورت لزوم
		state := getGroupState(chatID, senderID)

		// --- لایه اول: بررسی قفل‌ها و سکوت قبل از هر چیز ---
		// اگر کاربر ساکت بود، پیامش را حذف کن
		if state.MutedUsers[senderID] {
			log.Printf("کاربر %s ساکت است، پیام حذف شد.", senderID)
			return client.DeleteMessage(ctx, &ParsRubika.DeleteMessageRequest{ChatID: chatID, MessageID: message.MessageID})
		}

		// اگر پیام دستور نبود، قفل‌ها را بررسی کن
		if !strings.HasPrefix(text, "/") {
			if err := enforceLocks(ctx, client, chatID, message, state); err != nil {
				log.Printf("خطا در اجرای قفل‌ها: %v", err)
			}
			return nil // پیام غیردستوری پردازش نمی‌شود
		}

		// --- لایه دوم: بررسی دستور هوش مصنوعی ---
		if strings.HasPrefix(text, "+") {
			return handleAIChat(ctx, client, chatID, text, senderID)
		}

		// --- لایه سوم: پردازش دستورات اصلی ---
		command := strings.TrimPrefix(text, "/")
		log.Printf("دستور جدید دریافت شد: %s از کاربر %s در چت %s", command, senderID, chatID)

		// دستوراتی که نیاز به مالک بودن دارند (دستورات قدرتمند API)
		if isOwner(state, senderID) {
			switch {
			case strings.HasPrefix(command, "قفل "):
				return handleLockCommand(ctx, client, chatID, state, strings.TrimPrefix(command, "قفل "), true)
			case strings.HasPrefix(command, "باز کردن "):
				return handleLockCommand(ctx, client, chatID, state, strings.TrimPrefix(command, "باز کردن "), false)
			case strings.HasPrefix(command, "افزودن مالک "):
				return handleManageUserCommand(ctx, client, chatID, state, strings.TrimPrefix(command, "افزودن مالک "), "owner", true)
			case strings.HasPrefix(command, "حذف مالک "):
				return handleManageUserCommand(ctx, client, chatID, state, strings.TrimPrefix(command, "حذف مالک "), "owner", false)
			case strings.HasPrefix(command, "افزودن ادمین "):
				return handleManageUserCommand(ctx, client, chatID, state, strings.TrimPrefix(command, "افزودن ادمین "), "admin", true)
			case strings.HasPrefix(command, "حذف ادمین "):
				return handleManageUserCommand(ctx, client, chatID, state, strings.TrimPrefix(command, "حذف ادمین "), "admin", false)
			case command == "بن":
				return handleBan(ctx, client, chatID, message, &messageCache)
			case command == "آزاد":
				return handleUnban(ctx, client, chatID, message, &messageCache)
			case command == "ادمین":
				return handlePromote(ctx, client, chatID, message, &messageCache)
			}
		}

		// دستوراتی که نیاز به ادمین بودن (ربات یا روبیکا) دارند
		if isAdmin(state, senderID) {
			switch command {
			case "سکوت":
				return handleMuteCommand(ctx, client, chatID, message, &messageCache, state, true)
			case "رفع سکوت":
				return handleMuteCommand(ctx, client, chatID, message, &messageCache, state, false)
			case "پاکسازی لیست سکوت":
				return handleClearMutedCommand(ctx, client, chatID, state)
			case "پین":
				return handlePin(ctx, client, chatID, message)
			case "حذف_پین":
				return handleUnpin(ctx, client, chatID, message)
			}
		}

		// دستورات عمومی (برای همه اعضا)
		switch command {
		case "راهنما":
			return handleHelpCommand(ctx, client, chatID)
		case "وضعیت":
			return handleStatusCommand(ctx, client, chatID, state)
		case "شناسه من":
			return handleMyIDCommand(ctx, client, chatID, senderID)
		case "جوک":
			return handleJokeCommand(ctx, client, chatID)
		case "چالش":
			return handleChallengeCommand(ctx, client, chatID)
		case "انگیزشی":
			return handleMotivationalCommand(ctx, client, chatID)
		case "داستان":
			return handleStoryCommand(ctx, client, chatID)
		case "دانستنی":
			return handleFactCommand(ctx, client, chatID)
		case "فال حافظ":
			return handleFalCommand(ctx, client, chatID)
		case "نرخ ارز":
			return handleCurrencyCommand(ctx, client, chatID)
		case "نرخ طلا":
			return handleGoldCommand(ctx, client, chatID)
		case "بیوگرافی":
			return handleBioCommand(ctx, client, chatID)
		case "تاس":
			return handleDiceCommand(ctx, client, chatID)
		case "اطلاعات":
			return handleGetChat(ctx, client, chatID)
		case "تعداد":
			return handleMemberCount(ctx, client, chatID)
		case "مدیران":
			return handleGetAdmins(ctx, client, chatID)
		}

		return nil
	}

	pollingOpts := ParsRubika.PollingOptions{
		Handler:           handler,
		RetryTimeout:      5 * time.Second,
		Limit:             100,
		AllowEmptyUpdates: false,
		PollInterval:      1 * time.Second,
	}

	log.Println("🤖 ربات مدیریت گروه پیشرفته با موفقیت شروع به کار کرد")
	if err := client.StartPolling(context.Background(), pollingOpts); err != nil {
		log.Fatalf("خطا در اجرای ربات: %v", err)
	}
}

// --- توابع کمکی مدیریت وضعیت ---

func getGroupState(chatID string, senderID string) *GroupState {
	if state, ok := groupStates.Load(chatID); ok {
		return state.(*GroupState)
	}
	state := &GroupState{
		Locks:      make(map[string]bool),
		MutedUsers: make(map[string]bool),
		Owners:     make(map[string]bool),
		Admins:     make(map[string]bool),
	}
	// اولین نفری که در گروه جدید دستوری را اجرا می‌کند، به عنوان مالک اصلی تعیین می‌شود.
	if senderID != "" {
		state.Owners[senderID] = true
		log.Printf("کاربر %s به عنوان مالک اولیه گروه %s تعیین شد.", senderID, chatID)
	}
	groupStates.Store(chatID, state)
	return state
}

func isOwner(state *GroupState, userID string) bool {
	return state.Owners[userID]
}

func isAdmin(state *GroupState, userID string) bool {
	return isOwner(state, userID) || state.Admins[userID]
}

// --- توابع اصلی دستورات ---

func handleHelpCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	helpText := `💬 لیست دستورات و راهنما

🔒 **قفل‌ها (فقط مالک)**
● قفل لینک | باز کردن لینک
● قفل یوزرنیم | باز کردن یوزرنیم
● قفل فروارد | باز کردن فروارد
● قفل ریپلای | باز کردن ریپلای
● قفل فحش | باز کردن فحش
● قفل فایل | باز کردن فایل
● قفل هشتگ | باز کردن هشتگ
● قفل سخنگو | باز کردن سخنگو
● قفل دستورات عمومی | باز کردن دستورات عمومی
● قفل هایپرلینک | باز کردن هایپرلینک

👑 **مدیریت کاربران (فقط مالک)**
● بن (ریپلای روی کاربر)
● آزاد (ریپلای روی کاربر)
● ادمین (ریپلای روی کاربر)
● افزودن مالک <شناسه کاربر>
● حذف مالک <شناسه کاربر>
● افزودن ادمین <شناسه کاربر>
● حذف ادمین <شناسه کاربر>

🔇 **مدیریت سکوت (فقط ادمین‌های ربات)**
• سکوت (ریپلای روی کاربر)
• رفع سکوت (ریپلای روی کاربر)
• پاکسازی لیست سکوت

📌 **دستورات عمومی و اطلاعاتی**
● وضعیت | شناسه من | اطلاعات | تعداد | مدیران
● جوک | چالش | انگیزشی | داستان | دانستنی
● فال حافظ | نرخ ارز | نرخ طلا | بیوگرافی | تاس

🤖 **هوش مصنوعی**
برای گفتگو با هوش مصنوعی ابتدای پیام خود '+' قرار دهید.`
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: helpText})
	return err
}

func handleStatusCommand(ctx context.Context, client *ParsRubika.Client, chatID string, state *GroupState) error {
	var status strings.Builder
	status.WriteString("📊 **وضعیت فعلی گروه:**\n\n")

	status.WriteString("🔒 **قفل‌ها:**\n")
	for lockType, isLocked := range state.Locks {
		statusStr := "✅ باز"
		if isLocked {
			statusStr = "🔒 قفل"
		}
		status.WriteString(fmt.Sprintf("• %s: %s\n", lockType, statusStr))
	}

	status.WriteString("\n🔇 **لیست سکوت:**\n")
	if len(state.MutedUsers) == 0 {
		status.WriteString("• کاربر ساکت شده‌ای وجود ندارد.\n")
	} else {
		for userID := range state.MutedUsers {
			status.WriteString(fmt.Sprintf("• %s\n", userID))
		}
	}

	status.WriteString("\n👑 **مالکان (ربات):**\n")
	if len(state.Owners) == 0 {
		status.WriteString("• مالکی تعریف نشده است.\n")
	} else {
		for userID := range state.Owners {
			status.WriteString(fmt.Sprintf("• %s\n", userID))
		}
	}

	status.WriteString("\n🛡️ **ادمین‌های ربات:**\n")
	if len(state.Admins) == 0 {
		status.WriteString("• ادمینی تعریف نشده است.\n")
	} else {
		for userID := range state.Admins {
			status.WriteString(fmt.Sprintf("• %s\n", userID))
		}
	}

	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: status.String()})
	return err
}

// --- توابع API مدیریتی (از نسخه اول) ---

func handleBan(ctx context.Context, client *ParsRubika.Client, chatID string, message *ParsRubika.Message, cache *sync.Map) error {
	targetUserID, err := getUserIDFromReply(message, cache)
	if err != nil {
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: err.Error()})
		return err
	}
	err = client.BanChatMember(ctx, &ParsRubika.BanChatMemberRequest{ChatID: chatID, UserID: targetUserID})
	if err != nil {
		log.Printf("خطا در بن کردن کاربر: %v", err)
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("خطا در بن کردن کاربر: %v", err)})
		return err
	}
	_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "✅ کاربر با موفقیت بن شد."})
	return err
}

func handleUnban(ctx context.Context, client *ParsRubika.Client, chatID string, message *ParsRubika.Message, cache *sync.Map) error {
	targetUserID, err := getUserIDFromReply(message, cache)
	if err != nil {
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: err.Error()})
		return err
	}
	err = client.UnbanChatMember(ctx, &ParsRubika.UnbanChatMemberRequest{ChatID: chatID, UserID: targetUserID})
	if err != nil {
		log.Printf("خطا در آزاد کردن کاربر: %v", err)
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("خطا در آزاد کردن کاربر: %v", err)})
		return err
	}
	_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "✅ کاربر با موفقیت آزاد شد."})
	return err
}

func handlePromote(ctx context.Context, client *ParsRubika.Client, chatID string, message *ParsRubika.Message, cache *sync.Map) error {
	targetUserID, err := getUserIDFromReply(message, cache)
	if err != nil {
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: err.Error()})
		return err
	}
	isAdmin := true
	err = client.PromoteChatMember(ctx, &ParsRubika.PromoteChatMemberRequest{ChatID: chatID, UserID: targetUserID, IsAdministrator: &isAdmin})
	if err != nil {
		log.Printf("خطا در ترفیع کاربر: %v", err)
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("خطا در ترفیع کاربر: %v", err)})
		return err
	}
	_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "✅ کاربر با موفقیت به ادمین ترفیع یافت."})
	return err
}

func handlePin(ctx context.Context, client *ParsRubika.Client, chatID string, message *ParsRubika.Message) error {
	if message.ReplyToMessageID == "" {
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "❌ لطفاً به پیامی که می‌خواهید پین کنید پاسخ دهید."})
		return err
	}
	err := client.PinChatMessage(ctx, &ParsRubika.PinChatMessageRequest{ChatID: chatID, MessageID: message.ReplyToMessageID})
	if err != nil {
		log.Printf("خطا در پین کردن پیام: %v", err)
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("خطا در پین کردن پیام: %v", err)})
		return err
	}
	_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "✅ پیام با موفقیت پین شد."})
	return err
}

func handleUnpin(ctx context.Context, client *ParsRubika.Client, chatID string, message *ParsRubika.Message) error {
	var err error
	if message.ReplyToMessageID != "" {
		err = client.UnpinChatMessage(ctx, &ParsRubika.UnpinChatMessageRequest{ChatID: chatID, MessageID: message.ReplyToMessageID})
	} else {
		err = client.UnpinAllChatMessages(ctx, chatID)
	}
	if err != nil {
		log.Printf("خطا در حذف پین: %v", err)
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("خطا در حذف پین: %v", err)})
		return err
	}
	_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "✅ عملیات حذف پین با موفقیت انجام شد."})
	return err
}

func handleGetChat(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	chat, err := client.GetChat(ctx, chatID)
	if err != nil {
		return err
	}
	infoText := fmt.Sprintf(`اطلاعات گروه:
عنوان: %s
نوع: %s
شناسه: %s`, chat.Title, chat.ChatType, chat.ChatID)
	_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: infoText})
	return err
}

func handleMemberCount(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	count, err := client.GetChatMemberCount(ctx, &ParsRubika.GetChatMemberCountRequest{ChatID: chatID})
	if err != nil {
		log.Printf("خطا در دریافت تعداد اعضا: %v", err)
		_, sendErr := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("خطا در دریافت تعداد اعضا. آیا ربات ادمین است؟\nخطای API: %v", err)})
		return sendErr
	}
	_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("تعداد کل اعضای گروه: %d", count.Count)})
	return err
}

func handleGetAdmins(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	admins, err := client.GetChatAdministrators(ctx, &ParsRubika.GetChatAdministratorsRequest{ChatID: chatID})
	if err != nil {
		log.Printf("خطا در دریافت لیست مدیران: %v", err)
		_, sendErr := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("خطا در دریافت لیست مدیران. آیا ربات ادمین است؟\nخطای API: %v", err)})
		return sendErr
	}
	if len(admins.Administrators) == 0 {
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "هیچ مدیری در این گروه یافت نشد (یا ربات دسترسی لازم را ندارد)."})
		return err
	}
	var adminList strings.Builder
	adminList.WriteString("مدیران گروه (از طریق API):\n")
	for _, admin := range admins.Administrators {
		name := admin.User.FirstName
		if admin.User.LastName != "" {
			name += " " + admin.User.LastName
		}
		if admin.User.Username != "" {
			name += " (@" + admin.User.Username + ")"
		}
		adminList.WriteString(fmt.Sprintf("- %s\n", name))
	}
	_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: adminList.String()})
	return err
}

// --- بقیه توابع (بدون تغییر از نسخه پیشرفته) ---

func handleMyIDCommand(ctx context.Context, client *ParsRubika.Client, chatID, userID string) error {
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("شناسه شما:\n`%s`", userID)})
	return err
}

func handleLockCommand(ctx context.Context, client *ParsRubika.Client, chatID string, state *GroupState, lockType string, isLocked bool) error {
	state.Locks[lockType] = isLocked
	action := "قفل"
	if !isLocked {
		action = "باز کردن"
	}
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("✅ %s با موفقیت %s شد.", lockType, action)})
	return err
}

func handleManageUserCommand(ctx context.Context, client *ParsRubika.Client, chatID string, state *GroupState, targetID, role string, add bool) error {
	targetID = strings.TrimSpace(targetID)
	if targetID == "" {
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "❌ شناسه کاربر ارسال نشده است."})
		return err
	}
	var targetMap map[string]bool
	var roleTitle string
	if role == "owner" {
		targetMap = state.Owners
		roleTitle = "مالک"
	} else {
		targetMap = state.Admins
		roleTitle = "ادمین"
	}
	action := "حذف"
	if add {
		action = "افزودن"
		targetMap[targetID] = true
	} else {
		delete(targetMap, targetID)
	}
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("✅ کاربر `%s` با موفقیت از لیست %s‌های ربات %s شد.", targetID, roleTitle, action)})
	return err
}

func handleMuteCommand(ctx context.Context, client *ParsRubika.Client, chatID string, message *ParsRubika.Message, cache *sync.Map, state *GroupState, mute bool) error {
	targetUserID, err := getUserIDFromReply(message, cache)
	if err != nil {
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: err.Error()})
		return err
	}
	action := "ساکت"
	if !mute {
		action = "رفع سکوت"
		delete(state.MutedUsers, targetUserID)
	} else {
		state.MutedUsers[targetUserID] = true
	}
	_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("✅ کاربر با موفقیت %s شد.", action)})
	return err
}

func handleClearMutedCommand(ctx context.Context, client *ParsRubika.Client, chatID string, state *GroupState) error {
	state.MutedUsers = make(map[string]bool)
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "✅ لیست سکوت با موفقیت پاکسازی شد."})
	return err
}

func handleAIChat(ctx context.Context, client *ParsRubika.Client, chatID, text, userID string) error {
	aiResponse := fmt.Sprintf("🤖 پاسخ هوش مصنوعی به `%s`:\nشما گفتید: %s", userID, strings.TrimPrefix(text, "+"))
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: aiResponse})
	return err
}

// --- توابع سرگرمی و کاربردی (با داده‌های نمونه) ---

func handleJokeCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	jokes := []string{"چرا ریاضی دان به روانشناس مراجعه کرد؟ چون فکر می‌کرد جمع و تفریق داره!", "یه روز رفتم بیمارستان، دکتر گفت: بیماریت خطرناکه. گفتم: باشه، حالا خوب میشم. گفت: نه، منو می‌گم!"}
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: jokes[rand.Intn(len(jokes))]})
	return err
}

func handleChallengeCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	challenges := []string{"چالش: ۱ دقیقه بدون چشمک زدن!", "چالش: یک عکس از غذات بفرست و بگو اسمش چیه!", "چالش: بهترین خاطره‌ات رو در ۳ خط بنویس!"}
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: challenges[rand.Intn(len(challenges))]})
	return err
}

func handleMotivationalCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	quotes := []string{"موفقیت، یعنی از نقطه‌ای که شکست خورده‌ای، دوباره شروع کنی.", "فقط کسانی که دیوانه‌وار تلاش می‌کنند، می‌توانند به موفقیت‌های دیوانه‌وار دست پیدا کنند."}
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: quotes[rand.Intn(len(quotes))]})
	return err
}

func handleStoryCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	stories := []string{"داستان کوتاه: روزی روزگاری، کسی بود که نبود... و کسی نبود که بود... و تمام!", "داستان کوتاه: یک برنامه‌نویس به همسرش گفت: عشق من بی‌نهایت مثل یک حلقه بی‌نهایت است که هیچ وقت تموم نمیشه... تا اینکه سیستم کرش کرد."}
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: stories[rand.Intn(len(stories))]})
	return err
}

func handleFactCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	facts := []string{"آیا می‌دانستید قلب ماهی مرکب سه قلب دارد؟", "آیا می‌دانستید یک گروه قو را 'بادی' (bevy) می‌نامند؟"}
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: facts[rand.Intn(len(facts))]})
	return err
}

func handleFalCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	omens := []string{"فال شما:ال بی‌تو چون دیده نباشد چشم / جان بی‌تو چون نومیده نباشد دم", "فال شما:دل خوش باش زانکه جهان بقاست / خرم آن دل که ز جهان بی‌خبر است"}
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: omens[rand.Intn(len(omens))]})
	return err
}

func handleCurrencyCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "📈 نرخ ارز (نمونه):\nدلار: ۵۵,۰۰۰ تومان\nیورو: ۶۰,۰۰۰ تومان"})
	return err
}

func handleGoldCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "🥇 نرخ طلا (نمونه):\nمثقال طلا: ۱,۳۰۰,۰۰۰ تومان\nگرم طلا ۱۸: ۳۰۰,۰۰۰ تومان"})
	return err
}

func handleBioCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	bios := []string{"بیوگرافی: استیو جابز، بنیان‌گذار اپل، نمادی از نوآوری و تفکر خارج از چارچوب.", "بیوگرافی: ماری کوری، فیزیکدان و شیمی‌دان برنده دو جایزه نوبل، پیشگام در تحقیقات رادیواکتیو."}
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: bios[rand.Intn(len(bios))]})
	return err
}

func handleDiceCommand(ctx context.Context, client *ParsRubika.Client, chatID string) error {
	diceRoll := rand.Intn(6) + 1
	_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: fmt.Sprintf("🎲 تاس شما: %d", diceRoll)})
	return err
}

// --- توابع کمکی دیگر ---

func getUserIDFromReply(message *ParsRubika.Message, cache *sync.Map) (string, error) {
	if message.ReplyToMessageID == "" {
		return "", fmt.Errorf("❌ لطفاً به پیام کاربر مورد نظر ریپلای کنید.")
	}
	userIDInterface, ok := cache.Load(message.ReplyToMessageID)
	if !ok {
		return "", fmt.Errorf("❌ پیام اصلی پیدا نشد. لطفاً دوباره تلاش کنید.")
	}
	return userIDInterface.(string), nil
}

func enforceLocks(ctx context.Context, client *ParsRubika.Client, chatID string, message *ParsRubika.Message, state *GroupState) error {
	text := message.Text
	var shouldDelete bool
	var reason string
	if state.Locks["لینک"] && strings.Contains(text, "http") {
		shouldDelete = true
		reason = "ارسال لینک ممنوع است."
	}
	if state.Locks["یوزرنیم"] && strings.Contains(text, "@") {
		shouldDelete = true
		reason = "ارسال یوزرنیم ممنوع است."
	}
	if state.Locks["هشتگ"] && strings.Contains(text, "#") {
		shouldDelete = true
		reason = "ارسال هشتگ ممنوع است."
	}
	if state.Locks["فحش"] && containsBadWords(text) {
		shouldDelete = true
		reason = "استفاده از کلمات نامناسب ممنوع است."
	}
	if shouldDelete {
		log.Printf("پیام به دلیل قفل گروه حذف شد. دلیل: %s", reason)
		client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: reason, ReplyToMessageID: message.MessageID})
		return client.DeleteMessage(ctx, &ParsRubika.DeleteMessageRequest{ChatID: chatID, MessageID: message.MessageID})
	}
	return nil
}

func containsBadWords(text string) bool {
	badWords := []string{"توهین", "فحش"}
	lowerText := strings.ToLower(text)
	for _, word := range badWords {
		if strings.Contains(lowerText, word) {
			return true
		}
	}
	return false
}
