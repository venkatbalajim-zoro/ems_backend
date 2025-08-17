package models

type TestCase struct {
	Name               string            `json:"name"`
	Method             string            `json:"method"`
	Endpoint           string            `json:"endpoint"`
	Headers            map[string]string `json:"headers"`
	Body               map[string]any    `json:"body"`
	ExpectedStatusCode int               `json:"expected_status_code"`
	ExpectedMessage    string            `json:"expected_message"`
}
