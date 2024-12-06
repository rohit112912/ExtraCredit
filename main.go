package main

import (
	"fmt"
	"log"
	"sync"
)

// BankAccount struct with balance and mutex
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

var once sync.Once
var logger *log.Logger

// InitializeLogger ensures logger is initialized only once
func InitializeLogger() {
	logger = log.Default()
	logger.Println("Logger initialized.")
}

// Deposit adds amount to the balance
func (acc *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	acc.mu.Lock()
	defer acc.mu.Unlock()

	acc.balance += amount
	logger.Printf("Deposited %d, New Balance: %d\n", amount, acc.balance)
}

// Withdraw subtracts amount from the balance
func (acc *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if amount > acc.balance {
		logger.Printf("Failed to withdraw %d, Insufficient Balance: %d\n", amount, acc.balance)
		return
	}

	acc.balance -= amount
	logger.Printf("Withdrew %d, New Balance: %d\n", amount, acc.balance)
}

// GetBalance safely retrieves the balance
func (acc *BankAccount) GetBalance() int {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	return acc.balance
}

func main() {
	// Ensure logger is initialized once
	once.Do(InitializeLogger)

	var wg sync.WaitGroup
	account := &BankAccount{balance: 1000} // Initial balance

	// Simulate concurrent deposits and withdrawals
	wg.Add(5)
	go account.Deposit(500, &wg)
	go account.Withdraw(200, &wg)
	go account.Withdraw(1500, &wg) // Should fail due to insufficient balance
	go account.Deposit(100, &wg)
	go account.Withdraw(300, &wg)

	// Wait for all goroutines to complete
	wg.Wait()

	// Print final balance
	fmt.Printf("Final Balance: %d\n", account.GetBalance())
}

