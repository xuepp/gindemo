package v1

import (
	"gindemo/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddMember(c *gin.Context) {

	res := entity.Result{}
	mem := entity.Member{}

	if err := c.ShouldBind(&mem); err != nil {
		res.SetCode(entity.CODE_ERROR)
		res.SetMessage(err.Error())
		c.JSON(http.StatusForbidden, res)
		c.Abort()
		return
	}

	// 处理业务(下次再分享)

	data := map[string]interface{}{
		"name": mem.Name,
		"age":  mem.Age,
	}
	res.SetCode(entity.CODE_ERROR)
	res.SetData(data)
	c.JSON(http.StatusOK, res)
}
