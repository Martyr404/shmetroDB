package psql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

// 配置结构体，与cfg.yml中的字段对应
type Config struct {
	Host     string `yaml:"host"`     // 数据库主机地址
	Port     int    `yaml:"port"`     // 数据库端口，默认5432
	User     string `yaml:"user"`     // 用户名
	Password string `yaml:"password"` // 密码
	DBName   string `yaml:"dbname"`   // 数据库名
	SSLMode  string `yaml:"sslmode"`  // SSL模式，如"require"、"disable"
}

var db *sql.DB // 全局数据库连接实例

// 读取配置文件
func loadConfig() (Config, error) {
	var cfg Config
	//abs_wd, _ := os.Getwd()
	abs_wd := "/Users/pigeonlx/Desktop/code/shmetroDB/api"
	data, err := os.ReadFile(abs_wd + "/psql/cfg.yml")
	if err != nil {
		return cfg, err
	}

	// 解析YAML配置
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	// 设置默认值（如果配置文件中未指定）
	if cfg.Port == 0 {
		cfg.Port = 5432
	}
	if cfg.SSLMode == "" {
		cfg.SSLMode = "require" // 生产环境建议使用require
	}

	return cfg, nil
}

// 初始化数据库连接
func Init() error {
	// 读取同目录下的cfg.yml配置
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	// 构建PostgreSQL连接字符串
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	// 打开数据库连接
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	// 验证连接是否有效
	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}

// 获取数据库连接实例
func GetDB() *sql.DB {
	return db
}

// 关闭数据库连接
func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
