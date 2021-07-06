package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"double":  "D",
	"float":   "F",
}

func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

func toDescriptor(name string) string {
	if name[0] == '[' {
		return name
	}

	if d, ok := primitiveTypes[name]; ok {
		return d
	}

	return "L" + name + ";"
}

func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}

	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}

	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)

}
