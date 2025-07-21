package git

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

func TestWorktrees(t *testing.T) {
	t.Run("RemoveWorktree err bubbles up to result", func(t *testing.T) {
		expectedErrMsg := "fatal: 'name' is not a working tree"
		expected := []string{expectedErrMsg}

		runner := ErrRunner{
			errs: expected,
		}

		actual := RemoveWorktree(runner, "name", "branch", false, false)

		if len(expected) != len(actual) {
			t.Errorf("Expected %d errors, got %d", len(expected), len(actual))
		}

		for i, err := range actual {
			if err != expected[i] {
				t.Errorf("Expected error %d to be '%s', got '%s'", i, expected[i], err)
			}
		}
	})

	t.Run("RemoveWorktree no errors handles nil returned", func(t *testing.T) {
		runner := OutRunner{
			res: nil,
		}

		actual := RemoveWorktree(runner, "name", "branch", false, false)

		if actual != nil {
			t.Errorf("Expected no errors, got %v", actual)
		}
	})

	t.Run("RemoveWorktree no errors handles empty slice", func(t *testing.T) {
		runner := OutRunner{
			res: []string{},
		}

		actual := RemoveWorktree(runner, "name", "branch", false, false)

		if len(actual) != 0 {
			t.Errorf("Expected 0 errors, got %d", len(actual))
		}
	})
}
