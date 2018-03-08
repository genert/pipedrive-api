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

// Type of actions.
type EventAction string

const (
	ACTION_ADDED   EventAction = "added"
	ACTION_UPDATED EventAction = "updated"
	ACTION_MERGED  EventAction = "merged"
	ACTION_DELETED EventAction = "deleted"
	ACTION_ALL     EventAction = "all"
)

// Type of objects.
type EventObject string

const (
	OBJECT_ACTIVITY      EventObject = "activity"
	OBJECT_ACTIVTIY_TYPE EventObject = "activity_type"
	OBJECT_DEAL          EventObject = "deal"
	OBJECT_NOTE          EventObject = "note"
	OBJECT_ORGANIZATION  EventObject = "organization"
	OBJECT_PERSON        EventObject = "person"
	OBJECT_PIPELINE      EventObject = "pipeline"
	OBJECT_PRODUCT       EventObject = "product"
	OBJECT_STAGE         EventObject = "stage"
	OBJECT_USER          EventObject = "user"
	OBJECT_ALL_          EventObject = "*"
)

// Active flags
type ActiveFlag uint8

const (
	ActiveFlagEnabled  ActiveFlag = 1
	ActiveFlagDisabled ActiveFlag = 0
)

// Field types
type FieldType string

const (
	FieldTypeVarchar     FieldType = "varchar"
	FieldTypeVarcharAuto FieldType = "varchar_auto"
	FieldTypeText        FieldType = "text"
	FieldTypeDouble      FieldType = "double"
	FieldTypeMonetary    FieldType = "monetary"
	FieldTypeDate        FieldType = "date"
	FieldTypeSet         FieldType = "set"
	FieldTypeEnum        FieldType = "enum"
	FieldTypeUser        FieldType = "user"
	FieldTypeOrg         FieldType = "org"
	FieldTypePeople      FieldType = "people"
	FieldTypePhone       FieldType = "phone"
	FieldTypeTime        FieldType = "time"
	FieldTypeTimerange   FieldType = "timerange"
	FieldTypeDaterange   FieldType = "daterange"
)

// Visiblity
type VisibleTo uint8

const (
	VisibleToOwnersFollowers VisibleTo = 1
	VisibleToEntireCompany   VisibleTo = 3
)
