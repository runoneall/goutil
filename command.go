package goutil

import (
	"os"
)

func Command_Args() map[string]string {
	args := os.Args[1:]
	configMap := make(map[string]string)
	for index, item := range args {
		if index%2 == 0 && index+1 < len(args) {
			configMap[item] = args[index+1]
		}
	}
	return configMap
}

func Command_GetArg(arg string) string {
	args := Command_Args()
	value, exists := args[arg]
	if exists {
		return value
	}
	return ""
}
