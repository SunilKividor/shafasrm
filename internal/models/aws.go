package models

type PreSignedUrlReq struct {
	Size string `json:"size"`
	MIME string `json:"mime"`
}

type PreSignedUrlRes struct {
	Key string `json:"key"`
	Url string `json:"url"`
}
