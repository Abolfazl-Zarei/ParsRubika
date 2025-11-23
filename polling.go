package ParsRubika

import (
	"context"
	"fmt"
	"log"
	"time"
)

// PollingOptions تنظیمات فرآیند Polling
type PollingOptions struct {
	Handler           HandlerFunc
	RetryTimeout      time.Duration
	Limit             int
	AllowEmptyUpdates bool
	PollInterval      time.Duration
	Timeout           time.Duration
}

// StartPolling شروع دریافت آپدیت‌ها با پولینگ
func (c *BotClient) StartPolling(ctx context.Context, opts PollingOptions) error {
	if opts.Handler == nil {
		opts.Handler = c.ProcessUpdate
	}

	// مقادیر پیش‌فرض
	if opts.RetryTimeout == 0 {
		opts.RetryTimeout = 5 * time.Second
	}
	if opts.Limit == 0 {
		opts.Limit = 100
	}
	if opts.PollInterval == 0 {
		opts.PollInterval = 2 * time.Second
	}
	if opts.Timeout == 0 {
		opts.Timeout = 30 * time.Second
	}

	// دریافت شناسه بات
	if err := c.getBotID(ctx); err != nil {
		return fmt.Errorf("could not start polling without bot ID: %w", err)
	}

	var offset *string

	// پاک‌سازی صف آپدیت‌های قدیمی
	log.Println("در حال پاک‌سازی صف آپدیت‌های قدیمی...")
	if err := c.clearUpdateQueue(ctx, &offset); err != nil {
		log.Printf("اخطار در پاک‌سازی صف: %v", err)
	}

	log.Println("شروع پولینگ...")

	pollingTicker := time.NewTicker(opts.PollInterval)
	defer pollingTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("پولینگ توسط context متوقف شد")
			return ctx.Err()

		case <-c.stopChan:
			log.Println("پولینگ توسط بات متوقف شد")
			return nil

		case <-pollingTicker.C:
			updates, err := c.GetUpdates(ctx, offset, opts.Limit)
			if err != nil {
				log.Printf("خطا در دریافت آپدیت‌ها: %v. تلاش مجدد در %v", err, opts.RetryTimeout)
				time.Sleep(opts.RetryTimeout)
				continue
			}

			if len(updates.Updates) > 0 {
				log.Printf("تعداد %d آپدیت جدید دریافت شد", len(updates.Updates))

				for _, update := range updates.Updates {
					if err := opts.Handler(ctx, &update); err != nil {
						log.Printf("خطا در پردازش آپدیت: %v", err)
					}
				}

				// به‌روزرسانی آفست
				if updates.NextOffsetID != "" {
					offset = &updates.NextOffsetID
				}
			} else if opts.AllowEmptyUpdates {
				// اگر هیچ آپدیتی نبود اما اجازه داده شده، هندلر را صدا بزن
				emptyUpdate := &Update{Type: NewMessage}
				if err := opts.Handler(ctx, emptyUpdate); err != nil {
					log.Printf("خطا در پردازش آپدیت خالی: %v", err)
				}
			}
		}
	}
}

// clearUpdateQueue پاک‌سازی صف آپدیت‌های قدیمی
func (c *BotClient) clearUpdateQueue(ctx context.Context, offset **string) error {
	for {
		discardUpdates, err := c.GetUpdates(ctx, *offset, 100)
		if err != nil {
			return fmt.Errorf("خطا در پاک‌سازی صف: %w", err)
		}

		if len(discardUpdates.Updates) == 0 {
			log.Println("صف آپدیت‌ها با موفقیت پاک شد")
			break
		}

		log.Printf("تعداد %d آپدیت قدیمی دور ریخته شد", len(discardUpdates.Updates))

		if discardUpdates.NextOffsetID != "" {
			*offset = &discardUpdates.NextOffsetID
		} else {
			break
		}
	}
	return nil
}
