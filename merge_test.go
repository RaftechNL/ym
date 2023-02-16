// file: main_test.go
package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeYAML_TwoFiles(t *testing.T) {
	result, err := MergeYAML("tests/test1.yaml", "tests/test2.yaml")
	assert.NoError(t, err)
	assert.Equal(t, strings.TrimSpace(`
modules:
    my_module:
        providerAliasRef: my_provider2
        source: github.com/example/module
        version: v1.6.0
    my_module2:
        providerAliasRef: my_provider2
        source: github.com/example/module
        version: v1.6.0
    my_module3:
        providerAliasRef: my_provider2
        source: github.com/example/module
        version: v1.6.0
providers:
    my_provider:
        auth:
            ssh_key: ssh:key:2312312
        providerType: github
`), strings.TrimSpace(string(result)))
}

func TestMergeYAML_MissingFile(t *testing.T) {
	result, err := MergeYAML("tests/test1.yaml", "tests/test3.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such file or directory")
	assert.Equal(t, "", string(result))
}

func TestMergeYAML_InvalidYAML(t *testing.T) {
	result, err := MergeYAML("tests/test1.yaml", "tests/invalid.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal data from file")
	assert.Equal(t, "", string(result))
}
