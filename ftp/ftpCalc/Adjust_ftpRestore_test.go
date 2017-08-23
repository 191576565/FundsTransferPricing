package ftpCalc

import (
	"fmt"
	"testing"
)

func Test_getFtpRestoreCurve(t *testing.T) {
	rst := getFtpRestoreCurve("FTP")
	fmt.Println("resutl is ", rst)
}
