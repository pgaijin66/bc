# BetterCommit (bc)

![better-commit](assets/banner.png)

Make yourself write better commit messages.

### Before 
```
commit b81c906fc00eb28421bcbf3e1579b9f3cbf6cdd6
Author: Prabesh Thapa <sthapaprabesh2020@gmail.com>
Date:   Tue Nov 1 18:04:06 2022 -0700

    removing unwanted listings.go file
```

### After

```
commit 1b1ef83XXXXXXXXXXX3614
Author: Prabesh Thapa <sthapaprabesh2020@gmail.com>
Date:   Wed Nov 2 16:04:28 2022 -0700

    add(commit message): add validation and formatting of the commit message and added usage as well

```

## Setup
```
% make build
```
## Usage

```
BetterCommit (bc) is a small utility which will help you make an habit of writing better commit messages.

Usage: bc <subcommands>

Available commands:
    add         Adds all the changes to staging area
    commit      Adds and commits all the changes.

Example usage:

    bc add

```


## Help
```
% make help

build                          Installs binary to standard library PATH
help                           List targets & descriptions
uninstall                      Uninstalls application
```


## Features

- [X] Prettify and format commit message
- [ ] Allow use to interactively choose which files to commit.
- [ ] Make compatible with `bash` as well. Currently works on `zsh` 
- [ ] interactively cherry pick commits from one branch to another.
- [ ] Make reabasing easy

## Contributing guidelines

Please create a PR for contribution to this project.

## Disclaimer ;)

```
This is a tool i made for myself, as time and again i would get lazy and and skip the most important part while programming. There is no doubt that this work can be done with a simple function alias to `~/.zshrc` or `~/.bashrc` as well.
```