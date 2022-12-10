package factory

import (
	"fmt"
	"strings"
)

type (
	// Activity  is simple factory
	Activity interface{ Say(name string) }
	// SayHi simple factory implement by Activity
	SayHi struct{}

	// SayHello simple factory implement by Activity
	SayHello struct{}

	// SayBye simple factory implement by Activity
	SayBye struct{}
)

// NewActivity is Activity's constructor
func NewActivity(name string) Activity {
	switch strings.ToLower(name) {
	case `hi`:
		return &SayHi{}
	case `hello`:
		return &SayHello{}
	case `bye`:
		return &SayBye{}
	default:
		return nil
	}
}

func (s *SayHi) Say(name string)    { fmt.Printf("Hi, %s\n", name) }
func (s *SayHello) Say(name string) { fmt.Printf("Hello, %s\n", name) }
func (s *SayBye) Say(name string)   { fmt.Printf("Bye, %s\n", name) }
