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

		<title>整体试算</title>

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
		<link rel="stylesheet" type="text/css" href="/static/inspinia/css/plugins/steps/jquery.steps.css" />
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
		<script src="/static/inspinia/js/plugins/staps/jquery.steps.min.js" type="text/javascript" charset="utf-8"></script>
		<!-- Data picker -->
		<script src="/static/inspinia/js/plugins/datapicker/bootstrap-datepicker.js"></script>
		<script src="/static/inspinia/js/jquery.form.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/bootstrap-filestyle.min.js"></script>
		<script src="/static/inspinia/js/excellentexport.min.js" type="text/javascript" charset="utf-8"></script>
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
			
			hr {
				margin: 5px 0px;
			}
			
			.base .form-group {
				width: 20%;
				float: left;
				padding: 5px 10px;
			}
			
			.paymentInfo .form-group,
			.repeatInfo .form-group {
				width: 25%;
				float: left;
				padding: 5px 10px;
			}
			
			h3 {
				clear: both;
			}
			
			fieldset {
				border: 1px dashed #CCC;
				padding: 15px 10px;
				margin: 20px 0px;
			}
			
			fieldset legend {
				margin-bottom: 5px;
			}
			
			.input-group-addon {
				text-align: left;
			}
			
			.wizard .content {
				background: rgba(245, 245, 245, 0.69) !important;
			}
			
			.wizard .content .body {
				width: 100% !important;
				height: 100% !important;
				padding: 10px !important;
			}
			
			.form-horizontal label {
				padding: 2px;
				/*width: 80px;*/
				/*text-align: center !important;*/
			}
			
			.form-horizontal div {
				padding: 0px;
			}
			
			tbody tr td {
				text-align: center;
				vertical-align: middle !important;
			}
			/*.noPadding{
				padding: 0px !important;
				height: 33px !important;
			}
			.tdInput{
				height: 33px;
				width: 100%;
			}
			
			tr td:first-child{
				width: 100px;
				text-align: center;
			}*/
			
			tbody td {
				padding: 2px !important;
			}
			
			.wizard> .steps a,
			.wizard> .steps a:hover,
			.wizard> .steps a:active {
				background: #eee;
				color: #aaa;
				cursor: default;
			}
			
			.step-content {
				overflow-y: auto;
				overflow-x: hidden;
			}
			
			th,
			td {
				text-align: center;
			}
			
			.cd tbody tr td,
			.qh tbody tr td {
				width: 50%;
			}
			.modal-header .close{
				padding-top: 4px !important;
				padding-right: 8px !important;
			}
		</style>
	</head>

	<body>
		<div id="wrapper" style="height: 100%;">
			<div class="wrapper wrapper-content animated fadeInRight" style="height: 100%;padding: 20px 20px;">
				<div class="row" style="height: 100%;">
					<div class="ibox float-e-margins">
						<div class="ibox-title" style="padding-left: 20px;">
							<div class="form-horizontal">
								<!--<div class="col-sm-9">
									<label class="control-label" style="text-align: center;float: left;width: 50px;">币种：</label>
									<div class="col-sm-2">
										<select class="form-control" name="Iso_currency_cd" id="curreny">
										</select>
									</div>
								</div>-->
								<div class="col-sm-9" style="padding-left: 8px;">
								<!--<div style="margin-left: 5px;">-->
									<button resid="30208040201000" class="btn btn-primary res " type="button" onclick="showUploadDialog()"> <i class="fa fa-upload"></i> 手工计算</button>
									<button resid="30208040202000" class="btn btn-info res " type="button" onclick="sysCalc()"> <i class="fa fa-plus"></i> 系统计算</button>
									<!--<a class="btn btn-warning" id="export" href="#" style="display: none;">导出</a>-->
								</div>
								<div class="col-sm-3">
									<a class="btn btn-warning pull-right" id="export" href="#" style="display: none;">导出</a>
								</div>
							</div>
						</div>
						<div class="ibox-content" style="overflow-x: auto;">
							<div id="wizard" style="height: 100%;">
								<h1>资产业务FTP利差分析</h1>
								<div class="step-content">
									<div class="wait">等待计算中。。。。。</div>
									<div class="row">
										<table class="table table-striped table-bordered table-hover zcyw single" id="zcyw">
											<thead>
												<tr>
													<td>统计项</td>
													<td>原始期限</td>
													<td>业务单元编码</td>
													<td>存量余额</td>
													<td>占比</td>
													<td>加权利率</td>
													<td>收入</td>
													<td>FTP价格</td>
													<td>司库还原</td>
													<td>调节后FTP价格</td>
													<td>FTP利差</td>
													<td>调整前FTP成本</td>
													<td>调整后FTP成本</td>
													<td>FTP利润</td>
												</tr>
											</thead>
											<tbody>
												<tr class='loan type1'>
													<td rowspan="6">贷款</td>
													<td>1Y以内</td>
													<td rowspan="6" class="busiz"></td>
												</tr>
												<tr class='loan'>
													<td>2Y</td>
												</tr>
												<tr class='loan'>
													<td>3Y</td>
												</tr>
												<tr class='loan'>
													<td>5Y</td>
												</tr>
												<tr class='loan'>
													<td>10Y</td>
												</tr>
												<tr class='loan'>
													<td>10Y以上</td>
												</tr>

												<tr class='zt'>
													<td rowspan="9">直贴</td>
													<td>1天</td>
													<td rowspan="9" class="busiz"></td>
												</tr>
												<tr class='zt'>
													<td>7天</td>
												</tr>
												<tr class='zt'>
													<td>14天</td>
												</tr>
												<tr class='zt'>
													<td>1个月</td>
												</tr>
												<tr class='zt'>
													<td>2个月</td>
												</tr>
												<tr class='zt'>
													<td>3个月</td>
												</tr>
												<tr class='zt'>
													<td>4个月</td>
												</tr>
												<tr class='zt'>
													<td>5个月</td>
												</tr>
												<tr class='zt'>
													<td>6个月</td>
												</tr>

												<tr class='ztx'>
													<td rowspan="9">转贴现</td>
													<td>1天</td>
													<td rowspan="9" class="busiz"></td>
												</tr>
												<tr class='ztx'>
													<td>7天</td>
												</tr>
												<tr class='ztx'>
													<td>14天</td>
												</tr>
												<tr class='ztx'>
													<td>1个月</td>
												</tr>
												<tr class='ztx'>
													<td>2个月</td>
												</tr>
												<tr class='ztx'>
													<td>3个月</td>
												</tr>
												<tr class='ztx'>
													<td>4个月</td>
												</tr>
												<tr class='ztx'>
													<td>5个月</td>
												</tr>
												<tr class='ztx'>
													<td>6个月</td>
												</tr>

												<tr class='jyx'>
													<td colspan="2">交易性金融资产</td>
													<td class="busiz"></td>
												</tr>

												<tr class='kgcs'>
													<td colspan="2">可供出售金融资产</td>
													<td class="busiz"></td>
												</tr>

												<tr class='qt'>
													<td rowspan="12">其他资金资产</td>
													<td>活期(1天)</td>
													<td rowspan="12" class="busiz"></td>
												</tr>
												<tr class='qt'>
													<td>7天</td>
												</tr>
												<tr class='qt'>
													<td>14天</td>
												</tr>
												<tr class='qt'>
													<td>1个月</td>
												</tr>
												<tr class='qt'>
													<td>2个月</td>
												</tr>
												<tr class='qt'>
													<td>3个月</td>
												</tr>
												<tr class='qt'>
													<td>6个月</td>
												</tr>
												<tr class='qt'>
													<td>1年</td>
												</tr>
												<tr class='qt'>
													<td>2年</td>
												</tr>
												<tr class='qt'>
													<td>3年</td>
												</tr>
												<tr class='qt'>
													<td>5年</td>
												</tr>
												<tr class='qt'>
													<td>5年以上</td>
												</tr>
											</tbody>
											<tfoot>
												<tr>
													<td>统计项</td>
													<td>原始期限</td>
													<td>业务单元编码</td>
													<td>存量余额</td>
													<td>占比</td>
													<td>加权利率</td>
													<td>收入</td>
													<td>FTP价格</td>
													<td>司库还原</td>
													<td>调节后FTP价格</td>
													<td>FTP利差</td>
													<td>调整前FTP成本</td>
													<td>调整后FTP成本</td>
													<td>FTP利润</td>
												</tr>
											</tfoot>
										</table>
									</div>
								</div>

								<h1>负债业务FTP利差分析</h1>
								<div class="step-content">
									<div class="wait">等待计算中。。。。。</div>
									<div class="row">
										<table class="table table-striped table-bordered table-hover fzyw single">
											<thead>
												<tr>
													<td>统计项</td>
													<td>原始期限</td>
													<td>业务单元编码</td>
													<td>存量余额</td>
													<td>占比</td>
													<td>加权利率</td>
													<td>成本</td>
													<td>FTP价格</td>
													<td>准备金</td>
													<td>调节后FTP价格</td>
													<td>FTP利差</td>
													<td>调整前FTP收入</td>
													<td>调整后FTP收入</td>
													<td>FTP利润</td>
												</tr>
											</thead>
											<tbody>
												<tr class='ck'>
													<td rowspan="8">存款</td>
													<td>活期</td>
													<td rowspan="" class="busiz"></td>
												</tr>
												<!--<tr class='ck'>
													<td>通知存款</td>
												</tr>-->
												<tr class='ck'>
													<td>3个月</td>
													<td rowspan="7" class="busiz"></td>
												</tr>
												<tr class='ck'>
													<td>6个月</td>
												</tr>
												<tr class='ck'>
													<td>1年</td>
												</tr>
												<tr class='ck'>
													<td>2年</td>
												</tr>
												<tr class='ck'>
													<td>3年</td>
												</tr>
												<tr class='ck'>
													<td>5年</td>
												</tr>
												<tr class='ck'>
													<td>5年以上</td>
												</tr>

												<tr class='zjfz'>
													<td rowspan="12">资金负债</td>
													<td>活期(1天)</td>
													<td rowspan="12" class="busiz"></td>
												</tr>
												<tr class='zjfz'>
													<td>7天</td>
												</tr>
												<tr class='zjfz'>
													<td>14天</td>
												</tr>
												<tr class='zjfz'>
													<td>1个月</td>
												</tr>
												<tr class='zjfz'>
													<td>2个月</td>
												</tr>
												<tr class='zjfz'>
													<td>3个月</td>
												</tr>
												<tr class='zjfz'>
													<td>6个月</td>
												</tr>
												<tr class='zjfz'>
													<td>1年</td>
												</tr>
												<tr class='zjfz'>
													<td>2年</td>
												</tr>
												<tr class='zjfz'>
													<td>3年</td>
												</tr>
												<tr class='zjfz'>
													<td>5年</td>
												</tr>
												<tr class='zjfz'>
													<td>5年以上</td>
												</tr>
											</tbody>
											<tfoot>
												<tr>
													<td>统计项</td>
													<td>原始期限</td>
													<td>业务单元编码</td>
													<td>存量余额</td>
													<td>占比</td>
													<td>加权利率</td>
													<td>成本</td>
													<td>FTP价格</td>
													<td>准备金</td>
													<td>调节后FTP价格</td>
													<td>FTP利差</td>
													<td>调整前FTP收入</td>
													<td>调整后FTP收入</td>
													<td>FTP利润</td>
												</tr>
											</tfoot>
										</table>
									</div>
								</div>

								<h1>合计结果</h1>
								<div class="step-content">
									<div class="wait">等待计算中。。。。。</div>
									<div class="row">
										<div class="col-lg-12">
											<h3>1、各项统计结果</h3>
										</div>
										<div class="col-lg-12">
											<table class="table table-striped table-bordered table-hover total detail">
												<thead>
													<tr>
														<td>统计项</td>
														<td>存量余额</td>
														<td>外部利率</td>
														<td>外部利息</td>
														<td>FTP价格</td>
														<td>调节后FTP价格</td>
														<td>FTP利差</td>
														<td>调整前FTP成本</td>
														<td>调整后FTP成本</td>
														<td>FTP利润</td>
													</tr>
												</thead>
												<tbody>
													<tr class="dt">
														<td>贷款小计</td>
													</tr>
													<tr class="dt">
														<td>直贴合计</td>
													</tr>
													<tr class="dt">
														<td>转贴现合计</td>
													</tr>
													<tr class="dt">
														<td>交易性金融资产</td>
													</tr>
													<tr class="dt">
														<td>可供出售金融资产</td>
													</tr>
													<tr class="dt">
														<td>其他资金资产合计</td>
													</tr>
													<tr class="dt">
														<td>存款合计</td>
													</tr>
													<tr class="dt">
														<td>资金负债合计</td>
													</tr>
												</tbody>
											</table>
										</div>
										<div class="col-lg-12">
											<h3>2、统计项合计结果</h3>
										</div>
										<div class="col-lg-12">
											<table class="table table-striped table-bordered table-hover total all">
												<thead>
													<tr>
														<td>统计项</td>
														<td>存量余额</td>
														<td>外部利率</td>
														<td>外部利息</td>
														<td>FTP价格</td>
														<td>调节后FTP价格</td>
														<td>FTP利差</td>
														<td>调整前FTP成本</td>
														<td>调整后FTP成本</td>
														<td>FTP利润</td>
													</tr>
												</thead>
												<tbody>
													<tr class="t">
														<td>所有资产合计</td>
													</tr>
													<tr class="t">
														<td>所有负债合计</td>
													</tr>
												</tbody>
											</table>
										</div>
										<div class="col-lg-12">
											<h3>3、利差分析</h3>
										</div>
										<div class="col-lg-6">
											<table class="table table-striped table-bordered table-hover cd">
												<thead>
													<tr>
														<td colspan="2">存贷利差切割分析</td>
													</tr>
												</thead>
												<tbody>
													<tr>
														<td>存贷利差</td>
														<td id="cdlc">0</td>
													</tr>
													<tr>
														<td>贷款利差</td>
														<td id="dklc">0</td>
													</tr>
													<tr>
														<td>存款利差</td>
														<td id="cklc">0</td>
													</tr>
													<tr>
														<td>存贷错配利差</td>
														<td id="cdcplc">0</td>
													</tr>
												</tbody>
											</table>
										</div>
										<div class="col-lg-6">
											<table class="table table-striped table-bordered table-hover qh">
												<thead>
													<tr>
														<td colspan="2">生息资产和付息负债利差切割分析</td>
													</tr>
												</thead>
												<tbody>
													<tr>
														<td>全行利差</td>
														<td id="qhlc">0</td>
													</tr>
													<tr>
														<td>资产利差</td>
														<td id="zclc">0</td>
													</tr>
													<tr>
														<td>负债利差</td>
														<td id="fzlc">0</td>
													</tr>
													<tr>
														<td>全行错配利差</td>
														<td id="qhcplc">0</td>
													</tr>
												</tbody>
											</table>
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>

		<div class="modal inmodal" id="uploadDialog" tabindex="-2" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2>手工计算数据导入</h2>
					</div>
					<div class="modal-body">
						<div class="row">
							<form id="uploadForm" enctype="multipart/form-data" method="post">
								<input type="hidden" id="curr" name="IsoCurrencyCd" />
								<div class="col-xs-8">
									<div class="form-group">
										<input type="file" id="uploadFile" name="uploadFile">
									</div>
								</div>
								<div class="col-xs-4" style="padding-left: 0px;">
									<button class="btn btn-info" type="submit">提交</button>
									<a class="btn btn-link" style="text-decoration: underline;" href="/updownload/EnsembleCalc.xlsx">下载模板</a>
								</div>
							</form>
						</div>
					</div>
				</div>
			</div>
		</div>

		<!--<div class="modal inmodal" id="sysCalcDialog" tabindex="-2" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2>系统计算--选择业务单元</h2>
					</div>
					<div class="modal-body">
						<div class="row">
							<form id="sysCalc">
								<div class="form-group col-sm-6">
									<label>贷款:</label>
									<select class="form-control" name="Adj_id" id="adjId">
									</select>
								</div>
								<div class="form-group col-sm-6">
									<label>直贴:</label>
									<select class="form-control" name="Adj_id" id="adjId">
									</select>
								</div>
								<div class="form-group col-sm-6">
									<label>转贴现:</label>
									<select class="form-control" name="Adj_id" id="adjId">
									</select>
								</div>
								<div class="form-group col-sm-6">
									<label>交易性金融资产:</label>
									<select class="form-control" name="Adj_id" id="adjId">
									</select>
								</div>
								<div class="form-group col-sm-6">
									<label>可供出售金融资产:</label>
									<select class="form-control" name="Adj_id" id="adjId">
									</select>
								</div>
								<div class="form-group col-sm-6">
									<label>其他资金资产:</label>
									<select class="form-control" name="Adj_id" id="adjId">
									</select>
								</div>
								<div class="form-group col-sm-6">
									<label>存款:</label>
									<select class="form-control" name="Adj_id" id="adjId">
									</select>
								</div>
								<div class="form-group col-sm-6">
									<label>资金负债:</label>
									<select class="form-control" name="Adj_id" id="adjId">
									</select>
								</div>
							</form>
						</div>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-primary" onclick="sysCalc()">计算</button>
					</div>
				</div>
			</div>
		</div>-->
	</body>
	<script type="text/javascript">
		$(function() {
			//控制按钮权限
			$.get('/platform/DefaultMenu', {
				TypeId: 2,
				Id: '30208040200000',
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

			$("#wizard").steps({
				headerTag: "h1",
				bodyTag: "div",
				transitionEffect: "slideLeft",
				autoFocus: true,
				enableAllSteps: true,
				enableFinishButton: false,
				enableCancelButton: false,
				labels: {
					next: "下一步",
					previous: "上一步",
					loading: "加载中 ..."
				}
			});

			$('.content').height($('#wrapper').height() - 238);

//			$.get('/platform/MasDimCurrency', {r:Math.random()*10000000000000}, function(data) {
//				appendOption('curreny', data, 'IsoCurrencyCd', 'IsoCurrencyName');
//			});

			initTable();

			$('#uploadFile').filestyle({
				buttonName: 'btn-warning',
				buttonText: '选择文件'
			});

			$('.step-content .row').hide();

			var options = {
				contentType: "application/x-www-form-urlencoded;charset=utf-8",
				url: '/platform/FtpEnsembleCala?r='+Math.random()*10000000000000,
				success: function(data) {
					var rs=JSON.parse(data);
					$.unblockUI();
					if(rs.ErrorCode === '1') {
						fillTableData(rs);
						$('#export').attr('href', rs.DfileName).show();
						toastr.info('点击导出可导出计算统计结果');
					} else {
						toastr.error(rs.ErrorMsg);
						return false;
					}
					
				},
				data: {
					IsoCurrencyCd: 'CNY'
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
							message: '<h1>上传计算中。。。</h1>',
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

			$.get('/platform/MyDomainID', {r:Math.random()*10000000000000}, function(data) {
				var rs = JSON.parse(data);
				$('#region').text(rs.DomainName);
			});
		});

		function showUploadDialog() {
			$('#curr').val($('#curreny').val());
			$('#uploadForm')[0].reset();
			$('#uploadDialog').modal('show');
		}

		function sysCalc() { //系统计算
			//			toastr.info('系统统计计算中....');
			$.blockUI({
				message: '<h1>系统计算中....</h1>',
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
			$.get('/platform/FtpEnsembleCalcAuto', {
				IsoCurrencyCd: 'CNY',
				r:Math.random()*10000000000000
			}, function(data) {
				var rs=JSON.parse(data);
				$.unblockUI();
				if(rs.ErrorCode === '1') {
					fillTableData(rs);
					$('#export').attr('href', rs.DfileName).show();
					toastr.info('点击导出可导出计算统计结果');
				} else {
					toastr.error(rs.ErrorMsg);
					return false;
				}
			});
		}

		function initTable() { //初始化结果表格
			//			appendTd(['loan', 'zt', 'ztx', 'jyx', 'kgcs', 'qt', 'ck', 'zjfz'], 11);
			//			appendTd(['dt', 't'], 9);
			appendSingleTd(['loan', 'zt', 'ztx', 'jyx', 'kgcs', 'qt', 'ck', 'zjfz']);
			appendResultTd(['dt', 't']);
		}

		function appendSingleTd(dists) {
			dists.forEach(function(dist) {
				$('.' + dist).each(function() {
					$(this).append("<td class='StockBalance'>0</td><td class='Ratio'>0</td><td class='WeightRate'>0</td><td class='Income'>0</td><td class='FtpValue'>0</td><td class='ProfitRe'>0</td><td class='AftpValue'>0</td><td class='FtpDiffer'>0</td><td class='BeforeFtpCost'>0</td><td class='AfterFtpCost'>0</td><td class='FtpProfit'>0</td>");
				});
			});
		}

		function appendResultTd(dists) {
			dists.forEach(function(dist) {
				$('.' + dist).each(function() {
					$(this).append("<td class='StockBalance'>0</td><td class='ExternTa'>0</td><td class='ExternTax'>0</td><td class='FtpValue'>0</td><td class='AftpValue'>0</td><td class='FtpDiffer'>0</td><td class='BeforeFtpCost'>0</td><td class='AfterFtpCost'>0</td><td class='FtpProfit'>0</td>");
				});
			});
		}

		//		function appendTd(dists, number){
		//			dists.forEach(function(dist){
		//				$('.'+dist).each(function(){
		//					for(var i=0; i<number; i++){
		//						$(this).append("<td>0</td>");
		//					}
		//				});
		//			});
		//		}

		function fillTableData(rs) { //填充表格结果
//			var rs = JSON.parse(data);
			if(rs.ErrorCode != '0') {
				var single = rs.Single, //单项
					total = rs.Result; //合计

				//资产负债业务
				var singleParmsAll = [];
				single.forEach(function(e) {
					$(".single td.busiz:eq(" + (e.Id - 1) + ")").text(e.Bid);
					Array.prototype.push.apply(singleParmsAll, e.Params);
				});
				singleParmsAll.forEach(function(d, index) {
					//					$("td.StockBalance:eq("+index+")").text(d.StockBalance);
					for(var key in d) {
						$(".single td." + key + ":eq(" + index + ")").text(d[key]);
					}
				});

				//合计
				total.forEach(function(d, index) {
					//					$("td.StockBalance:eq("+index+")").text(d.StockBalance);
					for(var key in d) {
						$(".total td." + key + ":eq(" + index + ")").text(d[key]);
					}
				});

				//利差分析
				var cd = $('#cdlc').text(($(".detail tbody tr:eq(0) td:eq(2)").text() - $(".detail tbody tr:eq(6) td:eq(2)").text()).toFixed(4)).text();
				var dk = $('#dklc').text($(".detail tbody tr:eq(0) td:eq(6)").text()).text();
				var ck = $('#cklc').text($(".detail tbody tr:eq(6) td:eq(6)").text()).text();
				$('#cdcplc').text((cd - dk - ck).toFixed(4));

				var qh = $('#qhlc').text(($(".all tbody tr:eq(0) td:eq(2)").text() - $(".all tbody tr:eq(1) td:eq(2)").text()).toFixed(4)).text();
				var zc = $('#zclc').text($(".all tbody tr:eq(0) td:eq(6)").text()).text();
				var fz = $('#fzlc').text($(".all tbody tr:eq(1) td:eq(6)").text()).text();
				$('#qhcplc').text((qh - zc - fz).toFixed(4));
				toastr.success('计算完成');
				$('.step-content .row').show();
				$('.step-content .wait').hide();
			} else {
				toastr.warning(rs.ErrorMsg);
			}
		}
	</script>

</html>