package models

type PreSignedUrlReq struct {
	FileName string `json:"filename"`
	Size     string `json:"size"`
	MIME     string `json:"mime"`
}

type PreSignedUrlRes struct {
	Url string `json:"url"`
}
