#################################################################################
#                                                                               #
#     🤖 کتابخانه ParsRubika Bot Go | مستندات رسمی و کامل 🇮🇷                     #
#                                                                               #
#################################################################################
#                                                                               #
#  🚀 ساخت ربات‌های حرفه‌ای و هوشمند برای پلتفرم روبیکا با زبان گولنگ            #
#                                                                               #
#################################################################################
#                                                                               #
#    ✍️ نویسنده: ابوالفضل زارعی                                                    #
#    🔗 آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go       #
#                                                                               #
#################################################################################

<div align="center">

<img src="https://img.icons8.com/color/120/000000/rubika.png" alt="Rubika Icon"/>
<img src="https://img.icons8.com/color/120/000000/golang.png" alt="Golang Gopher"/>
<img src="https://img.icons8.com/color/120/000000/robot.png" alt="Bot Icon"/>

### **✨ قدرتمندترین کتابخانه Golang برای API روبیکا ✨**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/Abolfazl-Zarei/ParsRubika-bot-go?style=for-the-badge&color=gold)](https://github.com/Abolfazl-Zarei/ParsRubika-bot-go/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/Abolfazl-Zarei/ParsRubika-bot-go?style=for-the-badge&color=blue)](https://github.com/Abolfazl-Zarei/ParsRubika-bot-go/network/members)
[![GitHub issues](https://img.shields.io/github/issues/Abolfazl-Zarei/ParsRubika-bot-go?style=for-the-badge&color=orange)](https://github.com/Abolfazl-Zarei/ParsRubika-bot-go/issues)

**🔥 اولین و کامل‌ترین کتابخانه Golang برای API روبیکا 🔥**

</div>

================================================================================
📚 فهرست مطالب
================================================================================

1.  🌟 معرفی و چرا ParsRubika؟
2.  🚀 ویژگی‌های کلیدی و انحصاری
3.  💻 نصب و راه‌اندازی سریع
4.  🎯 شروع سریع: اولین بات شما در ۵ دقیقه!
5.  🏗️ مفاهیم اصلی: Client, Polling, Webhook
6.  📖 مستندات کامل API
    -   📤 ارسال پیام (متن، کیبورد، مدیا)
    -   📁 مدیریت فایل (آپلود، ارسال، دانلود)
    -   📍 ارسال موقعیت، مخاطب و نظرسنجی
    -   ✏️ ویرایش، حذف و فوروارد پیام
    -   👥 مدیریت گروه و کانال
    -   📱 مدیریت صفحات و استوری‌ها (توابع غیررسمی)
7.  🎯 مثال‌های پیشرفته
    -   🧠 مدیریت State (حالت کاربر) در فرم‌های چند مرحله‌ای
    -   🎨 ساخت کیبوردهای پیشرفته (تقویم، انتخابگر و...)
8.  🤝 مشارکت در پروژه و لایسنس
9.  📞 ارتباط با ما

================================================================================
1. 🌟 معرفی و چرا ParsRubika؟
================================================================================

**ParsRubika Bot Go** یک کتابخانه جامع، مدرن و قدرتمند برای زبان Go (Golang) است که به شما امکان می‌دهد به سادگی و با کمترین کدنویسی، ربات‌های پیشرفته و هوشمند برای پلتفرم روبیکا ایجاد کنید. این کتابخانه با هدف پوشش کامل API روبیکا و ارائه یک رابط کاربری ساده و intuitional طراحی شده است.

✨ **چرا این کتابخانه را انتخاب کنید؟**
------------------------------------
-   **سادگی فوق‌العاده:** با چند خط کد، بات خود را راه‌اندازی کنید.
-   **کامل‌ترین پوشش API:** از تمام متدهای رسمی و بسیاری از متدهای غیررسمی (صفحات، استوری) پشتیبانی می‌کند.
-   **عملکرد بالا:** بهینه‌سازی شده برای سرعت و مصرف منابع کم.
-   **کد تمیز و خوانا:** ساختار کد به گونه‌ای است که توسعه و نگهداری آن آسان باشد.
-   **مستندات کامل:** مثال‌های کاربردی برای هر بخش، یادگیری را آسان می‌کند.
-   **پشتیبانی فعال:** ما برای حل مشکلات شما آماده‌ایم!

================================================================================
2. 🚀 ویژگی‌های کلیدی و انحصاری
================================================================================

### 🎯 **ویژگی‌های اصلی**

| دسته‌بندی | ویژگی‌ها                                                                                             |
|-----------|------------------------------------------------------------------------------------------------------|
| 📡 **ارتباط** | 🔄 Polling هوشمند • 🌐 Webhook با Graceful Shutdown • ⚡ دریافت آپدیت‌ها در لحظه               |
| 💬 **پیام‌رسانی** | 📝 متن • 📁 فایل‌های گوناگون (عکس، ویدیو، موسیقی) • 📍 موقعیت مکانی • 👤 کارت مخاطب • 📊 نظرسنجی |
| 🎨 **رابط کاربری** | ⌨️ کیبورد اصلی و اینلاین • 🔘 دکمه‌های پیشرفته (انتخاب، تقویم، موقعیت) • 🎨 شخصی‌سازی کامل        |
| 👥 **مدیریت** | 🏢 گروه‌ها • 📢 کانال‌ها • 👮 مدیریت اعضا (Ban, Promote) • 📌 پیام‌های سنجاق شده              |
| 📱 **صفحات** | 📸 پست‌ها • 🎬 استوری‌ها • 🌟 هایلایت‌ها • 👥 دنبال کردن و لایک کردن (توابع غیررسمی)           |

### ⚡ **مزایای کلیدی**

-   🚀 **سرعت و کارایی بالا:** بهینه‌سازی شده برای عملکرد عالی و مصرف منابع کم.
-   🛡️ **مدیریت هوشمند خطا:** مدیریت خودکار خطاهای شبکه و API با قابلیت تلاش مجدد.
-   🧠 **State Manager داخلی:** ابزاری قدرتمند برای ذخیره اطلاعات کاربران بین پیام‌ها.
-   📊 **لاگ‌گیری پیشرفته:** لاگ‌های دقیق برای دیباگینگ آسان‌تر.
-   🔄 **آپدیت مداوم:** پشتیبانی از جدیدترین ویژگی‌های API روبیکا.

================================================================================
3. 💻 نصب و راه‌اندازی سریع
================================================================================

نصب این کتابخانه به سادگی یک دستور است. ترمینال خود را باز کرده و دستور زیر را اجرا کنید:

```bash
go get github.com/Abolfazl-Zarei/ParsRubika-bot-go
```

**پیش‌نیاز:**
قبل از شروع، شما به یک توکن بات از [@BotFather](https://rubika.ir/botfather) در روبیکا نیاز دارید.

================================================================================
4. 🎯 شروع سریع: اولین بات شما در ۵ دقیقه!
================================================================================

کد زیر یک بات ساده "اِکو" را نشان می‌دهد که هر پیامی را دریافت کرده و همان را به عنوان پاسخ برمی‌گرداند. این بهترین نقطه برای شروع است!

```go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/Abolfazl-Zarei/ParsRubika-bot-go"
)

func main() {
    // 🔑 توکن بات خود را اینجا قرار دهید
    botToken := "YOUR_BOT_TOKEN"
    
    // 🤖 ایجاد یک نمونه جدید از کلاینت بات
    bot := ParsRubika.NewClient(botToken)
    
    // 🛡️ ایجاد یک کانتکست برای مدیریت صحیح منابع و خروج Graceful
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    // 🛑 مدیریت سیگنال‌های سیستم (مانند Ctrl+C) برای خروج صحیح برنامه
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigChan
        log.Println("🛑 سیگنال خروج دریافت شد. بات در حال خاموش شدن...")
        cancel()
    }()
    
    // 🎯 تعریف تابع پردازش‌گر آپدیت‌ها (Handler)
    // این تابع برای هر آپدیت جدید از سمت سرور روبیکا فراخوانی می‌شود
    handler := func(ctx context.Context, update *ParsRubika.Update) error {
        // 📬 فقط پیام‌های جدید را پردازش کن
        if update.NewMessage != nil {
            message := update.NewMessage
            log.Printf("📩 پیام جدید از کاربر %s: %s", update.ChatID, message.Text)
            
            // 📤 ارسال پاسخ (همان متن پیام دریافتی)
            _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
                ChatID: update.ChatID,
                Text:   "شما نوشتید: " + message.Text,
            })
            if err != nil {
                log.Printf("❌ خطا در ارسال پاسخ: %v", err)
            }
        }
        return nil
    }
    
    // 🚀 شروع فرآیند دریافت آپدیت‌ها با متد Polling
    log.Println("✅ بات با موفقیت راه‌اندازی شد و منتظر پیام‌هاست...")
    err := bot.StartPolling(ctx, ParsRubika.PollingOptions{
        Handler:      handler,       // تابع پردازش‌گر
        PollInterval: 3 * time.Second, // فاصله زمانی بین هر درخواست
        Limit:        100,           // حداکثر تعداد آپدیت در هر بار
    })
    
    // ⚠️ بررسی خطا در صورت بروز (غیر از خطای لغو کانتکست)
    if err != nil && err != context.Canceled {
        log.Fatalf("💥 بات با خطا متوقف شد: %v", err)
    }
    
    log.Println("👋 بات با موفقیت متوقف شد.")
}
```

================================================================================
5. 🏗️ مفاهیم اصلی: Client, Polling, Webhook
================================================================================

### 🤖 Client (کلاینت)
Client قلب کتابخانه است و تمام درخواست‌ها به API روبیکا از طریق آن انجام می‌شود.

```go
// ساخت کلاینت ساده
bot := ParsRubika.NewClient("YOUR_BOT_TOKEN")

// ساخت کلاینت با تنظیمات سفارشی
bot := ParsRubika.NewClient("YOUR_BOT_TOKEN",
    ParsRubika.WithBaseURL("https://custom-api-url.com"),
    ParsRubika.WithHTTPClient(&http.Client{Timeout: 60 * time.Second}),
)
```

### 🔄 Polling (دریافت آپدیت)
در این متد، کتابخانه به صورت مداوم از سرور روبیکا آپدیت‌های جدید را درخواست می‌کند. این روش ساده‌ترین راه برای شروع کار است و در مثال "شروع سریع" مشاهده کردید.

### 🌐 Webhook (دریافت آپدیت با وب‌هوک)
در این متد، شما یک سرور HTTP راه‌اندازی می‌کنید و سرور روبیکا آپدیت‌های جدید را به صورت مستقیم به شما ارسال می‌کند. این روش برای بات‌های با ترافیک بالا بسیار مناسب‌تر است.

```go
// ... (تعریف handler مانند مثال قبل)

webhookOpts := ParsRubika.WebhookOptions{
    Port:    8080,             // پورت سرور شما
    Path:    "/webhook",       // مسیری که روبیکا باید به آن درخواست ارسال کند
    Handler: handler,          // تابع پردازش‌گر آپدیت‌ها
}

log.Println("🌐 سرور وب‌هوک روی پورت 8080 در حال راه‌اندازی است...")
err := bot.StartWebhook(ctx, webhookOpts)
if err != nil {
    log.Fatalf("💥 خطا در راه‌اندازی وب‌هوک: %v", err)
}
```

================================================================================
6. 📖 مستندات کامل API
================================================================================

--- 📤 ارسال پیام (متن، کیبورد، مدیا) ---

**مثال ۱: ارسال پیام ساده**
```go
messageID, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID: "USER_OR_CHAT_ID",
    Text:   "سلام! به ربات من خوش آمدید 🎉",
})
```

**مثال ۲: ارسال پیام با کیبورد اینلاین (Inline Keyboard)**
```go
inlineKeyboard := &ParsRubika.Keypad{
    Rows: []ParsRubika.KeypadRow{
        {
            Buttons: []ParsRubika.Button{
                {ID: "btn_like", Type: ParsRubika.ButtonTypeSimple, ButtonText: "👍 لایک"},
                {ID: "btn_dislike", Type: ParsRubika.ButtonTypeSimple, ButtonText: "👎 دیسلایک"},
            },
        },
        {
            Buttons: []ParsRubika.Button{
                {ID: "btn_help", Type: ParsRubika.ButtonTypeSimple, ButtonText: "❓ راهنما"},
            },
        },
    },
}

messageID, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:       "USER_OR_CHAT_ID",
    Text:         "لطفاً یک گزینه را انتخاب کنید:",
    InlineKeypad: inlineKeyboard,
})
```

**مثال ۳: ارسال پیام با کیبورد اصلی (Chat Keyboard)**
```go
replyKeyboard := &ParsRubika.Keypad{
    Rows: []ParsRubika.KeypadRow{
        {
            Buttons: []ParsRubika.Button{
                {ID: "btn_profile", Type: ParsRubika.ButtonTypeSimple, ButtonText: "پروفایل"},
                {ID: "btn_settings", Type: ParsRubika.ButtonTypeSimple, ButtonText: "تنظیمات"},
            },
        },
    },
    ResizeKeyboard: true, // کیبورد کوچک‌تر نمایش داده شود
}

messageID, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
    ChatID:         "USER_OR_CHAT_ID",
    Text:           "منوی اصلی:",
    ChatKeypad:     replyKeyboard,
    ChatKeypadType: ParsRubika.NewKeypad, // نوع کیبورد: جدید
})
```

--- 📁 مدیریت فایل (آپلود، ارسال، دانلود) ---

**مثال ۱: آپلود و ارسال فایل مستقیم**
```go
// ابتدا فایل را آپلود می‌کنیم
uploadedFile, err := bot.UploadFileDirectly(ctx, "/path/to/image.jpg", ParsRubika.ImageType)
if err != nil {
    log.Fatal(err)
}

// سپس با استفاده از file_id دریافتی، آن را ارسال می‌کنیم
messageID, err := bot.SendFile(ctx, &ParsRubika.SendFileRequest{
    ChatID: "USER_OR_CHAT_ID",
    FileID: uploadedFile.FileID,
    Text:   "این یک تصویر است که مستقیماً آپلود شد.",
})
```

**مثال ۲: ارسال فایل با استفاده از FileID موجود**
```go
messageID, err := bot.SendFile(ctx, &ParsRubika.SendFileRequest{
    ChatID: "USER_OR_CHAT_ID",
    FileID: "EXISTING_FILE_ID_FROM_RUBIKA",
    Text:   "این یک فایل از قبل موجود است.",
})
```

**مثال ۳: دریافت اطلاعات فایل (برای دانلود)**
```go
fileInfo, err := bot.GetFile(ctx, "FILE_ID")
if err == nil {
    log.Printf("لینک دانلود فایل: %s", fileInfo.DownloadURL)
    log.Printf("نام فایل: %s", fileInfo.FileName)
}
```

--- 📍 ارسال موقعیت، مخاطب و نظرسنجی ---

**مثال ۱: ارسال موقعیت مکانی (Location)**
```go
messageID, err := bot.SendLocation(ctx, &ParsRubika.SendLocationRequest{
    ChatID:    "USER_OR_CHAT_ID",
    Latitude:  "35.6892",  // عرض جغرافیایی
    Longitude: "51.3890", // طول جغرافیایی
})
```

**مثال ۲: ارسال اطلاعات تماس (Contact)**
```go
messageID, err := bot.SendContact(ctx, &ParsRubika.SendContactRequest{
    ChatID:      "USER_OR_CHAT_ID",
    PhoneNumber: "+989123456789",
    FirstName:   "ابوالفضل",
    LastName:    "زارعی",
})
```

**مثال ۳: ایجاد نظرسنجی (Poll)**
```go
pollID, err := bot.SendPoll(ctx, &ParsRubika.SendPollRequest{
    ChatID:   "USER_OR_CHAT_ID",
    Question: "بهترین زبان برنامه‌نویسی چیست؟",
    Options:  []string{"Go", "Python", "JavaScript", "Rust"},
})
```

--- ✏️ ویرایش، حذف و فوروارد پیام ---

**مثال ۱: ویرایش متن یک پیام**
```go
err := bot.EditMessageText(ctx, &ParsRubika.EditMessageTextRequest{
    ChatID:    "USER_OR_CHAT_ID",
    MessageID: "MESSAGE_ID_TO_EDIT",
    Text:      "متن جدید و ویرایش شده! ✍️",
})
```

**مثال ۲: ویرایش کیبورد اینلاین یک پیام**
```go
newKeyboard := &ParsRubika.Keypad{ /* ... تعریف کیبورد جدید ... */ }
err := bot.EditMessageKeypad(ctx, &ParsRubika.EditMessageKeypadRequest{
    ChatID:       "USER_OR_CHAT_ID",
    MessageID:    "MESSAGE_ID_TO_EDIT",
    InlineKeypad: newKeyboard,
})
```

**مثال ۳: حذف یک پیام**
```go
err := bot.DeleteMessage(ctx, &ParsRubika.DeleteMessageRequest{
    ChatID:    "USER_OR_CHAT_ID",
    MessageID: "MESSAGE_ID_TO_DELETE",
})
```

**مثال ۴: فوروارد کردن یک پیام**
```go
newMessageID, err := bot.ForwardMessage(ctx, &ParsRubika.ForwardMessageRequest{
    FromChatID: "SOURCE_CHAT_ID", // آیدی چتی که پیام از آنجا است
    MessageID:  "MESSAGE_ID_TO_FORWARD",
    ToChatID:   "DESTINATION_CHAT_ID", // آیدی چتی که پیام باید به آن ارسال شود
})
```

--- 👥 مدیریت گروه و کانال ---

**مثال ۱: دریافت اطلاعات یک چت**
```go
chatInfo, err := bot.GetChat(ctx, "CHAT_ID")
if err == nil {
    log.Printf("عنوان چت: %s", chatInfo.Title)
    log.Printf("نوع چت: %s", chatInfo.ChatType)
}
```

**مثال ۲: مسدود کردن (Ban) یک عضو**
```go
err := bot.BanChatMember(ctx, &ParsRubika.BanChatMemberRequest{
    ChatID: "GROUP_OR_CHANNEL_ID",
    UserID: "USER_ID_TO_BAN",
})
```

**مثال ۳: ترفیع دادن یک عضو به مدیر (Promote)**
```go
isTrue := true
err := bot.PromoteChatMember(ctx, &ParsRubika.PromoteChatMemberRequest{
    ChatID:              "GROUP_OR_CHANNEL_ID",
    UserID:              "USER_ID_TO_PROMOTE",
    IsAdministrator:     &isTrue,
    CanChangeInfo:       &isTrue,
    CanDeleteMessages:   &isTrue,
    CanPinMessages:      &isTrue,
})
```

**مثال ۴: دریافت لیست مدیران یک گروه/کانال**
```go
admins, err := bot.GetChatAdministrators(ctx, &ParsRubika.GetChatAdministratorsRequest{
    ChatID: "GROUP_OR_CHANNEL_ID",
})
if err == nil {
    for _, admin := range admins.Administrators {
        log.Printf("مدیر: %s (%s)", admin.User.FirstName, admin.User.UserID)
    }
}
```

--- 📱 مدیریت صفحات و استوری‌ها (توابع غیررسمی) ---
این توابع به شما امکان می‌دهند عملیات مربوط به صفحات شخصی، پست‌ها و استوری‌ها را انجام دهید.

**مثال ۱: لایک کردن یک پست**
```go
err := bot.LikePost(ctx, "POST_ID")
```

**مثال ۲: افزودن یک استوری جدید**
```go
err := bot.AddStory(ctx, "UPLOADED_FILE_ID", "این کپشن استوری است 🎬")
```

**مثال ۳: دریافت پست‌های یک پروفایل**
```go
posts, err := bot.GetProfilePosts(ctx, "USER_PAGE_ID")
if err == nil {
    for _, post := range posts {
        log.Printf("پست یافت شد: %s", post.PostID)
    }
}
```

================================================================================
7. 🎯 مثال‌های پیشرفته
================================================================================

--- 🧠 مدیریت State (حالت کاربر) در فرم‌های چند مرحله‌ای ---
گاهی اوقات نیاز دارید اطلاعات کاربر را بین پیام‌های مختلف ذخیره کنید (مثلاً در یک فرم چند مرحله‌ای). کتابخانه یک StateManager ساده برای این کار ارائه می‌دهد.

```go
// در ابتدای برنامه StateManager را بسازید
stateManager := ParsRubika.NewStateManager()

// در Handler خود
handler := func(ctx context.Context, update *ParsRubika.Update) error {
    if update.NewMessage == nil { return nil }
    
    userID := update.ChatID
    messageText := update.NewMessage.Text
    
    // مرحله اول: دریافت نام
    if messageText == "/start" {
        stateManager.SetState(userID, "step", "get_name")
        _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: userID,
            Text:   "سلام! لطفاً نام خود را وارد کنید:",
        })
        return err
    }
    
    // بررسی مرحله فعلی کاربر
    currentStep, found := stateManager.GetState(userID, "step")
    if !found { return nil } // اگر در فرآیندی نبود، کاری نکن
    
    // مرحله دریافت نام
    if currentStep == "get_name" {
        stateManager.SetState(userID, "name", messageText)
        stateManager.SetState(userID, "step", "get_age")
        _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: userID,
            Text:   "عالی! حالا لطفاً سن خود را وارد کنید:",
        })
        return err
    }
    
    // مرحله دریافت سن
    if currentStep == "get_age" {
        name, _ := stateManager.GetState(userID, "name")
        
        // نمایش نتیجه نهایی و پاک کردن state
        _, err := bot.SendMessage(ctx, &ParsRubika.SendMessageRequest{
            ChatID: userID,
            Text:   fmt.Sprintf("تشکر! نام شما %s و سن شما %s سال است.", name, messageText),
        })
        
        stateManager.DeleteUserState(userID) // پاک کردن اطلاعات کاربر
        return err
    }
    
    return nil
}
```

--- 🎨 ساخت کیبوردهای پیشرفته ---
کتابخانه از انواع دکمه‌های پیشرفته پشتیبانی می‌کند.

**مثال ۱: دکمه انتخابگر (Selection Button)**
```go
selectionButton := &ParsRubika.ButtonSelection{
    SelectionID: "color_selector",
    Title:       "رنگ مورد نظر خود را انتخاب کنید:",
    Items: []ParsRubika.ButtonSelectionItem{
        {Text: "قرمز", ImageUrl: "URL_TO_RED_IMAGE"},
        {Text: "سبز", ImageUrl: "URL_TO_GREEN_IMAGE"},
        {Text: "آبی", ImageUrl: "URL_TO_BLUE_IMAGE"},
    },
}

keyboard := &ParsRubika.Keypad{
    Rows: []ParsRubika.KeypadRow{
        {Buttons: []ParsRubika.Button{
            ID:              "btn_select_color",
            Type:            ParsRubika.ButtonTypeSelection,
            ButtonText:      "انتخاب رنگ",
            ButtonSelection: selectionButton,
        }},
    },
}
// ... ارسال پیام با این کیبورد
```

**مثال ۲: دکمه تقویم (Calendar Button)**
```go
calendarButton := &ParsRubika.ButtonCalendar{
    Title:   "تاریخ تولد خود را انتخاب کنید:",
    Type:    "DatePersian", // یا "DateGregorian"
    MinYear: "1350",
    MaxYear: "1410",
}

keyboard := &ParsRubika.Keypad{
    Rows: []ParsRubika.KeypadRow{
        {Buttons: []ParsRubika.Button{
            ID:             "btn_calendar",
            Type:           ParsRubika.ButtonTypeCalendar,
            ButtonText:     "انتخاب تاریخ",
            ButtonCalendar: calendarButton,
        }},
    },
}
// ... ارسال پیام با این کیبورد
```

================================================================================
8. 🤝 مشارکت در پروژه و لایسنس
================================================================================

🤝 **مشارکت:**
ما از مشارکت‌های شما استقبال می‌کنیم! اگر پیشنهادی برای بهبود دارید، باگ پیدا کردید یا می‌خواهید قابلیت جدیدی اضافه کنید، لطفاً یک Issue یا Pull Request در گیت‌هاب ایجاد کنید.

📄 **لایسنس:**
این پروژه تحت لایسنس MIT منتشر شده است. برای اطلاعات بیشتر، فایل LICENSE را در مخزن گیت‌هاب مشاهده کنید.

================================================================================
9. 📞 ارتباط با ما
================================================================================
گیت‌هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

با تشکر از استفاده شما از این کتابخانه! ❤️
```
