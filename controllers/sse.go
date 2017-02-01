package controllers

import (
	"encoding/json"
	"fmt"
	"go-sse-rest/models"

	"gopkg.in/gin-gonic/gin.v1"
)

type (
	SseController struct{}
)

func NewSseController() *SseController {
	return &SseController{}
}

func (ctrl SseController) GetUser(ctx *gin.Context) {
	getUser := ctx.Param("user")
	fmt.Println("User:", getUser)

	testUser := models.Sse{
		UserID: "kevinadi",
		Val:    123.45,
		Tags:   []string{"tag1", "tag2", "tag3"},
	}

	userJSON, _ := json.Marshal(testUser)
	fmt.Fprintf(ctx.Writer, "%s", userJSON)
}
