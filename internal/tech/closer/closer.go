package closer

import "github.com/rs/zerolog/log"

type ErrorLogger interface {
	ErrorF(format string, a ...interface{})
}

type Effector func() error

var globalCloser = New()

func Add(f ...Effector) {
	globalCloser.Add(f...)
}

func CloseAll() {
	globalCloser.CloseAll()
}

type Closer struct {
	lg          ErrorLogger
	closerFuncs []Effector
}

func New() *Closer {
	return &Closer{
		closerFuncs: make([]Effector, 0),
	}
}

func (c *Closer) SetErrorLogger(lg ErrorLogger) {
	c.lg = lg
}

func (c *Closer) Add(f ...Effector) {
	c.closerFuncs = append(c.closerFuncs, f...)
}

func (c *Closer) CloseAll() {
	for _, f := range c.closerFuncs {
		if err := f(); err != nil {
			if c.lg != nil {
				log.Printf("Error on close: %v", err)
			}
		}
	}

	c.closerFuncs = make([]Effector, 0)
}
