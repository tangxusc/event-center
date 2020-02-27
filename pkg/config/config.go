package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const debugArgName = "debug"

func InitLog() {
	if viper.GetBool(debugArgName) {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetReportCaller(true)
		logrus.Debug("已开启debug模式...")
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	Instance.Debug = viper.GetBool(debugArgName)
}

func BindParameter(cmd *cobra.Command) {
	viper.SetEnvPrefix("event")
	viper.AutomaticEnv()

	cmd.PersistentFlags().BoolVarP(&Instance.Debug, debugArgName, "v", false, "debug mod")
	cmd.PersistentFlags().UintVarP(&Instance.Server.Duration, "server-duration", "d", 60, "间隔时间(分)")

	cmd.PersistentFlags().StringVarP(&Instance.Mongo.Address, "mongo-address", "", "localhost", "mongo数据库连接地址")
	cmd.PersistentFlags().StringVarP(&Instance.Mongo.Port, "mongo-port", "", "27017", "mongo数据库端口")
	cmd.PersistentFlags().StringVarP(&Instance.Mongo.Username, "mongo-Username", "", "root", "数据库用户名")
	cmd.PersistentFlags().StringVarP(&Instance.Mongo.Password, "mongo-Password", "", "123456", "数据库密码")
	cmd.PersistentFlags().IntVarP(&Instance.Mongo.LocalThreshold, "mongo-LocalThreshold", "", 3, "本地阀值")
	cmd.PersistentFlags().IntVarP(&Instance.Mongo.MaxPoolSize, "mongo-MaxPoolSize", "", 10, "最大连接数")
	cmd.PersistentFlags().IntVarP(&Instance.Mongo.MaxConnIdleTime, "mongo-MaxConnIdleTime", "", 5, "最大等待时间")
	cmd.PersistentFlags().StringVarP(&Instance.Mongo.DbName, "mongo-DbName", "", "events", "mongo数据库名称")

	//_ = viper.BindPFlag(debugArgName, cmd.PersistentFlags().Lookup(debugArgName))
	//
	//_ = viper.BindPFlag("mongo-address", cmd.PersistentFlags().Lookup("mongo-address"))
	//_ = viper.BindPFlag("mongo-port", cmd.PersistentFlags().Lookup("mongo-port"))
	//_ = viper.BindPFlag("mongo-Username", cmd.PersistentFlags().Lookup("mongo-Username"))
	//_ = viper.BindPFlag("mongo-Password", cmd.PersistentFlags().Lookup("mongo-Password"))
	//_ = viper.BindPFlag("mongo-LocalThreshold", cmd.PersistentFlags().Lookup("mongo-LocalThreshold"))
	//_ = viper.BindPFlag("mongo-MaxPoolSize", cmd.PersistentFlags().Lookup("mongo-MaxPoolSize"))
	//_ = viper.BindPFlag("mongo-MaxConnIdleTime", cmd.PersistentFlags().Lookup("mongo-MaxConnIdleTime"))
	//_ = viper.BindPFlag("mongo-DbName", cmd.PersistentFlags().Lookup("mongo-DbName"))
	//_ = viper.BindPFlag("mongo-EventCollectionName", cmd.PersistentFlags().Lookup("mongo-EventCollectionName"))
	//_ = viper.BindPFlag("mongo-SnapshotCollectionName", cmd.PersistentFlags().Lookup("mongo-SnapshotCollectionName"))
	_ = viper.BindPFlags(cmd.Flags())
}

type MongoConfig struct {
	Address  string
	Port     string
	Username string
	Password string

	LocalThreshold  int
	MaxPoolSize     int
	MaxConnIdleTime int
	DbName          string
}

type Config struct {
	Debug  bool
	Server *ServerConfig
	Mongo  *MongoConfig
}

type ServerConfig struct {
	Duration uint
}

var Instance = &Config{
	Server: &ServerConfig{},
	Mongo:  &MongoConfig{},
}
