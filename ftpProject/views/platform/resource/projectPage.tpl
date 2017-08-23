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

		<title>域定义</title>

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
		<script src="/static/inspinia/js/plugins/validate/jquery.validate.min.js" type="text/javascript" charset="utf-8"></script>

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
			<!--<div class="row wrapper border-bottom white-bg page-heading">
                <div class="col-sm-4">
                    <h2>曲线定义</h2>
                    <ol class="breadcrumb">
                    		<li>
							<a href="#">首页</a>
						</li>
                        <li>
                            <a href="#">FTP</a>
                        </li>
                        <li class="active">
                            <strong>曲线定义</strong>
                        </li>
                    </ol>
                </div>
            </div>-->
			<div class="wrapper wrapper-content animated fadeInRight">
				<div class="row">
					<div class="col-xs-12" id="">
						<div class="ibox float-e-margins">
							<div class="ibox-title">
								<div class="row">
									<div class="col-xs-4">
										<button resid='101040101000' class="btn btn-primary res" type="button" onclick="showEditDialog('1')"> <i class="fa fa-plus"></i> 新增</button>
										<button resid='101040102000' class="btn btn-info editBtn res" type="button" onclick="showEditDialog('2')"><i class="fa fa-paste"></i> 编辑</button>
										<button resid='101040103000' class="btn btn-danger deleteBtn res" type="button" onclick="deleteDomain()"><i class="fa fa-times"></i> 删除</button>
									</div>
									<div class="col-xs-8 text-right">
										<div class="form-inline">
											<div class="form-group">
												<input type="text" placeholder="请输入域编号或名称" style="width: 200px;" id="keyword" class="form-control" value="">
											</div>
											<button class="btn btn-default searchBtn" type="button" onclick="search()" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
										</div>
									</div>
								</div>

							</div>
							<div class="ibox-content">
								<div class="table-responsive">
									<table class="table table-striped table-bordered table-hover domain"  style="width: 99%;">
										<thead>
											<tr>
												<th style="width: 40px;">
												</th>
												<th>域编号</th>
												<th>域名称</th>
												<!--<th>域状态</th>-->
												<th>创建日期</th>
												<th>用户</th>
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
						<h2 id="addDomain">新增域信息</h2>
						<h2 id="editDomain">编辑域信息</h2>
						<input type="hidden" id="type"/>
					</div>
					<form id="domainInfo">
					<div class="modal-body">
						
							<div class="row">
								<div class="col-xs-12">
										<div class="form-group">
											<label>域编码:</label>
											<input type="text" placeholder="" class="form-control" name="domainId" id="domainId">
										</div>
								</div>
								<div class="col-xs-12">
										<div class="form-group">
											<label>域名称:</label>
											<input type="text" placeholder="" class="form-control" name="domainDesc" id="domainDesc">
										</div>
										
								</div>
								<!--<div class="col-xs-12">
									<div class="form-group">
											<label>域状态:</label>
											<select name="domainStatus" id="domainStatus" class="form-control">
												<option value="0">正常</option>
						                        <option value="1">锁定</option>
						                        <option value="2">失效</option>
											</select>
										</div>
								</div>-->
							</div>
						
					</div>

					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="saveDomain()">保存</button>
					</div>
					</form>
				</div>
			</div>
		</div>
	</body>
	<script type="text/javascript">
		var domain; 
		$(function(){
			//控制按钮权限
			$.get('/platform/DefaultMenu',{TypeId:2,Id:'101040100000',r:Math.random()*10000000000000}, function(data){
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
			
//			$('#domainInfo').validate({
//				rules:{
//					domainId: "required",
//   	 			domainDesc: "required"
//				},
//				submitHandler: function(){
////			    		saveDomain();
//			   }
//			});
			
			$('.domain tbody').on('click', 'tr', function() {
				$(this).find('input:radio').get(0).checked = true;
			});
			
//			$('input:radio').click(function(e){
////				$(this)[0].checked;
//				//阻止冒泡
////				e = e || event;
////				e.stopPropagation ? e.stopPropagation() : e.cancelBubble = true;
//			});
		});
		
		function initTable() {
			domain = $('.domain').DataTable({
				sAjaxSource: '/platform/DomainMgr?r='+Math.random()*10000000000000,
				info: true,
				iDisplayLength: 10,
				aoColumns: [{
					"data": null
				}, {
					"data": "Project_id"
				}, {
					"data": "Project_name"
				}/*, {
					"data": "Project_status"
				}*/, {
					"data": "Maintance_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "User_id"
				}],
				fnRowCallback: function(nRow, aData, iDisplayIndex, iDisplayIndexFull) {
					$('td:eq(0)', nRow).html('<div class="radio"><input type="radio" value="' + iDisplayIndex + '" name="tableRadio"><label></label></div>');
				},
				fnServerParams: function(aoData) {
					aoData.push({
						name: "keyword",
						value: $('#keyword').val()
					});
				}
			});
		}
		
		function showEditDialog(type){ 
			$('#type').val(type);
			if(type==='1'){
				$('#addDomain').show();
				$('#editDomain').hide();
				
				$('#domainInfo')[0].reset();
				$('#domainId').removeAttr('readonly');
			}else if(type==='2'){
				if($('.domain input:checked').length===0){
					toastr.warning('请选择一行域信息进行编辑');
					return;
				}
				
				$('#addDomain').hide();
				$('#editDomain').show();
				$('#domainId').attr('readonly', 'readonly');
				
				var info = domain.data()[$('.domain input:checked').val()];
				
				$('#domainId').val(info.Project_id);
				$('#domainDesc').val(info.Project_name);
//				if(info.Project_status=='正常'){
//					$('#domainStatus').val('0');
//				}else if(info.Project_status=='锁定'){
//					$('#domainStatus').val('1');
//				}else if(info.Project_status=='失效'){
//					$('#domainStatus').val('2');
//				}
			}
			$('#editDialog').modal('show');
		}
		
		function saveDomain(){
			//判断是否是数字或字母 
    			var id = /^[0-9a-zA-Z\_]+$/;
    			//判断是否是汉字、字母、数字组成 
    			var name = /^[0-9a-zA-Z\u4e00-\u9fa5]+$/;
    			
    			var domainId=$('#domainId').val();
    			var domainDesc=$('#domainDesc').val();
    			if(!id.test(domainId)){
    				toastr.warning('域编码必须是数字或字母组成');
    				return;
    			}
    			if(!name.test(domainDesc)){
    				toastr.warning('域名称必须是数字、汉字或字母组成');
    				return;
    			}
			
			if($('#type').val()==='1'){ //新增
				$.post('/platform/DomainMgr?r='+Math.random()*10000000000000, $('#domainInfo').serialize()+"&domainUpId=Root",function(data){
//					var rs=JSON.parse(data);
//		            if(rs.ErrorCode==='1'){
//		                cvs.ajax.reload(null, true);
//		                toastr.success('曲线点值保存成功！');
//		            }else {
//		            		toastr.error(rs.ErrorMsg, '错误');
//		                return false;
//		            }
					var rs=JSON.parse(data);
	                if(rs.ErrorCode==='1'){
		                toastr.success('域信息保存成功！');
//						domain.ajax.reload(null, true);
						domain.ajax.url('/platform/DomainMgr?r='+Math.random()*10000000000000).load();
						$('#editDialog').modal('hide');
	                }else {
	                    toastr.error(rs.ErrorMsg);
	                    return ;
	                }
					
				});
			}else if($('#type').val()==='2'){ //编辑
				$.ajax({
					type:"put",
					url:"/platform/DomainMgr?"+$('#domainInfo').serialize()+"&domainUpId=Root"+"&r="+Math.random()*10000000000000,
					async:true,
					error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
//						var rs=JSON.parse(data);
//		                if(rs.ErrorCode==='1'){
////		                    toastr.success(rs.ErrorMsg);
//		                    cvs.ajax.reload(null, true);
//		                    swal("删除!", "曲线已删除", "success");
//		                }else {
//		                    toastr.error('rs.ErrorMsg','错误！');
//		                    return ;
//		                }
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
			                toastr.success('域信息保存成功！');
//							domain.ajax.reload(null, true);
							domain.ajax.url('/platform/DomainMgr?r='+Math.random()*10000000000000).load();
							$('#editDialog').modal('hide');
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
		            }
				});
			}
			
		}
		
		function deleteDomain(){
			if($('.domain input:checked').length===0){
				toastr.warning('请选择一行域信息进行删除');
				return;
			}
			
			var info = domain.data()[$('.domain input:checked').val()];
			swal({
				title: "删除！",
				text: "是否删除选中的域信息？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				var json=[{'Project_id':info.Project_id}];
				
				$.ajax({
		            type:"delete",
		            url:"/platform/DomainMgr?JSON="+JSON.stringify(json)+"&r="+Math.random()*10000000000000,
		            error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
		                    swal("删除!", "域信息已删除", "success");
//							domain.ajax.reload(null, true);
							domain.ajax.url('/platform/DomainMgr?r='+Math.random()*10000000000000).load();
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
		            }
		        });
			});
		}
		
		function search(){
//			domain.ajax.reload(null, true);
			domain.ajax.url('/platform/DomainMgr?r='+Math.random()*10000000000000).load();
		}
	</script>

</html>