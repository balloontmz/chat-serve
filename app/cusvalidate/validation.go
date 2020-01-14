package cusvalidate

import "gopkg.in/go-playground/validator.v9"

//MsgListQuery 消息列表查询条件
type MsgListQuery struct {
	GroupIDS []int  `json:"group_ids" form:"group_ids" query:"group_ids[]" validate:"dive,required"` //TODO: 此处注意 query 和 form 的[]!!!
}

//CustomValidator 自定义验证器
type CustomValidator struct {
	Validator *validator.Validate
}

//Validate 实现接口
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}