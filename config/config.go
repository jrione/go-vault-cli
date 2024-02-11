package config

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Vault struct {
		Token    string `json:"token"`
		Port     string `json:"port"`
		Url      string `json:"url"`
		Protocol string `json:"protocol"`
	} `json:"vault"`
}

var Env *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Env); err != nil {
		panic(err)
	}
}

func LoadExternalEnv() {
	os.Clearenv()
	godotenv.Load(".env")

}

func GetExternalEnv() map[string]any {
	LoadExternalEnv()
	newEnv := make(map[string]any)
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			newEnv[e[:i]] = e[i+1:]
		}
	}
	return newEnv
}

var (
	filename string
	help     string
	APP_NAME string
	BRANCH   string
)

func RunFlag() {
	flag.StringVar(&filename, "file", "", "File location want to save")
	flag.StringVar(&APP_NAME, "app", "", "Application Name")
	flag.StringVar(&BRANCH, "branch", "", "Branch Location")
	flag.StringVar(&help, "help", "", "Format: go run main.go [flags] [cmd]")
	flag.Parse()

	cmd := RunArgs()
	switch cmd {
	case "save":
		isExist := fileCheck(filename)
		if !isExist {
			log.Fatalf("File: %v not found\n", filename)
			return
		}
	}

	if APP_NAME == "" || BRANCH == "" {
		log.Fatal("Flags -app -branch Required")
		return
	}
	os.Setenv("APP_NAME", APP_NAME)
	os.Setenv("BRANCH", BRANCH)

}

func fileCheck(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err) //true=exist
}

func RunArgs() string {
	cmd := os.Args[len(os.Args)-1]
	return cmd
}

func ToMethod(cmd string) string {
	cmd = cases.Title(language.English, cases.Compact).String(strings.ToLower(cmd))
	return cmd
}
