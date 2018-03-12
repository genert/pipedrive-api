package pipedrive

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// FilesService handles files related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Files
type FilesService service

// File represents a Pipedrive file.
type File struct {
	ID             int         `json:"id"`
	UserID         int         `json:"user_id"`
	DealID         int         `json:"deal_id"`
	PersonID       int         `json:"person_id"`
	OrgID          int         `json:"org_id"`
	ProductID      interface{} `json:"product_id"`
	EmailMessageID interface{} `json:"email_message_id"`
	ActivityID     interface{} `json:"activity_id"`
	NoteID         interface{} `json:"note_id"`
	LogID          interface{} `json:"log_id"`
	AddTime        string      `json:"add_time"`
	UpdateTime     string      `json:"update_time"`
	FileName       string      `json:"file_name"`
	FileType       string      `json:"file_type"`
	FileSize       int         `json:"file_size"`
	ActiveFlag     bool        `json:"active_flag"`
	InlineFlag     bool        `json:"inline_flag"`
	RemoteLocation string      `json:"remote_location"`
	RemoteID       string      `json:"remote_id"`
	Cid            interface{} `json:"cid"`
	S3Bucket       interface{} `json:"s3_bucket"`
	MailMessageID  interface{} `json:"mail_message_id"`
	DealName       string      `json:"deal_name"`
	PersonName     string      `json:"person_name"`
	OrgName        string      `json:"org_name"`
	ProductName    interface{} `json:"product_name"`
	URL            string      `json:"url"`
	Name           string      `json:"name"`
	Description    interface{} `json:"description"`
}

func (f File) String() string {
	return Stringify(f)
}

// FileResponse represents single file response.
type FileResponse struct {
	Success        bool           `json:"success"`
	Data           File           `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// FilesResponse represents multiple files response.
type FilesResponse struct {
	Success        bool           `json:"success"`
	Data           []File         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// List all files.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Files/get_files
func (s *FilesService) List(ctx context.Context) (*FilesResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/files", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *FilesResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetByID returns specific file.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Files/get_files_id
func (s *FilesService) GetByID(ctx context.Context, id int) (*FileResponse, *Response, error) {
	uri := fmt.Sprintf("/files/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *FileResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// GetDownloadLinkByID returns link for specific file.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Files/get_files_id_download
func (s *FilesService) GetDownloadLinkByID(id int) (string, *http.Request, error) {
	uri := fmt.Sprintf("/files/%v/download", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return "", nil, err
	}

	return string(req.URL.Scheme + "://" + req.URL.Host + req.URL.Path), req, nil
}

// Upload a file.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Files/post_files
func (s *FilesService) Upload(ctx context.Context, fileName string, filePath string) (*FileResponse, *Response, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, nil, err
	}

	fileInfo, err := file.Stat()

	if err != nil {
		return nil, nil, err
	}

	var body *bytes.Buffer
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileName, fileInfo.Name())

	if err != nil {
		return nil, nil, err
	}

	part.Write(fileContents)

	err = writer.Close()

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, "/files", nil, body)

	if err != nil {
		return nil, nil, err
	}

	var record *FileResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// CreateRemoteLinkedFileOptions specifices the optional parameters to the
// FilesService.CreateRemoteLinkedFile method.
type CreateRemoteLinkedFileOptions struct {
	FileType       string `url:"file_type,omitempty"`
	Title          string `url:"title,omitempty"`
	ItemType       string `url:"item_type,omitempty"`
	ItemID         uint   `url:"item_id,omitempty"`
	RemoteLocation string `url:"remote_location,omitempty"`
}

// CreateRemoteLinkedFile creates a remote file and link it to an item.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Files/post_files_remote
func (s *FilesService) CreateRemoteLinkedFile(ctx context.Context, opt *CreateRemoteLinkedFileOptions) (*FileResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/files/remote", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *FileResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// LinkRemoteFileToItemOptions specifices the optional parameters to the
// FilesService.LinkRemoteFileToItem method.
type LinkRemoteFileToItemOptions struct {
	ItemType       string `url:"item_type"`
	ItemID         uint   `url:"item_id"`
	RemoteID       uint   `url:"remote_id"`
	RemoteLocation string `url:"remote_location"`
}

// LinkRemoteFileToItem links an existing remote file (googledrive, etc) to the item you supply.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Files/post_files_remoteLink
func (s *FilesService) LinkRemoteFileToItem(ctx context.Context, opt *LinkRemoteFileToItemOptions) (*FileResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/files/remoteLink", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *FileResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// UpdateFileDetailsOptions specifices the optional parameters to the
// FilesService.Update method.
type UpdateFileDetailsOptions struct {
	Name        string `url:"name"`
	Description string `url:"description"`
}

// Update the properties of a file.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Files/put_files_id
func (s *FilesService) Update(ctx context.Context, id int, opt *UpdateFileDetailsOptions) (*FileResponse, *Response, error) {
	uri := fmt.Sprintf("/files/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *FileResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Delete marks a file as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Files/delete_files_id
func (s *FilesService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/files/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
