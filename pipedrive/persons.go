package pipedrive

type PersonsService service

type Person struct {
	ID        int `json:"id"`
	CompanyID int `json:"company_id"`
	OwnerID   struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		HasPic     bool   `json:"has_pic"`
		PicHash    string `json:"pic_hash"`
		ActiveFlag bool   `json:"active_flag"`
		Value      int    `json:"value"`
	} `json:"owner_id"`
	OrgID                       interface{} `json:"org_id"`
	Name                        string      `json:"name"`
	FirstName                   string      `json:"first_name"`
	LastName                    string      `json:"last_name"`
	OpenDealsCount              int         `json:"open_deals_count"`
	RelatedOpenDealsCount       int         `json:"related_open_deals_count"`
	ClosedDealsCount            int         `json:"closed_deals_count"`
	RelatedClosedDealsCount     int         `json:"related_closed_deals_count"`
	ParticipantOpenDealsCount   int         `json:"participant_open_deals_count"`
	ParticipantClosedDealsCount int         `json:"participant_closed_deals_count"`
	EmailMessagesCount          int         `json:"email_messages_count"`
	ActivitiesCount             int         `json:"activities_count"`
	DoneActivitiesCount         int         `json:"done_activities_count"`
	UndoneActivitiesCount       int         `json:"undone_activities_count"`
	ReferenceActivitiesCount    int         `json:"reference_activities_count"`
	FilesCount                  int         `json:"files_count"`
	NotesCount                  int         `json:"notes_count"`
	FollowersCount              int         `json:"followers_count"`
	WonDealsCount               int         `json:"won_deals_count"`
	RelatedWonDealsCount        int         `json:"related_won_deals_count"`
	LostDealsCount              int         `json:"lost_deals_count"`
	RelatedLostDealsCount       int         `json:"related_lost_deals_count"`
	ActiveFlag                  bool        `json:"active_flag"`
	Phone                       []struct {
		Value   string `json:"value"`
		Primary bool   `json:"primary"`
	} `json:"phone"`
	Email []struct {
		Value   string `json:"value"`
		Primary bool   `json:"primary"`
	} `json:"email"`
	FirstChar                       string      `json:"first_char"`
	UpdateTime                      string      `json:"update_time"`
	AddTime                         string      `json:"add_time"`
	VisibleTo                       string      `json:"visible_to"`
	PictureID                       interface{} `json:"picture_id"`
	NextActivityDate                interface{} `json:"next_activity_date"`
	NextActivityTime                interface{} `json:"next_activity_time"`
	NextActivityID                  interface{} `json:"next_activity_id"`
	LastActivityID                  int         `json:"last_activity_id"`
	LastActivityDate                string      `json:"last_activity_date"`
	TimelineLastActivityTime        interface{} `json:"timeline_last_activity_time"`
	TimelineLastActivityTimeByOwner interface{} `json:"timeline_last_activity_time_by_owner"`
	LastIncomingMailTime            interface{} `json:"last_incoming_mail_time"`
	LastOutgoingMailTime            interface{} `json:"last_outgoing_mail_time"`
	OrgName                         interface{} `json:"org_name"`
	OwnerName                       string      `json:"owner_name"`
	CcEmail                         string      `json:"cc_email"`
}

type PersonsRespose struct {
	Success        bool           `json:"success"`
	Data           []Person       `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}
