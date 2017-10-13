package pipedrive

type MailThreadService service

type MailThreadsListOptions struct {
	Folder string `url:"folder"`
	Start  int    `url:"start"`
	Limit  int    `url:"limit"`
}

func (s *MailThreadService) List(opt *MailThreadsListOptions) (*Currencies, *Response, error) {
	uri := s.client.CreateRequestUrl("/mailbox/mailThreads")
	req, err := s.client.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Currencies

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
