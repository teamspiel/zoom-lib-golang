package zoom

import "fmt"

type (
	MeetingRegistrantCreateParams struct {
		MeetingID    int    `url:"-"`
		OccurrenceID string `url:"occurrence_id,omitempty"`
	}

	MeetingRegistrantCreateInput struct {
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Email       string `json:"email,omitempty"`
		Address     string `json:"address,omitempty"`
		City        string `json:"city,omitempty"`
		State       string `json:"state,omitempty"`
		Zip         string `json:"zip,omitempty"`
		Country     string `json:"country,omitempty"`
		Phone       string `json:"phone,omitempty"`
		Comments    string `json:"comments,omitempty"`
		Industry    string `json:"industry,omitempty"`
		JobTitle    string `json:"job_title,omitempty"`
		Org         string `json:"org,omitempty"`
		AutoApprove bool   `json:"auto_approve,omitempty"`
	}

	MeetingRegistrantCreateResponse struct {
		ID           int          `json:"id,omitempty"`
		JoinURL      string       `json:"join_url,omitempty"`
		RegistrantId string       `json:"registrant_id,omitempty"`
		StartTime    Time         `json:"start_time,omitempty"`
		Topic        string       `json:"topic,omitempty"`
		Occurrences  []Occurrence `json:"occurrences"`
	}
)

func (m *MeetingRegistrantCreateInput) Validate() error {
	if m.FirstName == "" {
		return ErrMeetingRegistrantFirstNameRequired
	}

	if m.Email == "" {
		return ErrMeetingRegistrantEmailRequired
	}

	return nil
}

// MeetingRegistrantCreatePath - Add a meeting registrant
const MeetingRegistrantCreatePath = "/meetings/%d/registrants"

func MeetingRegistrantCreate(params MeetingRegistrantCreateParams, input MeetingRegistrantCreateInput) (MeetingRegistrantCreateResponse, error) {
	return defaultClient.MeetingRegistrantCreate(params, input)
}

func (c *Client) MeetingRegistrantCreate(params MeetingRegistrantCreateParams, input MeetingRegistrantCreateInput) (MeetingRegistrantCreateResponse, error) {
	var response MeetingRegistrantCreateResponse

	return response, c.requestV2(requestV2Opts{
		Method:         Post,
		Path:           fmt.Sprintf(MeetingRegistrantCreatePath, params.MeetingID),
		DataParameters: &input,
		URLParameters:  &params,
		HeadResponse:   true,
		Ret:            &response,
	})
}
