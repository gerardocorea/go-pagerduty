package pagerduty

type Filter struct {
	SubjectMode    string `json:"subject_mode,omitempty"`
	SubjectRegex   string `json:"subject_regex,omitempty"`
	BodyMode       string `json:"body_mode,omitempty"`
	BodyRegex      string `json:"body_regex,omitempty"`
	FromEmailMode  string `json:"from_email_mode,omitempty"`
	FromEmailRegex string `json:"from_email_regex,omitempty"`
	ID             string `json:"id,omitempty"`
}
