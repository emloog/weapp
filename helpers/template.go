package helpers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//IsActive checks uri against currently active (uri, or nil) and returns "active" if they are equal
func IsActive(active interface{}, uri string) string {
	if s, ok := active.(string); ok {
		if s == uri {
			return "active"
		}
	}
	return ""
}

//DateTime prints timestamp in human format
func DateTime(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

//StringInSlice returns true if value is in list slice
func StringInSlice(value string, list []string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

//DefaultH returns common to all pages template data
func DefaultH(c *gin.Context) gin.H {
	user, _ := c.Get("User")
	return gin.H{
		"ActiveUser":    user,          //signed in models.User
		"Active":        "",            //active uri shortening for menu item highlight
		"Title":         "",            //page title:w
	}
}