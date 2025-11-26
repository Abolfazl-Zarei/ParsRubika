package ParsRubika

// سازنده: ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import (
	"log"
	"sync"
	"time"
)

// ReloadManager مدیریت قابلیت بارگذاری مجدد (Hot-Reload)
type ReloadManager struct {
	client      *BotClient
	isWatching  bool
	mu          sync.Mutex
	reloadFuncs []func() // لیستی از توابعی که پس از هر بارگذاری مجدد اجرا می‌شوند
}

// NewReloadManager ایجاد یک نمونه جدید از ReloadManager
func NewReloadManager(client *BotClient) *ReloadManager {
	return &ReloadManager{
		client:      client,
		reloadFuncs: make([]func(), 0),
	}
}

// StartWatching شروع نظارت بر تغییرات فایل‌ها
func (rm *ReloadManager) StartWatching() {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if rm.isWatching {
		log.Println("Hot-Reload در حال حاضر در حال اجراست.")
		return
	}

	if !rm.client.IsHotReloadEnabled() {
		log.Println("Hot-Reload در کلاینت غیرفعال است. برای فعال‌سازی از گزینه WithHotReload استفاده کنید.")
		return
	}

	rm.isWatching = true
	log.Println("Hot-Reload watcher شروع به کار کرد.")

	// در یک پیاده‌سازی واقعی، در اینجا از کتابخانه‌ای مانند fsnotify استفاده می‌شود
	// تا تغییرات فایل‌های .go در پروژه نظارت شود.
	// برای مثال:
	// watcher, err := fsnotify.NewWatcher()
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer watcher.Close()
	//
	// go func() {
	//     for {
	//         select {
	//         case event, ok := <-watcher.Events:
	//             if !ok {
	//                 return
	//             }
	//             if event.Op&fsnotify.Write == fsnotify.Write {
	//                 log.Println("فایل تغییر کرده:", event.Name)
	//                 rm.TriggerReload()
	//             }
	//         case err, ok := <-watcher.Errors:
	//            if !ok {
	//                 return
	//             }
	//             log.Println("خطا در watcher:", err)
	//         }
	//     }
	// }()
	//
	// err = watcher.Add(".")
	// if err != nil {
	//     log.Fatal(err)
	// }

	// برای این مثال، ما فقط یک پیام لاگ می‌کنیم و یک تیکر شبیه‌سازی می‌کنیم
	go func() {
		ticker := time.NewTicker(10 * time.Second) // هر 10 ثانیه یک بار بررسی کن
		defer ticker.Stop()
		for {
			select {
			case <-rm.client.stopChan:
				log.Println("Hot-Reload watcher متوقف شد.")
				return
			case <-ticker.C:
				// در اینجا باید منطق بررسی تغییرات فایل قرار گیرد
				// rm.checkForChanges()
			}
		}
	}()
}

// TriggerReload اجرای فرآیند بارگذاری مجدد
func (rm *ReloadManager) TriggerReload() {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	log.Println("بارگذاری مجدد (Hot-Reload) فعال شد...")

	// در اینجا می‌توان منطق‌های مربوط به بارگذاری مجدد هندلرها، میدلورها و غیره را پیاده‌سازی کرد.
	// برای مثال، می‌توان فایل‌های کد را دوباره کامپایل و بخش‌های مربوطه را در حافظه جایگزین کرد.

	// اجرای توابع ثبت شده برای پس از بارگذاری مجدد
	for _, fn := range rm.reloadFuncs {
		if fn != nil {
			fn()
		}
	}

	log.Println("بارگذاری مجدد با موفقیت انجام شد.")
}

// OnReload ثبت یک تابع برای اجرا پس از هر بارگذاری مجدد
func (rm *ReloadManager) OnReload(fn func()) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.reloadFuncs = append(rm.reloadFuncs, fn)
}

// StopWatching توقف نظارت بر تغییرات فایل‌ها
func (rm *ReloadManager) StopWatching() {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if !rm.isWatching {
		return
	}

	rm.isWatching = false
	log.Println("Hot-Reload watcher متوقف شد.")
	// در اینجا منابع مربوط به watcher (مانند fsnotify) باید آزاد شوند.
}
