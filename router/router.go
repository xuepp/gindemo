package router

import (
	"fmt"
	"gindemo/common"
	v1 "gindemo/controller/v1"
	v2 "gindemo/controller/v2"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/test", Test)

	r.GET("/sn", SignDemo)

	// v1 版本
	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/product/add", v1.AddProduct)
		GroupV1.Any("/member/add", v1.AddMember)
	}

	// v2 版本
	GroupV2 := r.Group("/v2", common.VerifySign)
	{
		GroupV2.Any("/product/add", v2.AddProduct)
		GroupV2.Any("/member/add", v2.AddMember)
	}
}

func SignDemo(c *gin.Context) {
	fmt.Println("SignDemo")
	ts := strconv.FormatInt(common.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"name":  []string{"a"},
		"price": []string{"10"},
		"ts":    []string{ts},
	}
	res["sn"] = common.CreateSign(params)
	res["ts"] = ts
	common.RetJson("200", "", res, c)
}

func Test(c *gin.Context) {
	fmt.Println("Test66666")
}
