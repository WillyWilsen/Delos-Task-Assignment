package model

type Pond struct {
    ID      int     `json:"id,omitempty"`
    Name    string  `json:"name,omitempty"`
	FarmID	int		`json:"farm_id,omitempty"`
}