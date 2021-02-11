package users

import (
	// "encoding/json"
	// "io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twinemarron/bookstore_users-api/domain/users"
	"github.com/twinemarron/bookstore_users-api/services"
	"github.com/twinemarron/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO: Handle error
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(user)
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	// TODO: Handle json error
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// 上記処理と同じ事を以下で行う
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
