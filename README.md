

# 🤖 ParsRubika - کتابخانه کامل Go برای ربات‌های روبیکا

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
| **🔄 Hot-Reload** | بارگذاری مجدد کد بدون توقف |
| **🛡️ Anti-Spam** | جلوگیری از اسپم کاربران |
| **🌐 Network Stability** | مدیریت پایداری شبکه |

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
├── 🛡️ antispam.go        # 🔒 سیستم ضد اسپم
├── 🔄 reload.go          # 🔄 مدیریت Hot-Reload
├── 🌐 network.go         # 🌐 مدیریت پایداری شبکه
├── 📝 formatting.go      # 📝 فرمت‌بندی پیام‌ها و کیبوردها
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
    CallbackQuery  UpdateTypeEnum = "CallbackQuery"   // 🔘 کوئری دکمه
    InlineQuery    UpdateTypeEnum = "InlineQuery"     // 🔍 کوئری اینلاین
)
```

### 💌 ساختار پیام

```go
// 📝 ساختار کامل یک پیام
type Message struct {
    MessageID        string            `json:"message_id"`         // 🔢 شناسه پیام
    Text             string            `json:"text"`               // 📄 متن پیام
    Time             string            `json:"time"`               // ⏰ زمان ارسال
    IsEdited         bool              `json:"is_edited"`          // ✏️ آیا ویرایش شده
    SenderType       MessageSenderEnum `json:"sender_type"`        // 👤 نوع فرستنده
    SenderID         string            `json:"sender_id"`          // 🆔 شناسه فرستنده
    AuxData          *AuxData          `json:"aux_data"`           // 🔗 داده‌های کمکی
    File             *File             `json:"file"`               // 📁 فایل پیام
    ReplyToMessageID string            `json:"reply_to_message_id"`// ↩️ پاسخ به پیام
    ForwardedFrom    *ForwardedFrom    `json:"forwarded_from"`     // ↪️ اطلاعات فوروارد
    Location         *Location         `json:"location"`           // 📍 موقعیت مکانی
    Sticker          *Sticker          `json:"sticker"`            // 🎨 استیکر
    ContactMessage   *ContactMessage   `json:"contact_message"`    // 👥 مخاطب
    Poll             *Poll             `json:"poll"`               // 📊 نظرسنجی
    Payment          *PaymentStatus    `json:"payment"`            // 💰 وضعیت پرداخت
}
```

### 🤖 ساختار کلاینت

```go
// 🤖 ساختار اصلی کلاینت بات
type BotClient struct {
    token        string
    baseURL      string
    httpClient   *http.Client
    botID        string
    mu           sync.RWMutex
    lastSentTime time.Time

    // مدیریت وضعیت
    isRunning bool
    stopChan  chan struct{}

    // هندلرها و میدلورها
    handlers      []Handler
    middlewares   []MiddlewareFunc
    updateFilters []FilterFunc

    // نرخ محدودیت
    rateLimitDuration time.Duration
    maxRetries        int

    // قابلیت‌های جدید
    antiSpam                *AntiSpam
    hotReloadEnabled        bool
    stateManager            *StateManager
    ignoreTimeout           bool
    metadata                map[string]interface{}
    reloadManager           *ReloadManager
    networkStabilityManager *NetworkStabilityManager
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
    
    // 🤖 ایجاد نمونه ربات با گزینه‌های پیشرفته
    bot := ParsRubika.NewClient(botToken,
        ParsRubika.WithRateLimitDelay(1*time.Second), // تأخیر بین درخواست‌ها
        ParsRubika.WithMaxRetries(3),                 // حداکثر تلاش مجدد
        ParsRubika.WithIgnoreTimeout(true),           // نادیده گرفتن خطاهای timeout
        ParsRubika.WithHotReload(true),               // فعال‌سازی Hot-Reload
    )
    
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
    "time"
    
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
    
    // 🤖 ایجاد نمونه ربات با تمام قابلیت‌ها
    bot := ParsRubika.NewClient(token,
        ParsRubika.WithRateLimitDelay(1*time.Second),
        ParsRubika.WithMaxRetries(3),
        ParsRubika.WithIgnoreTimeout(true),
        ParsRubika.WithHotReload(true),
    )
    
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
        userID := msg.SenderID
        
        // 🛡️ بررسی ضد اسپم
        if !bot.CheckAntiSpam(userID) {
            _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
                ChatID: chatID,
                Text:   "⚠️ لطفاً کمی صبر کنید و درخواست‌های خود را با فاصله ارسال کنید.",
            })
            return err
        }
        
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
        case msg.Text == "/state":
            return handleStateDemo(ctx, bot, chatID, userID)
        default:
            return handleDefault(ctx, bot, chatID, msg.Text)
        }
    })
    
    // 🔔 هندلر شروع رات
    bot.OnStart(func(ctx context.Context, update *ParsRubika.Update) error {
        log.Println("✅ ربات با موفقیت راه‌اندازی شد!")
        return nil
    })
    
    // 🎨 هندلر برای پیام‌های عکس
    bot.OnPhoto(func(ctx context.Context, update *ParsRubika.Update) error {
        _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: update.ChatID,
            Text:   "🖼️ عکس زیبایی ارسال کردید! 📸",
        })
        return err
    })
    
    // 🎵 هندلر برای پیام‌های صوتی
    bot.OnAudio(func(ctx context.Context, update *ParsRubika.Update) error {
        _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: update.ChatID,
            Text:   "🎵 موسیقی زیبایی است! 🎶",
        })
        return err
    })
    
    // 📍 هندلر برای پیام‌های موقعیت مکانی
    bot.OnLocation(func(ctx context.Context, update *ParsRubika.Update) error {
        _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: update.ChatID,
            Text:   "📍 مکان شما دریافت شد! 🗺️",
        })
        return err
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
/state - نمایش وضعیت فعلی شما

🔧 **ساخته شده با:** 
• زبان Go 🦫
• کتابخانه ParsRubika 🚀
• توسط NinjaCode 👨‍💻

💡 **شروع کنید:** یک دستور ارسال کنید یا پیام دلخواه بفرستید!`
    
    // 📨 ارسال پیام با کیبورد
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   welcomeText,
        ReplyKeyboardMarkup: createMainMenuKeyboard(),
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
🔹 /state - نمایش وضعیت فعلی

🛠 **امکانات پیشرفته:**
• 📨 ارسال پیام‌های متنی
• 🖼 ارسال عکس و مدیا
• 📍 ارسال موقعیت مکانی
• 📊 ایجاد نظرسنجی
• ⌨️ کیبوردهای تعاملی
• 🛡️ سیستم ضد اسپم
• 💾 مدیریت وضعیت

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

// 💾 هندلر دستور /state
func handleStateDemo(ctx context.Context, bot *ParsRubika.BotClient, chatID, userID string) error {
    // 💾 ذخیره وضعیت کاربر
    bot.SetState(userID, "last_command", "/state")
    bot.SetState(userID, "visit_count", 1)
    
    // 🔍 بازیابی وضعیت کاربر
    lastCmd, _ := bot.GetState(userID, "last_command")
    visitCount, _ := bot.GetState(userID, "visit_count")
    
    // 🔄 افزایش تعداد بازدید
    if count, ok := visitCount.(int); ok {
        bot.SetState(userID, "visit_count", count+1)
        visitCount = count + 1
    }
    
    infoText := fmt.Sprintf(`💾 **وضعیت فعلی شما**

🔧 **آخرین دستور:** %s
🔢 **تعداد بازدید:** %v
⏰ **زمان آخرین فعالیت:** %s

💡 **این اطلاعات در سرور ذخیره شده‌اند و در جلسات بعدی باقی می‌مانند.**`,
        lastCmd,
        visitCount,
        time.Now().Format("2006/01/02 15:04:05"),
    )
    
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
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

// 🔧 ایجاد کیبورد اصلی
func createMainMenuKeyboard() *ParsRubika.ReplyKeyboardMarkup {
    return &ParsRubika.ReplyKeyboardMarkup{
        Keyboard: [][]*ParsRubika.KeyboardButton{
            {
                {
                    Text: "📚 راهنما",
                },
                {
                    Text: "👤 اطلاعات",
                },
            },
            {
                {
                    Text: "🔊 تکرار متن",
                },
                {
                    Text: "💾 وضعیت",
                },
            },
        },
        ResizeKeyboard:  true,
        OneTimeKeyboard: false,
    }
}

// 🎯 ارسال پیام با کیبورد
func sendMessageWithKeyboard(ctx context.Context, bot *ParsRubika.BotClient, chatID, text string) error {
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID:              chatID,
        Text:                text,
        ReplyKeyboardMarkup: createMainMenuKeyboard(),
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

// 📨 ارسال پیام با فرمت‌بندی
func sendFormattedMessage(ctx context.Context, bot *ParsRubika.BotClient, chatID, text string) error {
    // 🎨 استفاده از فرمت‌بندی Markdown
    formattedText := ParsRubika.Bold("پیام مهم") + "\n\n" +
        ParsRubika.Italic("این یک متن کج است") + "\n" +
        ParsRubika.Code("کد نمونه") + "\n" +
        ParsRubika.Link("کلیک کنید", "https://github.com/Abolfazl-Zarei/ParsRubika-bot-go")
    
    _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: chatID,
        Text:   formattedText,
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
    ChatID:               "CHAT_ID",              // 💬 شناسه چت مقصد
    Text:                 "متن پیام شما",         // 📝 متن پیام
    ReplyToMessageID:     "MSG_ID",              // ↩️ پاسخ به پیام خاص (اختیاری)
    InlineKeyboardMarkup:  inlineKeyboard,        // ⌨️ کیبورد اینلاین (اختیاری)
    ReplyKeyboardMarkup:  replyKeyboard,         // ⌨️ کیبورد پاسخ (اختیاری)
    DisableNotification:  false,                 // 🔕 بی‌صدا کردن (اختیاری)
})
```

#### ⌨️ ارسال پیام با کیبورد اینلاین
```go
// 🎮 ایجاد کیبورد اینلاین
inlineKeyboard := &ParsRubika.InlineKeyboardMarkup{
    InlineKeyboard: [][]*ParsRubika.InlineKeyboardButton{
        {
            {
                Text:         "🎯 دکمه ۱",
                CallbackData: "btn1",
            },
            {
                Text:         "🚀 دکمه ۲",
                CallbackData: "btn2",
            },
        },
        {
            {
                Text: "🌐 وب‌سایت",
                URL:  "https://github.com/Abolfazl-Zarei/ParsRubika-bot-go",
            },
        },
    },
}

// 📨 ارسال پیام با کیبورد اینلاین
messageID, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:              "CHAT_ID",
    Text:                "پیام با کیبورد تعاملی 🎮",
    InlineKeyboardMarkup: inlineKeyboard,
})
```

#### ⌨️ ارسال پیام با کیبورد پاسخ
```go
// 🎮 ایجاد کیبورد پاسخ
replyKeyboard := &ParsRubika.ReplyKeyboardMarkup{
    Keyboard: [][]*ParsRubika.KeyboardButton{
        {
            {
                Text: "📚 راهنما",
            },
            {
                Text: "👤 اطلاعات",
            },
        },
        {
            {
                Text:            "📞 شماره تماس",
                RequestContact:  true,
            },
            {
                Text:            "📍 موقعیت مکانی",
                RequestLocation: true,
            },
        },
    },
    ResizeKeyboard:  true,   // 📱 تنظیم سایز برای موبایل
    OneTimeKeyboard: false,  // 🔁 نمایش دائمی
}

// 📨 ارسال پیام با کیبورد پاسخ
messageID, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:             "CHAT_ID",
    Text:               "پیام با کیبورد پاسخ ⌨️",
    ReplyKeyboardMarkup: replyKeyboard,
})
```

#### ✏️ ویرایش پیام
```go
err := bot.EditMessageText(ctx, &ParsRubika.EditMessageTextRequest{
    ChatID:              "CHAT_ID",      // 💬 شناسه چت
    MessageID:           "MESSAGE_ID",   // 🔢 شناسه پیام
    Text:                "متن جدید",     // 📝 متن جدید
    InlineKeyboardMarkup: newKeyboard,    // ⌨️ کیبورد جدید (اختیاری)
})
```

#### ✏️ ویرایش کیبورد پیام
```go
err := bot.EditInlineKeypad(ctx, &ParsRubika.EditMessageKeypadRequest{
    ChatID:              "CHAT_ID",      // 💬 شناسه چت
    MessageID:           "MESSAGE_ID",   // 🔢 شناسه پیام
    InlineKeyboardMarkup: newKeyboard,    // ⌨️ کیبورد جدید
})
```

#### 🗑️ حذف پیام
```go
err := bot.DeleteMessage(ctx, &ParsRubika.DeleteMessageRequest{
    ChatID:    "CHAT_ID",      // 💬 شناسه چت
    MessageID: "MESSAGE_ID",   // 🔢 شناسه پیام
})
```

#### 📌 پین کردن پیام
```go
err := bot.PinChatMessage(ctx, &ParsRubika.PinChatMessageRequest{
    ChatID:              "CHAT_ID",      // 💬 شناسه چت
    MessageID:           "MESSAGE_ID",   // 🔢 شناسه پیام
    DisableNotification:  false,          // 🔕 بی‌صدا کردن (اختیاری)
})
```

#### 📌 آنپین کردن پیام
```go
err := bot.UnpinChatMessage(ctx, &ParsRubika.UnpinChatMessageRequest{
    ChatID:    "CHAT_ID",      // 💬 شناسه چت
    MessageID: "MESSAGE_ID",   // 🔢 شناسه پیام
})
```

#### 📌 آنپین کردن همه پیام‌ها
```go
err := bot.UnpinAllChatMessages(ctx, "CHAT_ID")  // 💬 شناسه چت
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

#### 🎵 ارسال موسیقی
```go
messageID, err := bot.SendMusic(ctx, "CHAT_ID", "path/to/music.mp3", "عنوان موسیقی")
```

#### 📄 ارسال فایل
```go
messageID, err := bot.SendDocument(ctx, "CHAT_ID", "path/to/file.pdf", "عنوان فایل")
```

#### 🎙️ ارسال پیام صوتی
```go
messageID, err := bot.SendVoice(ctx, "CHAT_ID", "path/to/voice.ogg", "عنوان پیام صوتی")
```

#### 🎨 ارسال گیف
```go
messageID, err := bot.SendGif(ctx, "CHAT_ID", "path/to/animation.gif", "عنوان گیف")
```

#### 🎨 ارسال استیکر
```go
messageID, err := bot.SendSticker(ctx, "CHAT_ID", "path/to/sticker.webp")
```

#### 📍 ارسال موقعیت
```go
messageID, err := bot.SendLocation(ctx, &ParsRubika.SendLocationRequest{
    ChatID:              "CHAT_ID",
    Latitude:            "35.6892",   // 📍 عرض جغرافیایی
    Longitude:           "51.3890",   // 📍 طول جغرافیایی
    ReplyToMessageID:    "MSG_ID",    // ↩️ پاسخ به پیام (اختیاری)
    InlineKeyboardMarkup: keyboard,    // ⌨️ کیبورد (اختیاری)
    DisableNotification: false,       // 🔕 بی‌صدا کردن (اختیاری)
})
```

#### 👥 ارسال مخاطب
```go
messageID, err := bot.SendContact(ctx, &ParsRubika.SendContactRequest{
    ChatID:              "CHAT_ID",
    FirstName:           "نام",
    LastName:            "نام خانوادگی", 
    PhoneNumber:         "09123456789",
    ReplyToMessageID:    "MSG_ID",    // ↩️ پاسخ به پیام (اختیاری)
    InlineKeyboardMarkup: keyboard,    // ⌨️ کیبورد (اختیاری)
    DisableNotification: false,       // 🔕 بی‌صدا کردن (اختیاری)
})
```

#### 📊 ارسال نظرسنجی
```go
messageID, err := bot.SendPoll(ctx, &ParsRubika.SendPollRequest{
    ChatID:              "CHAT_ID",
    Question:            "سوال نظرسنجی",
    Options:             []string{"گزینه اول", "گزینه دوم", "گزینه سوم"},
    ReplyToMessageID:    "MSG_ID",    // ↩️ پاسخ به پیام (اختیاری)
    InlineKeyboardMarkup: keyboard,    // ⌨️ کیبورد (اختیاری)
    DisableNotification: false,       // 🔕 بی‌صدا کردن (اختیاری)
})
```

#### 📝 ارسال فعالیت چت
```go
err := bot.SendChatActivity(ctx, "CHAT_ID", "typing")  // 📝 در حال تایپ
err = bot.SendChatActivity(ctx, "CHAT_ID", "upload_photo")  // 📷 در حال آپلود عکس
err = bot.SendChatActivity(ctx, "CHAT_ID", "record_video")  // 🎬 در حال ضبط ویدیو
```

### 🔄 فوروارد پیام‌ها

#### 📤 فوروارد یک پیام
```go
newMessageID, err := bot.ForwardMessage(ctx, &ParsRubika.ForwardMessageRequest{
    FromChatID:          "SOURCE_CHAT_ID",  // 📤 شناسه چت مبدا
    MessageID:           "MESSAGE_ID",      // 🔢 شناسه پیام
    ToChatID:            "TARGET_CHAT_ID",  // 📥 شناسه چت مقصد
    DisableNotification: false,             // 🔕 بی‌صدا کردن (اختیاری)
})
```

#### 📤 فوروارد چندین پیام
```go
messageIDs := []string{"MSG_ID_1", "MSG_ID_2", "MSG_ID_3"}
newMessageIDs, err := bot.ForwardMessages(ctx, "SOURCE_CHAT_ID", messageIDs, "TARGET_CHAT_ID")
```

### 💬 مدیریت چت و کاربران

#### 💬 اطلاعات چت
```go
chat, err := bot.GetChat(ctx, "CHAT_ID")
fmt.Printf("نام چت: %s\n", chat.Title)
fmt.Printf("نوع چت: %s\n", chat.ChatType)
```

#### 👤 اطلاعات کاربر
```go
user, err := bot.GetUserInfo(ctx, "USER_ID")
fmt.Printf("نام کاربر: %s %s\n", user.FirstName, user.LastName)
fmt.Printf("یوزرنیم: @%s\n", user.Username)
fmt.Printf("تأیید شده: %v\n", user.IsVerified)
```

#### 👥 لیست اعضا
```go
members, err := bot.GetMembers(ctx, "CHAT_ID")
for _, member := range members {
    fmt.Printf("عضو: %s\n", member.User.FirstName)
}
```

#### 👤 اطلاعات عضو خاص
```go
member, err := bot.GetChatMember(ctx, &ParsRubika.GetChatMemberRequest{
    ChatID: "CHAT_ID",
    UserID: "USER_ID",
})
fmt.Printf("وضعیت عضو: %s\n", member.Status)
```

#### 🔢 تعداد اعضا
```go
count, err := bot.GetChatMemberCount(ctx, &ParsRubika.GetChatMemberCountRequest{
    ChatID: "CHAT_ID",
})
fmt.Printf("تعداد اعضا: %d\n", count.Count)
```

#### 👥 مدیران چت
```go
admins, err := bot.GetChatAdministrators(ctx, &ParsRubika.GetChatAdministratorsRequest{
    ChatID: "CHAT_ID",
})
for _, admin := range admins.Administrators {
    fmt.Printf("مدیر: %s (%s)\n", admin.User.FirstName, admin.Status)
}
```

#### 🛡️ مدیریت دسترسی اعضا
```go
// 🚫 مسدود کردن عضو
err := bot.BanChatMember(ctx, &ParsRubika.BanChatMemberRequest{
    ChatID: "CHAT_ID",
    UserID: "USER_ID",
})

// ✅ رفع مسدودیت عضو
err = bot.UnbanChatMember(ctx, &ParsRubika.UnbanChatMemberRequest{
    ChatID: "CHAT_ID",
    UserID: "USER_ID",
})

// 📈 ارتقای عضو به مدیر
err = bot.PromoteChatMember(ctx, &ParsRubika.PromoteChatMemberRequest{
    ChatID:          "CHAT_ID",
    UserID:          "USER_ID",
    IsAdministrator: &[]bool{true}[0],  // 📌 تبدیل به مدیر
    CanChangeInfo:   &[]bool{true}[0],  // ✏️ امکان تغییر اطلاعات
    CanDeleteMessages: &[]bool{true}[0], // 🗑️ امکان حذف پیام‌ها
    CanInviteUsers:  &[]bool{true}[0],  // 👥 امکان دعوت کاربران
    CanPinMessages:  &[]bool{true}[0],  // 📌 امکان پین کردن پیام‌ها
})
```

#### ⚙️ تنظیم مجوزهای چت
```go
permissions := map[string]bool{
    "can_send_messages":       true,
    "can_send_media_messages": true,
    "can_send_polls":          false,
    "can_add_web_page_previews": false,
}

err := bot.SetChatPermissions(ctx, "CHAT_ID", permissions)
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
fmt.Printf("وضعیت: %s\n", status.State)
fmt.Printf("تعداد آراء: %d\n", status.TotalVote)
fmt.Printf("درصد آراء: %v\n", status.PercentVoteOptions)
```

### 🤖 اطلاعات بات

#### 🤖 اطلاعات بات
```go
bot, err := bot.GetMe(ctx)
fmt.Printf("نام بات: %s\n", bot.BotTitle)
fmt.Printf("یوزرنیم: @%s\n", bot.Username)
fmt.Printf("شناسه: %s\n", bot.BotID)
fmt.Printf("توضیحات: %s\n", bot.Description)
fmt.Printf("لینک اشتراک: %s\n", bot.ShareURL)
```

#### ⚙️ تنظیم دستورات بات
```go
commands := []ParsRubika.BotCommand{
    {
        Command:     "start",
        Description: "شروع کار با ربات",
    },
    {
        Command:     "help",
        Description: "راهنمای ربات",
    },
    {
        Command:     "info",
        Description: "اطلاعات کاربر",
    },
}

err := bot.SetCommands(ctx, &ParsRubika.SetCommandsRequest{
    BotCommands: commands,
})
```

#### 🔄 به‌روزرسانی endpointهای بات
```go
err := bot.UpdateBotEndpoints(ctx, "https://your-domain.com/webhook", ParsRubika.ReceiveUpdate)
```

### 📁 مدیریت فایل‌ها

#### 📤 درخواست آدرس آپلود
```go
uploadResp, err := bot.RequestSendFile(ctx, ParsRubika.ImageType)
fmt.Printf("آدرس آپلود: %s\n", uploadResp.UploadURL)
```

#### 📤 آپلود فایل
```go
uploadResp, err := bot.UploadFile("https://upload.url", "path/to/file.jpg")
fmt.Printf("شناسه فایل: %s\n", uploadResp.FileID)
```

#### 📄 اطلاعات فایل
```go
fileInfo, err := bot.GetFile(ctx, "FILE_ID")
fmt.Printf("نام فایل: %s\n", fileInfo.FileName)
fmt.Printf("سایز: %d\n", fileInfo.Size)
fmt.Printf("آدرس دانلود: %s\n", fileInfo.DownloadURL)
```

#### 📥 دانلود فایل
```go
err := bot.Download(ctx, "FILE_ID", "path/to/save/file.jpg")
```

#### 🖼️ دانلود عکس پروفایل
```go
err := bot.DownloadProfilePicture(ctx, "USER_ID", "path/to/save/avatar.jpg")
```

#### 📤 آپلود فایل مستقیم
```go
file, err := bot.UploadFileDirectly(ctx, "path/to/file.jpg", ParsRubika.ImageType)
fmt.Printf("شناسه فایل: %s\n", file.FileID)
fmt.Printf("نام فایل: %s\n", file.FileName)
fmt.Printf("سایز: %d\n", file.Size)
```

### 🎛 مدیریت وضعیت کاربر

#### 💾 ذخیره وضعیت
```go
// 🎯 ذخیره وضعیت‌های مختلف کاربر
bot.SetState(userID, "current_menu", "main")
bot.SetState(userID, "selected_item", "item_123")
bot.SetState(userID, "step", "2")
bot.SetState(userID, "form_data", map[string]string{
    "name": "John",
    "age":  "30",
})
```

#### 🔍 بازیابی وضعیت
```go
// 🔍 بازیابی وضعیت کاربر
menu, exists := bot.GetState(userID, "current_menu")
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

step, exists := bot.GetState(userID, "step")
if exists {
    // 🔄 ادامه فرآیند از مرحله ذخیره شده
}
```

#### 🗑️ مدیریت وضعیت
```go
// 🗑️ حذف یک کلید خاص
bot.DeleteState(userID, "selected_item")

// 🗑️ حذف تمام وضعیت‌های کاربر
bot.DeleteUserState(userID)
```

### 🛡️ سیستم ضد اسپم

#### ⚙️ تنظیمات ضد اسپم
```go
// 🛡️ دریافت مدیر ضد اسپم
antiSpam := bot.GetAntiSpam()

// ⏰ تنظیم زمان کول‌داون (پیش‌فرض: 3 ثانیه)
antiSpam.SetCooldown(5 * time.Second)
```

#### 🔍 بررسی ضد اسپم
```go
// 🛡️ بررسی اینکه آیا کاربر اسپم می‌کند یا خیر
if bot.CheckAntiSpam(userID) {
    // ✅ کاربر اسپم نمی‌کند، پردازش ادامه می‌یابد
    processUserMessage(ctx, update)
} else {
    // ⚠️ کاربر در حال اسپم کردن است
    bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
        ChatID: update.ChatID,
        Text:   "⚠️ لطفاً کمی صبر کنید و درخواست‌های خود را با فاصله ارسال کنید.",
    })
}
```

#### 🔄 بازنشانی ضد اسپم
```go
// 🔄 بازنشانی وضعیت ضد اسپم برای یک کاربر خاص
bot.ResetAntiSpam(userID)
```

### 🔄 قابلیت Hot-Reload

#### ⚙️ فعال‌سازی Hot-Reload
```go
// 🔄 ایجاد نمونه ربات با Hot-Reload فعال
bot := ParsRubika.NewClient(token, ParsRubika.WithHotReload(true))

// 🔄 یا فعال‌سازی پس از ساخت
bot.EnableHotReload()
```

#### 🔍 بررسی وضعیت Hot-Reload
```go
// 🔍 بررسی اینکه آیا Hot-Reload فعال است یا خیر
if bot.IsHotReloadEnabled() {
    fmt.Println("Hot-Reload فعال است")
} else {
    fmt.Println("Hot-Reload غیرفعال است")
}
```

#### 🔄 مدیریت Hot-Reload
```go
// 🔄 دریافت مدیر Hot-Reload
reloadManager := bot.GetReloadManager()

// 🎯 ثبت تابع برای اجرا پس از بارگذاری مجدد
reloadManager.OnReload(func() {
    fmt.Println("بارگذاری مجدد انجام شد!")
})

// 🔄 شروع نظارت بر تغییرات فایل‌ها
reloadManager.StartWatching()

// 🛑 توقف نظارت بر تغییرات فایل‌ها
reloadManager.StopWatching()

// 🔄 اجرای دستی بارگذاری مجدد
reloadManager.TriggerReload()
```

### 🌐 مدیریت پایداری شبکه

#### 🌐 دریافت مدیر پایداری شبکه
```go
// 🌐 دریافت مدیر پایداری شبکه
networkManager := bot.networkStabilityManager
```

#### ⏰ محاسبه تأخیر برای تلاش مجدد
```go
// ⏰ محاسبه تأخیر برای تلاش مجدد با الگوریتم نمایی و Jitter
delay := networkManager.CalculateBackoffDelay(retryCount)
fmt.Printf("تأخیر برای تلاش مجدد: %v\n", delay)
```

#### 🔍 بررسی خطای قابل تلاش مجدد
```go
// 🔍 بررسی اینکه آیا یک خطا قابل تلاش مجدد است یا خیر
if networkManager.IsRetryableError(err) {
    fmt.Println("خطا قابل تلاش مجدد است")
} else {
    fmt.Println("خطا قابل تلاش مجدد نیست")
}
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
func createMainMenuKeyboard() *ParsRubika.ReplyKeyboardMarkup {
    return &ParsRubika.ReplyKeyboardMarkup{
        Keyboard: [][]*ParsRubika.KeyboardButton{
            {
                {
                    Text: "📚 راهنما",
                },
                {
                    Text: "👤 اطلاعات",
                },
            },
            {
                {
                    Text: "🔊 تکرار متن",
                },
                {
                    Text: "💾 وضعیت",
                },
            },
        },
        ResizeKeyboard:  true,   // 📱 تنظیم سایز برای موبایل
        OneTimeKeyboard: false,  // 🔁 نمایش دائمی
    }
}
```

### 🎯 کیبوردهای پیشرفته

#### 🔘 دکمه انتخاب (Selection)
```go
selectionBtn := ParsRubika.KeyboardButton{
    Text: "📁 انتخاب آیتم",
    Type: ParsRubika.ButtonTypeSelection,
    // در پیاده‌سازی کامل، باید فیلدهای مربوط به ButtonSelection تنظیم شوند
}

// 📨 ارسال پیام با دکمه انتخاب
_, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:             chatID,
    Text:               "لطفا یک آیتم انتخاب کنید:",
    ReplyKeyboardMarkup: &ParsRubika.ReplyKeyboardMarkup{
        Keyboard: [][]*ParsRubika.KeyboardButton{
            {&selectionBtn},
        },
    },
})
```

#### 📅 دکمه تقویم
```go
calendarBtn := ParsRubika.KeyboardButton{
    Text: "📅 انتخاب تاریخ",
    Type: ParsRubika.ButtonTypeCalendar,
    // در پیاده‌سازی کامل، باید فیلدهای مربوط به ButtonCalendar تنظیم شوند
}

// 📨 ارسال پیام با دکمه تقویم
_, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:             chatID,
    Text:               "تاریخ مورد نظر را انتخاب کنید:",
    ReplyKeyboardMarkup: &ParsRubika.ReplyKeyboardMarkup{
        Keyboard: [][]*ParsRubika.KeyboardButton{
            {&calendarBtn},
        },
    },
})
```

#### 🔢 دکمه انتخاب عدد
```go
numberPickerBtn := ParsRubika.KeyboardButton{
    Text: "🔢 انتخاب عدد",
    Type: ParsRubika.ButtonTypeNumberPicker,
    // در پیاده‌سازی کامل، باید فیلدهای مربوط به ButtonNumberPicker تنظیم شوند
}

// 📨 ارسال پیام با دکمه انتخاب عدد
_, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:             chatID,
    Text:               "عدد مورد نظر را انتخاب کنید:",
    ReplyKeyboardMarkup: &ParsRubika.ReplyKeyboardMarkup{
        Keyboard: [][]*ParsRubika.KeyboardButton{
            {&numberPickerBtn},
        },
    },
})
```

### 🎨 استفاده از کیبوردها

```go
// 📨 ارسال پیام با کیبورد اصلی
_, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:             chatID,
    Text:               "🎮 **منوی اصلی ربات**\n\nلطفا یک گزینه انتخاب کنید:",
    ReplyKeyboardMarkup: createMainMenuKeyboard(),
})

// ✏️ ویرایش کیبورد پیام موجود
err := bot.EditChatKeypad(ctx, &ParsRubika.EditChatKeypadRequest{
    ChatID:             chatID,
    ReplyKeyboardMarkup: createNewKeyboard(),
    ChatKeypadType:     ParsRubika.NewKeypad,
})

// 🗑️ حذف کیبورد
err = bot.EditChatKeypad(ctx, &ParsRubika.EditChatKeypadRequest{
    ChatID:         chatID,
    RemoveKeyboard: &ParsRubika.ReplyKeyboardRemove{
        RemoveKeyboard: true,
    },
    ChatKeypadType: ParsRubika.RemoveKeypad,
})
```

### 🎨 کیبوردهای اینلاین (Inline)

```go
// 🎮 ایجاد کیبورد اینلاین
inlineKeyboard := &ParsRubika.InlineKeyboardMarkup{
    InlineKeyboard: [][]*ParsRubika.InlineKeyboardButton{
        {
            {
                Text:         "🎯 دکمه ۱",
                CallbackData: "btn1",
            },
            {
                Text:         "🚀 دکمه ۲",
                CallbackData: "btn2",
            },
        },
        {
            {
                Text: "🌐 وب‌سایت",
                URL:  "https://github.com/Abolfazl-Zarei/ParsRubika-bot-go",
            },
            {
                Text:              "🔍 جستجو",
                SwitchInlineQuery: "search_query",
            },
        },
    },
}

// 📨 ارسال پیام با کیبورد اینلاین
_, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:              chatID,
    Text:                "پیام با کیبورد تعاملی 🎮",
    InlineKeyboardMarkup: inlineKeyboard,
})

// ✏️ ویرایش کیبورد اینلاین پیام موجود
err = bot.EditInlineKeypad(ctx, &ParsRubika.EditMessageKeypadRequest{
    ChatID:              chatID,
    MessageID:           messageID,
    InlineKeyboardMarkup: newInlineKeyboard,
})
```

### 🎨 توابع کمکی برای ساخت کیبوردها

```go
// 🎮 ایجاد کیبورد اینلاین
func NewInlineKeyboard(rows ...[]*ParsRubika.InlineKeyboardButton) *ParsRubika.InlineKeyboardMarkup {
    keyboard := &ParsRubika.InlineKeyboardMarkup{
        InlineKeyboard: make([][]*ParsRubika.InlineKeyboardButton, len(rows)),
    }

    for i, row := range rows {
        keyboard.InlineKeyboard[i] = make([]*ParsRubika.InlineKeyboardButton, len(row))
        copy(keyboard.InlineKeyboard[i], row)
    }

    return keyboard
}

// 🎮 ایجاد دکمه کیبورد اینلاین
func NewInlineKeyboardButton(text string) *ParsRubika.InlineKeyboardButton {
    return &ParsRubika.InlineKeyboardButton{Text: text}
}

// 🔗 تنظیم URL برای دکمه
func (b *ParsRubika.InlineKeyboardButton) WithURL(url string) *ParsRubika.InlineKeyboardButton {
    b.URL = url
    return b
}

// 🔗 تنظیم داده callback برای دکمه
func (b *ParsRubika.InlineKeyboardButton) WithCallbackData(data string) *ParsRubika.InlineKeyboardButton {
    b.CallbackData = data
    return b
}

// 🔍 تنظیم سوئیچ اینلاین کوئری برای دکمه
func (b *ParsRubika.InlineKeyboardButton) WithSwitchInlineQuery(query string) *ParsRubika.InlineKeyboardButton {
    b.SwitchInlineQuery = query
    return b
}

// 🎮 ایجاد کیبورد پاسخ
func NewReplyKeyboard(rows ...[]*ParsRubika.KeyboardButton) *ParsRubika.ReplyKeyboardMarkup {
    keyboard := &ParsRubika.ReplyKeyboardMarkup{
        Keyboard:       make([][]*ParsRubika.KeyboardButton, len(rows)),
        ResizeKeyboard: true,
    }

    for i, row := range rows {
        keyboard.Keyboard[i] = make([]*ParsRubika.KeyboardButton, len(row))
        copy(keyboard.Keyboard[i], row)
    }

    return keyboard
}

// 🎮 ایجاد دکمه کیبورد پاسخ
func NewReplyKeyboardButton(text string) *ParsRubika.KeyboardButton {
    return &ParsRubika.KeyboardButton{Text: text}
}

// 📞 تنظیم درخواست شماره تلفن برای دکمه
func (b *ParsRubika.KeyboardButton) WithRequestContact() *ParsRubika.KeyboardButton {
    b.RequestContact = true
    return b
}

// 📍 تنظیم درخواست موقعیت مکانی برای دکمه
func (b *ParsRubika.KeyboardButton) WithRequestLocation() *ParsRubika.KeyboardButton {
    b.RequestLocation = true
    return b
}

// 📱 تنظیم تغییر اندازه کیبورد
func (kb *ParsRubika.ReplyKeyboardMarkup) WithResizeKeyboard(resize bool) *ParsRubika.ReplyKeyboardMarkup {
    kb.ResizeKeyboard = resize
    return kb
}

// 🔁 تنظیم کیبورد یکبار مصرف
func (kb *ParsRubika.ReplyKeyboardMarkup) WithOneTimeKeyboard(oneTime bool) *ParsRubika.ReplyKeyboardMarkup {
    kb.OneTimeKeyboard = oneTime
    return kb
}

// 🗑️ ایجاد دستور برای حذف کیبورد
func NewRemoveKeyboard() *ParsRubika.ReplyKeyboardRemove {
    return &ParsRubika.ReplyKeyboardRemove{
        RemoveKeyboard: true,
    }
}
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

### 📤 آپلود آواتار

```go
// 🖼️ آپلود آواتار
avatar, err := bot.UploadAvatar(ctx, "path/to/avatar.jpg")
if err != nil {
    return fmt.Errorf("❌ خطا در آپلود آواتار: %w", err)
}

// 📨 ارسال پیام با اطلاعات آواتار
infoText := fmt.Sprintf(`🖼️ **آواتار آپلود شد!**

🏷️ نام: %s
📊 سایز: %s
🆔 شناسه: %s`, avatar.FileName, avatar.Size, avatar.FileID)

_, err = bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID: chatID,
    Text:   infoText,
})
```

### 🗑️ حذف آواتار

```go
// 🗑️ حذف آواتار
err := bot.DeleteAvatar(ctx)
if err != nil {
    return fmt.Errorf("❌ خطا در حذف آواتار: %w", err)
}

_, err = bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID: chatID,
    Text:   "✅ آواتار با موفقیت حذف شد!",
})
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
        Category:    "الکترونی
        

```go
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

RUN apk --no-cache add ca-certificates tzdata
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
      - ./data:/app/data
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s

  # 🗄️ دیتابیس اختیاری (PostgreSQL)
  postgres:
    image: postgres:15-alpine
    container_name: rubika-bot-db
    environment:
      - POSTGRES_DB=rubika_bot
      - POSTGRES_USER=bot_user
      - POSTGRES_PASSWORD=${DB_PASSWORD}
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql
    restart: unless-stopped
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U bot_user -d rubika_bot"]
      interval: 10s
      timeout: 5s
      retries: 5

  # 📊 Redis برای کش و State Management
  redis:
    image: redis:7-alpine
    container_name: rubika-bot-redis
    volumes:
      - redis_data:/data
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 5s
      retries: 3

volumes:
  postgres_data:
  redis_data:
```

### 🚀 استقرار روی سرور

#### 1. تنظیمات systemd
```ini
# /etc/systemd/system/rubika-bot.service
[Unit]
Description=Rubika Bot Service
After=network.target postgresql.service redis.service
Wants=network.target

[Service]
Type=simple
User=botuser
Group=botuser
WorkingDirectory=/opt/rubika-bot
Environment="RUBIKA_BOT_TOKEN=your_bot_token_here"
Environment="DATABASE_URL=postgresql://bot_user:password@localhost:5432/rubika_bot"
Environment="REDIS_URL=redis://localhost:6379"
ExecStart=/opt/rubika-bot/rubika-bot
ExecReload=/bin/kill -HUP $MAINPID
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
ReadWritePaths=/opt/rubika-bot/data /opt/rubika-bot/logs

[Install]
WantedBy=multi-user.target
```

#### 2. اسکریپت استقرار
```bash
#!/bin/bash
# deploy.sh - اسکریپت استقرار خودکار

set -e

echo "🚀 شروع استقرار ربات..."

# 🔄 توقف سرویس فعلی
echo "🛑 توقف سرویس ربات..."
sudo systemctl stop rubika-bot || true

# 📥 دریافت آخرین تغییرات
echo "📥 دریافت آخرین نسخه از گیت‌هاب..."
git pull origin main

# 📦 ساخت پروژه
echo "🏗 ساخت باینری پروژه..."
go build -ldflags="-s -w" -o rubika-bot .

# 🔒 تنظیم مجوزها
echo "🔒 تنظیم مجوزهای فایل‌ها..."
chmod +x rubika-bot
sudo chown botuser:botuser rubika-bot

# 🗄️ اجرای Migrations دیتابیس (اگر نیاز باشد)
# echo "🗄️ اجرای Migrations دیتابیس..."
# ./rubika-bot migrate up

# 🚀 راه‌اندازی سرویس
echo "🚀 راه‌اندازی مجدد سرویس..."
sudo systemctl daemon-reload
sudo systemctl start rubika-bot
sudo systemctl enable rubika-bot

# ✅ بررسی وضعیت
echo "✅ بررسی وضعیت سرویس..."
sleep 5
if sudo systemctl is-active --quiet rubika-bot; then
    echo "🎉 استقرار با موفقیت انجام شد!"
    sudo systemctl status rubika-bot --no-pager -l
else
    echo "❌ خطا در راه‌اندازی سرویس!"
    sudo journalctl -u rubika-bot -n 20 --no-pager
    exit 1
fi
```

### ☁️ استقرار ابری (AWS/GCP/Azure)

#### Kubernetes Deployment
```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: rubika-bot
  labels:
    app: rubika-bot
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
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
          name: http
        env:
        - name: RUBIKA_BOT_TOKEN
          valueFrom:
            secretKeyRef:
              name: bot-secrets
              key: token
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: bot-secrets
              key: database-url
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
            port: http
          initialDelaySeconds: 30
          periodSeconds: 10
          timeoutSeconds: 5
          failureThreshold: 3
        readinessProbe:
          httpGet:
            path: /ready
            port: http
          initialDelaySeconds: 5
          periodSeconds: 5
          timeoutSeconds: 3
          failureThreshold: 3
---
apiVersion: v1
kind: Service
metadata:
  name: rubika-bot-service
spec:
  selector:
    app: rubika-bot
  ports:
    - protocol: TCP
      port: 80
      targetPort: http
  type: ClusterIP
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: rubika-bot-ingress
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
spec:
  tls:
  - hosts:
    - bot.yourdomain.com
    secretName: rubika-bot-tls
  rules:
  - host: bot.yourdomain.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: rubika-bot-service
            port:
              number: 80
```

#### GitHub Actions CI/CD
```yaml
# .github/workflows/deploy.yml
name: Build and Deploy

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
    
    - name: Cache Go modules
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
        restore-keys: |
          ${{ runner.os }}-go-
    
    - name: Download dependencies
      run: go mod download
    
    - name: Run tests
      run: go test -v -race -coverprofile=coverage.out ./...
    
    - name: Upload coverage to Codecov
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out

  build-and-push:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Log in to Docker Hub
      uses: docker/login-action@v2
      with:
        username: ${{ secrets.DOCKER_USERNAME }}
        password: ${{ secrets.DOCKER_PASSWORD }}
    
    - name: Extract metadata
      id: meta
      uses: docker/metadata-action@v4
      with:
        images: your-registry/rubika-bot
        tags: |
          type=ref,event=branch
          type=ref,event=pr
          type=sha,prefix={{branch}}-
    
    - name: Build and push Docker image
      uses: docker/build-push-action@v4
      with:
        context: .
        push: true
        tags: ${{ steps.meta.outputs.tags }}
        labels: ${{ steps.meta.outputs.labels }}
        cache-from: type=gha
        cache-to: type=gha,mode=max

  deploy:
    needs: build-and-push
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up kubectl
      uses: azure/setup-kubectl@v3
      with:
        version: 'v1.24.0'
    
    - name: Configure kubectl
      run: |
        echo "${{ secrets.KUBE_CONFIG }}" | base64 -d > kubeconfig
        export KUBECONFIG=kubeconfig
    
    - name: Deploy to Kubernetes
      run: |
        export KUBECONFIG=kubeconfig
        kubectl set image deployment/rubika-bot rubika-bot=your-registry/rubika-bot:main -n rubika-bot
        kubectl rollout status deployment/rubika-bot -n rubika-bot
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
        log.Println("💡 راه‌حل: تایم‌اوت را افزایش دهید یا شبکه را بررسی کنید")
    } else if strings.Contains(err.Error(), "no such host") {
        log.Println("🌐 خطای DNS: نام دامنه پیدا نشد")
        log.Println("💡 راه‌حل: تنظیمات DNS و فایل hosts را بررسی کنید")
    }
}
```

#### 2. خطای توکن
```go
// 🔑 مشکل: توکن نامعتبر
func handleTokenError(err error) {
    if strings.Contains(err.Error(), "unauthorized") || strings.Contains(err.Error(), "401") {
        log.Println("❌ خطای احراز هویت: توکن نامعتبر است")
        log.Println("💡 راه‌حل:")
        log.Println("  1. توکن را از @rubika_bot دریافت کنید")
        log.Println("  2. مطمئن شوید ربات فعال است")
        log.Println("  3. توکن را درست کپی کنید (فضاهای خالی را حذف کنید)")
        log.Println("  4. از صحیح بودن متغیر محیطی RUBIKA_BOT_TOKEN اطمینان حاصل کنید")
    }
}
```

#### 3. محدودیت نرخ ارسال
```go
// 🚦 مشکل: ارسال درخواست‌های زیاد
func handleRateLimit(bot *ParsRubika.BotClient) {
    // ⏳ افزایش تأخیر بین درخواست‌ها
    log.Println("🚦 خطای محدودیت نرخ ارسال دریافت شد")
    log.Println("💡 راه‌حل: افزایش تأخیر بین درخواست‌ها")
    
    // 🔄 ساخت مجدد کلاینت با تنظیمات جدید
    bot = ParsRubika.NewClient(token,
        ParsRubika.WithRateLimitDelay(2 * time.Second),
        ParsRubika.WithMaxRetries(5),
    )
}
```

### 📊 مانیتورینگ و سلامت

#### سلامت‌سنجی (Health Check)
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
                    // منطق بازنشانی توکن یا ارسال هشدار به ادمین
                }
            } else {
                log.Println("✅ ربات در وضعیت سالم")
            }
            
            cancel()
        }
    }()
}

// 🌐 افزودن endpoint سلامت برای وب‌هوک
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status": "ok", "timestamp": "` + time.Now().Format(time.RFC3339) + `"}`))
}
```

#### میدلور لاگینگ پیشرفته
```go
type LoggingMiddleware struct {
    bot *ParsRubika.BotClient
}

func (lm *LoggingMiddleware) HandleUpdate(ctx context.Context, update *ParsRubika.Update, next ParsRubika.HandlerFunc) error {
    start := time.Now()
    
    log.Printf("📥 آپدیت دریافت شد - نوع: %s, چت: %s, کاربر: %s", 
        update.Type, update.ChatID, getSenderID(update))
    
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

func getSenderID(update *ParsRubika.Update) string {
    if update.NewMessage != nil {
        return update.NewMessage.SenderID
    }
    if update.UpdatedMessage != nil {
        return update.UpdatedMessage.SenderID
    }
    return "Unknown"
}

func (lm *LoggingMiddleware) sendErrorReport(ctx context.Context, update *ParsRubika.Update, err error) {
    // 📨 ارسال گزارش خطا به ادمین
    errorMsg := fmt.Sprintf(`🚨 **گزارش خطا در ربات**

💬 چت: %s
👤 کاربر: %s
❌ خطا: %v
⏰ زمان: %s
📄 آپدیت: %+v`,
        update.ChatID,
        getSenderID(update),
        err,
        time.Now().Format("2006/01/02 15:04:05"),
        update,
    )
    
    // ارسال به ادمین‌ها
    adminIDs := []string{"ADMIN_ID_1", "ADMIN_ID_2"}
    for _, adminID := range adminIDs {
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

set -e

echo "🔧 محیط توسعه ParsRubika"

case "$1" in
    "run")
        echo "🚀 اجرای ربات در حالت توسعه..."
        RUBIKA_BOT_TOKEN=${RUBIKA_BOT_TOKEN:-"dev_token"} go run main.go handlers.go utils.go
        ;;
    "test")
        echo "🧪 اجرای تست‌ها..."
        go test -v -race -coverprofile=coverage.out ./...
        go tool cover -html=coverage.out -o coverage.html
        echo "📊 گزارش پوشش در coverage.html ذخیره شد"
        ;;
    "build")
        echo "🏗 ساخت باینری..."
        go build -ldflags="-s -w" -o rubika-bot .
        ;;
    "lint")
        echo "📝 بررسی کد با golangci-lint..."
        golangci-lint run
        ;;
    "fmt")
        echo "🎨 فرمت‌بندی کد..."
        go fmt ./...
        goimports -w .
        ;;
    "clean")
        echo "🧹 پاک‌سازی..."
        go clean
        rm -f rubika-bot coverage.out coverage.html
        ;;
    "docker-build")
        echo "🐳 ساخت ایمیج داکر..."
        docker build -t rubika-bot:dev .
        ;;
    "docker-run")
        echo "🐳 اجرا با داکر..."
        docker run --rm -it -p 8080:8080 -e RUBIKA_BOT_TOKEN=$RUBIKA_BOT_TOKEN rubika-bot:dev
        ;;
    *)
        echo "💡 استفاده: ./dev.sh [run|test|build|lint|fmt|clean|docker-build|docker-run]"
        echo ""
        echo "📋 دستورات:"
        echo "  run           - اجرای ربات در حالت توسعه"
        echo "  test          - اجرای تست‌ها و گزارش پوشش"
        echo "  build         - ساخت باینری"
        echo "  lint          - بررسی کد با linter"
        echo "  fmt           - فرمت‌بندی کد"
        echo "  clean         - پاک‌سازی فایل‌های اضافی"
        echo "  docker-build  - ساخت ایمیج داکر"
        echo "  docker-run    - اجرا با داکر"
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
    secret: "${WEBHOOK_SECRET}"
  polling:
    enabled: true
    interval: "2s"
    limit: 100
    timeout: "30s"

database:
  url: "${DATABASE_URL}"
  max_connections: 20
  max_idle_connections: 5
  connection_max_lifetime: "1h"

redis:
  url: "${REDIS_URL}"
  pool_size: 10

logging:
  level: "info"
  format: "json" # json or text
  file: "logs/bot.log"
  max_size: 100 # MB
  max_backups: 3
  max_age: 28 # days

features:
  state_management: true
  file_upload: true
  admin_panel: true
  anti_spam: true
  hot_reload: false

monitoring:
  prometheus:
    enabled: true
    port: 9090
  health_check:
    enabled: true
    path: "/health"
```

---

## 📞 پشتیبانی

### 🔗 ارتباط با توسعه‌دهنده

- **👤 ایدی روبیکا:** `NinjaCode`
- **📢 چنل روبیکا:** `Ninja_code`
- **📧 ایمیل:** `ninjacode.ir@gmail.com`
- **🐙 گیت‌هاب:** [Abolfazl-Zarei](https://github.com/Abolfazl-Zarei)
- **🌐 وب‌سایت:** [ninjacode.ir](https://ninjacode.ir) (در صورت وجود)

### 📝 گزارش مشکل

برای گزارش باگ یا درخواست ویژگی جدید، لطفاً مراحل زیر را دنبال کنید:

1.  **جستجو:** ابتدا در [ایسوها](https://github.com/Abolfazl-Zarei/ParsRubika-bot-go/issues) جستجو کنید تا مطمئن شوید مشکل تکراری نیست.
2.  **ایجاد ایسو جدید:**
    *   از یک عنوان واضح و مشخص استفاده کنید.
    *   توضیح دهید که مشکل چیست و چگونه می‌توان آن را بازتولید کرد.
    *   قطعه کد مربوطه را ارسال کنید.
    *   خروجی خطا را کامل کپی کنید.
    *   محیط خود را مشخص کنید (نسخه Go, سیستم‌عامل, نسخه کتابخانه).

```bash
# 1. بررسی نسخه‌ها
go version
go list -m all | grep ParsRubika

# 2. بررسی لاگ‌ها
tail -f logs/bot.log

# 3. جمع‌آوری اطلاعات سیستم
go env
uname -a
```

### 🤝 مشارکت در توسعه

مشارکت شما باعث خوشحالی ماست! برای شروع:

1.  **Fork کردن پروژه:** روی دکمه Fork در بالای صفحه گیت‌هاب کلیک کنید.
2.  **کلون کردن مخزن:**
    ```bash
    git clone https://github.com/YOUR_USERNAME/ParsRubika-bot-go.git
    cd ParsRubika-bot-go
    ```
3.  **ایجاد شاخه جدید:**
    ```bash
    git checkout -b feature/amazing-feature
    ```
4.  **انجام تغییرات:**
    *   کد خود را بنویسید.
    *   تست‌های مربوطه را اضافه کنید.
    *   مطمئن شوید که همه تست‌ها پاس می‌شوند: `go test ./...`
    *   کد را با `go fmt` و `goimports` فرمت کنید.
5.  **کامیت کردن تغییرات:**
    ```bash
    git add .
    git commit -m "اضافه کردن قابلیت جدید: توضیحی کوتاه در مورد تغییرات"
    ```
6.  **Push کردن به شاخه خود:**
    ```bash
    git push origin feature/amazing-feature
    ```
7.  **ایجاد Pull Request:**
    *   به صفحه مخزن خود در گیت‌هاب بروید.
    *   روی دکمه "New pull request" کلیک کنید.
    *   شاخه خود را با شاخه `main` اصلی مقایسه کنید.
    *   یک توضیح واضح برای PR خود بنویسید.

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
- 🚀 **مثال‌های پیشرفته** واقعی (فروشگاه، نظرسنجی، مدیریت کانال)
- ☁ **استقرار و دیپلوی** حرفه‌ای (داکر، کوبرنتیز)
- 🔧 **عیب‌یابی** و مانیتورینگ
- 🛡️ **قابلیت‌های پیشرفته** مانند Hot-Reload و Anti-Spam

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
    // 🤖 ایجاد ربات با تمام قابلیت‌ها
    bot := ParsRubika.NewClient(os.Getenv("RUBIKA_BOT_TOKEN"),
        ParsRubika.WithRateLimitDelay(1*time.Second),
        ParsRubika.WithMaxRetries(3),
        ParsRubika.WithIgnoreTimeout(true),
        ParsRubika.WithHotReload(true),
    )
    
    // 🎯 هندلر ساده و قدرتمند
    bot.OnMessageUpdates(func(ctx context.Context, update *ParsRubika.Update) error {
        if update.NewMessage != nil {
            // 🛡️ بررسی ضد اسپم
            if !bot.CheckAntiSpam(update.NewMessage.SenderID) {
                return nil
            }
            
            // 💾 ذخیره وضعیت کاربر
            bot.SetState(update.NewMessage.SenderID, "last_message", update.NewMessage.Text)
            
            // 📨 ارسال پاسخ
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
    if err := bot.Run(ctx); err != nil {
        log.Fatal("💥 خطا در اجرای ربات:", err)
    }
}
```

### 📚 منابع بیشتر

- 📖 [مستندات رسمی روبیکا](https://rubika.ir/docs)
- 💻 [مخزن گیت‌هاب](https://github.com/Abolfazl-Zarei/ParsRubika-bot-go)
- 🐦 [کانال اطلاع‌رسانی](https://rubika.ir/Ninja_code)
- 📝 [ایسوها و باگ‌ها](https://github.com/Abolfazl-Zarei/ParsRubika-bot-go/issues)
- 📖 [مستندات Go](https://golang.org/doc/)

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
