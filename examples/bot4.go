package main

// سازنده ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

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

	"github.com/Abolfazl-Zarei/ParsRubika-bot-go"
)

var (
	client       *ParsRubika.Client
	stateManager = ParsRubika.NewStateManager()
)

// generateMainMenuKeyboard ساختار اصلی منوی ربات است
func generateMainMenuKeyboard() *ParsRubika.Keypad {
	return &ParsRubika.Keypad{
		Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{{ID: "menu_msg_ops", Type: ParsRubika.ButtonTypeSimple, ButtonText: "📝 عملیات پیام"}}},
			{Buttons: []ParsRubika.Button{{ID: "menu_file_ops", Type: ParsRubika.ButtonTypeSimple, ButtonText: "📎 عملیات فایل"}}},
			{Buttons: []ParsRubika.Button{{ID: "menu_group_ops", Type: ParsRubika.ButtonTypeSimple, ButtonText: "👥 مدیریت گروه/کانال"}}},
			{Buttons: []ParsRubika.Button{{ID: "menu_bot_settings", Type: ParsRubika.ButtonTypeSimple, ButtonText: "⚙️ تنظیمات بات"}}},
			{Buttons: []ParsRubika.Button{{ID: "menu_button_showcase", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🎛️ انواع کیبورد و دکمه"}}},
			{Buttons: []ParsRubika.Button{{ID: "menu_unofficial_api", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🌐 API غیررسمی (صفحه/استوری)"}}},
		},
	}
}

// handleUpdate نقطه ورود اصلی برای تمام آپدیت‌ها
func handleUpdate(ctx context.Context, update *ParsRubika.Update) error {
	if update.NewMessage == nil {
		return nil
	}

	chatID := update.ChatID
	messageText := update.NewMessage.Text
	senderID := update.NewMessage.SenderID

	log.Printf("[UPDATE] from %s in %s: %s", senderID, chatID, messageText)

	// 1. مدیریت دستورات اصلی
	if messageText == "/start" {
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{
			ChatID:       chatID,
			Text:         "🤖 به ربات جامع و آچار به دست ParsRubika خوش آمدید!\n\nاز منوی زیر یکی از گزینه‌ها را انتخاب کنید:",
			InlineKeypad: generateMainMenuKeyboard(),
		})
		return err
	}

	// 2. مدیریت کلیک دکمه‌های شیشه‌ای (Inline Keypad)
	if update.NewMessage.AuxData != nil && update.NewMessage.AuxData.ButtonID != nil {
		buttonID := *update.NewMessage.AuxData.ButtonID
		return handleInlineButtons(ctx, chatID, senderID, buttonID)
	}

	// 3. مدیریت پاسخ‌های متنی بر اساس وضعیت فعلی کاربر
	return handleTextResponses(ctx, chatID, senderID, messageText)
}

// handleInlineButtons کلیک روی دکمه‌های شیشه‌ای را مدیریت می‌کند
func handleInlineButtons(ctx context.Context, chatID, senderID, buttonID string) error {
	switch buttonID {
	case "back_to_main":
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "به منوی اصلی بازگشتید:", InlineKeypad: generateMainMenuKeyboard()})
		return err

	// --- منوی عملیات پیام ---
	case "menu_msg_ops":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{{ID: "act_send_text", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ارسال متن"}}},
			{Buttons: []ParsRubika.Button{{ID: "act_send_poll", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ارسال نظرسنجی"}}},
			{Buttons: []ParsRubika.Button{{ID: "act_forward_msg", Type: ParsRubika.ButtonTypeSimple, ButtonText: "فوروارد پیام"}}},
			{Buttons: []ParsRubika.Button{{ID: "act_edit_text", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ویرایش متن پیام"}}},
			{Buttons: []ParsRubika.Button{{ID: "act_delete_msg", Type: ParsRubika.ButtonTypeSimple, ButtonText: "حذف پیام"}}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "📝 عملیات پیام:", InlineKeypad: kb})
		return err
	case "act_send_text":
		stateManager.SetState(senderID, "action", "send_text")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "متن خود را برای ارسال بنویسید:"})
		return err
	case "act_send_poll":
		stateManager.SetState(senderID, "action", "poll_q")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "سوال نظرسنجی را بنویسید:"})
		return err
	case "act_forward_msg":
		stateManager.SetState(senderID, "action", "forward_from")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "از کدام چت؟ (chat_id)"})
		return err
	case "act_edit_text":
		stateManager.SetState(senderID, "action", "edit_id")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "آیدی پیام برای ویرایش:"})
		return err
	case "act_delete_msg":
		stateManager.SetState(senderID, "action", "delete_id")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "آیدی پیام برای حذف:"})
		return err

	// --- منوی عملیات فایل ---
	case "menu_file_ops":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{{ID: "act_upload_file", Type: ParsRubika.ButtonTypeSimple, ButtonText: "آپلود فایل"}}},
			{Buttons: []ParsRubika.Button{{ID: "act_get_file_info", Type: ParsRubika.ButtonTypeSimple, ButtonText: "دریافت اطلاعات فایل"}}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "📎 عملیات فایل:", InlineKeypad: kb})
		return err
	case "act_upload_file":
		stateManager.SetState(senderID, "action", "upload_path")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "مسیر فایل را بفرست (مثال: C:\\image.jpg):"})
		return err
	case "act_get_file_info":
		stateManager.SetState(senderID, "action", "get_file_id")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "file_id را بفرست:"})
		return err

	// --- منوی مدیریت گروه ---
	case "menu_group_ops":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{{ID: "act_get_chat", Type: ParsRubika.ButtonTypeSimple, ButtonText: "GetChat"}}},
			{Buttons: []ParsRubika.Button{{ID: "act_ban_member", Type: ParsRubika.ButtonTypeSimple, ButtonText: "BanMember"}}},
			{Buttons: []ParsRubika.Button{{ID: "act_promote_member", Type: ParsRubika.ButtonTypeSimple, ButtonText: "PromoteMember"}}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "👥 مدیریت گروه/کانال:", InlineKeypad: kb})
		return err
	case "act_get_chat":
		stateManager.SetState(senderID, "action", "get_chat_id")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "chat_id را بفرست:"})
		return err
	case "act_ban_member":
		stateManager.SetState(senderID, "action", "ban_chat_id")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "chat_id گروه را بفرست:"})
		return err
	case "act_promote_member":
		stateManager.SetState(senderID, "action", "promote_chat_id")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "chat_id گروه را بفرست:"})
		return err

	// --- منوی تنظیمات بات ---
	case "menu_bot_settings":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{{ID: "set_get_me", Type: ParsRubika.ButtonTypeSimple, ButtonText: "GetMe"}}},
			{Buttons: []ParsRubika.Button{{ID: "set_commands", Type: ParsRubika.ButtonTypeSimple, ButtonText: "SetCommands"}}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "⚙️ تنظیمات بات:", InlineKeypad: kb})
		return err
	case "set_get_me":
		bot, err := client.GetMe(ctx)
		if err != nil {
			return err
		}
		text := fmt.Sprintf("Bot Info: %s (@%s)", bot.BotTitle, bot.Username)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
		return err
	case "set_commands":
		commands := []ParsRubika.BotCommand{{Command: "start", Description: "شروع"}, {Command: "help", Description: "راهنما"}}
		err := client.SetCommands(ctx, &ParsRubika.SetCommandsRequest{BotCommands: commands})
		if err != nil {
			return err
		}
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "دستورات تنظیم شد.", InlineKeypad: generateMainMenuKeyboard()})
		return err

	// --- منوی نمایش انواع کیبورد و دکمه ---
	case "menu_button_showcase":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{{ID: "show_reply_kb", Type: ParsRubika.ButtonTypeSimple, ButtonText: "Reply Keyboard"}}},
			{Buttons: []ParsRubika.Button{{ID: "show_inline_kb", Type: ParsRubika.ButtonTypeSimple, ButtonText: "Advanced Inline Keyboard"}}},
			{Buttons: []ParsRubika.Button{{ID: "show_other_markups", Type: ParsRubika.ButtonTypeSimple, ButtonText: "Other Markups"}}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "🎛️ انتخاب نوع کیبورد:", InlineKeypad: kb})
		return err

	// --- Reply Keyboard ---
	case "show_reply_kb":
		kb := &ParsRubika.Keypad{
			Rows: []ParsRubika.KeypadRow{
				{Buttons: []ParsRubika.Button{{ID: "reply_btn_1", Type: ParsRubika.ButtonTypeSimple, ButtonText: "دکمه ساده ۱"}, {ID: "reply_btn_2", Type: ParsRubika.ButtonTypeSimple, ButtonText: "دکمه ساده ۲"}}},
				{Buttons: []ParsRubika.Button{{ID: "act_send_poll", Type: ParsRubika.ButtonTypeSimple, ButtonText: "📊 Request Poll"}}},
				{Buttons: []ParsRubika.Button{{ID: "btn_request_contact", Type: ParsRubika.ButtonTypeAskMyPhoneNumber, ButtonText: "📞 Request Contact"}}},
				{Buttons: []ParsRubika.Button{{ID: "btn_request_location", Type: ParsRubika.ButtonTypeAskMyLocation, ButtonText: "📍 Request Location"}}},
			},
			ResizeKeyboard: true,
		}
		err := client.EditChatKeypad(ctx, &ParsRubika.EditChatKeypadRequest{ChatID: chatID, ChatKeypad: kb, ChatKeypadType: ParsRubika.NewKeypad})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "Reply Keyboard فعال شد. روی دکمه‌ها کلیک کنید."})
		}
		return err
	case "reply_btn_1":
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "شما روی دکمه ساده ۱ کلیک کردید!", InlineKeypad: generateMainMenuKeyboard()})
		return err
	case "reply_btn_2":
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "شما روی دکمه ساده ۲ کلیک کردید!", InlineKeypad: generateMainMenuKeyboard()})
		return err

	// --- Advanced Inline Keyboard ---
	case "show_inline_kb":
		// --- اصلاح شده: برای فیلدهای *string و *bool ابتدا متغیر ساخته و سپس آدرس آن داده می‌شود ---
		googleLink := "https://google.com"
		switchInlineText := "test_query"
		loginForwardText := "ورود به سایت"
		isPaid := true
		var emptyGame = ""

		loginURL := &ParsRubika.LoginUrl{
			URL:         "https://example.com/auth",
			ForwardText: &loginForwardText,
		}
		webAppURL := &ParsRubika.WebAppInfo{URL: "https://example.com/webapp"}

		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{{ID: "btn_callback", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔘 Callback Button"}}},
			{Buttons: []ParsRubika.Button{{ID: "btn_url", Type: ParsRubika.ButtonTypeLink, ButtonText: "🔗 URL Button", Url: &googleLink}}},
			{Buttons: []ParsRubika.Button{{ID: "btn_switch_inline", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔄 Switch Inline", SwitchInlineQuery: &switchInlineText}}},
			{Buttons: []ParsRubika.Button{{ID: "btn_login", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔐 Login URL", LoginUrl: loginURL}}},
			{Buttons: []ParsRubika.Button{{ID: "btn_game", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🎮 Game Button", CallbackGame: &emptyGame}}},
			{Buttons: []ParsRubika.Button{{ID: "btn_pay", Type: ParsRubika.ButtonTypeSimple, ButtonText: "💳 Buy Button ($1)", Pay: &isPaid}}},
			{Buttons: []ParsRubika.Button{{ID: "btn_webapp", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🌐 Web App Button", WebApp: webAppURL}}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "نمایش انواع دکمه‌های Inline پیشرفته:", InlineKeypad: kb})
		return err
	case "btn_callback":
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "شما روی Callback Button کلیک کردید!", InlineKeypad: generateMainMenuKeyboard()})
		return err

	// --- Other Markups ---
	case "show_other_markups":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{{ID: "act_force_reply", Type: ParsRubika.ButtonTypeSimple, ButtonText: "Force Reply"}}},
			{Buttons: []ParsRubika.Button{{ID: "remove_reply_kb", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ReplyKeyboardRemove"}}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "سایر حالت‌های نشانه‌گذاری:", InlineKeypad: kb})
		return err
	case "act_force_reply":
		stateManager.SetState(senderID, "action", "force_reply")
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "لطفاً به این پیام پاسخ دهید:"})
		return err
	case "remove_reply_kb":
		err := client.EditChatKeypad(ctx, &ParsRubika.EditChatKeypadRequest{ChatID: chatID, ChatKeypadType: ParsRubika.RemoveKeypad})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "Reply Keyboard حذف شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
		return err

	// --- منوی API غیررسمی ---
	case "menu_unofficial_api":
		kb := &ParsRubika.Keypad{Rows: []ParsRubika.KeypadRow{
			{Buttons: []ParsRubika.Button{{ID: "unoff_get_my_info", Type: ParsRubika.ButtonTypeSimple, ButtonText: "GetMyInfo"}}},
			{Buttons: []ParsRubika.Button{{ID: "back_to_main", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🔙 بازگشت"}}},
		}}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "🌐 API غیررسمی (صفحه/استوری):\n⚠️ استفاده از این موارد با ریسک همراه است", InlineKeypad: kb})
		return err
	case "unoff_get_my_info":
		user, err := client.GetMyInfo(ctx)
		if err != nil {
			return err
		}
		text := fmt.Sprintf("Your Info:\nName: %s %s\nUsername: @%s", user.FirstName, user.LastName, user.Username)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
		return err
	}
	return nil
}

// handleTextResponses پاسخ‌های متنی کاربر را بر اساس وضعیت فعلی مدیریت می‌کند
func handleTextResponses(ctx context.Context, chatID, senderID, messageText string) error {
	action, ok := stateManager.GetState(senderID, "action")
	if !ok {
		return nil
	}
	var err error

	// --- اصلاح شده: استفاده از type assertion برای مقادیر interface{} ---
	switch action {
	case "send_text":
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: messageText, InlineKeypad: generateMainMenuKeyboard()})
	case "poll_q":
		stateManager.SetState(senderID, "action", "poll_o")
		stateManager.SetState(senderID, "poll_q", messageText)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "گزینه‌ها را با کاما جدا کنید:"})
	case "poll_o":
		question, _ := stateManager.GetState(senderID, "poll_q")
		opts := strings.Split(messageText, ",")
		_, err = client.SendPoll(ctx, &ParsRubika.SendPollRequest{ChatID: chatID, Question: question.(string), Options: opts})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "نظرسنجی ارسال شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
	case "forward_from":
		stateManager.SetState(senderID, "action", "forward_msg_id")
		stateManager.SetState(senderID, "forward_from", messageText)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "message_id را بفرست:"})
	case "forward_msg_id":
		from, _ := stateManager.GetState(senderID, "forward_from")
		_, err = client.ForwardMessage(ctx, &ParsRubika.ForwardMessageRequest{FromChatID: from.(string), MessageID: messageText, ToChatID: chatID})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "فوروارد شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
	case "edit_id":
		stateManager.SetState(senderID, "action", "edit_text")
		stateManager.SetState(senderID, "edit_id", messageText)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "متن جدید را بفرست:"})
	case "edit_text":
		id, _ := stateManager.GetState(senderID, "edit_id")
		err = client.EditMessageText(ctx, &ParsRubika.EditMessageTextRequest{ChatID: chatID, MessageID: id.(string), Text: messageText})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "ویرایش شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
	case "delete_id":
		err = client.DeleteMessage(ctx, &ParsRubika.DeleteMessageRequest{ChatID: chatID, MessageID: messageText})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "حذف شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
	case "upload_path":
		file, err := client.UploadFileDirectly(ctx, messageText, ParsRubika.ImageType)
		if err != nil {
			return fmt.Errorf("upload error: %w", err)
		}
		_, err = client.SendFile(ctx, &ParsRubika.SendFileRequest{ChatID: chatID, FileID: file.FileID, Text: "فایل آپلود شد."})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "موفقیت آمیز.", InlineKeypad: generateMainMenuKeyboard()})
		}
	case "get_file_id":
		info, err := client.GetFile(ctx, messageText)
		if err != nil {
			return err
		}
		text := fmt.Sprintf("File: %s\nSize: %s", info.FileName, info.Size)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
	case "get_chat_id":
		chat, err := client.GetChat(ctx, messageText)
		if err != nil {
			return err
		}
		text := fmt.Sprintf("Chat: %s (%s)", chat.Title, chat.ChatType)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: text, InlineKeypad: generateMainMenuKeyboard()})
	case "ban_chat_id":
		stateManager.SetState(senderID, "action", "ban_user_id")
		stateManager.SetState(senderID, "ban_chat_id", messageText)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "user_id برای بن کردن را بفرست:"})
	case "ban_user_id":
		chatIDFromState, _ := stateManager.GetState(senderID, "ban_chat_id")
		err = client.BanChatMember(ctx, &ParsRubika.BanChatMemberRequest{ChatID: chatIDFromState.(string), UserID: messageText})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "کاربر بن شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
	case "promote_chat_id":
		stateManager.SetState(senderID, "action", "promote_user_id")
		stateManager.SetState(senderID, "promote_chat_id", messageText)
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "user_id برای ادمین کردن را بفرست:"})
	case "promote_user_id":
		chatIDFromState, _ := stateManager.GetState(senderID, "promote_chat_id")
		isAdmin := true
		err = client.PromoteChatMember(ctx, &ParsRubika.PromoteChatMemberRequest{ChatID: chatIDFromState.(string), UserID: messageText, IsAdministrator: &isAdmin})
		if err == nil {
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "کاربر ادمین شد.", InlineKeypad: generateMainMenuKeyboard()})
		}
	case "force_reply":
		_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "پاسخ شما دریافت شد: " + messageText, InlineKeypad: generateMainMenuKeyboard()})
	}

	if err == nil {
		stateManager.DeleteUserState(senderID)
	}
	return err
}

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("BOT_TOKEN not set")
	}
	mode := flag.String("mode", "polling", "mode")
	flag.Parse()

	client = ParsRubika.NewClient(botToken)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() { <-c; log.Println("Shutting down..."); cancel() }()

	switch *mode {
	case "polling":
		log.Println("🚀 Ultimate Bot started in Polling mode...")
		pollOpts := ParsRubika.PollingOptions{Handler: handleUpdate, PollInterval: 3 * time.Second}
		if err := client.StartPolling(ctx, pollOpts); err != nil {
			log.Fatalf("Polling error: %v", err)
		}
	case "webhook":
		log.Println("Webhook mode not implemented in this example.")
	}
	log.Println("Bot stopped.")
}
