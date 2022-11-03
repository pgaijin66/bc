# BetterCommit (bc)

![better-commit](assets/banner.png)

Make yourself write better commit messages.

### Usage

```
BetterCommit (bc) is a small utility which will help you make an habit of writing better commit messages.

Usage: bc <subcommands>

Available commands:
    add         Adds all the changes to staging area
    commit      Adds and commits all the changes.

Example usage:

    bc add

```

### See it live in action

https://github.com/pgaijin66/bc/commits/main

### How it works

```
% bc commit
Is this commit related to any projects tickets / Components / features (eg: JIRA-124, button, vpc): bc
Enter which git operation did you performed (eg: add, update, del): update
INFO: Files modified
 M src/bc
What did you update: updated script to show modified files to help user add more context to the commit
[main be3e6a8] update(bc): updated script to show modified files to help user add more context to the commit
 1 file changed, 2 insertions(+), 2 deletions(-)
```

### Setup

1. Clone repository `git clone https://github.com/pgaijin66/bc.git`

2. Run `make build`

3. Reload terminal `sournce ~/.zshrc`



### Help
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

### Contributing guidelines

Please create a PR for contribution to this project.

### Disclaimer ;)

```
This is a tool i made for myself, as time and again i would get lazy and and skip the most important part while programming ( write good commit messages ). There is no doubt that this work can be done with a simple function alias to `~/.zshrc` or `~/.bashrc` as well but why make it easy when you can add multiple layers of encapsulation ;). After all that is what automation is right. :P 
```