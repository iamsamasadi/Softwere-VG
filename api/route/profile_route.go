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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
