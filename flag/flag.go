/*
 * MIT License
 *
 * Copyright (c) 2025 Andrei Casu-Pop
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
 * documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
 * Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
 * WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
 * OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package flag

// Flag boolean flag
type Flag struct {
	bool
}

// New creates a new flag
func New() Flag {
	return Flag{false}
}

// Set sets the flag to true
func (l *Flag) Set() {
	l.bool = true
}

// Clear clears the flag to false
func (l *Flag) Clear() {
	l.bool = false
}

// IsSet returns true if the flag is set
func (l *Flag) IsSet() bool {
	return l.bool
}

// IsCleared returns true if the flag is cleared
func (l *Flag) IsCleared() bool {
	return !l.bool
}
