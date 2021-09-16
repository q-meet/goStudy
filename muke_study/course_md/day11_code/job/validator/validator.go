package main
/*
import "github.com/go-playground/validator/v10"
import "fmt"

func main() {

	var req = RegisterReq{
		Username:       "Xargin",
		PasswordNew:    "ohno",
		PasswordRepeat: "ohn",
		Email:          "alex@abc.com",
	}

	err := validateFunc(req)
	fmt.Println(err)
}

type RegisterReq struct {
	// 字符串的 gt=0 表示长度必须 > 0，gt = greater than
	Username string `validate:"gt=0"`
	// 同上
	PasswordNew string `validate:"gt=0"`
	// eqfield 跨字段相等校验
	//PasswordRepeat string `validate:"eqfield=PasswordNew"`
	PasswordRepeat string `validate:"gt=0"`
	// 合法 email 格式校验
	Email string `validate:"email"`
}

var validate = validator.New()

func validateFunc(req RegisterReq) error {
	return validate.Struct(req)
}
*/