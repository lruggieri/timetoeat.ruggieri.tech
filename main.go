package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	internalError = "internal error, sorry for the inconvenience"
)

var (
	logger = logrus.New()
)

func main() {
	r := gin.Default()
	r.POST("/", handlerGeneral)
	port := "5001"
	fmt.Printf("running server on port %s", port)
	panic(r.Run(fmt.Sprintf(":%s", port)))
}

func handlerGeneral(c *gin.Context){
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, getFormattedError(internalError))
		return
	}
	urlParsed,err := url.Parse(string(body))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, getFormattedError("invalid request body"))
		return
	}
	urlParts, _ := url.ParseQuery(urlParsed.Path)

	c.JSON(200,
		generalResponse{
			ResponseType: "in_channel",
			Text: generateResponse(urlParts.Get("user_name")),
		},
	)
}

type generalResponse struct {
	ResponseType string `json:"response_type"`
	Text string `json:"text"`
}

func generateResponse(userName string) string{
	switch{
	case strings.Contains(userName,"abdel"):
		return "Abdel...just go to Saveur D'Asie and grab some sushi..."
	default:
		return fmt.Sprintf("grab something at %s", randomPlaceGenerator())
	}
}

func randomPlaceGenerator() string {
	places := []string{
		"Saveur D'Asie",
		"Libanese",
		"a random restaurant",
		"that nice creperie",
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return places[r1.Intn(len(places))]
}

func getFormattedError(error string) string{
	return `{"error": "`+error+`"}`
}