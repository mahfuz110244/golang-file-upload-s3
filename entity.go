package main

type Response struct {
	StatusCode int          `json:"status_code"`
	Success    bool         `json:"success"`
	Message    string       `json:"message,omitempty"`
	Errors     []FieldError `json:"errors,omitempty"`
	Data       interface{}  `json:"data,omitempty"`
}

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type FileUploadResponse struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	Size         int    `json:"size"`
	Extension    string `json:"extension"`
	UploadStatus bool   `json:"upload_status"`
	Message      string `json:"message,omitempty"`
}
