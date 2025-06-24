package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	AIModel = "deepseek-coder-v2"
	Pass    = "SummarizeModeActivatedPleaseGila"
)

func readFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(b)
}

func parseLog() ([]string, [][]string) {
	history := []string{}
	lastLines := []string{}
	mode := ""

	file, err := os.Open("config/log.txt")
	if err != nil {
		return history, [][]string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lower := strings.ToLower(line)
		if lower == "history:" {
			mode = "history"
			continue
		} else if lower == "last:" {
			mode = "last"
			continue
		}
		if mode == "history" {
			history = append(history, line)
		} else if mode == "last" {
			lastLines = append(lastLines, line)
		}
	}

	group := [][]string{}
	block := []string{}
	for _, ln := range lastLines {
		if strings.HasPrefix(ln, "- User:") || strings.HasPrefix(ln, "- AI:") {
			if len(block) > 0 {
				group = append(group, block)
			}
			block = []string{ln}
		} else {
			if len(block) > 0 {
				block = append(block, ln)
			} else if len(group) > 0 {
				group[len(group)-1] = append(group[len(group)-1], ln)
			}
		}
	}
	if len(block) > 0 {
		group = append(group, block)
	}
	return history, group
}

func sendPrompt(prompt string) string {
	body := map[string]string{
		"model":  AIModel,
		"prompt": prompt,
	}
	bs, _ := json.Marshal(body)

	resp, err := http.Post("http://127.0.0.1:11434/api/generate", "application/json", bytes.NewReader(bs))
	if err != nil {
		return "Request Failed"
	}
	defer resp.Body.Close()

	var result string
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		var msg map[string]interface{}
		json.Unmarshal(line, &msg)
		if res, ok := msg["response"].(string); ok {
			result += res
		}
	}
	return strings.TrimSpace(result)
}

func Summarize() {
	history, last := parseLog()
	var prompt strings.Builder
	prompt.WriteString("Summarize logs, Respond only with summary, Remove fluff, Intent focused, 3 Lines max, Keep important details:\n\n")
	if len(history) > 0 {
		prompt.WriteString("History:\n" + strings.Join(history, "\n") + "\n\n")
	}
	for _, blk := range last {
		for _, line := range blk {
			prompt.WriteString(line + "\n")
		}
	}
	result := sendPrompt(prompt.String())
	os.WriteFile("config/sum.txt", []byte(result), 0644)
	fmt.Println("Summarized to config/sum.txt")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input")
		os.Exit(1)
	}

	if os.Args[1] == Pass {
		Summarize()
		return
	}

	Msg := strings.Join(os.Args[1:], " ")

	history, last := parseLog()
	personality := readFile("config/personality.txt")
	userInfo := readFile("config/userInfo.txt")
	files := strings.TrimSpace(readFile("config/files.txt"))

	var promptBuilder strings.Builder
	if personality != "" {
		promptBuilder.WriteString("Personality:\n" + personality + "\n\n")
	}
	if userInfo != "" {
		promptBuilder.WriteString("UserInfo:\n" + userInfo + "\n\n")
	}
	if files != "" {
		promptBuilder.WriteString("Files.txt content:\n" + files + "\n\n")
	}
	if len(history) > 0 {
		promptBuilder.WriteString("History:\n" + strings.Join(history, "\n") + "\n\n")
	}
	if len(last) > 0 {
		flat := []string{}
		for _, blk := range last {
			flat = append(flat, blk...)
		}
		promptBuilder.WriteString("Last:\n" + strings.Join(flat, "\n") + "\n\n")
	}
	promptBuilder.WriteString("User: " + Msg + "\nAI:")

	fullPrompt := promptBuilder.String()
	os.WriteFile("config/debug.txt", []byte("Prompt -> "+fullPrompt+"\n"), 0644)

	response := sendPrompt(fullPrompt)
	fmt.Println("Response ->", response)

	f, _ := os.OpenFile("config/debug.txt", os.O_APPEND|os.O_WRONLY, 0644)
	f.WriteString("Response -> " + response + "\n\n")
	f.Close()

	for _, blk := range last {
		history = append(history, blk...)
	}
	var logBuilder strings.Builder
	logBuilder.WriteString("history:\n")
	if len(history) > 0 {
		logBuilder.WriteString(strings.Join(history, "\n") + "\n\n")
	} else {
		logBuilder.WriteString("\n")
	}
	logBuilder.WriteString("last:\n")
	logBuilder.WriteString("- User: " + Msg + "\n")
	respLines := strings.Split(response, "\n")
	if len(respLines) > 0 {
		logBuilder.WriteString("- AI: " + respLines[0] + "\n")
		for _, ln := range respLines[1:] {
			logBuilder.WriteString(ln + "\n")
		}
	} else {
		logBuilder.WriteString("- AI:\n")
	}
	os.WriteFile("config/log.txt", []byte(logBuilder.String()), 0644)
}
