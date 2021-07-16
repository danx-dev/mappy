package mappy

import (
	"errors"
	"log"
	"reflect"
)

func DoMap(dest interface{}, source interface{}) interface{} {
	log.Println("DoMap Start")

	v := reflect.ValueOf(source)
	if v.Kind() != reflect.Ptr {
		log.Fatalln("non-pointer ", v.Type())
		return nil
	}
	t := reflect.TypeOf(source)
	s := v.Elem()
	mapList := make(map[string]reflect.Value)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		mapField := t.Elem().Field(i).Tag.Get("mappy")
		if mapField != "" {
			mapList[mapField] = f
		}
	}
	log.Println("Input done")
	return setOutput(dest, mapList)

}
func setOutput(dest interface{}, mapList map[string]reflect.Value) interface{} {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		log.Fatalln("setOut non-pointer ", v.Type())
		return nil
	}
	t := reflect.TypeOf(dest).Elem()
	d := v.Elem()
	for i := 0; i < d.NumField(); i++ {
		f := d.Field(i)
		if f.Kind() == reflect.Struct {
			setValueStuct(&f, t.Field(i), mapList)
		}
		if val, ok := mapList[t.Field(i).Name]; ok {
			err := setValue(&f, val)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return dest
}

func setValueStuct(f *reflect.Value, field reflect.StructField, mapList map[string]reflect.Value) {
	t := field.Type
	for i := 0; i < t.NumField(); i++ {
		path := field.Name + "." + t.Field(i).Name
		if val, ok := mapList[path]; ok {
			subF := f.FieldByName(t.Field(i).Name)
			err := setValue(&subF, val)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}

func setValue(f *reflect.Value, v reflect.Value) error {
	switch f.Kind() {
	case reflect.Int:
		f.SetInt(v.Int())
	case reflect.String:
		f.SetString(v.String())
	case reflect.Bool:
		f.SetBool(v.Bool())
	case reflect.Struct:
		//Skip
	default:
		return errors.New("Unsupported kind: " + f.Kind().String())

	}

	return nil
}
