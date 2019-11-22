package response

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

//ReadBody 读取相应的body数据
func ReadBody(reader io.Reader, data interface{}) error {
	//context.WithValue()
	bts, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	//fmt.Printf("%s\n", bts)
	e := new(Error)
	if err := json.Unmarshal(bts, e); err != nil {
		return err
	}
	if err := e.Check(); err != nil {
		return err
	}

	if data != nil {
		switch v := data.(type) {
		case *bytes.Buffer:
			v.Write(bts)
			return nil
		default:
			return json.Unmarshal(bts, v)
		}
	}
	return nil
}
