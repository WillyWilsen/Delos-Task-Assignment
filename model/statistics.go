package model

type EndpointStatistics struct {
	Count           int64 `json:"count"`
	UniqueUserAgent int64 `json:"unique_user_agent"`
}

type Statistics map[string]*EndpointStatistics