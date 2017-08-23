package utils

import "strings"

func HandleSqlKey(keyword string) string {
	ss := strings.Replace(keyword, `\`, `\\`, -1)
	ss = strings.Replace(ss, `%`, `\%`, -1)
	ss = strings.Replace(ss, `_`, `\_`, -1)
	ss = strings.Replace(ss, `'`, `''`, -1)
	return ss
}
