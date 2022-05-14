package yasuna

import "strconv"

func fmtIds(ids []int64) string {
	var s string
	for _, id := range ids {
		s += strconv.FormatInt(id, 10) + ","
	}
	return s[:len(s)-1]
}

func fmtScope(scope ...Scope) Scope {
	var s string
	for _, sc := range scope {
		s += string(sc) + " "
	}
	return Scope(s[:len(s)-1])
}
