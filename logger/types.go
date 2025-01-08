package logger

type Message struct {
	Text string
	Type MessageType
}

type MessageType int

const (
	Err MessageType = iota
	Warning
	Info
)
