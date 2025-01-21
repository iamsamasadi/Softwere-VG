package route

import (
	"time"

	"github.com/iamsamasadi/Softwere-VG/api/controller"
	"github.com/iamsamasadi/Softwere-VG/bootstrap"
	"github.com/iamsamasadi/Softwere-VG/domain"
	"github.com/iamsamasadi/Softwere-VG/mongo"
	"github.com/iamsamasadi/Softwere-VG/repository"
	"github.com/iamsamasadi/Softwere-VG/usecase"
	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
