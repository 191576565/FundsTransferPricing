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

		<title>曲线定义</title>

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
				width: 60px;
			}
			
			.selected {
				background-color: #b7eafb !important;
			}
			.input-group-addon{
				border: 1px solid #E5E6E7 !important;
			}
			tr td:first-child{
				text-align: center;
				width: 40px !important;
			}
		</style>
	</head>

	<body>
		<div id="wrapper">
			<div class="wrapper wrapper-content animated fadeInRight">
				<div class="row">
					<div class="col-lg-12" id="cvInfos" style="display: none;">
						<div class="ibox float-e-margins" style="border: 1px solid #dcdada;">
							<div class="ibox-title" style="border: 0px;">
								<div class="ibox-tools">
									<h3 style="float: left;margin-top: 0px;" id="curveTitle"></h3>
									<input type="hidden" name="Curve_id" value="" />
									<input type="hidden" name="Domain_id" value="" />
									<a class="" style="float: right;" onclick="closeCurveWindow()">
										<i class="fa fa-times"></i>
									</a>
								</div>
							</div>
							<div class="ibox-content">
								<div class="row">
									<div class="col-xs-4">
										<button resid="208010104010" class="btn btn-primary res " type="button" onclick="showCVSDialog('add')"> <i class="fa fa-plus"></i> 新增</button>
										<button resid="208010104020" class="btn btn-info editBtn res" type="button" onclick="showCVSDialog('edit')"><i class="fa fa-paste"></i> 编辑</button>
										<button resid="208010104030" class="btn btn-danger res" type="button" onclick="deleteCVS()"><i class="fa fa-times"></i> 删除</button>
									</div>
									<div class="col-xs-8 text-right">
										<div class="form-inline">
											<div class="input-group date">
												<input type="text" class="form-control input-group-addon" placeholder="起始日期" id="startDay">
											</div> -->
											<div class="input-group date">
												<input type="text" class="form-control input-group-addon" placeholder="结束日期" id="endDay">
											</div>
											<button class="btn btn-default searchBtn" onclick="searchC('2')" type="button" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
										</div>
									</div>
								</div>
								<div class="table-responsive">
									<table class="table table-striped table-bordered table-hover cvs" style="width: 99%;">
										<thead>
											<tr>
											</tr>
										</thead>
									</table>
								</div>
							</div>
						</div>
					</div>

					<div class="col-xs-12" id="curveInfos">
						<div class="ibox float-e-margins">
							<div class="ibox-title">
								<div class="row">
									<div class="col-xs-6">
										<button resid="208010101000" class="btn btn-primary res" type="button" onclick="showEditCuverDialog('add')"> <i class="fa fa-plus"></i> 新增</button>
										<button resid="208010102000" class="btn btn-info editBtn res" type="button" onclick="showEditCuverDialog('edit')"><i class="fa fa-paste"></i> 编辑</button>
										<button resid="208010103000" class="btn btn-danger deleteBtn res" type="button"><i class="fa fa-times"></i> 删除</button>
										<button resid="208010104000" class="btn btn-warning infoBtn res" type="button" onclick="showCurveValueData()"><i class="fa fa-flag"></i> 查看曲线点值</button>
									</div>
									<div class="col-xs-6 text-right">
										<div class="form-inline">
											<div class="form-group">
												<input type="text" placeholder="请输入曲线编号或名称" style="width: 200px;" id="keyword" class="form-control">
											</div>
											<button class="btn btn-default searchBtn" type="button" onclick="searchC('1')" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
										</div>
									</div>
								</div>

							</div>
							<div class="ibox-content">
								<div class="table-responsive">
									<table class="table table-striped table-bordered table-hover curves"  style="width: 99%;">
										<thead>
											<tr>
												<th style="width: 40px;">
												</th>
												<th>曲线编号</th>
												<th>曲线名称</th>
												<th>币种名称</th>
												<th>曲线类型</th>
												<th>创建日期</th>
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

		<div class="modal inmodal" id="editCDialog" tabindex="-1" role="dialog" aria-hidden="false">
			<div class="modal-dialog modal-lg">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 id="addCurve">新增曲线信息</h2>
						<h2 id="editCurve">编辑曲线信息</h2>
						<input type="hidden" id="type" value="add" />
					</div>
					<div class="modal-body">
						<div class="row">
							<div class="col-sm-12 b-r">
								<form role="form" id="curveForm">
									<div class="form-group col-sm-6">
										<label>曲线编号:</label>
										<input type="number" placeholder="" class="form-control" name="Curve_id" id="curveCode">
									</div>
									<div class="form-group col-sm-6">
										<label>曲线名称:</label>
										<input type="text" placeholder="" class="form-control" name="Curve_desc" id="curveName">
									</div>
									<div class="form-group col-sm-6">
										<label>曲线类型:</label>
										<select class="form-control m-b" name="Curve_type" id="Curve_type" onchange="inputOnchange(this)">
										</select>
									</div>
									<div class="form-group col-sm-6">
										<label>币种:</label>
										<select class="form-control m-b" name="Iso_currency_cd" id="curveCurr">
										</select>
									</div>
									<div class="form-group col-sm-6" style="display: none;">
										<label>重定价频率:</label>
										<select class="form-control m-b" name="" id="repeat">
										</select>
									</div>
									<div class="form-group col-sm-12">
										<label class="" style="padding: 0px 0px;">曲线期限选择:</label>
										<div class="termCheck">
										</div>
									</div>
								</form>
							</div>
						</div>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="saveCurveInfo()">保存</button>
					</div>
				</div>
			</div>
		</div>

		<div class="modal inmodal" id="cvsDialog" tabindex="-1" role="dialog" aria-hidden="true">
			<div class="modal-dialog modal-lg">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 id="cvsTitle"></h2>
					</div>
					<div class="modal-body">
						<input type="hidden" value="" name="type" id="opType">
						<form role="form" id="cvsForm">
							<div class="row">
								<div class="col-sm-3 dataDate">
									<div class="form-group">
										<input type="hidden" id="dataDate" name="dataDate"/>
										<label>数据日期:</label>
										<input type="text" class="form-control input-group-addon" placeholder="" id="day" name="day" onchange="checkDate(this)">
									</div>
								</div>
							</div>
						</form>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="saveCVSInfo()">保存</button>
					</div>
				</div>
			</div>
		</div>

	</body>
	<script type="text/javascript">
		var curves, curveInfos, cvs;
		$(function() {
			//控制按钮权限
			$.get('/platform/DefaultMenu',{TypeId:2,Id:'208010100000',r:Math.random()*10000000000000}, function(data){
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
			
			
			$('#cvInfos .ibox').height(document.body.clientHeight-100);
			
			$('.input-group-addon').datepicker({
				autoclose: true,
				format: 'yyyy-mm-dd'
			});

			$.ajaxSetup({
				async: false,
				cache:false
			});

			initCurveTable();

			$('.curves tbody').on('click', 'tr', function() {
//				if($(this).find('input:radio').get(0).checked) {
//					$(this).find('input:radio').get(0).checked = false;
//				} else {
					$(this).find('input:radio').get(0).checked = true;
//				}
				//				$(this).toggleClass('selected');
			});

			$.get('/platform/FtpCurveSave', {r:Math.random()*10000000000000}, function(data) {
				var rs = JSON.parse(data);
				var ap = "";
				$('.termCheck').html('');
				if(rs!=null){
					for(var i = 0; i < rs.length; i++) {
						r = rs[i].Struct_code;
						ap = "<label class='checkbox-inline i-checks'> <input type='checkbox' value='" + r + "' name='" + r + "'>" + r + "</label>"
						$('.termCheck').append(ap);
						ap = "";
					}
				}
			});

			iCheckInit();

			$.get('/platform/MasDimCurrency', {r:Math.random()*10000000000000}, function(data) {
				appendOption('curveCurr', data, 'IsoCurrencyCd', 'IsoCurrencyName');
			});
			
			$.get('/platform/FtpCurveType', {r:Math.random()*10000000000000}, function(data) {
				appendOption('Curve_type', data, 'CuType', 'CuTypeDesc');
			});

			$.get('/platform/FtpRepType', {r:Math.random()*10000000000000}, function(data){
				appendOption('repeat', data, 'FtpRepId', 'FtpRepDesc');
			});

//			$.get('/platform/SysDomainInfo', function(data) {
//				appendOption('curveRegion', data, 'DomainId', 'DomainName');
//				$('#curveRegion').val('FTP').change();
//			});

			$('.deleteBtn').click(function() {
				if(!$('.curves input:checked').val()) {
					toastr.warning('请选择曲线信息');
					return;
				}
				var info = curves.data()[$('.curves input:checked').val()];
				swal({
					title: "删除！",
					text: "是否删除该条曲线信息？",
					type: "warning",
					showCancelButton: true,
					confirmButtonColor: "#DD6B55",
					//			        confirmButtonText: "",
					closeOnConfirm: false
				}, function() {
					$.ajax({ //删除曲线
						type: "delete",
						url: "/platform/FtpCurveSave?curveCode=" + info.Curve_id + "&Domain_id=" + info.Domain_id+"&r="+Math.random()*10000000000000,
						error: function(msg) {
							console.log(msg.responseText);
						},
						success: function(data) {
							var rs = JSON.parse(data);
							if(rs.ErrorCode === '1') {
								//刷新表格
//								curves.ajax.reload(null, true);
								curves.ajax.url('/platform/FtpCurveDef?r='+Math.random()*10000000000000).load();
								swal("删除!", "曲线已删除", "success");
							} else {
								toastr.error(rs.ErrorMsg);
								return;
							}
						}
					});
				});
			});
		});
		
		function iCheckInit(){
			$('.i-checks').iCheck({
				checkboxClass: 'icheckbox_square-green',
				radioClass: 'iradio_square-green',
			});
		}

		function showEditCuverDialog(type) {
			$('#type').val(type);
			//取消所有checkbox勾选
			$('.termCheck input:checkbox').parent().iCheck('uncheck');
			if(type === 'add') { //新增
				$('#addCurve').show();
				$('#editCurve').hide();
				$('#curveCode').attr('disabled', false);
//				$('#curveRegion').attr('disabled', false);
				$('#curveForm').get(0).reset();
				
			} else if(type === 'edit') { //编辑
				$('#addCurve').hide();
				$('#editCurve').show();
				$('#curveCode').attr('disabled', true);
//				$('#curveRegion').attr('disabled', true);

				if(!$('.curves input:checked').val()) {
					toastr.warning('请选择曲线信息');
					return;
				}

				//赋值及勾选已有点
				//          		var info=curveInfos[$('.curves input:checked').val()];
				var info = curves.data()[$('.curves input:checked').val()];

				$('#curveCode').val(info.Curve_id);
				$('#curveName').val(info.Curve_desc);
				$('#curveCurr').val(info.Iso_currency_cd);
				$('#Curve_type').val(info.Curve_type);
//				$('#curveRegion').val(info.Domain_id);
				$('#repeat').val(info.Rep_id);

				var spots = info.All_Struct_Code.split(',');
				for(var i = 0; i < spots.length; i++) {
					$(".termCheck input[name=" + spots[i] + "]").parent().iCheck('check');
				}
			}
			$('#Curve_type').change();
			$('#editCDialog').modal('show');
		}

		function saveCurveInfo() {
			//保存
			var code = $('#curveCode').val();
			var name = $('#curveName').val();
			var curr = $('#curveCurr').val();
			var region = $('#curveRegion').val();
			var type = $('#Curve_type').val();
			var rep = $('#repeat').val();

			var terms = '';
			$(".termCheck .checked").each(function() {
				terms = terms + $(this).find('input').val() + ',';
			});

			//输入校验
			if((code === "" || code.indexOf('-') >= 0 || code.indexOf('.') >= 0) || name === "") {
				toastr.warning('请正确输入曲线编号及名称');
				return;
			}

			if(terms === '') {
				toastr.warning('请选择至少一个曲线期限');
				return;
			}
			$.post('/platform/FtpCurveSave', {
				Curve_id: code,
				Curve_desc: name,
				Iso_currency_cd: curr,
				Domain_id: region,
				terms: terms.substr(0, terms.length - 1),
				Curve_type:type,
				type: $('#type').val(),
				Rep_Id:rep,
				r:Math.random()*10000000000000
			}, function(data) {
				var rs = JSON.parse(data);
				if(rs.ErrorCode === '1') {
					toastr.success('曲线信息保存成功！');
					//刷新表格
//					curves.ajax.reload(null, true);
					curves.ajax.url('/platform/FtpCurveDef?r='+Math.random()*10000000000000).load();
					$('#editCDialog').modal('hide');
				} else {
					toastr.error(rs.ErrorMsg);
					return;
				}
			});
		}

		function initCurveTable() {
			curves = $('.curves').DataTable({
				sAjaxSource: '/platform/FtpCurveDef?r='+Math.random()*10000000000000,
				bPaginate: false,
				info: false,
				aoColumns: [{
					"data": null
				}, {
					"data": "Curve_id"
				}, {
					"data": "Curve_desc"
				}, {
					"data": "Iso_currency_desc"
				},{
					"data":"Curve_type_desc"
				}, {
					"data": "Create_date",
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
						value: $('#keyword').val()
					});
				}
			});
		}

		function showCurveValueData() { //显示曲线的点值表格
			if(!$('.curves input:checked').val()) {
				toastr.warning('请选择曲线信息');
				return;
			}
//			$('#startDay').val('');
//			$('#endDay').val('');
			$('#startDay').datepicker('update', '');
			$('#endDay').datepicker('update', '');

			var info = curves.data()[$('.curves input:checked').val()];
			var columns = [{
				"data": null
			}, {
				"data": 'As_of_date'
			}];
			//表头
			$.get('/platform/FtpCurveInfoPage', {
//				DomainId: info.Domain_id,
				CurveId: info.Curve_id,
				r:Math.random()*10000000000000
			}, function(data) {
				$('.cvs thead tr').html("<th style='width: 40px;'></th><th>数据日期</th>");
				var rs = JSON.parse(data);
				rs.Cstruct.forEach(function(e) {
					$('.cvs thead tr').append("<th>" + e.Struct_code + "</th>");
					columns.push({
						"data": e.Struct_code
					});
				});
				$('.cvs tbody').remove();
//				$('#startDay').val('');
//				$('#end').val('');
				$('#startDay').datepicker('update', '');
				$('#endDay').datepicker('update', '');
			});

			cvs = $('.cvs').DataTable({
				sAjaxSource: '/platform/FtpCurveInfo?r='+Math.random()*10000000000000,
				aoColumns: columns,
				iDisplayLength: 10,
				bDestroy: true,
				fnRowCallback: function(nRow, aData, iDisplayIndex, iDisplayIndexFull) {
					$('td:eq(0)', nRow).html('<div class="checkbox"><input type="checkbox" value="' + iDisplayIndex + '" name="tableCheckbox"><label></label></div>');
				},
				fnServerParams: function(aoData) {
					aoData.push({
						name: "CurveId",
						value: info.Curve_id
					});
//					aoData.push({
//						name: "DomainId",
//						value: info.Domain_id
//					});
					aoData.push({
						name: "StartDate",
						value: $('#startDay').val()
					});
					aoData.push({
						name: "EndDate",
						value: $('#endDay').val()
					});
				}
			});

//			$('.cvs tbody').on('click', 'tr', function() {
////				if($(this).find('input:checkbox').get(0).checked){
////					$(this).find('input:checkbox').get(0).checked=false;
////				}else {
////					$(this).find('input:checkbox').get(0).checked=true;
////				}
//				$(this).find('input:checkbox').get(0).checked=!$(this).find('input:checkbox').get(0).checked;
////				$(this).find('input:checkbox').attr('checked', !$(this).find('input:checkbox').attr('checked'));
//			});
			
			$('#curveTitle').text(info.Curve_desc);
			$("input[name='Curve_id']").val(info.Curve_id);
			$("input[name='Domain_id']").val(info.Domain_id);
			$('#cvInfos').slideDown();
			$('#curveInfos').fadeOut();
		}

		function closeCurveWindow() {
			$('#cvInfos').slideUp();
			$('#curveInfos').fadeIn();
		}

		function searchC(type) { //搜索
			if(type === '1') {
//				curves.ajax.reload(null, true);
				curves.ajax.url('/platform/FtpCurveDef?r='+Math.random()*10000000000000).load();
			} else if(type === '2') {
//				cvs.ajax.reload(null, true);
				cvs.ajax.url('/platform/FtpCurveInfo?r='+Math.random()*10000000000000).load();
			}
		}

		function showCVSDialog(type) {
			$('.dataDate').nextAll().remove();
			$('.cvs thead tr th:gt(1)').each(function(index){
				var temp="<div class='col-sm-3 dataDate'><div class='form-group'><label>"+$(this).text()+"</label><input type='text' placeholder='' class='form-control' name='"+$(this).text()+"'></div></div>";
				if(index%4===3){
					$('.dataDate').parent().append("<div style='clear:both'></div>");
				}
				$('.dataDate').parent().append(temp);
			});
			
			if(type === 'add') {
				$('#cvsTitle').text('新增曲线日期点值');
				$('#opType').val('add');
				$('#cvsForm')[0].reset();
				
				$('#day').removeAttr('disabled');
				$('#day').datepicker('update', '');
			} else if(type === 'edit') {
				$('#cvsTitle').text('编辑曲线日期点值');
				$('#opType').val('edit');
				$('#day').attr('disabled','disabled');
				 
				if($('.cvs :checked').length==0){
					toastr.warning('请选择一行曲线点值信息进行编辑');
					return ;
				}
				var info=cvs.data()[$('.cvs :checked')[0].value];
				$("#cvsForm input[name=day]").val((info.As_of_date).substr(0, 10));
	            $("#cvsForm input[name!=day]").each(function(data){
	            		$(this).val(info[$(this).attr('name')]);
	            });
	            $('#dataDate').val($('#day').val());
			}
			
			$('#cvsDialog').modal('show');
		}
		
		function saveCVSInfo(){ //需求修改点值不能为空
//			$('#startDay').val('');
//			$('#endDay').val('');
			$('#day').attr('disabled','disabled');
			
			var parms=$('#cvsForm').serialize().split('&');
			var day=parms[0].split('=')[1]; //数据日期
			var Curve_Id=$("input[name='Curve_id']").val();
    			var Domain_Id=$("input[name='Domain_id']").val();
    			var spots=[];
    			var flag=0;
    			var p=parms.splice(1, parms.length-1);
    			
    			var re = /^[0-9]+\.?[0-9]*$/;
    			p.forEach(function(e){
				var spot=e.split('=');
				var temp={};
				temp.StructCode=spot[0];
				temp.StructValue=spot[1];
				if(spot[1]!=''){ //不为空时需要是小于35的正数
					if(spot[1].indexOf('-')==-1 && re.test(spot[1])){ //正数
						var s=parseFloat(spot[1]);
						if(s>0 && s<35){
							flag=flag+1;	
						}
					}
				}
				spots.push(temp);
	    		});
	    		if(flag!=p.length){
	    			toastr.warning('曲线点值必填，为大于0小于35的数字');
	            return ;
	    		}
	    		
	    		$.post('/platform/FtpCurveInfoStruct', {r:Math.random()*10000000000000, Curve_Id:Curve_Id, Domain_Id:Domain_Id, date:day, type:$('#opType').val(), JSON:JSON.stringify(spots)}, function(data){
	    			var rs=JSON.parse(data);
	            if(rs.ErrorCode==='1'){
//	                cvs.ajax.reload(null, true);
					cvs.ajax.url('/platform/FtpCurveInfo?r='+Math.random()*10000000000000).load();
	                toastr.success('曲线点值保存成功！');
	                $('#cvsDialog').modal('hide');
	            }else {
	            		toastr.error(rs.ErrorMsg);
	            		$('#day').removeAttr('disabled');
	                return false;
	            }
	    		});
		}
		
		function checkDate(obj){
	    		if(new Date($(obj).val())>new Date()){
	            toastr.warning('请选择今天及以前的日期');
	            $(obj).val('');
	            return ;
	    		}
	    		$('#dataDate').val($(obj).val());
	    }
		
		function deleteCVS(){
			if($('.cvs :checked').length==0){
				toastr.warning('请选择一行曲线点值信息进行删除');
				return ;
			}
			
			var days=[];
			$('.cvs input:checked').each(function(){
				var info=cvs.data()[$(this).val()];
				days.push(info.As_of_date);
			});
			var json={Domain_id:$("input[name='Domain_id']").val(),Curve_id:$("input[name='Curve_id']").val(), date:days};
			
			swal({
				title: "删除！",
				text: "是否删除选中的曲线日期点值？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				//			        confirmButtonText: "",
				closeOnConfirm: false
			}, function() {
				$.ajax({
		            type:"delete",
		            url:"/platform/FtpCurveInfoStruct?JSON="+JSON.stringify(json)+"&r="+Math.random()*10000000000000,
		            error: function(msg){
		                console.log(msg.responseText);
		            },
		            success: function(data){
						var rs=JSON.parse(data);
		                if(rs.ErrorCode==='1'){
//		                    toastr.success(rs.ErrorMsg);
//		                    cvs.ajax.reload(null, true);
							cvs.ajax.url('/platform/FtpCurveInfo?r='+Math.random()*10000000000000).load();
		                    swal("删除!", "曲线已删除", "success");
		                }else {
		                    toastr.error('rs.ErrorMsg');
		                    return ;
		                }
		            }
		        });
			});
		}
		
		function inputOnchange(obj){
			var type=$(obj).val();	
			if(type==='1'){
				$('#repeat').parent().show();
			}else {
				$('#repeat').parent().hide();
			}
		}
	</script>

</html>