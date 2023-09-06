package unique

import (
	"slices"
)

func Unique(lines []string) []string {
	return slices.Compact(lines)
}
