package assert

import (
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func structfmt(v reflect.Value) (string, string) {
	typ := v.Type()
	nf := typ.NumField()
	str := "\n{\n"

	inner := "\t"
	for i := 0; i < nf; i++ {
		tf := typ.Field(i)
		fv := v.Field(i)
		v, t := format(fv)
		inner += fmt.Sprintf("%s:\t%s(%s)\n", tf.Name, v, t)
	}
	inner = strings.Replace(inner, "\n", "\n\t", -1)
	inner = strings.TrimSuffix(inner, "\t")
	str += inner

	str += "}"
	return str, v.Type().String()
}

func strfmt(v reflect.Value) (string, string) {
	return fmt.Sprintf("%q", v.String()), v.Type().String()
}

func intfmt(v reflect.Value) (string, string) {
	return fmt.Sprintf("%d", v.Int()), v.Type().String()
}

func uintfmt(v reflect.Value) (string, string) {
	return fmt.Sprintf("%d", v.Uint()), v.Type().String()
}

func boolfmt(v reflect.Value) (string, string) {
	return fmt.Sprintf("%t", v.Bool()), v.Type().String()
}

func slicefmt(v reflect.Value) (string, string) {
	length := v.Len()
	slice := v.Slice(0, length)

	str := "["
	for i := 0; i < length; i += 1 {
		v, _ := format(slice.Index(i))
		str = fmt.Sprintf("%s%v, ", str, v)
	}
	str += "\b\b]"
	return fmt.Sprintf("%v", str), fmt.Sprintf("%s[%d]", v.Type().String(), length)
}

func format(v reflect.Value) (string, string) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intfmt(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintfmt(v)
	case reflect.String:
		return strfmt(v)
	case reflect.Bool:
		return boolfmt(v)
	case reflect.Slice:
		return slicefmt(v)
	case reflect.Struct:
		return structfmt(v)
	case reflect.Ptr:
		v := reflect.Indirect(v)
		vs, ts := format(v)
		return vs, "&" + ts
	}
	return "", ""
}

func getInfo() string {
	_, file, line, _ := runtime.Caller(2)
	file = filepath.Base(file)
	return fmt.Sprintf("%s:%d", file, line)
}

func Equal(t *testing.T, actual, expected interface{}) {
	if reflect.DeepEqual(actual, expected) {
		// Do Nothing while its went well.
	} else {
		a := reflect.ValueOf(actual)
		e := reflect.ValueOf(expected)

		av, at := format(a)
		ev, et := format(e)

		message := "\n"
		message += getInfo() + "\n"
		message += fmt.Sprintf("[actual]  :%s(%s)\n", av, at)
		message += fmt.Sprintf("[expected]:%s(%s)\n", ev, et)

		t.Error(message)
	}
}
