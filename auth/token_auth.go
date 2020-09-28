package auth

import "github.com/jinzhu/gorm"

//Token ... struct to store a boolean value
type Token struct {
	ID string `json:"id" gorm:"primary_key"`
}

type tokenRepo struct {
	db *gorm.DB
}

type tokenRepoInterface interface {
	CreateToken(token Token) error
	ExistToken(id string) bool
}

//TokenRepo ...
var (
	TokenRepo tokenRepoInterface
)

//SetTokenRepo ...
func SetTokenRepo(db *gorm.DB) {
	TokenRepo = &tokenRepo{db: db}
}

func (tr *tokenRepo) CreateToken(token Token) error {
	err := tr.db.Create(&token).Error
	return err
}

func (tr *tokenRepo) ExistToken(id string) bool {
	token := Token{}
	err := tr.db.Where(`id = ?`, id).Find(&token).Error
	if err != nil {
		return false
	}
	return true
}
