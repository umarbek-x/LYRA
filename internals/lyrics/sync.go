package lyrics

import "time"

func CurrentLine(lines []Line, pos time.Duration) int {
	current := 0
	for i, line := range lines {
		if pos >= line.Time {
			current = i
		} else {
			break
		}
	}
	return current
}
