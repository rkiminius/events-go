package events

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

var videoQualities = []string{
	"720p",
	"1080p",
	"1440p",
}

var audioQualities = []string{
	"Low",
	"Medium",
	"High",
}

// validateEventInput This method is used to validate the input data for an event,
// represented by the Event struct, to ensure that it meets the required criteria.
func validateEventInput(event *Event) error {
	var maxInvitees = 100

	if event.Name == "" {
		return errors.New("missing event name")
	}

	if event.Date.IsZero() {
		return errors.New("missing event date")
	}

	if time.Now().After(event.Date) {
		return errors.New("incorrect event date: event can be arranged only for a future")
	}

	if len(event.Languages) <= 0 {
		return errors.New("missing event languages")
	}

	if len(event.VideoQuality) <= 0 {
		return errors.New("missing event videoQuality")
	}

	for _, video := range event.VideoQuality {
		if valid := stringExistsInSlice(video, videoQualities); !valid {
			return errors.New(fmt.Sprintf("'%s' video quality format is not supported", video))
		}
	}

	if len(event.AudioQuality) <= 0 {
		return errors.New("missing event audioQuality")
	}

	for _, audio := range event.AudioQuality {
		if valid := stringExistsInSlice(audio, audioQualities); !valid {
			return errors.New(fmt.Sprintf("'%s' audio quality format is not supported", audio))
		}
	}

	if len(event.Invitees) <= 0 {
		return errors.New("missing event invitees")
	}

	for _, email := range event.Invitees {
		valid := IsEmailValid(email)
		if !valid {
			return errors.New(fmt.Sprintf("inviter email '%s' in not valid", email))
		}
	}

	if len(event.Invitees) > maxInvitees {
		return errors.New(fmt.Sprintf("more than %d invitees are not allowed", maxInvitees))
	}

	return nil
}

// IsEmailValid checks if the given email address is valid using regex
func IsEmailValid(email string) bool {
	// Regular expression for email validation
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

// stringExistsInSlice checks if a given string exists in a slice of strings.
func stringExistsInSlice(s string, slice []string) bool {
	for _, elem := range slice {
		if elem == s {
			return true
		}
	}
	return false
}
