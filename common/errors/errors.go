package errors

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/kran891/go-lang-api/logger"
)

var enableLogging = true

func DisableLogging() {
	enableLogging = false
}

func Wrap(err error, msg string) error {
	logErrorf("%v : %+v", msg, err)
	return errors.Wrap(err, msg)
}

func Wrapf(err error, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	logErrorf("%v : +%v", msg, err)
	return errors.Wrap(err, msg)
}

func WithStack(err error) error {
	return errors.WithStack(err)
}

func WithMsg(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func WithMsgf(err error, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	logErrorf("%v : +%v", msg, err)
	return errors.Wrapf(err, format, args...)
}

func UnWrap(err error) error {
	return errors.Unwrap(err)
}

func New(msg string) error {
	return errors.New(msg)
}
func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
func Cause(err error) error {
	return errors.Cause(err)
}

func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

func logErrorf(format string, args ...interface{}) {
	if enableLogging {
		logger.Debugf(format, args...)
	}
}

type Const string

func (e Const) Error() string {
	return string(e)
}
