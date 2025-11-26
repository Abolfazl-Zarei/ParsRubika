package ParsRubika

// سازنده: ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import (
	"math/rand"
	"net"
	"strings"
	"time"
)

// NetworkStabilityManager مدیریت پایداری شبکه
type NetworkStabilityManager struct {
	client *BotClient
}

// NewNetworkStabilityManager ایجاد یک نمونه جدید
func NewNetworkStabilityManager(client *BotClient) *NetworkStabilityManager {
	return &NetworkStabilityManager{client: client}
}

// CalculateBackoffDelay محاسبه تأخیر برای تلاش مجدد با الگوریتم نمایی و Jitter
func (nsm *NetworkStabilityManager) CalculateBackoffDelay(retryCount int) time.Duration {
	// الگوریتم نمایی: baseDelay * (2 ^ retryCount)
	baseDelay := 1 * time.Second
	maxDelay := 30 * time.Second

	exponentialDelay := time.Duration(1<<uint(retryCount)) * baseDelay
	if exponentialDelay > maxDelay {
		exponentialDelay = maxDelay
	}

	// همیشه Jitter (تأخیر تصادفی) اضافه می‌کنیم.
	// این کار از ارسال همزمان درخواست‌های متعدد در زمان‌های مشخص جلوگیری کرده
	// و رفتار ربات را طبیعی‌تر می‌کند.
	// Jitter تا 25% از تأخیر اصلی
	jitter := time.Duration(rand.Int63n(int64(exponentialDelay / 4)))
	return exponentialDelay + jitter
}

// IsRetryableError بررسی اینکه آیا یک خطا قابل تلاش مجدد است یا خیر
func (nsm *NetworkStabilityManager) IsRetryableError(err error) bool {
	if err == nil {
		return false
	}

	// اگر گزینه ignore_timeout فعال باشد، خطاهای timeout نادیده گرفته شده و قابل تلاش مجدد هستند
	if nsm.client.ignoreTimeout && isTimeoutError(err) {
		return true
	}

	// بررسی خطاهای شبکه دیگر
	if isNetworkError(err) {
		return true
	}

	return false
}

// isTimeoutError بررسی اینکه آیا خطا از نوع timeout است
func isTimeoutError(err error) bool {
	errStr := err.Error()
	return strings.Contains(errStr, "timeout") ||
		strings.Contains(errStr, "deadline exceeded") ||
		strings.Contains(errStr, "context deadline exceeded")
}

// isNetworkError بررسی اینکه آیا خطا از نوع شبکه است
func isNetworkError(err error) bool {
	if netErr, ok := err.(net.Error); ok {
		// اگر خطا موقتی بود، قابل تلاش مجدد است
		return netErr.Temporary() || netErr.Timeout()
	}
	errStr := err.Error()
	return strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "connection reset") ||
		strings.Contains(errStr, "no such host")
}
