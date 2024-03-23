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
