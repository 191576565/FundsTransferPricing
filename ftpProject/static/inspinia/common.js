$.extend($.fn.dataTable.defaults, {
	iDisplayLength: 15,
	bLengthChange: false,
	bFilter: false,
	bProcessing: true,
	bServerSide: true,
	autoWidth:false,
//	aaSorting: [
//		[0, "desc"]
//	],
	bSort:false,
	sPaginationType: "full_numbers",
	"language": {
					"processing": "处理中...",
					"lengthMenu": "显示 _MENU_ 项结果",
					"zeroRecords": "没有匹配结果",
					"info": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
					"infoEmpty": "显示第 0 至 0 项结果，共 0 项",
					"infoFiltered": "(由 _MAX_ 项结果过滤)",
					"infoPostFix": "",
					"search": "搜索:",
					"searchPlaceholder": "搜索...",
					"url": "",
					"emptyTable": "表中数据为空",
					"loadingRecords": "载入中...",
					"infoThousands": ",",
					"paginate": {
						"first": "首页",
						"previous": "上页",
						"next": "下页",
						"last": "末页"
					},
					"aria": {
						paginate: {
							first: '首页',
							previous: '上页',
							next: '下页',
							last: '末页'
						},
						"sortAscending": ": 以升序排列此列",
						"sortDescending": ": 以降序排列此列"
					},
					"decimal": "-",
					"thousands": ""
			},
	fnServerData: function(url, aoData, fnCallback, oSettings) {
		var sortArr = [];
		var orderArr = [];
		var params = [];
		$(aoData).each(function() {
			if(this.name == 'iDisplayStart') {
				params.push({
					name: 'offset',
					value: this.value
				});
			} else if(this.name == 'iDisplayLength') {
				params.push({
					name: 'limit',
					value: this.value
				});
			} else if(this.name.indexOf('iSortCol') > -1) {
				sortArr.push(oSettings.aoColumns[this.value].data);
			} else if(this.name.indexOf('sSortDir') > -1) {
				orderArr.push(this.value);
			} else if(this.name != 'sEcho' && this.name != 'iColumns' && this.name != 'sColumns' && this.name.indexOf('mDataProp') == -1 && this.name.indexOf('bSortable') == -1 && this.name != 'iSortingCols') {
				params.push(this)
			}
		})
		if(sortArr.length > 0) {
			params.push({
				name: 'sort',
				value: sortArr.join(',')
			});
			params.push({
				name: 'order',
				value: orderArr.join(',')
			});
		}
		params.push({
			name: 'rr',
			value: Math.random()*100000
		});
		$.get(url, params, function(result) {
			var json = {};
			var rs=JSON.parse(result);
			if(rs && rs.total != null) {
				json.iTotalRecords = rs.total;
				json.iTotalDisplayRecords = rs.total;
				if(rs.total!=0){
					json.aaData = rs.rows
				}else {
					json.aaData = []
				}
			} else {
				if(rs===null){
					json.aaData = []
				}else {
					json.aaData = rs
				}
			}
			$(oSettings.oInstance).trigger('xhr', [oSettings, json]);
			fnCallback(json);
		})
	}
});

Date.prototype.Format = function(fmt) { //author: meizz
	var o = {
		"M+": this.getMonth() + 1, //月份
		"d+": this.getDate(), //日
		"h+": this.getHours(), //小时
		"m+": this.getMinutes(), //分
		"s+": this.getSeconds(), //秒
		"q+": Math.floor((this.getMonth() + 3) / 3), //季度
		"S": this.getMilliseconds() //毫秒
	};
	if(/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
	for(var k in o)
		if(new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
	return fmt;
}

function appendOption(dist, data, value, text){
	var rs=JSON.parse(data);
	if(rs!=null){
		rs.forEach(function(obj){
			$("#"+dist).append("<option value='"+obj[value]+"'>"+obj[text]+"</option>");
		});
	}
}


function getTopWinow(){
    var p = window;
    while(p != p.parent){
        p = p.parent;
    }
    return p;
}

function ReplaceAll(str, array){
	array.forEach(function(e){
		 while (str.indexOf(e[0]) >= 0){
           str = str.replace(e[0], e[1]);
        }
	});
    return str;
 }

$(function(){
	$.ajaxSetup({
		cache: false,
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
	
	$('input, textarea').placeholder();
})