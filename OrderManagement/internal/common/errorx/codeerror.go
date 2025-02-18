package errorx

import (
	"OrderManagement/OrderManagement/internal/common/errorx/errorcode"
	"errors"
	"fmt"
)

type (
	CodeError interface {
		error
		Status() int
		Code() int
		PortalError() string
	}

	codeError struct {
		status     int
		code       int
		desc       string
		portalDesc string // 是经过产品定义的严谨可对外的错误描述
	}
)

func (err *codeError) Error() string {
	return err.desc
}
func (err *codeError) PortalError() string {
	return err.portalDesc
}

func (err *codeError) Code() int {
	return err.code
}

func (err *codeError) Status() int {
	return err.status
}

func (err *codeError) String() string {
	return fmt.Sprintf("Status: %d, Code: %d, Desc: %s", err.status, err.code, err.desc)
}

func NewCodeError(code int, desc string) CodeError {
	return NewStatCodeError(400, code, desc)
}

func NewDefaultError(desc string) CodeError {
	return NewStatCodeError(400, 406001, desc)
}

// NewPortalError 创建用户可见的(portalDesc)的error类型
// 假如这个error最终被跑到外部portal，则前端会暴露portalDesc
// portalDesc 是经过产品定义的严谨可对外的错误描述
func NewPortalError(status, code int, desc, portalDesc string) CodeError {
	return &codeError{
		status:     status,
		code:       code,
		desc:       desc,
		portalDesc: portalDesc,
	}
}

func NewDefaultPortalError(desc, portalDesc string) CodeError {
	return NewPortalError(400, 406001, desc, portalDesc)
}
func NewCodePortalError(code int, desc, portalDesc string) CodeError {
	return NewPortalError(400, code, desc, portalDesc)
}

func NewDBError(desc string) CodeError {
	return NewStatCodeError(500, errorcode.ErrDBError, fmt.Sprintf("数据库异常: %s", desc))
}

func NewStatCodeError(status, code int, desc string) CodeError {
	return &codeError{
		status: status,
		code:   code,
		desc:   desc,
	}
}

// ----------------------------------------------------------------------------
//

func FromError(err error) (CodeError, bool) {
	if err == nil {
		return nil, false
	}
	var ce CodeError
	if ok := errors.As(err, &ce); ok {
		return ce, ok
	}

	return nil, false
}
