package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/fatih/color"
	flags "github.com/jessevdk/go-flags"
	"github.com/mattn/go-colorable"
)

var (
	colors = map[string]color.Attribute{
		"black":   color.FgBlack,
		"white":   color.FgWhite,
		"red":     color.FgRed,
		"green":   color.FgGreen,
		"blue":    color.FgBlue,
		"magenta": color.FgMagenta,
		"cyan":    color.FgCyan,
		"yellow":  color.FgYellow,
	}
)

type Options struct {
	Color     string `short:"c" long:"color" description:"output color" choise:"black white red green blue magenta cyan yellow" default:"white"`
	Bold      bool   `short:"b" long:"bold" description:"make output with bold"`
	Under     bool   `short:"u" long:"under" description:"make output with underline"`
	Separater string `short:"s" long:"sep" description:"separate character if multiple args given" default:" "`
	End       string `short:"e" logn:"end" description:"EOF character" default:"\n"`

	List bool `short:"l" long:"list" description:"show avairable color list"`
}

var (
	opt    = new(Options)
	parser = flags.NewParser(opt, flags.Default)
)

func main() {
	args, err := parser.Parse()
	if err != nil {
		os.Exit(0)
	}

	out := colorable.NewColorableStdout()
	outColor := color.New(colors[opt.Color])
	defer color.Unset()

	// list flag
	if opt.List {
		fmt.Println("Avairable colors:")
		colorList := make([]string, len(colors))
		i := 0
		for k, _ := range colors {
			colorList[i] = k
			i++
		}
		sort.Strings(colorList)
		for _, c := range colorList {
			color.Set(colors[c])
			fmt.Println("  ", c)
			color.Unset()
		}
		os.Exit(0)
	}

	// create color object
	if opt.Bold {
		outColor.Add(color.Bold)
	}
	if opt.Under {
		outColor.Add(color.Underline)
	}
	outColor.Set()

	// write args to stdout by color
	for i, a := range args {
		if i > 0 {
			fmt.Fprint(out, opt.Separater)
		}
		fmt.Fprint(out, a)
	}
	fmt.Fprint(out, opt.End)
}
