package meituri

type LoginResp struct {
	User User `gson:"user"`

	Tab bool `gson:"has_tab"`//是否选择了tab 分类

	Bind bool `gson:"has_bind"`//登录设备是否绑定

}