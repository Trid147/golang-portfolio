# CLI Markdown Link Extractor

A robust, lightweight command-line tool written in Go to parse and isolate raw URLs from Markdown-formatted hyperlinks. This application is built using the industry-standard **Cobra** CLI framework, matching the design principles of enterprise open-source software like Kubernetes (`kubectl`) and GitHub CLI (`gh`).

## ⚙️ How It Works (Parsing Logic)

Instead of relying on heavy regular expressions, this project handles string manipulation efficiently using Go's optimized standard `strings` package:

1. **Cobra Validations:** Uses `cobra.ExactArgs(1)` to automatically validate terminal arguments, generating clean, standard user help messages if inputs are missing.
2. **Context-Aware Matching:** Looks explicitly for the `](` character junction. This completely prevents parsing conflicts with descriptive text containing parentheses ahead of the link (e.g., `(Update) [Google](https://google.com)` resolves seamlessly to `https://google.com`).
3. **Safe Slicing:** Dynamically calculates lookahead index positions relative to the markdown markers, avoiding negative bounds or out-of-range slicing errors that would trigger a runtime panic.

---

## 🚀 Installation & Usage

Ensure you have a working Go installation and have fetched project dependencies using `go mod tidy`.

### Extracting Links
Pass the Markdown snippet or raw link string directly as a command-line argument:

```bash
go run main.go "[GitHub](https://github.com)"
```

### Edge-Case Resiliency Example
The custom structural parsing reliably avoids false-positive markers in the text prefix:

```bash
go run main.go "(Hotfix v1.0) [Go Language](https://go.dev)"
# Output -> Result: https://go.dev
```

### Native Built-in Help System
Since this application relies on Cobra structures, it automatically yields a clean, standardized help menu if prompted with the standard flags:

```bash
go run main.go --help
```