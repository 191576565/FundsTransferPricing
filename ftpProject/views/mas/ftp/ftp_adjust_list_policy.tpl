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

		<title>政策性调节项配置</title>

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
		
		<!--<script src="/static/inspinia/js/ajaxfileupload.js"></script>-->
		<script src="/static/inspinia/js/jquery.form.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/bootstrap-filestyle.min.js"></script>
		<script src="/static/inspinia/jquery.blockUI.js" type="text/javascript" charset="utf-8"></script>

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
			
			.selected {
				background-color: #b0bed9;
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
									<div class="" style="float: left;">
										<button resid="30208020301000" class="btn btn-primary res " type="button" onclick="showEditDialog('add')"> <i class="fa fa-plus"></i> 新增</button>
										<button resid="30208020302000" class="btn btn-info editBtn res" type="button" onclick="showEditDialog('edit')"><i class="fa fa-paste"></i> 编辑</button>
										<button resid="30208020303000" class="btn btn-danger res" type="button" onclick="deletePolicy()"><i class="fa fa-times"></i> 删除</button>
										<button resid="30208020304000" class="btn btn-warning res" type="button" onclick="showCheckDialog()"><i class="fa fa-legal"></i> 数据校验</button>
									</div>
									<div class="" style="float: right;">
										<div class="form-inline">
											<div class="form-group">
												<select class="form-control" id="keyword1" style="width: 200px;">
													<option value="">------请选择调节项类型------</option>
												</select>
											</div>
											<div class="form-group">
												<input type="text" placeholder="请输入调节项维度" style="width: 200px;" id="keyword2" class="form-control" value="">
											</div>
											<button class="btn btn-default searchBtn" type="button" onclick="search()" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
											<button resid="30208020305000"  class="btn btn-default searchBtn res"  onclick="exportData()" style="margin-bottom: 0px;"><i class="fa fa-download"></i>导出</button>
											<!--<button resid="208020306000" class="btn btn-default searchBtn res" type="button" onclick="showUploadDialog()" style="margin-bottom: 0px;"><i class="fa fa-upload"></i>导入</button>-->
										</div>
									</div>
								</div>

							</div>
							<div class="ibox-content">
								<div class="table-responsive">
									<table class="table table-striped table-bordered table-hover policys" style="width: 99%;">
										<thead>
											<tr>
												<th style="width: 40px;">
												</th>
												<th>序号</th>
												<th>调节项类型</th>
												<th>机构</th>
												<th>币种</th>
												<th>调整项维度</th>
												<th>区间值-起</th>
												<th>区间值-止</th>
												<th>最新值日期</th>
												<th>生效起始日</th>
												<th>生效终止日</th>
												<th>存量业务生效起始日</th>
												<th>存量业务生效终止日</th>
												<th>调节点差</th>
												<!--<th>所属域</th>-->
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

		<div class="modal inmodal" id="editDialog" tabindex="-1" role="dialog" aria-hidden="false">
			<div class="modal-dialog modal-lg">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 id="add">新增政策性调节项信息</h2>
						<h2 id="edit">编辑政策性调节项信息</h2>
						<input type="hidden" id="type" value="add" />
					</div>
					<div class="modal-body">
						<form role="form" id="policyForm">
							<input type="hidden" id="uuid" name="Uuid" value=""/>
							<div class="row">
								<div class="form-group col-sm-4">
									<label>调节项类型:</label>
									<select class="form-control" name="Adj_id" id="adjId" onchange="changeInput(this)">
									</select>
								</div>
								<div class="form-group col-sm-4">
									<label>机构:</label>
									<div class="input-group">
										<input type="hidden" id="OrgID" name="Org_unit_id" />
										<input type="text" class="form-control" id="OrgDesc" name="OrgUpDesc" readonly="">
										<span class="input-group-btn orgBtn" onclick="showOrgDialog()">
											<button type="button" class="btn orgBtn" id="orgBtn"><i class="fa fa-level-up"></i></button> 
										</span>
									</div>
								</div>
								<div class="form-group col-sm-4">
									<label>币种:</label>
									<select class="form-control" name="Iso_currency_cd" id="curreny">
									</select>
								</div>
								<!--<div class="form-group col-sm-4">
										<label>调节项维度:</label>
										<input type="text" placeholder="" class="form-control" name="Curve_desc" id="curveName">
									</div>-->
							</div>
							<div class="row">
								
								<!-- 维度 -->
								<div class="form-group col-sm-4 dimension">
									<div>
										<label>资产负债标识:</label>
										<select class="form-control" name="Adj_dyn_dim" id="identify" data-type="">
											<!--<option value=""></option>-->
										</select>
									</div>
								</div>
								<div class="form-group col-sm-4 dimension">
									<label>行业:</label>
									<!--<input type="text" placeholder="" class="form-control" name="Adj_dyn_dim" id="industry">-->
									<div class="input-group">
										<input type="hidden" id="indId" name="Adj_dyn_dim" />
										<input type="text" class="form-control" id="indDesc" name="" readonly="">
										<span class="input-group-btn" onclick="showIndDialog()">
											<button type="button" class="btn"><i class="fa fa-level-up"></i></button> 
										</span>
									</div>
								</div>
								<div class="col-sm-4 form-group dimension">
									<div>
										<label>业务单元:</label>
										<select class="form-control" name="Adj_dyn_dim" id="busizId">
											<!--<option value=""></option>-->
										</select>
									</div>
								</div>
								
								<div class="form-group col-sm-4 dimension">
									<label>产品:</label>
									<div class="input-group">
										<input type="hidden" id="proId" name="Adj_dyn_dim" />
										<input type="text" class="form-control" id="proDesc" name="" readonly="">
										<span class="input-group-btn" onclick="showProDialog()">
											<button type="button" class="btn"><i class="fa fa-level-up"></i></button> 
										</span>
									</div>
								</div>
								<div class="form-group col-sm-4 dimension">
									<div>
										<label>账号:</label>
										<input type="text" placeholder="" class="form-control" name="Adj_dyn_dim" id="accounts">
									</div>
								</div>
								<div class="form-group col-sm-4 dimension">
									<div>
										<label>起始期限(月):</label>
										<input type="text" placeholder="" class="form-control" name="Term_str" id="startTerm">
									</div>
								</div>
								<div class="form-group col-sm-4 dimension">
									<div>
										<label>截止期限(月):</label>
										<input type="text" placeholder="" class="form-control" name="Term_end" id="endTerm">
									</div>
								</div>
								<div class="form-group col-sm-4 dimension">
									<div>
										<label>起始规模(元):</label>
										<input type="text" placeholder="" class="form-control" name="Term_str" id="startScale">
									</div>
								</div>
								<div class="form-group col-sm-4 dimension">
									<div>
										<label>截止规模(元):</label>
										<input type="text" placeholder="" class="form-control" name="Term_end" id="endScale">
									</div>
								</div>
								
								
								<!-- 共有 -->
								<div class="form-group col-sm-4">
									<label>调节点差:</label>
									<input type="text" placeholder="" class="form-control" name="Adj_bp" id="spot">
								</div>
								<div class="form-group col-sm-4">
									<label>生效起始日:</label>
									<div class="input-group date full-width">
										<input type="text" class="form-control input-group-addon" id="startEffectDay" name="Eff_str_date">
									</div>
								</div>
								<div class="form-group col-sm-4">
									<label>生效终止日:</label>
									<div class="input-group date full-width">
										<input type="text" class="form-control input-group-addon" id="endEffectDay" name="Eff_end_date">
									</div>
								</div>
								<div class="form-group col-sm-4">
									<label>存量业务生效起始日:</label>
									<div class="input-group date full-width">
										<input type="text" class="form-control input-group-addon" id="startBusizDay" name="Buz_str_date">
									</div>
								</div>
								<div class="form-group col-sm-4">
									<label>存量业务生效终止日:</label>
									<div class="input-group date full-width">
										<input type="text" class="form-control input-group-addon" id="endBusizDay" name="Buz_end_date">
									</div>
								</div>
							</div>
						</form>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="savePolicyInfo()">保存</button>
					</div>
				</div>
			</div>
		</div>
		
		<div class="modal inmodal" id="orgTree" tabindex="-2" role="dialog" data-backdrop="false">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" onclick="$('#orgTree').hide()"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 >选择机构</h2>
					</div>
					<div class="modal-body">
						<ul id="org" class="ztree"></ul>
					</div>
				</div>
			</div>
		</div>
		
		<div class="modal inmodal" id="proTree" tabindex="-2" role="dialog" data-backdrop="false">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" onclick="$('#proTree').hide()"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 >选择产品</h2>
					</div>
					<div class="modal-body">
						<ul id="pro" class="ztree"></ul>
					</div>
				</div>
			</div>
		</div>
		
		<div class="modal inmodal" id="indTree" tabindex="-2" role="dialog" data-backdrop="false">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" onclick="$('#indTree').hide()"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 >选择行业</h2>
					</div>
					<div class="modal-body">
						<ul id="ind" class="ztree"></ul>
					</div>
				</div>
			</div>
		</div>
		
		<div class="modal inmodal" id="uploadDialog" tabindex="-2" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 >导入政策性调节项信息</h2>
					</div>
					<div class="modal-body">
						<div class="row">
							<form  id="uploadForm" enctype ="multipart/form-data" method="post">
								<div class="col-xs-8">
									<div class="form-group">
										<input type="file" id="uploadFile" name="uploadFile">
									</div>
								</div>
								<div class="col-xs-4" style="padding-left: 0px;">
									<button class="btn btn-info" type="submit">提交</button>
									<a class="btn btn-link" style="text-decoration: underline;" href="/updownload/TempLateData.xlsx">下载模板</a>
								</div>
							</form>
						</div>
					</div>
				</div>
			</div>
		</div>
		
		<!--<div class="modal inmodal" id="downloadDialog" tabindex="-2" role="dialog" data-backdrop="false">
			<div class="modal-dialog modal-sm">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" onclick="$('#downloadDialog').hide()"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2>文件下载</h2>
					</div>
					<div class="modal-body">
						<div class="row" style="text-align: center;">
							<a  id="export" href="/updownload/ExpPolicyData.xlsx">文件生成完成，请点击下载</a>
						</div>
					</div>
				</div>
			</div>
		</div>-->
		
		<div class="modal inmodal" id="checkDialog" tabindex="-2" role="dialog">
			<div class="modal-dialog" style="width: 99%;">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close"  data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2>校验结果</h2>
					</div>
					<div class="modal-body">
						<div class="row">
							<div class="" style="float: left;">
								<button resid="" class="btn btn-warning " type="button" onclick="checkData()"> <i class="fa fa-legal"></i> 校验</button>
							</div>
							<div class="form-inline" style="float: right;">
								<div class="form-group">
									<select class="form-control" id="key1" style="width: 200px;">
										<option value="">------请选择调节项类型------</option>
									</select>
								</div>
								<div class="form-group">
									<input type="text" placeholder="请输入调节项维度" style="width: 200px;" id="key2" class="form-control" value="">
								</div>
								<button class="btn btn-default searchBtn" type="button" onclick="searchCheck()" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
							</div>
						</div><br />
						<div class="row">
							<div class="table-responsive">
								<table class="table table-striped table-bordered table-hover check" style="width: 99%;">
									<thead>
										<tr>
											<th>序号</th>
											<th>调节项类型</th>
											<th>机构</th>
											<th>币种</th>
											<th>调整项维度</th>
											<th>区间值-起</th>
											<th>区间值-止</th>
											<th>最新值日期</th>
											<th>生效起始日</th>
											<th>生效终止日</th>
											<th>存量业务生效起始日</th>
											<th>存量业务生效终止日</th>
											<th>调节点差</th>
											<th>冲突编码</th>
										</tr>
									</thead>
								</table>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
		
		<!--<form action="/mas/ftp/FtpAdjDownload" style="display: none;" id="downloadForm" method="post">
			<input type="hidden" name="keyword1" id="key1" value="" />
			<input type="hidden" name="keyword2" id="key2" value="" />
		</form>-->
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
		var treeType='';
		function zTreeOnDblClickSimple(event, treeId, treeNode) {
			//alert(treeNode ? treeNode.tId + ", " + treeNode.name : "isRoot");
			if(treeType==='org'){
				$('#OrgID').val(treeNode.id);
				$('#OrgDesc').val(treeNode.name);
	
				$('#orgTree').hide();
			}else if(treeType==='pro'){
				$('#proId').val(treeNode.id);
				$('#proDesc').val(treeNode.name);
	
				$('#proTree').hide();
			}else if(treeType==='ind'){
				$('#indId').val(treeNode.id);
				$('#indDesc').val(treeNode.name);
	
				$('#indTree').hide();
			}
		};

		function zTreeOnClickSimple(event, treeId, treeNode) {
			var treeObj = $.fn.zTree.getZTreeObj(treeId);
			treeObj.expandNode(treeNode, null, null, null);

			return true;
		};
	
	
		var policys, check;
		$(function() {
			//控制按钮权限
			$.get('/platform/DefaultMenu',{TypeId:2,Id:'30208020300000', r:Math.random()*10000000000000}, function(data){
				var rs=JSON.parse(data);
				$('.res').hide();
				if(rs!=null){
					rs.forEach(function(e){
						$("button[resid='"+e.Res_id+"']").show();
					});
				}
			})

			initTable();

			$('.policys tbody').on('click', 'tr', function() {
				$(this).find('input:radio').get(0).checked = true;
			});
			
			$.get('/platform/MasDimCurrency', {r:Math.random()*10000000000000}, function(data) {
				appendOption('curreny', data, 'IsoCurrencyCd', 'IsoCurrencyName');
			});
			
			$.get('/platform/BusizInfoCalc', {r:Math.random()*10000000000000}, function(data) {
				appendOption('busizId', data, 'Busiz_id', 'Busiz_desc');
			});
			
			$.get('/mas/ftp/FtpAlType', {r:Math.random()*10000000000000}, function(data) {
				appendOption('identify', data, 'Altypeid', 'Altypedesc');
			});
			
			$.get('/mas/ftp/adjust/info?TypeId=2', {r:Math.random()*10000000000000}, function(data) {
				appendOption('adjId', data, 'Adjustment_id', 'Adjustment_name');
				appendOption('keyword1', data, 'Adjustment_id', 'Adjustment_name');
				appendOption('key1', data, 'Adjustment_id', 'Adjustment_name');
			});
			
			$('.input-group-addon').datepicker({
				autoclose: true,
				format: 'yyyy-mm-dd'
			});
			
			$('#uploadFile').filestyle({
				buttonName : 'btn-warning',
				buttonText:'选择文件'
			});
			
			
			var options = { 
				contentType: "application/x-www-form-urlencoded;charset=utf-8",
			    url:'/mas/ftp/FtpAdjUpload?r='+Math.random()*10000000000000, 
			    success:function(data) { 
			        var rs=JSON.parse(data);
			        $.unblockUI();
		          	if(rs.ErrorCode==='1'){
			          	toastr.success(rs.ErrorMsg);
//						policys.ajax.reload(null, true);
						policys.ajax.url('/mas/ftp/FtpAdjustPolicy?r='+Math.random()*10000000000000).load();
						$('#uploadDialog').modal('hide');
	                }else {
	                    toastr.error(rs.ErrorMsg);
	                    return ;
	                }
			    },
				beforeSubmit: function(arr, $form, options) {
					var file=$('#uploadFile').val(),
						fileType=file.split('.');
					if(file==='' || fileType[fileType.length-1]!='xlsx'){
						toastr.warning("请选择要上传的Excel文件");
						return false;
					}
					$.get('/platform/MenuPage', function(){
						$('#uploadDialog').modal('hide');
							$.blockUI({
								message: '<h1>上传导入中。。。</h1>',
								css: {
									border: 'none',
									padding: '15px',
									backgroundColor: '#000',
									'-webkit-border-radius': '10px',
									'-moz-border-radius': '10px',
									opacity: .5,
									color: '#fff'
								}
						});
						return true;
					});
				}
			}; 
			$('#uploadForm').ajaxForm(options); 
		});

		function initTable() {
			policys = $('.policys').DataTable({
				sAjaxSource: '/mas/ftp/FtpAdjustPolicy?r='+Math.random()*10000000000000,
				info: true,
				bPaginate: true,
				sScrollX: "100%",
//				iDisplayLength: 5,
				//sScrollXInner: "150%",
				bScrollCollapse: true,
//				dom: 'Bfrt<ip>',
				sRowSelect: 'single',
//				buttons: [{
//						extend: 'copy'
//					},
//					//                  {extend: 'csv'},
//					{
//						extend: 'excel',
//						title: '政策性调节项'
//					}
//					//                  {extend: 'pdf', title: 'ExampleFile'},
//					//                  {extend: 'print',
//					//	                     customize: function (win){
//					//	                            $(win.document.body).addClass('white-bg');
//					//	                            $(win.document.body).css('font-size', '10px');
//					//	
//					//	                            $(win.document.body).find('table')
//					//	                                    .addClass('compact')
//					//	                                    .css('font-size', 'inherit');
//					//	                    }
//					//                  }
//				],
//buttons:['excelHtml5'],
				"sScrollX": true,
				aoColumns: [{
					"data": null
				},{
					"data": "Uuid"
				},{
					"data": "Adj_desc"
				}, {
					"data": "Org_unit_desc"
				}, {
					"data": "Iso_currency_cd"
				}, {
					"data": "Dyn_name"
				}, {
					"data": "Term_str"
				}, {
					"data": "Term_end"
				}, {
					"data": "Last_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Eff_str_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Eff_end_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Buz_str_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Buz_end_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Adj_bp"
				}],
				fnRowCallback: function(nRow, aData, iDisplayIndex, iDisplayIndexFull) {
					$('td:eq(0)', nRow).html('<div class="radio"><input type="radio" value="' + iDisplayIndex + '" name="tableRadio"><label></label></div>');
				},
				fnServerParams: function(aoData) {
					aoData.push({
						name: "keyword1",
						value: $('#keyword1').val()
					});
					aoData.push({
						name: "keyword2",
						value: $('#keyword2').val()
					});
				}
			});
			
			
			check = $('.check').DataTable({
				sAjaxSource: '/mas/ftp/FtpPCheckResult?r='+Math.random()*10000000000000,
				info: true,
				bPaginate: true,
				bScrollCollapse: true,
//				iDisplayLength: 3,
				aoColumns: [{
					"data": "Uuid"
				}, {
					"data": "Adj_desc"
				}, {
					"data": "Org_unit_desc"
				}, {
					"data": "Iso_currency_cd"
				}, {
					"data": "Dyn_name"
				}, {
					"data": "Term_str"
				}, {
					"data": "Term_end"
				}, {
					"data": "Last_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Eff_str_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Eff_end_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Buz_str_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Buz_end_date",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "Adj_bp"
				}, {
					"data": "Memo"
				}],
				fnServerParams: function(aoData) {
					aoData.push({
						name: "keyword1",
						value: $('#key1').val()
					});
					aoData.push({
						name: "keyword2",
						value: $('#key2').val()
					});
				}
			});
		}
		
		function search(){
//			policys.ajax.reload(null, true);
			policys.ajax.url('/mas/ftp/FtpAdjustPolicy?r='+Math.random()*10000000000000).load();
		}

		function showEditDialog(type) {
			$('#type').val(type);
			
			if(type === 'add') {
				$('#add').show();
				$('#edit').hide();
				
				$('#policyForm')[0].reset();
				$('#adjId').change();
				//机构默认取总行
				$.get('/mas/ftp/TopOrg', {r:Math.random()*10000000000000}, function(data){
					var rs=JSON.parse(data);
					$('#OrgID').val(rs.OrgUnitId);
					$('#OrgDesc').val(rs.OrgUnitDesc);
				});
				$('.input-group-addon').datepicker('update', '');
				$('#adjId').removeAttr('disabled');
			} else if(type === 'edit') {
				$('#add').hide();
				$('#edit').show();
				$('#policyForm')[0].reset();
				
				if($('.policys input:checked').length===0){
					toastr.warning('请选择一行进行编辑');
					return;
				}
				$('#adjId').attr('disabled', 'disabled');
				
				var info=policys.data()[$('.policys input:checked').val()];
				$('#adjId').val(info.Adj_id).change();
				$('#OrgID').val(info.Org_unit_id);
				$('#OrgDesc').val(info.Org_unit_desc);
				$('#curreny').val(info.Iso_currency_cd);
				if(info.Adj_id==='801'){ //产品调节项
					$('#proId').val(info.Adj_dyn_dim);
					$('#proDesc').val(info.Dyn_name);
				}else if(info.Adj_id==='802'){ //机构调节项
					//$('#adjId').val(info.Adj_id)
					$('#identify').val(info.Adj_dyn_dim);
				}else if(info.Adj_id==='803'){ //期限调节项
					$('#busizId').val(info.Adj_dyn_dim);
					$('#startTerm').val(info.Term_str);
					$('#endTerm').val(info.Term_end);
				}else if(info.Adj_id==='804'){ //行业调节项
					$('#indId').val(info.Adj_dyn_dim);
					$('#indDesc').val(info.Dyn_name);
				}else if(info.Adj_id==='805'){ //规模调节项
					$('#busizId').val(info.Adj_dyn_dim);
					$('#startScale').val(info.Term_str);
					$('#endScale').val(info.Term_end);
				}else if(info.Adj_id==='806'){ //账户调节项
					$('#accounts').val(info.Adj_dyn_dim);
				}
				
				$('#spot').val(info.Adj_bp);
				$('#startEffectDay').val(info.Eff_str_date.substring(0, 10));
				$('#endEffectDay').val(info.Eff_end_date.substring(0, 10));
				$('#startBusizDay').val(info.Buz_str_date.substring(0, 10));
				$('#endBusizDay').val(info.Buz_end_date.substring(0, 10));
				
				$('#uuid').val(info.Uuid);
			}
			
			$('#editDialog').modal('show');
		}
		
		function savePolicyInfo(){
			//校验
			var spot=$('#spot').val();
			var startEffectDay=$('#startEffectDay').val();
			var endEffectDay=$('#endEffectDay').val();
			var startBusizDay=$('#startBusizDay').val();
			var endBusizDay=$('#endBusizDay').val();
			
			var re = /^[-]{0,1}(\d+)$/; //整数可为负
			var regu = /^[0-9]+$/; //整数
			var regus = /^[0-9]+\.?[0-9]*$/; //数字
			
			if(spot==='' || !re.test(spot)){
				toastr.warning('请输入整数格式的点差值');
				return ;
			}else{
				if(parseInt(spot)<-9999 || parseInt(spot)>9999){
					toastr.warning('请输入介于-9999-9999的整数点差值');
					return ;
				}
			}
			
			if(startEffectDay==='' || endEffectDay===''){
				toastr.warning('请选择生效起止日');
				return ;
			}else {
				if(startEffectDay>=endEffectDay){
					toastr.warning('生效起始日需要小于生效终止日');
					return ;
				}
			}
			
			if(startBusizDay!='' && endBusizDay!=''){
				if(startBusizDay>=endBusizDay){
					toastr.warning('存量业务生效起始日需要小于存量业务生效终止日');
					return ;
				}
			}	
			
			var adjId=$('#adjId').val();
			if(adjId==='801'){ //产品调节项
				if($('proId').val()===''){
					toastr.warning('请选择产品');
					return ;
				}
			}else if(adjId==='802'){ //机构调节项
				if($('#identify').val()===''){
					toastr.warning('请选择资产负债类型');
					return ;
				}
			}else if(adjId==='803'){ //期限调节项
				if($('#busizId').val()===''){
					toastr.warning('请选择业务单元');
					return ;
				}
				
				var s=$('#startTerm').val();
				var e=$('#endTerm').val();
				if(s==='' || e === '' || !regu.test(s) || !regu.test(e)){
					toastr.warning('请输入正整数形式的起始期限和截止期限');
					return ;
				}else{
					if(parseInt(s)>parseInt(e)){
						toastr.warning('起始期限不能大于截止期限');
						return ;
					}
					
					if(parseInt(s)>360){
						toastr.warning('起始期限不能大于30年');
						return ;
					}
					if(parseInt(e)>360){
						toastr.warning('起始期限不能大于30年');
						return ;
					}
				}
			}else if(adjId==='804'){ //行业调节项
				if($('#indId').val()===''){
					toastr.warning('请选择行业');
					return ;
				}
			}else if(adjId==='805'){ //规模调节项
				if($('#busizId').val()===''){
					toastr.warning('请选择业务单元');
					return ;
				}
				
				var s=$('#startScale').val();
				var e=$('#endScale').val();
				if(s==='' || e === '' || !regus.test(s) || !regus.test(e)){
					toastr.warning('请输入数字形式的起始规模和截止规模');
					return ;
				}else{
					if(parseFloat(s)>parseFloat(e)){
						toastr.warning('起始规模不能大于截止规模');
						return ;
					}
				}
			}else if(adjId==='806'){ //账户调节项
				if($('#accounts').val()===''){
					toastr.warning('请输入帐号');
					return ;
				}
			}
			
			$('#adjId').removeAttr('disabled');
			var d=$('#policyForm').serialize();
			d=ReplaceAll(d, [['Adj_dyn_dim=&', ''], ['Term_end=&', ''], ['Term_str=&', '']]);
			if($('#type').val()==='add'){
				$.post('/mas/ftp/FtpAdjustPolicy?r='+Math.random()*10000000000000, d,function(data){
					var rs=JSON.parse(data);
		            if(rs.ErrorCode==='1'){
		                toastr.success('调节项信息保存成功！');
//						policys.ajax.reload(null, true);
						policys.ajax.url('/mas/ftp/FtpAdjustPolicy?r='+Math.random()*10000000000000).load();
						$('#editDialog').modal('hide');
		            }else {
		            		toastr.error(rs.ErrorMsg);
		                return false;
		            }
					
				});
			}else if($('#type').val()==='edit'){
				$.ajax({
					type:"put",
					url:"/mas/ftp/FtpAdjustPolicy?r="+Math.random()*10000000000000+"&"+d,
					async:true,
					error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
//		                    toastr.success(rs.ErrorMsg);
							toastr.success(rs.ErrorMsg);
//							policys.ajax.reload(null, true);
							policys.ajax.url('/mas/ftp/FtpAdjustPolicy?r='+Math.random()*10000000000000).load();
							$('#editDialog').modal('hide');
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
		            }
				});
			}
		}
		
		function changeInput(obj){
			$('.dimension').hide();
			$('.dimension input').val('');
			$('.dimension select').val('');
			
			$('.orgBtn').removeAttr('disabled');
			$('#startBusizDay').parent().parent().show();
			$('#endBusizDay').parent().parent().show();
			
			var info=[];
			if($(obj).val()==='801'){ //产品调节项
				info=['proId'];
			}else if($(obj).val()==='802'){ //机构调节项
				info=['identify'];
			}else if($(obj).val()==='803'){ //期限调节项
				info=['startTerm', 'endTerm', 'busizId'];
			}else if($(obj).val()==='804'){ //行业调节项
				info=['indId'];
			}else if($(obj).val()==='805'){ //规模调节项
				info=['startScale', 'endScale', 'busizId'];
			}else if($(obj).val()==='806'){ //账户调节项
				info=['accounts'];
				//机构不可选 //存量业务起止日不需要选
				$('.orgBtn').attr('disabled', 'disabled');
				$('#startBusizDay').datepicker('update', '');
				$('#endBusizDay').datepicker('update', '');
				$('#startBusizDay').parent().parent().hide();
				$('#endBusizDay').parent().parent().hide();
			}
			showMoreInfos(info);
		}
		
		function showMoreInfos(infos){
			infos.forEach(function(e){
				$("#"+e).parent().parent().show();
			});
		}
		
		function showOrgDialog(){
			if($('#adjId').val()==='806'){
				return ;
			}
			
			$.get('/platform/OrgTree' + '?r=' + Math.random(100000000000000), function(data) {
				$.fn.zTree.init($("#org"), s, JSON.parse(data)); //树
			});
			
			treeType='org';
			$('#orgTree').show();
		}
		
		function showProDialog(){
			$.get('/mas/ftp/ProductTree' + '?r=' + Math.random(100000000000000), function(data) {
				$.fn.zTree.init($("#pro"), s, JSON.parse(data)); //树
			});
			
			treeType='pro';
			$('#proTree').show();
		}
		
		function showIndDialog(){
			$.get('/mas/ftp/IndustryTree' + '?r=' + Math.random(100000000000000), function(data) {
				$.fn.zTree.init($("#ind"), s, JSON.parse(data)); //树
			});
			
			treeType='ind';
			$('#indTree').show();
		}
		
		function deletePolicy(){
			if($('.policys input:checked').length===0){
				toastr.warning('请选择一行进行编辑');
				return;
			}
			
			var info=policys.data()[$('.policys input:checked').val()];
			swal({
				title: "删除！",
				text: "是否删除选中的调节项信息？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				$.ajax({
		            type:"delete",
		            url:"/mas/ftp/FtpAdjustPolicy?Uuid="+info.Uuid+"&r="+Math.random()*10000000000000,
		            error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
//		                    toastr.success(rs.ErrorMsg);
		                    swal("删除!", rs.ErrorMsg, "success");
//							policys.ajax.reload(null, true);
							policys.ajax.url('/mas/ftp/FtpAdjustPolicy?r='+Math.random()*10000000000000).load();
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
		            }
		        });
			});
		}
		
		function exportData(){ //导出
			$.get('/platform/MenuPage', function(){
				var keyword1=$('#keyword1').val();
				var keyword2=$('#keyword2').val();
				
				var form=$("<form>");//定义一个form表单
				form.attr("style","display:none");
				form.attr("target","");
				form.attr("method","get");
				form.attr("action","/mas/ftp/FtpAdjDownload?r="+Math.random()*10000000000000);
				var input1=$("<input>");
				input1.attr("type","hidden");
				input1.attr("name","keyword1");
				input1.attr("value",keyword1);
				var input2=$("<input>");
				input2.attr("type","hidden");
				input2.attr("name","keyword2");
				input2.attr("value",keyword2);
				$("body").append(form);//将表单放置在web中
				form.append(input1);
				form.append(input2);
				
				form.submit();//表单提交
			});
		}
		
		function showUploadDialog(){
			$('#uploadForm')[0].reset();
			$('#uploadDialog').modal('show');
		}
		
//		function upload(){
//			$.ajaxFileUpload
//          (
//              {
//                  url: '/mas/ftp/FtpAdjUpload', //用于文件上传的服务器端请求地址
//                  secureuri: false, //是否需要安全协议，一般设置为false
//                  fileElementId: 'uploadFile', //文件上传域的ID
//                  dataType: 'json', //返回值类型 一般设置为json
//                  success: function (data, status)  //服务器成功响应处理函数
//                  {
//                  }
//              }
//          )
//          return false;
//		}

		function showCheckDialog(){ //校验完成后显示结果
//			check.ajax.reload(null, true);
			check.ajax.url('/mas/ftp/FtpPCheckResult?r='+Math.random()*10000000000000).load();
			$('#checkDialog').modal('show');
		}
		
		function checkData(){
			toastr.info('校验调节项数据中。。。。');
			$.post('/mas/ftp/FtpCallPCheckProc', {r:Math.random()*10000000000000}, function(data){
				var rs=JSON.parse(data);
	            if(rs.ErrorCode==='1'){
	                toastr.success('校验调节项数据成功！');
//					check.ajax.reload(null, true);
					check.ajax.url('/mas/ftp/FtpPCheckResult?r='+Math.random()*10000000000000).load();
	            }else {
	            		toastr.error(rs.ErrorMsg);
	                return false;
	            }
			});
		}
		
		function searchCheck(){
//			check.ajax.reload(null, true);
			check.ajax.url('/mas/ftp/FtpPCheckResult?r='+Math.random()*10000000000000).load();
		}
	</script>

</html>