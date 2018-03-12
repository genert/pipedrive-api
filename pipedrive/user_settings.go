package pipedrive

import (
	"context"
	"net/http"
)

// UserSettingsService handles user settings related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/UserSettings
type UserSettingsService service

// UserSettings represents a Pipedrive user settings.
type UserSettings struct {
	Success bool `json:"success"`
	Data    struct {
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
		NylasSync                            bool          `json:"nylas_sync"`
		OnboardingComplete                   bool          `json:"onboarding_complete"`
		ActivityEmailReminders               bool          `json:"activity_email_reminders"`
		ActivityEmailRemindersSendType       string        `json:"activity_email_reminders_send_type"`
		ActivityEmailRemindersAmountBefore   int           `json:"activity_email_reminders_amount_before"`
		ActivityNotificationsLanguageID      int           `json:"activity_notifications_language_id"`
		FileConvertAllowed                   bool          `json:"file_convert_allowed"`
		DefaultCurrency                      string        `json:"default_currency"`
		SendEmailNotifications               string        `json:"send_email_notifications"`
		ShowUpdateNotifications              bool          `json:"show_update_notifications"`
		CreateFolderInGoogleDrive            bool          `json:"create_folder_in_google_drive"`
		ShareGoogleDriveFilesWithCompany     bool          `json:"share_google_drive_files_with_company"`
		DealsTimelineDefaultField            string        `json:"deals_timeline_default_field"`
		DealsTimelineInterval                string        `json:"deals_timeline_interval"`
		DealsTimelineArrangeBy               string        `json:"deals_timeline_arrange_by"`
		DealsTimelineColumnCount             int           `json:"deals_timeline_column_count"`
		DealsTimelineShowProgress            bool          `json:"deals_timeline_show_progress"`
		ShareIncomingEmails                  bool          `json:"share_incoming_emails"`
		ConnectThreadsWithDeals              bool          `json:"connect_threads_with_deals"`
		EmailSignature                       bool          `json:"email_signature"`
		EmailSignatureText                   string        `json:"email_signature_text"`
		GlobalSearchSurveyLinkClicked        bool          `json:"global_search_survey_link_clicked"`
		GoogleCalendarActivityType           string        `json:"google_calendar_activity_type"`
		GoogleCalendarIgnoreActivityTypes    []interface{} `json:"google_calendar_ignore_activity_types"`
		UploadAllVisiblePersonsToGoogle      bool          `json:"upload_all_visible_persons_to_google"`
		FormatPhoneNumbersEnabled            bool          `json:"format_phone_numbers_enabled"`
		OnboardingCompletedTours             []interface{} `json:"onboarding_completed_tours"`
		OpenEmailLinksInNewTab               bool          `json:"open_email_links_in_new_tab"`
		TotalsConvertCurrency                bool          `json:"totals_convert_currency"`
		EmailSyncFilter                      string        `json:"email_sync_filter"`
		EmailSyncFilterLabels                []interface{} `json:"email_sync_filter_labels"`
		DealDetailsOpen                      bool          `json:"deal_details_open"`
		PersonDetailsOpen                    bool          `json:"person_details_open"`
		OrganizationDetailsOpen              bool          `json:"organization_details_open"`
		GoogleCalendarSubjectFormat          string        `json:"google_calendar_subject_format"`
		GoogleCalendarActivityReminders      []interface{} `json:"google_calendar_activity_reminders"`
		HideEmailSettingsPromotionBanner     bool          `json:"hide_email_settings_promotion_banner"`
		UserAdvancedDebugLogs                bool          `json:"user_advanced_debug_logs"`
		ShowFiltercolumnsTutorial            bool          `json:"show_filtercolumns_tutorial"`
		ShowDuplicates                       bool          `json:"show_duplicates"`
		ShowImportTutorial                   bool          `json:"show_import_tutorial"`
		ShowStatisticsTutorial               bool          `json:"show_statistics_tutorial"`
		ShowActivitycrossitemTutorial        bool          `json:"show_activitycrossitem_tutorial"`
		HasSeenDropboxOldDetailsDeprecNotice bool          `json:"has_seen_dropbox_old_details_deprec_notice"`
		LinkPersonToOrg                      bool          `json:"link_person_to_org"`
		TimezoneAutomaticUpdate              bool          `json:"timezone_automatic_update"`
		UsePipedriveMailtoLinks              bool          `json:"use_pipedrive_mailto_links"`
		AbMailPromotion1Enabled              bool          `json:"ab_mail_promotion1_enabled"`
		MailFiltersInbox                     []interface{} `json:"mail_filters_inbox"`
		MailFiltersDrafts                    []interface{} `json:"mail_filters_drafts"`
		MailFiltersSent                      []interface{} `json:"mail_filters_sent"`
		MailFiltersArchive                   []interface{} `json:"mail_filters_archive"`
		PromoteGsFilters                     bool          `json:"promote_gs_filters"`
		PromoteGsNotesAndFields              bool          `json:"promote_gs_notes_and_fields"`
		HasSeenHelpTooltip                   bool          `json:"has_seen_help_tooltip"`
		InvitationLastTime                   int           `json:"invitation_last_time"`
		InvitationResentCount                int           `json:"invitation_resent_count"`
		ExpandedCalendarAllday               bool          `json:"expanded_calendar_allday"`
		TimelinePanelOpen                    bool          `json:"timeline_panel_open"`
		TimelinePanelOnboardingDone          bool          `json:"timeline_panel_onboarding_done"`
		HasNotifications                     bool          `json:"has_notifications"`
		ActivitiesViewMode                   string        `json:"activities_view_mode"`
		ActivityQuickfilterRange             string        `json:"activity_quickfilter_range"`
		CurrentPipelineID                    int           `json:"current_pipeline_id"`
		DealsTimelineShowWeightedValues      bool          `json:"deals_timeline_show_weighted_values"`
		DealsViewMode                        string        `json:"deals_view_mode"`
		FeedViewType                         string        `json:"feed_view_type"`
		FilterActivities                     string        `json:"filter_activities"`
		FilterDeals                          string        `json:"filter_deals"`
		FilterOrg                            string        `json:"filter_org"`
		FilterPeople                         string        `json:"filter_people"`
		FilterPipeline1                      string        `json:"filter_pipeline_1"`
		FilterPipeline2                      string        `json:"filter_pipeline_2"`
		FilterProducts                       string        `json:"filter_products"`
		GoogledocsFolderID                   string        `json:"googledocs_folder_id"`
		GoogleCalendarSync                   bool          `json:"google_calendar_sync"`
	} `json:"data"`
}

// List settings of authorized user.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/UserSettings/get_userSettings
func (s *UserSettingsService) List(ctx context.Context) (*UserSettings, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/userSettings", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *UserSettings

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
