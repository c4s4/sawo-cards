package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

// Configuration est la configuration de l'application
type Configuration struct {
	Port  int
	Users map[string]User
}

// User est la configuration d'un utilisateur
type User struct {
	Email string
	Name  string
}

// configuration est la configuration de l'application
var configuration Configuration

// logger est le logger
var logger = log.New(os.Stdout, "", log.Ldate+log.Ltime)

// Index est un channel pour distribuer les index des images
var Index = make(chan int, 1)

// init lance la goroutine de l'index
func init() {
	go Indexer()
}

// Error affiche un message d'erreur et quitte en erreur
func Error(message string) {
	println(message)
	os.Exit(1)
}

// RunCommand exécute une commande de manière synchrone
func RunCommand(command string, params []string) (string, error) {
	cmd := exec.Command(command, params...)
	bytes, err := cmd.CombinedOutput()
	return string(bytes), err
}

// LoadConfiguration charge la configuration à partir du fichier passé en paramètre
func LoadConfiguration() {
	if len(os.Args) != 2 {
		Error("You must pass configuration file on command line")
	}
	source, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		Error(fmt.Sprintf("Error loading configuration file: %v", err))
	}
	yaml.UnmarshalStrict(source, &configuration)
}

// Indexer est appelé dans une goroutine pour générer les indexs des publications
func Indexer() {
	i := 1
	for {
		Index <- i
		i++
	}
}

// CheckCredentials looks for token un configuration.Users map
// Return authenticated user or error if access is not granted
func CheckCredentials(token string) (*User, error) {
	user, found := configuration.Users[token]
	if !found {
		logger.Printf("Unauthorized access: '%s'", token)
		return nil, fmt.Errorf("Invalid token '%s'", token)
	}
	return &user, nil
}

// SetupRouter set routes
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/:token/:img1/:img2/:img3/:img4/:img5/:img6/:img7/:img8", Planche4x2)
	router.GET("/:token/:img1/:img2/:img3/:img4", Planche2x2)
	router.GET("/:token/:img1/:img2", Planche2x1)
	router.GET("/:token", Planche)
	return router
}

func main() {
	LoadConfiguration()
	router := SetupRouter()
	router.Run(fmt.Sprintf(":%d", configuration.Port))
}
