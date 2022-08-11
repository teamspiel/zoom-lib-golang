package zoom

import "errors"

var (
	// ErrMeetingRegistrantFirstNameRequired is returned when the first_name is not provided
	ErrMeetingRegistrantFirstNameRequired = errors.New("first_name is required")
	ErrMeetingRegistrantEmailRequired     = errors.New("email is required")
)
