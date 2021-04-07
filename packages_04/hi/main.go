package main

import (
	"github.com/headfirstgo/greeting"
	"github.com/headfirstgo/greeting/dansk"
	"github.com/headfirstgo/greeting/deutsch"
)

// Package naming conventions.
// A package name should be all lowercase.
// The name should be abbreviated if the meaning is ovbious.
// One word if possible.

func main() {
	greeting.Hi()
	greeting.Hello()
	dansk.GodMorgen()
	deutsch.GutenTag()
}
