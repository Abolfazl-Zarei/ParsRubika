
---

# 📘 **ParsRubika Bot Go – Full Documentation**

### قدرتمندترین کتابخانه Golang برای ساخت ربات‌های روبیکا 🇮🇷🤖

<div align="center">

<img src="https://sae22.ir/wp-content/uploads/2022/12/00.webp" width="180" />
<img src="https://img.icons8.com/color/120/000000/golang.png" width="160"/>
<img src="https://img.icons8.com/color/120/000000/robot.png" width="150"/>

**سازنده:** *ابوالفضل زارعی*
🔗 GitHub: [https://github.com/Abolfazl-Zarei/ParsRubika-bot-go](https://github.com/Abolfazl-Zarei/ParsRubika-bot-go)

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge\&logo=go)]()
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)]()
[![Stars](https://img.shields.io/github/stars/Abolfazl-Zarei/ParsRubika-bot-go?style=for-the-badge\&color=gold)]()
[![Issues](https://img.shields.io/github/issues/Abolfazl-Zarei/ParsRubika-bot-go?style=for-the-badge\&color=orange)]()

</div>

---

# 📚 فهرست مطالب

* [معرفی](#-معرفی)
* [ویژگی‌های کلیدی](#-ویژگیهای-کلیدی)
* [نصب](#-نصب)
* [شروع سریع: ربات Echo](#-شروع-سریع-ربات-echo)
* [ساختارهای اصلی](#-ساختارهای-اصلی)
* [مستندات کامل API](#️-مستندات-کامل-api)
* [مدیریت گروه و کانال](#-مدیریت-گروه-و-کانال)
* [صفحات، پست، استوری](#-صفحات-پست-استوری)
* [State Manager](#-state-manager)
* [مثال‌های پیشرفته](#-مثالهای-پیشرفته)
* [لایسنس و مشارکت](#-لایسنس-و-مشارکت)
* [سازنده](#-سازنده)

---

# 🌟 معرفی

**ParsRubika Bot Go** یک کتابخانه سریع، سبک و کاملاً حرفه‌ای برای ساخت ربات در پلتفرم روبیکا است.
این کتابخانه به گونه‌ای طراحی شده که:

* توسعه را ساده کند
* سرعت بالا داشته باشد
* امکانات کامل API روبیکا را پوشش دهد
* و حتی قابلیت‌هایی که روبیکا مستقیماً ارائه نکرده را نیز فراهم کند

📌 *این کتابخانه هم برای پروژه‌های کوچک مناسب است و هم برای ربات‌های بزرگ و پیچیده.*

---

# 🚀 ویژگی‌های کلیدی

جدول کامل قابلیت‌ها:

| دسته           | ویژگی‌ها                                                  |
| -------------- | --------------------------------------------------------- |
| 📡 ارتباط      | Polling پایدار، Webhook، مدیریت خطا، Timeout هوشمند       |
| 💬 پیام‌رسانی  | متن، عکس، ویدیو، ویس، فایل، استیکر، لوکیشن، تماس، نظرسنجی |
| ⌨ رابط کاربری  | کیبورد اصلی، اینلاین، دکمه‌های انتخابی                    |
| 👥 مدیریت گروه | Ban، Unban، Promote، Pin، Admin list، Kick                |
| 📱 صفحات       | پست، استوری، هایلایت، لایک، فالو، آنفالو، کامنت           |
| 🧠 مدیریت حالت | State Manager داخلی                                       |
| ⚙ کارایی       | بدون وابستگی خارجی، مصرف RAM پایین، سرعت بالا             |

---

# 💻 نصب

در ترمینال اجرا کنید:

```bash
go get github.com/Abolfazl-Zarei/ParsRubika-bot-go
```

---

# 🎯 شروع سریع: ربات Echo

```go
package main

import (
	"context"
	"log"

	ParsRubika "github.com/Abolfazl-Zarei/ParsRubika-bot-go"
)

func main() {
	bot := ParsRubika.NewClient("TOKEN")
	ctx := context.Background()

	bot.StartPolling(ctx, ParsRubika.PollingOptions{
		Handler: func(ctx context.Context, u *ParsRubika.Update) error {

			if u.NewMessage != nil {
				bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
					ChatID: u.ChatID,
					Text:   "Echo: " + u.NewMessage.Text,
				})
			}

			return nil
		},
	})
}
```

---

# 🏗 ساختارهای اصلی

## 📨 Update

```go
type Update struct {
	Type             UpdateTypeEnum
	ChatID           string
	NewMessage       *Message
	UpdatedMessage   *Message
	RemovedMessageID *string
}
```

## 💬 Message

```go
type Message struct {
	MessageID string
	Text      string
	File      *File
	Location  *Location
	Sticker   *Sticker
	SenderID  string
}
```

---

# 🛠️ مستندات کامل API

## 📩 ارسال پیام

```go
bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
	ChatID: "ID",
	Text:   "سلام! 👋",
})
```

---

## 📎 ارسال فایل

```go
file, _ := bot.UploadFileDirectly(ctx, "pic.png", ParsRubika.ImageType)

bot.SendFile(ctx, &ParsRubika.SendFileRequest{
	ChatID: "ID",
	FileID: file.FileID,
})
```

---

## 📍 ارسال لوکیشن

```go
bot.SendLocation(ctx, &ParsRubika.SendLocationRequest{
	ChatID:    "ID",
	Latitude:  "35.6892",
	Longitude: "51.3890",
})
```

---

## ☎ ارسال مخاطب

```go
bot.SendContact(ctx, &ParsRubika.SendContactRequest{
	ChatID:     "ID",
	FirstName:  "Ali",
	PhoneNumber:"+98...",
})
```

---

## 🗳 ارسال نظرسنجی

```go
bot.SendPoll(ctx, &ParsRubika.SendPollRequest{
	ChatID:   "ID",
	Question: "نظرت چیه؟",
	Options:  []string{"عالی", "خوب", "ضعیف"},
})
```

---

## ✏ ویرایش پیام

```go
bot.EditMessageText(ctx, &ParsRubika.EditMessageTextRequest{
	ChatID:    "ID",
	MessageID: "MID",
	Text:      "ویرایش شد ✏️",
})
```

---

## ❌ حذف پیام

```go
bot.DeleteMessage(ctx, &ParsRubika.DeleteMessageRequest{
	ChatID: "ID",
	MessageID: "MID",
})
```

---

# 👥 مدیریت گروه و کانال

### Ban

```go
bot.BanChatMember(ctx, &ParsRubika.BanChatMemberRequest{
	ChatID: "GROUP_ID",
	UserID: "USER_ID",
})
```

### Promote

```go
yes := true

bot.PromoteChatMember(ctx, &ParsRubika.PromoteChatMemberRequest{
	ChatID:          "GROUP",
	UserID:          "USER",
	IsAdministrator: &yes,
})
```

### Admin List

```go
admins, _ := bot.GetChatAdministrators(ctx, &ParsRubika.GetChatAdministratorsRequest{
	ChatID: "GROUP",
})
```

---

# 📱 صفحات، پست، استوری

### پست

```go
bot.AddPost(ctx, "Hello", nil)
```

### استوری

```go
bot.AddStory(ctx, fileID, "caption")
```

### هایلایت

```go
bot.CreateHighlight(ctx, "title", []string{storyID})
```

### Explore

```go
posts, _ := bot.GetExplorePosts(ctx)
```

---

# 🧠 State Manager

```go
sm := ParsRubika.NewStateManager()

sm.SetState("UserID", "step", 1)

step, ok := sm.GetState("UserID", "step")
```

---

# 🎯 مثال‌های پیشرفته

اگر خواستی می‌توانم ایجاد کنم:

* ربات مدیریت گروه
* ربات فروشگاهی
* ربات پرسش و پاسخ با State
* ربات ضد اسپم
* ربات صفحه روبیکا

---

# 📜 لایسنس

این پروژه تحت لایسنس MIT منتشر شده است.

---

# 👨‍💻 سازنده

**ابوالفضل زارعی**
توسعه‌دهنده Go و طراح کتابخانه ParsRubika Bot Go
GitHub: [https://github.com/Abolfazl-Zarei](https://github.com/Abolfazl-Zarei)

---

# ❤️ پایان

اگر می‌خواهی:

✨ نسخه انگلیسی
✨ نسخه کوتاه Dev-Friendly
✨ نسخه با رنگ‌بندی گرافیکی
✨ اضافه کردن بنر اختصاصی

فقط بگو تا بسازم.
