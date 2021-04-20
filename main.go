package main

import (
	"flag"
	"log"
	"os"
	"regexp"

	"github.com/keyneston/mktable/table"
)

type Config struct {
	SkipHeaders bool
	Seperator   string
	MaxPadding  int

	fset *flag.FlagSet
}

func (c *Config) Register(f *flag.FlagSet) *Config {
	c.fset = f
	c.fset.BoolVar(&c.SkipHeaders, "no-header", false, "Skip Setting Headers")
	c.fset.StringVar(&c.Seperator, "s", `[ \t]*\t[ \t]*`, "Regexp of Delimiter to Build Table on")
	c.fset.IntVar(&c.MaxPadding, "max-padding", -1, "Maximum units of padding. Set to a negative number for unlimited")

	return c
}

func (c *Config) Parse(args []string) error {
	return c.fset.Parse(args)
}

func (c *Config) CompileSeperator() (*regexp.Regexp, error) {
	return regexp.Compile(c.Seperator)
}

func main() {
	fset := flag.NewFlagSet("", flag.ExitOnError)
	c := (&Config{}).Register(fset)
	if err := c.Parse(os.Args[1:]); err != nil {
		log.Fatalf("Error: %v", err)
	}

	tb := table.NewTable(regexp.MustCompile(c.Seperator))
	tb.SetMaxPadding(c.MaxPadding)
	tb.Read(os.Stdin)
	tb.Write(os.Stdout)
}
