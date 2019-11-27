package meituri

type Home struct {
	Banners  []Banner  `json:"banners"`
	Companys []Company `json:"companys"`
	Models   []Model   `json:"models"`
	//Colums   []Album   `json:"colums"`
}
