package random

import (
	"regexp"
	"strconv"
	"strings"
)

func RecursiveString(recursive string) string {
	if !strings.Contains(recursive, "[") {
		return recursive
	}

	builder := strings.Builder{}

	start := strings.Index(recursive, "[")
	end := strings.LastIndex(recursive, "]")

	reg := regexp.MustCompile("[0-9]+")
	strs := reg.FindAllString(recursive[:start], -1)
	nums := []int{}

	for _, v := range strs {
		num, err := strconv.Atoi(v)

		if err == nil {
			nums = append(nums, num)
		}
	}

	repeats := nums[len(nums)-1]

	base := strings.TrimSuffix(recursive[:start], strs[len(strs)-1])

	builder.WriteString(base)
	substr := recursive[start+1 : end]

	after := RecursiveString(substr)

	for i := 0; i < repeats; i++ {
		builder.WriteString(after)
	}

	return builder.String()
}
