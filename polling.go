package ParsRubika

// سازنده ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import (
	"context"
	"fmt"
	"log"
	"time"
)

// اصلاح شده: نام فیلد برای وضوح بیشتر تغییر یافت
// PollingOptions تنظیمات فرآیند Polling
type PollingOptions struct {
	Handler           HandlerFunc
	RetryTimeout      time.Duration // زمان انتظار بین هر درخواست در صورت بروز خطا
	Limit             int           // تعداد آپدیت در هر درخواست
	AllowEmptyUpdates bool          // آیا به آپدیت‌های خالی نیز هندلر فراخوانی شود؟
	PollInterval      time.Duration // فاصله زمانی بین هر درخواست موفق
}

// StartPolling فرآیند دریافت مداوم آپدیت‌ها را با متد Polling آغاز می‌کند
func (c *Client) StartPolling(ctx context.Context, opts PollingOptions) error {
	if opts.Handler == nil {
		return fmt.Errorf("handler function cannot be nil")
	}
	// مقادیر پیش‌فرض برای تنظیمات
	if opts.RetryTimeout == 0 {
		opts.RetryTimeout = 5 * time.Second
	}
	if opts.Limit == 0 {
		opts.Limit = 100
	}
	if opts.PollInterval == 0 {
		opts.PollInterval = 3 * time.Second
	}

	// دریافت شناسه بات برای جلوگیری از واکنش به پیام‌های خودش
	if err := c.getBotID(ctx); err != nil {
		return fmt.Errorf("could not start polling without bot ID: %w", err)
	}

	var offset *string

	// --- لایه دفاعی اول: پاک‌سازی کامل و قطعی صف ---
	log.Println("در حال پاک‌سازی کامل پیام‌های قدیمی از صف...")
	for {
		// تا ۱۰۰ آپدیت قدیمی را برای خالی کردن صف دریافت می‌کنیم
		discardUpdates, err := c.GetUpdates(ctx, offset, 100)
		if err != nil {
			log.Printf("اخطار: خطا در دریافت آپدیت‌های اولیه برای دور ریختن: %v", err)
			break // در صورت خطا از حلقه خارج شو
		}
		if len(discardUpdates.Updates) == 0 {
			log.Println("صف با موفقیت خالی شد. هیچ آپدیت قدیمی باقی نمانده است.")
			break // وقتی آدیتی وجود نداشت، از حلقه خارج شو
		}
		log.Printf("تعداد %d آپدیت قدیمی دور ریخته شد. در حال بررسی مجدد...", len(discardUpdates.Updates))
		// آفست را به آخرین آپدیت دریافت‌شده تنظیم می‌کنیم تا از شروع کنیم
		if discardUpdates.NextOffsetID != "" {
			offset = &discardUpdates.NextOffsetID
		}
	}
	log.Println("پاک‌سازی صف تمام شد. حلقه اصلی Polling شروع به کار کرد...")
	// --- پایان لایه دفاعی ---

	for {
		select {
		case <-ctx.Done():
			log.Println("Polling متوقف شد توسط context.")
			return ctx.Err()
		default:
			// --- لایه دفاعی دوم: خنثی‌سازی ---
			if c.isInCooldown() {
				log.Println("در دوره خنثی‌سازی (Cooldown). آپدیت نادیده گرفته می‌شود.")
				// اصلاح شده: برای جلوگیری از حلقه فشرده، به مدت PollInterval استاندارد صبر می‌کنیم
				time.Sleep(opts.PollInterval)
				continue
			}

			updates, err := c.GetUpdates(ctx, offset, opts.Limit)
			if err != nil {
				log.Printf("خطا در دریافت آپدیت‌ها: %v. تلاش مجدد در %v...", err, opts.RetryTimeout)
				time.Sleep(opts.RetryTimeout)
				continue
			}

			for _, update := range updates.Updates {
				// --- دیباگ قوی ---
				senderID := "N/A"
				if update.NewMessage != nil {
					senderID = update.NewMessage.SenderID
				}
				log.Printf("آپدیت دریافت شد. فرستنده: %s | بات: %s", senderID, *c.botID)

				// --- لایه دفاعی سوم: نادیده گرفتن پیام خود بات ---
				if update.NewMessage != nil && c.botID != nil && update.NewMessage.SenderID == *c.botID {
					log.Printf("پیام از طرف خود بات نادیده گرفته شد.")
					continue
				}

				if err := opts.Handler(ctx, &update); err != nil {
					log.Printf("هندلر برای آپدیت %s با خطا مواجه شد: %v", update.NewMessage.MessageID, err)
				}
			}

			if updates.NextOffsetID != "" {
				offset = &updates.NextOffsetID
			}

			time.Sleep(opts.PollInterval)
		}
	}
}
