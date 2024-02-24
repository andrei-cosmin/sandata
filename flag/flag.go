package flag

type Flag struct {
	bool
}

func New() Flag {
	return Flag{false}
}

func (l *Flag) Set() {
	l.bool = true
}

func (l *Flag) Clear() {
	l.bool = false
}

func (l *Flag) IsSet() bool {
	return l.bool
}

func (l *Flag) IsClear() bool {
	return !l.bool
}
