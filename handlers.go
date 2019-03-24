package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

// RegexpImage is the regexp for image names
var RegexpImage = regexp.MustCompile(`^\w+$`)

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
		ctx.Data(http.StatusOK, "text/html", []byte("Bad image size"))
		return
	}
	for i := 0; i < len(images); i++ {
		if !RegexpImage.MatchString(images[i]) {
			ctx.Data(http.StatusOK, "text/html", []byte("Bad image name"))
			return
		}
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
		ctx.Data(http.StatusOK, "text/html", []byte(output))
		return
	}
	data, err := ioutil.ReadFile(image)
	if err != nil {
		ctx.Data(http.StatusOK, "text/html", []byte(err.Error()))
		return
	}
	ctx.Data(http.StatusOK, "image/png", data)
}

// Planche4x2
func Planche4x2(ctx *gin.Context) {
	token := ctx.Param("token")
	user, err := CheckCredentials(token)
	if err != nil {
		ctx.Data(http.StatusOK, "text/html", []byte("Accès non permis"))
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
		ctx.Data(http.StatusOK, "text/html", []byte(err.Error()))
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
		ctx.Data(200, "text/html", []byte(err.Error()))
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
		ctx.Data(401, "text/html", []byte(err.Error()))
		return
	}
	logger.Printf("Planche pour %s", user.Name)
	image := fmt.Sprintf("/tmp/planche-%d.png", <-Index)
	var images []string
	files, err := ioutil.ReadDir("./")
	if err != nil {
		ctx.Data(http.StatusOK, "text/html", []byte(err.Error()))
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
	arguments := []string{"-mode", "concatenate", "-label", "%t", "-pointsize", "48", "-tile", "5x4"}
	for _, image := range images {
		arguments = append(arguments, image)
	}
	arguments = append(arguments, image)
	output, err := RunCommand(command, arguments)
	if err != nil {
		ctx.Data(http.StatusOK, "text/html", []byte(output))
		return
	}
	data, err := ioutil.ReadFile(image)
	if err != nil {
		ctx.Data(http.StatusOK, "text/html", []byte(err.Error()))
		return
	}
	ctx.Data(http.StatusOK, "image/png", data)
}
