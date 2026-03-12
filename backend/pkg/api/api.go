package api

import (
	"backend/pkg/db_query"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	TrafficKey = "X-Request-Id"
)

type Api struct {
	Context *gin.Context
	Logger  *zap.SugaredLogger
	Errors  error
}

// generateMsgIDFromContext 生成msgID
func generateMsgIDFromContext(c *gin.Context) string {
	requestId := c.GetHeader(TrafficKey)
	if requestId == "" {
		requestId = uuid.New().String()
		c.Header(TrafficKey, requestId)
	}
	return requestId
}

func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	return e
}

func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		// 使用 ShouldBind，不进行验证
		err = e.Context.ShouldBind(d)
	} else {
		for i := range bindings {
			// 直接绑定，不验证
			if bindings[i] == binding.Query {
				err = e.Context.ShouldBindQuery(d)
			} else if bindings[i] == binding.JSON {
				err = e.Context.ShouldBindJSON(d)
			} else {
				err = e.Context.ShouldBindWith(d, bindings[i])
			}
			if err != nil {
				break
			}
		}
	}
	if err != nil {
		e.Errors = err
	}
	return e
}

func (e *Api) Error(code int, err error, msg string) {
	e.Context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func (e *Api) ErrorApi(err error) {
	e.Context.JSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"msg":  err.Error(),
	})
}

type DetailResponse struct {
	RequestId string      `json:"requestId"`
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
}

func (e *Api) OK(data interface{}, msg string) {
	e.Context.JSON(http.StatusOK, DetailResponse{
		RequestId: generateMsgIDFromContext(e.Context),
		Code:      200,
		Data:      data,
	})
}

type PageResponse struct {
	Code int      `json:"code"`
	Data PageData `json:"data"`
}

type PageData struct {
	Total     int64       `json:"total"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	List      interface{} `json:"list"`
}

func (e *Api) PageOK(pagination *db_query.PaginationQ) {
	e.Context.JSON(http.StatusOK, PageResponse{
		Code: 200,
		Data: PageData{
			Total:     pagination.Total,
			PageIndex: pagination.Page,
			PageSize:  pagination.Size,
			List:      pagination.Data,
		},
	})
}
