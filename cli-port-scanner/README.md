# Simple Concurrent Port Scanner

A lightweight, concurrent network utility written in Go designed to audit active TCP services on a single host. Instead of scanning ports one by one, this tool leverages Go's native goroutines to scan all specified ports simultaneously.

## ⚙️ How It Works (Architecture)

The tool processes target ports concurrently using standard Go primitives:

1. **Goroutine Per Port:** The program spawns an independent background thread (goroutine) for every single port in the list.
2. **Safe Dialing:** Each goroutine attempts a connection using `net.DialTimeout`. If the connection succeeds, it securely closes the socket to avoid resource leaks.
3. **Synchronized Channel:** A thread-safe channel (`portChan`) collects the open port numbers as the goroutines find them.
4. **WaitGroup Control:** A `sync.WaitGroup` keeps the main program running until every single port scan finishes.
5. **Sorted Output:** The final collection is sorted sequentially before printing to keep the terminal results readable.

---

## 🚀 How to Run

Ensure you have Go installed on your machine.

1. Open `main.go` and adjust the target IP or port array inside the `main()` function configuration block if needed:
   ```go
   var IP = "192.168.0.1"
   ports := []int{21, 22, 23, 25, 80, 110, 139, 443, 445, 3389}
   ```

2. Run the application from your terminal:
   ```bash
   go run main.go
   ```

---

## 📈 Planned Roadmap
* [ ] Integrate Go's `flag` package to allow passing IPs directly from the CLI.
* [ ] Implement a full `/24` subnet loop to scan up to 254 hosts.
* [ ] Restructure the concurrency layout into a fixed `Worker Pool` to control network connection spikes safely.
