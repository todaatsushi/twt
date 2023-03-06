# twt
## Tmux / Git worktree workflow manager

`twt` is an interface for managing git worktrees in different tmux sessions.

# Quickstart
1. Clone a bare repository to work in
```
git clone --bare <repo> <dir>
```

2. Check ability to run command
```
twt check
```

3. To checkout any branch in a new session
```
twt go <branch>
```

3. To remove a session and / or worktree
```
twt rm <branch>
```

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
