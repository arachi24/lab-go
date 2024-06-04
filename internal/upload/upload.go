package upload

type UploadResponse struct {
	Url string `json:"url"`
}

type UploadRequest struct {
	FileType      string `json:"file_type"`
	ContentUpload []byte `json:"content_upload"`
}
