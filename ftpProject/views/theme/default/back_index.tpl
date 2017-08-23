<style type="text/css">
	#wrap .col-sm-12{
		margin: 20px 10px;
	}
</style>
<div style="background-size: cover;" class="theme-bg-color">
	<div id="h-left-tool-bar" class="h-left-tool-bar" style="z-index: 1;">
		<div class="h-left-btn-desk" data-toggle="tooltip" data-placement="right" title="菜单列表"><i class="icon-columns"></i></div>
		<div class="h-left-btn-user" data-toggle="tooltip" data-placement="right" title="用户" onclick="showChangePassDialog()"><i class="icon-user"></i></div>
		<div class="h-left-btn-off" data-toggle="tooltip" data-placement="right" title="安全退出" onclick="LogOut()"><i class="icon-off"></i></div>
		<div class="h-left-btn-menu" data-toggle="tooltip" data-placement="right" title="返回主菜单" onclick="H_LEFT_BAR.H_HomePage()"><i class="icon-th-large"></i></div>
	</div>
	<div id="bigdata-platform-subsystem" class="container-fluid" style="padding-left:55px; overflow: hidden;">
		<div class="row" style="position: relative; height: 60px;">
			<div class="col-sm-12 col-md-12 col-lg-12" style="font-size: 30px;width: 100%;padding-left: 35px;text-align: left;padding-top: 5px;padding-bottom: 8px;">
				<h4 id="huuid" style="color: #ffffff; font-size: 30px; font-weight: 700; height: 40px; line-height: 40px;">天健金管资金转移定价系统</h4>
			</div>
		</div>
		<div id="wrap" class="row" style="height: 10px; padding-right: 1px; overflow: hidden;">
			<div id="h-system-service" class="col-sm-12">
			</div>
			<div id="h-mas-service" class="col-sm-12">
			</div>
			<div id="h-other-service"  class="col-sm-12">
			</div>
		</div>
	</div>
</div>

<div class="modal inmodal info" id="changePassDialog" tabindex="-1" role="dialog" aria-hidden="true">
	<div class="modal-dialog">
		<div class="modal-content animated bounceInRight">
			<div class="modal-header" style="padding: 0px 0px;">
				<h2 id="addUserInfo">修改密码</h2>
				<input type="hidden" id="type"/>
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
				<button type="button" class="btn btn-primary" onclick="changPass()">保存</button>
				<button type="button" class="btn btn-warning" data-dismiss="modal">取消</button>
			</div>
		</div>
	</div>
</div>

<script type="text/javascript">
	$(function(){
		$.ajaxSetup({
		    complete: function(xhr,status) {
		        var sessionStatus = xhr.getResponseHeader('status');
		        if(sessionStatus == '403') {
		            var top = getTopWinow();
		            swal("提示!", "由于您长时间没有操作, session已过期, 请重新登录.", "error");
		            setTimeout(function(){
		            		top.location.href = '/'; 
		            }, 2000);
		        }
		    }
		});
	});
	
	function getTopWinow(){
	    var p = window;
	    while(p != p.parent){
	        p = p.parent;
	    }
	    return p;
	}

	function showChangePassDialog(){
		$("#plat-change-passwd")[0].reset();
		$('#changePassDialog').modal('show');
	}
	
	function changPass(){
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
				var rs=JSON.parse(data);
                if(rs.ErrorCode==='1'){
					toastr.success('密码修改成功！');
					window.location.href = "/"
                }else {
                    toastr.error(rs.ErrorMsg);
                    return ;
                }
			}
		})
	}
	
	function LogOut(){
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
					window.location.href = "/"
				}
			})
		});
	}

	/*
	 * 自动调整icon在框中的位置,使其水平与竖直方向居中显示
	 * */

	function go_entry(e){

		var id = $(e).attr("data-id");

		var quit = function(){
			window.location.href="/"
		};

		var err = function(dt){
			alert(dt)
		};

		var succ = function(d){
			$("#indexHtmlContent").html($("body").html())
			$("#bigdata-platform-subsystem").html(d)
		};

//		HAjaxRequest({
//			url:'/platform/select',
//			data:{Id:id},
//			dataType:'text',
//			error:err,
//			success:succ
//		});
		$.get('/platform/select', {Id:id,r:Math.random()*100000000000}, succ);
	}


	$(document).ready(function(){
		var succ = function(data){
			var rs=JSON.parse(data);
			if(rs!=null){
				$('#wrap').html('');
				rs.forEach(function(e){
					if(e.Res_up_id==='-1'){
						var temp="<div class='' style='margin-left: 80px;margin-bottom:20px'><span class='tile-group-title' style='font-size:20px;color:white;'>"+e.Res_name+"</span><div class='tile-container' style='padding-top:10px;' resid='"+e.Res_id+"'></div></div>"
						$('#wrap').append(temp);
					}else {
						var temp="<div data-id='"+e.Res_id+"'  onclick='go_entry(this)' data-url='"+e.Res_url+"' class='"+e.Res_class+" fg-white' data-role='tile' data-role='tile' style='background-color:"+e.Res_bg_color+"'>"
								+"<div class='tile-content iconic'><span class='icon'><img src='"+e.Res_img+"'></span></div><div class='tile-label'>"+e.Res_name+"</div>";
						$("div[resid='"+e.Res_up_id+"']").append(temp);
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
		$.get('/platform/MenuPage', {TypeId:0,Id:'-1',r:Math.random()*1000000000}, succ);
	});

	//调整主菜单的长度和宽度
	$(document).ready(function(){
		var hh = document.documentElement.clientHeight;
		$(".container-fluid").height(hh);
		$("#h-left-tool-bar").height(hh);
		$("#wrap").height(hh-60);
		$(".tile").click(function(){
			go_entry(this)
		})
		$("#wrap").mCustomScrollbar();

	});

	window.onload = function () {
		if (typeof history.pushState === "function") {
			history.pushState("jibberish", null, null);
			window.onpopstate = function () {
				history.pushState('newjibberish', null, null);
				// Handle the back (or forward) buttons here
				// Will NOT handle refresh, use onbeforeunload for this.
				H_LEFT_BAR.H_HomePage()
			};
		}
		else {
			var ignoreHashChange = true;
			window.onhashchange = function () {
				if (!ignoreHashChange) {
					ignoreHashChange = true;
					window.location.hash = Math.random();

					// Detect and redirect change here
					// Works in older FF and IE9
					// * it does mess with your hash symbol (anchor?) pound sign
					// delimiter on the end of the URL
				}
				else {
					ignoreHashChange = false;
				}
			};
		}
	}
	$("[data-toggle='tooltip']").tooltip();
</script>
