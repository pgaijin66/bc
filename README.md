# BetterCommit (bc)

![better-commit](assets/banner.png)

This project started to make yourself write better commit messages but later got extended to make your git journey much easier.

### Usage

```
% bc help

BetterCommit (bc) is a small utility which will help you make an habit of writing better commit messages.

Usage: bc <subcommands>

Available commands:
    add         Adds all the changes to staging area
    commit      Adds and commits all the changes
    branch      Creates a new branch

Example usage:

    bc add


```

### Commands

##### `bc commit`
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

##### `bc push`

bc does not allow to push in main branch.
```
% bc push

ERROR: Not recommended to push to main branch. Please create a new branch using "bc branch"
```

bc prompts you if try to push and still have some pending changes to be committed.

```
% bc push
There are still changes to be committed. Are you sure you want to push? [y/N]

```

##### `bc branch`
```
% bc branch

Is this commit related to any projects tickets / Components / features (eg: JIRA-124, button, vpc): bc
What kind of work is this (eg: bugfix, feat, demo, test): feat
What work will be done in this branch: update readme
Switched to a new branch 'feat/bc/update-readme'
```

### Setup

1. Clone repository `git clone https://github.com/pgaijin66/bc.git`

2. Run `make build`

3. Reload terminal `sournce ~/.zshrc` if you are using `zsh` shell.


### Setup help
```
% make help

build                          Installs binary to standard library PATH
help                           List targets & descriptions
uninstall                      Uninstalls application
```


## Features

- [X] Prettify and format commit message.
- [X] Stop committting and pushing changes to main branch.
- [X] Create querable branches.
- [X] Make compatible with `bash` as well. Currently works on `zsh` 
- [ ] Allow use to interactively choose which files to commit.
- [ ] interactively cherry pick commits from one branch to another.
- [ ] Make reabasing easy

### Contributing guidelines

Please create a PR for contribution to this project.

### Disclaimer ;)

```
This is a tool i made for myself, as time and again i would get lazy and and skip the most important part while programming ( write good commit messages ). There is no doubt that this work can be done with a simple function alias to `~/.zshrc` or `~/.bashrc` as well but why make it easy when you can add multiple layers of encapsulation ;). After all that is what automation is right. :P 
```


### See it live in action

Commits: https://github.com/pgaijin66/bc/commits/main

Branches: https://github.com/pgaijin66/bc/branches