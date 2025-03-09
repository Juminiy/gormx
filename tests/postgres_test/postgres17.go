package postgres17

import (
	"fmt"
	"github.com/Juminiy/gormx/tests/db_decl"
	"github.com/Juminiy/kube/pkg/util"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"os"
)

var _cfg = struct {
	Host     string `yaml:"Host"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
	Port     int    `yaml:"Port"`
	SSLMode  string `yaml:"SSLMode"`
	TimeZone string `yaml:"TimeZone"`
}{}

func Orm() *gorm.DB {
	cfgPath, err := os.Open("testdata/env/postgres17.yaml")
	util.Must(err)
	cfgBytes, err := io.ReadAll(cfgPath)
	util.Must(err)
	util.Must(yaml.Unmarshal(cfgBytes, &_cfg))
	return db_decl.Orm(postgres.Open(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			_cfg.Host, _cfg.User, _cfg.Password, _cfg.DBName, _cfg.Port, _cfg.SSLMode, _cfg.TimeZone)))
}
