package services

type DiaryHomeService struct {
	Page int `json:"page"binding:"required"`
	Size int `json:"size"binding:"required"`
}

//func (this *DiaryHomeService) GetAllDiary() []*serializer.Diary {
//
//	page := this.Page
//	size := this.Size
//
//	start := (page - 1) * size
//
//	var diary []*models.Diary
//
//	models.DB.Limit(size).Offset(start).Find(&diary)
//	return serializer.BuildDiarys(diary)
//
//}
