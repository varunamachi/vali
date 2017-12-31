package vnet

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/varunamachi/vaali/vlog"
)

//GetOffsetLimit - retrieves offset and limit as integers provided in URL as
//query parameters. These parameters should have name - offset and limit
//respectively
func GetOffsetLimit(ctx echo.Context) (offset, limit int, has bool) {
	has = false
	offset = 0
	limit = 0
	strOffset := ctx.QueryParam("offset")
	strLimit := ctx.QueryParam("limit")
	if len(strOffset) == 0 || len(strLimit) == 0 {

		has = false
		return
	}
	var err error
	offset, err = strconv.Atoi(strOffset)
	if err != nil {
		offset = 0
		return
	}
	limit, err = strconv.Atoi(strLimit)
	if err != nil {
		limit = 0
		return
	}
	has = true
	return
}

//DefMS - gives default message and status, used for convenience
func DefMS(oprn string) (int, string) {
	return http.StatusOK, oprn + " - successful"
}

//GetUserID - retrieves user ID from context
func GetUserID(ctx echo.Context) string {
	ui := ctx.Get("userID")
	userID, ok := ui.(string)
	if ok {
		return userID
	}
	return ""
}

//AuditedSend - sends result as JSON while logging it as event. The event data
//is same as the data present in the result
func AuditedSend(ctx echo.Context, res *Result) (err error) {
	err = ctx.JSON(res.Status, res)
	vlog.LogEvent(res.Op, GetUserID(ctx), res.OK, res.Err, res.Data)
	return err
}

//AuditedSendX - sends result as JSON while logging it as event. This method
//logs event data which is seperate from result data
func AuditedSendX(ctx echo.Context,
	data interface{},
	res *Result) (err error) {
	err = ctx.JSON(res.Status, res)
	vlog.LogEvent(res.Op, GetUserID(ctx), res.OK, res.Err, data)
	return err
}