package ftpCalc

import (
	"dbobj"
	"fmt"
)

type ftpRestore struct {
	curve_id string
}

func (this *ftpRestore) get(domain_id string) (map[string]string, error) {
	sql := `select distinct busiz_id, curve_id from ftp_adjust_ftp_restore t where domain_id = :1 order by busiz_id`
	rows, err := dbobj.Default.Query(sql, domain_id)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return make(map[string]string), err
	}
	rst := make(map[string]string)
	for rows.Next() {
		bid := ""
		cid := ""
		err = rows.Scan(&bid, &cid)
		if err != nil {
			fmt.Println(err)
			return make(map[string]string), err
		}
		rst[bid] = cid
	}
	return rst, nil
}

func getFtpRestoreCurve(domain_id string) map[string]string {
	r := new(ftpRestore)
	rst, err := r.get(domain_id)
	if err != nil {
		fmt.Println("init ftp restore curve failed.")
		return make(map[string]string)
	}
	return rst
}
