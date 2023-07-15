package model

type Log struct {
    Endpoint    string  `json:"name,omitempty"`
	UserAgent	string	`json:"user_agent,omitempty"`
}