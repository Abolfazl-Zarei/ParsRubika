# راهنمای جامع کتابخانه ParsRubika

این راهنما شما را برای استفاده از کتابخانه `ParsRubika` برای ساخت ربات‌های روبیکا راهنمایی می‌کند.

## فهرست مطالب

- [مقدمه](#مقدمه)
- [نصب و راه‌اندازی](#نصب-و-راهاندازی)
- [شروع سریع](#شروع-سریع)
- [مفاهیم اصلی](#مفاهیم-اصلی)
    - [کلاینت (Client)](#کلاینت-client)
    - [دریافت آپدیت‌ها (Polling و Webhook)](#دریافت-آپدیتها-polling-و-webhook)
    - [مدیریت وضعیت (State Management)](#مدیریت-وضعیت-state-management)
- [مرجع API](#مرجع-api)
    - [انواع دکمه‌ها (ButtonTypeEnum)](#انواع-دکمهها-buttontypeenum)
    - [انواع فایل‌ها (FileTypeEnum)](#انواع-فایلها-filetypeenum)
    - [انواع چت (ChatTypeEnum)](#انواع-چت-chattypeenum)
    - [مدل‌های داده (Models)](#مدلهای-داده-models)
    - [متدهای رسمی API بات](#متدهای-رسمی-api-بات)
    - [متدهای غیررسمی (مدیریت صفحه و استوری)](#متدهای-غیررسمی-مدیریت-صفحه-و-استوری)
- [مثال کامل](#مثال-کامل)

---

## مقدمه

کتابخانه `ParsRubika` یک پوشش (wrapper) کامل و ساده برای API رسمی ربات‌های روبیکا و همچنین برخی متدهای غیررسمی است. این کتابخانه به شما اجازه می‌دهد تا به راحتی ربات‌های تعاملی، مدیریت‌گر گروه و کانال، یا ربات‌های پیچیده بسازید.

## نصب و راه‌اندازی

برای استفاده از این کتابخانه، ابتدا آن را به پروژه گولنگ خود اضافه کنید. ساده‌ترین راه، کلون کردن از گیت‌هاب است:

```bash
git clone https://github.com/Abolfazl-Zarei/ParsRubika-bot-go.git


#آموزش ران کردن ربات polling 
$env:BOT_TOKEN="BOT_TOKEN"
go mod tidy
go run examples/{nameBot}.go -mode=polling
go run examples/bot1.go -mode=polling
go run examples/bot2.go -mode=polling
go run examples/bot3.go -mode=polling
go run examples/bot4.go -mode=polling



#آموزش ران کردن ربات webhook 
# آدرس ngrok خود را اینجا قرار دهید
$env:BOT_TOKEN="BOT_TOKEN" 
PUBLIC_URL="https://random-string.ngrok-free.app" go run {nameBot}.go -mode=webhook

PUBLIC_URL="https://random-string.ngrok-free.app" go run bot1.go -mode=webhook
PUBLIC_URL="https://random-string.ngrok-free.app" go run bot2.go -mode=webhook
PUBLIC_URL="https://random-string.ngrok-free.app" go run bot3.go -mode=webhook
PUBLIC_URL="https://random-string.ngrok-free.app" go run bot4.go -mode=webhook


#حالا باید ببینید که بات با موفقیت راه‌اندازی شده و وب‌هوک ثبت شده است.
Registering webhook URL: https://random-string.ngrok-free.app/rubika-webhook
✅ Webhook registered successfully.
🌐 Starting webhook server on port 8080...

