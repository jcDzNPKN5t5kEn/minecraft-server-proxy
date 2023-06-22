package API

import (
	"minecraft-proxy/Utils"
)

var whiteList []string

func InWhiteList(name string) bool {
	return Utils.ContainsInStringArray(name, whiteList)
}

func AddAllowedName(name string) {
	whiteList = append(whiteList, name)
}
func RemoveAllowedName(name string) {
	whiteList = Utils.RemoveFromStringArray(name, whiteList)
}
func GetAllowedName() string {
	allowedNameStr := ""
		for _, name := range whiteList {
			allowedNameStr += name + "\n"
		}
		return allowedNameStr
}
