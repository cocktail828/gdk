package httpd

//go:generate stringer -type state -linecomment
type state int

const (
	initing state = iota // initing
	running              // running
)
