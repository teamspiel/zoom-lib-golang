package zoom

import "fmt"

type (
	MeetingRegistrantCreateBatchParams struct {
		MeetingID int `url:"-"`
	}

	Registrant struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	MeetingRegistrantCreateBatchInput struct {
		AutoApprove bool          `json:"auto_approve,omitempty"`
		Registrants []*Registrant `json:"registrants"`
	}

	RegistrantResult struct {
		Email        string `json:"email"`
		JoinUrl      string `json:"join_url"`
		RegistrantId string `json:"registrant_id"`
	}

	MeetingRegistrantCreateBatchResponse struct {
		Registrants []*RegistrantResult `json:"registrants"`
	}
)

// MeetingRegistrantCreateBatchPath - Add many meeting registrants
const MeetingRegistrantCreateBatchPath = "/meetings/%d/batch_registrants"

func MeetingRegistrantCreateBatch(params MeetingRegistrantCreateBatchParams, input MeetingRegistrantCreateBatchInput) (MeetingRegistrantCreateBatchResponse, error) {
	return defaultClient.MeetingRegistrantCreateBatch(params, input)
}

func (c *Client) MeetingRegistrantCreateBatch(params MeetingRegistrantCreateBatchParams, input MeetingRegistrantCreateBatchInput) (MeetingRegistrantCreateBatchResponse, error) {
	var response MeetingRegistrantCreateBatchResponse

	return response, c.requestV2(requestV2Opts{
		Method:        Post,
		Path:          fmt.Sprintf(MeetingRegistrantCreateBatchPath, params.MeetingID),
		URLParameters: &input,
		Ret:           &response,
	})
}
