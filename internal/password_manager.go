package passwordmanager

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type PMEntry struct {
	Uuid     string
	Username string
	Password string
	Site     string
}

type PasswordManager interface {
	Add(entry *PMEntry) (err error)
	Update(entry *PMEntry) (err error)
	Delete(username string, site string) (err error)
	List() (users []*PMEntry, err error)
	FindEntries(username string, site string) (users []*PMEntry, err error)
}

type passwordManagerService struct {
}

func NewPasswordManagerService() PasswordManager {
	return &passwordManagerService{}
}

func (p *passwordManagerService) Add(entry *PMEntry) error {
	entry.Uuid = uuid.NewString()

	hashedPassword, err := Encrypt(entry.Password)
	if err != nil {
		return err
	}

	entry.Password = hashedPassword

	result := db.Create(entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *passwordManagerService) Update(entry *PMEntry) error {
	entry.Uuid = uuid.NewString()

	hashedPassword, err := Encrypt(entry.Password)
	if err != nil {
		return err
	}

	result := db.Model(&PMEntry{}).Where("username=? and site=?", entry.Username, entry.Site).Update("password", hashedPassword)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (p *passwordManagerService) Delete(username string, site string) error {
	result := db.Where("username=? and site=?", username, site).Delete(&PMEntry{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DecryptPasswords(users []*PMEntry) {
	for _, entry := range users {
		password, _ := Decrypt(entry.Password)
		entry.Password = password
	}
}

func (p *passwordManagerService) List() (users []*PMEntry, err error) {
	result := db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	DecryptPasswords(users)
	return users, nil
}

func (p *passwordManagerService) FindEntries(username string, site string) (users []*PMEntry, err error) {
	var result *gorm.DB
	if username == "" {
		result = db.Where("site=?", site).Find(&users)
	} else if site == "" {
		result = db.Where("username=?", username).Find(&users)
	} else {
		result = db.Where("username=? and site=?", username, site).Find(&users)
	}

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	DecryptPasswords(users)
	return users, nil
}

func init() {
	var err error

	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database: ", err)
	}

	err = db.AutoMigrate(&PMEntry{})
	if err != nil {
		log.Fatalln("failed to connect database: ", err)
	}
}
