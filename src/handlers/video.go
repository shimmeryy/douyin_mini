package handlers

type OSSUploadResponse struct {
	Response
	Url string `json:"url"`
}
