package tmux

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

func TestGetCurrentSessionName(t *testing.T) {
	t.Run("GetCurrentSessionName - get session", func(t *testing.T) {
		runner := OutRunner{
			res: []string{"test-session"},
		}

		name, err := GetCurrentSessionName(runner)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if name != "test-session" {
			t.Errorf("Expected session name 'test-session', got '%s'", name)
		}
	})
	t.Run("GetCurrentSessionName - no session", func(t *testing.T) {
		runner := OutRunner{
			res: []string{},
		}

		_, err := GetCurrentSessionName(runner)
		if err == nil {
			t.Error("Expected an error, got nil")
		}

		expected := "Couldn't fetch current tmux session name"
		actual := err.Error()

		if expected != actual {
			t.Errorf("Expected error '%s', got '%s'", expected, actual)
		}
	})
}

func TestListSessions(t *testing.T) {
	t.Run("ListSessions - get sessions", func(t *testing.T) {
		expected := []string{"session1", "session2"}
		runner := OutRunner{
			res: expected,
		}

		// justNames doesn't really matter
		sessions, err := ListSessions(runner, true)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(sessions) != len(expected) {
			t.Errorf("Expected %d sessions, got %d", len(expected), len(sessions))
		}

		for i, session := range sessions {
			if session != expected[i] {
				t.Errorf("Expected session '%s', got '%s'", expected[i], session)
			}
		}
	})

	t.Run("ListSessions - no sessions", func(t *testing.T) {
		runner := OutRunner{
			res: []string{},
		}

		_, err := ListSessions(runner, true)
		if err == nil {
			t.Error("Expected an error, got nil")
		}

		expected := "Couldn't fetch current tmux session name"
		actual := err.Error()

		if expected != actual {
			t.Errorf("Expected error '%s', got '%s'", expected, actual)
		}
	})
}

func TestHasSession(t *testing.T) {
	t.Run("HasSession - session exists", func(t *testing.T) {
		runner := OutRunner{
			res: []string{},
		}

		exists := HasSession(runner, "test-session")
		if !exists {
			t.Error("Expected session to exist, got false")
		}
	})

	t.Run("HasSession - session does not exist", func(t *testing.T) {
		runner := ErrRunner{
			errs: []string{"no such session"},
		}

		exists := HasSession(runner, "nonexistent-session")
		if exists {
			t.Error("Expected session to not exist, got true")
		}
	})
}
