package models

type Doc struct {
	ParentId string `json:"parentId"`
	// '1' 文件夹  '2' markdown文档
	DocType  string `json:"docType"`
	Title    string `json:"title"`
	FileId   string `json:"fileId"`
	Children []Doc  `json:"children"`
	*BaseModelImpl
}

type DocFile struct {
	List []Doc `json:"list"`
}
