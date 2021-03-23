package iteration

import "strings"

func RepeatBuildIn(ch string, times int) string {
	return strings.Repeat(ch, times)
}

func Repeat(ch string, times int) string {
	var b strings.Builder
	b.Grow(times)
	for b.Len() < times {
		if b.Len() <= times/2 {
			b.WriteString(ch)
		} else {
			b.WriteString(b.String()[:times-b.Len()])
			break
		}
	}
	return b.String()
}
