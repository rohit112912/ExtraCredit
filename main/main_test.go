// File: main_test.go
package main

import (
	"sync"
	"testing"
)

func TestBankOperations(t *testing.T) {
	var once sync.Once
	once.Do(InitializeBank)

	b := GetBankInstance()

	// Reset balance for test.
	b.mu.Lock()
	b.balance = 0
	b.mu.Unlock()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(amount int) {
			defer wg.Done()
			b.Deposit(amount)
		}(10)

		go func(amount int) {
			defer wg.Done()
			b.Withdraw(amount)
		}(5)
	}
	wg.Wait()

	expectedBalance := 50
	actualBalance := b.Balance()
	if actualBalance != expectedBalance {
		t.Errorf("Expected balance %d, got %d", expectedBalance, actualBalance)
	}
}
