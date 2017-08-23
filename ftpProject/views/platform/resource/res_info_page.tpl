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

		<title>菜单管理</title>

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

	<body style="overflow: hidden;">
		<div id="wrapper" class="full-height">
			<div class="wrapper wrapper-content animated fadeInRight full-height">
				<div class="row">
					<div class="col-xs-5" id="">
						<div class="ibox float-e-margins">
							<div class="ibox-title">
								<div class="row">
									<h3><i class="fa fa-paper-plane-o"></i>&nbsp;角色分类</h3>
								</div>

							</div>
							<div class="ibox-content">
								<div class="table-responsive">
									<table class="table table-striped table-bordered table-hover roleType"  style="width: 99%;">
										<thead>
											<tr>
												<th style="width: 40px;">
												</th>
												<th>角色描述</th>
												<th>角色编号</th>
											</tr>
										</thead>
									</table>
								</div>
							</div>
						</div>
					</div>
					
					<div class="col-xs-7">
						<div class="ibox float-e-margins">
							<div class="ibox-title">
								<div class="row">
									<h3 style="float: left;"><i class="fa fa-anchor"></i>&nbsp;角色可分配资源</h3>
									<input type="hidden" id="roleType" value="" />
									<button resid="101030101000" class="btn btn-info updateBtn res" style="float: right;margin: 0px;border: 0px;" disabled="disabled" onclick="updateRoleMenu()">更新</button>
								</div>
							</div>
							<div class="ibox-content">
								<div id="opRes"  style="overflow: auto;">
									<ul id="resTree" class="ztree"></ul>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</body>
	<script type="text/javascript">
		var s = {
			view: {
				dblClickExpand: false,
				selectedMulti: false
			},
			data: {
				simpleData: {
					enable: true
				}
			},
			callback: {
				onDblClick: zTreeOnDblClickSimple,
				onClick: zTreeOnClickSimple
			},
			check: {
				enable: true,
				chkboxType: { "Y": "ps", "N": "s" }
			}
		};
		function zTreeOnDblClickSimple(event, treeId, treeNode) {
			//alert(treeNode ? treeNode.tId + ", " + treeNode.name : "isRoot");
//			$('#userOrgUnitId').val(treeNode.id);
//			$('#userOrgUnitName').val(treeNode.name);
//
////			$('#orgTree').modal('hide');
//			$('#orgTree').hide();
		};

		function zTreeOnClickSimple(event, treeId, treeNode) {
			var treeObj = $.fn.zTree.getZTreeObj(treeId);
			treeObj.expandNode(treeNode, null, null, null);

			return true;
		};
	
		var roleType; 
		$(function(){
			//控制按钮权限
			$.get('/platform/DefaultMenu',{TypeId:2,Id:'101030100000',r:Math.random()*10000000000000}, function(data){
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
			
//			$('#wrapper').height(document.body.clientHeight);
//			$('#opRes').height(document.body.clientHeight-200);
			
			$('.roleType tbody').on('click', 'tr', function() {
				$(this).find('input:radio').get(0).checked = true;
				
				//获取角色类型对应可分配的最大已有菜单
				var roleType=$('input:checked').parent().parent().next().next().html();
				$('#roleType').val(roleType);
				$('.updateBtn').removeAttr('disabled');
				$.get('/platform/RoleTypeRes', {RoleTypeId:roleType,r:Math.random()*10000000000000}, function(data){
					var rs=JSON.parse(data);
					var treeObj = $.fn.zTree.getZTreeObj("resTree");
					treeObj.expandAll(false);
					treeObj.checkAllNodes(false);
					if(rs!=null){
						rs.forEach(function(e){
							var node = treeObj.getNodeByParam('id', e.ResId, null);
							if(node){
								treeObj.checkNode(node, true, false);
							}
						});
					}
					
					var nodes = treeObj.getCheckedNodes(true);
					nodes.forEach(function(e){
						treeObj.expandNode(e, true, false, false);
					});
				})
			});
			
			$.get('/platform/AllTree' + '?r=' + Math.random(100000000000000), function(data) {
				$.fn.zTree.init($("#resTree"), s, JSON.parse(data)); //树
			});
		});
		
		function initTable() {
			roleType = $('.roleType').DataTable({
				sAjaxSource: '/platform/RoleStage?r='+Math.random()*10000000000000,
				info: false,
				bPaginate: false,
				aoColumns: [{
					"data": null
				}, {
					"data": "RoletypeDesc"
				}, {
					"data": "RoletypeId"
				}],
				fnRowCallback: function(nRow, aData, iDisplayIndex, iDisplayIndexFull) {
					$('td:eq(0)', nRow).html('<div class="radio"><input type="radio" value="' + iDisplayIndex + '" name="tableRadio"><label></label></div>');
				}
			});
		}
		
		function updateRoleMenu(){
			var roleType=$('#roleType').val();
			var treeObj = $.fn.zTree.getZTreeObj("resTree");
			var nodes = treeObj.getCheckedNodes(true);
			var json=[];
			nodes.forEach(function(e){
				json.push({"ResId":e.id});
			});
			
			$.ajax({
				type:"put",
				url:"/platform/RoleTypeRes?JSON="+JSON.stringify(json)+"&RoleTypeId="+roleType+"&r="+Math.random()*10000000000000,
				async:true,
				error: function(msg){
	                console.log(msg.responseText);
	            },
	            success: function(data){
					var r=JSON.parse(data);
	                if(r.ErrorCode==='1'){
						toastr.success('角色菜单保存成功！');
	                }else {
	                    toastr.error(r.ErrorMsg);
	                    return ;
	                }
	            }
			});
		}
	</script>

</html>