package meituri

type Home struct {
	Banners  []Banner  `json:"banners"`
	Apps []App `json:"apps"`
	Companys []Company `json:"companys"`
	Models   []Model   `json:"models"`
	//Colums   []Album   `json:"colums"`

	Tab bool `json:"has_tab"` //是否选择了tab 分类
	Bind bool `json:"has_bind"` //登录设备是否绑定
}
