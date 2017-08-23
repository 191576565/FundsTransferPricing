package dbobj

import "ftpProject/logs"

func init() {

	if DefaultDB() == "db2" {
		logs.Debug("init sql")
		logs.Info("dbobj包初始化DB2数据库的sql")
		LOG_TO_DB = "insert into sys_op_logs(op_user_id,op_org,op_type,op_content,op_ip,op_date,op_role,op_app) values(?,?,?,?,?,?,?,?)"
	}

}
