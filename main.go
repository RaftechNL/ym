package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

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

func mergeMaps(dst, src map[string]interface{}) {
	for k, v := range src {
		if _, ok := dst[k]; !ok {
			// key does not exist in dst, just copy from src
			dst[k] = v
			continue
		}

		// key exists in both dst and src, need to merge recursively
		dstValue := dst[k]
		switch dstValue := dstValue.(type) {
		case map[string]interface{}:
			srcValue, ok := v.(map[string]interface{})
			if !ok {
				// type mismatch, just copy from src
				dst[k] = v
				continue
			}
			mergeMaps(dstValue, srcValue)
		default:
			// type mismatch or dst has scalar value, just copy from src
			dst[k] = v
			continue
		}
	}
}

func MergeYAML(filenames ...string) ([]byte, error) {
	if len(filenames) <= 0 {
		return nil, errors.New("You must provide at least one filename for reading Values")
	}

	resultValues := make(map[string]interface{})
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
		mergeMaps(resultValues, override)
	}

	bs, err := yaml.Marshal(resultValues)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return bs, nil
}
