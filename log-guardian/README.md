# Log Guardian: Real-Time Log Watcher Daemon

`log-guardian` is a lightweight, real-time security daemon written in Go designed to inspect web server access logs dynamically as they append. Utilizing the enterprise-grade **Cobra** CLI framework and high-performance file polling abstractions, it detects common web attack signatures natively and streams security context into structured `.jsonl` audit records for downstream incident response pipelines.

---

## ⚙️ Core Technical Features

1. **Persistent Tail Engine:** Employs the `hpcloud/tail` event wrapper configured with `os.SEEK_END` rulesets. The utility skips stale historic lines on startup, monitoring only live runtime additions without memory leak degradation.
2. **Contextual Token Slicing:** Avoids common array slicing out-of-bounds panics by implementing lookahead token gates when string-indexing Nmap or custom bracket sequences (`[...]`) to parse timestamps natively.
3. **Multi-Vector Threat Signature Matrix:** Evaluates log payloads against a hardcoded lookup index to categorize security compromise attempts:
   * **Path Traversal:** Identifies directory escapes (`/etc/passwd`).
   * **SQL Injection (SQLi):** Flags standard token manipulation signatures (`UNION SELECT`).
   * **Cross-Site Scripting (XSS):** Filters reflective HTML element injection markers (`<script>`).
   * **Remote Command Execution (RCE):** Captures automated exploit directory crawling payloads (`cgi-bin`).

---

## 📊 Structured JSONL Telemetry

When a malicious vector maps to a known signature, the scanner generates a strict JSON tracking schema to ensure third-party tools (like SIEM aggregators or automated firewall ban triggers) can parse incident indicators natively:

```json
{"ip":"10.0.0.99","timestamp":"04/Nov/2026:19:15:44 +0300","attack_type":"SQL Injection"}
```

---

## 🚀 Setup & Execution Blueprint

### 1. Initialize Log Pipeline Simulation
To simulate a production infrastructure load environment, execute your custom multi-vector log generator script in a separate background terminal window:

```bash
chmod +x model.sh
./model.sh
```
This loop dynamically appends clean user visits interspersed with malicious attacker payloads into `mock_access.log`.

### 2. Launch the Security Daemon
Compile or run the daemon utility by defining flag arguments using Cobra bindings:

```bash
# Run the real-time scanning agent
go run main.go --file=mock_access.log --output=alerts.jsonl
```

### 3. Review Help Constraints
To view argument flags and usage parameters directly from your console shell interface, request the built-in documentation menu:

```bash
go run main.go --help
```

```text
A Log Watcher Daemon

Usage:
  log-guardian [flags]

Flags:
  -f, --file string     Path to log .log file
  -h, --help            help for log-guardian
  -o, --output string   Path to output .jsonl file
```