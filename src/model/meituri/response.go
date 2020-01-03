package meituri

type LoginResp struct {
	User User `json:"user"`
	Tab bool `json:"has_tab"` //是否选择了tab 分类
	Bind bool `json:"has_bind"` //登录设备是否绑定
}
