package main

type Project struct {
	Id          int      `orm:"pk;auto;unique" json:"id"`
	ProjectId   string   `json:"project_id"`                 //uuid
	Name        string   `json:"name" form:"name"`           //name
	Type        string   `json:"type" form:"type"`           //github,gitlab
	Url         string   `json:"url" form:"url"`             //仓库地址
	Path        string   `json:"path" form:"path"`           //file 地址
	Branch      string   `json:"branch" form:"branch"`       //分支
	MainPath    string   `json:"main_path" form:"main_path"` //main 文件地址
	SecretToken string   `json:"secret_token" form:"secret_token"`
	Build       []*Build `orm:"reverse(many)" json:"-"`
}

type Build struct {
	Id        int      `orm:"pk;auto;unique" json:"id"` //主键
	Name      string   `json:"name" json:"name"`
	Result    string   `json:"result" json:"result"`
	Project   *Project `orm:"rel(fk)" json:"-"`
	IsSuccess bool     `json:"is_success"`
}
