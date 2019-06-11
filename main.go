package helper3

import (
	"encoding/json"

	"github.com/ypapax/glog3"
)

func ToJsonB(inp interface{}) string {
	b, err := json.MarshalIndent(&inp, "", "   ")
	if err != nil {
		glog3.Error(err)
		return err.Error()
	}
	return string(b)
}
