package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	Easy   = "🟢"
	Medium = "🟡"
	Hard   = "🔴"
)

const RootReadme = "../README.md"

func main() {
	data, err := os.ReadFile(RootReadme)
	if err != nil {
		panic(err)
	}

	updated := UpdateReadmeContent(string(data))

	if bytes.Equal(data, []byte(updated)) {
		fmt.Println("README is already up to date")
		return
	}

	if err := os.WriteFile(RootReadme, []byte(updated), 0644); err != nil {
		panic(err)
	}

	fmt.Println("README stats updated")
}

func countStats(content string) (easy, medium, hard, total int) {
	easy = strings.Count(content, Easy)
	medium = strings.Count(content, Medium)
	hard = strings.Count(content, Hard)
	total = easy + medium + hard

	return
}

func UpdateReadmeContent(content string) string {
	fmt.Println("Counting...")
	easy, medium, hard, total := countStats(content)

	stats := fmt.Sprintf(
		"- Total solved: %d\n- Easy: %d 🟢\n- Medium: %d 🟡\n- Hard: %d 🔴",
		total, easy, medium, hard,
	)

	re := regexp.MustCompile(`(?s)<!-- STATS_START -->.*<!-- STATS_END -->`)
	replacement := fmt.Sprintf("<!-- STATS_START -->\n%s\n<!-- STATS_END -->", stats)

	fmt.Println("Writing...")
	return re.ReplaceAllString(content, replacement)
}
