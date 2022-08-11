package zoom

import "fmt"

type UpdateMeetingLivestreamParams struct {
	MeetingID int `url:"-"`
}

type UpdateMeetingLivestreamParamsOptions struct {
	PageUrl   string `json:"page_url"`
	StreamKey string `json:"stream_key"`
	StreamUrl string `json:"stream_url"`
}

// UpdateMeetingLivestreamPath - v2 update a meeting livestream
const UpdateMeetingLivestreamPath = "/meetings/%d/livestream"

func (c *Client) UpdateMeetingLivestream(params UpdateMeetingLivestreamParams, opts UpdateMeetingLivestreamParamsOptions) error {
	return c.requestV2(requestV2Opts{
		Method:         Patch,
		Path:           fmt.Sprintf(UpdateMeetingLivestreamPath, params.MeetingID),
		DataParameters: &opts,
		URLParameters:  &params,
		HeadResponse:   true,
	})
}
