package main

import (
	"flag"
	"log"
	"os"
	"regexp"
)

type Config struct {
	SkipHeaders bool
	Seperator   string

	fset *flag.FlagSet
}

func (c *Config) Register(f *flag.FlagSet) *Config {
	c.fset = f

	flag.BoolVar(&c.SkipHeaders, "no-header", false, "Skip Setting Headers")
	flag.StringVar(&c.Seperator, "s", `[\t]+`, "Regexp of Delimiter to Build Table on")

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
	if err := c.Parse(os.Args); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
