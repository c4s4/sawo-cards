package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

// BuildImage run command to build the image
func BuildImage(images []string, ctx *gin.Context) {
	image := fmt.Sprintf("/tmp/planche-%d.png", <-Index)
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
	for i := 0; i < len(images); i++ {
		images[i] = images[i] + ".png"
	}
	command := "montage"
	arguments := []string{"-mode", "concatenate", "-tile", size}
	for _, image := range images {
		arguments = append(arguments, image)
	}
	arguments = append(arguments, image)
	output, err := RunCommand(command, arguments)
	if err != nil {
		ctx.AbortWithStatusJSON(400, output)
		return
	}
	data, err := ioutil.ReadFile(image)
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

// Planche
func Planche(ctx *gin.Context) {
	token := ctx.Param("token")
	user, err := CheckCredentials(token)
	if err != nil {
		ctx.AbortWithError(401, err)
		return
	}
	logger.Printf("Planche pour %s", user.Name)
	image := fmt.Sprintf("/tmp/planche-%d.png", <-Index)
	var images []string
	files, err := ioutil.ReadDir("./")
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".png") {
			images = append(images, name)
		}
	}
	sort.Strings(images)
	command := "montage"
	arguments := []string{"-mode", "concatenate", "-label", "%t", "-pointsize", "48", "-tile", "5x3"}
	for _, image := range images {
		arguments = append(arguments, image)
	}
	arguments = append(arguments, image)
	output, err := RunCommand(command, arguments)
	if err != nil {
		ctx.AbortWithError(400, fmt.Errorf(output))
		return
	}
	data, err := ioutil.ReadFile(image)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	ctx.Data(http.StatusOK, "image/png", data)
}
