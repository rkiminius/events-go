package events

import (
	"testing"
	"time"
)

func TestValidateEventInput(t *testing.T) {
	incorrectDataTestCases := []struct {
		name          string
		event         Event
		expectedError string
	}{
		{
			name: "EventNameNotDefined",
			event: Event{
				Name:         "",
				Date:         time.Now().Add(24 * time.Hour),
				Languages:    []string{"English"},
				VideoQuality: []string{"1080p"},
				AudioQuality: []string{"High"},
				Invitees:     []string{"example1@gmail.com"},
			},
			expectedError: "missing event name",
		},

		{
			name: "VideoQualityFormatNotSupported",
			event: Event{
				Name:         "My event",
				Date:         time.Now().Add(24 * time.Hour),
				Languages:    []string{"English"},
				VideoQuality: []string{"4k"},
				AudioQuality: []string{"High"},
				Invitees:     []string{"example1@gmail.com"},
			},
			expectedError: "'4k' video quality format is not supported",
		},

		{
			name: "VideoQualityFormatNotSupported",
			event: Event{
				Name:         "My event",
				Date:         time.Now().Add(-24 * time.Hour),
				Languages:    []string{"English"},
				VideoQuality: []string{"1080p"},
				AudioQuality: []string{"High"},
				Invitees:     []string{"example1@gmail.com"},
			},
			expectedError: "incorrect event date: event can be arranged only for a future",
		},
	}

	for _, tc := range incorrectDataTestCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validateEventInput(&tc.event)
			if result == nil {
				t.Error("Expected error")
			} else if result.Error() != tc.expectedError {
				t.Errorf("%s: Expected error to be '%s', but got '%s'", tc.name, tc.expectedError, result.Error())
			}
		})
	}
}

func TestIsEmailValid(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "ValidEmail",
			email:    "example@example.com",
			expected: true,
		},
		{
			name:     "InvalidEmailWithoutAtSymbol",
			email:    "example.com",
			expected: false,
		},
		{
			name:     "InvalidEmailWithoutDomain",
			email:    "example@.com",
			expected: false,
		},
		{
			name:     "InvalidEmailWithoutTopLevelDomain",
			email:    "example@example",
			expected: false,
		},
		{
			name:     "InvalidEmailWithSpecialCharacters",
			email:    "example@#$%.com",
			expected: false,
		},
		{
			name:     "InvalidEmailWithMultipleAtSymbols",
			email:    "example@@example.com",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsEmailValid(tc.email)
			if result != tc.expected {
				t.Errorf("Expected IsEmailValid(%s) to be %v, but got %v", tc.email, tc.expected, result)
			}
		})
	}
}
