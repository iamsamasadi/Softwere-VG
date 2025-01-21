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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
