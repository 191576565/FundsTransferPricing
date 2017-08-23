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

		<title>币种定义</title>

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
										<button resid="30101040201000" class="btn btn-primary res" type="button" onclick="showEditDialog('1')"> <i class="fa fa-plus"></i> 新增</button>
										<button resid="30101040202000" class="btn btn-info editBtn res" type="button" onclick="showEditDialog('2')"><i class="fa fa-paste"></i> 编辑</button>
										<!--<button resid="101040203000" class="btn btn-danger deleteBtn res" type="button" onclick="deleteCurrInfo()"><i class="fa fa-times"></i> 删除</button>-->
									</div>
									<div class="col-xs-4 text-right">
										<div class="form-inline">
											<div class="form-group">
												<input type="text" placeholder="请输入币种编码或名称" style="width: 200px;" id="keyword" class="form-control" value="">
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
												<th>币种编码</th>
												<th>币种名称</th>
												<th>状态</th>
												<th>排序编号</th>
												<th>创建人</th>
												<th>创建时间</th>
												<th>备注</th>
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
						<h2 id="addCurrInfo">新增币种信息</h2>
						<h2 id="editCurrInfo">编辑币种信息</h2>
						<input type="hidden" id="type"/>
					</div>
					<div class="modal-body">
						<form id="currInfo">
							<div class="row">
								<div class="col-xs-6">
										<div class="form-group">
											<label>币种编码:</label>
											<input type="text" placeholder="" class="form-control" name="IsoCurrencyCd" id="currId">
										</div>
								</div>
								<div class="col-xs-6">
										<div class="form-group">
											<label>币种名称:</label>
											<input type="text" placeholder="" class="form-control" name="IsoCurrencyDesc" id="currDesc">
										</div>
								</div>
								<div class="col-xs-6">
										<div class="form-group">
											<label>排序编号:</label>
											<input type="text" placeholder="" class="form-control" name="SortOrder" id="currSort">
										</div>
								</div>
								
								<div class="col-xs-6">
									<div class="form-group">
										<label>状态:</label>
										<select name="IsoCurrencyStatus" id="currStatus" class="form-control">
											<!--<option value="0">启用</option>
					                        <!--<option value="1">锁定</option>-->
					                        <!--<option value="1">失效</option>-->
										</select>
									</div>
								</div>
								<div class="col-xs-6">
										<div class="form-group">
											<label>备注:</label>
											<input type="text" placeholder="" class="form-control" name="MemoBk" id="currMemo">
										</div>
								</div>
							</div>
						</form>
					</div>

					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="saveCurrInfo()">保存</button>
					</div>
				</div>
			</div>
		</div>
		
	</body>
	<script type="text/javascript">
		var currs; 
		$(function(){
			//控制按钮权限
			$.get('/platform/DefaultMenu',{TypeId:2,Id:'30101040200000',r:Math.random()*10000000000000}, function(data){
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
			
//			$('.currs').height(document.body.clientHeight-200);
			
			$('.currs tbody').on('click', 'tr', function() {
//				if($(this).find('input:radio').get(0).checked) {
//					$(this).find('input:radio').get(0).checked = false;
//				} else {
//					$(this).find('input:radio').get(0).checked = true;
//				}
				$(this).find('input:radio').get(0).checked = true;
			});
			
			$.get('/platform/CurrencyStatus', {r:Math.random()*10000000000000}, function(data){
				appendOption('currStatus', data, 'Cstatus', 'CstatusDesc');
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
				sAjaxSource: '/platform/ComCurrency?r='+Math.random()*10000000000000,
				info: false,
				bPaginate: false,
				aoColumns: [{
					"data": null
				}, {
					"data": "IsoCurrencyCd"
				}, {
					"data": "IsoCurrencyDesc"
				}, {
					"data": "IsoCurrencyStatusDesc"
				}, {
					"data": "SortOrder"
				}, {
					"data": "CurrencyOwner"
				}, {
					"data": "EffectiveDate",
					render: function(data, type, row) {
						return data.substring(0, 10);
					}
				}, {
					"data": "MemoBk"
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
				$('#addCurrInfo').show();
				$('#editCurrInfo').hide();
				
				$('#currInfo')[0].reset();
				$('#currId').removeAttr('readonly');
			}else if(type==='2'){
				if($('.currs input:checked').length===0){
					toastr.warning('请选择一行币种信息进行编辑');
					return;
				}
				
				$('#addCurrInfo').hide();
				$('#editCurrInfo').show();
				$('#currId').attr('readonly', 'readonly');
				
				var info = currs.data()[$('.currs input:checked').val()];
				$('#currId').val(info.IsoCurrencyCd);
				$('#currDesc').val(info.IsoCurrencyDesc);
				$('#currMemo').val(info.MemoBk);
				$('#currSort').val(info.SortOrder);
				
				$('#currStatus').val(info.IsoCurrencyStatus);
			}
			$('#editDialog').modal('show');
		}
		
		function saveCurrInfo(){
			//输入校验
			//判断是否是字母 
    			var id = /^[a-zA-Z]+$/;
    			//判断是否是汉字、字母、数字组成 
    			var name = /^[0-9a-zA-Z\u4e00-\u9fa5]+$/;
    			
    			var regu = /^[0-9]+$/;
    			
    			var currId=$('#currId').val();
    			var currDesc=$('#currDesc').val();
    			var currSort=$('#currSort').val();
    			
    			if(!id.test(currId)){
    				toastr.warning('币种编码由字母组成');
    				return;
    			}
    			if(currDesc===''){
    				toastr.warning('请输入币种名称');
    				return;
    			}
    			if(regu.test(currSort)){
    				if(parseInt(currSort)<=0 || parseInt(currSort)>9999){
    					toastr.warning('请输入1-9999的正整数形式的排序编号');
    					return;
    				}
    			}else{
    				toastr.warning('请输入1-9999的正整数形式的排序编号');
    				return;
    			}
			
			if($('#type').val()==='1'){ //新增
				$.post('/platform/ComCurrency?r='+Math.random()*10000000000000, $('#currInfo').serialize(),function(data){
					var rs=JSON.parse(data);
		            if(rs.ErrorCode==='1'){
		                toastr.success('币种信息保存成功！');
//						currs.ajax.reload(null, true);
						currs.ajax.url('/platform/ComCurrency?r='+Math.random()*10000000000000).load();
						$('#editDialog').modal('hide');
		            }else {
		            		toastr.error(rs.ErrorMsg);
		                return false;
		            }
					
				});
			}else if($('#type').val()==='2'){ //编辑
				$.ajax({
					type:"put",
					url:"/platform/ComCurrency?"+$('#currInfo').serialize()+"&r="+Math.random()*10000000000000,
					async:true,
					error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
//		                    toastr.success(rs.ErrorMsg);
							toastr.success('币种信息编辑成功！');
//							currs.ajax.reload(null, true);
							currs.ajax.url('/platform/ComCurrency?r='+Math.random()*10000000000000).load();
							$('#editDialog').modal('hide');
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
						
		            }
				});
			}
			
		}
		
		function deleteCurrInfo(){
			if($('.currs input:checked').length===0){
				toastr.warning('请选择一行币种信息进行删除');
				return;
			}
			
			var info = currs.data()[$('.currs input:checked').val()];
			swal({
				title: "删除！",
				text: "是否删除选中的币种信息？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				$.ajax({
		            type:"delete",
		            url:"/platform/ComCurrency?IsoCurrencyCd="+info.IsoCurrencyCd+"&r="+Math.random()*10000000000000,
		            error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
//		                    toastr.success(rs.ErrorMsg);
		                    swal("删除!", "币种信息已删除", "success");
//							currs.ajax.reload(null, true);
							currs.ajax.url('/platform/ComCurrency?r='+Math.random()*10000000000000).load();
		                }else {
		                    toastr.error(rs.ErrorMsg);
		                    return ;
		                }
		            }
		        });
			});
		}
		
		function search(){
//			currs.ajax.reload(null, true);
			currs.ajax.url('/platform/ComCurrency?r='+Math.random()*10000000000000).load();
		}
	</script>

</html>