package app

import (
	"github.com/astaxie/beego/validation"
	"space/pkg/logutil"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logutil.Log.Errorf("markerrors  error  %v ,%v",err.Key, err.Message)
	}

	return
}
