package pipedrive

type Pagination struct {
	Start                 int  `json:"start"`
	Limit                 int  `json:"limit"`
	MoreItemsInCollection bool `json:"more_items_in_collection"`
}

type AdditionalData struct {
	SinceTimestamp      string     `json:"since_timestamp"`
	LastTimestampOnPage string     `json:"last_timestamp_on_page"`
	Pagination          Pagination `json:"pagination"`
}
