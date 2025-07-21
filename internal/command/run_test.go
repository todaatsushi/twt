package command_test

import (
	"errors"
	"testing"

	"github.com/todaatsushi/twt/internal/command"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		name            string
		input           string
		expectedSuccess bool
		expectedError   error
	}{
		{
			name:            "Valid branch",
			input:           "branch/task",
			expectedSuccess: true,
			expectedError:   nil,
		},
		{
			name:            "Has spaces",
			input:           "cat branch/task",
			expectedSuccess: false,
			expectedError:   errors.New("Error: illegal character(s): \" \""),
		},
		{
			name:            "Has colon",
			input:           "branch/task;branch/task",
			expectedSuccess: false,
			expectedError:   errors.New("Error: illegal character(s): \";\""),
		},
		{
			name:            "Has colon",
			input:           "branch/task\nbranch/task",
			expectedSuccess: false,
			expectedError:   errors.New("Error: illegal character(s): \"\\n\""),
		},
	}

	var branch string
	var err error
	for _, c := range cases {
		branch, err = command.Validate(c.input)
		if c.expectedSuccess && err != nil {
			t.Fatalf("%s: Expected success but got error: %s (%s)", c.name, err, c.input)
		} else {
			if c.expectedSuccess {
				if branch != c.input {
					t.Fatalf("%s: Expected %s but got %s", c.name, c.input, branch)
				}
			} else {
				if err == nil {
					t.Fatalf("%s: Expected error but got success (%s)", c.name, c.input)
				}
				if err.Error() != c.expectedError.Error() {
					t.Fatalf("%s: Expected error %s but got %s (%s)", c.name, c.expectedError, err, c.input)
				}
			}
		}
	}
}

func TestTerminalRunner(t *testing.T) {
	t.Run("Catch stdout", func(t *testing.T) {
		runner := command.Terminal{}

		out, err := runner.Run("echo", "Hello, World!")
		if len(err) > 0 {
			t.Fatalf("Expected no error but got: %v", err)
		}
		if len(out) == 0 || out[0] != "Hello, World!" {
			t.Fatalf("Expected output 'Hello, World!' but got: %v", out)
		}
	})

	t.Run("Catch stderr", func(t *testing.T) {
		runner := command.Terminal{}

		out, err := runner.Run("sh", "-c", "echo 'Hello, World!' >&2")
		if len(out) > 0 {
			t.Fatalf("Expected no stdout output but got: %v", out)
		}
		if len(err) == 0 || err[0] != "Hello, World!" {
			t.Fatalf("Expected stderr output 'Hello, World!' but got: %v", err)
		}
	})
}
