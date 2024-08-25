package validator

import (
	"ginblog/utils/errmsg"
	"log"
	"reflect"

	"github.com/go-playground/locales/zh_Hans_CN"
	universalTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

func Validate(data interface{}) (string, int) {
	validater := validator.New()                                                 //验证器实例
	trans, _ := universalTrans.New(zh_Hans_CN.New()).GetTranslator("zh_Hans_CN") //创建翻译器对象
	err := zhTrans.RegisterDefaultTranslations(validater, trans)
	if err != nil {
		log.Panicln(err)
	}

	validater.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	err = validater.Struct(data) //执行对 data 结构体的验证
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) { // 类型断言,将值赋给变量
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
