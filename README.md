# barnacle

An ohmyzsh plugin that automatically puts the npm modules for the current directory on your path.

### Pre-Requisites

* ohmyzsh (https://ohmyz.sh/)
* Go (https://go.dev/)
* GoReleaser (https://goreleaser.com/install/)

## Installation

Create a symlink from ohmyzsh's custom plugins to this repository.

```
ln -s $ZSH_CUSTOM/plugins/barnacle ~/settings/barnacle
```

Load the plugin in your `~/.zshrc`.

```
# where ... are other plugins
plugins=(... barnacle ...)
``` 

Source your `~/.zshrc`

```
source ~/.zshrc
```
