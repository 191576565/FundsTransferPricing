package ftpCalc

import (
	"fmt"
	"testing"
)

func Test_initPayment(t *testing.T) {
	r := new(customPay)
	fmt.Println(r.initPayment("2016-07-15"))
	var abc acctInfo
	abc.account_number = "542354325442354"
	abc.as_of_date = "2016-07-15"
	abc.origination_date = "2016-01-01"
	fmt.Println(r.get(abc))
}
