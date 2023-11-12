# BetterCommit (bc)

![better-commit](assets/banner.png)

This project started to make yourself write better commit messages but later got extended to make your git journey much easier.

# Table of contents

- [BetterCommit (bc)](#bettercommit-bc)
- [Table of contents](#table-of-contents)
    - [Usage](#usage)
    - [Commands](#commands)
        - [`bc commit`](#bc-commit)
        - [`bc push`](#bc-push)
        - [`bc branch`](#bc-branch)
        - [`bc pr`](#bc-pr)
    - [Steps to generate GH PAT( Personal Access Token):*](#steps-to-generate-gh-pat-personal-access-token)
    - [Setup](#setup)
    - [Help](#help)
  - [Features](#features)
    - [Contributing guidelines](#contributing-guidelines)
    - [Disclaimer ;)](#disclaimer-)
    - [See it live in action](#see-it-live-in-action)

### Usage

```
% bc help

BetterCommit (bc) is a small utility which will help you make an habit of writing better commit messages.

Usage: bc <subcommands>

Available commands:
    add         Adds all the changes to staging area
    commit      Adds and commits all the changes
    branch      Creates a new branch
    pr          Creates a new pull request

Example usage:

    $ bc add

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

bc allows you to create queryable git branches.

```
% bc branch

Is this commit related to any projects tickets / Components / features (eg: JIRA-124, button, vpc): certs
What kind of work is this (eg: bugfix, feat, demo, test): feat
What work will be done in this branch: Add support for application to run TLS natively
Switched to a new branch 'feat/certs/Add-support-for-application-to-run-TLS-natively'
```

At the end you are left with this clean sexy looking git friendly branch

`feat/certs/Add-support-for-application-to-run-TLS-natively`

which you can query ( suppose if i want to see how many people are working on features ) now i could just do

```
git branch --list "feat/*"
  feat/cert/create-a-new-self-signed-certificate
* feat/certs/Add-support-for-application-to-run-TLS-natively
```

##### `bc pr`

bc allows you to create pr from terminal. bc looks for env var called `GH_TOKEN` when creating a PR, so please generate one before you go about creating a new PR. Steps to generate PAT shown below.

```
% bc pr

Title of the Pull Request: Add new command to create PR
Is this PR associated with any ticket (eg: JIRA-124]): bc-124
Explain work done in this PR (When finished hit ctrl-d on a new line to proceed):
Add new command to create PR from terminal
Updated Makefile to substitute variable based on OS Type
PR type (eg: SHOW, SHIP. ASK): SHIP
What kind of change is this (eg: Bufix, Feature, Breaking Change, Doc update): Feature
Source branch name: feat/bc/Add-pr-from-terminal-command
Destination branch name: main
```

When you are done, `bc` will check if the branch has been pushed to origin or not, and if its pushed, then it will create a new PR.

*NOTE: To use this command, you would need to generate GH PAT ( Personal Access Token)*

### Steps to generate GH PAT( Personal Access Token):*

1. Go to 'Profile > Settting'
2. Click on 'Developer settings'
3. Click 'Personal acces tokens' > 'Fine grained tokens'
4. Click 'Generate new token'
5. Given 'Token name', 'Expiration'
6. In Repository access, select whether you want to use bc to create PR in all respsitory. If that is the case then select 'All repositories'. If not then, select individual repositories
7. Once 'Repository access' is selected, add permission via 'Permissions > Repository permissions'
8. Provide access to create PR by clicking drop down on 'Pull requests > Access: Read and Write' 
9. Click 'Generate token'
10. Token will be shown only once hence save it someplace safe.
11. Add token export to your rc file.

If you are on mac or using zsh then do this
```shell
echo 'export GH_TOKEN="github_pat_1....REDACTED"' >> ~/.zshrc
```

If you are on linux or using bash then do this
```shell
echo 'export GH_TOKEN="github_pat_1....REDACTED"' >> ~/.bash_rc
```

### Setup

1. Clone repository `git clone https://github.com/pgaijin66/bc.git`

2. Navigate to the cloned directory.

3. Run `make install` to install


### Help

```
% make help

help                           List targets & descriptions
install                        Installs binary to standard library PATH4
tag-n-release                  Tags and pushes
uninstall                      Uninstalls application
```


## Features

- [X] Prettify and format commit message.
- [X] Stop committting and pushing changes to main branch.
- [X] Create querable branches.
- [X] Make compatible with `bash` as well. Currently works on `zsh` 
- [X] Create GH PR from terminal.
- [ ] Allow use to interactively choose which files to commit.
- [ ] interactively cherry pick commits from one branch to another.
- [ ] Make reabasing easy

### Contributing guidelines

Please create a PR for contribution to this project.

### Disclaimer

This is a tool i made for myself, as time and again i would get lazy and and skip the most important part while programming ( write good commit messages ). There is no doubt that this work can be done with a simple function alias to `~/.zshrc` or `~/.bashrc` as well but why make it easy when you can add multiple layers of encapsulation. After all that is what automation is right. :P 

### See it live in action

Commits: https://github.com/pgaijin66/bc/commits/main

Branches: https://github.com/pgaijin66/bc/branches