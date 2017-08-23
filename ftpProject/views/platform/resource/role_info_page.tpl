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

		<title>角色管理</title>

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
			
			.tab-pane,
			.panel-body {
				height: 100%;
			}
			
			.panel-body .batch {
				float: left;
				margin: 20px 20px;
				text-align: center;
			}
			
			hr {
				margin: 5px 0px;
			}
			
			tr td:first-child {
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
										<button resid="101050201000" class="btn btn-primary res" type="button" onclick="showEditDialog('1')"> <i class="fa fa-plus"></i> 新增</button>
										<button resid="101050202000" class="btn btn-info editBtn res" type="button" onclick="showEditDialog('2')"><i class="fa fa-paste"></i> 编辑</button>
										<button resid="101050203000" class="btn btn-danger deleteBtn res" type="button" onclick="deleteRoleInfo()"><i class="fa fa-times"></i> 删除</button>
										<button resid="101050204000" class="btn btn-warning res" type="button" onclick="showMenuDialog()"><i class="fa fa-lock"></i> 菜单赋权</button>
									</div>
									<div class="col-xs-8 text-right">
										<div class="form-inline">
											<div class="form-group">
												<input type="text" placeholder="请输入角色编号或名称" style="width: 200px;" id="keyword" class="form-control" value="">
											</div>
											<button class="btn btn-default searchBtn" type="button" onclick="search()" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
										</div>
									</div>
								</div>

							</div>
							<div class="ibox-content">
								<div class="table-responsive">
									<table class="table table-striped table-bordered table-hover rs" style="width: 99%;">
										<thead>
											<tr>
												<th style="width: 40px;">
												</th>
												<th>角色编码</th>
												<th>角色名称</th>
												<th>操作人</th>
												<th>创建日期</th>
												<!--<th>角色状态</th>-->
												<th>角色类型</th>
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
						<h2 id="addRoleInfo">新增角色信息</h2>
						<h2 id="editRoleInfo">编辑角色信息</h2>
						<input type="hidden" id="type" />
					</div>
					<div class="modal-body">
						<form id="roleInfo">
							<div class="row">
								<div class="col-xs-12">
									<div class="form-group">
										<label>角色编码:</label>
										<input type="text" placeholder="" class="form-control" name="role_id" id="role_id">
									</div>
								</div>
								<div class="col-xs-12">
									<div class="form-group">
										<label>角色名称:</label>
										<input type="text" placeholder="" class="form-control" name="role_name" id="role_name">
									</div>
								</div>
								<!--<div class="col-xs-12">
									<div class="form-group">
										<label>角色状态:</label>
										<select name="role_status" id="role_status" class="form-control">
											<option value="0">正常</option>
					                        <option value="1">锁定</option>
					                        <option value="2">失效</option>
										</select>
									</div>
								</div>-->
								<div class="col-xs-12">
									<div class="form-group">
										<label>角色类型:</label>
										<select name="role_stage" id="role_type" class="form-control">
										</select>
									</div>
								</div>
							</div>
						</form>
					</div>

					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="saveRoleInfo()">保存</button>
					</div>
				</div>
			</div>
		</div>

		<div class="modal inmodal" id="resTree" tabindex="-1" role="dialog" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2>菜单赋权</h2>
					</div>
					<div class="modal-body">
						<ul id="treeDemo" class="ztree"></ul>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="saveMenuRole()">保存</button>
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
				chkboxType: {
					"Y": "ps",
					"N": "s"
				}
			}
		};

		function zTreeOnDblClickSimple(event, treeId, treeNode) {
			//alert(treeNode ? treeNode.tId + ", " + treeNode.name : "isRoot");
			//			$('#userOrgUnitId').val(treeNode.id);
			//			$('#userOrgUnitName').val(treeNode.name);

			//			$('#orgTree').modal('hide');
		};

		function zTreeOnClickSimple(event, treeId, treeNode) {
			var treeObj = $.fn.zTree.getZTreeObj(treeId);
			treeObj.expandNode(treeNode, null, null, null);

			return true;
		};

		var rs;
		$(function() {
			//控制按钮权限
			$.get('/platform/DefaultMenu', {
				TypeId: 2,
				Id: '101050200000',
				r:Math.random()*10000000000000
			}, function(data) {
				var rs = JSON.parse(data);
				$('.res').hide();
				if(rs != null) {
					rs.forEach(function(e) {
						$("button[resid='" + e.Res_id + "']").show();
					});
				}
			})

			initTable();

			$('.rs tbody').on('click', 'tr', function() {
				//				if($(this).find('input:radio').get(0).checked) {
				//					$(this).find('input:radio').get(0).checked = false;
				//				} else {
				//					$(this).find('input:radio').get(0).checked = true;
				//				}
				$(this).find('input:radio').get(0).checked = true;
			});

			//			$('input:radio').click(function(e){
			//				$(this)[0].checked;
			//				//阻止冒泡
			//				e = e || event;
			//				e.stopPropagation ? e.stopPropagation() : e.cancelBubble = true;
			//			});

			$.get('/platform/RoleStage', {r:Math.random()*10000000000000}, function(data) {
				appendOption('role_type', data, 'RoletypeId', 'RoletypeDesc');
			});
		});

		function initTable() {
			rs = $('.rs').DataTable({
				sAjaxSource: '/platform/RoleInfo?r='+Math.random()*10000000000000,
				info: true,
				iDisplayLength: 10,
				aoColumns: [{
						"data": null
					}, {
						"data": "Role_id"
					}, {
						"data": "Role_name"
					}, {
						"data": "Role_owner"
					}, {
						"data": "Role_create_date",
						render: function(data, type, row) {
							return data.substring(0, 10);
						}
					}
					/*, {
										"data": "Role_status_desc"
									}*/
					, {
						"data": "Role_type_desc"
					}
				],
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

		function search() {
//			rs.ajax.reload(null, true);
			rs.ajax.url('/platform/RoleInfo?r='+Math.random()*10000000000000).load();
		}

		function showEditDialog(type) {
			$('#type').val(type);
			if(type === '1') {
				$('#addRoleInfo').show();
				$('#editRoleInfo').hide();

				$('#roleInfo')[0].reset();
				$('#role_id').removeAttr('readonly');
			} else if(type === '2') {
				if($('.rs input:checked').length === 0) {
					toastr.warning('请选择一行角色信息进行编辑');
					return;
				}

				$('#addRoleInfo').hide();
				$('#editRoleInfo').show();
				$('#role_id').attr('readonly', 'readonly');

				var info = rs.data()[$('.rs input:checked').val()];

				$('#role_id').val(info.Role_id);
				$('#role_name').val(info.Role_name);
				//				$('#role_status').val(info.Role_status);
				$('#role_type').val(info.Role_type_id);
			}
			$('#editDialog').modal('show');
		}
		//		
		function saveRoleInfo() {
			//判断是否是数字或字母 
			var id = /^[0-9a-zA-Z\_]+$/;
			//判断是否是汉字、字母、数字组成 
			var name = /^[0-9a-zA-Z\u4e00-\u9fa5]+$/;

			var role_id = $('#role_id').val();
			var role_name = $('#role_name').val();
			if(!id.test(role_id)) {
				toastr.warning('角色编码必须是数字或字母组成');
				return;
			}
			if(!name.test(role_name)) {
				toastr.warning('角色名称必须是数字、汉字或字母组成');
				return;
			}

			if($('#type').val() === '1') { //新增
				$.post('/platform/RoleInfo?r='+Math.random()*10000000000000, $('#roleInfo').serialize(), function(data) {
					var r = JSON.parse(data);
					if(r.ErrorCode === '1') {
						toastr.success(r.ErrorMsg);
//						rs.ajax.reload(null, true);
						rs.ajax.url('/platform/RoleInfo?r='+Math.random()*10000000000000).load();
						$('#editDialog').modal('hide');
					} else {
						toastr.error(r.ErrorMsg);
						return false;
					}

				});
			} else if($('#type').val() === '2') { //编辑
				$.ajax({
					type: "put",
					url: "/platform/RoleInfo?" + $('#roleInfo').serialize()+"&r="+Math.random()*10000000000000,
					async: true,
					error: function(msg) {
						console.log(msg.responseText);
					},
					success: function(data) {
						var r = JSON.parse(data);
						if(r.ErrorCode === '1') {
							toastr.success(r.ErrorMsg);
//							rs.ajax.reload(null, true);
							rs.ajax.url('/platform/RoleInfo?r='+Math.random()*10000000000000).load();
							$('#editDialog').modal('hide');
						} else {
							toastr.error(r.ErrorMsg);
							return;
						}

					}
				});
			}

		}
		//		
		function deleteRoleInfo() {
			if($('.rs input:checked').length === 0) {
				toastr.warning('请选择一行角色信息进行删除');
				return;
			}

			var info = rs.data()[$('.rs input:checked').val()];
			swal({
				title: "删除！",
				text: "是否删除选中的角色信息？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				$.ajax({
					type: "delete",
					url: "/platform/RoleInfo?role_id=" + info.Role_id+"&r="+Math.random()*10000000000000,
					error: function(msg) {
						console.log(msg.responseText);
					},
					success: function(data) {
						var r = JSON.parse(data);
						if(r.ErrorCode === '1') {
							//		                    toastr.success(rs.ErrorMsg);
							swal("删除!", rs.ErrorMsg, "success");
//							rs.ajax.reload(null, true);
							rs.ajax.url('/platform/RoleInfo?r='+Math.random()*10000000000000).load();
						} else {
							toastr.error(r.ErrorMsg);
							return;
						}
					}
				});
			});
		}

		function showMenuDialog() {
			if($('.rs input:checked').length === 0) {
				toastr.warning('请选择一行角色信息进行赋权');
				return;
			}

			var info = rs.data()[$('.rs input:checked').val()];
			$.get('/platform/ResTree' + '?r=' + Math.random(100000000000000) + "&RoleTypeId=" + info.Role_type_id, function(data) {
				$.fn.zTree.init($("#treeDemo"), s, JSON.parse(data)); //树

				var info = rs.data()[$('.rs input:checked').val()];
				//获取角色已赋予的资源
				$.get('/platform/RoleResource', {
					RoleId: info.Role_id,
					r:Math.random()*10000000000000
				}, function(data) {
					var rs = JSON.parse(data);
					var treeObj = $.fn.zTree.getZTreeObj("treeDemo");
					if(rs != null) {
						rs.forEach(function(e) {
							var node = treeObj.getNodeByParam('id', e.Role_id, null);
							if(node) {
								treeObj.checkNode(node, true, false);
							}
						});
					}

					var nodes = treeObj.getCheckedNodes(true);
					nodes.forEach(function(e) {
						treeObj.expandNode(e, true, false, false);
					});
				});
			});

			$('#resTree').modal('show');
		}

		function saveMenuRole() {
			var treeObj = $.fn.zTree.getZTreeObj("treeDemo");
			var nodes = treeObj.getCheckedNodes(true);
			var json = [];
			nodes.forEach(function(e) {
				json.push(e.id);
			})
			var info = rs.data()[$('.rs input:checked').val()];
			$.ajax({
				type: "put",
				url: "/platform/RoleResource?JSON=" + JSON.stringify(json) + "&RoleId=" + info.Role_id+"&r="+Math.random()*10000000000000,
				async: true,
				error: function(msg) {
					console.log(msg.responseText);
				},
				success: function(data) {
					var r = JSON.parse(data);
					if(r.ErrorCode === '1') {
						toastr.success('菜单赋权成功！');

						$('#resTree').modal('hide');
					} else {
						toastr.error(r.ErrorMsg);
						return;
					}
				}
			});
		}
	</script>

</html>