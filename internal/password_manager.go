package passwordmanager

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type PMUser struct {
	Uuid     string
	Username string
	Password string
	Platform string
}

type PasswordManager interface {
	Add(user *PMUser) (err error)
}

type passwordManagerService struct {
}

func NewPasswordManagerService() PasswordManager {
	return &passwordManagerService{}
}

func (p *passwordManagerService) Add(user *PMUser) error {
	user.Uuid = uuid.NewString()
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func init() {
	var err error

	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database: ", err)
	}

	err = db.AutoMigrate(&PMUser{})
	if err != nil {
		log.Fatalln("failed to connect database: ", err)
	}
}
