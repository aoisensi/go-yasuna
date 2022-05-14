package yasuna

import "strconv"

func convIds(ids []int64) string {
	var s string
	for _, id := range ids {
		s += strconv.FormatInt(id, 10) + ","
	}
	return s[:len(s)-1]
}
