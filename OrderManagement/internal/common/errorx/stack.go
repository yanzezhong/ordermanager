package errorx

import "github.com/pkg/errors"

var (
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)
