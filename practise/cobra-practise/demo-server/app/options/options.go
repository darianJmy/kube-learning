package options

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/pflag"
	cliflag "k8s.io/component-base/cli/flag"
	"kube-learning/practise/cobra-practise/demo-server/app/config"
	"log"
)

const (
	defaultConfigFile = "democonfig.yaml"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age string  `json:"age"`
}

type Options struct {
	ComponentConfig *config.Config
	// ConfigFile is the location of the autoscaler's configuration file.
	ConfigFile string

	Master string

	DB *gorm.DB

	Fs *pflag.FlagSet
}

func NewOptions() (*Options, error) {
	return &Options{
		Master: "demo-master",
		ComponentConfig: config.NewComponentConfig(),
	}, nil
}

func (o *Options) Complete() error {
	configFile := defaultConfigFile
	if len(o.ConfigFile) != 0 {
		configFile = o.ConfigFile
	}

	cfg, err := loadConfigFromFile(configFile)
	if err != nil {
		return err
	}
	o.ComponentConfig = cfg

	// 注册数据库链接池
	if err = o.registerDatabase(); err != nil {
		return err
	}

	return nil
}

func (o *Options) Run() error {

	router := gin.Default()
	// 这里很关键, 我们的 login.html 是写在当前目录的 templates 目录中的, 所以必须指定模板所在的目录
	// templates/* 表示从templates目录中加载模板文件
	router.LoadHTMLGlob("index.html")
	router.Any("/login", o.loginHandler)
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
	// 打印测试
	return nil
}

func (o *Options) loginHandler (context *gin.Context) {
	if context.Request.Method == "GET" {
		// 调用context.HTML 渲染模板
		// 状态码、模板名、参数( 用于渲染模板中的 {{}}, 这里我们没有使用模板语法, 所以传个 gin.H{} 即可 )
		context.HTML(200, `index.html`, nil)
	} else {
		// 如果不存在的话, 得到的是空字符串, 但是我们也可以设置默认值, 和Query是类似的
		var user User
		user.Name = context.PostForm("username")
		user.Age = context.PostForm("password")
		o.DB.Create(&user)
	}
}

func (o *Options) registerDatabase() error {
	sqlConfig := o.ComponentConfig.Mysql
	passwd, _ := o.Fs.GetString("password")
	if &passwd != nil {
		sqlConfig.Password = passwd
	}
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlConfig.User,
		sqlConfig.Password,
		sqlConfig.Host,
		sqlConfig.Port,
		sqlConfig.DBName)
	DB, err := gorm.Open("mysql", dbConnection)
	if err != nil {
		return err
	}

	// set the connect pools
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	o.DB = DB

	return nil
}

func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.ComponentConfig.Mysql.AddFlags(fss.FlagSet("mysql"))
	return fss
}