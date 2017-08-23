var todate = function(a) {
		return a.substr(0, 10)
	},
	HAjaxRequest = function(a) {
		var b = {
			type: "get",
			url: "",
			data: "",
			cache: !0,
			async: !1,
			dataType: "json",
			error: function(b) {
				console.log(b.readyState, b.responseText, b.statusText)
			},
			success: function(b) {}
		};
		$.extend(!0, b, a);
		"delete" == b.type.toLowerCase() ? (b.data._Method = "Delete", $.ajax({
			type: "post",
			url: b.url,
			cache: b.cache,
			async: b.async,
			data: b.data,
			dataType: b.dataType,
			error: function(a) {
				b.error(a)
			},
			success: function(a) {
				b.success(a)
			}
		})) : $.ajax({
			type: b.type,
			url: b.url,
			cache: b.cache,
			async: b.async,
			data: b.data,
			dataType: b.dataType,
			error: function(a) {
				b.error(a)
			},
			success: function(a) {
				b.success(a)
			}
		})
	},
	H_LEFT_BAR = {
		H_HomePage: function() {
//			window.event.cancelBubble = !0;
//			$("body").load("/platform/IndexPage")
			$.get('/platform/MenuPage', function(){
				setTimeout(window.location.reload(), 1000);
			});
		},
		HLogOut: function() {
			window.event.cancelBubble = !0;
			modal.confirm(function() {
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
			})
		},
		UserMgrInfo: function() {
			modal.newModal(function() {
				$.ajax({
					type: "post",
					url: "/platform/passwd",
					data: $("#plat-change-passwd").serialize(),
					cache: !1,
					async: !1,
					dataType: "text",
					error: function(a, b, e) {
						alert("error");
						alert(a.readyState);
						alert(a.responseText);
						alert(a.statusText)
					},
					success: function(a) {
						alert("success");
						alert(a)
					}
				})
			}, "\u5bc6\u7801\u4fee\u6539", $("#mas-passwd-prop").html())
		}
	},
	subSysIndex = {
		hideWrapper: function() {
			$("#wrapper").width(0);
			$("#page-wrapper").css("right", "0")
		},
		showWrapper: function() {
			"0" == $("#wrapper").width() && ($("#wrapper").css("width", "100%"), $("#page-wrapper").css("right",
				"-100%"))
//			if("0" == $("#wrapper").width()){
//				document.getElementById('wrapper').style.width="100%";
//				document.getElementById('page-wrapper').style.right="-100%";
//			}
		},
		adjustTabShow: function() {
			var a = [],
				b = 0;
			$(".H-tabs-index").find("span").each(function(e, f) {
				"none" == $(f).css("display") && (a[b++] = f)
			});
			setTimeout(function(b, a) {
				do {
					var c;
					c = $("#H-tab-right").width();
					var d = $("#wrapper").width();
					c = document.documentElement.clientWidth - d - c;
					d = $("#H-tab-left").width();
					if(d >= c - 20) break;
					else {
						var g = b[--a];
						$(g).show()
					}
				} while (d <= c - 20 && 0 < a)
			}(a, b), 1E3)
		}
	},
	HzwTools = {
		isSupportHTML5: function() {
			return window.applicationCache ? !0 : !1
		},
		handleCheckBox: function(a) {
			a = $(a).closest("tr");
			$(a).find("th:eq(0) input").is(":checked") ? $(a).nextAll("tr").each(function(b, a) {
				$(a).find("td:eq(0) input").prop("checked", !0)
			}) : $(a).nextAll("tr").each(function(a, e) {
				$(e).find("td:eq(0) input").prop("checked", !1)
			})
		}
	};