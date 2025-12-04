package main

import "log"

/*
This example is a basic way to select a demo example to be executed when running
the program (without the usage of the package applet, just to show that is base
simple). This simple pattern can be used in most of cases (when developping an
algorithm and testing differents use cases, without needing an application
framework).

The hypothesis is the same than for applet, i.e. a use case function is a
function with no argument and that return an error.

*/

var tests map[string]func() error = map[string]func() error{
	"T01": test01,
	"T02": test02,
	"T03": test03,
}

func main() {
	name := "T02"
	test := tests[name]
	log.Printf("executing test %s", name)
	if err := test(); err != nil {
		log.Fatal(err)
	}
}
