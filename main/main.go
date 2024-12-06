package main

import (
	"fmt"
	"sync"
)

// Bank struct representing a bank with balance
type Bank struct {
	balance int
	mu      sync.Mutex
}

// Singleton instance and sync.Once to ensure one-time initialization
var bankInstance *Bank
var once sync.Once

// InitializeBank initializes the singleton Bank instance
func InitializeBank() *Bank {
	once.Do(func() {
		bankInstance = &Bank{balance: 0}
	})
	return bankInstance
}

// GetBankInstance retrieves the singleton Bank instance
func GetBankInstance() *Bank {
	return bankInstance
}

// Deposit adds an amount to the bank's balance
func (b *Bank) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}

// Withdraw subtracts an amount from the bank's balance
func (b *Bank) Withdraw(amount int) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if amount > b.balance {
		return fmt.Errorf("insufficient funds")
	}
	b.balance -= amount
	return nil
}

// GetBalance returns the current balance of the bank
func (b *Bank) GetBalance() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.balance
}

