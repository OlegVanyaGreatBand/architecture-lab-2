package main

import (
	"flag"
	lab2 "github.com/OlegVanyaGreatBand/architecture-lab-2"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile = flag.String("f", "", "Expression to compute")
	outputFile = flag.String("o", "", "Expression to compute")
)

func main() {
	flag.Parse()

	var source io.Reader
	var dest io.Writer

	if *inputExpression != "" {
		source = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		data, err := ioutil.ReadFile(*inputFile)
		if err != nil {
			_, _ = os.Stderr.WriteString(err.Error())
			return
		}
		source = strings.NewReader(string(data))
	} else {
		_, _ = os.Stderr.WriteString("No expression provided")
		return
	}

	if *outputFile != "" {
		if file, err := os.Create(*outputFile); err == nil {
			dest = file
		} else {
			_, _ = os.Stderr.WriteString("Error with output file")
			return
		}
	} else {
		dest = os.Stdout
	}

	handler := lab2.ComputeHandler{
		In:  source,
		Out: dest,
	}

	if err := handler.Compute(); err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return
	}
}
