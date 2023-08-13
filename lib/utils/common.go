package utils

import "fmt"

//org增加前缀
func AddSubjectPrefix(prefix, sub string) string {
	return fmt.Sprintf("%s%s", prefix, sub)
}
