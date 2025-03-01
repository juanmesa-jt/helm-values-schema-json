package pkg

import (
	"bytes"
	"flag"
	"fmt"
)

// Parse flags
func ParseFlags(progname string, args []string) (config *Config, output string, err error) {
	flags := flag.NewFlagSet(progname, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var conf Config
	flags.Var(&conf.input, "input", "Multiple yaml files as inputs (comma-separated)")
	flags.StringVar(&conf.outputPath, "output", "values.schema.json", "Output file path")
	flags.IntVar(&conf.draft, "draft", 2020, "Draft version (4, 6, 7, 2019, or 2020)")
	flags.IntVar(&conf.indent, "indent", 4, "Indentation spaces (even number)")

	// Nested SchemaRoot flags
	flags.StringVar(&conf.SchemaRoot.ID, "schemaRoot.id", "", "JSON schema ID")
	flags.StringVar(&conf.SchemaRoot.Title, "schemaRoot.title", "", "JSON schema title")
	flags.StringVar(&conf.SchemaRoot.Description, "schemaRoot.description", "", "JSON schema description")
	flags.Var(&conf.SchemaRoot.AdditionalProperties, "schemaRoot.additionalProperties", "JSON schema additional properties (true/false)")

	err = flags.Parse(args)
	if err != nil {
		fmt.Println("Usage: helm schema [-input STR] [-draft INT] [-output STR]")
		return nil, buf.String(), err
	}

	conf.args = flags.Args()
	return &conf, buf.String(), nil
}
