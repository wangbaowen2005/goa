// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Code Generator
//
// Command:
// $ goa

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"goa.design/goa/codegen/generator"
	"goa.design/goa/eval"
	_ "goa.design/goa/examples/calc/design"
	"goa.design/goa/pkg"
)

func main() {
	var (
		out     = flag.String("output", "", "")
		version = flag.String("version", "", "")
		cmdl    = flag.String("cmd", "", "")
	)
	{
		flag.Parse()
		if *out == "" {
			fail("missing output flag")
		}
		if *version == "" {
			fail("missing version flag")
		}
		if *cmdl == "" {
			fail("missing cmd flag")
		}
	}

	if *version != pkg.Version() {
		fail("cannot run generator produced by goa version %s and compiled with goa version %s\n", *version, pkg.Version())
	}
	if err := eval.Context.Errors; err != nil {
		fail(err.Error())
	}
	if err := eval.RunDSL(); err != nil {
		fail(err.Error())
	}
	if err := os.RemoveAll("/home/nitinmohan/rightscale/go/src/goa.design/goa/examples/calc/gen"); err != nil {
		fail(err.Error())
	}

	outputs, err := generator.Generate(*out, "gen")
	if err != nil {
		fail(err.Error())
	}

	fmt.Println(strings.Join(outputs, "\n"))
}

func fail(msg string, vals ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, vals...)
	os.Exit(1)
}
