package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BuildImage run command to build the image
func BuildImage(images []string, ctx *gin.Context) {
	var size string
	if len(images) == 8 {
		size = "4x2"
	} else if len(images) == 4 {
		size = "2x2"
	} else if len(images) == 2 {
		size = "2x1"
	} else {
		ctx.AbortWithError(400, fmt.Errorf("Bad image size"))
		return
	}
	for i:=0; i<len(images); i++ {
		images[i] = images[i] + ".png"
	}
	command := "montage"
	arguments := []string{"-mode", "concatenate", "-tile", size}
	for _, image := range images {
		arguments = append(arguments, image)
	}
	arguments = append(arguments, "/tmp/planche.png")
	output, err := RunCommand(command, arguments)
	if err != nil {
		ctx.AbortWithError(400, fmt.Errorf(output))
		return
	}
	data, err := ioutil.ReadFile("/tmp/planche.png")
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.Data(http.StatusOK, "image/png", data)
}

// Planche4x2
func Planche4x2(ctx *gin.Context) {
	token := ctx.Param("token")
	user, err := CheckCredentials(token)
	if err != nil {
		ctx.AbortWithError(401, err)
		return
	}
	images := []string{ctx.Param("img1"), ctx.Param("img2"), ctx.Param("img3"), ctx.Param("img4"),
		ctx.Param("img5"), ctx.Param("img6"), ctx.Param("img7"), ctx.Param("img8")}
	logger.Printf("Planche pour %s: %v", user.Name, images)
	BuildImage(images, ctx)
}

// Planche2x2
func Planche2x2(ctx *gin.Context) {
	token := ctx.Param("token")
	user, err := CheckCredentials(token)
	if err != nil {
		ctx.AbortWithError(401, err)
		return
	}
	images := []string{ctx.Param("img1"), ctx.Param("img2"), ctx.Param("img3"), ctx.Param("img4")}
	logger.Printf("Planche pour %s: %v", user.Name, images)
	BuildImage(images, ctx)
}

// Planche2x1
func Planche2x1(ctx *gin.Context) {
	token := ctx.Param("token")
	user, err := CheckCredentials(token)
	if err != nil {
		ctx.AbortWithError(401, err)
		return
	}
	images := []string{ctx.Param("img1"), ctx.Param("img2")}
	logger.Printf("Planche pour %s: %v", user.Name, images)
	BuildImage(images, ctx)
}
