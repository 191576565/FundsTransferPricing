<!DOCTYPE html>
<html>

	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, curr-scalable=no">
		<meta name="description" content="Metro, a sleek, intuitive, and powerful framework for faster and easier web development for Windows Metro Style.">
		<meta name="keywords" content="HTML, CSS, JS, JavaScript, framework, metro, front-end, frontend, web development">
		<meta name="author" content="Sergey Pimenov and Metro UI CSS contributors">

		<link rel='shortcut icon' type='image/x-icon' href='/static/inspinia/favicon.ico' />

		<title>存款准备金调节</title>

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
				text-align: left;
			}
			.tab-pane,.panel-body{
				height: 100%;
			}
			.panel-body .batch{
				float: left;
				margin: 20px 20px;
				text-align: center;
			}
			hr{
				margin: 5px 0px;	
			}
			tr td:first-child{
				text-align: center;
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
									<div class="col-xs-8">
										<button class="btn btn-primary" type="button" onclick="goBack()"> <i class="fa fa-reply"></i> 返回</button>
										<!--<button resid="208020201010" class="btn btn-primary editBtn res" type="button" onclick="showEditDialog('2')"><i class="fa fa-wrench"></i> 配置</button>-->
										<!--<button resid="101040203000" class="btn btn-danger deleteBtn res" type="button" onclick="deleteCurrInfo()"><i class="fa fa-times"></i> 删除</button>-->
										<button resid="" class="btn btn-primary " type="button" onclick="showEditDialog('1')"> <i class="fa fa-plus"></i> 新增</button>
										<button resid="" class="btn btn-info editBtn " type="button" onclick="showEditDialog('2')"><i class="fa fa-paste"></i> 编辑</button>
										<button resid="" class="btn btn-danger deleteBtn " type="button" onclick="deleteRe()"><i class="fa fa-times"></i> 删除</button>
									</div>
									<div class="col-xs-4 text-right">
										<div class="form-inline">
											<div class="form-group">
												<input type="text" placeholder="请输入业务单元编码或名称" style="width: 200px;" id="keyword" class="form-control" value="">
											</div>
											<button class="btn btn-default searchBtn" type="button" onclick="search()" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
										</div>
									</div>
								</div>

							</div>
							<div class="ibox-content">
								<div class="table-responsive">
									<table class="table table-striped table-bordered table-hover currs"  style="width: 99%;">
										<thead>
											<tr>
												<th style="width: 40px;">
												</th>
												<th>业务单元编码</th>
												<th>业务单元名称</th>
												<th>准备金比例(%)</th>
												<th>准备金利率(%)</th>
												<th>生效起始日</th>
												<th>生效截止日</th>
												<th>域描述</th>
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
		
		<div class="modal inmodal" id="editDialog" tabindex="-1" role="dialog" aria-hidden="true">
			<div class="modal-dialog modal-lg">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 id="addInfo">新增准备金调节项</h2>
						<h2 id="editInfo">编辑准备金调节项</h2>
						<input type="hidden" id="type"/>
					</div>
					<div class="modal-body">
						<form id="reserveInfo">
							<input type="hidden" name='Uuid' id='Uuid'/>
							<div class="row">
								<div class="col-xs-6">
										<div class="form-group">
											<label>业务单元:</label>
											<select name="busiz_id" id="Busiz_id" class="form-control">
											</select>
										</div>
										<!--<div class="form-group">
											<label>业务单元编码:</label>
											<input type="text" placeholder="" class="form-control" name="busiz_id" id="Busiz_id" readonly="readonly">
										</div>-->
								</div>
								<!--<div class="col-xs-6">
										<div class="form-group">
											<label>业务单元名称:</label>
											<input type="text" placeholder="" class="form-control" name="busiz_desc" id="Busiz_desc" >
										</div>
								</div>-->
								<div class="col-xs-6">
										<div class="form-group">
											<label>准备金比例(%):</label>
											<input type="text" placeholder="" class="form-control" name="reserve_percent" id="Reserve_percent">
										</div>
								</div>
								<div class="col-xs-6">
										<div class="form-group">
											<label>准备金利率(%):</label>
											<input type="text" placeholder="" class="form-control" name="reserve_rate" id="Reserve_rate">
										</div>
								</div>
								<div class="col-xs-6">
										<div class="form-group">
											<label>生效起始日:</label>
											<div class="input-group date full-width">
												<input type="text" class="form-control input-group-addon" id="startEffectDay" name="str_date">
											</div>
										</div>
								</div>
								<div class="col-xs-6">
									<div class="form-group">
										<label>生效截止日:</label>
										<div class="input-group date full-width">
											<input type="text" class="form-control input-group-addon" id="endEffectDay" name="end_date">
										</div>
									</div>
								</div>
							</div>
						</form>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="saveReserveInfo()">保存</button>
					</div>
				</div>
			</div>
		</div>
		
	</body>
	<script type="text/javascript">
		var currs; 
		$(function(){
			//控制按钮权限
			$.get('/platform/DefaultMenu',{TypeId:2,Id:'30208020201000',r:Math.random()*10000000000000}, function(data){
				var rs=JSON.parse(data);
				//将所有有权限编码的按钮先禁用
//				$('.res').attr('disabled', true);
				$('.res').hide();
				if(rs!=null){
					rs.forEach(function(e){
//						$("button[resid='"+e.Res_id+"']").attr('disabled', false);
						$("button[resid='"+e.Res_id+"']").show();
					});
				}
			})
			
			initTable();
			
			$('.currs tbody').on('click', 'tr', function() {
				$(this).find('input:radio').get(0).checked = true;
			});
			
			$.get('/mas/ftp/FtpReBusiz', {r:Math.random()*10000000000000}, function(data){
				appendOption('Busiz_id', data, 'BusizId', 'BusizDesc');
			});
			
			$('.input-group-addon').datepicker({
				autoclose: true,
				format: 'yyyy-mm-dd'
			});	
		});
		
		function initTable() {
			currs = $('.currs').DataTable({
				sAjaxSource: '/mas/ftp/adjust/reserve/config?AdjId=604&r='+Math.random()*10000000000000,
				info: true,
				iDisplayLength: 15,
				aoColumns: [{
					"data": null
				}, {
					"data": "Busiz_id"
				}, {
					"data": "Busiz_desc"
				}, {
					"data": "Reserve_percent"
				}, {
					"data": "Reserve_rate"
				}, {
					"data": "Eff_start",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Eff_end",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Domain_desc"
				}],
				fnRowCallback: function(nRow, aData, iDisplayIndex, iDisplayIndexFull) {
					$('td:eq(0)', nRow).html('<div class="radio"><input type="radio" value="' + iDisplayIndex + '" name="tableRadio"><label></label></div>');
				},
				fnServerParams: function(aoData) {
					aoData.push({
						name: "keyword",
						value: $('#keyword').val().toUpperCase()
					});
				}
			});
		}
		
		function showEditDialog(type){ 
			$('#type').val(type);
			if(type==='1'){
				$('#addInfo').show();
				$('#editInfo').hide();
				
				$('#reserveInfo')[0].reset();
				$('#Busiz_id').removeAttr('readonly');
				$('.input-group-addon').datepicker('update', '');
			}else if(type==='2'){
				if($('.currs input:checked').length===0){
					toastr.warning('请选择一行');
					return;
				}
				
				$('#addInfo').hide();
				$('#editInfo').show();
				$('#Busiz_id').attr('readonly', 'readonly');
				
				var info=currs.data()[$('.currs input:checked').val()];
				$('#Busiz_id').val(info.Busiz_id);
				$('#Reserve_percent').val(info.Reserve_percent);
				$('#Reserve_rate').val(info.Reserve_rate);
				$('#startEffectDay').val(info.Eff_start.substring(0, 10));
				$('#endEffectDay').val(info.Eff_end.substring(0, 10));
				$('#Uuid').val(info.Uuid);
			}
			
			$('#editDialog').modal('show');
		}
		
		function saveReserveInfo(){
			var percent=$('#Reserve_percent').val();
			var rate=$('#Reserve_rate').val();
			var startEffectDay=$('#startEffectDay').val();
			var endEffectDay=$('#endEffectDay').val();
			
			var regu = /^[0-9]+\.?[0-9]*$/;
			if(!regu.test(percent)){ //正数
				toastr.warning('准备金比例为正数');
				return;
			}else {
				if(parseFloat(percent)<0 || parseFloat(percent)>100){
					toastr.warning('准备金比例值介于0-100，单位为%');
					return;
				}
			}
			
			if(!regu.test(rate)){ //正数
				toastr.warning('准备金利率为正数');
				return;
			}else {
				if(parseFloat(rate)<0 || parseFloat(rate)>100){
					toastr.warning('准备金利率值介于0-100，单位为%');
					return;
				}
			}
			
			if(startEffectDay==='' || endEffectDay===''){
				toastr.warning('请选择生效起止日');
				return ;
			}else {
				if(startEffectDay>=endEffectDay){
					toastr.warning('生效起始日需要小于生效截止日');
					return ;
				}
			}
			
			if(startEffectDay!='' && endEffectDay!=''){
				if(startEffectDay>=endEffectDay){
					toastr.warning('生效起始日需要小于生效截止日');
					return ;
				}
			}	
			
			var type=$('#type').val();
			if(type==='1'){
				$.post('/mas/ftp/adjust/reserve/config?r='+Math.random()*10000000000000, $('#reserveInfo').serialize(),function(data){
					var rs=JSON.parse(data);
		            if(rs.ErrorCode==='1'){
		                toastr.success('调节项信息保存成功！');
//						currs.ajax.reload(null, true);
						currs.ajax.url('/mas/ftp/adjust/reserve/config?AdjId=604&r='+Math.random()*10000000000000).load();
						$('#editDialog').modal('hide');
		            }else {
		            		toastr.error(rs.ErrorMsg);
		                return false;
		            }
				});
			}else {
				$.ajax({
					type:"put",
					url:"/mas/ftp/adjust/reserve/config?r="+Math.random()*10000000000000+"&"+$('#reserveInfo').serialize(),
					async:true,
					error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
							toastr.success('调节项信息编辑成功！');
//							currs.ajax.reload(null, true);
							currs.ajax.url('/mas/ftp/adjust/reserve/config?AdjId=604&r='+Math.random()*10000000000000).load();
							$('#editDialog').modal('hide');
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
		            }
				});
			}
		}
		
		function search(){
//			currs.ajax.reload(null, true);
			currs.ajax.url('/mas/ftp/adjust/reserve/config?AdjId=604&r='+Math.random()*10000000000000).load();
		}
		
		function goBack(){
			window.location.href='/mas/ftp/adjust/page?r='+Math.random()*10000000000000;
		}
		
		function deleteRe(){
			if($('.currs input:checked').length===0){
				toastr.warning('请选择一行进行编辑');
				return;
			}
			
			var info=currs.data()[$('.currs input:checked').val()];
			swal({
				title: "删除！",
				text: "是否删除选中的调节项信息？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				var json=[{Uuid:info.Uuid,Busiz_id:info.Busiz_id}];
				$.ajax({
		            type:"delete",
		            url:"/mas/ftp/adjust/reserve/config?JSON="+JSON.stringify(json)+"&r="+Math.random()*10000000000000,
		            error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
//		                    toastr.success(rs.ErrorMsg);
		                    swal("删除!", rs.ErrorMsg, "success");
//							policys.ajax.reload(null, true);
							currs.ajax.url('/mas/ftp/adjust/reserve/config?AdjId=604&r='+Math.random()*10000000000000).load();
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
		            }
		        });
			});
		}
	</script>

</html>