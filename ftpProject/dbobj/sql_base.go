package dbobj

var (
	LOG_TO_DB = "insert into sys_op_logs(op_user_id,op_org,op_type,op_content,op_ip,op_date,op_role,op_app) values(:1,:2,:3,:4,:5,:6,:7,:8)"
)
