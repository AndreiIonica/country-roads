package posthttp

import (
	"backend/roralis/jwt"
	"backend/roralis/post"
	httpresponse "backend/roralis/shared/http_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin controller for creating a post, needs jwt key
func (r *PostController) Create(c *gin.Context) {

	claimsRaw, exists := c.Get(r.tokenKey)

	if !exists {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: "JWT claims object is missing"})
		return
	}

	claims, ok := claimsRaw.(*jwt.JWTClaims)

	if !ok {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: "JWT claims is not of correct shape"})
		return
	}

	if claims.Role < 5 {
		c.JSON(http.StatusForbidden, httpresponse.Response{Message: "Your email is not verified or you don't have enough permissions"})
		return
	}

	var json post.Post

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, httpresponse.Response{Message: err.Error()})
		return
	}

	json.UserID = claims.ID

	err := r.repo.Create(&json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpresponse.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, json)

}
