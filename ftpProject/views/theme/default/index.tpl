<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">

	<head>
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
		<meta name="description" content="Metro, a sleek, intuitive, and powerful framework for faster and easier web development for Windows Metro Style.">
		<meta name="keywords" content="HTML, CSS, JS, JavaScript, framework, metro, front-end, frontend, web development">
		<meta name="author" content="Sergey Pimenov and Metro UI CSS contributors">
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta http-equiv="Expires" content="0">
		<meta http-equiv="Pragma" content="no-cache">
		<meta http-equiv="Cache-control" content="no-cache">
		<meta http-equiv="Cache" content="no-cache">

		<title>天健科技大数据应用分析平台</title>

		<link rel="stylesheet" href="/static/css/bootstrap.min.css" />
		<link rel="stylesheet" href="/static/css/font-awesome.min.css" />

		<link rel="stylesheet" href="/static/css/metro.css">
		<link rel="stylesheet" href="/static/css/metro-icons.css">
		<link rel="stylesheet" href="/static/css/jquery.mCustomScrollbar.min.css" type="text/css" />
		<link rel="stylesheet" href="/static/css/bootstrap-table.min.css" type="text/css" />
		<link rel="stylesheet" href="/static/theme/default/index.css" type="text/css" />
		<link rel="stylesheet" href="/static/theme/default/sysconfig.css" type="text/css" />
		<link href="/static/inspinia/css/plugins/toastr/toastr.min.css" rel="stylesheet">
		<link href="/static/inspinia/css/plugins/iCheck/custom.css" rel="stylesheet">

		<script type="text/javascript" src="/static/js/jquery-1.12.3.min.js"></script>
		<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>

		<script type="text/javascript" src="/static/js/metro.js"></script>

		<script type="text/javascript" src="/static/js/bootstrap-table.min.js"></script>
		<script type="text/javascript" src="/static/js/bootstrap-table-zh-CN.min.js"></script>
		<script type="text/javascript" src="/static/js/bootstrap-table-export.min.js"></script>
		<script type="text/javascript" src="/static/js/bootstrap-table-tableExport.js"></script>
		<script type="text/javascript" src="/static/js/echarts-all.js"></script>

		<script type="text/javascript" src="/static/js/hzwy23table-from-bootstrap-table.js"></script>
		<script type="text/javascript" src="/static/js/modal.js"></script>
		<script type="text/javascript" src="/static/js/hzwy23.js"></script>
		<script type="text/javascript" src="/static/js/jquery.mCustomScrollbar.concat.min.js"></script>

		<link href="/static/inspinia/css/animate.css" rel="stylesheet">
		<link href="/static/inspinia/css/style.css" rel="stylesheet">
		<!-- Toastr style -->
		<!-- Sweet Alert -->
		<link href="/static/inspinia/css/plugins/sweetalert/sweetalert.css" rel="stylesheet">

		<!-- Custom and plugin javascript -->
		<!--<script src="/static/inspinia/js/plugins/metisMenu/jquery.metisMenu.js"></script>
		<script src="/static/inspinia/js/inspinia.js"></script>-->
		<script src="/static/inspinia/js/plugins/pace/pace.min.js"></script>
		<script src="/static/inspinia/js/plugins/iCheck/icheck.min.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/plugins/toastr/toastr.min.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/plugins/sweetalert/sweetalert.min.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/jquery.form.js" type="text/javascript" charset="utf-8"></script>
		<script src="/static/inspinia/js/bootstrap-filestyle.min.js"></script>
		<script src="/static/inspinia/jquery.blockUI.js" type="text/javascript" charset="utf-8"></script>
		<!--<script src="/static/inspinia/ftpCommon.js" type="text/javascript" charset="utf-8"></script>-->

		<!--<script type="text/javascript" src="/static/js/angular.min.js"></script>-->
		<style type="text/css">
			#wrap .col-sm-12 {
				margin: 20px 10px;
			}
			
			.modal-header .close {
				padding-right: 8px !important;			
			}
		   .tile-group .tile-group-title {
			    color: #ffffff;
			    font-size: 25px;
			    line-height: 20px;
			    position: absolute;
			    top: 10px;
			    left: 0;
				font-weight:bold;
			}
			.tile .tile-label {
			    position: absolute;
			    bottom: 0px;
			    left: 0.625rem;
			    z-index: 999;
			    padding: 0.425rem 0.25rem;
				font-size: 15px;
				font-weight:bold;
			}	
		</style>
	</head>

	<body>
		<div style="background-size: cover;" class="theme-bg-color">
			<div id="h-left-tool-bar" class="h-left-tool-bar" style="z-index: 8888;">
				<div class="h-left-btn-desk" type data-toggle="tooltip" data-placement="right" title="菜单列表"><i class="icon-columns"></i></div>
				<div class="h-left-btn-off" data-toggle="tooltip" data-placement="right" title="安全退出" onclick="LogOut()"><i class="icon-off"></i></div>
			</div>
			<div id="bigdata-platform-subsystem" class="container-fluid" style="padding-left:55px; overflow: hidden;">
				<!--<div class="row" style="position: relative; height: 60px;">
					
				</div>-->
				<div id="wrap" class="row" style="height: 10px; padding-right: 1px; ">
					<div class="col-sm-6 col-md-6 col-lg-6" style="font-size: 30px;width: 100%;padding-left: 35px;text-align: left;padding-top: 5px;padding-bottom: 0px;">
						<!--<h4 id="huuid" style="color: #ffffff; font-size: 30px; font-weight: 700; height: 40px; line-height: 40px;margin-bottom: 20px;">天健金管资金转移定价系统</h4>-->
						<img src="/static/theme/default/img/logo.png" alt="" width="500px" height="80px" />
					</div>
					<!--<div id="h-system-service" class="col-sm-12">
			</div>
			<div id="h-mas-service" class="col-sm-12">
			</div>
			<div id="h-other-service"  class="col-sm-12">
			</div>-->
					<div id="h-service" class="col-sm-12" style="margin: 0px;padding: 0px;">
					</div>
				</div>
				<div style="background-color: #086d87;height: 30px;">
					<div class="row full-height" id="copyRight">
						<div class="col-md-12" style="height: 30px;text-align: center;color: white;line-height: 30px;">
							<!--<div class="col-md-3" style="text-align: center;color: white;">Copyright © 2016-2017</div>
							<div class="col-md-6" style="text-align: center;color: white;">重庆天健金管科技服务有限公司</div>
							<div class="col-md-3" style="text-align: center;color: white;">All rights reserved</div>-->
							Copyright © 2016-2017&nbsp;重庆天健金管科技服务有限公司&nbsp;All rights reserved
						</div>
					</div>
				</div>
			</div>
		</div>

		<div class="modal inmodal info" id="changePassDialog" tabindex="-1" role="dialog" aria-hidden="true" data-backdrop="true">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<h2 id="addUserInfo">修改密码</h2>
						<input type="hidden" id="type" />
					</div>
					<div class="modal-body">
						<form id="plat-change-passwd">
							<div class="row">
								<div class="col-xs-12">
									<div class="form-group">
										<label>原密码:</label>
										<input placeholder="" class="form-control" type="password" name="ora_passwd">
									</div>
								</div>
								<div class="col-xs-12">
									<div class="form-group">
										<label>新密码:</label>
										<input placeholder="" class="form-control" type="password" name="new_passwd">
									</div>
								</div>
								<div class="col-xs-12">
									<div class="form-group">
										<label>确认密码:</label>
										<input placeholder="" class="form-control" type="password" name="sure_passwd">
									</div>
								</div>
							</div>
						</form>
					</div>

					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="changPass()">保存</button>
					</div>
				</div>
			</div>
		</div>

		<div class="modal inmodal info" id="exportDialog" tabindex="-1" role="dialog" aria-hidden="true" data-backdrop="true">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 id="addUserInfo">批量数据导出</h2>
						<input type="hidden" id="type" />
					</div>
					<div class="modal-body">
						<!--<form id="plat-change-passwd">
							<div class="row">
								<div class="col-md-4"><label class="checkbox-inline i-checks" down="DownCurve" style="width: 30px;"> <input type="checkbox"></label>曲线定义</div>
								<div class="col-md-4"><label class="checkbox-inline i-checks" down="DownBusiz" style="width: 30px;"> <input type="checkbox"></label>规则配置</div>
								<div class="col-md-4"><label class="checkbox-inline i-checks" down="DownAdj" style="width: 30px;"> <input type="checkbox"></label>调节项信息</div>
							</div>
						</form>-->
						<button type="button" class="btn btn-primary btn-block" onclick="exportData()"><i class="fa fa-download"></i>导出</button>
					</div>

					<!--<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
						<button type="button" class="btn btn-primary" onclick="exportData()">导出</button>
					</div>-->
				</div>
			</div>
		</div>

		<div class="modal inmodal" id="uploadDialog" tabindex="-2" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2>批量导入业务方案数据</h2>
					</div>
					<div class="modal-body">
						<div class="row">
							<form id="uploadFormAll" enctype="multipart/form-data" method="post">
								<div class="col-xs-12">
									<h3>1、全量导入</h3>
								</div>
								<div class="col-xs-8">
									<div class="form-group">
										<input type="file" id="uploadFileAll" name="uploadFile">
									</div>
								</div>
								<div class="col-xs-4" style="padding-left: 0px;">
									<button class="btn btn-info" type="">提交</button>
								</div>
							</form>
							<form id="uploadFormCurve" enctype="multipart/form-data" method="post">
								<div class="col-xs-12">
									<h3>2、曲线点值增量导入</h3>
								</div>
								<div class="col-xs-8">
									<div class="form-group">
										<input type="file" id="uploadFileCurve" name="uploadFile">
									</div>
								</div>
								<div class="col-xs-4" style="padding-left: 0px;">
									<button class="btn btn-info" type="submit">提交</button>
									<a class="btn btn-link" style="text-decoration: underline;" href="/updownload/CurveInfo.xlsx">下载模板</a>
								</div>
							</form>
							<form id="uploadFormPolicy" enctype="multipart/form-data" method="post">
								<div class="col-xs-12">
									<h3>3、政策性调节项增量导入</h3>
								</div>
								<div class="col-xs-8">
									<div class="form-group">
										<input type="file" id="uploadFilePolicy" name="uploadFile">
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
	</body>
	<script type="text/javascript">
		$(function() {
			$.ajaxSetup({
				scriptCharset: 'utf-8',
				complete: function(xhr, status) {
					//					var x=JSON.stringify(xhr);
					var sessionStatus = xhr.status;
					if(sessionStatus == '403') {
						var top = getTopWinow();
						swal("提示!", "由于您长时间没有操作, session已过期, 请重新登录.", "error");
						setTimeout(function() {
							top.location.href = '/';
						}, 2000);
					}
				}
			});

			iCheckInit();

			$('#uploadFileAll').filestyle({
				buttonName: 'btn-warning',
				buttonText: '选择文件',
				icon: false
			});
			$('#uploadFileCurve').filestyle({
				buttonName: 'btn-warning',
				buttonText: '选择文件',
				icon: false
			});
			$('#uploadFilePolicy').filestyle({
				buttonName: 'btn-warning',
				buttonText: '选择文件',
				icon: false
			});

			//全量导入
			var options = {
				contentType: "application/x-www-form-urlencoded;charset=utf-8",
				url: '/mas/ftp/FtpFormBackup?r=' + Math.random() * 10000000000000,
				success: function(data) {
					var rs = JSON.parse(data);
					$.unblockUI();
					if(rs.ErrorCode === '1') {
						toastr.success(rs.ErrorMsg);
					} else {
						toastr.error(rs.ErrorMsg);
						return;
					}
				},
				beforeSubmit: function(arr, $form, options) {
					var file = $('#uploadFileAll').val(),
						fileType = file.split('.');
					if(file === '' || fileType[fileType.length - 1] != 'xlsx') {
						toastr.warning("请选择要上传的Excel文件");
						return false;
					}

					$.get('/platform/MenuPage', function() {
						$.blockUI({
							message: '<h1>上传导入中。。。</h1>',
							baseZ: 9999,
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
			$('#uploadFormAll').ajaxForm(options);

			//曲线点值增量导入
			var options = {
				contentType: "application/x-www-form-urlencoded;charset=utf-8",
				url: '/mas/ftp/CurveInfoInput?r=' + Math.random() * 10000000000000,
				success: function(data) {
					var rs = JSON.parse(data);
					$.unblockUI();
					if(rs.ErrorCode === '1') {
						toastr.success(rs.ErrorMsg);
					} else {
						toastr.error(rs.ErrorMsg);
						return;
					}
				},
				beforeSubmit: function(arr, $form, options) {
					var file = $('#uploadFileCurve').val(),
						fileType = file.split('.');
					if(file === '' || fileType[fileType.length - 1] != 'xlsx') {
						toastr.warning("请选择要上传的Excel文件");
						return false;
					}
					$.get('/platform/MenuPage', function() {
						$.blockUI({
							message: '<h1>上传导入中。。。</h1>',
							baseZ: 9999,
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
			$('#uploadFormCurve').ajaxForm(options);

			//政策性调节项增量导入
			var options = {
				contentType: "application/x-www-form-urlencoded;charset=utf-8",
				url: '/mas/ftp/FtpAdjUpload?r=' + Math.random() * 10000000000000,
				success: function(data) {
					var rs = JSON.parse(data);
					$.unblockUI();
					if(rs.ErrorCode === '1') {
						toastr.success(rs.ErrorMsg);
					} else {
						toastr.error(rs.ErrorMsg);
						return;
					}
				},
				beforeSubmit: function(arr, $form, options) {
					var file = $('#uploadFilePolicy').val(),
						fileType = file.split('.');
					if(file === '' || fileType[fileType.length - 1] != 'xlsx') {
						toastr.warning("请选择要上传的Excel文件");
						return false;
					}
					$.get('/platform/MenuPage', function() {
						$.blockUI({
							message: '<h1>上传导入中。。。</h1>',
							baseZ: 9999,
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
			$('#uploadFormPolicy').ajaxForm(options);
		});

		function iCheckInit() {
			$('.i-checks').iCheck({
				checkboxClass: 'icheckbox_square-green',
				radioClass: 'iradio_square-green',
			});
		}

		function getTopWinow() {
			var p = window;
			while(p != p.parent) {
				p = p.parent;
			}
			return p;
		}

		function showChangePassDialog() {
			$("#plat-change-passwd")[0].reset();
			$('#changePassDialog').modal('show');
		}

		function changPass() {
			$.ajax({
				type: "post",
				url: "/platform/passwd",
				data: $("#plat-change-passwd").serialize(),
				cache: !1,
				async: !1,
				dataType: "text",
				error: function(a, b, e) {
					toastr.error('密码修改出现异常');
				},
				success: function(data) {
					var rs = JSON.parse(data);
					if(rs.ErrorCode === '1') {
						toastr.success('密码修改成功！');
						window.location.href = "/"
					} else {
						toastr.error(rs.ErrorMsg);
						return;
					}
				}
			})
		}

		function exportData() {
			var d = { "DownCurve": 1, "DownBusiz": 1, "DownAdj": 1 };
			//			if($('.checked').length===0){
			//				toastr.warning('请至少选择一个要导出数据的模块');
			//				return;
			//			}
			//			
			//			$('.checked').each(function(){
			//				d[$(this).parent().attr('down')]=1;
			//			});

			var form = $("<form>"); //定义一个form表单
			form.attr("style", "display:none");
			form.attr("target", "");
			form.attr("method", "get");
			form.attr("action", '/mas/ftp/FtpFormBackup?r=' + Math.random() * 10000000000000);
			var input1 = $("<input>");
			input1.attr("type", "hidden");
			input1.attr("name", "DownCurve");
			input1.attr("value", d['DownCurve']);
			var input2 = $("<input>");
			input2.attr("type", "hidden");
			input2.attr("name", "DownBusiz");
			input2.attr("value", d['DownBusiz']);
			var input3 = $("<input>");
			input3.attr("type", "hidden");
			input3.attr("name", "DownAdj");
			input3.attr("value", d['DownAdj']);
			$("body").append(form); //将表单放置在web中
			form.append(input1);
			form.append(input2);
			form.append(input3);

			form.submit(); //表单提交 
		}

		function LogOut() {
			swal({
				title: "登出！",
				text: "是否退出登录？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				closeOnConfirm: false
			}, function() {
				$.ajax({
					type: "Get",
					url: "/logout",
					cache: !1,
					async: !1,
					dataType: "text",
					error: function() {
						window.location.href = "/"
					},
					success: function(a) {
						window.location.href = a
					}
				})
			});
		}

		/*
		 * 自动调整icon在框中的位置,使其水平与竖直方向居中显示
		 * */

		function go_entry(e) {

			var id = $(e).attr("data-id");

			var quit = function() {
				window.location.href = "/"
			};

			var err = function(dt) {
				alert(dt)
			};

			var succ = function(d) {
				//				$("#indexHtmlContent").html($("body").html())
				$("#bigdata-platform-subsystem").html(d)
			};

			$.get('/platform/select', {
				Id: id,
				r: Math.random() * 1000000000
			}, succ);
		}

		function go_entryfirst() {

			var id = 208000000000

			var quit = function() {
				window.location.href = "/"
			};

			var err = function(dt) {
				alert(dt)
			};

			var succ = function(d) {
				//				$("#indexHtmlContent").html($("body").html())
				$("#bigdata-platform-subsystem").html(d)
			};

			$.get('/platform/select', {
				Id: id,
				r: Math.random() * 1000000000
			}, succ);
		}

		$(document).ready(function() {
			go_entryfirst();
			var succ = function(data) {
				var rs = JSON.parse(data);
				//			var rs=data;
				if(rs != null) {
					$('#wrap #h-service').html('');
					rs.forEach(function(e) {
						if(e.Res_up_id === '-1') {
							var temp = "<div class='' style='margin-left: 90px;margin-bottom:20px'><span class='tile-group-title' style='font-size:20px;color:white;'>" + e.Res_name + "</span><div class='tile-container'  style='padding-top:10px;' resid='" + e.Res_id + "'></div></div>"
							$('#wrap #h-service').append(temp);
						} else {
							var temp = "<div data-id='" + e.Res_id + "' onclick='go_entry(this)'  data-url='" + e.Res_url + "' class='" + e.Res_class + " fg-white' data-role='tile' data-role='tile' style='background-color:" + e.Res_bg_color + "'>" +
								"<div class='tile-content iconic'><span class='icon'><img src='" + e.Res_img + "'></span></div><div class='tile-label'>" + e.Res_name + "</div>";
							$("div[resid='" + e.Res_up_id + "']").append(temp);
						}
					});
				}

				$(function() {
					//取消水平滑动的插件
					//$.StartScreen();

					var tiles = $(".tile, .tile-small, .tile-sqaure, .tile-wide, .tile-large, .tile-big, .tile-super");

					$.each(tiles, function() {
						var tile = $(this);
						setTimeout(function() {
							tile.css({
								opacity: 1,
								"-webkit-transform": "scale(1)",
								"transform": "scale(1)",
								"-webkit-transition": ".3s",
								"transition": ".3s"
							});
						}, Math.floor(Math.random() * 500));
					});

					$(".tile-group").animate({
						left: 0
					});
				});
			}

			//		HAjaxRequest({
			//			url:'/platform/MenuPage',
			//			data:{TypeId:0,Id:'-1'},
			//			success:succ,
			//		})
			
		});

		//调整主菜单的长度和宽度
		$(document).ready(function() {
			var hh = document.documentElement.clientHeight;
			$(".container-fluid").height(hh);
			$("#h-left-tool-bar").height(hh);
			$("#wrap").height(hh - 30);
			$(".tile").click(function() {
				go_entry(this)
			})
			$("#wrap").mCustomScrollbar({
				axis: "y",
				theme: "dark-thin",
				scrollSpeed: 100
			});

			$('.mCSB_scrollTools').width(0);
		});

		window.onload = function() {
			if(typeof history.pushState === "function") {
				history.pushState("jibberish", null, null);
				window.onpopstate = function() {
					history.pushState('newjibberish', null, null);
					// Handle the back (or forward) buttons here
					// Will NOT handle refresh, use onbeforeunload for this.
					H_LEFT_BAR.H_HomePage()
				};
			} else {
				var ignoreHashChange = true;
				window.onhashchange = function() {
					if(!ignoreHashChange) {
						ignoreHashChange = true;
						window.location.hash = Math.random();

						// Detect and redirect change here
						// Works in older FF and IE9
						// * it does mess with your hash symbol (anchor?) pound sign
						// delimiter on the end of the URL
					} else {
						ignoreHashChange = false;
					}
				};
			}
		}
		$("[data-toggle='tooltip']").tooltip();

		function getEvent() {
			if(window.event) { return window.event; }
			func = getEvent.caller;
			while(func != null) {
				var arg0 = func.arguments[0];
				if(arg0) {
					if((arg0.constructor == Event || arg0.constructor == MouseEvent ||
							arg0.constructor == KeyboardEvent) ||
						(typeof(arg0) == "object" && arg0.preventDefault &&
							arg0.stopPropagation)) {
						return arg0;
					}
				}
				func = func.caller;
			}
			return null;
		}
		//阻止冒泡
		function cancelBubble() {
			var e = getEvent();
			if(window.event) {
				//e.returnValue=false;//阻止自身行为
				e.cancelBubble = true; //阻止冒泡
			} else if(e.preventDefault) {
				//e.preventDefault();//阻止自身行为
				e.stopPropagation(); //阻止冒泡
			}
		}
		window.onbeforeunload = function() {
			//			$.get('/platform/MenuPage');
			$.ajax({ url: '/platform/MenuPage', async: false });
		}
		
	</script>

</html>