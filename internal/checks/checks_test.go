package checks

import "testing"

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
