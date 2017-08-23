package ftpCalc

import (
	"fmt"
	"testing"
)

func Test_getReserveInfo(t *testing.T) {
	m := getReserveInfo("FTP")
	adjust := calcReserveRate("3.45", m["2000020401"])
	fmt.Println("准备金调节项值:", adjust)
}
