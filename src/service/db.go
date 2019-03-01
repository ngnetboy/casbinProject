package service

import (
	"github.com/tredoe/osutil/user/crypt/sha512_crypt"
	"model"

	log "github.com/Sirupsen/logrus"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

//var casbin *gormadapter.Adapter
var enForcer *casbin.Enforcer

func initAccount(name, role string) {
	account := &model.Account{}
	if err := db.Where("name=?", name).First(account).Error; err != nil {
		password, _ := sha512_crypt.New().Generate([]byte("123456"), nil)
		account.Name = name
		account.Role = role
		account.Password = password

		db.Create(account)
	}
}

func initTables() {
	initAccount("admin", "admin")
	initAccount("anonymous", "anonymous")
}

func ConnecDB() {
	var err error
	db, err = gorm.Open("postgres", model.Conf.DB)
	if err != nil {
		log.Fatalln("Open database failed:", err.Error())
	}

	if err = db.AutoMigrate(model.Models...).Error; nil != err {
		log.Fatalln("auto migrate tables failed:", err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)

	initTables()

	//初始化casbin的gorm适配器
	adapter := gormadapter.NewAdapterByDB(db)
	enForcer = casbin.NewEnforcer("D:/code/golang/casbinProject/src/conf/restModel.conf", adapter)
	enForcer.AddPolicy("admin", "/*", "(GET)|(POST)|(DELETE)")
	enForcer.LoadPolicy()
}

func CheckPolicy(role, path, method string) bool {
	res, err := enForcer.EnforceSafe(role, path, method)
	if err != nil {
		log.Debugln("check policy:", err.Error())
		return false
	}
	if res {
		return true
	} else {
		return false
	}
}

func DisconnectDB() {
	if db != nil {
		if err := db.Close(); nil != err {
			log.Fatalln("Disconnect from database failed:", err.Error())
		}
	}
}
