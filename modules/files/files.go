package files

import "mime/multipart"

type FileReq struct {
	File        *multipart.FileHeader `json:"file" form:"file"`
	Destination string                `json:"destination" form:"destination"`
	Extension   string
	FileName    string
}

type FileRes struct {
	FileName string `json:"file_name"`
	Url      string `json:"url"`
}

type DeleteFileReq struct {
	Destination string `json:"destination" form:"destination"`
}
