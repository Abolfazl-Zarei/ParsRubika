package ParsRubika

// سازنده ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import (
	"log"
	"sync"
)

// StateManager برای مدیریت وضعیت‌های کاربران
type StateManager struct {
	states map[string]map[string]interface{} // نگهداری وضعیت هر کاربر
	mu     sync.RWMutex                      // قفل برای دسترسی امن همزمان
}

// NewStateManager یک نمونه جدید از StateManager ایجاد می‌کند
func NewStateManager() *StateManager {
	return &StateManager{
		states: make(map[string]map[string]interface{}),
	}
}

// SetState یک مقدار را برای کاربر مشخص ذخیره می‌کند
func (sm *StateManager) SetState(userID, key string, value interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if _, ok := sm.states[userID]; !ok {
		sm.states[userID] = make(map[string]interface{})
	}
	sm.states[userID][key] = value
	log.Printf("State: Set for user %s, key '%s'", userID, key)
}

// GetState یک مقدار را برای کاربر و کلید مشخص برمی‌گرداند
func (sm *StateManager) GetState(userID, key string) (interface{}, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	if userState, ok := sm.states[userID]; ok {
		if value, ok := userState[key]; ok {
			return value, true
		}
	}
	return nil, false
}

// DeleteState یک کلید را برای کاربر مشخص حذف می‌کند
func (sm *StateManager) DeleteState(userID, key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if userState, ok := sm.states[userID]; ok {
		delete(userState, key)
	}
}

// DeleteUserState تمام وضعیت‌های یک کاربر را حذف می‌کند
func (sm *StateManager) DeleteUserState(userID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.states, userID)
}
