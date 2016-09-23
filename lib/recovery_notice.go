package lib

import (
	"SoftWeb/lib/jlog"
	"errors"
	"fmt"
	"net/http"
)

type RecoveryNotice struct {
}

func NewRecoveryNotice() *RecoveryNotice {
	return &RecoveryNotice{}
}

func (rec *RecoveryNotice) ServeHTTP(rw http.ResponseWriter,
	r *http.Request, next http.HandlerFunc) {
	defer func() {
		err := recover()
		if err != nil {
			jlog.LogErrorSendMail(errors.New(fmt.Sprintf("%s", err)))
		}
	}()

	next(rw, r)
}
