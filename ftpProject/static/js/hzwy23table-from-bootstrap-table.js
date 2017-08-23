/**
 * Created by hzwy23 on 2016/7/3.
 */

(function ($) {
    $.fn.extend({
        hzwTable:function(options){

            var defaults = {

                method: 'get',

                striped: true,

                pagination: true,

                showRefresh:true,

                showToggle:true,

                clickToSelect:true,

                pageSize: 20,

                showExport:true,

                showToggle:true,

                showPaginationSwitch:true,

                height: $(window).height() - 58,

                pageList: [10, 20, 30, 50, 100, 200,800],

                search: true,

                sidePagination: "server",

                showColumns: true,

                minimunCountColumns: 2,

            };

            $.extend(true,defaults,options);

            var $table = $(this).bootstrapTable({

                    method: defaults.method,

                    url: defaults.url,

                    striped: defaults.striped,

                    pagination: defaults.pagination,

                    showRefresh:defaults.showRefresh,

                    showToggle:defaults.showToggle,

                    clickToSelect:defaults.clickToSelect,

                    toolbar:defaults.toolbar,

                    pageSize: defaults.pageSize,

                    uniqueId:defaults.uniqueId,

                    showExport:defaults.showExport,

                    showToggle:defaults.showToggle,

                    showPaginationSwitch:defaults.showPaginationSwitch,

                    height: defaults.height,

                    pageList: defaults.pageList,

                    search: defaults.search,

                    sidePagination: defaults.sidePagination,

                    showColumns: defaults.showColumns,

                    minimunCountColumns: defaults.minimunCountColumns,

                    columns:defaults.columns
            });
            $('.columns').find("button").css("height","34px");
        },
        hzwInsertRow:function(e){
            if (e.handle != undefined){
                e.handle();
            }
            this.bootstrapTable('insertRow', {
                index: e.index,
                row:e.row
            });
        },

        hzwDeleteRow:function(e){


            var fieldName = e.filed;
            var ids = $.map(this.bootstrapTable('getSelections'), function (row) {
                return row[fieldName];
            });

            if (e.handle != undefined){
                var flag = e.handle(ids);
                if (flag){
                    this.bootstrapTable('remove', {
                        field: fieldName,
                        values: ids
                    });
                }
            }
        },

        hzwEditRow:function(e){
            
            var fieldName = e.filed;
            var ids = $.map(this.bootstrapTable('getSelections'), function (row) {
                return row[fieldName];
            });

            if (e.handle != undefined){
                e.handle(ids);
            }

            this.bootstrapTable('updateByUniqueId', {
                id: ids,
                row: e.row
            });
        }
    });
})(jQuery);