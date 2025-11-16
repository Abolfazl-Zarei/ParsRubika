package main

// سازنده ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Abolfazl-Zarei/ParsRubika-bot-go" // این خط را بررسی کنید
)

// متغیرهای سراسری برای دسترسی آسان در هندلر
var (
	client       *ParsRubika.Client
	stateManager = ParsRubika.NewStateManager()
)

// handleUpdate تابع اصلی پردازش آپدیت‌هاست که هم در polling و هم در webhook فراخوانی می‌شود
func handleUpdate(ctx context.Context, update *ParsRubika.Update) error {
	if update.NewMessage == nil {
		return nil
	}

	chatID := update.ChatID
	messageText := update.NewMessage.Text
	senderID := update.NewMessage.SenderID

	log.Printf("پیام جدید از %s در چت %s: %s", senderID, chatID, messageText)

	// 1. مدیریت دستورات
	if strings.HasPrefix(messageText, "/start") {
		welcomeText := "به ربات خوش آمدید! یکی از گزینه‌ها را انتخاب کنید:"
		inlineKeyboard := &ParsRubika.Keypad{
			Rows: []ParsRubika.KeypadRow{
				{
					Buttons: []ParsRubika.Button{
						{ID: "info_button", Type: ParsRubika.ButtonTypeSimple, ButtonText: "ℹ️ اطلاعات ربات"},
						{ID: "toggle_reply_keyboard", Type: ParsRubika.ButtonTypeSimple, ButtonText: "⌨️ تغییر کیبورد چت"},
					},
				},
				{
					Buttons: []ParsRubika.Button{
						{ID: "set_state_button", Type: ParsRubika.ButtonTypeSimple, ButtonText: "📝 تنظیم وضعیت"},
						{ID: "get_state_button", Type: ParsRubika.ButtonTypeSimple, ButtonText: "📖 خواندن وضعیت"},
						{ID: "clear_state_button", Type: ParsRubika.ButtonTypeSimple, ButtonText: "🗑️ پاک کردن وضعیت"},
					},
				},
			},
		}
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{
			ChatID:       chatID,
			Text:         welcomeText,
			InlineKeypad: inlineKeyboard,
		})
		return err
	}

	// 2. مدیریت کلیک دکمه‌های شیشه‌ای (Inline Keypad)
	if update.NewMessage.AuxData != nil && update.NewMessage.AuxData.ButtonID != nil {
		buttonID := *update.NewMessage.AuxData.ButtonID
		var err error

		switch buttonID {
		case "info_button":
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{
				ChatID: chatID,
				Text:   "این یک ربات نمونه برای کتابخانه ParsRubika است.",
			})
		case "toggle_reply_keyboard":
			// بررسی وضعیت فعلی کیبورد
			if val, ok := stateManager.GetState(senderID, "reply_keyboard_on"); ok && val == "true" {
				// کیبورد روشن است، پس آن را خاموش کن
				err = client.EditChatKeypad(ctx, &ParsRubika.EditChatKeypadRequest{
					ChatID:         chatID,
					ChatKeypadType: ParsRubika.RemoveKeypad,
				})
				stateManager.SetState(senderID, "reply_keyboard_on", "false")
				if err == nil {
					_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "کیبورد چت حذف شد."})
				}
			} else {
				// کیبورد خاموش است، پس آن را روشن کن
				replyKeyboard := &ParsRubika.Keypad{
					Rows: []ParsRubika.KeypadRow{
						{
							Buttons: []ParsRubika.Button{
								{ID: "opt1", Type: ParsRubika.ButtonTypeSimple, ButtonText: "گزینه ۱"},
								{ID: "opt2", Type: ParsRubika.ButtonTypeSimple, ButtonText: "گزینه ۲"},
							},
						},
					},
					ResizeKeyboard: true,
				}
				err = client.EditChatKeypad(ctx, &ParsRubika.EditChatKeypadRequest{
					ChatID:         chatID,
					ChatKeypad:     replyKeyboard,
					ChatKeypadType: ParsRubika.NewKeypad,
				})
				stateManager.SetState(senderID, "reply_keyboard_on", "true")
				if err == nil {
					_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "کیبورد چت فعال شد."})
				}
			}
		case "set_state_button":
			stateManager.SetState(senderID, "user_data", "این یک مقدار تستی است.")
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "وضعیت با موفقیت ذخیره شد."})
		case "get_state_button":
			if val, ok := stateManager.GetState(senderID, "user_data"); ok {
				_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "وضعیت فعلی شما:\n" + val.(string)})
			} else {
				_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "هیچ وضعیتی برای شما ذخیره نشده است."})
			}
		case "clear_state_button":
			stateManager.DeleteUserState(senderID)
			_, err = client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "تمام وضعیت‌های شما پاک شدند."})
		}
		return err
	}

	// 3. مدیریت متن‌های ساده (مثلاً دکمه‌های کیبورد چت)
	switch messageText {
	case "گزینه ۱":
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "شما گزینه ۱ را انتخاب کردید."})
		return err
	case "گزینه ۲":
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{ChatID: chatID, Text: "شما گزینه ۲ را انتخاب کردید."})
		return err
	default:
		// برای پیام‌های ناشناخته، یک پیام پیش‌فرض ارسال کن
		_, err := client.SendMessage(ctx, &ParsRubika.SendMessageRequest{
			ChatID: chatID,
			Text:   "پیام شما دریافت شد. برای شروع از دستور /start استفاده کنید.",
		})
		return err
	}
}

func main() {
	// --- 1. دریافت توکن از متغیر محیطی ---
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("متغیر محیطی BOT_TOKEN تنظیم نشده است.")
	}

	// --- 2. تنظیم فلگ برای انتخاب حالت اجرا ---
	mode := flag.String("mode", "polling", "حالت اجرا: 'polling' یا 'webhook'")
	flag.Parse()

	// --- 3. ساخت کلاینت ---
	client = ParsRubika.NewClient(botToken)

	// --- 4. ساخت Context برای مدیریت توقف برنامه ---
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// --- 5. مدیریت سیگنال‌های سیستمی (Ctrl+C) برای توقف elegance ---
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("سیگنال توقف دریافت شد. در حال بستن برنامه...")
		cancel()
	}()

	// --- 6. اجرای ربات بر اساس حالت انتخاب شده ---
	switch *mode {
	case "polling":
		log.Println("ربات در حالت Polling شروع به کار کرد...")
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
		log.Println("ربات در حالت Webhook شروع به کار کرد...")
		webhookURL := os.Getenv("WEBHOOK_URL") // مثلا: https://yourdomain.com/webhook
		if webhookURL == "" {
			log.Fatal("متغیر محیطی WEBHOOK_URL برای حالت وب‌هوک الزامی است.")
		}

		// تنظیم وب‌هوک در سرور روبیکا
		log.Printf("در حال تنظیم وب‌هوک به آدرس: %s", webhookURL)
		if err := client.UpdateBotEndpoints(ctx, webhookURL, ParsRubika.ReceiveUpdate); err != nil {
			log.Fatalf("خطا در تنظیم وب‌هوک: %v", err)
		}
		log.Println("وب‌هوک با موفقیت تنظیم شد.")

		webhookOpts := ParsRubika.WebhookOptions{
			Port:    8080,       // یا از متغیر محیطی بخوانید
			Path:    "/webhook", // باید با WEBHOOK_URL مطابقت داشته باشد
			Handler: handleUpdate,
		}

		if err := client.StartWebhook(ctx, webhookOpts); err != nil {
			log.Fatalf("خطا در راه‌اندازی سرور وب‌هوک: %v", err)
		}

	default:
		log.Fatalf("حالت اجرای نامعتبر است. از 'polling' یا 'webhook' استفاده کنید.")
	}

	log.Println("برنامه با موفقیت متوقف شد.")
}
