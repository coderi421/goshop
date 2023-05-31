package code

import (
	"net/http"

	"github.com/coderi421/gframework/pkg/common/util/sliceutil"
	"github.com/coderi421/gframework/pkg/errors"
)

type ErrCode struct {
	//错误码
	C int

	//http的状态码
	HTTP int

	//扩展字段
	Ext string

	//引用文档
	Ref string
}

func (e ErrCode) HTTPStatus() int {
	return e.HTTP
}

func (e ErrCode) String() string {
	return e.Ext
}

func (e ErrCode) Reference() string {
	return e.Ref
}

func (e ErrCode) Code() int {
	if e.C == 0 {
		return http.StatusInternalServerError
	}
	return e.C
}

func register(code int, httpStatus int, message string, refs ...string) {
	found := sliceutil.Find([]int{200, 400, 401, 403, 404, 500}, httpStatus)

	if !found {
		panic("http code not in `200,400,401,403,404,500`")
	}
	var ref string
	if len(refs) > 0 {
		ref = refs[0]
	}
	coder := ErrCode{
		C:    code,
		HTTP: httpStatus,
		Ext:  message,
		Ref:  ref,
	}
	errors.MustRegister(coder)
}

var _ errors.Coder = (*ErrCode)(nil)
