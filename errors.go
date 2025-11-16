package ParsRubika

// سازنده ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import "fmt"

type APIError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("خطای API روبیکا (وضعیت %d): %s", e.StatusCode, e.Message)
}

var ErrUnsupportedMethod = fmt.Errorf("این متد توسط API رسمی روبیکا پشتیبانی نمی‌شود")
