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

		<title>用户管理</title>

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
									<div class="col-xs-8">
										<button resid="101050101000" class="btn btn-primary res" type="button" onclick="showEditDialog('1')"> <i class="fa fa-plus"></i> 新增</button>
										<button resid="101050102000" class="btn btn-info editBtn res" type="button" onclick="showEditDialog('2')"><i class="fa fa-paste"></i> 编辑</button>
										<button resid="101050103000" class="btn btn-danger deleteBtn res" type="button" onclick="deleteUserInfo()"><i class="fa fa-times"></i> 删除</button>
										<button resid="101050104000" class="btn btn-warning res" type="button" onclick="resetPass()"><i class="fa fa-lock"></i> 重置密码</button>
										<button resid="101050105000" class="btn btn-info res" type="button" onclick="showRoleDialog()"><i class="fa fa-lock"></i> 角色赋权</button>
									</div>
									<div class="col-xs-4 text-right">
										<div class="form-inline">
											<div class="form-group">
												<input type="text" placeholder="请输入用户编号或名称" style="width: 200px;" id="keyword" class="form-control" value="">
											</div>
											<button class="btn btn-default searchBtn" type="button" onclick="search()" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
										</div>
									</div>
								</div>

							</div>
							<div class="ibox-content">
								<div class="table-responsive">
									<table class="table table-striped table-bordered table-hover users"  style="width: 99%;">
										<thead>
											<tr>
												<th style="width: 40px;">
												</th>
												<th>帐号</th>
												<th>用户名</th>
												<th>组织代码</th>
												<th>状态</th>
												<th>邮箱</th>
												<th>手机号</th>
												<th>创建时间</th>
												<th>创建人</th>
												<!--<th>是否在线</th>-->
												<th>用户角色</th>
												<th>所属域</th>
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
						<h2 id="addUserInfo">新增用户信息</h2>
						<h2 id="editUserInfo">编辑用户信息</h2>
						<input type="hidden" id="type"/>
					</div>
					<div class="modal-body">
						<form id="userInfo">
							<div class="row">
								<div class="col-xs-6">
										<div class="form-group">
											<label>帐号:</label>
											<input type="text" placeholder="" class="form-control" name="userId" id="userId">
										</div>
								</div>
								<div class="col-xs-6">
										<div class="form-group">
											<label>用户名:</label>
											<input type="text" placeholder="" class="form-control" name="userDesc" id="userDesc">
										</div>
								</div>
								<div class="col-xs-6">
										<div class="form-group">
											<label>邮箱:</label>
											<input type="text" placeholder="" class="form-control" name="userEmail" id="userEmail">
										</div>
								</div>
								<div class="col-xs-6">
										<div class="form-group">
											<label>手机号:</label>
											<input type="text" placeholder="" class="form-control" name="userPhone" id="userPhone">
										</div>
								</div>
								
								<div class="col-xs-6">
									<div class="form-group">
										<label>机构:</label>
										<div class="input-group">
											<input type="hidden" id="userOrgUnitId" name="userOrgUnitId" />
											<input type="text" class="form-control" id="userOrgUnitName" name="userOrgUnitName" readonly="">
											<span class="input-group-btn" onclick="showOrgDialog()">
												<button type="button" class="btn"><i class="fa fa-level-up"></i></button> 
											</span>
										</div>
									</div>
								</div>
								
								<div class="col-xs-6">
									<div class="form-group">
										<label>状态:</label>
										<select name="userStatus" id="userStatus" class="form-control">
											<option value="0">正常</option>
					                        <!--<option value="1">锁定</option>-->
					                        <option value="2">失效</option>
										</select>
									</div>
								</div>
								
								<div class="col-xs-6">
									<div class="form-group">
										<label>所属域:</label>
										<select name="domainId" id="domainId" class="form-control">
										</select>
									</div>
								</div>
							</div>
						</form>
					</div>

					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="saveUserInfo()">保存</button>
					</div>
				</div>
			</div>
		</div>
		
		<div class="modal inmodal" id="roleDialog" tabindex="-1" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2>用户角色赋权</h2>
					</div>
					<div class="modal-body"  style="max-height: 500px;overflow: auto;">
						<div class="row">
							<table class="table table-bordered roles">
								<thead>
									<tr>
										<th></th>
										<th>角色编码</th>
										<th>角色名称</th>
										<th>角色类型</th>
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
						<button type="button" class="btn btn-primary" onclick="saveUserRoles()">保存</button>
					</div>
				</div>
			</div>
		</div>
		
		<div class="modal inmodal" id="orgTree" tabindex="-1" role="dialog" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" onclick="$('#orgTree').hide()"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 >机构选择</h2>
					</div>
					<div class="modal-body">
						<ul id="treeDemo" class="ztree"></ul>
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
			}
		};
		function zTreeOnDblClickSimple(event, treeId, treeNode) {
			//alert(treeNode ? treeNode.tId + ", " + treeNode.name : "isRoot");
			$('#userOrgUnitId').val(treeNode.id);
			$('#userOrgUnitName').val(treeNode.name);

//			$('#orgTree').modal('hide');
			$('#orgTree').hide();
		};

		function zTreeOnClickSimple(event, treeId, treeNode) {
			var treeObj = $.fn.zTree.getZTreeObj(treeId);
			treeObj.expandNode(treeNode, null, null, null);

			return true;
		};
		
		var users; 
		$(function(){
			//控制按钮权限
			$.get('/platform/DefaultMenu',{TypeId:2,Id:'101050100000',r:Math.random()*10000000000000}, function(data){
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
			
//			$('.users').height(document.body.clientHeight-200);
			
			$('.users tbody').on('click', 'tr', function() {
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
			
			$.get('/platform/AllRole', {r:Math.random()*10000000000000}, function(data){
				var rs=JSON.parse(data);
//				<tr>
//					<td style='width: 40px;text-align:center;'><label class='checkbox-inline i-checks' style='width: 30px;'> <input type='checkbox'></label></td>
//					<td>aa</td>
//					<td>aaa</td>
//				</tr>
				if(rs!=null){
					$('.roles tbody').html('');
					rs.forEach(function(e, index){
						var temp="<tr><td style='width: 40px;text-align:center;'><label 	typeid='"+e.RoletypeId+"' roleid='"+e.RoleId+"' rolename='"+e.RoleName+"' class='checkbox-inline i-checks' style='width: 30px;'> <input type='checkbox'></label></td>"
								+"<td>"+e.RoleId+"</td><td>"+e.RoleName+"</td>"+"<td>"+e.RoletypeDesc+"</td>";
						$('.roles tbody').append(temp);
					});
				}
				iCheckInit();
			});
			
			$.get('/platform/SysDomainInfo', {r:Math.random()*10000000000000}, function(data){
				appendOption('domainId', data, 'DomainId', 'DomainName');
			});
		});
		
		function iCheckInit(){
			$('.i-checks').iCheck({
				checkboxClass: 'icheckbox_square-green',
				radioClass: 'iradio_square-green',
			});
		}
		
		function initTable() {
			users = $('.users').DataTable({
				sAjaxSource: '/platform/UserInfo?r='+Math.random()*10000000000000,
				info: true,
				iDisplayLength: 15,
				aoColumns: [{
					"data": null
				}, {
					"data": "User_id"
				}, {
					"data": "User_name"
				}, {
					"data": "Org_unit_desc"
				}, {
					"data": "User_status_desc"
				}, {
					"data": "User_email"
				}, {
					"data": "User_phone"
				}, {
					"data": "User_create_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "User_owner"
				}/*, {
					"data": "User_online_flag"
				}*/, {
					"data": "User_role",
					render: function(data, type, row) {
						if(data===''){
							return "<span style='color:red;'>待分配角色</span>";
						}
						return data;
					}
				}, {
					"data": "Domain_Desc"
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
				$('#addUserInfo').show();
				$('#editUserInfo').hide();
				
				$('#userInfo')[0].reset();
				$('#userId').removeAttr('readonly');
			}else if(type==='2'){
				if($('.users input:checked').length===0){
					toastr.warning('请选择一行用户信息进行编辑');
					return;
				}
				
				$('#addUserInfo').hide();
				$('#editUserInfo').show();
				$('#userId').attr('readonly', 'readonly');
				
				var info = users.data()[$('.users input:checked').val()];
				
				$('#userId').val(info.User_id);
				$('#userDesc').val(info.User_name);
				$('#userEmail').val(info.User_email);
				$('#userPhone').val(info.User_phone);
				
				$('#userOrgUnitId').val(info.Org_unit_id);
				$('#userOrgUnitName').val(info.Org_unit_desc);
				$('#domainId').val(info.Domain_id);
				if(info.User_status_desc=='正常'){
					$('#userStatus').val('0');
				}else if(info.User_status_desc=='锁定'){
					$('#userStatus').val('1');
				}else if(info.User_status_desc=='失效'){
					$('#userStatus').val('2');
				}
			}
			$('#editDialog').modal('show');
		}
		
		function saveUserInfo(){
			//输入校验
			//判断是否是数字或字母 
    			var id = /^[0-9a-zA-Z\_]+$/;
    			//判断是否是汉字、字母、数字组成 
    			var name = /^[0-9a-zA-Z\u4e00-\u9fa5]+$/;
    			//手机
    			var phone = /^[1][0-9]{10}$/;
    			//邮箱
    			var email = /^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$/;
    			
    			var userId=$('#userId').val();
    			var userDesc=$('#userDesc').val();
    			var userEmail=$('#userEmail').val();
    			var userPhone=$('#userPhone').val();
    			var userOrgUnitName=$('#userOrgUnitName').val();
    			
    			if(!id.test(userId)){
    				toastr.warning('帐号必须是数字或字母组成');
    				return;
    			}
    			if(!name.test(userDesc)){
    				toastr.warning('用户名必须是数字、汉字或字母组成');
    				return;
    			}
    			if(!phone.test(userPhone) && userPhone!=''){
    				toastr.warning('手机号码错误');
    				return;
    			}
    			if(!email.test(userEmail) && userEmail!=''){
    				toastr.warning('邮箱格式错误');
    				return;
    			}
    			if(userOrgUnitName==='' || userOrgUnitName===null){
    				toastr.warning('请选择机构');
    				return;
    			}
			
			if($('#type').val()==='1'){ //新增
				$.post('/platform/UserInfo?r='+Math.random()*10000000000000, $('#userInfo').serialize(),function(data){
					var rs=JSON.parse(data);
		            if(rs.ErrorCode==='1'){
						toastr.success(rs.ErrorMsg+", 用户初始密码为123456");
//						users.ajax.reload(null, true);
						users.ajax.url('/platform/UserInfo?r='+Math.random()*10000000000000).load();
						$('#editDialog').modal('hide');
		            }else {
		            		toastr.error(rs.ErrorMsg);
		                return false;
		            }
				});
			}else if($('#type').val()==='2'){ //编辑
				$.ajax({
					type:"put",
					url:"/platform/UserInfo?"+$('#userInfo').serialize()+"&r="+Math.random()*10000000000000,
					async:true,
					error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
//		                    toastr.success(rs.ErrorMsg);
							toastr.success(rs.ErrorMsg);
//							users.ajax.reload(null, true);
							users.ajax.url('/platform/UserInfo?r='+Math.random()*10000000000000).load();
							$('#editDialog').modal('hide');
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
						
		            }
				});
			}
		}
		
		function deleteUserInfo(){
			if($('.users input:checked').length===0){
				toastr.warning('请选择一行用户信息进行删除');
				return;
			}
			
			var info = users.data()[$('.users input:checked').val()];
			swal({
				title: "删除！",
				text: "是否删除选中的用户信息？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				var json=[{'User_id':info.User_id}];
				
				$.ajax({
		            type:"delete",
		            url:"/platform/UserInfo?JSON="+JSON.stringify(json)+"&r="+Math.random()*10000000000000,
		            error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
		                    swal("删除!", rs.ErrorMsg, "success");
//							users.ajax.reload(null, true);
							users.ajax.url('/platform/UserInfo?r='+Math.random()*10000000000000).load();
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
		            }
		        });
			});
		}
		
		function search(){
//			users.ajax.reload(null, true);
			users.ajax.url('/platform/UserInfo?r='+Math.random()*10000000000000).load();
		}
		
		function resetPass(){
			if($('.users input:checked').length===0){
				toastr.warning('请选择一行用户信息进行密码重置');
				return;
			}
			
			swal({
				title: "提示！",
				text: "是否重置该用户密码？初始密码为123456",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				var info = users.data()[$('.users input:checked').val()];
				
				$.ajax({
		            type:"post",
		            url:"/platform/ResetPassword?userId="+info.User_id+"&r="+Math.random()*10000000000000,
		            error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
		            		if(JSON.parse(data).ErrorCode==='1'){
		            			swal("提示!", "密码重置完成", "success");
		            		}else {
		            			toastr.error(JSON.parse(data).ErrorMsg);
		            		}
		            }
		        });
			});
		}
		
		function showRoleDialog(){
			if($('.users input:checked').length===0){
				toastr.warning('请选择一行用户信息进行角色赋权');
				return;
			}
			
			var info = users.data()[$('.users input:checked').val()];
			$.get('/platform/UserRole',{UserId:info.User_id,r:Math.random()*10000000000000},function(data){
				var rs=JSON.parse(data);
//				$('.icheckbox_square-green').removeClass('checked');
				$('.icheckbox_square-green').iCheck('uncheck');
				if(rs!=null){
					rs.forEach(function(e){
						$("label[roleid="+e.RoleId+"]").find('.icheckbox_square-green').iCheck('check');
					});
				}
			});
			
			$('#roleDialog').modal('show');
		}
		
		function saveUserRoles(){
			if($('.checked').length===0){
				toastr.warning('请至少选择一个角色赋予用户');
				return;
			}
			
			var UserId = users.data()[$('.users input:checked').val()].User_id;
			var json=[];
			var roleType=[];
			$('.checked').each(function(){
				json.push($(this).parent().attr('roleid'));
				roleType.push($(this).parent().attr('typeid'));
			});
			
			//检查选择的角色是否是相同类型的
			var temp=roleType[0];
			var flag=0;
			roleType.forEach(function(e){
				if(e==temp){
					flag+=1;
				}
			});
			
			if(flag!=roleType.length){
				toastr.warning('只能给用户赋予多个相同类型的角色！');
				return;
			}
			
			$.ajax({
	            type:"put",
	            url:"/platform/UserRole?UserId="+UserId+"&JSON="+JSON.stringify(json)+"&r="+Math.random()*10000000000000,
	            error: function(msg){
	                console.log(msg.responseText);
	            },
	            success: function(data){
//					toastr.success('角色赋权完成');
////					users.ajax.reload(null, true);
//					cvs.ajax.url('/platform/UserInfo?r='+Math.random()*10000000000000).load();
//					$('#roleDialog').modal('hide');
					if(JSON.parse(data).ErrorCode==='1'){
	            			toastr.success('角色赋权完成');
						users.ajax.url('/platform/UserInfo?r='+Math.random()*10000000000000).load();
						$('#roleDialog').modal('hide');
	            		}else {
	            			toastr.error(JSON.parse(data).ErrorMsg);
	            		}
	            }
	        });
		}
		
		function showOrgDialog(){
			$.get('/platform/OrgTree' + '?r=' + Math.random(100000000000000), function(data) {
				$.fn.zTree.init($("#treeDemo"), s, JSON.parse(data)); //树
			});
			
			$('#orgTree').show();
		}
	</script>

</html>