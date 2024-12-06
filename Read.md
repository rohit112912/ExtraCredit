# Concurrent Bank Account Program

This Go program demonstrates the use of concurrency mechanisms such as **goroutines**, **sync.Mutex**, and **sync.Once** to simulate a concurrent bank account system. It is designed to handle multiple deposits and withdrawals simultaneously, ensuring the safety and consistency of shared variables.

## Features

- **Concurrency with Goroutines**: Multiple goroutines perform concurrent operations like deposits and withdrawals.
- **Mutual Exclusion with sync.Mutex**: Ensures safe access to the shared `balance` variable to prevent data races.
- **One-Time Initialization with sync.Once**: Guarantees that the logger is initialized only once, regardless of the number of goroutines.
- **Race Condition Detection**: Program has been tested with the Go race detector to ensure no data races occur.
- **Concurrency-Safe Logging**: Logs transactions safely using a global logger.

---

## File Structure

- `main.go`: Contains the program's source code, implementing the concurrent bank account logic.
- `go.mod`: Go module file for dependency management.
- `.github/workflows/go.yml`: GitHub Actions workflow file to automate testing and building.

---

## Running the Program

1. **Prerequisites**:
   - Install [Go](https://golang.org/dl/) (version 1.20 or later).
   - Ensure your environment supports Go commands.

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/<your-username>/ExtraCredit-Concurrency.git
   cd ExtraCredit-Concurrency

