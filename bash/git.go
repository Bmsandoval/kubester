package bash

// Git
func GitFetchRecurse() string { return "git fetch --recurse-submodules" }
func GitPullRecurse() string { return "git pull --recurse-submodules" }
func GitStash() string { return "git stash" }
func GitStashRecurse() string { return "git submodule foreach 'git stash'" }
func GitResetHard() string { return "git reset --hard" }
func GitResetHardRecurse() string { return "git submodule foreach 'git reset --hard'" }

