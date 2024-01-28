package request

import (
	"github.com/benebobaa/harisenin-mini-project/domain/entity"
	"github.com/google/uuid"
)

type LoginRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequestDTO struct {
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func (this LoginRequestDTO) ToUserEntity() entity.User {
	return entity.User{
		Username: this.Username,
		Password: this.Password,
	}
}

func (this RegisterRequestDTO) ToUserEntity() entity.User {
	return entity.User{
		ID:       uuid.New(),
		Username: this.Username,
		Password: this.Password,
	}
}
