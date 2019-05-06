package closure

import "strings"

type closure struct {
}

// 单例模式。
func GetClosureInstance() closure {
	return closure{}
}

func (c closure) MakeSuffix(suffix string) func(string) string {

	return func(fileName string) string {
		if !(strings.HasSuffix(fileName, suffix)) {

			return fileName + suffix
		}
		return fileName
	}
}
