package ParsRubika

// سازنده: ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import (
	"context"
	"fmt"
	"log"
	"time"
)

// PollingOptions تنظیمات فرآیند Polling
type PollingOptions struct {
	Handler           HandlerFunc
	RetryTimeout      time.Duration // این فیلد دیگر به صورت مستقیم استفاده نمی‌شود و توسط NetworkStabilityManager مدیریت می‌شود
	Limit             int
	AllowEmptyUpdates bool
	PollInterval      time.Duration
	Timeout           time.Duration
}

// StartPolling شروع دریافت آپدیت‌ها با پولینگ (هوشمند شده)
func (c *BotClient) StartPolling(ctx context.Context, opts PollingOptions) error {
	if opts.Handler == nil {
		opts.Handler = c.ProcessUpdate
	}

	// مقادیر پیش‌فرض
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

	// پاک‌سازی صف آپدیت‌های قدیمی برای شروعی پاک
	log.Println("در حال پاک‌سازی صف آپدیت‌های قدیمی...")
	if err := c.clearUpdateQueue(ctx, &offset); err != nil {
		log.Printf("اخطار در پاک‌سازی صف: %v", err)
	}

	log.Println("شروع پولینگ هوشمند...")

	pollingTicker := time.NewTicker(opts.PollInterval)
	defer pollingTicker.Stop()

	var retryCount int
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
				// استفاده از NetworkStabilityManager برای تصمیم‌گیری هوشمندانه
				if c.networkStabilityManager.IsRetryableError(err) {
					delay := c.networkStabilityManager.CalculateBackoffDelay(retryCount)
					log.Printf("خطا در دریافت آپدیت‌ها: %v. تلاش مجدد در %v", err, delay)
					retryCount++
					time.Sleep(delay)
					continue
				} else {
					// اگر خطا قابل تلاش مجدد نبود، با خطا خارج می‌شویم
					log.Printf("خطای غیرقابل بازیابی در پولینگ: %v", err)
					return err
				}
			}

			// در صورت موفقیت، شمارنده تلاش مجدد ریست می‌شود
			retryCount = 0

			if len(updates.Updates) > 0 {
				log.Printf("تعداد %d آپدیت جدید دریافت شد", len(updates.Updates))

				for _, update := range updates.Updates {
					if err := opts.Handler(ctx, &update); err != nil {
						log.Printf("خطا در پردازش آپدیت: %v", err)
					}
				}

				// به‌روزرسانی آفست برای جلوگیری از دریافت آپدیت‌های تکراری
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

// clearUpdateQueue پاک‌سازی صف آپدیت‌های قدیمی (بهینه شده)
func (c *BotClient) clearUpdateQueue(ctx context.Context, offset **string) error {
	// یک حلقه برای خالی کردن کامل صف
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
			// اگر next_offset وجود نداشت، یعنی صف خالی شده است
			break
		}
	}
	return nil
}
