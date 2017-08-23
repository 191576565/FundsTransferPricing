<!DOCTYPE html>
<html>

	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
		<meta name="description" content="Metro, a sleek, intuitive, and powerful framework for faster and easier web development for Windows Metro Style.">
		<meta name="keywords" content="HTML, CSS, JS, JavaScript, framework, metro, front-end, frontend, web development">
		<meta name="author" content="Sergey Pimenov and Metro UI CSS contributors">

		<link rel='shortcut icon' type='image/x-icon' href='/static/inspinia/favicon.ico' />

		<title>操作日志</title>

		<link href="/static/inspinia/css/bootstrap.min.css" rel="stylesheet">
		<link href="/static/inspinia/font-awesome/css/font-awesome.css" rel="stylesheet">
		<link href="/static/inspinia/css/plugins/dataTables/datatables.min.css" rel="stylesheet">
		<link href="/static/inspinia/css/plugins/iCheck/custom.css" rel="stylesheet">
		<link href="/static/inspinia/css/animate.css" rel="stylesheet">
		<link href="/static/inspinia/css/style.css" rel="stylesheet">
		<!-- Toastr style -->
		<link href="/static/inspinia/css/plugins/toastr/toastr.min.css" rel="stylesheet">
		<link href="/static/inspinia/css/plugins/select2/select2.min.css" rel="stylesheet">
		<link href="/static/inspinia/css/plugins/datapicker/datepicker3.css" rel="stylesheet">
		<!-- Ladda style -->
		<link href="/static/inspinia/css/plugins/ladda/ladda-themeless.min.css" rel="stylesheet">
		<!-- Sweet Alert -->
		<link href="/static/inspinia/css/plugins/sweetalert/sweetalert.css" rel="stylesheet">

		<link href="/static/inspinia/css/plugins/datapicker/datepicker3.css" rel="stylesheet">
		<link rel="stylesheet" type="text/css" href="/static/inspinia/css/plugins/ztree/zTreeStyle.css" />

		<script src="/static/inspinia/js/jquery-1.12.3.min.js"></script>
		<script src="/static/inspinia/js/bootstrap.min.js"></script>
		<script src="/static/inspinia/js/jquery.placeholder.min.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/plugins/metisMenu/jquery.metisMenu.js"></script>
		<script src="/static/inspinia/js/plugins/slimscroll/jquery.slimscroll.min.js"></script>
		<script src="/static/inspinia/js/plugins/dataTables/datatables.min.js"></script>

		<!-- Custom and plugin javascript -->
		<script src="/static/inspinia/js/inspinia.js"></script>
		<script src="/static/inspinia/js/plugins/pace/pace.min.js"></script>
		<script src="/static/inspinia/js/plugins/iCheck/icheck.min.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/plugins/toastr/toastr.min.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/plugins/sweetalert/sweetalert.min.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/common.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/plugins/ztree/jquery.ztree.all.min.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/plugins/select2/select2.full.min.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/plugins/jsKnob/jquery.knob.js" type="text/javascript" charset="utf-8"></script>
		<!-- Data picker -->
		<script src="/static/inspinia/js/plugins/datapicker/bootstrap-datepicker.js"></script>

		<style type="text/css">
			/**{
				overflow-x: hidden;
			}*/
			#wrapper {
				width: 100%;
				height: 100%;
				overflow-y: auto;
			}
			.i-checks {
				margin-left: 0px !important;
				margin-bottom: 2px;
				width: 150px;
			}
			
			.selected {
				background-color: #b7eafb !important;
			}
			
			.input-group-addon {
				border: 1px solid #E5E6E7 !important;
			}
			hr{
				margin: 5px 0px;	
			}
			.modal-header .close{
				padding-top: 4px !important;
				padding-right: 8px !important;
			}
		</style>
	</head>

	<body>
		<div id="wrapper">
			<div class="wrapper wrapper-content animated fadeInRight">
				<div class="row">
					<div class="col-xs-12" id="">
						<div class="ibox float-e-margins">
							<div class="ibox-title">
								<div class="row">
									<div class="col-xs-4">
										<button class="btn btn-warning" type="button" onclick="toggleShow()"><i class="fa fa-search"></i> 条件搜索</button>
									</div>
									<div class="col-xs-8 text-right">
										<div class="form-inline">
											<div class="form-group">
												<input type="hidden" name="SearchTrue" id="SearchTrue" value=""/>
												<input type="text" placeholder="请输入用户帐号" style="width: 200px;" id="UserId" class="form-control ss" value="">
											</div>
											<button class="btn btn-default searchBtn ss" id="" type="button" onclick="simpleSearch()" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
										</div>
									</div>
								</div>
								
								<div class="row conditions" style="display: none;">
									<hr />
									<form id="conditionForm">
									<div class="col-xs-3">
										<div class="form-group"><label>帐号</label>
		                                    <input type="text" class="form-control" name="UserId" id="userId">
		                                </div>
									</div>
									<div class="col-xs-3">
										<div class="form-group"><label>开始时间</label>
		                                   <div class="input-group date full-width">
												<input type="text" class="form-control input-group-addon" id="StartDate" name="StartDate">
											</div>
		                                </div>
									</div>
									<div class="col-xs-3">
										<div class="form-group"><label>结束时间</label>
		                                    <div class="input-group date full-width">
												<input type="text" class="form-control input-group-addon" id="EndDate" name="EndDate">
											</div>
		                                </div>
									</div>
									<div class="col-xs-3">
										<div class="form-group"><label>子系统</label>
		                                    <input type="text" class="form-control" name="OpApp" id="OpApp">
		                                </div>
									</div>
									<div class="col-xs-4">
										<div class="form-group"><label>业务类型</label>
		                                    <input type="text" class="form-control" name="OpType" id="OpType">
		                                </div>
									</div>
			
									<div class="col-xs-4">
										<button class="btn btn-primary btn-block" type="button" onclick="search()" style="margin-top: 20px;">搜&nbsp;索</button>
									</div>
									</form>
								</div>
							</div>
							<div class="ibox-content">
								<div class="table-responsive">
									<table class="table table-striped table-bordered table-hover opLogs"  style="width: 99%;">
										<thead>
											<tr>
												<th>用户帐号</th>
												<th>所属机构</th>
												<th>子系统</th>
												<th>业务类型</th>
												<th>操作信息</th>
												<th>客户端IP</th>
												<th>操作时间</th>
											</tr>
										</thead>
									</table>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</body>
	<script type="text/javascript">
		var opLogs; 
		$(function(){
			initTable();
			
//			$('.opLogs tbody').on('click', 'tr', function() {
//				if($(this).find('input:radio').get(0).checked) {
//					$(this).find('input:radio').get(0).checked = false;
//				} else {
//					$(this).find('input:radio').get(0).checked = true;
//				}
//			});
			
			$.get('/platform/AllRole', {r:Math.random()*10000000000000}, function(data){
				appendOption('RoleType', data, 'RoleName', 'RoleName');
			});
			
			$('.input-group-addon').datepicker({
				autoclose: true,
				format: 'yyyy-mm-dd'
			});
		});
		
		function initTable() {
			opLogs = $('.opLogs').DataTable({
				sAjaxSource: '/platform/HandleLogs?r='+Math.random()*10000000000000,
				info: true,
				iDisplayLength: 15,
				aoColumns: [{
					"data": "OpUserId"
				}, {
					"data": "OpOrg"
				}, {
					"data": "OpApp"
				}, {
					"data": "OpType"
				}, {
					"data": "OpContent"
				}, {
					"data": "OpIp"
				}, {
					"data": "OpDate"
				}],
//				fnRowCallback: function(nRow, aData, iDisplayIndex, iDisplayIndexFull) {
//					$('td:eq(0)', nRow).html('<div class="radio"><input type="radio" value="' + iDisplayIndex + '" name="tableRadio"><label></label></div>');
//				},
				fnServerParams: function(aoData) {
					aoData.push({
						name: "SearchTrue",
						value: $('#SearchTrue').val()
					});
					aoData.push({
						name: "UserId",
						value: $('#UserId').val()
					});
					aoData.push({
						name: "StartDate",
						value: $('#StartDate').val()
					});
					aoData.push({
						name: "EndDate",
						value: $('#EndDate').val()
					});
					aoData.push({
						name: "OpApp",
						value: $('#OpApp').val()
					});
					aoData.push({
						name: "RoleType",
						value: $('#RoleType').val()
					});
					aoData.push({
						name: "OpType",
						value: $('#OpType').val()
					});
				}
			});
		}
		
		function simpleSearch(){
			$('#conditionForm')[0].reset();
			$('#SearchTrue').val('true');
//			opLogs.ajax.reload(null, true);
			opLogs.ajax.url('/platform/HandleLogs?r='+Math.random()*10000000000000).load();
		}
		
		function search(){
			$('#SearchTrue').val('true');
			$('#UserId').val($('#userId').val());
//			opLogs.ajax.reload(null, true);
			opLogs.ajax.url('/platform/HandleLogs?r='+Math.random()*10000000000000).load();
			$('#UserId').val('');
		}
		
		function toggleShow(){
			$('#conditionForm')[0].reset();
			$('.conditions').slideToggle();
			$('.ss').attr('disabled', !$('.ss').attr('disabled'));
		}
	</script>
</html>