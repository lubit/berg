package models

//import(
//	"flag"
//	"fmt"
//	"github.com/BurntSushi/toml"
//	"sync"
//)
//
////type MainConfig struct {
////	//总体行为控制
////	Main       mainConfig		`toml:"main"`
////	//Log文件位置以及设置级别
////	Log        logConfig		`toml:"log"`
////	//数据接收端 -- Kafka配置信息
////	Source      sourceConfig	`toml:"source"`
////	//数据写入端配置 -- MySQL\PgSQL\Kafka等
////	Target		targetConfig	`toml:"target"`
////	//失败后重试机制配置
////	Retry      retryConfig		`toml:"retry"`
////	//运行结构化消息配置
////	Record    recordConfig		`toml:"record"`
////}
//
//type recordConfig struct{
//	//总体处理量级统计日志
//	AuditDirectory string `toml:"audit_directory"`
//	//日志处理过程
//	LoseDirectory  string `toml:"lose_directory"`
//	//
//	//单个日志最大储存空间
//	MaxFileSize    int    `toml:"max_file_size"`
//	//日志名称前缀
//	Prefix         string `toml:"prefix"`
//}
//
//var g_cfg *MainConfig
//var cfg_once sync.Once
//
//func GetCfgInstance(filename string) *MainConfig {
//	cfg_once.Do(func() {
//		cfgFile := flag.String("c", filename, "config file")
//		flag.Parse()
//		if _, err := toml.DecodeFile(*cfgFile, &g_cfg); err != nil {
//			fmt.Println(err)
//			panic(err)
//		} else {
//			fmt.Printf("%s: %+v\n", filename, *g_cfg)
//		}
//	})
//	return g_cfg
//}