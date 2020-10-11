package service

import (
	"flag"
	"riji/config"
	"riji/dao"
	"riji/model"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

const (
	BingBackImgUrl = "https://www.cn.bing.com/"

	// error
	DBError           = 100001 // 数据库操作失败
	OtherHttpError    = 100002 // 第三方请求失败
	CodeError         = 100003 // json解析失败
	CosError          = 100004 // cos失败
	ParamError        = 100005 // 参数错误
	NormalError       = 100006 // 通常错误
	UserNotFound      = 100007 // 用户不存在
	UserPasswordError = 100008 //用户密码错误
	RecordNotFound    = 100009 // 记录不存在
)

var configPath string = "config"

func init() {
	flag.StringVar(&configPath, "conf", configPath, "server config path")
}

type RijiServer struct {
	dao  *dao.Dao
	conf *config.Config
}

func NewServer() (*RijiServer, error) {
	// 加载配置
	conf, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	// 连接数据库
	dao, err := dao.NewDao(conf)
	if err != nil {
		return nil, err
	}

	// 建表
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		db.SingularTable(true)
		return "t_" + defaultTableName
	}
	if !dao.Db.HasTable(&model.BackImgCos{}) {
		dao.Db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").CreateTable(&model.BackImgCos{})
	}
	if !dao.Db.HasTable(&model.MessageBoard{}) {
		dao.Db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").CreateTable(&model.MessageBoard{})
	}
	if !dao.Db.HasTable(&model.UserUploadFile{}) {
		dao.Db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").CreateTable(&model.UserUploadFile{})
	}
	if !dao.Db.HasTable(&model.User{}) {
		dao.Db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").CreateTable(&model.User{})
	}
	if !dao.Db.HasTable(&model.WordBook{}) {
		dao.Db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").CreateTable(&model.WordBook{})
	}
	if !dao.Db.HasTable(&model.Blog{}) {
		dao.Db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").CreateTable(&model.Blog{})
	}

	return &RijiServer{
		dao:  dao,
		conf: conf,
	}, nil
}

func LoadConfig() (*config.Config, error) {
	var (
		conf config.Config
	)

	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig() // 读取配置数据
	if err != nil {
		return nil, err
	}
	viper.Unmarshal(&conf) // 将配置信息绑定到结构体上

	return &conf, nil
}
