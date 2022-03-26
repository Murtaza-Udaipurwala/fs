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
	Header *multipart.FileHeader
	Ext    string
	Size   int64
	Onet   bool
	ID     string
}

const (
	maxUploadSize = 1024 * 1024 * 10
	uploadDir     = "uploads"
	scale         = 1024 * 1024
)

var chars = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E',
	'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
	'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
	'y', 'z',
}
