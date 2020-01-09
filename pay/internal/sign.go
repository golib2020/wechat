package internal

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/golib2020/wechat/internal"
)

type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		Text string `xml:",cdata"`
	}{Text: string(c)}, start)
}

//SignCheck 数据有效检查
func SignCheck(mchKey string, signType interface{}) string {
	sortK := joinKeyValue(signType)
	sort.SliceStable(sortK, func(i, j int) bool {
		return sortK[i].key < sortK[j].key
	})
	var buf bytes.Buffer
	for index, val := range sortK {
		if index != 0 {
			buf.WriteString("&")
		}
		buf.WriteString(fmt.Sprintf("%s=%s", val.key, val.value))
	}
	buf.WriteString(fmt.Sprintf("&key=%s", mchKey))
	fmt.Println(buf.String())
	hs := internal.Md5(buf.Bytes())
	return strings.ToUpper(hs)
}

type item struct {
	key   string
	value string
}

func joinKeyValue(signType interface{}) []item {
	var sortK []item
	valueOf := reflect.ValueOf(signType)
	typeOf := reflect.TypeOf(signType)

	switch typeOf.Kind() {
	case reflect.Ptr:
		typeOf = typeOf.Elem()
		valueOf = valueOf.Elem()
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return sortK
	}
	for i := 0; i < typeOf.NumField(); i++ {
		s := valueOf.Field(i).String()
		if s == "" || s == "0" {
			continue
		}
		if valueOf.Field(i).Kind() == reflect.Struct {
			sortK = append(sortK, joinKeyValue(valueOf.Field(i).Interface())...)
			continue
		}
		tag, ok := typeOf.Field(i).Tag.Lookup("xml")
		if !ok || tag == "" {
			continue
		}
		tags := strings.Split(tag, ",")
		if len(tags) < 0 {
			continue
		}
		key := strings.TrimSpace(tags[0])
		if key == "sign" {
			continue
		}
		sortK = append(sortK, item{
			key:   key,
			value: fmt.Sprintf("%v", valueOf.Field(i).Interface()),
		})
	}
	return sortK
}
