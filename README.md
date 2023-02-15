# cli-yaml-merger

![Logo](https://img.raftech.nl/white_logo_color1_background.png)

The yaml-merge CLI allows you to merge multiple YAML files into a single file.



#
[![License](https://img.shields.io/github/license/raftechnl/cli-yaml-merger)](./LICENSE)


## Functionality

The yaml-merge CLI reads one or more YAML files and merges them into a single YAML file. The merged YAML file contains all the key-value pairs from the input files. If there are any conflicts between the keys, the value from the last input file takes precedence.

    
## Example

To merge two YAML files (file1.yaml and file2.yaml) and save the result to a new file (merged.yaml), run:

```shell
yaml-merger -i file1.yaml -i file2.yaml -o merged.yaml
```

This command reads the input YAML files (file1.yaml and file2.yaml), merges them into a single YAML file, and saves the result to a new file (merged.yaml).

If you want to merge three or more YAML files, simply add more -i options to the command:

```shell
yaml-merge -i file1.yaml -i file2.yaml -i file3.yaml -o merged.yaml
```

This command reads the input YAML files (file1.yaml, file2.yaml, and file3.yaml), merges them into a single YAML file, and saves the result to a new file (merged.yaml).


## Contributing

Contributions are always welcome!


## Authors

- [@rafpe](https://www.github.com/rafpe)
