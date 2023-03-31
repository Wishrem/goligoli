package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wishrem/goligoli/erp"
	"github.com/wishrem/goligoli/pkg/util/jwt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func getDetails(st *status.Status) map[string]string {
	for _, d := range st.Details() {
		switch info := d.(type) {
		case *errdetails.QuotaFailure:
			dt := make(map[string]string)
			for _, item := range info.Violations {
				dt[item.Subject] = item.Description
			}
			return dt
		default:
			return nil
		}
	}
	return nil
}

func SendError(c *gin.Context, err error) {
	st := status.Convert(err)
	dt := getDetails(st)
	code := toHttpCode(st.Code())
	if len(dt) != 0 {
		c.JSON(code, gin.H{
			"error":   st.Message(),
			"details": dt,
		})
	} else {
		c.JSON(code, gin.H{
			"error": st.Message(),
		})
	}
}

/*
**********************

	New version

**********************
*/

func SendBadRequest(c *gin.Context) {
	c.JSON(erp.BadRequest.HttpCode(), erp.BadRequest)
}

func SendErrResp(c *gin.Context, err error) {
	erp := erp.Covert(err)
	c.JSON(erp.HttpCode(), erp)
}

func ParseToken(c *gin.Context) *jwt.Claims {
	token := new(Token)
	if err := c.ShouldBind(&token); err != nil {
		SendErrResp(c, erp.BadRequest)
		return nil
	}

	claims, err := jwt.Parse(token.SS)
	if err != nil {
		SendErrResp(c, erp.Unauthorized)
		return nil
	}

	return claims
}
