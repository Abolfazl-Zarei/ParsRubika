package ParsRubika

// سازنده ابوالفضل زارعی
// آدرس گیت هاب: https://github.com/Abolfazl-Zarei/ParsRubika-bot-go

import (
	"log"
	"sync"
)

type StateManager struct {
	states map[string]map[string]interface{}
	mu     sync.RWMutex
}

func NewStateManager() *StateManager {
	return &StateManager{
		states: make(map[string]map[string]interface{}),
	}
}

func (sm *StateManager) SetState(userID, key string, value interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if _, ok := sm.states[userID]; !ok {
		sm.states[userID] = make(map[string]interface{})
	}
	sm.states[userID][key] = value
	log.Printf("State: Set for user %s, key '%s'", userID, key)
}

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

func (sm *StateManager) DeleteState(userID, key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if userState, ok := sm.states[userID]; ok {
		delete(userState, key)
	}
}

func (sm *StateManager) DeleteUserState(userID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.states, userID)
}
