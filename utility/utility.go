package utility

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jwalton/go-supportscolor"
)

type tsHttp struct {
	HttpPort                          string `json:"http_port"`
}

type tsDatabase struct {
	Hostname     string `json:"hostname"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}

type Configuration struct {
	Http     tsHttp     `json:"http"`
	Database tsDatabase `json:"database"`

	AppPath string `json:"app_path"`
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func LoadApplicationConfiguration(removeSuffixPath string) (config Configuration, err error) {
	defer RecoverError()

	config.AppPath, err = os.Getwd()
	if err != nil {
		return
	}

	if removeSuffixPath != "" {
		config.AppPath = strings.TrimSuffix(config.AppPath, removeSuffixPath)
	}

	configPath := filepath.Join(config.AppPath, "config.json")
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	if err != nil {
		return
	} else {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &config)
		return
	}
}

func RecoverError() {
	if r := recover(); r != nil {
		PrintConsole(fmt.Sprintf("[ERROR][RECOVER]=> %v", r), "error")
	}
}

func PrintConsole(strPrint string, strStatus string) {
	defer RecoverError()

	if supportscolor.Stdout().SupportsColor {
		if strings.ToLower(strings.TrimSpace(strStatus)) == "info" {
			fmt.Println(Green + "[INFO]> " + strPrint + Reset)
		} else if strings.ToLower(strings.TrimSpace(strStatus)) == "error" {
			fmt.Println(Red + "[ERROR]> " + strPrint + Reset)
		} else if strings.ToLower(strings.TrimSpace(strStatus)) == "warning" {
			fmt.Println(Yellow + "[WARNING]> " + strPrint + Reset)
		} else if strings.ToLower(strings.TrimSpace(strStatus)) == "logo" {
			fmt.Println(Red + strPrint + Reset)
		} else {
			fmt.Println(Reset + strPrint + Reset)
		}
	} else {
		if strings.ToLower(strings.TrimSpace(strStatus)) == "info" {
			fmt.Println("[INFO]> " + strPrint)
		} else if strings.ToLower(strings.TrimSpace(strStatus)) == "error" {
			fmt.Println("[ERROR]> " + strPrint)
		} else if strings.ToLower(strings.TrimSpace(strStatus)) == "warning" {
			fmt.Println("[WARNING]> " + strPrint)
		} else if strings.ToLower(strings.TrimSpace(strStatus)) == "logo" {
			fmt.Println(strPrint)
		} else {
			fmt.Println(strPrint)
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, X-Auth-Token, Content-Type, Content-Length, Authorization, Access-Control-Allow-Headers, Accept, Access-Control-Allow-Methods, Access-Control-Allow-Origin, Access-Control-Allow-Credentials")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, HEAD, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}