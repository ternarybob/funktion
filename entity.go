package funktion

import "reflect"

func GetCollectionName(v interface{}) string {

	if typ, ok := v.(reflect.Type); ok {
		return GetTypeName(typ)
	}

	return GetElementName(v)
}

func GetEntityName(v interface{}) string {

	if typ, ok := v.(reflect.Type); ok {
		return GetTypeName(typ)
	}

	return GetElementName(v)
}

func GetElementName(v interface{}) string {

	rv := reflect.ValueOf(v)

	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	typ := rv.Type()

	return GetTypeName(typ)

}

func GetTypeName(typ reflect.Type) string {

	// for *Foo and **Foo the name we want to return is Foo
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return typ.Name()

}
