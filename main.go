package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var version = "v0.0.1"

func main() {
	app := &cli.App{
		Name:    "yaml-merge",
		Usage:   "Merge multiple YAML files",
		Version: version,
		Authors: []*cli.Author{
			{
				Name:  "NinjaOps by raftech.io",
				Email: "hello@raftech.nl",
			},
		},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "Input YAML files to merge",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output file path for merged YAML",
			},
		},
		Action: func(c *cli.Context) error {

			// Read input files
			inputFiles := c.StringSlice("input")
			if len(inputFiles) <= 0 {
				return cli.ShowAppHelp(c)
			}
			mergedData, err := MergeYAML(inputFiles...)
			if err != nil {
				return err
			}

			// Write merged data to output file
			outputFile := c.String("output")
			if outputFile != "" {
				if err := ioutil.WriteFile(outputFile, mergedData, 0644); err != nil {
					return err
				}
			} else {
				// Print merged data to console
				fmt.Println(string(mergedData))
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func removeLastLineIfEmpty(s string) string {
	lines := strings.Split(s, "\n")
	lastLine := strings.TrimSpace(lines[len(lines)-1])
	if lastLine == "" {
		lines = lines[:len(lines)-1]
	}
	return strings.Join(lines, "\n")
}

func MergeYAML(filenames ...string) ([]byte, error) {
	if len(filenames) <= 0 {
		return nil, errors.New("You must provide at least one filename for reading Values")
	}
	var resultValues map[string]interface{}
	for _, filename := range filenames {

		var override map[string]interface{}
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Info(err)
			return nil, fmt.Errorf("failed to read file %q: %w", filename, err)
		}
		if err := yaml.Unmarshal(bs, &override); err != nil {
			log.Info(err)
			return nil, fmt.Errorf("failed to unmarshal data from file %q: %w", filename, err)
		}

		//check if is nil. This will only happen for the first filename
		if resultValues == nil {
			resultValues = override
		} else {
			for k, v := range override {
				// Check if value is nil before adding to resultValues
				if v != nil {
					resultValues[k] = v
				}
			}
		}

	}

	bs, err := yaml.Marshal(resultValues)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	if err != nil {
		log.Info(err)
		return nil, err
	}

	return bs, nil
}
