package support

type ParamImage struct {
	FilePath    string `json:"path"`
	FileContent string `json:"content"`
	Destination string `json:"destination"`
	Filename    string `json:"filename"`
	MaxSize     uint   `json:"maxSize"`
	Quality     int    `json:"quality"`
}
