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

		<title>单笔试算</title>

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
			
			hr {
				margin: 5px 0px;
			}
			
			.base .form-group {
				width: 20%;
				float: left;
				padding: 5px 10px;
			}
			
			.paymentInfo .form-group, .repeatInfo .form-group{
				width: 25%;
				float: left;
				padding: 5px 10px;
			}
			h3{
				clear: both;
			}
			
			fieldset {
				border: 1px dashed #CCC;
				padding: 15px 10px;
				margin: 20px 0px;
			}
			
			fieldset legend{
				margin-bottom: 5px;
			}
			.input-group-addon{
				text-align: left;
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
					<div class="col-lg-12" id="calcResult" style="display: none;">
						<div class="ibox float-e-margins" style="border: 1px solid #dcdada;box-shadow:0 0 25px 0 rgba(118, 222, 118, 0.51);">
							<div class="ibox-title" style="border: 0px;background-color: #a4e2d5;">
								<div class="ibox-tools">
									<h3 style="float: left;margin-top: 0px;" id="curveTitle"><i class="fa fa-tag"></i>&nbsp;试算结果</h3>
									<input type="hidden" name="Curve_id" value="" />
									<input type="hidden" name="Domain_id" value="" />
									<a class="" style="float: right;color: #676a6c;" onclick="closeCalcWindow()">
										<i class="fa fa-times"></i>
									</a>
								</div>
							</div>
							<div class="ibox-content">
								<div class="row">
									<fieldset>
										<legend>基础价格</legend>
										<div class="row">
											<div class="col-xs-4"><span>业务单元：</span><span class="value busiz"></span></div>
											<div class="col-xs-4"><span>FTP计算价格（%）：</span><span class="value ftpPrice"></span></div>
											<div class="col-xs-4"><span>外部利率（%）：</span><span class="value exeRate"></span></div>
										</div>
									</fieldset>
	
									<fieldset>
										<legend>内生性调节项</legend>
										<div class="row inner">
										</div>
									</fieldset>
									<fieldset>
										<legend>最终结果</legend>
										<div class="row">
											<div class="col-xs-4"><span>最终FTP价格（%）：</span><span class="value lastPrice"></span></div>
											<div class="col-xs-4"><span>存续期内FTP利润：</span><span class="value ftpProfit"></span></div>
											<div class="col-xs-4"><span>FTP利润年累计：</span><span class="value ftpYearProfit"></span></div>
										</div>
									</fieldset>
									<fieldset>
										<legend>所匹配政策性调节项</legend>
										<div class="row">
	
										</div>
									</fieldset>
								</div>
							</div>
						</div>
					</div>
				</div>

				<div class="row">
					<div class="col-lg-12">
						<div class="ibox float-e-margins" style="border: 1px solid #dcdada;">
							<!--<div class="ibox-title" style="border: 0px;">
								<div class="ibox-tools">
									<h3 style="float: left;margin-top: 0px;" id="curveTitle"><i class="fa fa-info"></i>&nbsp;单笔试算</h3>
								</div>
							</div>-->
							<div class="ibox-content">
								<form id="calcInfos">
								<div class="row">
									<h3>基础属性字段</h3>
									<hr />
									<div class="base">
										<div class="form-group">
											<label>业务单元:</label>
											<select class="form-control m-b" name="BusinessId" id="BusinessId" onchange="changePrice(this, '3')">
												<option value="">请选择业务单元</option>
											</select>
										</div>
										<div class="form-group">
											<label>币种:</label>
											<select class="form-control m-b" name="IsoCurrencyCd" id="IsoCurrencyCd">
											</select>
										</div>
										<div class="form-group">
											<label>原始金额:</label>
											<input class="form-control" type="text" name="OrgParBal" id="OrgParBal" placeholder="请输入数字类型的原始金额">
										</div>
										<div class="form-group">
											<label>执行利率:</label>
											<input class="form-control" type="text" name="CurNetRate" id="CurNetRate" placeholder="请输入数字类型的执行利率">
										</div>
										<div class="form-group">
											<label>利息计提方式:</label>
											<select class="form-control m-b" name="InterestMode" id="InterestMode">
											</select>
										</div>
										<div class="form-group">
											<label>起息日:</label>
											<div class="input-group date full-width">
												<input type="text" class="form-control input-group-addon " id="OriginalDate" name="OriginalDate">
											</div>
										</div>
										<div class="form-group">
											<label>到期日:</label>
											<div class="input-group date full-width">
												<input type="text" class="form-control input-group-addon" id="MaturityDate" name="MaturityDate">
											</div>
										</div>
										<div class="form-group">
											<label>支付方式:</label>
											<select class="form-control m-b" name="PayInterestMode" id="PayInterestMode" onchange="changePrice(this, '1')">
											</select>
										</div>
										<div class="form-group">
											<label>利率调整方式:</label>
											<select class="form-control m-b" name="AdjustType" id="AdjustType" onchange="changePrice(this, '2')">
											</select>
										</div>
									</div>
									<div class="payment">
										<h3>支付信息字段</h3>
										<hr />
										<div class="paymentInfo">
											<div class="form-group">
												<label>支付频率:</label>
												<input type="text" placeholder="" class="form-control" name="PaymentFreq" id="PaymentFreq">
											</div>
											<div class="form-group">
												<label>支付频率单位:</label>
												<select class="form-control m-b" name="PaymentFreqMult" id="PaymentFreqMult">
																<option value="D">日</option>
																<option value="M" selected="selected">月</option>
																<option value="Y">年</option>
												</select>
											</div>
											<div class="form-group">
												<label>首次支付日:</label>
												<div class="input-group date full-width">
													<input type="text" class="form-control input-group-addon" id="NextPaymentDate" name="NextPaymentDate">
												</div>
											</div>
											<div class="form-group">
												<label>首次支付金额:</label>
												<input type="text" placeholder="" class="form-control" name="OrgPaymentAmt" id="OrgPaymentAmt">
											</div>
										</div>
									</div>
									
									<div class="repeat">
										<h3>重定价信息字段</h3>
										<hr />
										<div class="repeatInfo">
											<div class="form-group">
												<label>重定价频率:</label>
												<input type="text" placeholder="" class="form-control" name="RepriceFreq" id="RepriceFreq">
											</div>
											<div class="form-group">
												<label>重定价频率单位:</label>
												<select class="form-control m-b" name="RepriceFreqMult" id="RepriceFreqMult">
																<option value="D">日</option>
																<option value="M">月</option>
																<option value="Y" selected="selected">年</option>
												</select>
											</div>
											<div class="form-group">
												<label>首次重定价日:</label>
												<div class="input-group date full-width">
													<input type="text" class="form-control input-group-addon" id="NextRepriceDate" name="NextRepriceDate">
												</div>
											</div>
										</div>
									</div>
									<div>
										<button type="button" class="btn btn-primary btn-block" onclick="savePrice()">计算</button>
									</div>
								</div>
								</form>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</body>
	<script type="text/javascript">
		$(function(){
			$('.input-group-addon').datepicker({
				autoclose: true,
				format: 'yyyy-mm-dd'
			});
			
			//动态加载下拉框内容
			$('#BusinessId').select2();
			$('.select2-selection').height(32);

			//动态加载下拉框内容
			$.get('/platform/MasDimCurrency', {r:Math.random()*10000000000000}, function(data) {
				appendOption('IsoCurrencyCd', data, 'IsoCurrencyCd', 'IsoCurrencyName');
			});

			$.get('/platform/BusizInfoCalc', {r:Math.random()*10000000000000}, function(data) {
				var rs = JSON.parse(data);
				if(rs!=null){
					rs.forEach(function(obj) {
						$("#BusinessId").append("<option method='" + obj.Method_id + "' value='" + obj.Busiz_id + "'>" + obj.Busiz_desc + "</option>");
					});
				}
			});
			$.get('/platform/RateAdjustType', {r:Math.random()*10000000000000}, function(data) {
				appendOption('AdjustType', data, 'Adjustable_type_cd', 'Adjustable_type_desc');

//				if($("#AdjustType").val() == '250') {
//					$('#RepriceFreq').val('');
//					$('#NextRepriceDate').val('');
//
//					$('.repeat').slideDown();
//				}
				$("#AdjustType").change();
			});
			$.get('/platform/AccrualCdAttr', {r:Math.random()*10000000000000}, function(data) {
				appendOption('InterestMode', data, 'Accrual_basis_cd', 'Accrual_basis_desc');
				$('#InterestMode').val('3');
			});
			$.get('/platform/PaymentTypeAttr', {r:Math.random()*10000000000000}, function(data) {
				appendOption('PayInterestMode', data, 'Amrt_type_cd', 'Amrt_type_desc');

//				if($("#PayInterestMode").val() != '600') {
//					$('#PaymentFreq').val('');
//					$('#NextPaymentDate').val('');
//					$('#OrgPaymentAmt').val('');
//
//					$('.payment').slideDown();
//				}
				$("#PayInterestMode").change();
			});
		});
		
		function closeCalcWindow(){
			$('#calcResult').slideUp('slow');
		}
		
		function changePrice(obj, type) {
			if(type === '1') { //支付方式改变
				if($(obj).val() != '600') {
					$('.payment').slideDown();
				} else {
					$('.payment').slideUp();
				}
			} else if(type === '2') { //利率调整方式改变
				if($(obj).val() === '250') {
					$('.repeat').slideDown();
				} else {
					$('.repeat').slideUp();
				}
			} else if(type === '3') {
				var method = $(obj).find("option:selected").attr('method');
				$('.base .form-group').show();
				$('.base .form-group input').removeAttr('disabled');
				$('.base .form-group select').removeAttr('disabled');
				if('101' === method || '104' === method || '106' === method) { //指定期限法、沉淀率法、利率差额法
					//不需要支付方式
					$('#PayInterestMode').val('600').change();
					//					$('.payment').slideUp();
					$('#PayInterestMode').parent().hide();
					//利率调整方式为固定利率
					$('#AdjustType').val('0').change();
					$('#AdjustType').attr('disabled', 'disabled');
					//					$('.repeat').slideUp();
					//不需要到期日
					$('#MaturityDate').parent().parent().hide();
				} else if('102' === method) { //直接期限匹配法
					//支付方式默认到期还本付息
					$('#PayInterestMode').val('600').change();
					//					$('.payment').slideUp();
					$('#PayInterestMode').attr('disabled', 'disabled');

					//利率调整方式自选

				} else if('103' === method) { //现金流法

				} else if('105' === method) { //久期法
					//利率调整方式为固定利率
					$('#AdjustType').val('0').change();
					$('#AdjustType').attr('disabled', 'disabled');
				}
			}
		};
		
		function savePrice() {
			var rate = $('#CurNetRate').val(); //执行利率
			var busiz = $('#select2-BusinessId-container').attr('title'); //业务单元名称
			var amt = $('#OrgParBal').val(); //原始金额
			var start = $('#OriginalDate').val(); //起息日
			var end = $('#MaturityDate').val(); //到期日

			var s = new Date(start);
			var e = new Date(end);
			var d = new Date();

			//输入校验
			var method = $('#BusinessId').find("option:selected").attr('method');
			if(method === undefined) {
				toastr.warning('请选择业务单元');
				return;
			}

			var re = /^[0-9]+\.?[0-9]*$/;
			var zzs = /^[0-9]*[1-9][0-9]*$/;
			//原始金额
			if($('#OrgParBal').val() === '' || $('#OrgParBal').val().indexOf('-') >= 0 || !re.test($('#OrgParBal').val())) { //必输、非负
				toastr.warning('请输入正数的原始金额');
				return;
			} else {
				if(parseFloat($('#OrgParBal').val()) > 999999999999999999) { //最大18位正数
					toastr.warning('原始金额输入过大');
					return;
				}
			}

			//执行利率
			if($('#CurNetRate').val() === '' || $('#CurNetRate').val().indexOf('-') >= 0 || !re.test($('#CurNetRate').val())) { //必输、非负
				toastr.warning('请输入正数的执行利率');
				return;
			} else {
				if(parseFloat($('#CurNetRate').val()) > 35) { //最大35
					toastr.warning('执行利率输入过大');
					return;
				}
			}
			
			if(start === "") {
				toastr.warning('请选择起息日');
				return;
			}
			
			if(start > d.Format('yyyy-MM-dd')) {
				toastr.warning('起息日不能大于当前日期');
				return;
			}
			
			if('101' === method || '104' === method || '106' === method) { //指定期限法、沉淀率法、利率差额法
					
			} else if('102' === method || '103' === method || '105' === method) { //直接期限匹配法
				if(start === "" || end === "") {
					toastr.warning('请选择起始和结束日期');
					return;
				}

				if(start > end) {
					toastr.warning('起息日不能大于结束日期');
					return;
				}
				if('102' === method || '103' === method) { //直接期限匹配法 现金流法
					if($('#AdjustType').find("option:selected").val() === '250') { //如果是固定期限浮动
						if($('#RepriceFreq').val() === '' || $('#NextRepriceDate').val() === '') {
							toastr.warning('请完善重定价信息');
							return;
						} else if($('#RepriceFreq').val().indexOf('.') >= 0 || $('#RepriceFreq').val().indexOf('-') >= 0 || !zzs.test($('#RepriceFreq').val())) {
							toastr.warning('请输入正整数的重定价频率');
							return;
						} else {
							var flag = 0;
							if($('#RepriceFreqMult').val() === 'Y') {
								if($('#RepriceFreq').val() > 30) {
									flag = 1;
								}
							} else if($('#RepriceFreqMult').val() === 'M') {
								if($('#RepriceFreq').val() > 30 * 12) {
									flag = 1;
								}
							} else if($('#RepriceFreqMult').val() === 'D') {
								if($('#RepriceFreq').val() > 30 * 365) {
									flag = 1;
								}
							}

							if(flag == 1) {
								toastr.warning('请输入30年内的重定价频率');
								return;
							}
						}
					}
					if('103' === method) { //支付信息
						if($('#PayInterestMode').val() != '600') {
							if($('#PayInterestMode').val() === '' || $('#NextPaymentDate').val() === '' || $('#OrgPaymentAmt').val() === '') {
								toastr.warning('请完善支付信息');
								return;
							} else if($('#PaymentFreq').val().indexOf('.') >= 0 || $('#PaymentFreq').val().indexOf('-') >= 0 || !zzs.test($('#PaymentFreq').val())) {
								toastr.warning('请输入正整数的支付频率');
								return;
							} else {
								var flag = 0;
								if($('#PaymentFreqMult').val() === 'Y') {
									if($('#PaymentFreq').val() > 30) {
										flag = 1;
									}
								} else if($('#PaymentFreqMult').val() === 'M') {
									if($('#PaymentFreq').val() > 30 * 12) {
										flag = 1;
									}
								} else if($('#PaymentFreqMult').val() === 'D') {
									if($('#PaymentFreq').val() > 30 * 365) {
										flag = 1;
									}
								}

								if(flag == 1) {
									toastr.warning('请输入30年内的支付频率');
									return;
								}
							}

							if($('#OrgPaymentAmt').val() === '' || $('#OrgPaymentAmt').val().indexOf('-') >= 0 || !re.test($('#OrgPaymentAmt').val())) { //必输、非负
								toastr.warning('请输入正数的首次支付金额');
								return;
							} else {
								if(parseFloat($('#OrgPaymentAmt').val()) > 999999999999999999) { //最大18位正数
									toastr.warning('首次支付金额输入过大');
									return;
								}
							}
						}
					}
				} else if('105' === method) { //久期法
					if($('#PayInterestMode').val() != '600') {
						if($('#PayInterestMode').val() === '' || $('#NextPaymentDate').val() === '' || $('#OrgPaymentAmt').val() === '') {
							toastr.warning('请完善支付信息');
							return;
						} else if($('#PaymentFreq').val().indexOf('.') >= 0 || $('#PaymentFreq').val().indexOf('-') >= 0 || !zzs.test($('#PaymentFreq').val())) {
							toastr.warning('请输入正整数的支付频率');
							return;
						} else {
							var flag = 0;
							if($('#PaymentFreqMult').val() === 'Y') {
								if($('#PaymentFreq').val() > 30) {
									flag = 1;
								}
							} else if($('#PaymentFreqMult').val() === 'M') {
								if($('#PaymentFreq').val() > 30 * 12) {
									flag = 1;
								}
							} else if($('#PaymentFreqMult').val() === 'D') {
								if($('#PaymentFreq').val() > 30 * 365) {
									flag = 1;
								}
							}

							if(flag == 1) {
								toastr.warning('请输入30年内的支付频率');
								return;
							}
						}

						if($('#OrgPaymentAmt').val() === '' || $('#OrgPaymentAmt').val().indexOf('-') >= 0) { //必输、非负
							toastr.warning('请输入正数的首次支付金额');
							return;
						} else {
							if(parseFloat($('#OrgPaymentAmt').val()) > 999999999999999999) { //最大18位正数
								toastr.warning('首次支付金额输入过大');
								return;
							}
						}
					}
				}
			}

			var busiDay = 0; //本年业务存在天数
//			if(s.Format('yyyy') < d.Format('yyyy')) {
//				busiDay = 365;
//			} else {
//				busiDay = (new Date(d.Format('yyyy') + "-12-31") - s) / 86400000;
//			}
			if(end != ''){ //到期日不为空
				if(e.Format('yyyy') == d.Format('yyyy')){ //到期日在本年
					if(s.Format('yyyy') == d.Format('yyyy')){ //起息日在本年
						busiDay=(e-s)/86400000;
					}else if(s.Format('yyyy') < d.Format('yyyy')){ //起息日在本年之前
						busiDay=(e-new Date(d.Format('yyyy') + "-01-01"))/ 86400000;
					}
				}else if(e.Format('yyyy') > d.Format('yyyy')){ //到期日本年之后
					if(s.Format('yyyy') == d.Format('yyyy')){ //起息日本年
						busiDay=(new Date(d.Format('yyyy') + "-12-31") - s) / 86400000;
					}else if(s.Format('yyyy') < d.Format('yyyy')){ //起息日本年之前
						var InterestMode=$('#InterestMode').val();
						if(InterestMode=='3' || InterestMode == '5'){
							year=d.Format('yyyy');
							if((year % 4 == 0) && (year % 100 != 0 || year % 400 == 0)){
								busiDay=366;
							}else {
								busiDay=365;
							}
						}else if(InterestMode=='4' || InterestMode == '6'){
							busiDay=365;
						}else {
							busiDay=360;
						}
					}
				}
			}else { //到期日为空
				if(s.Format('yyyy') < d.Format('yyyy')) { //起息日本年之前
					var InterestMode=$('#InterestMode').val();
					if(InterestMode=='3' || InterestMode == '5'){
						year=d.Format('yyyy');
						if((year % 4 == 0) && (year % 100 != 0 || year % 400 == 0)){
							busiDay=366;
						}else {
							busiDay=365;
						}
					}else if(InterestMode=='4' || InterestMode == '6'){
						busiDay=365;
					}else {
						busiDay=360;
					}
				} else { //起息日本年内
					busiDay = (new Date(d.Format('yyyy') + "-12-31") - s) / 86400000;
				}
			}
			
			var busizId = $('#BusinessId').find("option:selected").attr('value');
			
			$.blockUI({
				message: '<h1>计算中。。。</h1>',
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
			
			$.post('/platform/FtpValueCalc?' + $('#calcInfos').serialize() + '&MethodID=' + method + '&r=' + Math.random(100000000000000), function(data) {
				var rs = JSON.parse(data);
				$.unblockUI();
				//组装结果
				$('.ftpPrice').text(rs.FtpValue); //ftp价格
				$('.busiz').text(busiz); //业务单元名
				$('.exeRate').text(rate); //执行利率

				if(rs.ErrorCode === '1') {
					var lPrice = parseFloat(rs.FtpValue);
					//<div class="cell"><span>外部利率：</span><span class="value exeRate"></span></div>
					var temp = "<div class='col-xs-4'><span>";
//					$('.inner').html(''); //处理内生性调节项
					if(rs.Insideinfo != null) {
						for(var i = 0; i < rs.Insideinfo.length; i++) {
							lPrice += parseFloat(rs.Insideinfo[i].Adjustvale);
							temp += rs.Insideinfo[i].Adjustname + ":" + "</span><span class='value'>" + rs.Insideinfo[i].Adjustvale + "</span></div>";
							$('.inner').append(temp);
							temp = "<div class='col-xs-4'><span>";
						}
					}

					var yProfit=0;
					if(busizId.substring(0, 1) === '1') { //资产业务
						yProfit = parseFloat(amt) * (parseFloat(rate) - lPrice) * busiDay / 365 / 100;
					} else if(busizId.substring(0, 1) === '2') { //负载业务
						yProfit = parseFloat(amt) * (lPrice - parseFloat(rate)) * busiDay / 365 / 100;
					}

					$('.lastPrice').text(lPrice.toFixed(6)); //最终价格
					$('.ftpYearProfit').text(yProfit.toFixed(2)); //FTP利润年累计

					var method = $('#BusinessId').find("option:selected").attr('method');
					if('101' === method || '104' === method || '106' === method) {
						$('.ftpProfit').parent().hide();
					} else {
						var profit;
						if(busizId.substring(0, 1) === '1') { //资产业务
							profit = (parseFloat(rate) - lPrice) / 100 * parseFloat(amt) * ((e - s) / 86400000 / 365);
						} else if(busizId.substring(0, 1) === '2') { //负载业务
							profit = (lPrice - parseFloat(rate)) / 100 * parseFloat(amt) * ((e - s) / 86400000 / 365);
						}

						$('.ftpProfit').text(profit.toFixed(2)); //存续期内FTP利润
						$('.ftpProfit').parent().show();
					}

					$('#calcResult').slideDown();
				} else {
					toastr.error(rs.ErrorMsg);
					return false;
				}
			});
		}
	</script>

</html>