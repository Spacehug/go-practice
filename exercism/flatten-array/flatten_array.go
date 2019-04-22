package flatten

func Flatten(in interface{}) []interface{} {
	fl := []interface{}{}
	for _, v := range in.([]interface{}) {
		if _, ok := v.([]interface{}); ok {fl = append(fl, Flatten(v)...)}
		if _, ok := v.(int); ok {fl = append(fl, v)}
	}
	return fl
}