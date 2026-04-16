package main_test

import (
	"strings"
	"testing"

	updateReadme "github/XOzymandiasz/leetcode-solutions/scripts"
)

func TestUpdateReadmeContent(t *testing.T) {
	input := `
		# Repo
		
		| Difficulty |
		|------------|
		| 🟢 |
		| 🟢 |
		| 🟡 |
		| 🔴 |
		
		<!-- STATS_START -->
		old stats
		<!-- STATS_END -->
	`

	output := updateReadme.UpdateReadmeContent(input)

	if !strings.Contains(output, "- Total solved: 4") {
		t.Fatalf("expected total solved = 4, got:\n%s", output)
	}
	if !strings.Contains(output, "- Easy: 2 🟢") {
		t.Fatalf("expected easy = 2, got:\n%s", output)
	}
	if !strings.Contains(output, "- Medium: 1 🟡") {
		t.Fatalf("expected medium = 1, got:\n%s", output)
	}
	if !strings.Contains(output, "- Hard: 1 🔴") {
		t.Fatalf("expected hard = 1, got:\n%s", output)
	}
}
