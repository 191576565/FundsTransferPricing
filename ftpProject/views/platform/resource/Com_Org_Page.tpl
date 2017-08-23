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

		<title>机构定义</title>

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
		<link rel="stylesheet" href="/static/css/jquery.treegrid.css">

		<script src="/static/inspinia/js/jquery-1.12.3.min.js"></script>
		<script src="/static/inspinia/js/bootstrap.min.js"></script>
		<script src="/static/inspinia/js/jquery.placeholder.min.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/plugins/metisMenu/jquery.metisMenu.js"></script>
		<script src="/static/inspinia/js/plugins/slimscroll/jquery.slimscroll.min.js"></script>
		<script src="/static/inspinia/js/plugins/dataTables/datatables.min.js"></script>
		<script type="text/javascript" src="/static/js/jquery.treegrid.min.js"></script>
        <script type="text/javascript" src="/static/js/jquery.treegrid.bootstrap3.js"></script>

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
									<div class="col-xs-12">
										<button resid="101040301000" class="btn btn-primary res" type="button" onclick="showEditDialog('1')"> <i class="fa fa-plus"></i> 新增</button>
										<button resid="101040302000" class="btn btn-info editBtn res" type="button" onclick="showEditDialog('2')"><i class="fa fa-paste"></i> 编辑</button>
										<!--<button resid="101040303000" class="btn btn-danger deleteBtn res" type="button" onclick="deleteOrgInfo()"><i class="fa fa-times"></i> 删除</button>-->
									</div>
									<!--<div class="col-xs-4 text-right">
										<div class="form-inline">
											<div class="form-group">
												<input type="text" placeholder="请输入机构编号或机构名称" style="width: 200px;" id="keyword" class="form-control" value="">
											</div>
											<button class="btn btn-default searchBtn" type="button" onclick="search()" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
										</div>
									</div>-->
								</div>

							</div>
							<div class="ibox-content">
								<div class="table-responsive">
									<table class="table table-striped table-bordered table-hover orgs"  style="width: 99%;">
										<thead>
											<tr>
												<td style='width:40px;'></td>
												<th>机构编码</th>
												<th>机构名称</th>
												<th>上级机构编码</th>
												<th>上级机构名称</th>
												<th>备注</th>
											</tr>
										</thead>
										<tbody>
						                  
						                </tbody>
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
						<h2 id="addOrgInfo">新增机构信息</h2>
						<h2 id="editOrgInfo">编辑机构信息</h2>
						<input type="hidden" id="type"/>
					</div>
					<div class="modal-body">
						<form id="orgInfo">
							<div class="row">
								<div class="col-xs-6">
										<div class="form-group">
											<label>机构编码:</label>
											<input type="text" placeholder="" class="form-control" name="OrgUintID" id="OrgUintID">
										</div>
								</div>
								<div class="col-xs-6">
										<div class="form-group">
											<label>机构名称:</label>
											<input type="text" placeholder="" class="form-control" name="OrgUintDesc" id="OrgUintDesc">
										</div>
								</div>
								<div class="col-xs-6">
									<div class="form-group">
										<label>上级机构:</label>
										<div class="input-group">
											<input type="hidden" id="OrgUpID" name="OrgUpID" />
											<input type="text" class="form-control" id="OrgUpDesc" name="OrgUpDesc" readonly="">
											<span class="input-group-btn" onclick="showOrgDialog()">
												<button type="button" class="btn"><i class="fa fa-level-up"></i></button> 
											</span>
										</div>
									</div>
								</div>
								<div class="col-xs-6">
										<div class="form-group">
											<label>备注:</label>
											<input type="text" placeholder="" class="form-control" name="MemoBk" id="MemoBk">
										</div>
								</div>
							</div>
						</form>
					</div>

					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="saveOrgInfo()">保存</button>
					</div>
				</div>
			</div>
		</div>
		
		<div class="modal inmodal" id="orgTree" tabindex="-2" role="dialog" data-backdrop="false">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" onclick="$('#orgTree').hide()"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 >选择上级机构</h2>
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
			if(treeNode.id===$('#OrgUintID').val()){
				toastr.warning('机构编码与上级机构编码冲突');
				return;
			}
			$('#OrgUpID').val(treeNode.id);
			$('#OrgUpDesc').val(treeNode.name);

//			$('#orgTree').modal('hide');
			$('#orgTree').hide();
		};

		function zTreeOnClickSimple(event, treeId, treeNode) {
			var treeObj = $.fn.zTree.getZTreeObj(treeId);
			treeObj.expandNode(treeNode, null, null, null);

			return true;
		};
	
		var orgs; 
		$(function(){
			//控制按钮权限
			$.get('/platform/DefaultMenu',{TypeId:2,Id:'101040300000',r:Math.random()*10000000000000}, function(data){
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
			});
			
			initTable();
			
//			$('.orgs').height(document.body.clientHeight-200);
			
//			$.get('/platform/SysDomainInfo', function(data){
//				appendOption('domainId', data, 'DomainId', 'DomainName');
//			});
		});
		
		function initTable() {
			$.get('/platform/ComOrg', {r:Math.random()*10000000000000}, function(data){
				var rs=JSON.parse(data);
				$('.orgs tbody').html('');
				var temp='';
				if(rs!=null){
					rs.forEach(function(e){
						if(e.Level==='1'){
							temp='<tr class="treegrid-'+e.OrgUintID+'" level="'+e.Level+'"><td style="width:40px;text-align:center;"><input type="radio" name="orgRadio"/></td><td id="'+e.OrgUintID+'">'+e.OrgUintID+'</td><td>'+e.OrgUintDesc+'</td><td>'
									+e.OrgUpID+'</td><td>'+e.OrgUpDesc+'</td><td>'+e.MemoBk+'</td></tr>';
							$('.orgs tbody').append(temp);
						}else {
							temp='<tr class="treegrid-'+e.OrgUintID+' treegrid-parent-'+e.OrgUpID+'" level="'+e.Level+'"><td style="width:40px;text-align:center;"><input type="radio" name="orgRadio"/></td><td id="'+e.OrgUintID+'">'+e.OrgUintID+'</td><td>'+e.OrgUintDesc+'</td><td>'
									+e.OrgUpID+'</td><td>'+e.OrgUpDesc+'</td><td>'+e.MemoBk+'</td></tr>';
							$('.orgs tbody tr.treegrid-'+e.OrgUpID).after(temp);
//							if($(".orgs tbody tr[level="+e.Level+"]").length===0){
//								$('.orgs tbody tr.treegrid-'+e.OrgUpID).after(temp);
//							}else {
//								$('.orgs tbody tr.treegrid-'+e.OrgUpID).
//							}
						}
					});
				}
				
//				$('.orgs tbody tr').prepend("<td style='width:40px;text-align:center;'><input type='radio' name='orgRadio'/></td>");
				$('.orgs').treegrid({
					treeColumn:1,
	                expanderExpandedClass: 'glyphicon glyphicon-minus',
	                expanderCollapsedClass: 'glyphicon glyphicon-plus'
	            });
				//收起第二级及一下的机构
            		$('.orgs tbody tr[level=2]').treegrid('collapseRecursive');
            		
            		//tr点击单选钮选中
	            $('.orgs tbody tr').click(function(){
	            		$(this).find('input:radio')[0].checked=true;
	            });
			});
		}
		
		function showEditDialog(type){ 
			$('#type').val(type);
			if(type==='1'){
				$('#addOrgInfo').show();
				$('#editOrgInfo').hide();
				
				$('#orgInfo')[0].reset();
				$('#OrgUintID').removeAttr('readonly');
				
				if($('.orgs input:checked').length!=0){
					var tds=$('.orgs input:checked').parent().nextAll();
				
					$('#OrgUpID').val($(tds[0]).attr('id'));
					$('#OrgUpDesc').val($(tds[1]).text());
				}
			}else if(type==='2'){
				if($('.orgs input:checked').length===0){
					toastr.warning('请选择一行机构信息进行编辑');
					return;
				}
				
				$('#addOrgInfo').hide();
				$('#editOrgInfo').show();
				$('#OrgUintID').attr('readonly', 'readonly');
				
//				var info = orgs.data()[$('.orgs input:checked').val()];
				var tds=$('.orgs input:checked').parent().nextAll();
				
				$('#OrgUintID').val($(tds[0]).attr('id'));
				$('#OrgUintDesc').val($(tds[1]).text());
				$('#OrgUpID').val($(tds[2]).text());
				$('#OrgUpDesc').val($(tds[3]).text());
				$('#MemoBk').val($(tds[4]).text());
				
			}
			$('#editDialog').modal('show');
		}
		
		function saveOrgInfo(){
//			//输入校验
//			//判断是否是字母 
    			var id = /^[a-zA-Z0-9]+$/;
//  			//判断是否是汉字、字母、数字组成 
//  			var name = /^[0-9a-zA-Z\u4e00-\u9fa5]+$/;
//  			
//  			var regu = /^[0-9]+$/;
    			
    			var OrgUintID=$('#OrgUintID').val();
    			var OrgUintDesc=$('#OrgUintDesc').val();
    			var OrgUpID=$('#OrgUpID').val();
    			
    			if(!id.test(OrgUintID)){
    				toastr.warning('请输入由字母或数字组成的机构编码');
    				return;
    			}
    			if(OrgUintDesc===''){
    				toastr.warning('请输入机构名称');
    				return;
    			}
    			if(OrgUpID===''){
    				toastr.warning('请选择上级机构');
    				return;
    			}
			
			if($('#type').val()==='1'){ //新增
				$.post('/platform/ComOrg?r='+Math.random()*10000000000000, $('#orgInfo').serialize(),function(data){
					var rs=JSON.parse(data);
		            if(rs.ErrorCode==='1'){
		                toastr.success('机构信息保存成功！');
						initTable();
//						var info=$('#orgInfo').serializeArray();
						//不刷新表格直接插入tr
//						var upId=$('#OrgUpID').val();
//						var level=$('.orgs tbody tr.treegrid-'+info[2].value).attr('Level');
//						var temp='<tr class="treegrid-'+info[0].value+' treegrid-parent-'+info[2].value+'"><td style="width:40px;text-align:center;"><input type="radio" name="orgRadio"/></td><td id="'+info[0].value+'">'
//						for(var i=0; i<(parseInt(level)+1); i++){
//							temp+='<span class="treegrid-indent"></span>';
//						}
//						temp=temp+info[0].value+'</td><td>'+info[1].value+'</td><td>'+info[2].value+'</td><td>'+info[3].value+'</td><td>'+info[4].value+'</td></tr>';
//						$('.orgs tbody tr.treegrid-'+info[2].value).after(temp);
////						$('.treegrid-'+info[0].value).prepend("<td style='width:40px;text-align:center;'><input type='radio' name='orgRadio'/></td>");
////						$('.orgs').treegrid({
////							treeColumn:1,
////				            expanderExpandedClass: 'glyphicon glyphicon-minus',
////				            expanderCollapsedClass: 'glyphicon glyphicon-plus'
////				        });
////						$('.treegrid-'+info[2].value).treegrid('render');
//						
//						//tr点击单选钮选中
//			            $('.orgs tbody tr').click(function(){
//			            		$(this).find('input:radio')[0].checked=true;
//			            });
						$('#editDialog').modal('hide');
		            }else {
		            		toastr.error(rs.ErrorMsg);
		                return false;
		            }
					
				});
			}else if($('#type').val()==='2'){ //编辑
				$.ajax({
					type:"put",
					url:"/platform/ComOrg?"+$('#orgInfo').serialize()+"&r="+Math.random()*10000000000000,
					async:true,
					error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
//		                    toastr.success(rs.ErrorMsg);
							toastr.success('机构信息编辑成功！');
							initTable();
							$('#editDialog').modal('hide');
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
						
		            }
				});
			}
			
		}
		
		function showOrgDialog(){
			$.get('/platform/OrgTree' + '?r=' + Math.random(100000000000000), function(data) {
				$.fn.zTree.init($("#treeDemo"), s, JSON.parse(data)); //树
			});
			
			$('#orgTree').show();
		}
		
		function deleteOrgInfo(){
			if($('.orgs input:checked').length===0){
				toastr.warning('请选择一行机构信息进行删除');
				return;
			}
			
//			var info = currs.data()[$('.currs input:checked').val()];
			var OrgUintID=$('.orgs input:checked').parent().next().attr('id');
			swal({
				title: "删除！",
				text: "是否删除选中的机构及其子机构信息？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				$.ajax({
		            type:"delete",
		            url:"/platform/ComOrg?OrgUintID="+OrgUintID+"&r="+Math.random()*10000000000000,
		            error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
//		                    toastr.success(rs.ErrorMsg);
		                    swal("删除!", "机构信息已删除", "success");
		                    initTable();
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