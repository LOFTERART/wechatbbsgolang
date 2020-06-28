package services

import (
 "QUZHIYOU/models"
 "QUZHIYOU/serializer"
)

//定义前段传递的数据字段
type Comment struct{
	Name string`json:"name"`
	DiaryID uint `json:"diary_id"`
	UserID uint `json:"user_id"`
}

//增加Comment
func (item *Comment)CreateComment() serializer.Response{

 info:=models.Comment{
  Name:    item.Name,
  DiaryID: item.DiaryID,
  UserID:  item.UserID,
 }
 models.DB.Create(&info)

 return serializer.Response{
  Code: 0,
  Data:  nil,
  Msg:   "评论成功",

 }
}
//删 Comment
func (item *Comment)DelComment() serializer.Response{

 return serializer.Response{
  Code: 0,
  Data:  nil,
  Msg:   "删除成功",
  Error: "",
 }
}
//查ALL Comment
func (item *Comment)GetALLComment() serializer.Response{

 var infos []*models.Comment

 models.DB.
  Preload("User").
  Where("diary_id=?",item.DiaryID).Find(&infos)


 return serializer.Response{
  Code: 0,
  Data:  serializer.BuildCommentSSerializers(infos),
  Msg:   "查询ALL成功",
 }
}
//查ONE Comment
func (item *Comment)GetONEComment() serializer.Response{

 return serializer.Response{
  Code: 0,
  Data:  nil,
  Msg:   "查询ONE成功",
 }
}

//更新 Comment
func (item *Comment)UpdateComment() serializer.Response{

 return serializer.Response{
  Code: 0,
  Data: nil,
  Msg:  "更新成功",
 }
}
