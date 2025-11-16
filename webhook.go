package ParsRubika

// سازنده ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type WebhookOptions struct {
	Port    int
	Path    string
	Handler HandlerFunc
}

// اصلاح شده: منطق توقف优雅 (graceful shutdown) اضافه شد
func (c *Client) StartWebhook(ctx context.Context, opts WebhookOptions) error {
	if opts.Handler == nil {
		return fmt.Errorf("تابع هندلر نمی‌تواند خالی باشد")
	}

	mux := http.NewServeMux()
	mux.HandleFunc(opts.Path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("خطا در خواندن بدنه درخواست: %v", err)
			http.Error(w, "خطا در خواندن بدنه", http.StatusBadRequest)
			return
		}

		var updateWrapper struct {
			Update *Update `json:"update"`
		}
		if err := json.Unmarshal(body, &updateWrapper); err == nil && updateWrapper.Update != nil {
			if err := opts.Handler(ctx, updateWrapper.Update); err != nil {
				log.Printf("هندلر با خطا مواجه شد: %v", err)
			}
			w.WriteHeader(http.StatusOK)
			return
		}

		var inlineWrapper struct {
			InlineMessage *InlineMessage `json:"inline_message"`
		}
		if err := json.Unmarshal(body, &inlineWrapper); err == nil && inlineWrapper.InlineMessage != nil {
			inlineMsg := inlineWrapper.InlineMessage
			update := &Update{
				Type:   NewMessage,
				ChatID: inlineMsg.ChatID,
				NewMessage: &Message{
					MessageID: inlineMsg.MessageID,
					Text:      inlineMsg.Text,
					SenderID:  inlineMsg.SenderID,
					AuxData:   inlineMsg.AuxData,
					File:      inlineMsg.File,
					Location:  inlineMsg.Location,
				},
			}
			if err := opts.Handler(ctx, update); err != nil {
				log.Printf("هندلر با خطا مواجه شد: %v", err)
			}
			w.WriteHeader(http.StatusOK)
			return
		}

		log.Printf("پارس کردن بدنه وب‌هوک ناموفق بود")
		http.Error(w, "فرمت بدنه نامعتبر است", http.StatusBadRequest)
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", opts.Port),
		Handler: mux,
	}

	// ایجاد یک goroutine برای مدیریت توقف سرور
	go func() {
		<-ctx.Done() // منتظر سیگنال توقف از context بمان
		log.Println("در حال بستن سرور وب‌هوک...")

		// یک context با timeout برای اطمینان از بسته شدن سرور
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("خطا در بستن سرور وب‌هوک: %v", err)
		} else {
			log.Println("سرور وب‌هوک با موفقیت بسته شد.")
		}
	}()

	log.Printf("سرور وب‌هوک روی پورت %d و مسیر %s راه‌اندازی شد", opts.Port, opts.Path)
	// server.ListenAndServe() مسدودکننده است و منتظر می‌ماند تا خطا رخ دهد یا سرور بسته شود
	return server.ListenAndServe()
}
