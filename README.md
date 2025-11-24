# 🤖 ParsRubika - کامل‌ترین کتابخانه Golang برای روبیکا

<div align="center">

![Rubika Bot](https://img.shields.io/badge/Rubika-Bot%20API-red?style=for-the-badge&logo=telegram&logoColor=white)
![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Stars](https://img.shields.io/github/stars/Abolfazl-Zarei/ParsRubika-bot-go?style=for-the-badge&color=gold)

<br />

<img src="https://img.icons8.com/color/96/000000/iran.png" width="80"/>
<img src="https://img.icons8.com/color/96/000000/golang.png" width="80"/>
<img src="https://img.icons8.com/color/96/000000/robot-2.png" width="80"/>
<img src="https://img.icons8.com/color/96/000000/api.png" width="80"/>

**🔗 مخزن گیت‌هاب:**  
[https://github.com/Abolfazl-Zarei/ParsRubika-bot-go](https://github.com/Abolfazl-Zarei/ParsRubika-bot-go)

**👤 توسعه‌دهنده:** ابوالفضل زارعی  
**📧 ایمیل:** `ninjacode.ir@gmail.com`  
**🆔 روبیکا:** `NinjaCode`  
**📢 چنل روبیکا:** `Ninja_code`

</div>

## 📚 فهرست جامع

1. [🎯 معرفی کامل](#-معرفی-کامل)
2. [⚡ نصب و راه‌اندازی](#-نصب-و-راهاندازی)
3. [🏗 مفاهیم پایه](#-مفاهیم-پایه)
4. [🚀 شروع سریع](#-شروع-سریع)
5. [🛠 آموزش قدم به قدم](#-آموزش-قدم-به-قدم)
6. [📡 API Reference](#-api-reference)
7. [🎛 مدیریت وضعیت](#-مدیریت-وضعیت)
8. [⌨ کیبوردهای پویا](#-کیبوردهای-پویا)
9. [📁 مدیریت فایل‌ها](#-مدیریت-فایلها)
10. [🌐 Webhook & Polling](#-webhook--polling)
11. [🚀 مثال‌های پیشرفته](#-مثالهای-پیشرفته)
12. [☁ استقرار و دیپلوی](#-استقرار-و-دیپلوی)
13. [🔧 عیب‌یابی](#-عیبیابی)
14. [📞 پشتیبانی](#-پشتیبانی)

---

## 🎯 معرفی کامل

### ✨ پارس روبیکا چیست؟

**ParsRubika** یک کتابخانه **کاملاً فارسی** و **قدرتمند** برای ساخت ربات‌های پیام‌رسان **روبیکا** با زبان **Go** است. این کتابخانه با معماری مدرن و امکانات پیشرفته، توسعه ربات‌های حرفه‌ای را برای شما ساده می‌کند.

### 🌟 چرا ParsRubika؟

| ویژگی | 🎯 توضیح |
|-------|-----------|
| **✅ پشتیبانی کامل API** | تمام متدهای رسمی روبیکا |
| **🔄 دو روش دریافت** | Polling و Webhook |
| **🎛 مدیریت وضعیت** | State Management پیشرفته |
| **📁 آپلود/دانلود** | پشتیبانی از تمام فرمت‌ها |
| **⌨ کیبوردهای پویا** | انواع دکمه و اینترفیس |
| **🛡 خطایابی هوشمند** | مدیریت خودکار خطاها |
| **⚡ Performance بالا** | بهینه‌شده برای سرعت |
| **📚 مستندات کامل** | آموزش‌های قدم به قدم |

### 🏗 معماری کتابخانه

```
📦 ParsRubika/
├── 📄 client.go          # 🎯 کلاینت اصلی و منطق کسب‌وکار
├── 🏷️ models.go          # 📊 مدل‌های داده و ساختارها
├── 🔢 enums.go           # 🎮 انواع شمارشی و ثابت‌ها
├── ❌ errors.go          # 🚨 مدیریت خطاها
├── 🔄 polling.go         # 📡 سیستم پولینگ
├── 🌐 webhook.go         # 🌍 سیستم وب‌هوک
├── 💾 state.go           # 💡 مدیریت وضعیت کاربران
└── 📋 go.mod            # 📦 وابستگی‌های پروژه
```

---

## ⚡ نصب و راه‌اندازی

### 📋 پیش‌نیازها

#### 1. نصب Go
```bash
# 🐧 اوبونتی/دبیان
sudo apt update && sudo apt install golang-go

# 🍎 مک
brew install go

# 🪟 ویندوز
# از سایت https://golang.org/dl دانلود کنید
```

#### 2. بررسی نسخه Go
```bash
go version
# خروجی باید باشد: go version go1.21.x یا بالاتر
```

#### 3. دریافت توکن ربات

1. در روبیکا به `@BotFather` مراجعه کنید
2. دستور `/newbot` را ارسال کنید
3. نام ربات را وارد کنید (مثال: `MyAwesomeBot`)
4. یوزرنیم ربات را وارد کنید (مثال: `my_awesome_bot`)
5. **توکن دریافتی** را ذخیره کنید

### 📥 نصب کتابخانه

#### روش 1: نصب مستقیم از گیت‌هاب
```bash
go get github.com/Abolfazl-Zarei/ParsRubika-bot-go
```

#### روش 2: کلون کردن مخزن
```bash
git clone https://github.com/Abolfazl-Zarei/ParsRubika-bot-go.git
cd ParsRubika-bot-go
```

#### روش 3: استفاده در پروژه جدید
```bash
mkdir my-rubika-bot
cd my-rubika-bot
go mod init my-rubika-bot
go get github.com/Abolfazl-Zarei/ParsRubika-bot-go
```

---

## 🏗 مفاهیم پایه

### 🎮 ساختار اصلی آپدیت‌ها

```go
// 📨 ساختار Update - اصلی‌ترین بخش کتابخانه
type Update struct {
    Type             UpdateTypeEnum `json:"type"`              // 🏷️ نوع آپدیت
    ChatID           string         `json:"chat_id"`           // 💬 شناسه چت
    RemovedMessageID *string        `json:"removed_message_id"`// 🗑️ شناسه پیام حذف شده
    NewMessage       *Message       `json:"new_message"`       // 📩 پیام جدید
    UpdatedMessage   *Message       `json:"updated_message"`   // ✏️ پیام ویرایش شده
    UpdatedPayment   *PaymentStatus `json:"updated_payment"`   // 💰 وضعیت پرداخت
}
```

### 📧 انواع آپدیت‌ها

```go
// 🎯 انواع مختلف آپدیت‌هایی که ربات دریافت می‌کند
const (
    UpdatedMessage UpdateTypeEnum = "UpdatedMessage"  // ✏️ ویرایش پیام
    NewMessage     UpdateTypeEnum = "NewMessage"      // 📩 پیام جدید
    RemovedMessage UpdateTypeEnum = "RemovedMessage"  // 🗑️ حذف پیام
    StartedBot     UpdateTypeEnum = "StartedBot"      // 🚀 شروع بات
    StoppedBot     UpdateTypeEnum = "StoppedBot"      // 🛑 توقف بات
)
```

### 💌 ساختار پیام

```go
// 📝 ساختار کامل یک پیام
type Message struct {
    MessageID        int64             `json:"message_id"`         // 🔢 شناسه پیام
    Text             string            `json:"text"`               // 📄 متن پیام
    Time             string            `json:"time"`               // ⏰ زمان ارسال
    SenderType       MessageSenderEnum `json:"sender_type"`        // 👤 نوع فرستنده
    SenderID         string            `json:"sender_id"`          // 🆔 شناسه فرستنده
    File             *File             `json:"file"`               // 📁 فایل پیام
    ReplyToMessageID string            `json:"reply_to_message_id"`// ↩️ پاسخ به پیام
    Location         *Location         `json:"location"`           // 📍 موقعیت مکانی
    Sticker          *Sticker          `json:"sticker"`            // 🎨 استیکر
    ContactMessage   *ContactMessage   `json:"contact_message"`    // 👥 مخاطب
    Poll             *Poll             `json:"poll"`               // 📊 نظرسنجی
}
```

---

## 🚀 شروع سریع

### 🎯 اولین ربات شما در 5 دقیقه!

#### 1. ایجاد فایل اصلی

```go
// 📄 main.go
package main

import (
    "context"
    "log"
    "os"
    
    // 📦 ایمپورت کتابخانه ParsRubika
    ParsRubika "github.com/Abolfazl-Zarei/ParsRubika-bot-go"
)

func main() {
    // 🔑 دریافت توکن از متغیر محیطی
    botToken := os.Getenv("RUBIKA_BOT_TOKEN")
    if botToken == "" {
        log.Fatal("❌ توکن ربات یافت نشد! لطفا متغیر محیطی RUBIKA_BOT_TOKEN را تنظیم کنید.")
    }
    
    // 🤖 ایجاد نمونه ربات
    bot := ParsRubika.NewClient(botToken)
    
    // 🎯 تنظیم هندلر برای پیام‌ها
    bot.OnMessageUpdates(func(ctx context.Context, update *ParsRubika.Update) error {
        if update.NewMessage != nil {
            // 📨 پاسخ به پیام کاربر
            _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
                ChatID: update.ChatID,
                Text:   "👋 سلام! من با ParsRubika ساخته شده‌ام! 🚀",
            })
            return err
        }
        return nil
    })
    
    // 🚀 اجرای ربات
    ctx := context.Background()
    log.Println("🤖 ربات در حال راه‌اندازی...")
    if err := bot.Run(ctx); err != nil {
        log.Fatal("💥 خطا در اجرای ربات:", err)
    }
}
```

#### 2. اجرای ربات

```bash
# 🔑 تنظیم توکن
export RUBIKA_BOT_TOKEN="your_bot_token_here"

# 🚀 اجرا
go run main.go
```

#### 3. خروجی مورد انتظار

```
🤖 ربات در حال راه‌اندازی...
✅ بات با شناسه [BOT_ID] مقداردهی اولیه شد
🚀 بات شروع به کار کرد
📡 شروع پولینگ...
```

---

## 🛠 آموزش قدم به قدم

### 📁 قدم 1: ایجاد ساختار پروژه

```bash
# 📂 ایجاد پوشه پروژه
mkdir my-advanced-bot
cd my-advanced-bot

# 📦 مقداردهی اولیه پروژه Go
go mod init my-advanced-bot

# 📥 نصب کتابخانه ParsRubika
go get github.com/Abolfazl-Zarei/ParsRubika-bot-go

# 📄 ایجاد فایل‌های پروژه
touch main.go handlers.go utils.go
```

### 📄 قدم 2: فایل اصلی (main.go)

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "strings"
    
    ParsRubika "github.com/Abolfazl-Zarei/ParsRubika-bot-go"
)

func main() {
    // 🔐 دریافت توکن ربات
    token := os.Getenv("RUBIKA_BOT_TOKEN")
    if token == "" {
        log.Fatal(`
        ❌ خطا: توکن ربات تنظیم نشده است!
        
        📝 راه‌حل:
        1. توکن ربات خود را از @rubika_bot دریافت کنید
        2. متغیر محیطی تنظیم کنید:
           export RUBIKA_BOT_TOKEN="your_token_here"
        3. یا مستقیماً در کد قرار دهید (غیرامن)
        `)
    }
    
    // 🤖 ایجاد نمونه ربات
    bot := ParsRubika.NewClient(token)
    
    // ⚙️ تنظیم هندلرها
    setupHandlers(bot)
    
    // 🚀 اجرای ربات
    ctx := context.Background()
    log.Println(`
    🎉 ربات در حال راه‌اندازی...
    📍 برای توقف: Ctrl+C
    📱 برای تست: به ربات پیام بدهید!
    `)
    
    if err := bot.Run(ctx); err != nil {
        log.Fatalf("💥 خطا در اجرای ربات: %v", err)
    }
}

// ⚙️ تنظیم تمام هندلرها
func setupHandlers(bot *ParsRubika.BotClient) {
    // 📨 هندلر اصلی پیام‌ها
    bot.OnMessageUpdates(func(ctx context.Context, update *ParsRubika.Update) error {
        if update.NewMessage == nil {
            return nil
        }
        
        msg := update.NewMessage
        chatID := update.ChatID
        
        // 🎯 مسیریابی دستورات
        switch {
        case msg.Text == "/start":
            return handleStart(ctx, bot, chatID)
        case msg.Text == "/help":
            return handleHelp(ctx, bot, chatID)
        case strings.HasPrefix(msg.Text, "/echo"):
            return handleEcho(ctx, bot, chatID, msg.Text)
        case msg.Text == "/info":
            return handleInfo(ctx, bot, update)
        default:
            return handleDefault(ctx, bot, chatID, msg.Text)
        }
    })
    
    // 🔔 هندلر شروع ربات
    bot.OnStart(func(ctx context.Context, update *ParsRubika.Update) error {
        log.Println("✅ ربات با موفقیت راه‌اندازی شد!")
        return nil
    })
}
```

### 📄 قدم 3: هندلرها (handlers.go)

```go
package main

import (
    "context"
    "fmt"
    "strings"
    
    ParsRubika "github.com/Abolfazl-Zarei/ParsRubika-bot-go"
)

// 🎯 هندلر دستور /start
func handleStart(ctx context.Context, bot *ParsRubika.BotClient, chatID string) error {
    welcomeText := `🎉 **به ربات خوش آمدید!**

🤖 **من یک ربات روبیکا هستم که با ParsRubika ساخته شده‌ام**

📋 **دستورات موجود:**
/start - نمایش این پیام
/help - راهنمای کامل ربات  
/echo [متن] - تکرار متن شما
/info - اطلاعات کاربر

🔧 **ساخته شده با:** 
• زبان Go 🦫
• کتابخانه ParsRubika 🚀
• توسط NinjaCode 👨‍💻

💡 **شروع کنید:** یک دستور ارسال کنید یا پیام دلخواه بفرستید!`
    
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   welcomeText,
    })
    return err
}

// 🆘 هندلر دستور /help
func handleHelp(ctx context.Context, bot *ParsRubika.BotClient, chatID string) error {
    helpText := `📚 **راهنمای کامل ربات**

🎯 **دستورات اصلی:**
🔹 /start - شروع کار با ربات
🔹 /help - نمایش این راهنما  
🔹 /echo [متن] - ارسال متن به ربات
🔹 /info - دریافت اطلاعات کاربر

🛠 **امکانات پیشرفته:**
• 📨 ارسال پیام‌های متنی
• 🖼 ارسال عکس و مدیا
• 📍 ارسال موقعیت مکانی
• 📊 ایجاد نظرسنجی
• ⌨️ کیبوردهای تعاملی

💡 **نکات:**
• شما می‌توانید هر متنی را برای ربات ارسال کنید
• ربات به تمام پیام‌ها پاسخ می‌دهد
• از دستورات برای قابلیت‌های خاص استفاده کنید

🔗 **پشتیبانی:** 
برای گزارش مشکل یا پیشنهاد به @NinjaCode پیام دهید`
    
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   helpText,
    })
    return err
}

// 🔄 هندلر دستور /echo
func handleEcho(ctx context.Context, bot *ParsRubika.BotClient, chatID, text string) error {
    // حذف "/echo" از متن
    echoText := strings.TrimSpace(strings.TrimPrefix(text, "/echo"))
    
    if echoText == "" {
        echoText = "📝 لطفا متنی برای تکرار وارد کنید.\n\n💡 **مثال:**\n/echo سلام دنیا!"
    }
    
    response := fmt.Sprintf("🔊 **تکرار متن شما:**\n\n%s", echoText)
    
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   response,
    })
    return err
}

// ℹ️ هندلر دستور /info
func handleInfo(ctx context.Context, bot *ParsRubika.BotClient, update *ParsRubika.Update) error {
    userInfo, err := bot.GetUserInfo(ctx, update.NewMessage.SenderID)
    if err != nil {
        // ❌ در صورت خطا، پیام مناسب ارسال کنید
        _, err = bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: update.ChatID,
            Text:   "❌ خطا در دریافت اطلاعات کاربر. لطفا بعدا تلاش کنید.",
        })
        return err
    }
    
    infoText := fmt.Sprintf(`👤 **اطلاعات کاربر**

🆔 **شناسه:** %s
👤 **نام:** %s %s
📛 **یوزرنیم:** @%s
📝 **بیو:** %s
✅ **تأیید شده:** %v
🔒 **خصوصی:** %v`,
        userInfo.UserID,
        userInfo.FirstName,
        userInfo.LastName,
        userInfo.Username,
        userInfo.Bio,
        userInfo.IsVerified,
        userInfo.IsPrivate,
    )
    
    _, err = bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: update.ChatID,
        Text:   infoText,
    })
    return err
}

// 💬 هندلر پیام‌های معمولی
func handleDefault(ctx context.Context, bot *ParsRubika.BotClient, chatID, text string) error {
    response := fmt.Sprintf(`💬 **پیام شما دریافت شد!**

📝 **متن شما:** %s

💡 **راهنما:** 
از دستور /help برای مشاهده امکانات استفاده کنید.`, text)
    
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   response,
    })
    return err
}
```

### 📄 قدم 4: ابزارهای کمکی (utils.go)

```go
package main

import (
    "context"
    "fmt"
    
    ParsRubika "github.com/Abolfazl-Zarei/ParsRubika-bot-go"
)

// 🔧 ایجاد کیبورد ساده
func createSimpleKeyboard() *ParsRubika.Keypad {
    return &ParsRubika.Keypad{
        Rows: []ParsRubika.KeypadRow{
            {
                Buttons: []ParsRubika.Button{
                    {
                        ID:         "btn_help",
                        Type:       ParsRubika.ButtonTypeSimple,
                        ButtonText: "📚 راهنما",
                    },
                    {
                        ID:         "btn_info", 
                        Type:       ParsRubika.ButtonTypeSimple,
                        ButtonText: "👤 اطلاعات",
                    },
                },
            },
            {
                Buttons: []ParsRubika.Button{
                    {
                        ID:         "btn_echo",
                        Type:       ParsRubika.ButtonTypeSimple,
                        ButtonText: "🔊 تکرار متن",
                    },
                },
            },
        },
        ResizeKeyboard: true,
    }
}

// 🎯 ارسال پیام با کیبورد
func sendMessageWithKeyboard(ctx context.Context, bot *ParsRubika.BotClient, chatID, text string) error {
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID:       chatID,
        Text:         text,
        InlineKeypad: createSimpleKeyboard(),
    })
    return err
}

// 📊 نمایش اطلاعات ربات
func displayBotInfo(ctx context.Context, bot *ParsRubika.BotClient, chatID string) error {
    botInfo, err := bot.GetMe(ctx)
    if err != nil {
        return err
    }
    
    infoText := fmt.Sprintf(`🤖 **اطلاعات ربات**

🏷️ **نام:** %s
📛 **یوزرنیم:** @%s
🆔 **شناسه:** %s
📝 **توضیحات:** %s
🔗 **لینک اشتراک:** %s`,
        botInfo.BotTitle,
        botInfo.Username,
        botInfo.BotID,
        botInfo.Description,
        botInfo.ShareURL,
    )
    
    _, err = bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   infoText,
    })
    return err
}
```

### 🚀 قدم 5: اجرای نهایی

```bash
# 1. 🔑 تنظیم توکن
export RUBIKA_BOT_TOKEN="your_actual_bot_token_here"

# 2. 📦 دانلود وابستگی‌ها
go mod tidy

# 3. 🏗 ساخت پروژه
go build -o my-bot

# 4. 🚀 اجرای ربات
./my-bot

# یا برای اجرای مستقیم:
go run main.go handlers.go utils.go
```

---

## 📡 API Reference کامل

### 💬 مدیریت پیام‌ها

#### 📨 ارسال پیام متنی
```go
messageID, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:   "CHAT_ID",              // 💬 شناسه چت مقصد
    Text:     "متن پیام شما",         // 📝 متن پیام
    ReplyToMessageID: "MSG_ID",       // ↩️ پاسخ به پیام خاص (اختیاری)
})
```

#### ⌨️ ارسال پیام با کیبورد
```go
// 🎮 ایجاد کیبورد
keypad := &ParsRubika.Keypad{
    Rows: []ParsRubika.KeypadRow{
        {
            Buttons: []ParsRubika.Button{
                {
                    ID:         "btn1",
                    Type:       ParsRubika.ButtonTypeSimple,
                    ButtonText: "🎯 دکمه ۱",
                },
                {
                    ID:         "btn2",
                    Type:       ParsRubika.ButtonTypeSimple, 
                    ButtonText: "🚀 دکمه ۲",
                },
            },
        },
    },
    ResizeKeyboard: true,  // 📱 تنظیم سایز برای موبایل
}

// 📨 ارسال پیام با کیبورد
messageID, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:       "CHAT_ID",
    Text:         "پیام با کیبورد تعاملی 🎮",
    InlineKeypad: keypad,
})
```

#### ✏️ ویرایش پیام
```go
err := bot.EditMessageText(ctx, &ParsRubika.EditMessageTextRequest{
    ChatID:    "CHAT_ID",      // 💬 شناسه چت
    MessageID: "MESSAGE_ID",   // 🔢 شناسه پیام
    Text:      "متن جدید",     // 📝 متن جدید
})
```

#### 🗑️ حذف پیام
```go
err := bot.DeleteMessage(ctx, &ParsRubika.DeleteMessageRequest{
    ChatID:    "CHAT_ID",      // 💬 شناسه چت
    MessageID: "MESSAGE_ID",   // 🔢 شناسه پیام
})
```

### 🖼 ارسال مدیا

#### 🖼️ ارسال عکس
```go
messageID, err := bot.SendPhoto(ctx, "CHAT_ID", "path/to/image.jpg", "عنوان عکس")
```

#### 🎬 ارسال ویدیو
```go
messageID, err := bot.SendVideo(ctx, "CHAT_ID", "path/to/video.mp4", "عنوان ویدیو")
```

#### 📄 ارسال فایل
```go
messageID, err := bot.SendDocument(ctx, "CHAT_ID", "path/to/file.pdf", "عنوان فایل")
```

#### 🎵 ارسال صدا
```go
messageID, err := bot.SendVoice(ctx, "CHAT_ID", "path/to/voice.ogg", "عنوان صدا")
```

#### 🎨 ارسال استیکر
```go
messageID, err := bot.SendSticker(ctx, "CHAT_ID", "path/to/sticker.webp")
```

#### 📍 ارسال موقعیت
```go
messageID, err := bot.SendLocation(ctx, &ParsRubika.SendLocationRequest{
    ChatID:    "CHAT_ID",
    Latitude:  "35.6892",   // 📍 عرض جغرافیایی
    Longitude: "51.3890",   // 📍 طول جغرافیایی
})
```

#### 👥 ارسال مخاطب
```go
messageID, err := bot.SendContact(ctx, &ParsRubika.SendContactRequest{
    ChatID:      "CHAT_ID",
    FirstName:   "نام",
    LastName:    "نام خانوادگی", 
    PhoneNumber: "09123456789",
})
```

### 💬 مدیریت چت و کاربران

#### 💬 اطلاعات چت
```go
chat, err := bot.GetChat(ctx, "CHAT_ID")
fmt.Printf("نام چت: %s\n", chat.Title)
```

#### 👤 اطلاعات کاربر
```go
user, err := bot.GetUserInfo(ctx, "USER_ID")
fmt.Printf("نام کاربر: %s %s\n", user.FirstName, user.LastName)
```

#### 👥 لیست اعضا
```go
members, err := bot.GetMembers(ctx, "CHAT_ID")
for _, member := range members {
    fmt.Printf("عضو: %s\n", member.User.FirstName)
}
```

#### 🛡️ مدیران چت
```go
admins, err := bot.GetChatAdministrators(ctx, &ParsRubika.GetChatAdministratorsRequest{
    ChatID: "CHAT_ID",
})
```

### 📊 نظرسنجی‌ها

#### 📊 ایجاد نظرسنجی
```go
messageID, err := bot.CreatePoll(ctx, "CHAT_ID", "سوال نظرسنجی", []string{
    "گزینه اول",
    "گزینه دوم", 
    "گزینه سوم",
})
```

#### 🗳️ رأی دادن
```go
err := bot.VotePoll(ctx, "CHAT_ID", "MESSAGE_ID", 0) // 0 = گزینه اول
```

#### 📈 وضعیت نظرسنجی
```go
status, err := bot.GetPollStatus(ctx, "CHAT_ID", "MESSAGE_ID")
fmt.Printf("تعداد آراء: %d\n", status.TotalVote)
```

---

## 🎛 مدیریت وضعیت (State Management)

### 💾 ایجاد State Manager

```go
// 🆕 ایجاد مدیر وضعیت
stateManager := ParsRubika.NewStateManager()

// 🔄 استفاده در هندلرها
bot.OnMessageUpdates(func(ctx context.Context, update *ParsRubika.Update) error {
    userID := update.NewMessage.SenderID
    
    // 💾 ذخیره وضعیت کاربر
    stateManager.SetState(userID, "current_menu", "main")
    stateManager.SetState(userID, "selected_item", "item_123")
    stateManager.SetState(userID, "step", "2")
    
    return nil
})
```

### 💽 ذخیره وضعیت کاربر

```go
// 💾 ذخیره داده‌های مختلف کاربر
stateManager.SetState(userID, "current_menu", "main")
stateManager.SetState(userID, "selected_item", "item_123")
stateManager.SetState(userID, "step", "2")
stateManager.SetState(userID, "form_data", map[string]string{
    "name": "John",
    "age":  "30",
})
```

### 🔍 بازیابی وضعیت

```go
// 🔍 بازیابی وضعیت کاربر
menu, exists := stateManager.GetState(userID, "current_menu")
if exists {
    switch menu {
    case "main":
        // 🏠 نمایش منوی اصلی
    case "settings":
        // ⚙️ نمایش تنظیمات
    case "profile":
        // 👤 نمایش پروفایل
    }
}

step, exists := stateManager.GetState(userID, "step")
if exists {
    // 🔄 ادامه فرآیند از مرحله ذخیره شده
}
```

### 🗑️ مدیریت وضعیت

```go
// 🗑️ حذف یک کلید خاص
stateManager.DeleteState(userID, "selected_item")

// 🗑️ حذف تمام وضعیت‌های کاربر
stateManager.DeleteUserState(userID)
```

### 🎯 مثال کامل State Management

```go
func setupStatefulHandlers(bot *ParsRubika.BotClient) {
    stateManager := ParsRubika.NewStateManager()
    
    bot.OnMessageUpdates(func(ctx context.Context, update *ParsRubika.Update) error {
        if update.NewMessage == nil {
            return nil
        }
        
        userID := update.NewMessage.SenderID
        text := update.NewMessage.Text
        chatID := update.ChatID
        
        // 🔍 بررسی وضعیت فعلی کاربر
        currentState, _ := stateManager.GetState(userID, "state")
        
        switch currentState {
        case "awaiting_name":
            return handleNameInput(ctx, bot, chatID, userID, stateManager, text)
        case "awaiting_age":
            return handleAgeInput(ctx, bot, chatID, userID, stateManager, text)
        case "awaiting_email":
            return handleEmailInput(ctx, bot, chatID, userID, stateManager, text)
        default:
            return handleInitialState(ctx, bot, chatID, userID, stateManager, text)
        }
    })
}

// 🏠 وضعیت اولیه
func handleInitialState(ctx context.Context, bot *ParsRubika.BotClient, chatID, userID string, stateManager *ParsRubika.StateManager, text string) error {
    if text == "/register" {
        // 💾 ذخیره وضعیت جدید
        stateManager.SetState(userID, "state", "awaiting_name")
        
        _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: chatID,
            Text:   "👤 لطفا نام خود را وارد کنید:",
        })
        return err
    }
    return nil
}

// 👤 دریافت نام
func handleNameInput(ctx context.Context, bot *ParsRubika.BotClient, chatID, userID string, stateManager *ParsRubika.StateManager, name string) error {
    // 💾 ذخیره نام
    stateManager.SetState(userID, "name", name)
    stateManager.SetState(userID, "state", "awaiting_age")
    
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   "🎂 لطفا سن خود را وارد کنید:",
    })
    return err
}

// 🎂 دریافت سن
func handleAgeInput(ctx context.Context, bot *ParsRubika.BotClient, chatID, userID string, stateManager *ParsRubika.StateManager, age string) error {
    // 💾 ذخیره سن
    stateManager.SetState(userID, "age", age)
    stateManager.SetState(userID, "state", "awaiting_email")
    
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   "📧 لطفا ایمیل خود را وارد کنید:",
    })
    return err
}

// 📧 دریافت ایمیل
func handleEmailInput(ctx context.Context, bot *ParsRubika.BotClient, chatID, userID string, stateManager *ParsRubika.StateManager, email string) error {
    // 💾 ذخیره ایمیل
    stateManager.SetState(userID, "email", email)
    
    // 🔍 بازیابی تمام اطلاعات
    name, _ := stateManager.GetState(userID, "name")
    age, _ := stateManager.GetState(userID, "age")
    
    // 🗑️ پاک کردن وضعیت
    stateManager.DeleteUserState(userID)
    
    // ✅ ارسال نتیجه نهایی
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text: fmt.Sprintf(`✅ **ثبت نام کامل شد!**

👤 نام: %s
🎂 سن: %s
📧 ایمیل: %s

🎉 از عضویت شما متشکریم!`, name, age, email),
    })
    return err
}
```

---

## ⌨ کیبوردهای پویا

### 🎮 ایجاد کیبورد ساده

```go
func createMainMenuKeyboard() *ParsRubika.Keypad {
    return &ParsRubika.Keypad{
        Rows: []ParsRubika.KeypadRow{
            {
                Buttons: []ParsRubika.Button{
                    {
                        ID:         "profile",
                        Type:       ParsRubika.ButtonTypeSimple,
                        ButtonText: "👤 پروفایل",
                    },
                    {
                        ID:         "settings", 
                        Type:       ParsRubika.ButtonTypeSimple,
                        ButtonText: "⚙️ تنظیمات",
                    },
                },
            },
            {
                Buttons: []ParsRubika.Button{
                    {
                        ID:         "help",
                        Type:       ParsRubika.ButtonTypeSimple,
                        ButtonText: "❓ راهنما",
                    },
                    {
                        ID:         "about",
                        Type:       ParsRubika.ButtonTypeSimple,
                        ButtonText: "ℹ️ درباره ما",
                    },
                },
            },
        },
        ResizeKeyboard: true,   // 📱 تنظیم سایز برای موبایل
        OnTimeKeyboard: false,  // ⏰ نمایش موقت
    }
}
```

### 🎯 کیبوردهای پیشرفته

#### 🔘 دکمه انتخاب (Selection)
```go
selectionBtn := ParsRubika.Button{
    ID:         "selection",
    Type:       ParsRubika.ButtonTypeSelection,
    ButtonText: "📁 انتخاب آیتم",
    ButtonSelection: &ParsRubika.ButtonSelection{
        SelectionID:      "my_selection",
        SearchType:       ParsRubika.ButtonSelectionSearchLocal,
        GetType:          ParsRubika.ButtonSelectionGetLocal,
        IsMultiSelection: false,
        ColumnsCount:     "2",
        Title:            "📋 لطفا یک آیتم انتخاب کنید",
        Items: []ParsRubika.ButtonSelectionItem{
            {
                Text:     "📱 آیتم ۱",
                ImageUrl: "https://example.com/image1.jpg",
                Type:     ParsRubika.ButtonSelectionTextImgThu,
            },
            {
                Text:     "💻 آیتم ۲",
                ImageUrl: "https://example.com/image2.jpg", 
                Type:     ParsRubika.ButtonSelectionTextImgThu,
            },
        },
    },
}
```

#### 📅 دکمه تقویم
```go
calendarBtn := ParsRubika.Button{
    ID:         "calendar",
    Type:       ParsRubika.ButtonTypeCalendar,
    ButtonText: "📅 انتخاب تاریخ",
    ButtonCalendar: &ParsRubika.ButtonCalendar{
        Type:  ParsRubika.ButtonCalendarDatePersian,
        Title: "🗓️ تاریخ مورد نظر را انتخاب کنید",
    },
}
```

#### 🔢 دکمه انتخاب عدد
```go
numberPickerBtn := ParsRubika.Button{
    ID:         "number_picker",
    Type:       ParsRubika.ButtonTypeNumberPicker,
    ButtonText: "🔢 انتخاب عدد",
    ButtonNumberPicker: &ParsRubika.ButtonNumberPicker{
        MinValue:     "1",
        MaxValue:     "100",
        DefaultValue: "50",
        Title:        "🔢 عدد مورد نظر را انتخاب کنید",
    },
}
```

### 🎨 استفاده از کیبوردها

```go
// 📨 ارسال پیام با کیبورد اصلی
_, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:       chatID,
    Text:         "🎮 **منوی اصلی ربات**\n\nلطفا یک گزینه انتخاب کنید:",
    InlineKeypad: createMainMenuKeyboard(),
})

// ✏️ ویرایش کیبورد پیام موجود
err := bot.EditInlineKeypad(ctx, &ParsRubika.EditMessageKeypadRequest{
    ChatID:       chatID,
    MessageID:    messageID,
    InlineKeypad: createNewKeyboard(),
})
```

---

## 📁 مدیریت فایل‌ها

### 📤 آپلود فایل‌ها

#### 🖼️ آپلود عکس
```go
fileID, err := bot.uploadFile(ctx, "path/to/image.jpg", ParsRubika.ImageType)
```

#### 🎬 آپلود ویدیو
```go
fileID, err := bot.uploadFile(ctx, "path/to/video.mp4", ParsRubika.VideoType)
```

#### 🎵 آپلود صدا
```go
fileID, err := bot.uploadFile(ctx, "path/to/audio.mp3", ParsRubika.VoiceType)
```

#### 🎨 آپلود استیکر
```go
fileID, err := bot.uploadFile(ctx, "path/to/sticker.webp", ParsRubika.StickerType)
```

#### 📄 آپلود فایل معمولی
```go
fileID, err := bot.uploadFile(ctx, "path/to/document.pdf", ParsRubika.FileType)
```

### 📥 دانلود فایل‌ها

#### 💾 دانلود فایل معمولی
```go
err := bot.Download(ctx, "file_id", "path/to/save/file")
```

#### 🖼️ دانلود عکس پروفایل
```go
err := bot.DownloadProfilePicture(ctx, "user_id", "path/to/save/avatar.jpg")
```

### 🎯 مثال کامل آپلود و دانلود

```go
// 📤 آپلود فایل و دریافت اطلاعات کامل
func uploadAndSendFile(ctx context.Context, bot *ParsRubika.BotClient, chatID, filePath string, fileType ParsRubika.FileTypeEnum) error {
    // 📤 آپلود فایل
    file, err := bot.UploadFileDirectly(ctx, filePath, fileType)
    if err != nil {
        return fmt.Errorf("❌ خطا در آپلود فایل: %w", err)
    }
    
    // 📨 ارسال پیام با اطلاعات فایل
    infoText := fmt.Sprintf(`📁 **فایل آپلود شد!**

🏷️ نام: %s
📊 سایز: %s
🆔 شناسه: %s`, file.FileName, file.Size, file.FileID)
    
    _, err = bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   infoText,
    })
    return err
}

// 📥 دانلود و ذخیره فایل
func downloadAndSaveFile(ctx context.Context, bot *ParsRubika.BotClient, fileID, savePath string) error {
    // ℹ️ دریافت اطلاعات فایل
    fileInfo, err := bot.GetFile(ctx, fileID)
    if err != nil {
        return fmt.Errorf("❌ خطا در دریافت اطلاعات فایل: %w", err)
    }
    
    // 📥 دانلود فایل
    err = bot.Download(ctx, fileID, savePath)
    if err != nil {
        return fmt.Errorf("❌ خطا در دانلود فایل: %w", err)
    }
    
    log.Printf("✅ فایل با موفقیت دانلود شد: %s", savePath)
    return nil
}
```

---

## 🌐 Webhook & Polling

### 📡 پولینگ (Polling)

#### ⚙️ تنظیمات پیشرفته پولینگ
```go
err := bot.StartPolling(ctx, ParsRubika.PollingOptions{
    Handler:           customHandler,      // 🎯 هندلر سفارشی
    RetryTimeout:      10 * time.Second,   // ⏰ زمان انتظار برای تلاش مجدد
    Limit:             100,                // 📨 حداکثر تعداد آپدیت در هر درخواست
    AllowEmptyUpdates: false,              // ❌ عدم دریافت آپدیت‌های خالی
    PollInterval:      2 * time.Second,    // ⏱️ فاصله بین درخواست‌ها
    Timeout:           30 * time.Second,   // ⏳ تایم‌اوت درخواست
})
```

#### 🎯 مثال کامل پولینگ
```go
func startAdvancedPolling(bot *ParsRubika.BotClient) error {
    ctx := context.Background()
    
    pollingOpts := ParsRubika.PollingOptions{
        Handler: func(ctx context.Context, update *ParsRubika.Update) error {
            log.Printf("📨 آپدیت دریافت شد - نوع: %s", update.Type)
            
            // 🎯 پردازش آپدیت
            if update.NewMessage != nil {
                return handleMessage(ctx, bot, update)
            }
            return nil
        },
        RetryTimeout:      15 * time.Second,
        Limit:             50,
        PollInterval:      1 * time.Second,
        Timeout:           60 * time.Second,
    }
    
    return bot.StartPolling(ctx, pollingOpts)
}
```

### 🌐 وب‌هوک (Webhook)

#### ⚙️ تنظیمات وب‌هوک
```go
webhookOpts := ParsRubika.WebhookOptions{
    Port:    8443,                           // 🚪 پورت سرور
    Path:    "/webhook",                     // 🌐 مسیر وب‌هوک
    Handler: customHandler,                  // 🎯 هندلر سفارشی
    Secret:  "your_webhook_secret",          // 🔐 کلید امنیتی
}

err := bot.StartWebhook(ctx, webhookOpts)
```

#### 🎯 مثال کامل وب‌هوک
```go
func startWebhookServer(bot *ParsRubika.BotClient) error {
    ctx := context.Background()
    
    webhookOpts := ParsRubika.WebhookOptions{
        Port:    8080,
        Path:    "/bot-webhook",
        Handler: handleWebhookUpdate,
        Secret:  "my_super_secret_key_123",
    }
    
    log.Println("🌐 سرور وب‌هوک در حال راه‌اندازی...")
    return bot.StartWebhook(ctx, webhookOpts)
}

func handleWebhookUpdate(ctx context.Context, update *ParsRubika.Update) error {
    log.Printf("🌐 وب‌هوک دریافت شد - نوع: %s", update.Type)
    
    // 🎯 پردازش آپدیت وب‌هوک
    if update.NewMessage != nil {
        return processWebhookMessage(ctx, update)
    }
    return nil
}
```

### 🔄 مقایسه Polling و Webhook

| ویژگی | 📡 Polling | 🌐 Webhook |
|--------|------------|------------|
| **سادگی** | ✅ بسیار ساده | ⚠️ نیاز به سرور |
| **Performance** | ⚠️ متوسط | ✅ بسیار بالا |
| **Real-time** | ❌ تأخیر دارد | ✅ فوری |
| **مصرف منابع** | ❌ بالا | ✅ پایین |
| **پیکربندی** | ✅ آسان | ⚠️ پیچیده |
| **مقیاس‌پذیری** | ⚠️ محدود | ✅ نامحدود |

---

## 🚀 مثال‌های پیشرفته

### 🏪 ربات فروشگاهی

```go
type Product struct {
    ID          string
    Name        string
    Description string
    Price       int
    ImagePath   string
    Category    string
}

type ShopBot struct {
    bot          *ParsRubika.BotClient
    stateManager *ParsRubika.StateManager
    products     map[string]Product
    carts        map[string][]string // userID -> productIDs
    orders       map[string]Order    // orderID -> Order
}

type Order struct {
    ID        string
    UserID    string
    Products  []string
    Total     int
    Status    string
    CreatedAt time.Time
}

func NewShopBot(token string) *ShopBot {
    bot := &ShopBot{
        bot:          ParsRubika.NewClient(token),
        stateManager: ParsRubika.NewStateManager(),
        products:     make(map[string]Product),
        carts:        make(map[string][]string),
        orders:       make(map[string]Order),
    }
    
    bot.initializeProducts()
    bot.setupHandlers()
    return bot
}

func (sb *ShopBot) initializeProducts() {
    // 📦 محصولات نمونه
    sb.products["1"] = Product{
        ID:          "1",
        Name:        "لپ‌تاپ گیمینگ",
        Description: "لپ‌تاپ گیمینگ با کارت گرافیک RTX 4060",
        Price:       45000000,
        ImagePath:   "images/laptop.jpg",
        Category:    "الکترونیک",
    }
    
    sb.products["2"] = Product{
        ID:          "2", 
        Name:        "هدفون بی‌سیم",
        Description: "هدفون بی‌سیم با نویزکنسلینگ",
        Price:       3500000,
        ImagePath:   "images/headphone.jpg",
        Category:    "الکترونیک",
    }
    
    sb.products["3"] = Product{
        ID:          "3",
        Name:        "کتاب Go Programming",
        Description: "کتاب جامع برنامه‌نویسی Go",
        Price:       150000,
        ImagePath:   "images/book.jpg", 
        Category:    "کتاب",
    }
}

func (sb *ShopBot) setupHandlers() {
    sb.bot.OnMessageUpdates(func(ctx context.Context, update *ParsRubika.Update) error {
        if update.NewMessage == nil {
            return nil
        }
        
        userID := update.NewMessage.SenderID
        text := update.NewMessage.Text
        
        switch {
        case text == "/start":
            return sb.showMainMenu(ctx, update.ChatID)
        case text == "/products":
            return sb.showProducts(ctx, update.ChatID)
        case text == "/cart":
            return sb.showCart(ctx, update.ChatID, userID)
        case text == "/orders":
            return sb.showOrders(ctx, update.ChatID, userID)
        case strings.HasPrefix(text, "/add_"):
            productID := strings.TrimPrefix(text, "/add_")
            return sb.addToCart(ctx, update.ChatID, userID, productID)
        case text == "/checkout":
            return sb.startCheckout(ctx, update.ChatID, userID)
        default:
            return sb.showMainMenu(ctx, update.ChatID)
        }
    })
}

func (sb *ShopBot) showMainMenu(ctx context.Context, chatID string) error {
    menuText := `🛍️ **فروشگاه آنلاین**

🎯 **دستورات اصلی:**
/products - مشاهده محصولات
/cart - سبد خرید
/orders - سفارشات من
/help - راهنما

💡 **برای شروع خرید از /products استفاده کنید**`
    
    _, err := sb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   menuText,
    })
    return err
}

func (sb *ShopBot) showProducts(ctx context.Context, chatID string) error {
    var productsText strings.Builder
    productsText.WriteString("🛍️ **محصولات موجود**\n\n")
    
    for id, product := range sb.products {
        productsText.WriteString(fmt.Sprintf(
            "📦 **%s**\n💵 قیمت: %s تومان\n🔹 %s\n\n🛒 اضافه به سبد: /add_%s\n\n%s\n",
            product.Name,
            humanize.Comma(int64(product.Price)),
            product.Description,
            id,
            strings.Repeat("─", 30),
        ))
    }
    
    _, err := sb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   productsText.String(),
    })
    return err
}

func (sb *ShopBot) addToCart(ctx context.Context, chatID, userID, productID string) error {
    product, exists := sb.products[productID]
    if !exists {
        _, err := sb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: chatID,
            Text:   "❌ محصول مورد نظر یافت نشد!",
        })
        return err
    }
    
    if sb.carts[userID] == nil {
        sb.carts[userID] = []string{}
    }
    sb.carts[userID] = append(sb.carts[userID], productID)
    
    _, err := sb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   fmt.Sprintf("✅ **%s** به سبد خرید اضافه شد! 🛒", product.Name),
    })
    return err
}

func (sb *ShopBot) showCart(ctx context.Context, chatID, userID string) error {
    cartItems := sb.carts[userID]
    if len(cartItems) == 0 {
        _, err := sb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: chatID,
            Text:   "🛒 سبد خرید شما خالی است!",
        })
        return err
    }
    
    var cartText strings.Builder
    cartText.WriteString("🛒 **سبد خرید شما**\n\n")
    
    total := 0
    itemCount := make(map[string]int)
    
    for _, productID := range cartItems {
        itemCount[productID]++
    }
    
    for productID, count := range itemCount {
        product := sb.products[productID]
        itemTotal := product.Price * count
        total += itemTotal
        
        cartText.WriteString(fmt.Sprintf(
            "📦 %s (تعداد: %d)\n💵 %s تومان\n\n",
            product.Name,
            count,
            humanize.Comma(int64(itemTotal)),
        ))
    }
    
    cartText.WriteString(fmt.Sprintf(
        "💰 **جمع کل: %s تومان**\n\n💳 برای تسویه حساب: /checkout",
        humanize.Comma(int64(total)),
    ))
    
    _, err := sb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   cartText.String(),
    })
    return err
}

func (sb *ShopBot) startCheckout(ctx context.Context, chatID, userID string) error {
    cartItems := sb.carts[userID]
    if len(cartItems) == 0 {
        _, err := sb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: chatID,
            Text:   "❌ سبد خرید شما خالی است!",
        })
        return err
    }
    
    // 🧮 محاسبه جمع کل
    total := 0
    for _, productID := range cartItems {
        product := sb.products[productID]
        total += product.Price
    }
    
    // 📝 ایجاد سفارش
    orderID := fmt.Sprintf("ORD-%d", time.Now().Unix())
    order := Order{
        ID:        orderID,
        UserID:    userID,
        Products:  cartItems,
        Total:     total,
        Status:    "pending",
        CreatedAt: time.Now(),
    }
    sb.orders[orderID] = order
    
    // 🗑️ خالی کردن سبد خرید
    delete(sb.carts, userID)
    
    // ✅ ارسال تأیید سفارش
    _, err := sb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text: fmt.Sprintf(`✅ **سفارش شما ثبت شد!**

🆔 کد سفارش: %s
💰 مبلغ قابل پرداخت: %s تومان
📦 تعداد آیتم: %d
📞 برای پیگیری با پشتیبانی تماس بگیرید.`,
            orderID,
            humanize.Comma(int64(total)),
            len(cartItems),
        ),
    })
    return err
}
```

### 📊 ربات نظرسنجی و آمار

```go
type PollBot struct {
    bot          *ParsRubika.BotClient
    stateManager *ParsRubika.StateManager
    activePolls  map[string]*PollData    // messageID -> PollData
    userVotes    map[string]map[string]int // userID -> map[pollID]optionIndex
}

type PollData struct {
    ID       string
    Question string
    Options  []string
    Votes    map[int]int // optionIndex -> voteCount
    Creator  string
    CreatedAt time.Time
}

func NewPollBot(token string) *PollBot {
    bot := &PollBot{
        bot:          ParsRubika.NewClient(token),
        stateManager: ParsRubika.NewStateManager(),
        activePolls:  make(map[string]*PollData),
        userVotes:    make(map[string]map[string]int),
    }
    
    bot.setupHandlers()
    return bot
}

func (pb *PollBot) setupHandlers() {
    pb.bot.OnMessageUpdates(func(ctx context.Context, update *ParsRubika.Update) error {
        if update.NewMessage == nil {
            return nil
        }
        
        text := update.NewMessage.Text
        chatID := update.ChatID
        userID := update.NewMessage.SenderID
        
        switch {
        case strings.HasPrefix(text, "/create_poll"):
            return pb.handleCreatePoll(ctx, chatID, userID, text)
        case strings.HasPrefix(text, "/poll_stats"):
            return pb.handlePollStats(ctx, chatID, text)
        case update.NewMessage.Poll != nil:
            return pb.handleVote(ctx, update)
        default:
            return pb.showPollMenu(ctx, chatID)
        }
    })
}

func (pb *PollBot) handleCreatePoll(ctx context.Context, chatID, userID, text string) error {
    // 📝 پارس کردن دستور: /create_poll سوال | گزینه۱ | گزینه۲ | گزینه۳
    parts := strings.Split(strings.TrimPrefix(text, "/create_poll "), "|")
    if len(parts) < 3 {
        _, err := pb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: chatID,
            Text:   "❌ فرمت دستور نامعتبر!\n\n💡 مثال:\n/create_poll بهترین زبان برنامه‌نویسی؟ | Go | Python | JavaScript",
        })
        return err
    }
    
    question := strings.TrimSpace(parts[0])
    options := make([]string, len(parts)-1)
    
    for i, part := range parts[1:] {
        options[i] = strings.TrimSpace(part)
    }
    
    // 📊 ایجاد نظرسنجی
    messageID, err := pb.bot.CreatePoll(ctx, chatID, question, options)
    if err != nil {
        return err
    }
    
    // 💾 ذخیره اطلاعات نظرسنجی
    pb.activePolls[messageID] = &PollData{
        ID:        messageID,
        Question:  question,
        Options:   options,
        Votes:     make(map[int]int),
        Creator:   userID,
        CreatedAt: time.Now(),
    }
    
    _, err = pb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   fmt.Sprintf("✅ نظرسنجی با موفقیت ایجاد شد!\n\n📊 برای مشاهده آمار: /poll_stats_%s", messageID),
    })
    return err
}

func (pb *PollBot) handleVote(ctx context.Context, update *ParsRubika.Update) error {
    if update.NewMessage.Poll == nil {
        return nil
    }
    
    userID := update.NewMessage.SenderID
    messageID := strconv.FormatInt(update.NewMessage.MessageID, 10)
    
    poll, exists := pb.activePolls[messageID]
    if !exists {
        return nil
    }
    
    // 🗳️ ثبت رأی کاربر
    selectedOption := update.NewMessage.Poll.SelectionIndex
    
    if pb.userVotes[userID] == nil {
        pb.userVotes[userID] = make(map[string]int)
    }
    
    // 🔄 اگر کاربر قبلاً رأی داده، رأی قبلی را حذف کن
    if prevVote, exists := pb.userVotes[userID][messageID]; exists {
        poll.Votes[prevVote]--
    }
    
    // ➕ ثبت رأی جدید
    poll.Votes[selectedOption]++
    pb.userVotes[userID][messageID] = selectedOption
    
    log.Printf("✅ کاربر %s به گزینه %d رأی داد", userID, selectedOption)
    return nil
}

func (pb *PollBot) handlePollStats(ctx context.Context, chatID, text string) error {
    messageID := strings.TrimPrefix(text, "/poll_stats_")
    
    poll, exists := pb.activePolls[messageID]
    if !exists {
        _, err := pb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: chatID,
            Text:   "❌ نظرسنجی مورد نظر یافت نشد!",
        })
        return err
    }
    
    // 📈 محاسبه آمار
    totalVotes := 0
    for _, count := range poll.Votes {
        totalVotes += count
    }
    
    var statsText strings.Builder
    statsText.WriteString(fmt.Sprintf("📊 **آمار نظرسنجی**\n\n❓ %s\n\n", poll.Question))
    
    for i, option := range poll.Options {
        voteCount := poll.Votes[i]
        percentage := 0
        if totalVotes > 0 {
            percentage = (voteCount * 100) / totalVotes
        }
        
        // 📊 ایجاد نمودار پیشرفت
        progressBar := createProgressBar(percentage, 20)
        
        statsText.WriteString(fmt.Sprintf(
            "🔹 %s\n%s %d%% (%d رأی)\n\n",
            option,
            progressBar,
            percentage,
            voteCount,
        ))
    }
    
    statsText.WriteString(fmt.Sprintf("👥 **کل آراء:** %d\n⏰ **زمان ایجاد:** %s", 
        totalVotes, 
        poll.CreatedAt.Format("2006/01/02 15:04"),
    ))
    
    _, err := pb.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   statsText.String(),
    })
    return err
}

func createProgressBar(percentage, length int) string {
    filled := (percentage * length) / 100
    empty := length - filled
    
    bar := "🟦" + strings.Repeat("🟦", filled) + strings.Repeat("⬜", empty)
    return bar
}
```

### 🔧 ربات مدیریت کانال

```go
type ChannelManagerBot struct {
    bot          *ParsRubika.BotClient
    stateManager *ParsRubika.StateManager
    adminUsers   map[string]bool
    channelStats map[string]ChannelStats
}

type ChannelStats struct {
    MemberCount  int
    MessageCount int
    ActiveUsers  map[string]int
    LastActivity time.Time
}

func NewChannelManagerBot(token string) *ChannelManagerBot {
    bot := &ChannelManagerBot{
        bot:          ParsRubika.NewClient(token),
        stateManager: ParsRubika.NewStateManager(),
        adminUsers:   make(map[string]bool),
        channelStats: make(map[string]ChannelStats),
    }
    
    // 👥 تنظیم ادمین‌ها
    bot.adminUsers["ADMIN_USER_ID_1"] = true
    bot.adminUsers["ADMIN_USER_ID_2"] = true
    
    bot.setupHandlers()
    return bot
}

func (cm *ChannelManagerBot) setupHandlers() {
    cm.bot.OnMessageUpdates(func(ctx context.Context, update *ParsRubika.Update) error {
        if update.NewMessage == nil {
            return nil
        }
        
        userID := update.NewMessage.SenderID
        chatID := update.ChatID
        text := update.NewMessage.Text
        
        // 🔐 بررسی دسترسی ادمین
        if !cm.adminUsers[userID] {
            _, err := cm.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
                ChatID: chatID,
                Text:   "⛔ شما دسترسی لازم برای استفاده از این ربات را ندارید!",
            })
            return err
        }
        
        // 🎯 مسیریابی دستورات ادمین
        switch {
        case strings.HasPrefix(text, "/broadcast"):
            return cm.handleBroadcast(ctx, chatID, userID, text)
        case text == "/stats":
            return cm.handleStats(ctx, chatID)
        case strings.HasPrefix(text, "/post"):
            return cm.handlePost(ctx, chatID, text)
        case strings.HasPrefix(text, "/pin"):
            return cm.handlePinMessage(ctx, chatID, text)
        case strings.HasPrefix(text, "/ban"):
            return cm.handleBanUser(ctx, chatID, text)
        default:
            return cm.showAdminMenu(ctx, chatID)
        }
    })
}

func (cm *ChannelManagerBot) showAdminMenu(ctx context.Context, chatID string) error {
    menuText := `🛠️ **پنل مدیریت کانال**

🎯 **دستورات مدیریتی:**
/broadcast [متن] - ارسال پیام همگانی
/stats - آمار کانال
/post [متن] - ارسال پست در کانال
/pin [message_id] - پین کردن پیام
/ban [user_id] - مسدود کردن کاربر

💡 **برای استفاده، دستور مورد نظر را وارد کنید.**`
    
    _, err := cm.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   menuText,
    })
    return err
}

func (cm *ChannelManagerBot) handleBroadcast(ctx context.Context, chatID, userID, text string) error {
    broadcastText := strings.TrimSpace(strings.TrimPrefix(text, "/broadcast"))
    
    if broadcastText == "" {
        _, err := cm.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: chatID,
            Text:   "❌ لطفا متن پیام همگانی را وارد کنید!\n\n💡 مثال:\n/broadcast سلام به همه کاربران عزیز! 👋",
        })
        return err
    }
    
    // 📨 ارسال پیام همگانی به کاربران
    // در اینجا باید لیست کاربران از دیتابیس خوانده شود
    users := []string{"USER_1", "USER_2", "USER_3"} // کاربران نمونه
    
    successCount := 0
    for _, userID := range users {
        _, err := cm.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: userID,
            Text:   fmt.Sprintf("📢 **پیام همگانی:**\n\n%s", broadcastText),
        })
        
        if err != nil {
            log.Printf("❌ خطا در ارسال به کاربر %s: %v", userID, err)
        } else {
            successCount++
        }
        
        // ⏳ تأخیر برای جلوگیری از محدودیت rate
        time.Sleep(100 * time.Millisecond)
    }
    
    // 📊 گزارش نتیجه
    _, err := cm.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text: fmt.Sprintf("✅ **گزارش ارسال همگانی**

📤 ارسال شده به: %d کاربر
❌ ناموفق: %d کاربر
📝 متن پیام: %s",
            successCount,
            len(users)-successCount,
            broadcastText,
        ),
    })
    return err
}

func (cm *ChannelManagerBot) handleStats(ctx context.Context, chatID string) error {
    // 📊 جمع‌آوری آمار از کانال‌ها
    var statsText strings.Builder
    statsText.WriteString("📊 **آمار کانال‌ها**\n\n")
    
    for channelID, stats := range cm.channelStats {
        channelInfo, err := cm.bot.GetChat(ctx, channelID)
        if err != nil {
            continue
        }
        
        statsText.WriteString(fmt.Sprintf(
            "📢 **%s**\n👥 اعضا: %d\n💬 پیام‌ها: %d\n🕒 آخرین فعالیت: %s\n\n%s\n",
            channelInfo.Title,
            stats.MemberCount,
            stats.MessageCount,
            stats.LastActivity.Format("15:04"),
            strings.Repeat("─", 30),
        ))
    }
    
    _, err := cm.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   statsText.String(),
    })
    return err
}
```

---

## ☁ استقرار و دیپلوی

### 🐳 داکرایز کردن

#### Dockerfile
```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 📦 کپی فایل‌های وابستگی
COPY go.mod go.sum ./
RUN go mod download

# 🏗 کپی سورس کد و ساخت
COPY . .
RUN go build -o rubika-bot .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

# 📋 کپی باینری ساخته شده
COPY --from=builder /app/rubika-bot .

# 🔑 کپی فایل‌های پیکربندی
COPY config.yaml ./

# 🚪 اکسپوز پورت
EXPOSE 8080

# 🚀 اجرای ربات
CMD ["./rubika-bot"]
```

#### docker-compose.yml
```yaml
version: '3.8'

services:
  rubika-bot:
    build: .
    ports:
      - "8080:8080"
    environment:
      - RUBIKA_BOT_TOKEN=${RUBIKA_BOT_TOKEN}
      - DATABASE_URL=${DATABASE_URL}
    volumes:
      - ./logs:/app/logs
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3

  # 🗄️ دیتابیس اختیاری
  postgres:
    image: postgres:13
    environment:
      - POSTGRES_DB=rubika_bot
      - POSTGRES_USER=bot_user
      - POSTGRES_PASSWORD=${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    restart: unless-stopped

volumes:
  postgres_data:
```

### 🚀 استقرار روی سرور

#### 1. تنظیمات systemd
```ini
# /etc/systemd/system/rubika-bot.service
[Unit]
Description=Rubika Bot
After=network.target
Requires=network.target

[Service]
Type=simple
User=botuser
Group=botuser
WorkingDirectory=/home/botuser/rubika-bot
Environment=RUBIKA_BOT_TOKEN=your_bot_token_here
Environment=DATABASE_URL=postgresql://user:pass@localhost:5432/rubika_bot
ExecStart=/home/botuser/rubika-bot/rubika-bot
Restart=always
RestartSec=10
StandardOutput=journal
StandardError=journal
SyslogIdentifier=rubika-bot

# 🔒 امنیت
NoNewPrivileges=yes
PrivateTmp=yes
ProtectSystem=strict
ProtectHome=yes

[Install]
WantedBy=multi-user.target
```

#### 2. اسکریپت استقرار
```bash
#!/bin/bash
# deploy.sh

set -e

echo "🚀 شروع استقرار ربات..."

# 🔄 توقف سرویس فعلی
sudo systemctl stop rubika-bot || true

# 📥 دریافت آخرین تغییرات
git pull origin main

# 📦 ساخت پروژه
go build -o rubika-bot .

# 🔒 تنظیم مجوزها
chmod +x rubika-bot

# 🗄️ migrate دیتابیس (اگر نیاز باشد)
# ./rubika-bot migrate

# 🚀 راه‌اندازی سرویس
sudo systemctl daemon-reload
sudo systemctl start rubika-bot
sudo systemctl enable rubika-bot

echo "✅ استقرار با موفقیت انجام شد!"
echo "📊 وضعیت سرویس:"
sudo systemctl status rubika-bot
```

### ☁️ استقرار ابری (AWS/GCP)

#### Dockerfile برای ابر
```dockerfile
FROM golang:1.21 as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

CMD ["./main"]
```

#### Kubernetes Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: rubika-bot
  labels:
    app: rubika-bot
spec:
  replicas: 2
  selector:
    matchLabels:
      app: rubika-bot
  template:
    metadata:
      labels:
        app: rubika-bot
    spec:
      containers:
      - name: rubika-bot
        image: your-registry/rubika-bot:latest
        ports:
        - containerPort: 8080
        env:
        - name: RUBIKA_BOT_TOKEN
          valueFrom:
            secretKeyRef:
              name: bot-secrets
              key: token
        resources:
          requests:
            memory: "64Mi"
            cpu: "50m"
          limits:
            memory: "128Mi"
            cpu: "100m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
```

---

## 🔧 عیب‌یابی

### 🚨 خطاهای رایج

#### 1. خطای اتصال
```go
// 🔌 مشکل: عدم اتصال به API روبیکا
func handleConnectionError(err error) {
    if strings.Contains(err.Error(), "connection refused") {
        log.Println("❌ خطای اتصال: سرور روبیکا در دسترس نیست")
        log.Println("💡 راه‌حل: اینترنت و فایروال را بررسی کنید")
    } else if strings.Contains(err.Error(), "timeout") {
        log.Println("⏰ خطای تایم‌اوت: پاسخ سرور طولانی شد")
        log.Println("💡 راه‌حل: تایم‌اوت را افزایش دهید")
    }
}
```

#### 2. خطای توکن
```go
// 🔑 مشکل: توکن نامعتبر
func handleTokenError(err error) {
    if strings.Contains(err.Error(), "unauthorized") {
        log.Println("❌ خطای احراز هویت: توکن نامعتبر است")
        log.Println("💡 راه‌حل:")
        log.Println("  1. توکن را از @rubika_bot دریافت کنید")
        log.Println("  2. مطمئن شوید ربات فعال است")
        log.Println("  3. توکن را درست کپی کنید")
    }
}
```

#### 3. محدودیت نرخ ارسال
```go
// 🚦 مشکل: ارسال درخواست‌های زیاد
func handleRateLimit(bot *ParsRubika.BotClient) {
    // ⏳ افزایش تأخیر بین درخواست‌ها
    bot = ParsRubika.NewClient(token,
        ParsRubika.WithRateLimitDelay(2 * time.Second),
        ParsRubika.WithMaxRetries(5),
    )
}
```

### 📊 مانیتورینگ و سلامت

#### سلامت‌سنجی
```go
func startHealthCheck(bot *ParsRubika.BotClient) {
    ticker := time.NewTicker(5 * time.Minute)
    
    go func() {
        for range ticker.C {
            ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
            
            // 🔍 بررسی وضعیت ربات
            _, err := bot.GetMe(ctx)
            if err != nil {
                log.Printf("⚠️ بررسی سلامت ناموفق: %v", err)
                
                // 🔄 تلاش برای بازیابی
                if strings.Contains(err.Error(), "token") {
                    log.Println("🔄 تلاش برای بازنشانی توکن...")
                    // منطق بازنشانی توکن
                }
            } else {
                log.Println("✅ ربات در وضعیت سالم")
            }
            
            cancel()
        }
    }()
}
```

#### میدلور لاگینگ
```go
type LoggingMiddleware struct {
    bot *ParsRubika.BotClient
}

func (lm *LoggingMiddleware) HandleUpdate(ctx context.Context, update *ParsRubika.Update, next ParsRubika.HandlerFunc) error {
    start := time.Now()
    
    log.Printf("📥 آپدیت دریافت شد - نوع: %s, چت: %s", update.Type, update.ChatID)
    
    err := next(ctx, update)
    
    duration := time.Since(start)
    if err != nil {
        log.Printf("❌ خطا در پردازش آپدیت پس از %v: %v", duration, err)
        
        // 📧 ارسال گزارش خطا (اختیاری)
        go lm.sendErrorReport(ctx, update, err)
    } else {
        log.Printf("✅ آپدیت با موفقیت پردازش شد در %v", duration)
    }
    
    return err
}

func (lm *LoggingMiddleware) sendErrorReport(ctx context.Context, update *ParsRubika.Update, err error) {
    // 📨 ارسال گزارش خطا به ادمین
    errorMsg := fmt.Sprintf("🚨 **گزارش خطا**\n\n💬 چت: %s\n❌ خطا: %v\n⏰ زمان: %s",
        update.ChatID, err, time.Now().Format("2006/01/02 15:04:05"))
    
    // ارسال به ادمین‌ها
    for adminID := range lm.bot.adminUsers {
        lm.bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: adminID,
            Text:   errorMsg,
        })
    }
}
```

### 🛠 ابزارهای توسعه

#### اسکریپت توسعه
```bash
#!/bin/bash
# dev.sh - اسکریپت توسعه

echo "🔧 محیط توسعه ParsRubika"

case "$1" in
    "run")
        echo "🚀 اجرای ربات در حالت توسعه..."
        go run main.go handlers.go
        ;;
    "test")
        echo "🧪 اجرای تست‌ها..."
        go test -v ./...
        ;;
    "build")
        echo "🏗 ساخت باینری..."
        go build -o rubika-bot .
        ;;
    "lint")
        echo "📝 بررسی کد..."
        golangci-lint run
        ;;
    "clean")
        echo "🧹 پاک‌سازی..."
        go clean
        rm -f rubika-bot
        ;;
    *)
        echo "💡 استفاده: ./dev.sh [run|test|build|lint|clean]"
        ;;
esac
```

#### پیکربندی محیط
```yaml
# config.yaml
bot:
  token: "${RUBIKA_BOT_TOKEN}"
  webhook:
    enabled: false
    port: 8080
    path: "/webhook"
  polling:
    enabled: true
    interval: "2s"
    limit: 100

database:
  url: "${DATABASE_URL}"
  max_connections: 20

logging:
  level: "info"
  file: "bot.log"

features:
  state_management: true
  file_upload: true
  admin_panel: true
```

---

## 📞 پشتیبانی

### 🔗 ارتباط با توسعه‌دهنده

- **👤 ایدی روبیکا:** `NinjaCode`
- **📢 چنل روبیکا:** `Ninja_code`
- **📧 ایمیل:** `ninjacode.ir@gmail.com`
- **🐙 گیت‌هاب:** [Abolfazl-Zarei](https://github.com/Abolfazl-Zarei)

### 📝 گزارش مشکل

```bash
# 1. بررسی نسخه‌ها
go version
go list -m all | grep ParsRubika

# 2. لاگ‌های خطا
tail -f bot.log

# 3. اطلاعات سیستم
uname -a
```

### 🤝 مشارکت در توسعه

```bash
# 1. Fork کردن پروژه
# از طریق گیت‌هاب پروژه را Fork کنید

# 2. کلون کردن
git clone https://github.com/YOUR_USERNAME/ParsRubika-bot-go.git
cd ParsRubika-bot-go

# 3. ایجاد branch جدید
git checkout -b feature/amazing-feature

# 4. کامیت تغییرات
git add .
git commit -m "اضافه کردن قابلیت جدید"

# 5. Push کردن
git push origin feature/amazing-feature

# 6. ایجاد Pull Request
```

---

## 🎉 نتیجه‌گیری

**ParsRubika** یک کتابخانه **کامل**، **قدرتمند** و **کاربردی** برای ساخت ربات‌های روبیکا با زبان Go است. با استفاده از این مستندات می‌توانید:

### ✅ آنچه یاد گرفتید:

- 🏗 **ساختار کلی کتابخانه** و مفاهیم پایه
- ⚡ **نصب و راه‌اندازی** سریع و آسان
- 🛠 **آموزش قدم به قدم** از مبتدی تا پیشرفته
- 📡 **API Reference کامل** برای تمام متدها
- 🎛 **مدیریت وضعیت** پیشرفته کاربران
- ⌨ **کیبوردهای پویا** و تعاملی
- 📁 **مدیریت فایل‌ها** و مدیا
- 🌐 **Webhook و Polling** برای دریافت آپدیت‌ها
- 🚀 **مثال‌های پیشرفته** واقعی
- ☁ **استقرار و دیپلوی** حرفه‌ای
- 🔧 **عیب‌یابی** و مانیتورینگ

### 🚀 شروع نهایی

```go
package main

import (
    "context"
    "log"
    "os"
    
    ParsRubika "github.com/Abolfazl-Zarei/ParsRubika-bot-go"
)

func main() {
    // 🤖 ایجاد ربات
    bot := ParsRubika.NewClient(os.Getenv("RUBIKA_BOT_TOKEN"))
    
    // 🎯 هندلر ساده
    bot.OnMessageUpdates(func(ctx context.Context, update *ParsRubika.Update) error {
        if update.NewMessage != nil {
            _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
                ChatID: update.ChatID,
                Text:   "👋 سلام! من با ParsRubika ساخته شده‌ام! 🚀",
            })
            return err
        }
        return nil
    })
    
    // 🚀 اجرا
    ctx := context.Background()
    log.Println("🎉 ربات شروع به کار کرد...")
    bot.Run(ctx)
}
```

### 📚 منابع بیشتر

- 📖 [مستندات رسمی روبیکا](https://rubika.ir/docs)
- 💻 [مخزن گیت‌هاب](https://github.com/Abolfazl-Zarei/ParsRubika-bot-go)
- 🐦 [کانال اطلاع‌رسانی](https://rubika.ir/Ninja_code)
- 📝 [ایسوها و باگ‌ها](https://github.com/Abolfazl-Zarei/ParsRubika-bot-go/issues)

---

<div align="center">

## 🎯 **همین حالا شروع کنید!**

**با ParsRubika، ربات‌های قدرتمند روبیکا بسازید 🚀**

![Footer](https://img.icons8.com/color/96/000000/hearts.png) 

**ساخته شده با ❤️ توسط ابوالفضل زارعی (NinjaCode)**

📧 [ninjacode.ir@gmail.com](mailto:ninjacode.ir@gmail.com) | 
🐙 [گیت‌هاب](https://github.com/Abolfazl-Zarei) |
📱 [روبیکا](https://rubika.ir/NinjaCode)

</div>
