package logger

type Logger interface {
	Log(message string)
}

type Adapter func(message string)

func (a Adapter) Log(message string) {
	a(message)
}
