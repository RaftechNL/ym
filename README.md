# ym

![Logo](https://img.raftech.nl/white_logo_color1_background.png)

The yaml-merge CLI allows you to merge multiple YAML files into a single file.



#
[![License](https://img.shields.io/github/license/raftechnl/cli-yaml-merger)](./LICENSE)


## Functionality

The yaml-merge CLI reads one or more YAML files and merges them into a single YAML file. The merged YAML file contains all the key-value pairs from the input files. If there are any conflicts between the keys, the value from the last input file takes precedence.

## Installing

### Download
> Check our release page to download a specific version

```shell
    #!/bin/bash

    # Fetch the latest release version from Github API
    VERSION=$(curl --silent "https://api.github.com/repos/RaftechNL/ym/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

    # Set the URL of the tarball for the latest release
    URL="https://github.com/RaftechNL/cli-yaml-merger/releases/download/${VERSION}/ym_${VERSION}_darwin_x86_64.tar.gz"

    # Download and install the latest release
    curl -L ${URL} | tar xz
    chmod +x ym
    sudo mv ym /usr/local/bin/
```

### Homebrew
```shell
brew tap  RaftechNL/toolbox
brew install raftechnl/toolbox/ym
```

## Example

To merge two YAML files (file1.yaml and file2.yaml) and save the result to a new file (merged.yaml), run:

```shell
ym -i file1.yaml -i file2.yaml -o merged.yaml
```

This command reads the input YAML files (file1.yaml and file2.yaml), merges them into a single YAML file, and saves the result to a new file (merged.yaml).

If you want to merge three or more YAML files, simply add more -i options to the command:

```shell
ym -i file1.yaml -i file2.yaml -i file3.yaml -o merged.yaml
```

This command reads the input YAML files (file1.yaml, file2.yaml, and file3.yaml), merges them into a single YAML file, and saves the result to a new file (merged.yaml).


## Contributing

Contributions are always welcome!


## Authors

- [@rafpe](https://www.github.com/rafpe)
