package utils

func TrimStringEnd(str string, num int) string {
	var ts string
	for i := 0; i < len(str)-num; i++ {
		ts = ts + string(str[i])
	}
	return ts
}
