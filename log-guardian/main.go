package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hpcloud/tail"
	"github.com/spf13/cobra"
)

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
)

var (
	filePath   string
	outputPath string
)

type LogEntry struct {
	IP         string `json:"ip"`
	Timestamp  string `json:"timestamp"`
	AttackType string `json:"attack_type"`
}

func addLineToAlerts(line string, output string) {
	ip := strings.Split(line, " ")[0]
	var date string

	openBracket := strings.Index(line, "[")

	if openBracket != -1 {
		closeBracket := strings.Index(line[openBracket:], "]")
		if closeBracket != -1 {
			closeBracket += openBracket
			date = line[openBracket+1 : closeBracket]
		}
	}

	if date == "" {
		date = "UNKNOWN_TIMESTAMP"
	}

	var attackType string

	if strings.Contains(line, "/etc/passwd") {
		attackType = "Path Traversal"
	} else if strings.Contains(line, "UNION SELECT") {
		attackType = "SQL Injection"
	} else if strings.Contains(line, "<script>") {
		attackType = "Cross-Site Scripting"
	} else if strings.Contains(line, "cgi-bin") {
		attackType = "Remote Command Execution"
	}

	entry := LogEntry{IP: ip, Timestamp: date, AttackType: attackType}

	jsonData, _ := json.Marshal(entry)

	file, _ := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()

	file.WriteString(string(jsonData) + "\n")
}

func isLineSuspicious(line string) bool {
	suspiciousRequests := []string{"/etc/passwd", "UNION SELECT", "<script>", "cgi-bin"}
	for _, request := range suspiciousRequests {
		if strings.Contains(line, request) {
			return true
		}
	}
	return false
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else {
		return false
	}
}

func startScan(cmd *cobra.Command, args []string) {
	if fileExists(filePath) {
		config := tail.Config{Follow: true, ReOpen: true, MustExist: true, Location: &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END}, Logger: log.New(io.Discard, "", 0)}

		t, err := tail.TailFile(filePath, config)
		if err != nil {
			fmt.Printf("Error openning file: %v\n", filePath)
		}

		fmt.Println("Scan started.")

		for line := range t.Lines {
			if isLineSuspicious(line.Text) {
				fmt.Printf("%s[WARNING] found a suspricious line: %s%s\n", ColorRed, line.Text, ColorReset)
				addLineToAlerts(line.Text, outputPath)
			}
		}
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "log-guardian",
		Short: "A Log Watcher Daemon",
		Run:   startScan,
	}

	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to log .log file")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Path to output .jsonl file")

	rootCmd.MarkFlagRequired("file")
	rootCmd.MarkFlagRequired("output")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
