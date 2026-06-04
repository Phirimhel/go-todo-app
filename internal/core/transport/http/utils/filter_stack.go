package core_http_utils

import "strings"

func FilterStack(stack []byte) string {
	lines := strings.Split(string(stack), "\n")
	var result []string

	for _, line := range lines {
		// keep only lines from your project
		if strings.Contains(line, "go-todo-app") {
			result = append(result, strings.TrimSpace(line))
		}
	}

	return strings.Join(result, "\n")
}
