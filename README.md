# twt
## Tmux / Git worktree workflow manager

`twt` is an interface for managing git worktrees in different tmux sessions.

Built for MacOS.

# Quickstart
1. Clone a bare repository to work in
```
git clone --bare <repo> <dir>
```

Move the executable from `build` into your PATH, or add the path to the executable to
your PATH.

2. Check ability to run command
```
twt check
```

3. To checkout any branch in a new session
```
twt go -b <branch>
```

3. To remove a session and / or worktree
```
twt rm -b <branch>
```

# Usage
`twt` has four commands:
 - `go`
 - `rm`
 - `common`
 - `check`

## `go`

Change to a session for a branch:

 - If the session exists, switch to it. Otherwise create one from the branch name.
 - If the worktree exists, check it out, otherwise create one.
 - If the branch exists, check it out, otherwise create one.

Must be run from within a bare repo or worktree, and within a tmux session.

## `rm`

Cleanup a worktree/session by removing both, if they exist. Options to delete and force
delete the branch/worktree.

Must be run from within a bare repo or worktree, and within a tmux session.

## Common files

In case your project has assets to be shared across branches (e.g. `.env` vars, docker
compose processes, etc.), they can be stored in a common files directory which lives in
the base dir (ie. the bare repo).

Common files are also useful for optionally running setup scripts when a worktree and
session starts. Those can be written and placed there for the command to automatically
pick it up and run.

### `common`
ie. `twt common`

Creates / switches to a session in the common files directory.

### `init`
ie. `twt common init`

In the case where a common files directory doesn't exist in the bare repo, create one
with the directories needed to run the setup scripts, along with an initial template.

## `check`

Check the viability of using `twt` features:

 - Is this run in a bare repo or in a worktree
 - Is this run in a tmux session
 - Are common files set up

## Usage with other tools

You can use regular shell tools like completions & fzf to streamline the development.
Check out `zsh-examples/bindings` for a sample script, which can speed up workflow.

# Why
## Git worktrees

Git worktrees are a way of managing branches in a cloned repository, by creating a bare
repository and spawining working directories within it for each branch.

It avoids the issue of being limited to one state of the repository at any one time, and
allows working on multiple branches at once.

* [Docs](https://git-scm.com/docs/git-worktree#:~:text=A%20git%20repository%20can%20support,others%20in%20the%20same%20repository.)

## tmux

Tmux is a session multiplexer, allowing multiple windows and sessions to be run in a
single terminal session.

* [GitHub](https://github.com/tmux/tmux/wiki)
