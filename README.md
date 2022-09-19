# barnacle

An ohmyzsh plugin that automatically puts the npm modules for the current directory on your path.

### Pre-Requisites

* ohmyzsh (https://ohmyz.sh/)
* Go (https://go.dev/)

## Installation

Ensure your `go` binary path is in your $PATH, e.g.

```
PATH=$HOME/go/bin:$PATH
...
```

Install `barnacle` binary.

```
go install
```

Create a symlink from ohmyzsh's custom plugins to this repository.

```
ln -s ~/settings/barnacle $ZSH_CUSTOM/plugins/barnacle
```

Load the plugin in your `~/.zshrc`.

```
# where ... are other plugins
plugins=(... barnacle ...)
``` 

Source your `~/.zshrc`.

```
source ~/.zshrc
```
