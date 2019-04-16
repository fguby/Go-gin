package v1

import (
	"Gin_demo/models"
	cs "Gin_demo/pkg/constant"
	"Gin_demo/pkg/util"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

var (
	ReturnResult = func(c *gin.Context, code int, msg string) {
		c.JSON(code, msg)
	}
)

// @Summary 查询文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/ [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	// data := make(map[string]interface{})
	var data util.Page
	var list interface{}
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	code := cs.INVALID_PARAMS
	if !valid.HasErrors() {
		code = cs.SUCCESS
		//获取数据总数
		count, err := models.GetTagTotal(maps)
		if err != nil {
			ReturnResult(c, cs.ERROR, cs.GetMsg(cs.ERROR))
		}
		// data["count"] = count
		//默认取第一页1-10条记录
		var pageSize int = 10
		var pageNo int = 1
		if pageSizeArg := c.Query("pageSize"); pageSizeArg != "" {
			pageSize = com.StrTo(pageSizeArg).MustInt()
		}
		if pageNoArg := c.Query("pageNo"); pageNoArg != "" {
			pageNo = com.StrTo(pageNoArg).MustInt()
		}
		var err1 error
		list, err1 = models.GetTags(pageNo, pageSize, maps)
		if err1 != nil {
			ReturnResult(c, cs.ERROR, cs.GetMsg(cs.ERROR))
		}
		data = util.GetPage(count, pageNo, pageSize, list)
	}
	c.JSON(cs.SUCCESS, gin.H{
		"code": code,
		"msg":  cs.GetMsg(code),
		"data": data,
	})
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	//var data interface{}
	var b models.Tag
	c.Bind(&b)
	if err := models.AddTag(&b); err != nil {
		c.JSON(cs.ERROR, cs.GetMsg(cs.ERROR))
	}
	c.JSON(cs.SUCCESS, cs.GetMsg(cs.SUCCESS))
}

// @Summary 修改文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	var b models.Tag
	b.Base_model.Id = id
	b.Name = c.Query("name")
	b.State = com.StrTo(c.Query("state")).MustInt()
	b.ModifiedBy = c.Query("modified_by")
	if err := models.UpdateTag(id, b); err != nil {
		ReturnResult(c, cs.ERROR, "修改数据失败")
	}
	c.JSON(cs.SUCCESS, "修改成功")
}

// @Summary 删除文章标签
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if err := models.DeleteTagById(id); err != nil {
		ReturnResult(c, cs.ERROR, "删除数据失败")
	}
	c.JSON(cs.SUCCESS, "删除成功")
}
