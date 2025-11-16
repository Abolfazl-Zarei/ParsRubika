package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Abolfazl-Zarei/ParsRubika-bot-go" // مسیر import اصلاح شد
)

var (
	client       *ParsRubika.Client
	stateManager = ParsRubika.NewStateManager()
)

// generateMainMenuKeyboard کیبورد اصلی ربات را می‌سازد
func generateMainMenuKeyboard() *ParsRubika.Keypad {
	return &ParsRubika.Keypad{
		Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{
				{ID: "msg_actions", Type: ParsRubika.ButtonTypeSimple, ButtonText: "📝 عملیات پیام"},
				{ID: "file_actions", Type: ParsRubika.ButtonTypeSimple, ButtonText: "📎 عملیات فایل"},
			}},
			{Buttons: []ParsRubika.Button{
				{ID: "group_actions", Type: ParsRubika.ButtonTypeSimple, ButtonText: "👥 مدیریت گروه"},
				{ID: "info_actions", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ℹ️ اطلاعات"},
			}},
			{Buttons: []ParsRubika.Button{
				{ID: "button_types", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🎛️ نمایش انواع دکمه"},
				{ID: "unofficial_api", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🌐 API غیررسمی"},
			}},
		},
	}
}

// handleUpdate تابع اصلی پردازش آپدیت‌ها
func handleUpdate(ctx context.Context, update *ParsRubika.Update) error {
	if update.NewMessage == nil {
		return nil
	}

	chatID := update.ChatID
	messageText := update.NewMessage.Text
	senderID := update.NewMessage.SenderID

	log.Printf("پیام جدید از %s در چت %s: %s", senderID, chatID, messageText)

	// 1. مدیریت دستورات اصلی
	if messageText == "/start" {
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{
			ChatID:       chatID,
			Text:         "به ربات جامع ParsRubika خوش آمدید! یکی از گزینه‌ها را انتخاب کنید:",
			InlineKeypad: generateMainMenuKeyboard(),
		})
		return err
	}

	// 2. مدیریت کلیک دکمه‌های شیشه‌ای (Inline Keypad)
	if update.NewMessage.AuxData != nil && update.NewMessage.AuxData.ButtonID != nil {
		buttonID := *update.NewMessage.AuxData.ButtonID
		return handleInlineButtons(ctx, chatID, senderID, buttonID)
	}

	// 3. مدیریت پاسخ‌های متنی بر اساس وضعیت کاربر
	return handleTextResponses(ctx, chatID, senderID, messageText)
}

// handleInlineButtons کلیک دکمه‌های شیشه‌ای را مدیریت می‌کند
func handleInlineButtons(ctx context.Context, chatID, senderID, buttonID string) error {
	switch buttonID {
	case "back_to_main":
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{
			ChatID:       chatID,
			Text:         "به منوی اصلی بازگشتید:",
			InlineKeypad: generateMainMenuKeyboard(),
		})
		return err

	case "msg_actions":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{
				{ID: "send_text", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ارسال متن"},
				{ID: "send_poll", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ارسال نظرسنجی"},
			}},
			{Buttons: []ParsRubika.Button{
				{ID: "send_location", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ارسال موقعیت"},
				{ID: "send_contact", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ارسال مخاطب"},
			}},
			{Buttons: []ParsRubika.Button{
				{ID: "edit_message", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ویرایش پیام"},
				{ID: "delete_message", Type: ParsRubika.ButtonTypeSimple, ButtonText: "حذف پیام"},
			}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "عملیات پیام:", InlineKeypad: kb})
		return err

	case "send_text":
		stateManager.SetState(senderID, "action", "send_text")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "متن مورد نظر خود را برای ارسال بنویسید:"})
		return err

	case "send_poll":
		stateManager.SetState(senderID, "action", "send_poll_question")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "سوال نظرسنجی را بنویسید:"})
		return err

	case "edit_message":
		stateManager.SetState(senderID, "action", "edit_message_id")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "آیدی پیامی که می‌خواهید ویرایش کنید را ارسال کنید:"})
		return err

	case "delete_message":
		stateManager.SetState(senderID, "action", "delete_message_id")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "آیدی پیامی که می‌خواهید حذف کنید را ارسال کنید:"})
		return err

	case "file_actions":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{
				{ID: "upload_file", Type: ParsRubika.ButtonTypeSimple, ButtonText: "آپلود و ارسال فایل"},
				{ID: "get_file_info", Type: ParsRubika.ButtonTypeSimple, ButtonText: "دریافت اطلاعات فایل"},
			}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "عملیات فایل:", InlineKeypad: kb})
		return err

	case "upload_file":
		stateManager.SetState(senderID, "action", "upload_file")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "مسیر کامل فایل روی کامپیوتر خود را ارسال کنید (مثال: C:\\image.jpg):"})
		return err

	case "get_file_info":
		stateManager.SetState(senderID, "action", "get_file_info")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "file_id فایل مورد نظر را ارسال کنید:"})
		return err

	case "group_actions":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{
				{ID: "get_chat_info", Type: ParsRubika.ButtonTypeSimple, ButtonText: "دریافت اطلاعات چت"},
				{ID: "get_chat_member", Type: ParsRubika.ButtonTypeSimple, ButtonText: "اطلاعات عضو"},
			}},
			{Buttons: []ParsRubika.Button{
				{ID: "get_admins", Type: ParsRubika.ButtonTypeSimple, ButtonText: "لیست ادمین‌ها"},
				{ID: "member_count", Type: ParsRubika.ButtonTypeSimple, ButtonText: "تعداد اعضا"},
			}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "مدیریت گروه و کانال:", InlineKeypad: kb})
		return err

	case "get_chat_info":
		stateManager.SetState(senderID, "action", "get_chat_info")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "chat_id گروه یا کانال را ارسال کنید:"})
		return err
	case "get_chat_member":
		stateManager.SetState(senderID, "action", "get_chat_member")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "chat_id و سپس user_id را با یک فاصله ارسال کنید:\n`chat_id user_id`"})
		return err
	case "get_admins":
		stateManager.SetState(senderID, "action", "get_admins")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "chat_id گروه یا کانال را برای دریافت لیست ادمین‌ها ارسال کنید:"})
		return err
	case "member_count":
		stateManager.SetState(senderID, "action", "member_count")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "chat_id گروه یا کانال را برای دریافت تعداد اعضا ارسال کنید:"})
		return err

	case "info_actions":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{
				{ID: "get_me", Type: ParsRubika.ButtonTypeSimple, ButtonText: "GetMe"},
				{ID: "set_commands", Type: ParsRubika.ButtonTypeSimple, ButtonText: "تنظیم دستورات"},
			}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "دریافت اطلاعات:", InlineKeypad: kb})
		return err

	case "get_me":
		bot, err := client.GetMe(ctx)
		if err != nil {
			return err
		}
		text := fmt.Sprintf("اطلاعات بات:\nID: %s\nنام: %s\nیوزرنیم: %s", bot.BotID, bot.BotTitle, bot.Username)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
		return err

	case "set_commands":
		commands := []ParsRubika.BotCommand{
			{Command: "start", Description: "شروع ربات"},
			{Command: "help", Description: "راهنما"},
		}
		err := client.SetCommands(ctx, &ParsRubika.SetCommandsRequest{BotCommands: commands})
		if err != nil {
			return err
		}
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "دستورات با موفقیت تنظیم شدند.", InlineKeypad: generateMainMenuKeyboard()})
		return err

	case "button_types":
		// --- اصلاح شده: برای فیلد *string ابتدا یک متغیر ساخته و سپس آدرس آن داده می‌شود ---
		googleLink := "https://google.com"

		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{
				{ID: "btn_simple", Type: ParsRubika.ButtonTypeSimple, ButtonText: "Simple"},
				{ID: "btn_link", Type: ParsRubika.ButtonTypeLink, ButtonText: "Link (گوگل)", Url: &googleLink},
			}},
			{Buttons: []ParsRubika.Button{
				{ID: "btn_request_phone", Type: ParsRubika.ButtonTypeAskMyPhoneNumber, ButtonText: "درخواست شماره تلفن"},
				{ID: "btn_request_location", Type: ParsRubika.ButtonTypeAskMyLocation, ButtonText: "درخواست موقعیت"},
			}},
			{Buttons: []ParsRubika.Button{
				{ID: "btn_textbox", Type: ParsRubika.ButtonTypeTextbox, ButtonText: "مربع متنی"},
				{ID: "btn_gallery_image", Type: ParsRubika.ButtonTypeGalleryImage, ButtonText: "ارسال عکس از گالری"},
			}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "انواع دکمه‌ها را تست کنید:", InlineKeypad: kb})
		return err

	case "unofficial_api":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{
				{ID: "get_my_info", Type: ParsRubika.ButtonTypeSimple, ButtonText: "GetMyInfo"},
				{ID: "get_suggested", Type: ParsRubika.ButtonTypeSimple, ButtonText: "GetSuggested"},
			}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "API غیررسمی (صفحه و استوری):", InlineKeypad: kb})
		return err

	case "get_my_info":
		user, err := client.GetMyInfo(ctx)
		if err != nil {
			return err
		}
		text := fmt.Sprintf("اطلاعات شما:\nID: %s\nنام: %s %s\nیوزرنیم: %s", user.UserID, user.FirstName, user.LastName, user.Username)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
		return err

	case "get_suggested":
		pages, err := client.GetSuggested(ctx)
		if err != nil {
			return err
		}
		text := "صفحات پیشنهادی:\n"
		for _, page := range pages {
			text += fmt.Sprintf("- %s (%s)\n", page.Title, page.Username)
		}
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
		return err
	}
	return nil
}

// handleTextResponses پاسخ‌های متنی کاربر را بر اساس وضعیت فعلی مدیریت می‌کند
func handleTextResponses(ctx context.Context, chatID, senderID, messageText string) error {
	action, ok := stateManager.GetState(senderID, "action")
	if !ok {
		return nil // اگر وضعتی تعریف نشده بود، کاری نکن
	}

	switch action {
	case "send_text":
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "پیام شما: " + messageText, InlineKeypad: generateMainMenuKeyboard()})
		stateManager.DeleteState(senderID, "action")
		return err

	case "send_poll_question":
		stateManager.SetState(senderID, "action", "send_poll_options")
		stateManager.SetState(senderID, "poll_question", messageText)
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "گزینه‌های نظرسنجی را با کاما (,) جدا کنید:"})
		return err
	case "send_poll_options":
		question, _ := stateManager.GetState(senderID, "poll_question")
		options := strings.Split(messageText, ",")
		_, err := client.SendPoll(ctx, &ParsRubika.SendPollRequest{ChatID: chatID, Question: question.(string), Options: options})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "نظرسنجی ارسال شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
		stateManager.DeleteUserState(senderID)
		return err

	case "edit_message_id":
		stateManager.SetState(senderID, "action", "edit_message_text")
		stateManager.SetState(senderID, "edit_message_id", messageText)
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "متن جدید برای پیام را بنویسید:"})
		return err
	case "edit_message_text":
		msgID, _ := stateManager.GetState(senderID, "edit_message_id")
		err := client.EditMessageText(ctx, &ParsRubika.EditMessageTextRequest{ChatID: chatID, MessageID: msgID.(string), Text: messageText})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "پیام ویرایش شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
		stateManager.DeleteUserState(senderID)
		return err

	case "delete_message_id":
		err := client.DeleteMessage(ctx, &ParsRubika.DeleteMessageRequest{ChatID: chatID, MessageID: messageText})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "پیام حذف شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
		stateManager.DeleteState(senderID, "action")
		return err

	case "upload_file":
		file, err := client.UploadFileDirectly(ctx, messageText, ParsRubika.ImageType)
		if err != nil {
			return fmt.Errorf("خطا در آپلود: %w", err)
		}
		_, err = client.SendFile(ctx, &ParsRubika.SendFileRequest{ChatID: chatID, FileID: file.FileID, Text: "فایل آپلود شد."})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "فایل با موفقیت آپلود و ارسال شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
		stateManager.DeleteState(senderID, "action")
		return err

	case "get_file_info":
		fileInfo, err := client.GetFile(ctx, messageText)
		if err != nil {
			return err
		}
		text := fmt.Sprintf("اطلاعات فایل:\nنام: %s\nحجم: %s\nلینک دانلود: %s", fileInfo.FileName, fileInfo.Size, fileInfo.DownloadURL)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
		stateManager.DeleteState(senderID, "action")
		return err

	case "get_chat_info":
		chat, err := client.GetChat(ctx, messageText)
		if err != nil {
			return err
		}
		text := fmt.Sprintf("اطلاعات چت:\nID: %s\nنوع: %s\nنام: %s", chat.ChatID, chat.ChatType, chat.Title)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
		stateManager.DeleteState(senderID, "action")
		return err

	case "get_chat_member":
		parts := strings.Split(messageText, " ")
		if len(parts) < 2 {
			_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "فرمت اشتباه است. لطفاً `chat_id user_id` را وارد کنید."})
			return err
		}
		member, err := client.GetChatMember(ctx, &ParsRubika.GetChatMemberRequest{ChatID: parts[0], UserID: parts[1]})
		if err != nil {
			return err
		}
		text := fmt.Sprintf("اطلاعات عضو:\nنام: %s %s\nوضعیت: %s", member.Member.User.FirstName, member.Member.User.LastName, member.Member.Status)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
		stateManager.DeleteState(senderID, "action")
		return err

	case "get_admins":
		admins, err := client.GetChatAdministrators(ctx, &ParsRubika.GetChatAdministratorsRequest{ChatID: messageText})
		if err != nil {
			return err
		}
		text := "لیست ادمین‌ها:\n"
		for _, admin := range admins.Administrators {
			text += fmt.Sprintf("- %s %s (%s)\n", admin.User.FirstName, admin.User.LastName, admin.User.Username)
		}
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
		stateManager.DeleteState(senderID, "action")
		return err

	case "member_count":
		count, err := client.GetChatMemberCount(ctx, &ParsRubika.GetChatMemberCountRequest{ChatID: messageText})
		if err != nil {
			return err
		}
		text := fmt.Sprintf("تعداد اعضای چت: %d", count.Count)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
		stateManager.DeleteState(senderID, "action")
		return err
	}

	return nil
}

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("متغیر محیطی BOT_TOKEN تنظیم نشده است.")
	}

	mode := flag.String("mode", "polling", "حالت اجرا: 'polling' یا 'webhook'")
	flag.Parse()

	client = ParsRubika.NewClient(botToken)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("سیگنال توقف دریافت شد. در حال بستن برنامه...")
		cancel()
	}()

	switch *mode {
	case "polling":
		log.Println("ربات جامع در حالت Polling شروع به کار کرد...")
		pollingOpts := ParsRubika.PollingOptions{
			Handler:      handleUpdate,
			RetryTimeout: 5 * time.Second,
			Limit:        100,
			PollInterval: 3 * time.Second,
		}
		if err := client.StartPolling(ctx, pollingOpts); err != nil {
			log.Fatalf("خطا در Polling: %v", err)
		}
	case "webhook":
		log.Println("حالت وب‌هوک در این مثال پیاده‌سازی نشده است.")
		// می‌توانید منطق وب‌هوک را مانند مثال قبلی اضافه کنید
	}

	log.Println("برنامه با موفقیت متوقف شد.")
}
