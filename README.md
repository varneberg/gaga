# Gaga
## Motivation
Gaga is a CLI-tool for adding labeling within github actions. Authentication is done within the workflow step, making gaga being authenticated with the same permissions as the Workflow it is called from.

Currently supports adding labels to pull requests

## Installation

Get the binary for operating system from releases

```bash
curl -fL -o gaga.tar.gz https://github.com/varneberg/gaga/releases/download/{gaga_version}/gaga_{os}_{architecture}.tar.gz
```

Extract the binary

```bash
tar -C /usr/bin -xzf ./gaga.tar.gz 
```

## Examples

Adding a label to a pull request:

```bash
gaga label -n <label_to_add>
```

## Github Workflow Implementation

```yaml
name: Gaga
permissions:
  contents: write
  issues: write
  pull-request: write

jobs:
  gaga:
    steps:
    - name: Setup gaga
      run: |
        sudo curl -fL -o gaga.tar.gz https://github.com/varneberg/gaga/releases/download/v0.0.1/gaga_linux_amd64.tar.gz
        sudo tar -C /usr/bin -xzf ./gaga.tar.gz
      
    - name: Gaga add labels
      run: gaga label -n newLabel
      env:
        GITHUB_TOKEN: ${{ github.token }}

```
