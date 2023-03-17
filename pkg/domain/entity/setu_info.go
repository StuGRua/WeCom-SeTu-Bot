package entity

// Query 请求参数
type Query struct {
	R18 int64 `json:"r18"`
	Num int64 `json:"num"`
	// Uid int `json:"uid"`
	// KeyWord string `json:"keyword"`
	Tag   []string `json:"tag"`
	Size  []string `json:"size"`
	Proxy string   `json:"proxy"`
	// DateAfter int `json:"dateAfter"`
	// DateBefore int `json:"dateBefore"`
	// Dsc bool `json:"dsc"`
}

// QueryResult 请求结果
type QueryResult struct {
	Error        string    `json:"error"`
	ArchiveSlice []Archive `json:"data"`
	//picPaths     []string
}

// PicUrl 图片链接
type PicUrl struct {
	Original string `json:"original"`
	Regular  string `json:"regular"`
	Small    string `json:"small"`
	Thumb    string `json:"thumb"`
	Mini     string `json:"mini"`
}

// Archive Pixiv投稿信息
type Archive struct {
	Pid     int64    `json:"pid"`
	P       int64    `json:"p"`
	Uid     int64    `json:"uid"`
	Title   string   `json:"title"`
	Author  string   `json:"author"`
	R18     bool     `json:"r18"`
	Width   int64    `json:"width"`
	Height  int64    `json:"height"`
	Tags    []string `json:"tags"`
	Ext     string   `json:"ext"`
	Date    int64    `json:"uploadDate"`
	Urls    PicUrl   `json:"urls"`
	DumpUrl string
}
