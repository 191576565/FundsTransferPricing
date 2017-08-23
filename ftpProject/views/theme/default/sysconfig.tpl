
<nav id="wrapper" class="navbar-static-side theme-bg-color" style="margin-left: 0px;">
    <div class="H-logo-area" style="font-size: 30px;width: 100%;padding-left: 20px;text-align: left;margin-bottom: 20px;">
        <img src="/static/theme/default/img/sys.png" alt="" width="500px" height="80px"/>
    </div>
    <div class="col-sm-12 col-md-12 col-lg-12" id="H-main-menu">
        <!--<div id="h-system-service" class="col-sm-12 col-md-6 col-lg-4">
        </div>
        <div id="h-mas-service" class="col-sm-12 col-md-6 col-lg-4">
        </div>
        <div id="h-other-service"  class="col-sm-12 col-md-6 col-lg-4">
        </div>-->
    </div>
</nav>
<div id="page-wrapper" class="container-fluid gray-bg" style="margin: 0px 0px;">
    <div class="row H-main-content">
        <div class="col-md-12 col-sm-12 col-lg-12" id="h-main-content" style="padding: 0px 0px;">
            <div data-id="homepage" class="active">

                <script>
//                  // 创建moudle1
//                  var rootMoudle = angular.module('moudle1', []);
//                  rootMoudle.controller("controller1",function($scope){$scope.name="aty"});
//
//                  // 创建moudle2
//                  var m2 = angular.module('moudle2', []);
//                  m2.controller("controller2",function($scope){$scope.name="aty"});
//
//
//                  // 页面加载完成后,再加载模块
//                  angular.element(document).ready(function() {
//                      angular.bootstrap(document.getElementById("div1"),["moudle1"]);
//                      angular.bootstrap(document.getElementById("div2"),["moudle2"]);
//                  });
                </script>
            </div>
        </div>
    </div>
    <div class="row H-content-tab">
        <div class="H-tab-bar pull-left" id="H-tab-left" style="position: relative;width: 100%;">
			<button class="H-left-tab" onclick="leftTabShow()" style="position: absolute;left: 0px;top: 0px;z-index: 233;"><i class="glyphicon glyphicon-backward"></i></button>
			<!--<button class="H-left-tab active-tab" data-id="homepage" onclick="changetab(this)">首页</button>-->
			<nav class="H-tabs-index" style="position: absolute; left: 40px;right: 40px;"></nav>
    			<button class="H-right-tab H-right-tab-padding" onclick="rightTabShow()" style="position: absolute;right: 0px;top: 0px;z-index: 233;"><i class="glyphicon glyphicon-forward"></i></button>
		</div>
        <div class="H-tab-bar pull-right" id="H-tab-right">
            <!--<button class="H-right-tab H-right-tab-padding" onclick="backIndex()"><i class="icon-reply"></i>&nbsp;返回</button>-->
            <!---->
            <!--<div class="H-right-tab dropup">
                <button class="dropdown" type="button" id="dropdownMenu1" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                    关闭操作<span class="caret"></span>
                </button>
                <ul role="menu" class="dropdown-menu dropdown-menu-right" style="text-align: center;">
                    <li onclick="lockCurrentTab()"><a>锁定当前选项卡</a></li>
                    <li class="divider"></li>
                    <li onclick="closeOtherTab()"><a>关闭其他选项卡</a></li>
                    <li onclick="closeAllTab()"><a>关闭全部选项卡</a></li>
                </ul>
            </div>-->
            <!--<button class="H-right-tab H-right-tab-padding" onclick="rightTabShow()"><i class="glyphicon glyphicon-forward"></i></button>-->
        </div>
    </div>
</div>

<script type="text/javascript">
	$(function(){
		$.ajaxSetup({
			scriptCharset:'utf-8',
		    complete: function(xhr,status) {
		        var sessionStatus = xhr.status;
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

    $(document).ready(function(){
        var succ = function(data){
			var rs=JSON.parse(data);
			if(rs!=null){
				$('#H-main-menu').html('');
				rs.forEach(function(e){
					if(e.Res_up_id==='101000000000'){
						var temp="<div class='col-sm-4'><div class='tile-group double' style='margin-left:60px;'><span class='tile-group-title'>"+e.Res_name+"</span><div class='tile-container' resid='"+e.Res_id+"'></div></div></div>"
						$('#H-main-menu').append(temp);
					}else {
						var temp="<div data-id='"+e.Res_id+"' data-url='"+e.Res_url+"' class='"+e.Res_class+" fg-white' data-role='tile' onclick='platMenuService(this)' data-role='tile' style='background-color:"+e.Res_bg_color+"'>"
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

        $.get('/platform/MenuPage',{TypeId:1,Id:'101000000000',r:Math.random()*100000000000},succ)
    })

    /*
    * 调整页面宽度和高度
    * */
    $(document).ready(function(){
        var hwindow = document.documentElement.clientHeight;
        var wwindow = document.documentElement.clientWidth;
        $("#wrapper").height(hwindow);
        $("#page-wrapper").height(hwindow);
        $("#main-menu-bar").height(hwindow);
        $("#h-main-content").height(hwindow - 42);
        $("#page-wrapper").width(wwindow-80);
        $(".navbar-static-side").width(wwindow-50);

        //$(".navbar-static-side").mCustomScrollbar();
        $(".navbar-static-side").mCustomScrollbar({
            axis:"y",
            theme:"dark-thin",
            scrollSpeed:100,
        });
        $('.mCSB_scrollTools').width(0);
    });


    var platMenuService = function(e){

        subSysIndex.hideWrapper();

        var flag = false;
        var url = $(e).attr("data-url");
        var data_id = $(e).attr("data-id");
        var name = $(e).find("div:last").html();
        var optHtml = '<span data-id="'+data_id+'" ' +
                'class="H-left-tab active-tab" ' +
                'onclick="changetab(this)">'+name+'' +
                '<i onclick="closetab(this)" ' +
                'class="icon-remove-sign H-gray-close"></i>&nbsp;' +
                '<i class="icon-refresh H-gray-close" onclick="refresh('+data_id+')"></i>'+
                '</span>'

        $(".H-tabs-index").find("span").each(function(index,element){
            if (data_id == $(element).attr("data-id")){
                changetab(element)
                flag = true;
                return false;
            }
        });
		
		$.get(url,function(data){}); //验证session是否失效
		if (flag == false){
			$(".active-tab").removeClass("active-tab");
			$(".H-tabs-index").prepend(optHtml);
	
			$("#h-main-content").find("div.active").removeClass("active").addClass("none")
	//					var cot = '<div data-type='frame' data-id=''+data_id+'" class='active'>'+data+'</div>';
			var cot = "<div data-type='frame' data-id='"+data_id+"' class='active' style='width:100%;height:100%;padding:0px 0px;'>"
						+"<iframe src='"+url+"' width='100%' height='100%' scrolling='no'></iframe>"
						+"</div>";
//			var cot = "<iframe src='"+"/views/platform/resource/UserInfoPage.tpl"+"' width='100%' height='100%'></iframe>";
			$("#h-main-content").prepend(cot);
		}

//      if (flag == false){
//          $.ajax({
//              type:"get",
//              url:url,
//              cache:false,
//              async:true,
//              dataType:"text", error: function(){
//                  setTimeout(redirectLoginPage,300);
//              },
//              success: function(data){
//                  $(".active-tab").removeClass("active-tab");
//                  $(".H-tabs-index").append(optHtml);
//
////                    (function(){
////                        do{
////                            var maxLT = (function(){
////                                var rt = $("#H-tab-right").width();
////                                var sl = $("#wrapper").width();
////                                var ww = document.documentElement.clientWidth;
////                                return ww - sl - rt;
////                            })();
////                            var lt = $("#H-tab-left").width();
////                            if (lt >= maxLT - 20){
////                                var status = $(".H-tabs-index").find("span:visible:eq(0)").css("display")
////
////                                if ( typeof status == "undefined"){
////                                    break
////                                }
////                                $(".H-tabs-index").find("span:visible:eq(0)").hide();
////                            }
////                        }while(lt >= maxLT - 20)
////                    })();
//
//                  $("#h-main-content").find("div.active").removeClass("active").addClass("none")
//                  var cot = '<div data-type="frame" data-id="'+data_id+'" class="active">'+data+'</div>';
//                  $("#h-main-content").append(cot);
//              }
//          });
//      }
    };

    var closetab = function(e){
        var id = $(e).parent().attr("data-id");
        if ($(e).parent().hasClass("active-tab")){
            var pobj = $(e).parent().prev("span");
            var pid = $(pobj).attr("data-id");
            var nobj = $(e).parent().next("span");
            var nid = $(nobj).attr("data-id");

            $(e).parent().remove();
            $("#h-main-content").find("div[data-id='"+id+"']").remove();
            if (pid == undefined){
                if (nid == undefined){
                    id = "homepage"
                } else {
                    id = nid
                }
            } else {
                id = pid
            }

            $("#h-main-content").find("div[data-id='"+id+"']").removeClass("none").addClass("active");
            $(".H-left-tab").each(function(index,element){
                if (id == $(element).attr("data-id")){
                    $(element).addClass("active-tab")
                }
            });
            var leftStyle = leftTabShow();
            if (leftStyle == false){
                rightTabShow()
            }
        }else{
            $(e).parent().remove();
            $("#h-main-content").find("div[data-id='"+id+"']").remove();
            var leftStyle = leftTabShow();
            if (leftStyle == false){
                rightTabShow()
            }
        }
        
        if($('.H-left-tab').length===1){ //所有tab已关闭
			subSysIndex.showWrapper();
		}
//      window.event.cancelBubble = true;
		cancelBubble();
    };

    var leftTabShow = function(){
//      var firstObj = $(".H-tabs-index").find("span:visible:eq(0)")
//      var preOjb = $(firstObj).prev("span").attr("data-id")
//      if (preOjb == undefined){
//          return false
//      }else{
//          var lastShowItem = null;
//          $(firstObj).nextAll("span").each(function(index,element){
//              if ("none" ==  $(element).css("display")){
//                  return false;
//              }else{
//                  lastShowItem = element;
//              }
//          });
//
//          $(firstObj).prev("span").show();
//          do {
//              var maxLT = (function(){
//                  var rt = $("#H-tab-right").width();
//                  var sl = $("#wrapper").width();
//                  var ww = document.documentElement.clientWidth;
//                  return ww - sl - rt;
//              })();
//              var lt = $("#H-tab-left").width();
//              if (lt >= maxLT - 20){
//                  $(lastShowItem).hide();
//                  lastShowItem = $(lastShowItem).prev("span");
//              }
//          }while( lt >= maxLT -20)
//          return true;
//      }
		var na=$('.H-tabs-index').width()
		,total=0;
		$('.H-tabs-index').children().each(function(){
			total+=$(this).width()+30;
		});
		
		if(na>total){
			return ;
		}else {
			$('.H-tabs-index span:first').insertAfter('.H-tabs-index span:last');
		}
    };

    var rightTabShow = function(){
//      var firstObj = $(".H-tabs-index").find("span:visible:eq(0)")
//      $(firstObj).nextAll("span").each(function(index,element){
//          if ("none" ==  $(element).css("display")){
//              $(element).show();
//              do {
//                  var maxLT = (function(){
//                      var rt = $("#H-tab-right").width();
//                      var sl = $("#wrapper").width();
//                      var ww = document.documentElement.clientWidth;
//                      return ww - sl - rt;
//                  })();
//                  var lt = $("#H-tab-left").width();
//                  if (lt >= maxLT - 20){
//                      $(".H-tabs-index").find("span:visible:eq(0)").hide();
//                  }
//              }while( lt >= maxLT -20)
//              return false;
//          }
//      });

		var na=$('.H-tabs-index').width()
		,total=0;
		$('.H-tabs-index').children().each(function(){
			total+=$(this).width()+30;
		});
		
		if(na>total){
			return ;
		}else {
			$('.H-tabs-index span:last').insertBefore('.H-tabs-index span:first');
		}
    }

    var changetab = function(e){
        $(".active-tab").removeClass("active-tab")
        $(e).addClass("active-tab")
        var id = $(e).attr("data-id");
        $("#h-main-content").find("div.active").removeClass("active").addClass("none")
        $("#h-main-content").find("div[data-id='"+id+"']").removeClass("none").addClass("active");
    };

    var closeAllTab = function(){
        $(".H-tabs-index").find("span").remove();
        $("#h-main-content").find("div[data-type='frame']").remove();
        $("#h-main-content").find("div[data-id='homepage']").removeClass("none").addClass("active");
        $(".H-left-tab").each(function(index,element){
            if ("homepage" == $(element).attr("data-id")){
                $(element).addClass("active-tab")
            }
        })
        
        if($('.H-left-tab').length===1){ //所有tab已关闭
			subSysIndex.showWrapper();
		}
    };

    var closeOtherTab = function(){
        var id = new Array();
        var i = 0;
        $(".H-tabs-index").find("span").each(function(index,element){
            if ($(element).hasClass("active-tab") || $(element).hasClass("tab-lock")){
                id[i++] = $(element).attr("data-id");
            } else {
                $(element).remove()
            }
        });

        $("#h-main-content").find("div[data-type='frame']").each(function(index,element){
            if( id.indexOf($(element).attr("data-id")) > -1){

            }else{
                $(element).remove()
            }
        });
    };

    var lockCurrentTab = function(){
        $(".H-tabs-index").find("span.active-tab").addClass("tab-lock").find("i").remove()
    };

    var getHomePage = function(){
        $.ajax({
            type:"get",
            url:"/platform/HomePage",
            cache:false,
            async:true,
            dataType:"text", error: function(){
                setTimeout(redirectLoginPage,300);
            },
            success: function(data){
                $("#h-main-content").html("<div data-id='homepage' class='active'>"+data+"</div>")
            }
        });
    }

    /*
    * 绑定内容框,当点击内容框时,隐藏左侧菜单栏
    * */
    $("#page-wrapper").click(function(){
        subSysIndex.hideWrapper();
    });

    /*
    * 左侧菜单栏事件绑定,当点击左侧菜单栏时,内容框右移30%
    * */
    $(".h-left-btn-desk").click(function(){
        subSysIndex.showWrapper();
    });

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
        
        $("[data-toggle='tooltip']").tooltip();
    });

    window.onresize = function(){
    		var hw = document.documentElement.clientWidth;
        var hh = document.documentElement.clientHeight;
        if($("#wrapper").width()==0){
        		$("#wrapper").width(0);
        }else {
        		$("#wrapper").width(hw);
        }
	    $("#page-wrapper").width(hw-80);
	    $("#page-wrapper").css('right', '0'); //ys
        
        $("#wrapper").height(hh);
        $("#page-wrapper").height(hh)
        $(".H-main-content").height(hh-42);
        $("#h-left-tool-bar").height(hh);
    }
    
    function refresh(dataId){ //刷新
    		var iframe=$('div[data-id="'+dataId+'"]').find('iframe')[0];
//  		iframe.attr('src', iframe.attr('src'));
		$.get('/platform/MenuPage', function(){
			iframe.contentWindow.location.reload(true);
		});
    }
</script>