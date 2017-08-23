package ftp

var (
	FTP_VALUECALC_GET1 = `select  F.BUSIZ_ID,F.BUSIZ_DESC ,F.AL_FLAG,FSM.FTP_METHOD_ID,FSM.FTP_METHOD_DESC   from FTP_BUSIZ_INFO F
left join ftp_busiz_method_relation FBMR  ON  F.BUSIZ_ID =FBMR.BUSIZ_ID  AND F.DOMAIN_ID=FBMR.DOMAIN_ID
left join FTP_SYS_METHOD             FSM  ON  FBMR.FTP_METHOD_ID=FSM.FTP_METHOD_ID
 where F.ftp_flag=0 AND F.domain_id  =:1 order by busiz_id` //liutie 1  add  :1
	//利率调整方式
	FTP_VALUECALC_GET2 = `select adjustable_type_cd,adjustable_type_desc from ftp_rate_adjust_attr`
	//利息计提
	FTP_VALUECALC_GET3 = `select accrual_basis_cd,accrual_basis_desc from ftp_accrual_cd_attr`
	//偿还方式
	FTP_VALUECALC_GET4 = `select amrt_type_cd,amrt_type_desc from ftp_payment_type_attr`

	FTP_CURVESAVE_POST1 = `insert into MAS_CURVE_DEFINE (curve_id,curve_desc,iso_currency_cd,create_date,domain_id,curve_type,rep_id)  values(:1,:2,:3,TO_DATE(:4,'YYYY-MM-DD'),:5,:6,:7)`
	FTP_CURVESAVE_POST2 = `insert into mas_curve_info_struct_node(curve_id,struct_code,domain_id,uuid) values(:1,:2,:3,:4)`
	FTP_CURVESAVE_PUT   = `update  MAS_CURVE_DEFINE set curve_desc=:1,curve_type=:2,iso_currency_cd=:3,rep_id=:4 where curve_id=:5 and domain_id=:6`
	FTP_CURVESAVE_PUT1  = `delete from mas_curve_info_struct_node where curve_id = :1 and domain_id=:2`
	FTP_CURVESAVE_PUT2  = `insert into mas_curve_info_struct_node(curve_id,struct_code,domain_id,uuid) values(:1,:2,:3,:4)`
	FTP_CURVESAVE_GET   = `
       select struct_code,
    term_cd,
    term_cd_mult,
    domain_id,
    sort_id
    from mas_curve_struct order by  sort_id
  `
	FTP_CURVESAVE_DELETE1 = `delete from mas_curve_define where curve_id = :1 and domain_id=:2`
	FTP_CURVESAVE_DELETE2 = `delete from mas_curve_info_struct_node where curve_id = :1 and domain_id=:2`
	FTP_CURVESAVE_DELETE3 = `delete from  MAS_CURVE_INFO t where t.curve_uuid like :1||'_'||:2||'_'||'%'`
	FTP_BUSIZSTRUCT_GET   = `select busiz_id,busiz_up_id,trim(busiz_desc) busiz_desc from ftp_busiz_info where domain_id=:1`
	FTP_BUSIZ_GET1        = `select busiz_id from ftp_busiz_info where busiz_up_id='-1' and domain_id=:1`
	FTP_BUSIZ_GET2        = `select  BUSIZ_ID
        ,BUSIZ_ID_DESC
        ,BUSIZ_UP_ID
        ,BUSIZ_TYPE
        ,FTP_FLAG
        ,AL_FLAG
        ,FTP_METHOD_DESC
        ,FTP_METHOD_ID
        ,TERM_CD
        ,TERM_CD_MULT
        ,POINT_VAL
        ,CURVE_ID
        ,DOMAIN_ID
        ,CASE WHEN BUSIZ_UP_ID = '-1' THEN 'true' ELSE 'false' END IS_ROOT
        ,CASE WHEN level <=2 then 'true' ELSE 'false' END
        ,ADJ_ID_COLL,level leve,rownum seq
   from (
              SELECT T0.BUSIZ_ID
                    ,T0.BUSIZ_ID||' '||T0.BUSIZ_DESC      AS  BUSIZ_ID_DESC
                    ,T0.BUSIZ_UP_ID
                    ,T0.BUSIZ_TYPE
                    ,T0.FTP_FLAG
                     ,T0.AL_FLAG
                    ,T5.FTP_METHOD_DESC
                    ,T5.FTP_METHOD_ID
                    ,T2.TERM_CD
                    ,T2.TERM_CD_MULT
                    ,T2.POINT_VAL
                    ,T2.CURVE_ID
                    ,T0.DOMAIN_ID
                    ,T6.ADJ_ID_COLL

                 FROM FTP_BUSIZ_INFO                  T0
                LEFT JOIN  FTP_BUSIZ_INFO             T1 ON T0.BUSIZ_UP_ID =T1.BUSIZ_ID    AND T0.DOMAIN_ID=T1.DOMAIN_ID      --Modified by lt 160908   added "AND T0.DOMAIN_ID=T1.DOMAIN_ID"
                LEFT JOIN  FTP_BUSIZ_METHOD_RELATION  T2 ON T0.BUSIZ_ID    =  T2.BUSIZ_ID  AND T0.DOMAIN_ID=T2.DOMAIN_ID     --Modified by lt 160908  (added "AND T0.DOMAIN_ID=T2..DOMAIN_ID"
                LEFT JOIN  FTP_BUSIZ_TYPE_ATTR        T3 ON T0.BUSIZ_TYPE=T3.BUSIZ_TYPE
                LEFT JOIN  FTP_BUSIZ_FLAG_ATTR        T4 ON T0.FTP_FLAG=T4.FTP_FLAG
                LEFT JOIN  FTP_SYS_METHOD             T5 ON T2.FTP_METHOD_ID=T5.FTP_METHOD_ID
                LEFT JOIN (SELECT t.busiz_id,
                                  LISTAGG(t.adj_id, ',') WITHIN GROUP(ORDER BY t.adj_id) AS ADJ_ID_COLL
                                  ,t.domain_id
                             FROM FTP_ADJUST_REL t
                          GROUP BY busiz_id,domain_id)          T6 ON T0.BUSIZ_ID =T6.busiz_id and t0.domain_id=t6.domain_id
                    where t0.domain_id=:1
         )
  start with  BUSIZ_ID=:2  CONNECT BY prior  busiz_id=  busiz_up_id   order by leve, busiz_id
`
	FTP_BUSIZ_POST1   = `insert into FTP_BUSIZ_INFO(busiz_id,busiz_desc,busiz_up_id,ftp_flag,busiz_type,domain_id,al_flag) values(:1,:2,:3,:4,:5,:6,:7)`
	FTP_BUSIZ_POST2   = `insert into ftp_busiz_method_relation(busiz_id,ftp_method_id,curve_id,term_cd,term_cd_mult,point_val,domain_id) values(:1,:2,:3,:4,:5,:6,:7)`
	FTP_BUSIZ_POST3   = `insert into ftp_adjust_rel(busiz_id,adj_id,adj_type_id,domain_id) values(:1,:2,:3,:4)` //liutie 2  add:4
	FTP_BUSIZ_PUT1    = `update  FTP_BUSIZ_INFO set busiz_desc=:1,busiz_up_id=:2,ftp_flag=:3,busiz_type=:4,al_flag=:5 where busiz_id=:6 AND domain_id=:7`
	FTP_BUSIZ_PUT2    = `delete from ftp_busiz_method_relation where busiz_id=:1 and domain_id=:2`
	FTP_BUSIZ_PUT3    = `insert into ftp_busiz_method_relation(busiz_id,ftp_method_id,curve_id,term_cd,term_cd_mult,point_val,domain_id) values(:1,:2,:3,:4,:5,:6,:7)`
	FTP_BUSIZ_PUT4    = `delete from ftp_adjust_rel where busiz_id=:1 and domain_id=:2`                                                                                                                                                                         //liutie  3 add :2
	FTP_BUSIZ_PUT5    = `insert into ftp_adjust_rel(busiz_id,adj_id,adj_type_id,domain_id) values(:1,:2,:3,:4)`                                                                                                                                                 //liutie  4 add :4
	FTP_BUSIZ_DELETE1 = `delete from FTP_BUSIZ_METHOD_RELATION t
      where (t.busiz_id,t.domain_id) in
      (select t.busiz_id ,t.domain_id from FTP_BUSIZ_INFO t where t.domain_id=:1 start with busiz_id=:2
      connect by prior busiz_id= busiz_up_id)` //liutie 5 add where :2

	FTP_BUSIZ_DELETE2 = `delete from ftp_adjust_rel t
      where (t.busiz_id,t.domain_id) in
      (select t.busiz_id,t.domain_id from FTP_BUSIZ_INFO t where t.domain_id=:1 start with busiz_id=:2
      connect by prior busiz_id= busiz_up_id)` //liutie 6 add where :2

	FTP_BUSIZ_DELETE3 = `DELETE  FROM FTP_BUSIZ_INFO  T1
       WHERE (T1.BUSIZ_ID,t1.domain_id) IN (
        SELECT T.BUSIZ_ID,t.domain_id  from FTP_BUSIZ_INFO t
       where t.domain_id=:1   start with busiz_id=:2 connect by prior busiz_id= busiz_up_id    )` //liutie 7 add where :3
	FTP_CURVEDATA_GET = `
  select struct_code
  from mas_curve_info_struct_node
  where curve_id=:1 and domain_id=:2
  `
	FTP_CURVEDATA_POST1 = `insert into MAS_CURVE_info(
             curve_uuid,
          as_of_date,
          yield)
                   values(:1,TO_DATE(:2,'YYYY-MM-DD'),:3)`

	FTP_CURVEDATA_POST2 = `insert into MAS_CURVE_info(
             curve_uuid,
          as_of_date,
          yield)
                   values(:1,TO_DATE(:2,'YYYY-MM-DD'),:3)`
	FTP_CURVEDATA_DELETE = `delete from  MAS_CURVE_INFO t where t.curve_uuid like :1||'_'||:2||'_'||'%' AND T.AS_OF_DATE = TO_DATE(:3,'YYYY-MM-DD') `
	FTP_CURVEDATA_PUT1   = `delete from  MAS_CURVE_INFO t where t.curve_uuid like :1||'_'||:2||'_'||'%' AND T.AS_OF_DATE = TO_DATE(:3,'YYYY-MM-DD')`
	FTP_CURVEDATA_PUT2   = `insert into MAS_CURVE_info(
             curve_uuid,
          as_of_date,
          yield)
                   values(:1,TO_DATE(:2,'YYYY-MM-DD'),:3)`
	FTP_CURVEDEF_GET1 = `select
        curve_id,
        curve_desc,
        curve_type,
        curve_type_desc,
        iso_currency_cd,
		iso_currency_desc,
        create_date,
        domain_id,
      
        rep_id,
        rep_desc,
      cnt
    from (
      select t.curve_id,
           t.curve_desc,
     k.curve_type,
           k.curve_type_desc,
           t.iso_currency_cd,
		 	m.iso_currency_desc,
           t.create_date,
           t.domain_id,         
           n.rep_id,
           n.rep_desc,
           count(*) over() as cnt,
         row_number() over(order by t.curve_id) as rk
      from MAS_CURVE_DEFINE t
      inner join mas_dim_currency m
      on t.iso_currency_cd = m.iso_currency_cd
      inner join mas_curve_type k
      on t.curve_type=k.curve_type
      left join FTP_REPRICE_FREQ n
      on t.rep_id=n.rep_id
      where t.domain_id  = :1
    order by t.curve_id asc)`

	FTP_CURVEDEF_GET2 = `select
        curve_id,
      curve_desc,
       curve_type,
        curve_type_desc,
        iso_currency_cd,
		iso_currency_desc,
		
        create_date,
        domain_id,
         rep_id,
           rep_desc,
      cnt
    from (
      select t.curve_id,
           t.curve_desc,
     k.curve_type,
           k.curve_type_desc,
           t.iso_currency_cd,
		m.iso_currency_desc,       
           t.create_date,
           t.domain_id,
            n.rep_id,
           n.rep_desc,
           count(*) over() as cnt,
         row_number() over(order by t.curve_id) as rk
      from MAS_CURVE_DEFINE t
	 inner join mas_dim_currency m
      on t.iso_currency_cd = m.iso_currency_cd
    inner join mas_curve_type k
      on t.curve_type=k.curve_type
      left join FTP_REPRICE_FREQ n
      on t.rep_id=n.rep_id
    where (upper(t.curve_id) like upper(:1) escape '\' or upper(t.curve_desc) like upper(:2) escape '\')
       and t.domain_id  =:3
    order by t.curve_id asc)
`
	FTP_CURVEDEF_GET3 = `select a.struct_code struct_code from mas_curve_info_struct_node a left join mas_curve_struct b on a.struct_code=b.struct_code
        where a.curve_id=:1 and a.domain_id=:2 order by b.sort_id` //liutie 8 add :2
	FTP_CURVEINFOPAGE_GET = `
	select T.STRUCT_CODE
      from mas_curve_info_struct_node T
     INNER JOIN MAS_CURVE_STRUCT S
        ON T.STRUCT_CODE = S.STRUCT_CODE
       where t.domain_id = :1  and t.curve_id = :2
     ORDER BY S.SORT_ID ASC
  `
	FTP_CURVEINFO_GET = ` select
           domain_id,
           curve_id,
           curve_desc,
           to_char(as_of_date,'YYYY-MM-DD'),
           struct_code,
             to_char(yield,'fm990.0000'),
           cnt
       from (
       select
           domain_id,
           curve_id,
           curve_desc,
           as_of_date,
           struct_code,
           yield,
           rk,
           max(rk) over() as cnt
       from (
         select
               t.domain_id,
               t.curve_id,
               t.curve_desc,
               i.as_of_date,
               n.struct_code,
               i.yield,
               dense_rank() over(order by t.curve_id,i.as_of_date desc) as rk
          from mas_curve_define t
          inner join mas_curve_info_struct_node n
            on t.curve_id = n.curve_id
           and t.domain_id = n.domain_id
          inner join mas_curve_struct s
          on n.struct_code = s.struct_code
          inner join mas_curve_info i
          on n.uuid = i.curve_uuid
          where i.as_of_date >= to_date(:1,'YYYY-MM-DD') and i.as_of_date <= to_date(:2,'YYYY-MM-DD')
          and t.curve_id = :3 and t.domain_id =:4
      and i.yield is not null
          order by t.curve_id,i.as_of_date desc,s.sort_id asc
        )
        ) where rk > :5 and rk <= :6
`
	FTP_REDEMPTION_GET = `select term_cd,term_cd_mult,replace(to_char(weight),'.','0.')
 from FTP_BUSIZ_REDEMPTION_CURVE where busiz_id=:1 and domain_id=:2 `
	FTP_REDEMPTION_PUT1    = `delete from FTP_BUSIZ_REDEMPTION_CURVE where busiz_id=:1 and domain_id=:2`
	FTP_REDEMPTION_PUT2    = `insert into FTP_BUSIZ_REDEMPTION_CURVE(busiz_id,term_cd,term_cd_mult,weight,domain_id) values(:1,:2,:3,:4,:5)`
	FTP_DISPATCHCALC_POST1 = `delete from ftp_dispatch_pro where dispatch_status='4'`
	FTP_DISPATCHCALC_POST2 = `delete from ftp_dispatch_pro where dispatch_id=:1 and dispatch_date=TO_DATE(:2,'YYYY-MM-DD') and domain_id=:3`
	FTP_DISPATCHCALC_POST3 = `insert into ftp_dispatch_pro(dispatch_id,dispatch_date,dispatch_status,cur_rows,all_rows,err_msg,dispatch_name,domain_id) values(:1,TO_DATE(:2,'YYYY-MM-DD'),:3,:4,:5,:6,:7,:8)`
	FTP_DISPATCHINFO_POST1 = `insert into ftp_dispatch_list(dispatch_id,dispatch_name,input_source_cd,output_result_cd,domain_id,start_offset,max_limit)
    values(:1,:2,:3,:4,:5,:6,:7)`
	FTP_DISPATCHINFO_GET1 = `select Dispatch_Id,Dispatch_Name,Input_Source_Cd,Output_Result_Cd,Domain_Id,Start_Offset,Max_Limit from ftp_dispatch_list where domain_id=:1`
	FTP_DISPATCHINFO_PUT1 = `update  ftp_dispatch_list set dispatch_name=:1,input_source_cd=:2,
    output_result_cd=:3,domain_id=:4,start_offset=:5,max_limit=:6 where dispatch_id=:7`
	FTP_DISPATCHINFO_DELETE1 = `select dispatch_status from ftp_dispatch_pro where domain_id=:1 and dispatch_id=:2`
	FTP_DISPATCHINFO_DELETE2 = `delete from ftp_dispatch_pro where dispatch_id=:1 and domain_id=:2`
	FTP_DISPATCHINFO_DELETE3 = `delete from ftp_dispatch_list where dispatch_id = :1 and domain_id  = :2`
	//	FTP_DISPATCHREALT_GET    = `select t.dispatch_id,t.dispatch_name,to_char(t.dispatch_date,'YYYY-MM-DD'),
	//t.dispatch_status,t.cur_rows,t.all_rows,t.err_msg,t1.output_result_cd,t.domain_id from ftp_dispatch_pro t
	//join FTP_DISPATCH_list  t1 on t.dispatch_id=t1.dispatch_id where t.domain_id=:1 and t.dispatch_status='1' order by t.dispatch_status
	//`
	FTP_DISPATCHREALT_GET = `select t.dispatch_id,t.dispatch_name,to_char(t.dispatch_date,'YYYY-MM-DD'),
t.dispatch_status,t.cur_rows,t.all_rows,t.err_msg,T1.INPUT_SOURCE_CD,t1.output_result_cd,t.domain_id,t1.start_offset from ftp_dispatch_pro t
join FTP_DISPATCH_list  t1 on t.dispatch_id=t1.dispatch_id where t.domain_id=:1 order by t.dispatch_status`
	FTP_DISPATCHREALT_PUT1 = `update  ftp_dispatch_pro set dispatch_status=:1,err_msg=:2
   where dispatch_id=:3 and dispatch_date=TO_DATE(:4,'YYYY-MM-DD') and domain_id=:5
    `
	FTP_DISPATCHREALT_DELETE1 = `delete from ftp_dispatch_pro where dispatch_id=:1 and dispatch_date=TO_DATE(:2,'YYYY-MM-DD') and domain_id=:3`

	FTP_ADJPOLICY_GET1 = `select
  uuid,
  adj_id,
adj_desc,
        org_unit_id,
      
       iso_currency_cd,
       adj_dyn_dim,
       DYN_NAME,
       term_str,
       term_end,
       to_char(last_date,'YYYY-MM-DD'),
       adj_bp,
       to_char(eff_str_date,'YYYY-MM-DD'),
       to_char(eff_end_date,'YYYY-MM-DD'),
       to_char(buz_str_date,'YYYY-MM-DD'),
       to_char(buz_end_date,'YYYY-MM-DD'),
       domain_id,
       cnt
  from (select t.uuid,
         t.adj_id,
               a.adj_desc,
               t.org_unit_id,
             
               t.iso_currency_cd,
               t.adj_dyn_dim,
               T.DYN_NAME,
               t.term_str,
               t.term_end,
               t.last_date,
               t.adj_bp,
               t.eff_str_date,
               t.eff_end_date,
               t.buz_str_date,
               t.buz_end_date,
               t.domain_id,
               count(t.uuid) over() as cnt,
               row_number() over(order by t.adj_id) as rk
          from (
                select  T.*,T1.PRODUCT_NAME DYN_NAME
                  from FTP_ADJUST_POLICY t
                LEFT JOIN FTP_PRODUCT_INFO  T1 ON  T.ADJ_DYN_DIM=T1.PRODUCT_ID   AND T.DOMAIN_ID=T1.DOMAIN_ID
                WHERE T.ADJ_ID='801'
                UNION ALL
                select  T.*,T1.al_flag_desc  DYN_NAME
                  from FTP_ADJUST_POLICY t
                LEFT JOIN FTP_AL_TYPE  T1 ON  T.ADJ_DYN_DIM=T1.AL_flag
                WHERE T.ADJ_ID='802'
                UNION ALL
                select  T.*,T1.BUSIZ_DESC  DYN_NAME
                from FTP_ADJUST_POLICY t
                LEFT JOIN  FTP_BUSIZ_INFO  T1 ON  T.ADJ_DYN_DIM=T1.BUSIZ_ID  AND T.DOMAIN_ID=T1.DOMAIN_ID
                WHERE T.ADJ_ID IN ('803' ,'805')
                UNION ALL
                select  T.*,T1.INDUSTRY_NAME  DYN_NAME
                from FTP_ADJUST_POLICY t
                LEFT JOIN  FTP_INDUSTRY_INFO  T1 ON  T.ADJ_DYN_DIM=T1.INDUSTRY_ID AND T.DOMAIN_ID=T1.DOMAIN_ID
                WHERE T.ADJ_ID IN ('804')
                UNION ALL
                select  T.*,T.ADJ_DYN_DIM  DYN_NAME
                from FTP_ADJUST_POLICY t
                WHERE T.ADJ_ID='806'
            ) t
          LEFT  join ftp_adj_info a on t.adj_id=a.adj_id  and a.domain_id=t.domain_id
          
      where t.domain_id=:1 and SSQQLL
      AND T.ORG_UNIT_ID IN ORGINPARA
       ORDER BY    domain_id,ADJ_ID ,ORG_UNIT_ID ,ISO_CURRENCY_CD ,ADJ_DYN_DIM,LAST_DATE,TERM_STR,TERM_END,EFF_STR_DATE,EFF_END_DATE,BUZ_STR_DATE,BUZ_END_DATE
          )
          where rk>:2 and rk<=:3
         `
	FTP_ADJPOLICY_POST1 = `insert into FTP_ADJUST_POLICY(
		adj_id
		,org_unit_id
		,iso_currency_cd
		,adj_dyn_dim
		,term_str
		,term_end
		,last_date
		,adj_bp
		,eff_str_date
		,eff_end_date
		,buz_str_date
		,buz_end_date
		,domain_id
		,uuid)
		values(:1,:2,:3,:4,:5,:6,to_date(:7,'YYYY-MM-DD'),:8,to_date(:9,'YYYY-MM-DD'),to_date(:10,'YYYY-MM-DD'),to_date(:11,'YYYY-MM-DD'),to_date(:12,'YYYY-MM-DD'),:13,SEQ_ADJ_CHK.NEXTVAL)`
	FTP_ADJPOLICY_PUT1 = `	update  FTP_ADJUST_POLICY set
  org_unit_id    =:1,
iso_currency_cd=:2,
adj_dyn_dim    =:3,
term_str       =:4,
term_end       =:5,
last_date      =to_date(:6,'YYYY-MM-DD'),
adj_bp         =:7,
eff_str_date   =to_date(:8,'YYYY-MM-DD'),
eff_end_date   =to_date(:9,'YYYY-MM-DD'),
buz_str_date   =to_date(:10,'YYYY-MM-DD'),
buz_end_date   =to_date(:11,'YYYY-MM-DD')
  where uuid=:12`
	FTP_ADJPOLICY_DELETE1   = `delete from FTP_ADJUST_POLICY where uuid=:1 and domain_id=:2`
	FTP_ADJPOLICYDOWN_POST1 = `insert into FTP_ADJUST_POLICY(
adj_id
,org_unit_id
,iso_currency_cd
,adj_dyn_dim
,term_str
,term_end
,last_date
,adj_bp
,eff_str_date
,eff_end_date
,buz_str_date
,buz_end_date
,domain_id
,memo
,uuid)
values(:1,:2,:3,:4,:5,:6,to_date(:7,'YYYY-MM-DD'),:8,to_date(:9,'YYYY-MM-DD'),to_date(:10,'YYYY-MM-DD'),to_date(:11,'YYYY-MM-DD'),to_date(:12,'YYYY-MM-DD'),:13,:14,SEQ_ADJ_CHK.NEXTVAL)`

	FTP_ADJPOLICYDOWN_GET1 = `select
  uuid,
  adj_id,
adj_desc,
        org_unit_id,
       iso_currency_cd,
       adj_dyn_dim,
       DYN_NAME,
       term_str,
       term_end,
       to_char(last_date,'YYYY-MM-DD'),
       adj_bp,
       to_char(eff_str_date,'YYYY-MM-DD'),
       to_char(eff_end_date,'YYYY-MM-DD'),
       to_char(buz_str_date,'YYYY-MM-DD'),
       to_char(buz_end_date,'YYYY-MM-DD'),
       domain_id,
     memo,
       cnt
  from (select t.uuid,
         t.adj_id,
               a.adj_desc,
               t.org_unit_id,
              
               t.iso_currency_cd,
               t.adj_dyn_dim,
               T.DYN_NAME,
               t.term_str,
               t.term_end,
               t.last_date,
               t.adj_bp,
               t.eff_str_date,
               t.eff_end_date,
               t.buz_str_date,
               t.buz_end_date,
               t.domain_id,             
         t.memo,
               count(t.uuid) over() as cnt
          from (
                select  T.*,T1.PRODUCT_NAME DYN_NAME
                  from FTP_ADJUST_POLICY t
                LEFT JOIN FTP_PRODUCT_INFO  T1 ON  T.ADJ_DYN_DIM=T1.PRODUCT_ID   AND T.DOMAIN_ID=T1.DOMAIN_ID
                WHERE T.ADJ_ID='801'
                UNION ALL
                select  T.*,T1.al_flag_desc  DYN_NAME
                  from FTP_ADJUST_POLICY t
                LEFT JOIN FTP_AL_TYPE  T1 ON  T.ADJ_DYN_DIM=T1.AL_flag
                WHERE T.ADJ_ID='802'
                UNION ALL
                select  T.*,T1.BUSIZ_DESC  DYN_NAME
                from FTP_ADJUST_POLICY t
                LEFT JOIN  FTP_BUSIZ_INFO  T1 ON  T.ADJ_DYN_DIM=T1.BUSIZ_ID  AND T.DOMAIN_ID=T1.DOMAIN_ID
                WHERE T.ADJ_ID IN ('803' ,'805')
                UNION ALL
                select  T.*,T1.INDUSTRY_NAME  DYN_NAME
                from FTP_ADJUST_POLICY t
                LEFT JOIN  FTP_INDUSTRY_INFO  T1 ON  T.ADJ_DYN_DIM=T1.INDUSTRY_ID AND T.DOMAIN_ID=T1.DOMAIN_ID
                WHERE T.ADJ_ID IN ('804')
                UNION ALL
                select  T.*,T.ADJ_DYN_DIM  DYN_NAME
                from FTP_ADJUST_POLICY t
                WHERE T.ADJ_ID='806'
            ) t
          LEFT  join ftp_adj_info a on t.adj_id=a.adj_id and a.domain_id=t.domain_id             
      where t.domain_id=:1 and SSQQLL
     AND T.ORG_UNIT_ID IN ORGINPARA
       ORDER BY    domain_id,ADJ_ID ,ORG_UNIT_ID ,ISO_CURRENCY_CD ,ADJ_DYN_DIM,LAST_DATE,TERM_STR,TERM_END,EFF_STR_DATE,EFF_END_DATE,BUZ_STR_DATE,BUZ_END_DATE
          )
         `
	FTP_PCHECK_GET1 = `select
  uuid,
  adj_id,
adj_desc,
        org_unit_id,
       iso_currency_cd,
       adj_dyn_dim,
       DYN_NAME,
       term_str,
       term_end,
       to_char(last_date,'YYYY-MM-DD'),
       adj_bp,
       to_char(eff_str_date,'YYYY-MM-DD'),
       to_char(eff_end_date,'YYYY-MM-DD'),
       to_char(buz_str_date,'YYYY-MM-DD'),
       to_char(buz_end_date,'YYYY-MM-DD'),
       domain_id,
       UUID_DOUB,
       cnt
  from (select t.uuid,
         t.adj_id,
               a.adj_desc,
               t.org_unit_id,
               t.iso_currency_cd,
               t.adj_dyn_dim,
               T.DYN_NAME,
               t.term_str,
               t.term_end,
               t.last_date,
               t.adj_bp,
               t.eff_str_date,
               t.eff_end_date,
               t.buz_str_date,
               t.buz_end_date,
               t.domain_id,           
               t.UUID_DOUB,
               count(t.uuid) over() as cnt,
               row_number() over(order by t.adj_id) as rk
          from (
              SELECT  DISTINCT  T.ADJ_ID, T.ORG_UNIT_ID, T.ISO_CURRENCY_CD, T.ADJ_DYN_DIM, T.TERM_STR, T.TERM_END, T.LAST_DATE, T.ADJ_BP, T.EFF_STR_DATE, T.EFF_END_DATE, T.BUZ_STR_DATE, T.BUZ_END_DATE, T.DOMAIN_ID
                 ,t.UUID
                 ,T2.MEMO  AS "UUID_DOUB"
                 ,T1.PRODUCT_NAME DYN_NAME
                  FROM FTP_ADJUST_POLICY  T
                LEFT JOIN FTP_ADJUST_POLICY_CHK  T2 ON T.UUID=T2.UUID
                LEFT JOIN FTP_PRODUCT_INFO  T1 ON  T.ADJ_DYN_DIM=T1.PRODUCT_ID   AND T.DOMAIN_ID=T1.DOMAIN_ID
                WHERE T.ADJ_ID='801' AND  T2.MEMO IS NOT NULL
                UNION ALL
              SELECT  DISTINCT  T.ADJ_ID, T.ORG_UNIT_ID, T.ISO_CURRENCY_CD, T.ADJ_DYN_DIM, T.TERM_STR, T.TERM_END, T.LAST_DATE, T.ADJ_BP, T.EFF_STR_DATE, T.EFF_END_DATE, T.BUZ_STR_DATE, T.BUZ_END_DATE, T.DOMAIN_ID
                 ,t.UUID
                 ,T2.MEMO  AS "UUID_DOUB"
                 ,T1.al_flag_desc DYN_NAME
                  FROM FTP_ADJUST_POLICY  T
                LEFT JOIN FTP_ADJUST_POLICY_CHK  T2 ON T.UUID=T2.UUID
                LEFT JOIN FTP_AL_TYPE  T1 ON  T.ADJ_DYN_DIM=T1.AL_flag
                WHERE T.ADJ_ID='802'  AND  T2.MEMO IS NOT NULL
                UNION ALL
               SELECT  DISTINCT  T.ADJ_ID, T.ORG_UNIT_ID, T.ISO_CURRENCY_CD, T.ADJ_DYN_DIM, T.TERM_STR, T.TERM_END, T.LAST_DATE, T.ADJ_BP, T.EFF_STR_DATE, T.EFF_END_DATE, T.BUZ_STR_DATE, T.BUZ_END_DATE, T.DOMAIN_ID
                 ,t.UUID
                 ,T2.MEMO  AS "UUID_DOUB"
                 ,T1.BUSIZ_DESC AS  DYN_NAME
                  FROM FTP_ADJUST_POLICY  T
                LEFT JOIN FTP_ADJUST_POLICY_CHK  T2 ON T.UUID=T2.UUID
                LEFT JOIN  FTP_BUSIZ_INFO  T1 ON  T.ADJ_DYN_DIM=T1.BUSIZ_ID  AND T.DOMAIN_ID=T1.DOMAIN_ID
                WHERE T.ADJ_ID IN ('803' ,'805') AND  T2.MEMO IS NOT NULL
                UNION ALL
                SELECT  DISTINCT  T.ADJ_ID, T.ORG_UNIT_ID, T.ISO_CURRENCY_CD, T.ADJ_DYN_DIM, T.TERM_STR, T.TERM_END, T.LAST_DATE, T.ADJ_BP, T.EFF_STR_DATE, T.EFF_END_DATE, T.BUZ_STR_DATE, T.BUZ_END_DATE, T.DOMAIN_ID
                 ,t.UUID
                 ,T2.MEMO  AS "UUID_DOUB"
                 ,T1.INDUSTRY_NAME DYN_NAME
                  FROM FTP_ADJUST_POLICY  T
                LEFT JOIN FTP_ADJUST_POLICY_CHK  T2 ON T.UUID=T2.UUID
                LEFT JOIN  FTP_INDUSTRY_INFO  T1 ON  T.ADJ_DYN_DIM=T1.INDUSTRY_ID AND T.DOMAIN_ID=T1.DOMAIN_ID
                WHERE T.ADJ_ID IN ('804')    AND  T2.MEMO IS NOT NULL
                UNION ALL
               SELECT  DISTINCT  T.ADJ_ID, T.ORG_UNIT_ID, T.ISO_CURRENCY_CD, T.ADJ_DYN_DIM, T.TERM_STR, T.TERM_END, T.LAST_DATE, T.ADJ_BP, T.EFF_STR_DATE, T.EFF_END_DATE, T.BUZ_STR_DATE, T.BUZ_END_DATE, T.DOMAIN_ID
                 ,t.UUID
                 ,T2.MEMO  AS "UUID_DOUB"
                 ,T.ADJ_DYN_DIM DYN_NAME
                  FROM FTP_ADJUST_POLICY  T
                LEFT JOIN FTP_ADJUST_POLICY_CHK  T2 ON T.UUID=T2.UUID
                WHERE T.ADJ_ID='806'     AND  T2.MEMO IS NOT NULL


            ) t
          LEFT  join ftp_adj_info a on t.adj_id=a.adj_id and a.domain_id=t.domain_id
      where t.domain_id=:1 and SSQQLL
      AND T.ORG_UNIT_ID IN ORGINPARA
       ORDER BY    domain_id,ADJ_ID ,ORG_UNIT_ID ,ISO_CURRENCY_CD ,ADJ_DYN_DIM,LAST_DATE,TERM_STR,TERM_END,EFF_STR_DATE,EFF_END_DATE,BUZ_STR_DATE,BUZ_END_DATE
          )
          where rk>:3 and rk<=:4`
	FTP_ADJINFO_GET1 = `select t.adj_id, t.adj_desc, t.adj_type_id, a.adj_type_name,t.status,b.status_desc
        from FTP_ADJ_INFO t
       inner join ftp_adj_type a
        on a.adj_type_id = t.adj_type_id
         inner join FTP_ADJ_STATUS b
        on t.status=b.status where t.domain_id=:1
      order by t.adj_type_id,t.adj_id,t.status`
	FTP_ADJINFO_GET2 = `select t.adj_id, t.adj_desc, t.adj_type_id, a.adj_type_name,t.status,b.status_desc
        from FTP_ADJ_INFO t
       inner join ftp_adj_type a
        on a.adj_type_id = t.adj_type_id
   			inner join FTP_ADJ_STATUS b
    		on t.status=b.status
         where t.domain_id=:1 and (upper(t.adj_id) like upper(:2) escape '\' or upper(t.adj_desc) like upper(:3) escape '\')
			order by t.adj_type_id`
	FTP_ADJINFO_POST1   = `insert into FTP_ADJ_INFO( adj_id,adj_desc,adj_type_id,status,domain_id) values(:1,:2,:3,:4,:5)`
	FTP_ADJINFO_PUT1    = `update FTP_ADJ_INFO set adj_desc=:1, adj_type_id=:2 ,status=:3 where adj_id=:4 and domain_id=:5`
	FTP_ADJINFO_DELETE1 = `delete from FTP_ADJ_INFO where adj_id=:1 and domain_id=:2`
	FTP_ADJLIST_GET1    = `select t.adj_id,
       t.adj_desc,
       t.adj_type_id,
       i.adj_type_name
  		from ftp_adj_info t
 		inner join ftp_adj_type i
   		 on t.adj_type_id = i.adj_type_id
 		where t.adj_type_id =:1 and t.status='0' and t.domain_id=:2
 		order by t.adj_type_id`
	FTP_ADJRESERVE_GET1 = `select
	Uuid,
BUSIZ_ID,
BUSIZ_DESC,
RESERVE_PERCENT,
RESERVE_RATE,
EFF_STR_DATE,
EFF_END_DATE,
DOMAIN_ID,
 cnt
from (
select t.Uuid,
t.BUSIZ_ID,
       T1.BUSIZ_DESC,
       t.RESERVE_PERCENT,
       t.RESERVE_RATE,
       t.EFF_STR_DATE,
       t.EFF_END_DATE,
       t.DOMAIN_ID,
       count (t.busiz_id) over() as cnt,
       row_number() over(order by t.busiz_id) as rk
  from FTP_ADJUST_CAPITAL_RESERVES t
 inner join ftp_busiz_info t1
    on t.busiz_id = t1.busiz_id
   and t.domain_id = t1.domain_id
 where t.domain_id = :1 )
 where rk>:2 and rk<=:3`
	FTP_ADJRESERVE_GET2 = `select
	Uuid,
BUSIZ_ID,
BUSIZ_DESC,
RESERVE_PERCENT,
RESERVE_RATE,
EFF_STR_DATE,
EFF_END_DATE,
DOMAIN_ID,
 cnt
from (
select t.Uuid,
t.BUSIZ_ID,
       T1.BUSIZ_DESC,
       t.RESERVE_PERCENT,
       t.RESERVE_RATE,
       t.EFF_STR_DATE,
       t.EFF_END_DATE,
       t.DOMAIN_ID,
       count (t.busiz_id) over() as cnt,
       row_number() over(order by t.busiz_id) as rk
  from FTP_ADJUST_CAPITAL_RESERVES t
 inner join ftp_busiz_info t1
    on t.busiz_id = t1.busiz_id
   and t.domain_id = t1.domain_id
 where t.domain_id = :1 and (upper(t.busiz_id) like upper(:2) escape
                '\' or upper(t1.busiz_desc) like upper(:3) escape '\'))
 where rk>:4 and rk<=:5
`
	FTP_ADJRESERVE_POST1   = `insert into  FTP_ADJUST_CAPITAL_RESERVES(busiz_id,Reserve_Percent,reserve_rate,eff_str_date,eff_end_date,Domain_Id) values(:1,:2,:3,to_date(:4,'YYYY-MM-DD'),to_date(:5,'YYYY-MM-DD'),:6)`
	FTP_ADJRESERVE_PUT1    = `update FTP_ADJUST_CAPITAL_RESERVES set Reserve_Percent = :1,reserve_rate = :2,eff_str_date=to_date(:3,'YYYY-MM-DD'),eff_end_date=to_date(:4,'YYYY-MM-DD')where  Uuid=:5`
	FTP_ADJRESERVE_DELETE1 = `delete from FTP_ADJUST_CAPITAL_RESERVES  where uuid = :1`
	FTP_ADJTERML_GET1      = `select
		busiz_id,
		busiz_desc,
		curve_id,
		curve_desc,
		domain_id,
		cnt
  	from (select
               r.busiz_id,
               f.busiz_desc,
               t.curve_id,
               t.curve_desc,
               f.domain_id,             
               count (r.busiz_id) over() as cnt,
			 row_number() over(order by r.busiz_id) as rk
          from FTP_ADJUST_REL r
          left join (select
                           t.busiz_id,
                           t.domain_id,
                       LISTAGG(t.curve_id  , ',') WITHIN GROUP(ORDER BY t.domain_id,T.BUSIZ_ID)   AS curve_id,
                       LISTAGG(d.curve_desc, ',') WITHIN GROUP(ORDER BY t.domain_id,T.BUSIZ_ID) AS curve_desc
                      from FTP_ADJUST_TERM_LIQUIDITY t
                     inner join mas_curve_define d
                        on t.curve_id = d.curve_id
                       and t.domain_id = d.domain_id
                     GROUP BY   t.domain_id,T.BUSIZ_ID
                       )  t
            on r.busiz_id = t.busiz_id and r.domain_id=t.domain_id
         inner join ftp_busiz_info f
            on r.busiz_id = f.busiz_id and f.domain_id=r.domain_id        
         where r.adj_id = :1
           and f.domain_id = :2
         order by r.busiz_id)
		  where rk>:3 and rk <=:4
`
	FTP_ADJTERML_GET2 = `select
busiz_id,
busiz_desc,
curve_id,
curve_desc,
domain_id,
cnt
  from (select
               r.busiz_id,
               f.busiz_desc,
               t.curve_id,
               t.curve_desc,
               f.domain_id,             
               count (r.busiz_id) over() as cnt,
			row_number() over(order by r.busiz_id) as rk
          from FTP_ADJUST_REL r
          left join (select
                           t.busiz_id,
                           t.domain_id,
                       LISTAGG(t.curve_id  , ',') WITHIN GROUP(ORDER BY t.domain_id,T.BUSIZ_ID)   AS curve_id,
                       LISTAGG(d.curve_desc, ',') WITHIN GROUP(ORDER BY t.domain_id,T.BUSIZ_ID) AS curve_desc
                      from FTP_ADJUST_TERM_LIQUIDITY t
                     inner join mas_curve_define d
                        on t.curve_id = d.curve_id
                       and t.domain_id = d.domain_id
                     GROUP BY   t.domain_id,T.BUSIZ_ID
                       ) t
            on r.busiz_id = t.busiz_id and r.domain_id=t.domain_id
         inner join ftp_busiz_info f
            on r.busiz_id = f.busiz_id and f.domain_id=r.domain_id        
         where r.adj_id = :1
           and f.domain_id = :2
           and (upper(r.busiz_id) like upper(:3) escape
                '\' or upper(f.busiz_desc) like upper(:4) escape '\')
         order by r.busiz_id)
		  where rk>:5 and rk <=:6
`
	FTP_ADJTERML_POST1 = `insert into FTP_ADJUST_TERM_LIQUIDITY(busiz_id,curve_id,domain_id,Reprice_Freq_Range) values(:1,:2,:3,:4)`
	FTP_ADJTERML_POST2 = `select a.rep_calc_attr  from MAS_CURVE_DEFINE t inner join ftp_reprice_freq  a on t.rep_id=a.rep_id
    where t.curve_id=:1`
	FTP_ADJTERML_PUT1    = `delete from FTP_ADJUST_TERM_LIQUIDITY where busiz_id=:1 and domain_id=:2`
	FTP_ADJTERML_PUT2    = `insert into FTP_ADJUST_TERM_LIQUIDITY(busiz_id,curve_id,domain_id) values(:1,:2,:3)`
	FTP_ADJTERML_DELETE1 = `delete from FTP_ADJUST_TERM_LIQUIDITY  where busiz_id = :1`
	FTP_ADJTLP_GET1      = `select t.curve_id, t.curve_desc
 			 from MAS_CURVE_DEFINE t
 			where curve_type = :1
  			 and domain_id = :2`
	FTP_ECALCAUTO_GET1 = `select
 t.item_id,
 t.item_name_l1,
 t.item_name_l2,
 t.busiz_id,
  to_char(t.org_par_bal,'fm999999999990'),
 to_char(t.ratio_bal,'fm9990.0000'),
  to_char(t.cur_net_rate,'fm9990.0000'),
 to_char(t.accd_int,'fm999999999990'),
 to_char(t.adj_int,'fm9999999990.0000'),
 to_char(t.ftp_rate_b,'fm990.0000'),
 to_char(t.ftp_rate_a,'fm990.0000'),
 to_char(t.ftp_margin_a,'fm990.0000'),
 to_char(t.ftp_int_b,'fm999999999990'),
 to_char(t.ftp_int_a,'fm999999999990'),
 to_char(t.ftp_profit_a,'fm999999999990')
  from FTP_ENSEMBLE_CALC t where busiz_id !=item_id and t.domain_id=:1 AND T.ISO_CURRENCY_CD=:2 ORDER BY SORT_ID`
	FTP_ECALCAUTO_GET2 = `with lv2 as
 (select item_name_l1,
         org_par_bal,
         cur_net_rate,
         accd_int,
         ftp_rate_b,
         ftp_rate_a,
         ftp_margin_a,
         ftp_int_b,
         ftp_int_a,
         ftp_profit_a,
          domain_id,
       ISO_CURRENCY_CD
    from FTP_ENSEMBLE_CALC t
   where regexp_like(t.item_id, '^[[:digit:]].[[:digit:]]$')
   order by sort_id),
lv1 as
 (select item_name_l1,
         org_par_bal,
         cur_net_rate,
         accd_int,
         ftp_rate_b,
         ftp_rate_a,
         ftp_margin_a,
         ftp_int_b,
         ftp_int_a,
         ftp_profit_a,
          domain_id,
       ISO_CURRENCY_CD
    from FTP_ENSEMBLE_CALC t
   where regexp_like(t.item_id, '^[[:digit:]]$')
   order by sort_id)
select item_name_l1,
      to_char( org_par_bal,'fm999999999990'),
      to_char(cur_net_rate,'fm9999990.0000'),
       to_char(accd_int,'fm999999999990'),
       to_char(ftp_rate_b,'fm990.0000'),
       to_char(ftp_rate_a,'fm990.0000'),
       to_char(ftp_margin_a,'fm990.0000'),
       to_char(ftp_int_b,'fm999999999990'),
       to_char(ftp_int_a,'fm999999999990'),
       to_char(ftp_profit_a,'fm999999999990'),
        domain_id,
       ISO_CURRENCY_CD
  from lv2
union all
select item_name_l1,
       to_char( org_par_bal,'fm999999999990'),
      to_char(cur_net_rate,'fm9999990.0000'),
       to_char(accd_int,'fm999999999990'),
       to_char(ftp_rate_b,'fm990.0000'),
       to_char(ftp_rate_a,'fm990.0000'),
       to_char(ftp_margin_a,'fm990.0000'),
       to_char(ftp_int_b,'fm999999999990'),
       to_char(ftp_int_a,'fm999999999990'),
       to_char(ftp_profit_a,'fm999999999990'),
       domain_id,
       ISO_CURRENCY_CD
  from lv1
 where domain_id = :1
   and ISO_CURRENCY_CD = :2
`
	FTP_PRODUCTINFO_GET1 = `select t.product_id,
       t.product_name,
       t.product_parent_id,
       t1.product_name,
       t.creation_time,
       t.creater,
	   t.domain_id,
       t.memo,
       level lvl
  from (select * from ftp_product_info t where t.domain_id=:1)  t
  left join (select * from ftp_product_info t where t.domain_id=:2)   t1
    on t.product_parent_id = t1.product_id
 start with t.product_parent_id = '-1'
connect by prior t.product_id = t.product_parent_id
 order by lvl, product_id desc`
	FTP_PRODUCTINFO_PUT1    = `update FTP_PRODUCT_INFO set product_name=:1,product_parent_id=:2,memo=:3 where product_id=:4 and domain_id=:5`
	FTP_PRODUCTINFO_DELETE1 = `delete from FTP_PRODUCT_INFO  i
      		where i.product_id in (
			select  t.product_id from
 			FTP_PRODUCT_INFO t start with t.product_id=:1
			connect by prior t.product_id=t.product_parent_id) and i.domain_id=:2`

	FTP_PRODUCTINFO_POST1 = `insert into FTP_PRODUCT_INFO( product_id,
     product_name,
     product_parent_id,
     creation_time,
     creater,
	 domain_id,
     memo) values(:1,:2,:3,to_date(:4,'YYYY-MM-DD'),:5,:6,:7)`
	FTP_ADJTPRO_GET1 = `select
		busiz_id,
		busiz_desc,
		curve_id,
		curve_desc,
		domain_id,
		cnt
  	from (select
               r.busiz_id,
               f.busiz_desc,
               t.curve_id,
               t.curve_desc,
               f.domain_id,
               count (r.busiz_id) over() as cnt
          from FTP_ADJUST_REL r
          left join (select
                           t.busiz_id,
                           t.domain_id,
                       LISTAGG(t.curve_id  , ',') WITHIN GROUP(ORDER BY t.domain_id,T.BUSIZ_ID)   AS curve_id,
                       LISTAGG(d.curve_desc, ',') WITHIN GROUP(ORDER BY t.domain_id,T.BUSIZ_ID) AS curve_desc
                      from FTP_ADJUST_FTP_RESTORE t
                     inner join mas_curve_define d
                        on t.curve_id = d.curve_id
                       and t.domain_id = d.domain_id
                     GROUP BY   t.domain_id,T.BUSIZ_ID
                       )  t
            on r.busiz_id = t.busiz_id and r.domain_id=t.domain_id
         inner join ftp_busiz_info f
            on r.busiz_id = f.busiz_id and r.domain_id=f.domain_id
         where r.adj_id = :1
           and f.domain_id = :2
         order by r.busiz_id)
`
	FTP_ADJTPRO_GET2 = `select
busiz_id,
busiz_desc,
curve_id,
curve_desc,
domain_id,
cnt
  from (select
               r.busiz_id,
               f.busiz_desc,
               t.curve_id,
               t.curve_desc,
               f.domain_id,
               count (r.busiz_id) over() as cnt
          from FTP_ADJUST_REL r
          left join (select
                           t.busiz_id,
                           t.domain_id,
                       LISTAGG(t.curve_id  , ',') WITHIN GROUP(ORDER BY t.domain_id,T.BUSIZ_ID)   AS curve_id,
                       LISTAGG(d.curve_desc, ',') WITHIN GROUP(ORDER BY t.domain_id,T.BUSIZ_ID) AS curve_desc
                      from FTP_ADJUST_FTP_RESTORE t
                     inner join mas_curve_define d
                        on t.curve_id = d.curve_id
                       and t.domain_id = d.domain_id
                     GROUP BY   t.domain_id,T.BUSIZ_ID
                       ) t
            on r.busiz_id = t.busiz_id and r.domain_id=t.domain_id
         inner join ftp_busiz_info f
            on r.busiz_id = f.busiz_id and r.domain_id=f.domain_id
         where r.adj_id = :1
           and f.domain_id = :2
           and (upper(r.busiz_id) like upper(:3) escape
                '\' or upper(f.busiz_desc) like upper(:4) escape '\')
         order by r.busiz_id)
`

	FTP_ADJTPRO_POST1   = `insert into FTP_ADJUST_FTP_RESTORE(busiz_id,curve_id,domain_id) values(:1,:2,:3)`
	FTP_ADJTPRO_PUT1    = `delete from FTP_ADJUST_FTP_RESTORE where busiz_id=:1 and domain_id=:2`
	FTP_ADJTPRO_PUT2    = `insert into FTP_ADJUST_FTP_RESTORE(busiz_id,curve_id,domain_id) values(:1,:2,:3)`
	FTP_ADJTPRO_DELETE1 = `delete from FTP_ADJUST_FTP_RESTORE where busiz_id=:1 and domain_id=:2`
	FTP_CURVEDEFINE_G1  = `select curve_id,curve_desc from mas_curve_define where domain_id=:1 AND CURVE_type='0'`
	FTP_ADJ_G1          = `select adj_id,adj_desc,adj_type_id from ftp_adj_info where status='0' and domain_id=:1 and adj_type_id='1' order by adj_type_id,adj_id`
	FTP_BP_G1           = `select t.busiz_id,a.busiz_desc
  from FTP_ADJUST_REL t
 inner join ftp_busiz_info a
    on t.busiz_id = a.busiz_id AND t.adj_type_id='2' and t.domain_id=a.domain_id
    where  t.domain_id=:1`
	FTP_INDUSTREE_G1 = `select t.industry_id,t.industry_parent_id ,t.industry_name from FTP_INDUSTRY_INFO t where t.domain_id=:1`
	FTP_REB_G1       = `select t.busiz_id,t1.busiz_desc from FTP_ADJUST_REL t
	inner join ftp_busiz_info t1 on t.busiz_id=t1.busiz_id and t.domain_id=t1.domain_id
	WHERE T.ADJ_ID='604' and t.domain_id=:1`
	FTP_PRDTREE_G1 = `select t.product_id,t.product_parent_id ,t.product_name from FTP_PRODUCT_INFO t where t.domain_id=:1`
	FTP_PRDT_D1    = `select count(*)
  from ftp_adjust_policy t
 where adj_id = '801'
   and t.domain_id = :1
   and ADJ_DYN_DIM in (select t1.product_id
                         from ftp_product_info t1
                        where t1.domain_id = :2
                        start with t1.product_id = :3
                       connect by prior t1.product_id = t1.product_parent_id
                       )`
	FTP_CURVE_D1 = `select count (*) from FTP_BUSIZ_METHOD_RELATION t where T.curve_id=:1 and T.domain_id=:2`
	FTP_CURVE_D2 = `select count (*) from FTP_ADJUST_TERM_LIQUIDITY t where T.curve_id=:1 and T.domain_id=:2`
	FTP_CURVE_D3 = `select count (*) from FTP_ADJUST_FTP_RESTORE t where T.curve_id=:1 and T.domain_id=:2`
	FTP_CURVE_U1 = `update FTP_ADJUST_TERM_LIQUIDITY set reprice_freq_range=:1 where curve_id=:2`
	FTP_CURVE_C1 = `select count(*) from MAS_CURVE_INFO where curve_uuid like :1||'_'||:2||'_'||'%' AND AS_OF_DATE = TO_DATE(:3,'YYYY-MM-DD')`
	FTP_BUSIZ_D1 = `delete from FTP_BUSIZ_REDEMPTION_CURVE t where t.domain_id=:1 and t.busiz_id=:2`
	FTP_BUSIZ_D2 = `select count(*) from FTP_ADJUST_TERM_LIQUIDITY where busiz_id=:1 and domain_id=:2`
	FTP_BUSIZ_D3 = `select count(*) from FTP_ADJUST_FTP_RESTORE where busiz_id=:1 and domain_id=:2`
	FTP_BUSIZ_D4 = `select count(*) from FTP_ADJUST_CAPITAL_RESERVES where busiz_id=:1 and domain_id=:2`
	FTP_BUSIZ_D5 = `delete from FTP_BUSIZ_REDEMPTION_CURVE t where t.domain_id=:1 and t.busiz_id=:2`
	P_HLOG_GET1  = `select op_user_id, op_org, op_app, op_type, op_content, op_ip, op_date, cnt
        from (select op_user_id,
               op_org,
               op_date,
               op_type,
               op_content,
               nvl(op_ip, ' ') as op_ip,
               op_role,
               op_app,
               count(rowid) over() as cnt,
               row_number() over(order by op_date desc) as rk
            from sys_op_logs
         where SQLPARAMS
         
             and op_org   in
               ORGINPARA
         )
         where rk > :2
         and rk <= :3`

	FTP_CURVE_EXPORT = `select 
t.CURVE_ID,
t.CURVE_DESC,
t.ISO_CURRENCY_CD,
to_char(t.CREATE_DATE,'YYYY-MM-DD'),
t.DOMAIN_ID,
t.CURVE_TYPE,
t.REP_ID
from MAS_CURVE_DEFINE t where t.domain_id=:1`
	FTP_CURVE_IMPORT       = `insert into MAS_CURVE_DEFINE values(:1,	:2,	:3,	to_date(:4,'YYYY-MM-DD'),:5,:6,:7)`
	FTP_CURVESTRUCT_EXPORT = `select t.CURVE_ID, t.STRUCT_CODE,t.DOMAIN_ID,t.UUID  
  from MAS_CURVE_INFO_STRUCT_NODE t
 where t.domain_id = :1`
	FTP_CURVESTRUCT_IMPORT = `insert into MAS_CURVE_INFO_STRUCT_NODE(CURVE_ID,STRUCT_CODE,DOMAIN_ID,uuid) values(:1,:2,:3,:4)`
	FTP_CURVEINFO_EXPORT   = `select 
	    t.uuid,
		t.CURVE_UUID,
		to_char(t.AS_OF_DATE,'YYYY-MM-DD'),
		t.YIELD
		 from MAS_CURVE_INFO t where t.curve_uuid  like :1||'%'`
	FTP_CURVEINFO_IMPORT = `insert into MAS_CURVE_INFO values(:1,:2,to_date(:3,'YYYY-MM-DD'),:4)`
	FTP_BUSIZ_EXPORT     = `select t.BUSIZ_ID,
t.BUSIZ_DESC,
t.BUSIZ_UP_ID,
t.FTP_FLAG,
t.BUSIZ_TYPE,
t.DOMAIN_ID,
t.al_flag
  from FTP_BUSIZ_INFO t where t.domain_id=:1`

	FTP_BUSIZ_IMPORT   = `insert into FTP_BUSIZ_INFO values(:1,:2,:3,	:4,	:5,	:6,:7)`
	FTP_BMETHOD_EXPORT = `select t.BUSIZ_ID,
       t.FTP_METHOD_ID,
       t.CURVE_ID,
       t.TERM_CD,
       t.TERM_CD_MULT,
       t.POINT_VAL,
       t.MOVE_DAYS,
       t.MOVE_DAYS_MULT,
       t.DOMAIN_ID
  from FTP_BUSIZ_METHOD_RELATION t
 where t.domain_id = :1`
	FTP_BMETHOD_IMPORT = `insert into FTP_BUSIZ_METHOD_RELATION values(:1,:2,:3,:4,:5,:6,:7,:8,:9)`
	FTP_BADJ_EXPORT    = `select t.BUSIZ_ID,
t.ADJ_ID,
t.ADJ_TYPE_ID,
t.DOMAIN_ID

from FTP_ADJUST_REL t
 where t.domain_id = :1`
	FTP_BADJ_IMPORT = `insert into FTP_ADJUST_REL values(:1,:2,:3,:4)`
	FTP_BCD_EXPORT  = `select
 t.UUID, t.BUSIZ_ID, t.TERM_CD, t.TERM_CD_MULT, t.WEIGHT, t.DOMAIN_ID
  from FTP_BUSIZ_REDEMPTION_CURVE t
 where t.domain_id = :1`
	FTP_BCD_IMPORT        = `insert into FTP_BUSIZ_REDEMPTION_CURVE values(:1,:2,:3,:4,:5,:6)`
	FTP_ADJREVERSE_EXPORT = `select t.BUSIZ_ID,
	t.RESERVE_PERCENT,
	t.RESERVE_RATE,
	t.DOMAIN_ID,
	t.uuid,
	to_char(t.eff_str_date,'YYYY-MM-DD'),
	to_char(t.eff_end_date,'YYYY-MM-DD')
	from FTP_ADJUST_CAPITAL_RESERVES t
    where t.domain_id = :1`
	FTP_ADJREVERSE_IMPORT = `insert into FTP_ADJUST_CAPITAL_RESERVES values(:1,:2,:3,:4,:5,to_date(:6,'YYYY-MM-DD'),to_date(:7,'YYYY-MM-DD'))`
	FTP_ADJRESTORE_EXPORT = `select 
t.BUSIZ_ID,
t.CURVE_ID,
t.DOMAIN_ID,
T.UUID
from FTP_ADJUST_FTP_RESTORE t
 where t.domain_id = :1`
	FTP_ADJRESTORE_IMPORT = `insert into FTP_ADJUST_FTP_RESTORE values(:1,:2,:3,:4)`
	FTP_TERMLIQU_EXPORT   = `select
 t.BUSIZ_ID, t.CURVE_ID, t.DOMAIN_ID, t.UUID, t.REPRICE_FREQ_RANGE
  from FTP_ADJUST_TERM_LIQUIDITY t
 where t.domain_id = :1`
	FTP_TERMLIQU_IMPORT = `insert into FTP_ADJUST_TERM_LIQUIDITY values(:1,:2,:3,:4,:5)`
	FTP_ADJPLOCY_EXPORT = `select
 t.ADJ_ID,
 t.ORG_UNIT_ID,
 t.ISO_CURRENCY_CD,
 t.ADJ_DYN_DIM,
 t.TERM_STR,
 t.TERM_END,
 to_char(t.LAST_DATE,'YYYY-MM-DD'),
 t.ADJ_BP,
to_char(t.EFF_STR_DATE,'YYYY-MM-DD'),
to_char(t.EFF_END_DATE,'YYYY-MM-DD'),
to_char(t.BUZ_STR_DATE,'YYYY-MM-DD'),
to_char(t.BUZ_END_DATE,'YYYY-MM-DD'),
 t.DOMAIN_ID,
 t.MEMO,
t.uuid
  from FTP_ADJUST_POLICY t
 where t.domain_id = :1
`
	FTP_ADJPLOCY_IMPORT = `insert into FTP_ADJUST_POLICY
		values(:1,:2,:3,:4,:5,:6,to_date(:7,'YYYY-MM-DD'),:8,to_date(:9,'YYYY-MM-DD'),to_date(:10,'YYYY-MM-DD'),to_date(:11,'YYYY-MM-DD'),to_date(:12,'YYYY-MM-DD'),:13,:14,:15)`
	FTP_CHECTPATCH_RUN = `select count(*) from FTP_DISPATCH_PRO t where t.dispatch_status='1' and t.domain_id=:1`

	FTP_CURVEINFO_INPUT = `insert into MAS_CURVE_INFO(curve_uuid,as_of_date,yield) values(:1,to_date(:2,'YYYY-MM-DD'),:3)`
	FTP_ADJINFO_P       = `select  sum(rn) from  (
	select count(*)  rn from FTP_ADJ_INFO t
     join ftp_adjust_rel  t1 on  t.adj_id=t1.adj_id  and t.domain_id=t1.domain_id
   where   t.adj_id=:1 and t.domain_id=:2
   union all    
   select count(*) rn  from FTP_ADJ_INFO t
     join ftp_adjust_policy  t1 on  t.adj_id=t1.adj_id  and t.domain_id=t1.domain_id
      where   t.adj_id=:3 and t.domain_id=:4) `
	FTP_ADJINFO_D = `select  sum(rn) from  (
	select count(*)  rn from FTP_ADJ_INFO t
     join ftp_adjust_rel  t1 on  t.adj_id=t1.adj_id  and t.domain_id=t1.domain_id
   where   t.adj_id=:1 and t.domain_id=:2
   union all    
   select count(*) rn  from FTP_ADJ_INFO t
     join ftp_adjust_policy  t1 on  t.adj_id=t1.adj_id  and t.domain_id=t1.domain_id
      where   t.adj_id=:3 and t.domain_id=:4) `
	FTP_ADJLIST_G1 = `
select t.adj_id, t.adj_desc, t.adj_type_id, i.adj_type_name
  from ftp_adj_info t
 inner join ftp_adj_type i
    on t.adj_type_id = i.adj_type_id
 where t.adj_type_id = :1
   and t.status = '0'
   and t.domain_id = :2
   
   and (upper(t.adj_id) like upper(:3) escape '\'
   or upper(t.adj_desc) like upper(:4) escape '\')
 order by t.adj_type_id
`
	FTP_ADJREVERSE_ADDC = `select to_char(eff_str_date,'YYYY-MM-DD'),to_char(eff_end_date,'YYYY-MM-DD') from FTP_ADJUST_CAPITAL_RESERVES  where domain_id=:1 and busiz_id=:2`
	FTP_ADJREVERSE_DEC  = `select to_char(eff_str_date,'YYYY-MM-DD'),to_char(eff_end_date,'YYYY-MM-DD') from FTP_ADJUST_CAPITAL_RESERVES  where domain_id=:1 and busiz_id=:2 and uuid!=:3`
	FTP_GETLATEST_CURVE = `select
				   d.domain_id
				   ,d.curve_id
				   ,max(i.as_of_date)
				from mas_curve_define d
				inner join mas_curve_info_struct_node n
				on d.curve_id = n.curve_id
				and d.domain_id = n.domain_id
				inner join mas_curve_info i
				on n.uuid = i.curve_uuid
				group by d.domain_id,d.curve_id
				`
	FTP_DISPATCHREAL_G = `select 
      count(t.account_number)
  from (
      select account_number, BUSIZ_ID, as_of_date, domain_id
        from (select account_number,
                     BUSIZ_ID,
                     as_of_date,
                     domain_id,
                     row_number() over(order by account_number) rk
                from IINNPPUUTT T
               where as_of_date = to_date(:1, 'YYYY-MM-DD')
                 and domain_id = :2)
      
       where rk > :3
         and rk <= :4
  ) t
inner join ftp_busiz_info t1
   on t.busiz_id = t1.busiz_id
   AND T.DOMAIN_ID = T1.DOMAIN_ID
   and t1.ftp_flag = '0'
  left join OOUUTTPPUUTT r
    on r.as_of_date = t.as_of_date
   and r.account_number = t.account_number
   and t.domain_id = r.domain_id
 where (r.ftp_rate is null or r.ftp_rate = '')
`
	FTP_DISPATREAL_G1 = `update FTP_DISPATCH_PRO set dispatch_status=:1 where dispatch_id=:2 and dispatch_date=to_date(:3,'YYYY-MM-DD') and domain_id=:4`
	FTP_DISPATREAL_G2 = `update FTP_DISPATCH_PRO set dispatch_status=:1,err_msg=:2 where dispatch_id=:3 and dispatch_date=to_date(:4,'YYYY-MM-DD') and domain_id=:5`
	P_CURRENCY_GET1   = `select t.iso_currency_cd,
       t.iso_currency_desc,
       t.owner,
       to_char(t.effective_date,'YYYY-MM-DD'),
       a.status,
       a.status_desc,
       t.sort_id,
       t.memo     
       from MAS_DIM_CURRENCY t
       inner join MAS_DIM_CURRENCY_STATUS a
       on t.status=a.status 
	   where upper(t.iso_currency_cd) like upper(:1) escape '\' or upper(t.iso_currency_desc) like upper(:2) escape '\' order by sort_id`
	P_CURRENCY_GET2 = `select t.iso_currency_cd,
       t.iso_currency_desc,
       t.owner,
       to_char(t.effective_date,'YYYY-MM-DD'),
       a.status,
       a.status_desc,
       t.sort_id,
	   t.memo     
 	   from MAS_DIM_CURRENCY t
       inner join MAS_DIM_CURRENCY_STATUS a
       on t.status=a.status order by sort_id`
	P_CURRENCY_POST1 = `insert into mas_dim_currency(iso_currency_cd, iso_currency_desc,owner,effective_date, status,sort_id,memo) values(:1,:2,:3,to_date(:4,'YYYY-MM-DD'),:5,:6,:7)`
	P_CURRENCY_PUT1  = `update  mas_dim_currency
			set iso_currency_desc=:1 ,status=:2,sort_id=:3,memo=:4
			where iso_currency_cd=:5`

	P_CURRENCY_DELETE1    = `delete from mas_dim_currency where iso_currency_cd=:1`
	P_CURRENCYSTATUS_GET1 = `select status,status_desc from MAS_DIM_CURRENCY_STATUS`
	P_CURRENCY_DC         = `select count (*) from MAS_CURVE_DEFINE t where t.iso_currency_cd=:1`
	P_DOMIAN_I1           = `insert into FTP_BUSIZ_INFO(busiz_id,busiz_desc,busiz_up_id,ftp_flag,busiz_type,domain_id,al_flag) values(:1,:2,:3,:4,:5,:6,:7)`
)
