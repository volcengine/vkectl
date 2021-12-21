package client

import (
	"fmt"
	"net/url"
	"reflect"
)

// DealtWitQuery return the http query.
func DealtWitQuery(req interface{}) url.Values {
	query := url.Values{}
	reqType := reflect.TypeOf(req)
	reqValue := reflect.ValueOf(req)
	for i := 0; i < reqType.NumField(); i++ {
		structure := reqType.Field(i)
		key, value := "", ""
		if queryKey := structure.Tag.Get("query"); queryKey != "" {
			key = queryKey
			if defalutVal := structure.Tag.Get("default"); defalutVal != "" {
				value = defalutVal
			}
			name := structure.Name
			switch structure.Type.Name() {
			case "bool":
				{
					value = fmt.Sprintf("%v", reqValue.FieldByName(name).Bool())
				}
			case "float32", "float64":
				{
					value = fmt.Sprintf("%v", reqValue.FieldByName(name).Float())
				}
			case "int", "int8", "int16", "int32", "int64":
				{
					if v := reqValue.FieldByName(name).Int(); v != 0 {
						value = fmt.Sprintf("%v", v)
					}
				}
			case "uint", "uint8", "uint16", "uint32", "uint64":
				{
					if v := reqValue.FieldByName(name).Uint(); v != 0 {
						value = fmt.Sprintf("%v", v)
					}
				}
			case "string":
				{
					if v := reqValue.FieldByName(name).String(); v != "" {
						value = v
					}
				}
			default:
				{
					if v := reqValue.FieldByName(name); v.Kind() == reflect.Ptr && !v.IsNil() {
						value = fmt.Sprintf("%v", v.Elem())
					}
				}
			}
			if value != "" {
				query.Add(key, value)
			}
		}
	}
	return query
}
