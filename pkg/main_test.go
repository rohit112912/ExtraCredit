package main

import (
      	"testing"
	"sync"
)

// Helper function to reset the Bank singleton instance
func resetBankInstance() {
	bankInstance = nil
	once = sync.Once{}
}

func TestInitializeBank(t *testing.T) {
	resetBankInstance()
	bank := InitializeBank()
	if bank == nil {
		t.Error("Failed to initialize bank instance")
	}
}

func TestGetBankInstance(t *testing.T) {
	resetBankInstance()
	_ = InitializeBank() // Ensure the instance is initialized
	bank := GetBankInstance()
	if bank == nil {
		t.Error("Failed to retrieve bank instance")
	}
}

func TestDeposit(t *testing.T) {
	resetBankInstance()
	bank := InitializeBank()
	bank.Deposit(100)
	if bank.GetBalance() != 100 {
		t.Errorf("Expected balance 100, got %d", bank.GetBalance())
	}
}

func TestWithdraw(t *testing.T) {
	resetBankInstance()
	bank := InitializeBank()
	bank.Deposit(100)

	err := bank.Withdraw(50)
	if err != nil {
		t.Errorf("Withdraw failed with error: %v", err)
	}

	if bank.GetBalance() != 50 {
		t.Errorf("Expected balance 50, got %d", bank.GetBalance())
	}

	err = bank.Withdraw(100)
	if err == nil {
		t.Error("Expected insufficient funds error, got nil")
	}
}

func TestGetBalance(t *testing.T) {
	resetBankInstance()
	bank := InitializeBank()
	bank.Deposit(200)
	if bank.GetBalance() != 200 {
		t.Errorf("Expected balance 200, got %d", bank.GetBalance())
	}
}

