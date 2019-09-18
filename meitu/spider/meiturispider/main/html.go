package main

//func AnalyzeModelHomePageHtml(client *http.Client, url string, modelId int, i int) int {
//	resp, err := client.Get(url)
//	defer resp.Body.Close()
//	if resp.StatusCode > 400 {
//		return _const.UrlInvalid
//	}
//	doc, err := goquery.NewDocument(url)
//	if err != nil {
//		return _const.AnalysisHtmlFaild
//	}
//	if i == 1 && _const.SaveUserInfo {
//		cover, _ := doc.Find("div.left").Find("img").Attr("src")
//		right := doc.Find("div.right")
//
//		var map1 = make(map[string]string)
//		right.Find("p").Each(func(i int, s *goquery.Selection) {
//			html, _ := s.Html()
//			text := strings.Replace(strings.Replace(html, "<span>", ";", -1), "</span>", "", -1)
//			texts := strings.Split(text, ";")
//			for _, t := range texts {
//				if len(strings.Trim(t, "")) == 0 {
//
//				} else {
//					kv := strings.Split(t, ":")
//					map1[kv[0]] = kv[1]
//				}
//			}
//			fmt.Print(texts)
//			fmt.Println(strconv.Itoa(len(texts)))
//			//texts:=strings.Split(text,"<span>")
//
//		})
//
//		nicknames := right.Find("h1").Text()
//		shuoming := doc.Find("div.shuoming")
//		more := shuoming.Text()
//		tags := shuoming.Find("p").Text()
//
//		model := model.Models{
//			ID:            modelId,
//			Cover:         cover,
//			Name:          strings.Split(nicknames, "、")[0],
//			Nicknames:     nicknames,
//			More:          more,
//			Tags:          tags,
//			Birthday:      map1["生日"],
//			Constellation: map1["星座"],
//			Height:        map1["身高"],
//			Weight:        map1["体重"],
//			Dimensions:    map1["罩杯"],
//			Address:       map1["来自"],
//			Jobs:          map1["职业"],
//			Interest:      map1["兴趣"],
//		}
//		gorm.SaveModelInfo(modelId, &model)
//	}
//	return AnalyzeModelColumPage(modelId, doc)
//}
//
//func AnalyzeCompanyHomePageHtml(client *http.Client, url string, companyId int, i int) int {
//	resp, err := client.Get(url)
//	defer resp.Body.Close()
//	if resp.StatusCode > 400 {
//		return _const.UrlInvalid
//	}
//	doc, err := goquery.NewDocument(url)
//	if err != nil {
//		return _const.AnalysisHtmlFaild
//	}
//	if i == 1 && _const.SaveCompanyGroupRatation {
//		doc.Find("div.fenlei").Find("p").Find("a").Each(func(i int, s *goquery.Selection) {
//			name := s.Text()
//			homepage, _ := s.Attr("href")
//			id := getIdFromUri(homepage)
//			group := model.Groups{
//				Id:       id,
//				Name:     name,
//				Homepage: homepage,
//				Belong:   companyId,
//			}
//			gorm.SaveGroupInfo(group)
//			//fmt.Print()
//		})
//	}
//	return AnalyzeCompanyPage(doc)
//}
//
//func AnalyzeCompanyPage(doc *goquery.Document) int {
//	println(doc.Url.String())
//	hezi := doc.Find("div.hezi")
//	if hezi != nil {
//		hezi.Find("ul").Find("li").Each(func(i int, s *goquery.Selection) {
//			//path, _ := s.Find("a").Attr("href")
//			//columId:=getIdFromUri(path)
//			s.Find("p").Each(func(i int, s1 *goquery.Selection) {
//				if i == 0 {
//					modelhomepage, _ := s1.Find("a").Attr("href")
//					fmt.Println(modelhomepage)
//					modelId := getIdFromUri(modelhomepage)
//					//downloadColum(modelId,columId)
//					getModelColums(modelId)
//				}
//			})
//		})
//	} else {
//		return _const.AnalysisHtmlFaild
//	}
//	return _const.Success
//}
//
//func AnalyzeModelColumPage(modelId int, doc *goquery.Document) int {
//	println(doc.Url.String())
//	hezi := doc.Find("div.hezi")
//	if hezi != nil {
//		hezi.Find("ul").Find("li").Each(func(i int, s *goquery.Selection) {
//			a := s.Find("a")
//			path, _ := a.Attr("href")
//			//cover, _ := a.Find("img").Attr("src")
//			//nums:=	s.Find("span").Text()
//			//num,_:=strconv.Atoi(nums[0 : len(nums)-1])
//
//			paths := strings.Split(path, "/")
//			columId, _ := strconv.Atoi(paths[len(paths)-2])
//			println(path + " " + strconv.Itoa(len(path)) + strconv.Itoa(columId))
//			//if(!util.PathExists(strconv.Itoa(colum))){
//			//
//			//}
//			gorm.SaveColumRelation(modelId, columId)
//			err := downloadSingleColum(modelId, columId)
//			if err < 0 {
//				//continue
//			}
//		})
//	} else {
//		return _const.AnalysisHtmlFaild
//	}
//	return _const.Success
//}