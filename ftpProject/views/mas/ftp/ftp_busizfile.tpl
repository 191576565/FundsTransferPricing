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

		<title>产品定义</title>

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
			#wrapper {
				width: 100%;
				height: 100%;
				overflow-y: auto;
			}
			
			#wrapper,
			.wrapper,
			.row {
				height: 100%;
			}
			
			h2{
				margin: 5px 0px;
			}
			
			.content .first,.content .second, .p{
				height: 98%;
			}
			.content .first .col-lg-12{
				height: 50%;
			}
			
			.content .second .col-lg-12{
				height: 40%;
			}
			.panel-group{
				height: 100%;
				overflow: auto;
				margin: 0px 0px;
			}
		</style>
	</head>

	<body>
		<div id="wrapper">
			<div class="wrapper wrapper-content animated fadeInRight">
				<div class="row content">
					<div class="col-lg-6 first">
						<div class="col-lg-12">
							<div class="panel panel-default p">
								<div class="panel-heading">
									<h2>业务方案类</h2>
								</div>
								<div class="panel-body" style="overflow: auto;height: 90%;">
									<div class="panel-group" id="busiz">
										<div class="panel panel-default">
											<div class="panel-heading">
												<h5 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#busiz" href="#1">需求分析下载</a>
                                                <a class="btn btn-link pull-right">上传</a>
                                            </h5>
											</div>
											<div id="1" class="panel-collapse collapse in">
												<div class="panel-body">
													<ol>
							                            <li><span>But I must explain </span></li>
							                            <li><span>To you how all this mistaken</span></li>
							                            <li><span>Idea of denouncing pleasure </span></li>
							                            <li><span>Great explorer of the truth</span></li>
							                            <li><span>To take a trivial example</span></li>
							                            <li><span>That they cannot foresee</span></li>
							                            <li><span>Who avoids a pain that produceg</span></li>
							                            <li><span>Consequences that are extremely </span></li>
							                        </ol>
												</div>
											</div>
										</div>
										<div class="panel panel-default">
											<div class="panel-heading">
												<h4 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#busiz" href="#2">方案设计下载</a>
                                            </h4>
											</div>
											<div id="2" class="panel-collapse collapse">
												<div class="panel-body">
												</div>
											</div>
										</div>
										<div class="panel panel-default">
											<div class="panel-heading">
												<h4 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#busiz" href="#3">模型设计下载</a>
                                            </h4>
											</div>
											<div id="3" class="panel-collapse collapse">
												<div class="panel-body">
												</div>
											</div>
										</div>
										<div class="panel panel-default">
											<div class="panel-heading">
												<h4 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#busiz" href="#4">产品配置下载</a>
                                            </h4>
											</div>
											<div id="4" class="panel-collapse collapse">
												<div class="panel-body">
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
						<div class="col-lg-12">
							<div class="panel panel-default p">
								<div class="panel-heading">
									<h2>技术方案类</h2>
								</div>
								<div class="panel-body" style="overflow: auto;height: 90%;">
									<div class="panel-group" id="tech">
										<div class="panel panel-default">
											<div class="panel-heading">
												<h5 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#tech" href="#5">设计方案下载</a>
                                            </h5>
											</div>
											<div id="5" class="panel-collapse collapse in">
												<div class="panel-body">
												</div>
											</div>
										</div>
										<div class="panel panel-default">
											<div class="panel-heading">
												<h4 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#tech" href="#6">数据处理方案下载</a>
                                            </h4>
											</div>
											<div id="6" class="panel-collapse collapse">
												<div class="panel-body">
												</div>
											</div>
										</div>
										<div class="panel panel-default">
											<div class="panel-heading">
												<h4 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#tech" href="#7">维护方案下载</a>
                                            </h4>
											</div>
											<div id="7" class="panel-collapse collapse">
												<div class="panel-body">
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
					<div class="col-lg-6 second">
						<div class="col-lg-12">
							<div class="panel panel-default p">
								<div class="panel-heading">
									<h2>分析方案类</h2>
								</div>
								<div class="panel-body" style="overflow: auto;height: 90%;">
									<div class="panel-group" id="analyze">
										<div class="panel panel-default">
											<div class="panel-heading">
												<h5 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#analyze" href="#8">结果分析模板下载</a>
                                            </h5>
											</div>
											<div id="8" class="panel-collapse collapse in">
												<div class="panel-body">
												</div>
											</div>
										</div>
										<div class="panel panel-default">
											<div class="panel-heading">
												<h4 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#analyze" href="#9">结果分析方案下载</a>
                                            </h4>
											</div>
											<div id="9" class="panel-collapse collapse">
												<div class="panel-body">
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
						<div class="col-lg-12">
							<div class="panel panel-default p">
								<div class="panel-heading">
									<h2>规则制度与办法</h2>
								</div>
								<div class="panel-body" style="overflow: auto;height: 90%;">
									<div class="panel-group" id="regulation">
										<div class="panel panel-default">
											<div class="panel-heading">
												<h5 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#regulation" href="#10">FTP相关规则制度下载</a>
                                            </h5>
											</div>
											<div id="10" class="panel-collapse collapse in">
												<div class="panel-body">
												</div>
											</div>
										</div>
										<div class="panel panel-default">
											<div class="panel-heading">
												<h4 class="panel-title">
                                                <a data-toggle="collapse" data-parent="#regulation" href="#11">FTP相关使用办法下载</a>
                                            </h4>
											</div>
											<div id="11" class="panel-collapse collapse">
												<div class="panel-body">
												</div>
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
						<div class="col-md-12">
							<div class="panel panel-default p">
								<div class="panel-heading">
									<h2>业务方案数据导出</h2>
								</div>
								<div class="panel-body" style="overflow: auto;height: 90%;">
									<button class="btn btn-block btn-primary dim" onclick=""><i class="fa fa-download"></i>&nbsp;导出</button>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</body>
	<script type="text/javascript">
		$(function() {
			$.get('/platform/BusizFile', function(data){
			})
		});
	</script>

</html>