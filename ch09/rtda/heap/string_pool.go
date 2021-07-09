package heap

import (
	"unicode/utf16"
)

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUtf16(goStr)
	//goStr ==> to char[]数组
	jChars := &Object{
		class: loader.LoadClass("[C"),
		data:  chars,
	}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr
}

func stringToUtf16(s string) []uint16 {
	runes := []rune(s)
	return utf16.Encode(runes)
}

func JStrToGoStr(ref *Object) string {
	field := ref.GetRefVar("value", "[C")
	return utf16ToString(field.Chars())
}

func utf16ToString(chars []uint16) string {
	runes := utf16.Decode(chars)
	return string(runes)
}

func InternString(jStr *Object) *Object {
	goStr := JStrToGoStr(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	internedStrings[goStr] = jStr
	return jStr
}
