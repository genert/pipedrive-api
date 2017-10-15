package pipedrive

import "net/http"

type AuthorizationsService service

type Authorizations struct {
	Success bool `json:"success"`
	Data    []struct {
		UserID    int    `json:"user_id"`
		CompanyID int    `json:"company_id"`
		APIToken  string `json:"api_token"`
		AddTime   string `json:"add_time"`
		Company   struct {
			Info struct {
				ID                 int         `json:"id"`
				Name               string      `json:"name"`
				CreatorCompanyID   interface{} `json:"creator_company_id"`
				PlanID             int         `json:"plan_id"`
				Identifier         string      `json:"identifier"`
				Domain             string      `json:"domain"`
				BillingCurrency    string      `json:"billing_currency"`
				AddTime            string      `json:"add_time"`
				Status             string      `json:"status"`
				TrialEnds          string      `json:"trial_ends"`
				CancelledFlag      bool        `json:"cancelled_flag"`
				CancelTime         interface{} `json:"cancel_time"`
				Country            string      `json:"country"`
				PromoCode          string      `json:"promo_code"`
				UsedPromoCodeKey   string      `json:"used_promo_code_key"`
				AccountIsOpen      bool        `json:"account_is_open"`
				AccountIsNotPaying bool        `json:"account_is_not_paying"`
			} `json:"info"`
			Features []string `json:"features"`
			Settings struct {
				ShowGettingStartedVideo                 bool          `json:"show_getting_started_video"`
				ListLimit                               int           `json:"list_limit"`
				BetaApp                                 bool          `json:"beta_app"`
				FileUploadDestination                   string        `json:"file_upload_destination"`
				CalltoLinkSyntax                        string        `json:"callto_link_syntax"`
				AutofillDealExpectedCloseDate           bool          `json:"autofill_deal_expected_close_date"`
				PersonDuplicateCondition                string        `json:"person_duplicate_condition"`
				OrganizationDuplicateCondition          string        `json:"organization_duplicate_condition"`
				AddFollowersWhenImporting               bool          `json:"add_followers_when_importing"`
				SearchBackend                           string        `json:"search_backend"`
				BillingManagedBySales                   bool          `json:"billing_managed_by_sales"`
				MaxDealAgeInAverageProgressCalculation  int           `json:"max_deal_age_in_average_progress_calculation"`
				ThirdPartyLinks                         []interface{} `json:"third_party_links"`
				ElasticWriteTargetDuringMigration       string        `json:"elastic_write_target_during_migration"`
				AutoCreateNewPersonsFromForwarderEmails bool          `json:"auto_create_new_persons_from_forwarder_emails"`
				CompanyAdvancedDebugLogs                bool          `json:"company_advanced_debug_logs"`
				DealBlockOrder                          []struct {
					Type    string `json:"type"`
					Visible bool   `json:"visible"`
				} `json:"deal_block_order"`
				PersonBlockOrder []struct {
					Type    string `json:"type"`
					Visible bool   `json:"visible"`
				} `json:"person_block_order"`
				OrganizationBlockOrder []struct {
					Type    string `json:"type"`
					Visible bool   `json:"visible"`
				} `json:"organization_block_order"`
				NylasSync          bool `json:"nylas_sync"`
				OnboardingComplete bool `json:"onboarding_complete"`
			} `json:"settings"`
		} `json:"company"`
	} `json:"data"`
	AdditionalData struct {
		User struct {
			Profile struct {
				ID              int         `json:"id"`
				Email           string      `json:"email"`
				Name            string      `json:"name"`
				IsAdmin         bool        `json:"is_admin"`
				DefaultCurrency string      `json:"default_currency"`
				IconURL         interface{} `json:"icon_url"`
				Activated       bool        `json:"activated"`
			} `json:"profile"`
			Locale struct {
				Language        string `json:"language"`
				Country         string `json:"country"`
				Uses12HourClock bool   `json:"uses_12_hour_clock"`
			} `json:"locale"`
			Timezone struct {
				Name   string `json:"name"`
				Offset int    `json:"offset"`
			} `json:"timezone"`
		} `json:"user"`
		MultipleCompanies bool `json:"multiple_companies"`
		DefaultCompanyID  int  `json:"default_company_id"`
	} `json:"additional_data"`
}

type AuthorizationsGetOptions struct {
	Email    string `url:"email"`
	Password string `url:"password"`
}

// Returns all authorizations for a particular user. Authorization objects contain
// the API tokens the user has with different company accounts in Pipedrive.
// These can be fetched without an API token but using the email and password of the user.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Authorizations
func (s *AuthorizationsService) Get(opt *AuthorizationsGetOptions) (*Authorizations, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/authorizations", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Authorizations

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
