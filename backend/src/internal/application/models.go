package application 

import (
	"github.com/Taielmolina01/gin-nextjs-template/src/internal/domains/users/models"
)

func GetAllModels() []interface{} {
	return []interface{}{
		&models.UserDB{},
	}
}
