package mysql8

import (
	"fmt"
	"github.com/Juminiy/gormx/tests/db_decl"
	"github.com/Juminiy/kube/pkg/util"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
)

var _cfg = struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Protocol string `yaml:"Protocol"`
	Addr     string `yaml:"Addr"`
	DBName   string `yaml:"DBName"`
}{}

func Orm() *gorm.DB {
	cfgPath, err := os.Open("testdata/env/mysql8.yaml")
	util.Must(err)
	cfgBytes, err := io.ReadAll(cfgPath)
	util.Must(err)
	util.Must(yaml.Unmarshal(cfgBytes, &_cfg))
	return db_decl.Orm(mysql.Open(
		fmt.Sprintf("%s:%s@%s(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			_cfg.Username, _cfg.Password, _cfg.Protocol, _cfg.Addr, _cfg.DBName)))
}
