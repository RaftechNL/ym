package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
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
			ctx := context.Background()

			// Read input files
			inputFiles := c.StringSlice("input")
			if len(inputFiles) <= 0 {
				return cli.ShowAppHelp(c)
			}
			mergedData, err := MergeYAML(ctx, inputFiles...)
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

func MergeYAML(ctx context.Context, filenames ...string) ([]byte, error) {
	if len(filenames) <= 0 {
		return nil, errors.New("You must provide at least one filename for merging YAML files")
	}

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)

	var resultValues sync.Map

	for _, filename := range filenames {
		// Check if the context has been cancelled
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		var override map[string]interface{}
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Info(err)
			continue
		}
		if err := yaml.Unmarshal(bs, &override); err != nil {
			log.Info(err)
			continue
		}

		// Merge override map with resultValues map
		resultValues.Range(func(key, value interface{}) bool {
			overrideKey := key.(string)
			overrideValue := value
			if _, ok := override[overrideKey]; !ok {
				override[overrideKey] = overrideValue
			}
			return true
		})

		// Store merged map in resultValues
		resultValues.Store(filename, override)

		// Encode override map to buffer
		if err := encoder.Encode(override); err != nil {
			log.Info(err)
			continue
		}
	}

	// Check if the context has been cancelled
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if err := encoder.Close(); err != nil {
		log.Info(err)
		return nil, err
	}

	var mergedMap map[string]interface{}
	resultValues.Range(func(key, value interface{}) bool {
		override := value.(map[string]interface{})
		if mergedMap == nil {
			mergedMap = override
		} else {
			for k, v := range override {
				mergedMap[k] = v
			}
		}
		return true
	})

	bs, err := yaml.Marshal(mergedMap)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	// Save result to file
	if err := ioutil.WriteFile("result.yaml", bs, 0644); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return bs, nil
}
