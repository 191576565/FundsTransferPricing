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

		<title>定价规则配置</title>

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
		<!-- Data picker -->
		<script src="/static/inspinia/js/plugins/datapicker/bootstrap-datepicker.js"></script>

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
				width: 60px;
			}
			
			.selected {
				background-color: #b7eafb !important;
			}
			
			.input-group-addon {
				border: 1px solid #E5E6E7 !important;
			}
			
			.ztree * {
				font-size: 16px;
			}
			/*.ztree li {
				line-height: 20px;
			}*/
			
			.ztree li a.curSelectedNode {
				height: 20px !important;
			}
			.select2-selection{
				border: 1px solid #e5e6e7 !important; 
			}
			.modal-header .close{
				padding-top: 4px !important;
				padding-right: 8px !important;
			}
		</style>
	</head>

	<body>
		<div id="wrapper" style="height: 100%;">
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
			<div class="wrapper wrapper-content animated fadeInRight full-height">
				<div class="row">
					<div class="col-sm-12">
						<div class="ibox float-e-margins full-height">
							<div class="ibox-title">
								<div class="row">
									<div class="col-xs-4">
										<h3>定价规则配置</h3>
									</div>
									<div class="col-xs-8 text-right">
										<div class="form-inline">
											<div class="form-group">
												<input type="text" placeholder="请输入业务单元编号" style="width: 200px;" id="searchBusiz" class="form-control">
											</div>
											<button class="btn btn-default searchBtn" type="button" onclick="searchBusiz()" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
										</div>
									</div>
								</div>
							</div>
							<div class="ibox-content" style="overflow-x: auto;height: 95%;">
								<div class="row full-height">
					<div class="col-xs-5 full-height">
						<!--<div class="ibox float-e-margins full-height">
							<div class="ibox-title">
								<div class="row">
									<div class="col-xs-4">
										<h3>定价规则配置</h3>
									</div>
									<div class="col-xs-8 text-right">
										<div class="form-inline">
											<div class="form-group">
												<input type="text" placeholder="请输入业务单元编号" style="width: 200px;" id="searchBusiz" class="form-control">
											</div>
											<button class="btn btn-default searchBtn" type="button" onclick="searchBusiz()" style="margin-bottom: 0px;"><i class="fa fa-search"></i></button>
										</div>
									</div>
								</div>
							</div>-->
							<!--<div class="ibox-content" style="overflow-x: auto;height: 95%;">-->
								<ul id="busizTree" class="ztree" style="padding-left: 30px;padding-top: 20px;"></ul>
							<!--</div>-->
						<!--</div>-->
					</div>
					<div class="col-xs-7" style="">
						<div class="ibox float-e-margins busizInfo" style="border: 1px solid #dcdada;">
							<div class="ibox-title" style="border: 0px;">
								<div class="ibox-tools">
									<h3 style="float: left;margin-top: 0px;" id="curveTitle"><i class="fa fa-info-circle"></i>&nbsp;规则配置信息</h3>
									<input type="hidden" name="Curve_id" value="" />
									<input type="hidden" name="Domain_id" value="" />
									<a class="" style="float: right;" onclick="">
										<!--<i class="fa fa-times"></i>-->
									</a>
									<button class="btn btn-info res" type="button" onclick="updateBusiz()" resid='30208010202000'>更新</button>
									<button class="btn btn-danger res" type="button" onclick="delBusiz()" resid='30208010203000'>删除</button>
								</div>
							</div>
							<div class="ibox-content">
								<form role="form" id="bInfo">
									<div class="row">
										<div class="col-xs-4">
											<div class="form-group">
												<label>业务单元编号:</label>
												<input type="text" placeholder="" class="form-control" id="bCode" name="Busiz_id" readonly="">
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>业务单元描述:</label>
												<input type="text" placeholder="" class="form-control" id="bName" name="Busiz_desc">
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>上级业务单元:</label>
												<div class="input-group">
													<input type="hidden" id="bPCode" name="Busiz_up_id" />
													<input type="text" class="form-control" id="bPName" name="Busiz_up_name" readonly="">
													<span class="input-group-btn" onclick="showLevelDialog('update')">
														<button type="button" class="btn" id="PBtn"><i class="fa fa-level-up"></i></button> 
													</span>
												</div>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>业务单元类型:</label>
												<select name="Busiz_type" id="bType" class="form-control">
													<option value="1">叶子</option>
													<option value="0">结点</option>
												</select>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>是否定价:</label>
												<select name="Ftp_flag" id="bPrice" onchange="inputOnchange(this, '2')" class="form-control">
													<option value="1">不定价</option>
													<option value="0">定价</option>
												</select>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>资产负债类型:</label>
												<select name="AL_flag" id="bAlType" class="form-control">
												</select>
											</div>
										</div>
										<!--<div class="col-xs-4">
											<div class="form-group">
												<label>所属域:</label>
												<select name="Domain_id" id="bRegion" class="form-control">
												</select>
											</div>
										</div>-->
									</div>

									<div class="row" id="priceInfo_B">
										<div class="col-xs-4">
											<div class="form-group">
												<label>定价方法:</label>
												<select name="ftp_method_id" id="bMethod" onchange="inputOnchange(this, '4')" class="form-control">
												</select>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>指定曲线:</label>
												<select name="curve_id" id="bCurve" class="form-control">
												</select>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>期限:</label>
												<input type="number" name="term_cd" id="bTerm" class="form-control">
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>期限单位:</label>
												<select name="term_cd_mult" id="bTermUnit" class="form-control">
													<option value="Y" selected="">年</option>
													<option value="M">月</option>
													<option value="D">日</option>
												</select>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>点差值:</label>
												<input type="text" placeholder="" class="form-control" name="point_val" id="bSpot">
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>期限结构:</label>
												<button type="button" class="btn btn-link form-control" id="term_struct_b" onclick="selectTermStructs('1')">期限结构选择</button>
											</div>
										</div>
										<div class="col-xs-12">
											<div class="form-group">
												<label>调节项:</label>
												<div class="input-group full-width">
													<select name="innerAdjust" id="innerAdjust" multiple="multiple" class="full-width" style="">
													</select>
												</div>
											</div>
										</div>
									</div>
									<!--<hr />-->
									<div class="row" style="">
										<!--<div class="col-xs-12" style="text-align: center;">
											<button class="btn btn-info res" type="button" onclick="updateBusiz()" resid='208010202000'>更新</button>
											<button class="btn btn-danger res" type="button" onclick="delBusiz()" resid='208010203000'>删除</button>
										</div>-->
										<!--<div class="col-xs-6">
											
										</div>-->
									</div>
								</form>
							</div>
						</div>
						
						<div class="ibox float-e-margins addBusizInfo" style="border: 1px solid #dcdada;display: none;">
							<div class="ibox-title" style="border: 0px;">
								<div class="ibox-tools">
									<h3 style="float: left;margin-top: 0px;" id="curveTitle"><i class="fa fa-info-circle"></i>&nbsp;新增配置信息</h3>
									<input type="hidden" name="Curve_id" value="" />
									<input type="hidden" name="Domain_id" value="" />
									<a class="" style="float: right;" onclick="closeAddWindow()">
										<i class="fa fa-times"></i>
									</a>
								</div>
							</div>
							<div class="ibox-content">
								<form role="form" id="a_bInfo">
									<div class="row">
										<div class="col-xs-4">
											<div class="form-group">
												<label>业务单元编号:</label>
												<input type="text" placeholder="" class="form-control" id="Busiz_id" name="Busiz_id">
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>业务单元描述:</label>
												<input type="text" placeholder="" class="form-control" id="Busiz_desc" name="Busiz_desc">
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>上级业务单元:</label>
												<div class="input-group">
													<input type="hidden" id="Busiz_up_id" name="Busiz_up_id"  />
													<input type="text" class="form-control" id="Busiz_up_name" name="Busiz_up_name" readonly="">
													<span class="input-group-btn" onclick="showLevelDialog('add')">
														<button type="button" class="btn"><i class="fa fa-tree"></i></button> 
													</span>
												</div>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>业务单元类型:</label>
												<select name="Busiz_type" id="Busiz_type" class="form-control">
													<option value="1">叶子</option>
													<option value="0">结点</option>
												</select>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>是否定价:</label>
												<select name="Ftp_flag" id="Ftp_flag" onchange="inputOnchange(this, '1')" class="form-control">
													<option value="1">不定价</option>
													<option value="0">定价</option>
												</select>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>资产负债类型:</label>
												<select name="AL_flag" id="AlType" class="form-control">
												</select>
											</div>
										</div>
										<!--<div class="col-xs-4">
											<div class="form-group">
												<label>所属域:</label>
												<select name="Domain_id" id="Domain_id_B" class="form-control">
												</select>
											</div>
										</div>-->
									</div>

									<div class="row" id="priceInfo">
										<div class="col-xs-4">
											<div class="form-group">
												<label>定价方法:</label>
												<select name="ftp_method_id" id="ftp_method_id" onchange="inputOnchange(this, '3')" class="form-control">
												</select>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>指定曲线:</label>
												<select name="curve_id" id="curve_id" class="form-control">
												</select>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>期限:</label>
												<input type="number" name="term_cd" id="term_cd" class="form-control">
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>期限单位:</label>
												<select name="term_cd_mult" id="term_cd_mult" class="form-control">
													<option value="Y" selected="">年</option>
													<option value="M">月</option>
													<option value="D">日</option>
												</select>
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>点差值:</label>
												<input type="text" placeholder="" class="form-control" name="point_val" id="point_val">
											</div>
										</div>
										<div class="col-xs-4">
											<div class="form-group">
												<label>期限结构:</label>
												<button type="button" class="btn btn-link form-control" id="term_struct" onclick="selectTermStructs('2')">期限结构选择</button>
											</div>
										</div>
										<div class="col-xs-12">
											<div class="form-group">
												<label>调节项:</label>
												<div class="input-group full-width">
													<select name="inner" id="inner" multiple="multiple" class="full-width" style="">
													</select>
												</div>
											</div>
										</div>
									</div>
									<div class="row">
										<div class="col-xs-12">
											<button class="btn btn-primary btn-block" type="button" onclick="saveBusiz()">保存</button>
										</div>
									</div>
								</form>
							</div>
						</div>
					</div>
				</div>
							</div>
						</div>
					</div>
				</div>
				
			</div>
		</div>

		<div class="modal inmodal" id="upBusiz" tabindex="-1" role="dialog" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 >上级业务单元选择</h2>
					</div>
					<div class="modal-body">
						<ul id="treeDemo" class="ztree"></ul>
					</div>
				</div>
			</div>
		</div>
		
		<div class="modal inmodal" id="termStruct" tabindex="-1" role="dialog" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content animated bounceInRight">
					<div class="modal-header" style="padding: 0px 0px;">
						<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
						<h2 >期限结构选择</h2>
						<input type="hidden" id="domainIdS"/>
						<input type="hidden" id="bCodeS"/>
					</div>
					<div class="modal-body">
						<div class="row s1">
							<div class="col-xs-12">
								<div class="form-group">
									<label>期限结构:</label>
									<div class="input-group full-width termCheckBase">
										<!--<label class="checkbox-inline i-checks"> <input type="checkbox" value="option1">a </label>
                                        <label class="checkbox-inline i-checks"> <input type="checkbox" value="option2"> b </label>
                                        <label class="checkbox-inline i-checks"> <input type="checkbox" value="option3"> c </label>-->
									</div>
								</div>
							</div>
						</div>
						<div class="row s2">
							<div class="col-xs-12">
								<table class="table table-striped table-bordered table-hover" id="termStructTable" style="width: 100%;text-align: center;">
									<thead>
										<tr  style="border-bottom: 1px solid #C1C3C4;">
											<!--<th><input type="checkbox" /></th>-->
											<th style="text-align: center;">期限段</th>
											<th style="text-align: center;">期限比例值</th>
										</tr>
									</thead>
									<tbody>
									</tbody>
								</table>
							</div>
						</div>
					</div>
					<div class="modal-footer steps">
						<button type="button" class="btn btn-white prev" onclick="goStepBusiz('1')" style="float: left;display: none;">上一步</button>
						<button type="button" class="btn btn-white next" onclick="goStepBusiz('2')">下一步</button>
						<button type="button" class="btn btn-primary save" onclick="saveTermStruct()" style="display: none;">保存</button>
					</div>
				</div>
			</div>
		</div>

	</body>
	<script type="text/javascript">
		var re = /^[0-9]+.?[0-9]*$/;
		var inte = /^[-]{0,1}[0-9]{1,}$/;
		var setting = {
			view: {
				addHoverDom: addHoverDom,
				removeHoverDom: removeHoverDom,
				dblClickExpand: false,
				selectedMulti: false
			},
			data: {
				simpleData: {
					enable: true
				}
			},
			callback: {
				//onDblClick: zTreeOnDblClick,
				beforeExpand: beforeExpand,
				onExpand: onExpand,
				onClick: zTreeOnClick
			}
		};
		
		var s = {
			view: {
				dblClickExpand: false,
				selectedMulti: false
			},
			data: {
				simpleData: {
					enable: true
				}
			},
			callback: {
				onDblClick: zTreeOnDblClickSimple,
				onClick: zTreeOnClickSimple
			}
		};
		
		var curExpandNode = null;
		function beforeExpand(treeId, treeNode) {
			var pNode = curExpandNode ? curExpandNode.getParentNode():null;
			var treeNodeP = treeNode.parentTId ? treeNode.getParentNode():null;
			var zTree = $.fn.zTree.getZTreeObj("busizTree");
			for(var i=0, l=!treeNodeP ? 0:treeNodeP.children.length; i<l; i++ ) {
				if (treeNode !== treeNodeP.children[i]) {
					zTree.expandNode(treeNodeP.children[i], false);
				}
			}
			while (pNode) {
				if (pNode === treeNode) {
					break;
				}
				pNode = pNode.getParentNode();
			}
			if (!pNode) {
				singlePath(treeNode);
			}

		}
		function singlePath(newNode) {
			if (newNode === curExpandNode) return;

            var zTree = $.fn.zTree.getZTreeObj("busizTree"),
                    rootNodes, tmpRoot, tmpTId, i, j, n;

            if (!curExpandNode) {
                tmpRoot = newNode;
                while (tmpRoot) {
                    tmpTId = tmpRoot.tId;
                    tmpRoot = tmpRoot.getParentNode();
                }
                rootNodes = zTree.getNodes();
                for (i=0, j=rootNodes.length; i<j; i++) {
                    n = rootNodes[i];
                    if (n.tId != tmpTId) {
                        zTree.expandNode(n, false);
                    }
                }
            } else if (curExpandNode && curExpandNode.open) {
				if (newNode.parentTId === curExpandNode.parentTId) {
					zTree.expandNode(curExpandNode, false);
				} else {
					var newParents = [];
					while (newNode) {
						newNode = newNode.getParentNode();
						if (newNode === curExpandNode) {
							newParents = null;
							break;
						} else if (newNode) {
							newParents.push(newNode);
						}
					}
					if (newParents!=null) {
						var oldNode = curExpandNode;
						var oldParents = [];
						while (oldNode) {
							oldNode = oldNode.getParentNode();
							if (oldNode) {
								oldParents.push(oldNode);
							}
						}
						if (newParents.length>0) {
							zTree.expandNode(oldParents[Math.abs(oldParents.length-newParents.length)-1], false);
						} else {
							zTree.expandNode(oldParents[oldParents.length-1], false);
						}
					}
				}
			}
			curExpandNode = newNode;
		}

		function onExpand(event, treeId, treeNode) {
			curExpandNode = treeNode;
		}
		
		var type;
		function zTreeOnDblClickSimple(event, treeId, treeNode) {
			//alert(treeNode ? treeNode.tId + ", " + treeNode.name : "isRoot");
			if(type === 'add') {
				$('#Busiz_up_id').val(treeNode.id);
				$('#Busiz_up_name').val(treeNode.name);
			} else {
				$('#bPCode').val(treeNode.id);
				$('#bPName').val(treeNode.name);
			}

			$('#upBusiz').modal('hide');
		};

		function zTreeOnClickSimple(event, treeId, treeNode) {
			var treeObj = $.fn.zTree.getZTreeObj(treeId);
			treeObj.expandNode(treeNode, null, null, null);

			return true;
		};
	
		var addBtn=false;
		function addHoverDom(treeId, treeNode) {
			if(treeNode.type === '1') {
				return;
			}
			
			if(!addBtn){ //如果没有新增的权限不显示新增
				return ;
			}
			var sObj = $("#" + treeNode.tId + "_span");
			if(treeNode.editNameFlag || $("#addBtn-" + treeNode.tId).length > 0) return;
			var addStr = "<span class='addBtn res' id='addBtn-" + treeNode.tId +
				"' title='add node' onfocus='this.blur();' resid='30208010201000'>+</span>";
			sObj.after(addStr);
			var btn = $("#addBtn-" + treeNode.tId);
			if(btn) btn.bind("click", function(e) {
				$('#a_bInfo')[0].reset();
				$('.busizInfo').hide();
				$('.addBusizInfo').show();
				var treeObj = $.fn.zTree.getZTreeObj("busizTree");
				var node = treeObj.getNodeByTId($(this).attr('id').split('-')[1]);
				$('#Busiz_up_name').val(node.name.split(' ')[1]);
				$('#Busiz_up_id').val(node.id);
				$('#Domain_id_B').val(node.region);
				//$('#Domain_id_B').attr('disabled','disabled');
				
				$('#Ftp_flag').change();
				$('#ftp_method_id').change();
				
				//阻止冒泡
				e = e || event;
				e.stopPropagation ? e.stopPropagation() : e.cancelBubble = true;
			});
		};

		function removeHoverDom(treeId, treeNode) {
			$("#addBtn-" + treeNode.tId).unbind().remove();
		};

		function zTreeOnClick(event, treeId, treeNode) {
//			var treeObj = $.fn.zTree.getZTreeObj(treeId);
//			treeObj.expandNode(treeNode, null, null, null);
			
			var treeObj = $.fn.zTree.getZTreeObj(treeId);
			treeObj.expandNode(treeNode, null, null, null, true);


			$('#bCode').val(treeNode.id);
			$('#bName').val(treeNode.name.split(' ')[1]);
			if(!treeNode.pId) {
				$('#bPCode').val('-1');
				$('#bPName').val('');
				$('#bPName').attr('disabled', true);
				$('#PBtn').attr('disabled', true);
				
				$('#bAlType').parent().parent().hide();
			} else {
				var parentNode = treeObj.getNodeByTId(treeNode.parentTId);
				$('#bPCode').val(parentNode.id);
				$('#bPName').val(parentNode.name.split(' ')[1]);
				$('#bPName').attr('disabled', false);
				$('#PBtn').attr('disabled', false);
				
				$('#bAlType').parent().parent().show();
			}
			$('#bType').val(treeNode.type);
			$('#bPrice').val(treeNode.isPrice).change();
//			$('#bRegion').val(treeNode.region);
			$('#bAlType').val(treeNode.Al_flag);

			$('#bMethod').val(treeNode.Ftp_method_id).change();
			$('#bCurve').val(treeNode.Curve_id);
			$('#bTerm').val(treeNode.Term_cd);
			$('#bTermUnit').val(treeNode.Term_cd_mult);
			$('#bSpot').val(treeNode.Point_val);

			$('.busizInfo').show();
			$('.addBusizInfo').hide();

			$.get('/platform/FtpAdjustment' + '?r=' + Math.random(100000000000000), function(data) {
				$('#innerAdjust').html('');
				var rs = JSON.parse(data);
				rs.forEach(function(obj) {
					var v=obj.AdjustmentId+" "+obj.AdjtypeId;
					if(treeNode.Adjment_info.indexOf(obj.AdjustmentId) >= 0) {
						$("#innerAdjust").append("<option selected value='" + v + "'>" + obj.AdjustmentName + "</option>");
					} else {
						$("#innerAdjust").append("<option value='" + v + "'>" + obj.AdjustmentName + "</option>");
					}
				});
				
				$('#innerAdjust').select2();
			});

			return true;
		};

		$(function() {
//			#('#wrapper').height(document.body.clientHeight);
			//控制按钮权限
			$.get('/platform/DefaultMenu',{TypeId:2,Id:'30208010200000',r:Math.random()*10000000000000}, function(data){
				var rs=JSON.parse(data);
				//将所有有权限编码的按钮先禁用
//				$('.res').attr('disabled', true);
				$('.res').hide();
				if(rs!=null){
					rs.forEach(function(e){
						if(e.Res_id==='30208010201000'){ //新增非button按钮
							addBtn=true;
						}else { 
//							$("button[resid='"+e.Res_id+"']").attr('disabled', false);
							$("button[resid='"+e.Res_id+"']").show();
						}
					});
				}
			})
			
			$('#innerAdjust').select2();
			$('#inner').select2();
			$('.select2').addClass('full-width');
			$('#bPrice').change();
			iCheckInit();
			
			//获取树数据
			$.get('/platform/FtpBusiz' + '?r=' + Math.random(100000000000000), function(data) {
				$.fn.zTree.init($("#busizTree"), setting, JSON.parse(data)); //树	

				var treeObj = $.fn.zTree.getZTreeObj("busizTree");
			});
			
			$.get('/platform/FtpBusizStruct' + '?r=' + Math.random(100000000000000), function(data) {
				$.fn.zTree.init($("#treeDemo"), s, JSON.parse(data)); //树
			});
			
			$.get('/platform/FtpSysMethod', {r:Math.random()*10000000000000}, function(data) {
				appendOption('ftp_method_id', data, 'FtpMethodId', 'FtpMethodName');
				appendOption('bMethod', data, 'FtpMethodId', 'FtpMethodName');
			});
			$.get('/platform/MasCurveDefine', {r:Math.random()*10000000000000}, function(data) {
				appendOption('curve_id', data, 'CurveId', 'CurveDesc');
				appendOption('bCurve', data, 'CurveId', 'CurveDesc');
			});
//			$.get('/platform/SysDomainInfo', {r:Math.random()*10000000000000}, function(data) {
//				appendOption('Domain_id_B', data, 'DomainId', 'DomainName');
//				appendOption('bRegion', data, 'DomainId', 'DomainName');
//			});
			//资产负债类型
			$.get('/mas/ftp/FtpAlType', {r:Math.random()*10000000000000}, function(data) {
				appendOption('AlType', data, 'Altypeid', 'Altypedesc');
				appendOption('bAlType', data, 'Altypeid', 'Altypedesc');
			});

			$.get('/platform/FtpAdjustment', {r:Math.random()*10000000000000}, function(data) {
//				appendOption('innerAdjust', data, 'AdjustmentId', 'AdjustmentName');
//				appendOption('inner', data, 'AdjustmentId', 'AdjustmentName');
				var rs=JSON.parse(data);
				if(rs!=null){
					rs.forEach(function(obj){
						var v=obj.AdjustmentId+" "+obj.AdjtypeId;
						$("#innerAdjust").append("<option value='"+v+"'>"+obj.AdjustmentName+"</option>");
						$("#inner").append("<option value='"+v+"'>"+obj.AdjustmentName+"</option>");
					});
				}
			});
			
		})
		
		function iCheckInit(){
			$('.i-checks').iCheck({
				checkboxClass: 'icheckbox_square-green',
				radioClass: 'iradio_square-green',
			});
		}
		
		function inputOnchange(obj, type) {
			var v = $(obj).val();
			if(type === '1') {
				if(v === '0') { //定价
					$('#priceInfo').slideDown();
					showAdjust();
				} else { //不定价
					$('#priceInfo').slideUp();
				}
			} else if(type === '2') {
				if(v === '0') { //定价
					$('#priceInfo_B').slideDown();
					$('#bMethod').change();
					showAdjust();
				} else { //不定价
					$('#priceInfo_B').slideUp();
				}
			} else if(type==='3'){
				var infos=[];
				$(obj).parent().parent().nextAll().hide();
				if(v==='103' || v==='102' || v==='105'){ //现金流加权期限法、直接期限法、久期法
					infos=['curve_id'];
					showMoreInfos(infos);
				}else if(v==='101'){ //指定期限法
					infos=['curve_id', 'term_cd', 'term_cd_mult'];
					showMoreInfos(infos);
				}else if(v==='104'){ //沉淀率法
					infos=['curve_id','term_struct'];
					showMoreInfos(infos);
				}else if(v==='106'){ //利率差额法
					infos=['point_val'];
					showMoreInfos(infos);
				}
				showAdjust();
//				refreshPosition();
			}else if(type==='4'){
				var infos=[];
				$(obj).parent().parent().nextAll().hide();
				if(v==='103' || v==='102' || v==='105'){ //现金流加权期限法、直接期限法、久期法
					infos=['bCurve'];
					showMoreInfos(infos);
				}else if(v==='101'){ //指定期限法
					infos=['bCurve', 'bTerm', 'bTermUnit'];
					showMoreInfos(infos);
				}else if(v==='104'){ //沉淀率法
					infos=['bCurve','term_struct_b'];
					showMoreInfos(infos);
				}else if(v==='106'){ //利率差额法
					infos=['bSpot'];
					showMoreInfos(infos);
				}
				showAdjust();
//				refreshPosition();
			}
		}
		
		function showAdjust(){
			$('#innerAdjust').parent().parent().parent().show();
			$('#inner').parent().parent().parent().show();
			
			$('#innerAdjust').select2();
			$('#inner').select2();
			$('.select2').addClass('full-width');
		}
		
		function showMoreInfos(infos){
			infos.forEach(function(e){
				$("#"+e).parent().parent().show();
			});
//			$('#innerAdjust').parent().parent().parent().show();
//			$('#inner').parent().parent().parent().show();
		}
		
		function closeAddWindow() {
			$('.addBusizInfo').hide();
			$('.busizInfo').show();
			$('#a_bInfo')[0].reset();
			$('#bPrice').change();
		}
		
		function showLevelDialog(t) {
			type = t;
			$('#upBusiz').modal('show');
		}
		
		
		function saveBusiz() { //新增
			var regu = /^[0-9]+$/;
			if($('#Busiz_id').val()==='' || !regu.test($('#Busiz_id').val())){
				toastr.warning('请输入正整数形式的业务单元编号');
				return ;
			}
			
			if($('#Busiz_desc').val()===''){
				toastr.warning('请输入业务单元描述');
				return ;
			}
			
			//输入校验
			if($('#Ftp_flag').val()==='0'){ //定价
				if($('#ftp_method_id').val()===null){
					toastr.warning("请选择定价方法");
					return;
				}else {
					var v=$('#ftp_method_id').val();
					if(v==='103' || v==='102' || v==='105'){ //现金流加权期限法、直接期限法、久期法
						if($('#curve_id').val()===null){
							toastr.warning("请选择曲线");
							return;
						}
					}else if(v==='101'){ //指定期限法
						if($('#curve_id').val()===null || $('#term_cd').val()==='' || $('#term_cd_mult').val()===null){
							toastr.warning("请完善曲线、期限、期限单位");
							return;
						}else if($('#term_cd').val().indexOf('.') >=0 || $('#term_cd').val().indexOf('-') >=0 ){
							toastr.warning("请输入正整数的期限值");
							return;
						}else{
							var flag=0;
							if($('#term_cd_mult').val()==='Y'){
								if($('#term_cd').val()>30){
									flag=1;
								}
							}else if($('#term_cd_mult').val()==='M'){
								if($('#term_cd').val()>30*12){
									flag=1;
								}
							}else if($('#term_cd_mult').val()==='D'){
								if($('#term_cd').val()>30*365){
									flag=1;
								}
							}
							
							if(flag==1){
								toastr.warning("请输入30年内的期限值");
								return;
							}
						}
					}else if(v==='104'){ //沉淀率法
						if($('#curve_id').val()===null){
							toastr.warning("请选择曲线");
							return;
						}
					}else if(v==='106'){ //利率差额法
						if($('#point_val').val()===''){
							toastr.warning("请输入点差值");
							return;
						}else {
							
							if(!inte.test($('#point_val').val())){
								toastr.warning("请输入整数的点差值");
								return;
							}else if($('#point_val').val()>9999 || $('#point_val').val()<-9999){
								toastr.warning("点差值输入介于-9999至9999之间");
								return;
							}
						}
					}
				}
			}
			
			var adjs=$('#inner').val(), json=[];
			if(adjs!=null){
				adjs.forEach(function(e){
					var v=e.split(' ');
					json.push({AdjId:v[0], AdjTypeId:v[1]});
				});
			}
			
			$.post('/platform/FtpBusiz', $('#a_bInfo').serialize() + '&JSON=' + JSON.stringify(json) + '&r=' + Math.random(100000000000000), function(data) {
				var rs = JSON.parse(data);
				if(rs.ErrorCode === '1') {
					toastr.success(rs.ErrorMsg);
					refreshTree($('#Busiz_id').val());
					$('#a_bInfo')[0].reset();
					closeAddWindow();
				} else {
					toastr.error(rs.ErrorMsg);
					return false;
				}
			});
		}
		
		function refreshTree(selectNode) {
			$.get('/platform/FtpBusizStruct?r=' + Math.random(100000000000000), function(data) {
				$.fn.zTree.init($("#treeDemo"), s, JSON.parse(data)); //树
			});

			$.get('/platform/FtpBusiz?r=' + Math.random(100000000000000), function(data) {
				$.fn.zTree.init($("#busizTree"), setting, JSON.parse(data)); //树	

				var treeObj = $.fn.zTree.getZTreeObj("busizTree");
				//				treeObj.expandAll(true);
				var node = treeObj.getNodeByParam("id", selectNode, null);
				treeObj.selectNode(node);
				
				zTreeOnClick(event,'busizTree',node);
				$('#a_bInfo')[0].reset();
			});
			
			$('#innerAdjust').select2();
			$('#inner').select2();
			$('.select2').addClass('full-width');
		}
		
		function updateBusiz() { //更新
			if($("#bCode").val() === '' || $("#bCode").val() === null) {
				toastr.warning('请选择业务单元');
				return;
			}
			//更新输入校验
			if($('#bPrice').val()==='0'){ //定价 根据不同方法的输入校验
				if($('#bMethod').val()===null){
					toastr.warning('请选择定价方法');
					return;
				}else {
					var v=$('#bMethod').val();
					if(v==='103' || v==='102' || v==='105'){ //现金流加权期限法、直接期限法、久期法
						if($('#bCurve').val()===null){
							toastr.warning('请选择曲线');
							return;
						}
					}else if(v==='101'){ //指定期限法
						if($('#bCurve').val()===null || $('#bTerm').val()==='' || $('#bTermUnit').val()===null){
							toastr.warning('请完善曲线、期限、期限单位');
							return;
						}else if($('#bTerm').val().indexOf('.') >=0 || $('#bTerm').val().indexOf('-') >=0 ){
							toastr.warning('请输入正整数的期限值');
							return;
						}else{
							var flag=0;
							if($('#bTermUnit').val()==='Y'){
								if($('#bTerm').val()>30){
									flag=1;
								}
							}else if($('#bTermUnit').val()==='M'){
								if($('#bTerm').val()>30*12){
									flag=1;
								}
							}else if($('#bTermUnit').val()==='D'){
								if($('#bTerm').val()>30*365){
									flag=1;
								}
							}
							
							if(flag==1){
								toastr.warning('请输入30年内的期限值');
								return;
							}
						}
					}else if(v==='104'){ //沉淀率法
						if($('#bCurve').val()===null){
							toastr.warning('请选择曲线');
							return;
						}
					}else if(v==='106'){ //利率差额法
						if($('#bSpot').val()===''){
							toastr.warning('请输入点差值');
							return;
						}else {
							if(!inte.test($('#bSpot').val())){
								toastr.warning('请输入整数的点差值');
								return;
							}else if($('#bSpot').val()>9999 || $('#bSpot').val()<-9999){
								toastr.warning('点差值输入介于-9999至9999之间');
								return;
							}
						}
					}
				}
			}
			
			var adjs=$('#innerAdjust').val(), json=[];
			if(adjs!=null){
				adjs.forEach(function(e){
					var v=e.split(' ');
					json.push({AdjId:v[0], AdjTypeId:v[1]});
				});
			}
			
			$.ajax({
				type: "put",
				url: "/platform/FtpBusiz?" + $('#bInfo').serialize() + '&JSON=' + JSON.stringify(json) + '&r=' + Math.random(100000000000000),
				error: function(msg) {
					console.log(msg.responseText);
				},
				success: function(data) {
					var rs = JSON.parse(data);
					if(rs.ErrorCode === '1') {
						toastr.success(rs.ErrorMsg);
						refreshTree($('#bCode').val());
					} else {
						toastr.error(rs.ErrorMsg);
						return false;
					}
				}
			});

			return false;
		}
		
		function delBusiz(){
			if($("#bCode").val() === '' || $("#bCode").val() === null) {
				toastr.warning("请选择业务单元");
				return;
			}
			
			if($('#bPCode').val()==='-1'){
				toastr.warning('不能直接删除根节点');
				return ;
			}
			
			swal({
				title: "删除！",
				text: "是否删除该业务单元及其子节点配置信息？",
				type: "warning",
				showCancelButton: true,
				confirmButtonColor: "#DD6B55",
				//			        confirmButtonText: "",
				closeOnConfirm: false
			}, function() {
				$.ajax({
					type: "delete",
					url: "/platform/FtpBusiz?Busiz_id=" + $("#bCode").val() + '&r=' + Math.random(100000000000000),
					error: function(msg) {
						console.log(msg.responseText);
					},
					success: function(data) {
						var rs = JSON.parse(data);
						if(rs.ErrorCode === '1') {
//							toastr.success(rs.ErrorMsg);
							swal("删除!", "业务单元已删除", "success");
							var treeObj = $.fn.zTree.getZTreeObj("busizTree");
							var node = treeObj.getNodeByParam("id", $('#bCode').val(), null);
	
							$('#bInfo')[0].reset();
							refreshTree(node.pId);
						} else {
							toastr.error(rs.ErrorMsg);
							return;
						}
					}
				});
			});
		}
		
		function searchBusiz() {
			var treeObj = $.fn.zTree.getZTreeObj("busizTree");
			var nodes = treeObj.getSelectedNodes(); //取消已选中的
			if (nodes.length>0) { 
				treeObj.cancelSelectedNode(nodes[0]);
			}
			var node = treeObj.getNodeByParam("id", $('#searchBusiz').val(), null);
			treeObj.selectNode(node);
			
			node=treeObj.getSelectedNodes();
			if(node.length==0){
				toastr.warning("未搜索到该业务单元");
				return ;
			}
			zTreeOnClick(event,'busizTree',node[0]);
		}
		
		function selectTermStructs(type){ //期限结构选择
			restoreBusiz();
			$('.termCheckBase input:checkbox').removeAttr('checked');
			
			var busizId,domainId;
			if(type==='1'){ //更新
				busizId=$('#bCode').val();
				if(busizId==='' || busizId===null){
					toastr.warning("请选择一个业务单元");
					return ;
				}
//				domainId=$('#bRegion').val();
				$('#bCodeS').val(busizId);
				$('#domainIdS').val(domainId);
				
			}else if(type==='2'){ //新增
				busizId=$('#Busiz_id').val();
				domainId=$('#Domain_id_B').val();
				if(busizId === '' || domainId === ''){
					toastr.warning("请先完善业务单元编码和域等基本信息");
					return ;
				}
				
				$('#bCodeS').val(busizId);
				$('#domainIdS').val(domainId);
			}
			
			createTermCheckBox($('#bCodeS').val(),$('#domainIdS').val());
			
			$('#termStruct').modal('show');
		}
		
		function restoreBusiz(){ //将控件是否显示还原
	        $('.s1').show();
	        $('.s2').hide();
	        $('.steps button').hide();
	        $('.next').show();
	    }
		
		function goStepBusiz(step){
			$('.steps button').hide();
			$('.s1').hide();
			$('.s2').hide();
	        if(step==='1'){ //prev
	        		$('.s1').show();
	        		$('.next').show();
	        }else if(step==='2'){ //next
	        		//至少选择勾选一个
	        		if($(".termCheckBase .checked").length===0){
					toastr.warning("请至少勾选一个期限");
					restoreBusiz();
					return ;
	        		}
	        		
	        		$('#termStructTable tbody').html('');
	        		$(".termCheckBase .checked input").each(function(){
					var temp="<tr id='term_'"+$(this).attr('name')+"><td class='termS' TermCd='"+$(this).attr('Term_cd')+"' TermCdMult='"+$(this).attr('Term_cd_mult')+"'>"+$(this).attr('name')
	            				+"</td><td ondblclick='editTermValue(this)'><span></span><input type='text' style='display: none;width:100px;' onfocusout='sureTermVaule(this)'/></td>";
	            		$('#termStructTable tbody').append(temp);
	        		});
	        		
	        		var busizId=$('#bCodeS').val();
				var domainId=$('#domainIdS').val();
	        		//填入已有值
	        		$.get('/platform/FtpRedemption?r='+Math.random(10000000000000)+"&BusizId="+busizId+"&DomainId="+domainId,function(data){ //获取改业务单元的期限结构信息
					var rs=JSON.parse(data);
					if(rs!=null){
						rs.forEach(function(e){
							var s=e.TermCd+e.TermCdMult;
							$('.termS').each(function(){
								if($(this).text()===s){
									$(this).next().find('span').text(e.Weight);
									return ;
								}
							});
						});
					}
				});
	        		
	        		$('.s2').show();
	        		$('.prev').show();
            		$('.save').show();
	        }
	    }
		
		function createTermCheckBox(busizId,domainId){
			$.ajax({
				type: "get",
				url: '/platform/FtpCurveSave?r='+Math.random()*10000000000000,
				async: false,
				error: function(msg) {
					console.log(msg.responseText);
				},
				success: function(data) {
					var rs=JSON.parse(data);
		            $('.termCheckBase').html('');
		            for(var i=0; i<rs.length; i++) {
		                r = rs[i].Struct_code;
						var ap="<label class='checkbox-inline i-checks'> <input type='checkbox' Term_cd='"+rs[i].Term_cd+"' Term_cd_mult='"+rs[i].Term_cd_mult+"' name='"+r+"' value='" + r + "'>"+r+"</label>"
//						<label class='checkbox-inline i-checks'> <input type='checkbox' value='option2'> b </label>
		                $('.termCheckBase').append(ap);
		            }
		            
		            iCheckInit();
		            $.ajax({
						type: "get",
						url: '/platform/FtpRedemption?r='+Math.random(10000000000000)+"&BusizId="+busizId+"&DomainId="+domainId,
						async: false,
						error: function(msg) {
							console.log(msg.responseText);
						},
						success: function(data) {
							var rs=JSON.parse(data);
							//选中已有期限值
							if(rs!=null){
								rs.forEach(function(e){
									var s=e.TermCd+e.TermCdMult;
									$(".termCheckBase input[name="+s+"]").parent().addClass('checked');
//									$(".termCheckBase input[name="+s+"]").attr('checked', 'checked');
								});
							}
						}
					});
				}
			});
		}
		
		function editTermValue(obj){
			$(obj).find('input').show();
			$(obj).find('input')[0].value=$(obj).find('span')[0].innerText;
			$(obj).find('input')[0].focus();
			$(obj).find('span').hide();
		}
		
		function sureTermVaule(obj){
			var v=$(obj).val();
			var re = /^[0,1]+\.?[0-9]*$/;
			if(v.indexOf('-')>=0 || !re.test(v)){ //数字、非负
				toastr.warning('请输入0-1的小数值');
				return;
			}else{
				if(parseFloat(v) > 1 || parseFloat(v) <= 0) { //0-1
					toastr.warning('请输入0-1的小数值');
					return;
				}
			}
			
			$(obj).prev()[0].innerText=$(obj).val();
			$(obj).hide();
			$(obj).prev().show();
		}
		
		function saveTermStruct(){ //保存业务单元的期限结构
			var termCd,termCdMult,weight;
			var terms=[],flag=0;
			var allWeight=0.0;
			$('.termS').each(function(){
				termCd=$(this).attr('TermCd')
				termCdMult=$(this).attr('TermCdMult');
				
				weight=$(this).next().find('span')[0].innerText;
				if(weight==''){
					flag=1;
				}
				allWeight+=parseFloat(weight)*100000000;
				terms.push({"TermCd":termCd,"TermCdMult":termCdMult,"Weight":weight});
			});
			
			if(flag==1){
				toastr.warning('请给每个期限输入比例值');
				return;
			}
			if(allWeight/100000000!=1){
				toastr.warning('总的期限比例值不等于1，请检查输入');
				return;
			}
			
			var busizId=$('#bCodeS').val();
			var domainId=$('#domainIdS').val();
			
			$.ajax({
	            type:"put",
	            url:'/platform/FtpRedemption?r='+Math.random(100000000000)+"&BusizId="+busizId+"&DomainId="+domainId+"&JSON="+JSON.stringify(terms), 
	            error: function(msg){
	                console.log(msg.responseText);
	            },
	            success: function(data){
					var rs = JSON.parse(data);
					if(rs.ErrorCode === '1') {
						toastr.success(rs.ErrorMsg);
						$('#termStruct').modal('hide');
					} else {
						toastr.error(rs.ErrorMsg);
						return false;
					}
	            }
	        });
		}
	</script>

</html>