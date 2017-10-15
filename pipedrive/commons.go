package pipedrive

const (
	VisibleToOwnersAndFollowers = 1
	VisibleToWholeCompany       = 3
)

type Pagination struct {
	Start                 int  `json:"start"`
	Limit                 int  `json:"limit"`
	MoreItemsInCollection bool `json:"more_items_in_collection"`
}

type AdditionalData struct {
	CompanyID           int        `json:"company_id"`
	SinceTimestamp      string     `json:"since_timestamp"`
	LastTimestampOnPage string     `json:"last_timestamp_on_page"`
	Pagination          Pagination `json:"pagination"`
}

type DeleteMultipleOptions struct {
	Ids string `url:"ids,omitempty"`
}

type ErrorFields struct {
	Error     string `json:"error"`
	ErrorInfo string `json:"error_info"`
}
