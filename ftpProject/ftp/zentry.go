package ftp

type GoEntry struct {
	RouteControl
}

func (this *GoEntry) Get() {
	this.TplName = "theme/default/ftpindex.tpl"
}
