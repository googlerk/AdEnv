package adenv

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetAdEnv(path string) (stt *AdEnv, err error) {
	strPath := pathOrDefaultPath(path)
	var res map[string]string
	res, err = godotenv.Read(strPath)
	if err != nil {
		log.Printf("\n godotenv %#v \n", err)
		return
	}
	jsonbody, err := json.Marshal(res)
	if err != nil {
		log.Printf("\n jsonbody %#v \n", err)
		return
	}
	if err = json.Unmarshal(jsonbody, &stt); err != nil {
		log.Printf("\n Unmarshal %#v \n", err)
		return
	}

	hostname, err := os.Hostname()
	constEnv := chkHostName(hostname)
	//fmt.Printf("\n\n constEnv = %s \n\n", constEnv)
	if constEnv == "" {
		return stt, errors.New("HostName : " + hostname + " not exist.")
	}
	stt.ConstEnv = constEnv
	return stt, err
}

func chkHostName(s string) (env string) {
	str := strings.ToLower(s)
	switch {
	case strings.Index(str, "prd") >= 0:
		return "Production"
	case strings.Index(str, "uat") >= 0:
		return "UAT"
	case strings.Index(str, "kb") >= 0:
		return "local"
	default:
		return ""
	}
}

func pathOrDefaultPath(path string) string {
	if len(path) == 0 {
		return constPathEnv
	}
	return path
}
