package nutils

import (
	"os"
	"strings"
)

func load(filename string, dest map[string]string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	buff := make([]byte, fi.Size())
	f.Read(buff)

	f.Close()
	str := string(buff)
	if !strings.HasSuffix(str, "\n") {
		str = str + "\n"
	}
	s2 := strings.Split(str, "\n")

	for i := 0; i < len(s2); {

		if strings.HasPrefix(s2[i], "#") {
			i++
		} else if strings.Contains(s2[i], "=") {
			key := strings.Trim(s2[i][0:strings.Index(s2[i], "=")], " ")
			val := strings.Trim(s2[i][strings.Index(s2[i], "=")+1:len(s2[i])], " ")

			i++
			val = strings.Replace(val, "\r", "", -1)

			dest[key] = val
		} else {
			i++
		}
	}
	return nil
}

// GetConfigValue read configuration parameter from configuration file
func GetConfigValue(configFile, key, defaultValue string) string {
	mymap := make(map[string]string)

	var value string

	dir, _ := os.Getwd()
	if !strings.Contains(configFile, string(os.PathSeparator)) {
		configFile = dir + string(os.PathSeparator) + configFile
	}

	err := load(configFile, mymap)
	if err != nil {
		println(err.Error())
		return defaultValue
	} else if mymap[key] != "" {
		value = mymap[key]
		return value
	}
	return defaultValue
}

func GetOSArgs() map[string]string {
	args := os.Args[1:]
	argsMap := make(map[string]string)
	for i := 0; i < len(args)-1; i += 2 {
		argsMap[args[i]] = args[i+1]
	}
	return argsMap
}
