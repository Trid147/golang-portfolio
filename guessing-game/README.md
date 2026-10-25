# CLI Number Guessing Game

A clean, interactive command-line terminal game written in Go. The application securely selects a random integer within a configurable limit and challenges the user to deduce it within a set number of attempts.

## ⚙️ Key Technical Features

Despite being a classic foundational project, this implementation applies idiomatic Go principles to handle real-world terminal usage:

1. **Modern Randomization:** Uses the updated standard library `math/rand/v2` engine for clean, pseudo-random integer distribution without needing legacy manual seed tracking.
2. **Robust Input Validation:** Wraps user input parsing inside an infinite validation loop. If a user inputs non-integer garbage strings (like characters or symbols), the engine catches the error, discards the broken terminal buffer via a temporary string scanner, and recovers seamlessly without panicking.
3. **Control Flow Loops:** Manages states and guess boundaries utilizing standard `switch` conditional branches and Go's versatile unified `for` loop patterns.

---

## 🛠️ Configuration & Rules

By default, the game is configured at the top of the `main()` function block with the following standard variables:

* **Max Number Range:** 1 to `100` (inclusive)
* **Maximum Attempts Allowed:** `7`

You can easily adjust these variables inside `main.go` to scale the game difficulty up or down.

---

## 🚀 How to Run

Ensure you have Go installed on your machine.

1. Open your Linux terminal and navigate into this project directory.
2. Launch the game using the native compiler:
   ```bash
   go run main.go
   ```

### Example Gameplay Interface
```text
I have thought of an integer. Try to guess it! You have 7 attempts.
Your guess: hello
Write a valid integer!
Your guess: 50
That's not right. My integer is bigger
Your attempts: 6
Your guess: 75
...
```
