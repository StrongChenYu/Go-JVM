package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	// 分隔符是; 如果path中包含;
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	// */test/java.go 这种形式
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	// 具体到某个文件
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	//除了上面三种情况外就是路径形式的
	return newDirEntry(path)
}
