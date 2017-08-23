<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">

	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="renderer" content="ie-stand">
		<title>天健科技大数据应用分析平台</title>
		<link rel="stylesheet" href="/static/css/bootstrap.min.css" />

		<link rel="stylesheet" href="/static/css/font-awesome.min.css" />

		<link rel="stylesheet" href="/static/css/common.css" />
		<link href="/static/inspinia/css/plugins/toastr/toastr.min.css" rel="stylesheet">

		<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
		<script src="/static/js/jquery-1.12.3.min.js"></script>

		<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
		<script src="/static/js/bootstrap.min.js"></script>
		<script src="/static/inspinia/js/plugins/toastr/toastr.min.js" type="text/javascript" charset="utf-8"></script>

		<style>
			.login_bg {
				background-image: url(/static/img/index_bg.jpg);
				background-repeat: no-repeat;
				width: 100%;
				padding: 0px;
				margin: 0px;
				background-size: cover;
			}
			
			.login_input_area {
				background-image: url(/static/img/login_info_bg.png);
				z-index: 10;
				height: 441px;
				width: 431px;
				position: fixed;
			}
			
			.inputZoon {
				margin-top: 200px;
				width: 380px;
				text-align: center;
				margin-left: 20px;
			}
			
			.form-group {
				margin: 15px 60px;
			}
		</style>
	</head>

	<body id="show">

		<div id='login_page' class="login_bg row col-sm-12 col-md-12 col-lg-12">

			<!--<div id="logo_area" class="logo_area">
             </div>
				<h4 class="col-sm-12 col-mg-12 col-lg-12" style="margin-top:15px; font-weight:bold; color:#FFF">天健金管资金转移定价系统</h4>
			</div>-->
			<div id='input_area' class='login_input_area'>
				<form id='inputZoon' class='inputZoon' action="" method="POST">
					<!--<div class="input-group" style="margin:0 60px;">
        <span class="input-group-addon" style="width:60px; color:#36F; background-color:inherit;"><i class="icon-user"></i></span>
        <input id="p_user_id" class="form-control" type="text" name="username" />
      </div>
      <br />
               
      <div  class="input-group" style="margin:0 60px;">
        <span class="input-group-addon" style="width:60px; color:#36F; background-color:inherit;"><i class="icon-lock"></i></span>
        <input id="p_passwd_id" class="form-control" type="password" name="password" />
      </div>
      <br />
               
      <div style="margin:0 60px;">
        <button type="submit" class="btn btn-block form-control" style="background-color:#0066cc; color:#FFF;" >登　录
        </button>
      </div>-->
					<div class="row">
						<div class="col-xs-12">
							<div class="form-group">
								<div class="input-group m-b"><span class="input-group-btn">
                        				<button type="button" class="btn btn-default" disabled="" style="padding-top: 10px;padding-bottom: 8px;width: 60px;"><i class="icon-user"  style="color:#086d87;"></i></button> </span> 
                        				<input id="p_user_id" name="username" type="text" class="form-control" placeholder="请输入用户名" value="">
								</div>
							</div>
							<div class="form-group">
								<div class="input-group m-b"><span class="input-group-btn">
                        				<button type="button" class="btn btn-default" disabled="" style="padding-top: 10px;padding-bottom: 8px;width: 60px;"><i class="icon-lock" style="color:#086d87;"></i></button> </span> 
                        				<input  id="p_passwd_id" name="password" type="password" class="form-control"  placeholder="请输入密码" value="">
								</div>
							</div>
							<div class="form-group">
								<button class="btn btn-block" type="button" onclick="login()" style="background-color:#086d87; color:#FFF;"><strong>登　录</strong></button>
							</div>
						</div>
					</div>

				</form>
			</div>
		</div>
	</body>

	<script>
		$(document).ready(function(e) {
			var h = $(document).innerHeight();
			var w = $(document).innerWidth();
			$("#login_page").height(h);
			var h1 = $("#logo_area").height();
			$("#input_area").css("margin-top", (h - 441) / 2 - h1 - 10);
			$("#input_area").css("margin-left", (w - 431) / 2);
		});

		function login() {
			if($('#p_user_id').val()=='' || $('#p_passwd_id').val()==''){
				toastr.warning('请输入用户名和密码');
				return ;
			}
			
//			$("#inputZoon").submit();
			$.post('/login', $("#inputZoon").serialize() ,function(data){
				var rs=JSON.parse(data);
                if(rs.ErrorCode==='0'){
                		toastr.error(rs.ErrorMsg);
                    return ;
                }else {
                    window.location.href='/platform/LoginIndexPage';
                }
			});
		}

		$(document).keydown(function(e) {
			if(e.keyCode == '13') {
				setTimeout(login, 200)
			}
		});
	</script>

</html>