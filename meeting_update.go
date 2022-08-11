package zoom

type UpdateMeetingParams struct {
	MeetingID    int    `url:"-"`
	OccurrenceID string `url:"occurrence_id,omitempty"`
}

type UpdateMeetingOptions struct {
	Agenda         string          `json:"agenda,omitempty"`
	Duration       int             `json:"duration,omitempty"`
	Password       string          `json:"password,omitempty"`
	Recurrence     Recurrence      `json:"recurrence"`
	Settings       MeetingSettings `json:"settings,omitempty"`
	StartTime      *Time           `json:"start_time,omitempty"`
	Timezone       string          `json:"timezone,omitempty"`
	Topic          string          `json:"topic,omitempty"`
	TrackingFields []TrackingField `json:"tracking_fields,omitempty"`
	Type           MeetingType     `json:"type,omitempty"`
}

// UpdateMeetingPath - v2 update a meeting
const UpdateMeetingPath = "/meetings/%d"

func UpdateMeeting(params UpdateMeetingParams, opts UpdateMeetingOptions) error {
	return defaultClient.UpdateMeeting(params, opts)
}

func (c *Client) UpdateMeeting(params UpdateMeetingParams, opts UpdateMeetingOptions) error {
	return c.requestV2(requestV2Opts{
		Method:         Patch,
		Path:           UpdateMeetingPath,
		DataParameters: &opts,
		URLParameters:  &params,
		HeadResponse:   true,
	})
}
