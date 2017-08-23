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

		<title>基础价格计算</title>

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
			#batchList{
				
			}
			#batchList li{
				/*display: inline;*/
			}
			#runningBatch,#errorBatch,#endBatch{
            		height: 100%;
            		/*border: 1px solid;*/
            		padding: 0px;
            }
            .modal-header .close{
				padding-top: 4px !important;
				padding-right: 8px !important;
			}
		</style>
	</head>

	<body>
		<div id="wrapper">
			<div class="wrapper wrapper-content animated fadeInRight full-height">
				<div class="row full-height content">
					<div class="col-xs-12" id="opList">
						<div class="ibox float-e-margins" style="margin-bottom: 0px;">
							<div class="ibox-title">
								<div class="row">
									<!--<div class="col-xs-4">
										<h3>批次列表</h3>
									</div>-->
									<div class="col-xs-10 text-left">
										<div class="form-inline">
											<button resid="30208030101000" class="btn btn-info res" type="button" onclick="showBatchDialog('1')" style="margin-bottom: 0px;">新增</button>
											<!--<button resid="208030102000" class="btn btn-primary res" type="button" onclick="showBatchDialog('2')" style="margin-bottom: 0px;">编辑</button>-->
											<button resid="30208030102000" type="button" class="btn btn-primary res" onclick="showBatchDialog('2')" style="margin-bottom: 0px;">编辑</button>
											<button resid="30208030107000" type="button" class="btn btn-danger res" onclick="deleteBatch()"  style="margin-bottom: 0px;">删除</button>
										</div>
									</div>
									<div class="col-xs-2">
										<button resid="30208030103000" class="btn btn-info btn-block res" type="button" onclick="runBatch()" style="">开始</button>
									</div>
								</div>

							</div>
							<div class="ibox-content lists" style="overflow-x: auto;padding:15px 0px 0px 0px;">
								<ul class="sortable-list connectList agile-list" id="batchList">
								</ul>
							</div>
						</div>
					</div>
					<!--<div class="col-xs-12" id="s">
						<div class="ibox float-e-margins" style="margin-bottom: 0px;border: 0px;">
							<div class="ibox-title" style="border: 0px;">
								<button resid="208030103000" class="btn btn-info btn-block res" type="button" onclick="runBatch()" style="">开始</button>
							</div>
						</div>
					</div>-->
					<div class="col-xs-12" id="rs">
						<div class="ibox float-e-margins full-height">
							<div class="ibox-content full-height" style="padding: 0px;">
								<div class="col-lg-12 full-height">
				                    <div id="runningBatch" class="tab-pane active col-sm-4" style="padding: 10px 10px 10px 0px;">
		                            		<div height="60px" width="100%" style="background-color: #89cbc1;text-align: center;"><img src="/static/theme/default/img/runningBatch.png" style="height: 60px;width: 180px;padding: 10px;"/></div>
		                                <div class="panel-body row" style="height: 95%;overflow: auto;border: 1px solid #cccccc;margin: 0px;">
		                                		
		                                </div>
		                            </div>
		                            <div id="endBatch" class="tab-pane col-sm-4" style="padding: 10px 0px 10px 10px;">
		                            		<div height="60px" width="100%" style="background-color: #a8cebf;text-align: center;"><img src="/static/theme/default/img/endBatch.png" style="height: 60px;width: 180px;padding: 10px;"/></div>
		                                <div class="panel-body" style="height: 95%;overflow: auto;border: 1px solid #cccccc;">
		                                		<button class="btn btn-link col-sm-6"  onclick="selectAll('2',this)">全选</button>
		                                		<button resid="30208030106000" class="btn btn-default res col-sm-6" onclick="deleteBatchs('2')">清除</button>
		                                		<div style="clear: both;"></div>
		                                		<hr />
		                                		<div class="row">
		                                			
		                                		</div>
		                                </div>
		                            </div>
		                            <div id="errorBatch" class="tab-pane col-sm-4" style="padding: 10px;">
		                            		<div height="60px" width="100%" style="background-color: #b7d9c0;text-align: center;"><img src="/static/theme/default/img/errorBatch.png" style="height: 60px;width: 180px;padding: 10px;"/></div>
		                                <div class="panel-body" style="height: 95%;overflow: auto;border: 1px solid #cccccc;">
		                                		<button class="btn btn-link col-sm-6" onclick="selectAll('1',this)">全选</button>
		                                		<button resid="30208030105000" class="btn btn-default res col-sm-6" onclick="deleteBatchs('1')">清除</button>
		                                		<div style="clear: both;"></div>
		                                		<hr />
		                                		<div class="row">
		                                		</div>
		                                </div>
		                            </div>
		                            
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

		<div class="modal inmodal" id="addBatchDialog" tabindex="-1" role="dialog" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2>新增批次信息</h2>
					</div>
					<div class="modal-body">
						<form id="batch">
							<div class="row">
								<div class="col-xs-12">
									<div class="col-xs-6">
										<div class="form-group">
											<label>批次编码:</label>
											<input type="text" placeholder="" class="form-control" name="DispatchId" id="DispatcId">
										</div>
									</div>
									<div class="col-xs-6">
										<div class="form-group">
											<label>批次名称:</label>
											<input type="text" placeholder="" class="form-control" name="DispatcName" id="DispatcName">
										</div>
									</div>
								</div>
								<div class="col-xs-12">
									<div class="col-xs-6">
										<div class="form-group">
											<label>输入:</label>
											<select name="InputSouceCd" id="input" class="form-control">
											</select>
										</div>
									</div>
									<div class="col-xs-6">
										<div class="form-group">
											<label>输出:</label>
											<select name="OutputResultCd" id="output" class="form-control">
											</select>
										</div>
									</div>
								</div>
								<div class="col-xs-12">
									<div class="col-xs-6">
										<div class="form-group">
											<label>偏移量:</label>
											<input type="text" placeholder="" class="form-control" name="StartOffset" id="StartOffset" value="0">
										</div>
									</div>
									<div class="col-xs-6">
										<div class="form-group">
											<label>最大限制:</label>
											<input type="text" placeholder="" class="form-control" name="MaxLimit" id="MaxLimit" value="10000000000">
										</div>
									</div>
								</div>
								<!--<div class="col-xs-12">
									<div class="col-xs-6">
										<div class="form-group">
											<label>所属域:</label>
											<select name="DomainId" id="DomainId" class="form-control">
											</select>
										</div>
									</div>
								</div>-->
							</div>
						</form>
					</div>

					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="submitBase()">保存</button>
					</div>
				</div>
			</div>
		</div>
		
		<div class="modal inmodal" id="editBatchDialog" tabindex="-1" role="dialog" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2>编辑批次信息</h2>
					</div>
					<div class="modal-body">
						<form id="batchInfo">
							<input type="hidden" name="index" id="index"/>
							<div class="row">
								<div class="col-xs-12">
									<div class="col-xs-6">
										<div class="form-group">
											<label>批次编码:</label>
											<input type="text" placeholder="" class="form-control"name="DispatchId" id="batchCode" readonly="">
										</div>
									</div>
									<div class="col-xs-6">
										<div class="form-group">
											<label>批次名称:</label>
											<input type="text" placeholder="" class="form-control" name="DispatcName" id="batchName">
										</div>
									</div>
								</div>
								<div class="col-xs-12">
									<div class="col-xs-6">
										<div class="form-group">
											<label>输入:</label>
											<select name="InputSouceCd" id="ftpDispatchInput" class="form-control">
											</select>
										</div>
									</div>
									<div class="col-xs-6">
										<div class="form-group">
											<label>输出:</label>
											<select name="OutputResultCd" id="ftpDispatchOutput" class="form-control">
											</select>
										</div>
									</div>
								</div>
								<div class="col-xs-12">
									<div class="col-xs-6">
										<div class="form-group">
											<label>偏移量:</label>
											<input type="text" placeholder="" class="form-control" name="StartOffset" id="startOffset" value="0">
										</div>
									</div>
									<div class="col-xs-6">
										<div class="form-group">
											<label>最大限制:</label>
											<input type="text" placeholder="" class="form-control" name="MaxLimit" id="maxLimit" value="10000000000">
										</div>
									</div>
								</div>
								<!--<div class="col-xs-12">
									<div class="col-xs-6">
										<div class="form-group">
											<label>所属域:</label>
											<select name="DomainId" id="Domain_id_Base" class="form-control">
											</select>
										</div>
									</div>
								</div>-->
							</div>
						</form>
					</div>

					<div class="modal-footer">
						<!--<button resid="208030102000" type="button" class="btn btn-primary res" onclick="updateBatch()">更新</button>
						<button resid="208030107000" type="button" class="btn btn-danger res" onclick="deleteBatch()" >删除</button>-->
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="updateBatch()">保存</button>
					</div>
				</div>
			</div>
		</div>
	</body>
	<script type="text/javascript">
		var batchList=[]; //批次信息列表
		var batchRunList = []; //正在运行或报错停止的批次列表
		$(function() {
			initKnob();
			
			$.ajaxSetup({
//				async: false,
				cache: false
			});
//			$('.lists').height(document.body.clientHeight - 200);
//			$('.resultInfo').height(document.body.clientHeight - 250);
			$('#rs').height(700);

			getBatchList();
			getBatchRunList();

			$.get('/platform/FtpDispatchInput', {r:Math.random()*10000000000000}, function(data) {
				appendOption('ftpDispatchInput', data, 'InputSourceCd', 'InputSourceDesc');
				appendOption('input', data, 'InputSourceCd', 'InputSourceDesc');
			})

			$.get('/platform/FtpDispatchOutput', {r:Math.random()*10000000000000}, function(data) {
				appendOption('ftpDispatchOutput', data, 'OutputResultCD', 'OutputResultDesc');
				appendOption('output', data, 'OutputResultCD', 'OutputResultDesc');
			})

//			$.get('/platform/SysDomainInfo', function(data) {
//				appendOption('Domain_id_Base', data, 'DomainId', 'DomainName');
//				appendOption('DomainId', data, 'DomainId', 'DomainName');
//				$('#Domain_id_Base').val('FTP');
//				$('#DomainId').val('FTP');
//			});
			
			setInterval("loopBatchStatus()", 5000); //定时取跑批信息
			
			initBtnPermission('30208030100000');
		});
		
		function initBtnPermission(resUpId){
			//控制按钮权限
			$.get('/platform/DefaultMenu',{TypeId:2,Id:resUpId,r:Math.random()*10000000000000}, function(data){
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
		}

		function initDate(){
			$('.selectDate').datepicker({
				autoclose: true,
				format: 'yyyy-mm-dd',
				language: 'zh-CN'
			});
//			$('.selectDate').val(new Date().Format('yyyy-MM-dd'));
//			$('.selectDate').setDate(new Date());
			$('.selectDate').datepicker('setDate', new Date());
		}
		function iCheckInit() {
			$('.i-checks').iCheck({
				checkboxClass: 'icheckbox_square-green',
				radioClass: 'iradio_square-green',
			});
		}
		function initKnob(){
			$(".dial").knob({
				readOnly:true,
				height:100,
				width:100
			});
		}

		var colors = ['warning', 'success', 'info', 'danger'];
		function getBatchList() {
			$.get('/platform/FtpBaseValueCal?r=' + Math.random(100000000000000), function(data) {
				batchList = JSON.parse(data);
				$('#batchList').html('');
				if(batchList!=null){
					batchList.forEach(function(e, index) {
						var c = colors[index % colors.length];
						
						appendBatchList(c, index, e);
					});
				}else {
					batchList=[];
				}
				
				initDate();
				iCheckInit();
			});
		}

		function appendBatchList(c, index, e){
			var temp = "<div class='col-sm-6'><li class=' " + c + "-element'><div class='input-group full-width' bid='" + index + "'>" +
						"<label class='checkbox-inline i-checks pull-left' style='margin: auto 0;'>" +
						"<input type='checkbox'/><span style='margin-left: 5px;margin-top: 1px;' class='caption' name='" + e.DispatchId + "'>" + e.DispatcName + "</span></label>" +
						"<div class='input-group date pull-right	'>" +
						"<input type='text' class='form-control input-group-addon selectDate' id='"+e.DispatchId+"' placeholder='跑批日期'>" +
						"</div></div></li><div>";
			$('#batchList').append(temp);
		}

		function showBatchDialog(type) {
			if(type==='1'){ //新增
				$('#batch')[0].reset();
				$('#addBatchDialog').modal('show');
			}else if(type==='2'){ //编辑
				if($('#batchList .checked').length===0){
					toastr.warning('请勾选要编辑的批次');
					return;
				}
				var index=$($('#batchList div.checked')[0]).parent().parent().attr('bid');
				
				$('#batchCode').val(batchList[index].DispatchId);
				$('#index').val(index);
				$('#batchName').val(batchList[index].DispatcName);
//				$('#Domain_id_Base').val(batchList[index].DomainId);
				$('#ftpDispatchInput').val(batchList[index].InputSouceCd);
				$('#ftpDispatchOutput').val(batchList[index].OutputResultCd);
				$('#startOffset').val(batchList[index].StartOffset);
				$('#maxLimit').val(batchList[index].MaxLimit);
				
				$('#editBatchDialog').modal('show');
			}
		}
		
		function submitBase(){
			//判断是否是字母、数字组成 
    			var reg = /^[0-9a-zA-Z]+$/;
    			if(!reg.test($('#DispatcId').val())){
    				toastr.warning('请输入由数字或字母组成的批次编码');
    				return;
    			}
    			
    			if($('#DispatcName').val()===''){
    				toastr.warning('请输入批次名称');
    				return;
    			}
    			
    			if($('#StartOffset').val()===''){
    				toastr.warning('请输入偏移量');
    				return;
    			}
    			
    			if($('#MaxLimit').val()===''){
    				toastr.warning('请输入最大限制');
    				return;
    			}
			
			$.post('/platform/FtpBaseValueCal?' + $('#batch').serialize()+'&r='+ Math.random(100000000000000), function(data) {
				var rs = JSON.parse(data);
				if(rs.ErrorCode === '1') {
					toastr.success(rs.ErrorMsg);
//					getBatchList();
					batchInfo={};
					$('#batch').serializeArray().forEach(function(e,index){
						batchInfo[e['name']]=e['value'];
					});
					batchList.push(batchInfo);
					appendBatchList(colors[(batchList.length-1) % colors.length], batchList.length-1, batchInfo);
					
//					initDate();
					$('input[id="'+batchInfo.DispatchId+'"]').datepicker({
						autoclose: true,
						format: 'yyyy-mm-dd'
					});
					$('input[id="'+batchInfo.DispatchId+'"]').val(new Date().Format('yyyy-MM-dd'));
					iCheckInit();
					$('#addBatchDialog').modal('hide');
				} else {
					toastr.error(rs.ErrorMsg);
					return;
				}
			});
		}
		
		function updateBatch() {
    			if($('#batchName').val()===''){
    				toastr.warning('请输入批次名称');
    				return;
    			}
    			
    			if($('#startOffset').val()===''){
    				toastr.warning('请输入偏移量');
    				return;
    			}
    			
    			if($('#maxLimit').val()===''){
    				toastr.warning('请输入最大限制');
    				return;
    			}
			
			$.ajax({
				type: "put",
				url: "/platform/FtpBaseValueCal?" + $('#batchInfo').serialize()+'&r='+ Math.random(100000000000000),
				error: function(msg) {
					console.log(msg.responseText);
				},
				success: function(data) {
					var rs = JSON.parse(data);
					if(rs.ErrorCode === '1') {
						toastr.success(rs.ErrorMsg);
						$('#batchInfo').serializeArray().forEach(function(e,index){
							batchList[$('#index').val()][e.name]=e.value;
						});
						$("span[name='"+$('#batchCode').val()+"']").text($('#batchName').val());
						$('#editBatchDialog').modal('hide');
					} else {
						toastr.error(rs.ErrorMsg);
						return;
					}
				}
			});
		}
		
		function deleteBatch(){
			var cs=$('#batchList div.checked'), json=[], index=[];
			cs.each(function(){
				var i=$(this).parent().parent().attr('bid');
				index.push(i)
				json.push(batchList[i].DispatchId);
			});
			
			swal({
				title: "删除！",
				text: "是否删除选中的批次信息？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				$.ajax({
					type: "delete",
					url: "/platform/FtpBaseValueCal?JSON=" + JSON.stringify(json) +'&r='+ Math.random(100000000000000),
					error: function(msg) {
						console.log(msg.responseText);
					},
					success: function(data) {
						var rs = JSON.parse(data);
						if(rs.ErrorCode === '1') {
							swal("删除!", "批次信息已删除", "success");
							$('#batchInfo')[0].reset();
							
							index.forEach(function(e){
								$("div[bid='"+e+"']").parent().parent().remove();
							})
							
							$('#editBatchDialog').modal('hide');
//							getBatchList();
						} else {
							toastr.error(rs.ErrorMsg);
							return;
						}
					}
				});
			});
		}
		
//		var overBatch=[];
		function getBatchRunList() {
			$("[data-toggle='tooltip']").tooltip('destroy');
			$.get('/platform/FtpDispatchRealt?r=' + Math.random(100000000000000), function(data) {
				$('#runningBatch .panel-body').html('');
				$('#errorBatch .panel-body .row').html('');
				$('#endBatch .panel-body .row').html('');
				var rs = JSON.parse(data);
				batchRunList = rs;

				if(rs!=null){
					var errorIndex=0, endIndex=0;
					rs.forEach(function(e) {
						if(e.DispatchStatus === '1') { //运行中的批次
//							if(Math.floor((parseInt(e.CurRows) / parseInt(e.AllRows)) * 100)===100){ //如果百分百了，提示统计中
//								var s={date:e.DispatchDate, id:e.DispatchId};
//								if($.inArray(s, overBatch)==-1){
//									overBatch.push(e);
//									toastr.warning('跑批已完成，统计结果中');
//								}
//							}
							var temp="<div class='batch col-xs-6' style='margin:0px;' data-container='body' data-toggle='popover' data-placement='bottom' data-content='"+e.ErrMsg+"'>"
									+"<label>"+e.DispatchName + " " + e.DispatchDate+"</label>"
									+"<div><input type='text' value='"+Math.floor((parseInt(e.CurRows) / parseInt(e.AllRows)) * 100)+"' class='dial'></div><div><button bid='" + e.DispatchId + "' bdate='" + e.DispatchDate + "' class='btn btn-warning res' resid='208030104000' onclick='stopBatch(this)'>停止</button></div></div>";
							$('#runningBatch .panel-body').append(temp);
						} else if(e.DispatchStatus === '2') { //错误批次
							errorIndex++;
							var temp="<div class='col-xs-6' style='margin-bottom:2px;'"
									+"'><label class='checkbox-inline i-checks pull-left' style='font-size: 14px;margin: auto 0;width:100%;'>"
									+"<input type='checkbox' />"+"<span data-container='body' data-toggle='tooltip' data-html='true' data-placement='bottom' data-original-title='<span style=\"color:red;\">"+e.ErrMsg+"</span>' style='margin-left: 5px;margin-top: 1px;' "
									+"bid='" + e.DispatchId + "' bdate='" + e.DispatchDate + "'>"+e.DispatchName+" "+e.DispatchDate+"</span></label></div>"
							if(errorIndex%2===1){
								$('#errorBatch .panel-body .row').append("<div style='clear:both;'></div>");
							}
							$('#errorBatch .panel-body .row').append(temp);
						} else if(e.DispatchStatus === '3') { //已完成批次
							endIndex++;
							var temp="<div class='col-xs-6' style='margin-bottom:2px;'"
									+"'><label  class='checkbox-inline i-checks pull-left' style='font-size: 14px;margin: auto 0;width:100%;'>"
									+"<input type='checkbox' />"+"<span data-container='body' data-toggle='tooltip' data-placement='bottom' title='"+e.ErrMsg+"' style='margin-left: 5px;margin-top: 1px;' "
									+"bid='" + e.DispatchId + "' bdate='" + e.DispatchDate + "'>"+e.DispatchName+" "+e.DispatchDate+"</span></label></div>"
							if(endIndex%2===1){
								$('#endBatch .panel-body .row').append("<div style='clear:both;'></div>");
							}
							$('#endBatch .panel-body .row').append(temp);
						}
					});
				}
				initKnob();
				iCheckInit();
				initBtnPermission('30208030100000');
				
				$("[data-toggle='tooltip']").tooltip();
			});
		}
		
		function loopBatchStatus() {
			getBatchRunList();
		}
		
		
		function runBatch(){
			if($('#batchList .checked').length===0){
				toastr.warning('请勾选要跑批的批次');
				return;
			}
			
			var runList = []; //要启动的批次
			//获取批次信息
			$('#batchList .checked').each(function() {
				var id = $(this).parent().parent().attr('bid');
				batchList[id].batchDate=$(this).parent().next().find('input').val();
				var checkBatch = batchList[id];

				var flag = 0;
				if(batchRunList != null) {
					batchRunList.forEach(function(e) {
//						if(e.DispatchId === checkBatch.DispatchId && e.DispatchDate === checkBatch.batchDate && e.DispatchStatus === '1') { //正在运行的批次
						if(e.DispatchId === checkBatch.DispatchId && e.DispatchStatus === '1') { //正在运行的批次
							toastr.warning(e.DispatchName + " " + "正在跑批中......");
							flag = 1;
						}
					});
				}

				if(flag === 0) {
					runList.push(checkBatch);
				}
			});

			if(runList.length == 0) {
				toastr("没有发起新的批次");
				return;
			}
			
			var input=[],output=[];
			runList.forEach(function(e){
				if(input.indexOf(e.InputSouceCd)==-1){
					input.push(e.InputSouceCd);
				}
				if(output.indexOf(e.OutputResultCd)==-1){
					output.push(e.OutputResultCd);
				}
			});
			
			if(input.length!=runList.length || output.length!=runList.length){
				toastr.error("批次输入输出含有重复，请检查批次信息");
				return;
			}
			
			$.blockUI({
				baseZ:9999,
				message: '<h1>开始跑批中。。。</h1>',
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
			
			$.ajax({
				type: "post",
				url: "/platform/FtpPatchCalc?JSON=" + encodeURIComponent(JSON.stringify(runList))+'&r='+ Math.random(100000000000000),
				contentType: "application/x-www-form-urlencoded; charset=utf-8",
				async: true,
				success: function(data) {
					var rs = JSON.parse(data);
					if(rs.ErrorCode === '1') {
						toastr.info(rs.ErrorMsg);
						getBatchRunList();
					} else {
						toastr.error(rs.ErrorMsg);
//						return false;
					}
					$.unblockUI();
				},
				error: function() {}
			});
		}
		
		var selectErrFlag=0,selectEndFlag=0;
		function selectAll(type,obj){ //全选
			if('1'===type){ //错误批次
				if(selectErrFlag==0){
					selectErrFlag++;
					$('#errorBatch .icheckbox_square-green').iCheck('check');
				}else {
					selectErrFlag--;
					$('#errorBatch .icheckbox_square-green').iCheck('uncheck');
				}
			}else if('2'===type){ //已完成批次
				if(selectEndFlag==0){
					selectEndFlag++;
					$('#endBatch .icheckbox_square-green').iCheck('check');
				}else {
					selectEndFlag--;
					$('#endBatch .icheckbox_square-green').iCheck('uncheck');
				}
				
			}
		}
		
		function deleteBatchs(type){
			swal({
				title: "删除！",
				text: "是否清除选中的批次？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				var batchs = [];
				if('1'===type){ //错误批次
					$('#errorBatch .checked').each(function(){
						batchs.push({
							"DispatchId": $(this).next().attr('bid'),
							"DispatchDate": $(this).next().attr('bdate')
						});
					});
				}else if('2'===type){ //已完成批次
					$('#endBatch .checked').each(function(){
						batchs.push({
							"DispatchId": $(this).next().attr('bid'),
							"DispatchDate": $(this).next().attr('bdate')
						});
					});
				}
				
				$.ajax({
					type: "delete",
					url: "/platform/FtpDispatchRealt?JSON=" + JSON.stringify(batchs)+'&r='+ Math.random(100000000000000),
					error: function(msg) {
						console.log(msg.responseText);
					},
					success: function(data) {
						var rs = JSON.parse(data);
						if(rs.ErrorCode === '1') {
							swal("删除!", "清除完成", "success");
							getBatchRunList();
						} else {
							toastr.error(rs.ErrorMsg);
							return;
						}
					}
				});
			});
		}
		
		function stopBatch(obj) { //停止某个批次
			swal({
				title: "警告！",
				text: "是否停止该批次？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				$.ajax({
					type: "put",
					url: "/platform/FtpDispatchRealt?DispatchId=" + $(obj).attr('bid') + "&DispatchDate=" + $(obj).attr('bdate')+'&r='+ Math.random(100000000000000),
					error: function(msg) {
						console.log(msg.responseText);
					},
					success: function(data) {
						var rs = JSON.parse(data);
						if(rs.ErrorCode === '1') {
							swal("提示!", "停止成功", "success");
							getBatchRunList();
						} else {
							toastr.error(rs.ErrorMsg);
							return;
						}
					}
				});
			});
		}
	</script>

</html>