package models

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	} `json:"data"`
}

type ResponseList struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		Key string `json:"Key"`
	} `json:"data"`
}

type Secret struct {
	Key   string
	Value string `json:",omitempty"`
}
