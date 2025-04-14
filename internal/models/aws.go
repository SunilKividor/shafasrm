package models

type PreSignedUrlReq struct {
	Size string `json:"size"`
	MIME string `json:"mime"`
}

type PreSignedUrlRes struct {
	Key string `json:"key"`
	Url string `json:"url"`
}

type PhotoObject struct {
	Key       string `json:"photo_key"`
	IsPrimary bool   `json:"is_primary"`
}

type PhotoResponse struct {
	URL       string `json:"url"`
	IsPrimary bool   `json:"is_primary"`
}
