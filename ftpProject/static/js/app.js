$(function() {
	//alert(window.location.href);

	//除登录界面外,检查sessionStorage中是否有登陆信息
//	var url = window.location.href.split('/');
//	var page = url[url.length - 1];
//	//alert(page);
//	if(page != 'login.html') {
//		var user = JSON.parse(window.sessionStorage.getItem("userInfo"));
//		if(!user) {
//			console.log(url);
//			window.location.href = "/"+url[3]+"/index.html";
//			return;
//		}
//
//		if(page == 'start-screen.html') {
//			$('.sub-header').html('你好,' + user.username);
//		}
//	}

});

//用户登录
//function login() {
//	//alert($('#loginForm').serialize());
//	var username = $('#user_name').val();
//	var password = $('#user_password').val();
//
//	var userInfo = {
//		username: username,
//		password: password
//	}
//
//	//ajax登录验证
////	if(username != '' && password != '') {
////		window.sessionStorage.setItem('userInfo', JSON.stringify(userInfo));
////		//页面跳转
////		window.location.href = 'index.html';
////	} else {
////		alert('请输入正确的用户名及密码！');
////	}
//	$.post('/login', userInfo, function(data){
////		console.log(data);
//	});
//}

//注销
function logout() {
	window.sessionStorage.clear();
	window.location.href = 'login.html';
}

//显示新增编辑窗口
function showDialog(id) {
	var dialog = $(id).data('dialog');
	dialog.open();
}