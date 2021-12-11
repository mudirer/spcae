package res

import (
	"fmt"
	"space/pkg/logutil"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"time"
)

type config struct {
	Server     Server
	DataBase   database
	Edgeservice edgeservice
	App App
}
type edgeservice struct {
	Ordersrv string
}


type database struct {
	Link string
	DSN    string
}
type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}



var (
	Config *config
)

func init() {
	viper.SetConfigName("config/config")
	//viper.SetConfigName("Desktop/workdir/space/config/config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config failed: %v", err)
		fmt.Println("config failed: %v", err)
	}

	viper.Unmarshal(&Config)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

	execpath, err := os.Executable()
	logutil.NewLogger(filepath.Join(filepath.Dir(execpath), "logs/blog.log"), "debug")

}
