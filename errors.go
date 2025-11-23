package ParsRubika

// سازنده ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import "fmt"

// APIError ساختاری برای خطاهای بازگشتی از API روبیکا
type APIError struct {
	StatusCode int    `json:"-"` // وضعیت HTTP
	Message    string `json:"message"`
}

// Error متد رابط خطا برای نمایش پیام خطا
func (e *APIError) Error() string {
	return fmt.Sprintf("خطای API روبیکا (وضعیت %d): %s", e.StatusCode, e.Message)
}

// ErrUnsupportedMethod خطایی برای متدهای پشتیبانی نشده
var ErrUnsupportedMethod = fmt.Errorf("این متد توسط API رسمی روبیکا پشتیبانی نمی‌شود")
