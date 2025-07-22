package checks

import (
	"testing"
)

type ErrRunner struct {
	errs []string
}

func (r ErrRunner) Run(app string, args ...string) (out, err []string) {
	return nil, r.errs
}

type OutRunner struct {
	res []string
}

func (r OutRunner) Run(app string, args ...string) (out, err []string) {
	return r.res, nil
}

func TestInGitDir(t *testing.T) {
	t.Run("InGitDir returning true", func(t *testing.T) {
		expected := true
		runner := OutRunner{
			res: []string{"true"},
		}

		actual := InGitDir(runner)

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("InGitDir returning false - no out", func(t *testing.T) {
		expected := false
		runner := OutRunner{
			res: []string{},
		}

		actual := InGitDir(runner)

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("InGitDir returning false - invalid output", func(t *testing.T) {
		expected := false
		runner := OutRunner{
			res: []string{"not true or false"},
		}

		actual := InGitDir(runner)

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("InGitDir returning false - false", func(t *testing.T) {
		expected := false
		runner := OutRunner{
			res: []string{"false"},
		}

		actual := InGitDir(runner)

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})
}

func TestIsInWorktree(t *testing.T) {
	t.Run("IsInWorktree returning true", func(t *testing.T) {
		expected := true
		runner := OutRunner{
			res: []string{"true"},
		}

		actual := IsInWorktree(runner)

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("IsInWorktree returning false - no out", func(t *testing.T) {
		expected := false
		runner := OutRunner{
			res: []string{},
		}

		actual := IsInWorktree(runner)

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("IsInWorktree returning false - invalid output", func(t *testing.T) {
		expected := false
		runner := OutRunner{
			res: []string{"not true or false"},
		}

		actual := IsInWorktree(runner)

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("IsInWorktree returning false - false", func(t *testing.T) {
		expected := false
		runner := OutRunner{
			res: []string{"false"},
		}

		actual := IsInWorktree(runner)

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})
}

func TestIsUsingBareRepo(t *testing.T) {
	t.Run("IsUsingBareRepo - true", func(t *testing.T) {
		expected := true
		runner := OutRunner{
			res: []string{"/path/to/repo.git                        (bare)"},
		}

		actual := IsUsingBareRepo(runner)

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	t.Run("IsUsingBareRepo - false", func(t *testing.T) {
		expected := false
		runner := OutRunner{
			res: []string{},
		}

		actual := IsUsingBareRepo(runner)

		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})
}

func TestAssertGit(t *testing.T) {
	t.Run("AssertGit - valid", func(t *testing.T) {
		runner := OutRunner{
			res: []string{"true"},
		}

		err := AssertGit(runner)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("AssertGit - invalid", func(t *testing.T) {
		runner := OutRunner{
			res: []string{"false"},
		}

		err := AssertGit(runner)
		if err == nil {
			t.Error("Expected an error, got nil")
		}

		expected := "\u2717 Git status invalid - must be in a .git folder (worktree base) or inside a worktree, and in a bare repository."
		actual := err.Error()

		if actual != expected {
			t.Errorf("Expected error message '%s', got '%s'", expected, actual)
		}
	})
}

func TestAssertReady(t *testing.T) {
	t.Run("AssertReady - all checks pass", func(t *testing.T) {
		runner := OutRunner{
			res: []string{"true"},
		}

		if AssertReady(runner) {
			t.Error("Expected AssertReady to return false, but it returned true")
		}
	})

	t.Run("AssertReady - any check fails", func(t *testing.T) {
		runner := ErrRunner{
			errs: []string{"Any error occurred"},
		}

		if !AssertReady(runner) {
			t.Error("Expected AssertReady to return true, but it returned false")
		}
	})
}
