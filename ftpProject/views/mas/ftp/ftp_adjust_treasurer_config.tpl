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

		<title>司库利润还原调节</title>

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
										<button resid="30208020201010" class="btn btn-info editBtn res" type="button" onclick="showEditDialog('2')"><i class="fa fa-wrench"></i> 配置</button>
										<!--<button resid="101040203000" class="btn btn-danger deleteBtn res" type="button" onclick="deleteCurrInfo()"><i class="fa fa-times"></i> 删除</button>-->
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
												<th>曲线编码</th>
												<th>曲线名称</th>
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
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2>选择曲线</h2>
					</div>
					<div class="modal-body"  style="max-height: 500px;overflow: auto;">
						<div class="row">
							<table class="table table-bordered tlps">
								<thead>
									<tr>
										<th></th>
										<th>曲线编码</th>
										<th>曲线描述</th>
									</tr>
								</thead>
								<tbody>
									<!--<tr>
										<td style="width: 40px;"><label class="checkbox-inline i-checks" style="width: 30px;"> <input type="checkbox"></label></td>
										<td>aa</td>
										<td>aaa</td>
									</tr>-->
								</tbody>
							</table>
						</div>
					</div>

					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="saveTermInfo()">保存</button>
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
			
			$.get('/mas/ftp/adjust/termLiq/tlp', {CurveType:2,r:Math.random()*10000000000000}, function(data){
				var rs=JSON.parse(data);
				if(rs!=null){
					$('.tlps tbody').html('');
					rs.forEach(function(e, index){
						var temp="<tr><td style='width: 40px;text-align:center;'><label curveid='"+e.Curve_id+"' curvename='"+e.Curve_desc+"' class='checkbox-inline i-checks' style='width: 30px;'> <input type='checkbox'></label></td>"
								+"<td>"+e.Curve_id+"</td><td>"+e.Curve_desc+"</td>";
						$('.tlps tbody').append(temp);
					});
				}
				iCheckInit();
			});
			
		});
		
		function iCheckInit(){
			$('.i-checks').iCheck({
				checkboxClass: 'icheckbox_square-green',
				radioClass: 'iradio_square-green',
			});
		}
		
		function initTable() {
			currs = $('.currs').DataTable({
				sAjaxSource: '/mas/ftp/adjust/treasurep/config?AdjId=603&r='+Math.random()*10000000000000,
				info: true,
				iDisplayLength: 15,
				aoColumns: [{
					"data": null
				}, {
					"data": "Busiz_id"
				}, {
					"data": "Busiz_desc"
				}, {
					"data": "Curve_id"
				}, {
					"data": "Curve_desc"
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
		
		var type=0;
		function showEditDialog(){ 
			if($('.currs input:checked').length===0){
				toastr.warning('请选择一行');
				return;
			}
			
			var info=currs.data()[$('.currs input:checked').val()];
			
			if(info.Curve_id==='' && info.Curve_desc===''){
				type=1; //新增
			}else {
				type=2; //编辑
			}
			
			//勾选已有曲线
			var curves=info.Curve_id.split(',');
			if(info.Curve_id!=''){
				$('.icheckbox_square-green').iCheck('uncheck');
				curves.forEach(function(e){
					$("label[curveid="+e+"]").find('.icheckbox_square-green').iCheck('check');
				});
			}
			
			$('#editDialog').modal('show');
		}
		
		function saveTermInfo(){
//			if($('.checked').length!=1){
//				toastr.warning('请选择且只能选择一条曲线');
//				return;
//			}
			
			var Busiz_id=currs.data()[$('.currs input:checked').val()].Busiz_id;
			var json=[];
			$('.checked').each(function(){
				json.push({"Curve_id":$(this).parent().attr('curveid'), "Busiz_id":Busiz_id });
			});
			
			if(type===1){
				$.post('/mas/ftp/adjust/treasurep/config?JSON='+JSON.stringify(json)+"&BusizId="+Busiz_id+"&r="+Math.random()*10000000000000,function(data){
					var rs=JSON.parse(data);
		            if(rs.ErrorCode==='1'){
		                toastr.success('调节项信息保存成功！');
//						currs.ajax.reload(null, true);
						currs.ajax.url('/mas/ftp/adjust/treasurep/config?AdjId=603&r='+Math.random()*10000000000000).load();
						$('#editDialog').modal('hide');
		            }else {
		            		toastr.error(rs.ErrorMsg);
		                return false;
		            }
				});
			}else {
				$.ajax({
					type:"put",
					url:'/mas/ftp/adjust/treasurep/config?JSON='+JSON.stringify(json)+"&BusizId="+Busiz_id+"&r="+Math.random()*10000000000000,
					async:true,
					error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
							toastr.success('调节项信息编辑成功！');
//							currs.ajax.reload(null, true);
							currs.ajax.url('/mas/ftp/adjust/treasurep/config?AdjId=603&r='+Math.random()*10000000000000).load();
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
			currs.ajax.url('/mas/ftp/adjust/treasurep/config?AdjId=603&r='+Math.random()*10000000000000).load();
		}
		
		function goBack(){
			window.location.href='/mas/ftp/adjust/page?r='+Math.random()*10000000000000;
		}
	</script>

</html>