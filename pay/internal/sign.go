package internal

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/golib2020/wechat/internal"
)

type CDATA struct {
	Text string `xml:",cdata"`
}

//SignCheck 数据有效检查
func SignCheck(mchKey string, signType interface{}) string {
	valueOf := reflect.ValueOf(signType)
	typeOf := reflect.TypeOf(signType)

	switch typeOf.Kind() {
	case reflect.Ptr:
		typeOf = typeOf.Elem()
		valueOf = valueOf.Elem()
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return ""
	}

	storeKv := make(map[string]string)
	var sortK []string
	for i := 0; i < typeOf.NumField(); i++ {
		tag, ok := typeOf.Field(i).Tag.Lookup("xml")
		if !ok || tag == "" {
			continue
		}
		tags := strings.Split(tag, ",")
		if len(tags) < 0 {
			continue
		}
		s := valueOf.Field(i).String()
		if s == "" || s == "0" {
			continue
		}
		key := strings.TrimSpace(tags[0])

		if key == "sign" {
			continue
		}

		switch valueOf.Field(i).Interface().(type) {
		case CDATA:
			storeKv[key] = valueOf.Field(i).Interface().(CDATA).Text
		default:
			storeKv[key] = fmt.Sprintf("%v", valueOf.Field(i).Interface())
		}
		sortK = append(sortK, key)
	}

	sort.Strings(sortK)

	var buf bytes.Buffer
	for index, val := range sortK {
		if index != 0 {
			buf.WriteString("&")
		}
		buf.WriteString(val + "=" + storeKv[val])
	}
	buf.WriteString("&key=")
	buf.WriteString(mchKey)
	hs := internal.Md5(buf.Bytes())
	return strings.ToUpper(hs)
}
