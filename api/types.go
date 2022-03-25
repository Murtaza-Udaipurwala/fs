package api

import (
	"mime/multipart"
	"time"
)

type MetaData struct {
	Expiry    time.Time `json:"expiry"`
	IsOneTime bool      `json:"is_one_time"`
}

type HTTPErr struct {
	Msg    string
	Status int
}

type Resp struct {
	Err string `json:"error,omitempty"`
	URL string `json:"url,omitempty"`
}

type File struct {
	File multipart.File
	Ext  string
}

const (
	maxUploadSize = 1024 * 1024 * 10
	uploadDir     = "uploads"
)
