package calcLog

type Calclog interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Error(v ...interface{})
	Warn(v ...interface{})
	Close()
}
