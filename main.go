package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/keyneston/mktable/table"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

type Config struct {
	Format      FormatValue
	MaxPadding  int
	Reformat    bool
	Seperator   string
	SkipHeaders bool
	Version     bool

	Alignments ParseAlignments

	fset *flag.FlagSet
}

func (c *Config) Register(f *flag.FlagSet) *Config {
	allFormats := strings.Join(table.AllFormats(), ",")

	c.fset = f
	c.fset.BoolVar(&c.SkipHeaders, "no-header", false, "Skip Setting Headers")
	c.fset.StringVar(&c.Seperator, "s", `[ \t]*\t[ \t]*`, "Regexp of Delimiter to Build Table on")
	c.fset.IntVar(&c.MaxPadding, "max-padding", -1, "Maximum units of padding. Set to a negative number for unlimited")
	c.fset.BoolVar(&c.Reformat, "r", false, "Read in markdown table and reformat")
	c.fset.BoolVar(&c.Reformat, "reformat", false, "Alias for -r")
	c.fset.Var(&c.Format, "f", fmt.Sprintf("Set the format. Available formats: %v", allFormats))
	c.fset.Var(&c.Format, "format", "Alias for -f")
	c.fset.Var(&c.Alignments, "a", "Set column alignments; Can be called multiple times and/or comma separated. Arrow indicates direction '<' left, '>' right, '=' center; Columns are zero indexed; e.g. -a '0<,1>,2='")
	c.fset.BoolVar(&c.Version, "v", false, "Print version info")
	c.fset.BoolVar(&c.Version, "version", false, "Alias for -v")

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

	if c.Version {
		PrintVersion(os.Args[0])
		os.Exit(0)
	}

	tableConfig := table.TableConfig{
		MaxPadding:  c.MaxPadding,
		SkipHeaders: c.SkipHeaders,
		Alignments:  c.Alignments.alignments,
		Seperator:   regexp.MustCompile(c.Seperator),
		Format:      c.Format.GetFormat(),
	}
	if c.Reformat {
		tableConfig.Format = table.FormatMK
	}

	tb := table.NewTable(tableConfig)
	tb.Read(os.Stdin)
	tb.Write(os.Stdout)
}

func PrintVersion(bin string) {
	fmt.Printf(strings.TrimLeft(`
%s
version: %s
commit: %v
built-on: %v
built-by: %v
formats: %v
`,
		" \t\n"),
		filepath.Base(bin),
		version,
		commit,
		date,
		builtBy,
		strings.Join(table.AllFormats(), ", "),
	)
}
