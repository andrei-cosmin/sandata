package flag

// Flag - struct for a boolean flag
type Flag struct {
	bool
}

// New - creates a new flag
func New() Flag {
	return Flag{false}
}

// Set - sets the flag to true
func (l *Flag) Set() {
	l.bool = true
}

// Clear - clears the flag to false
func (l *Flag) Clear() {
	l.bool = false
}

// IsSet - returns true if the flag is set
func (l *Flag) IsSet() bool {
	return l.bool
}

// IsClear - returns true if the flag is clear
func (l *Flag) IsClear() bool {
	return !l.bool
}
