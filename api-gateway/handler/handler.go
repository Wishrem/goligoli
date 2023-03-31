package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func SendBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Params Loss",
	})
}
