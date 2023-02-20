package main

import (
	"encoding/base64"
	"encoding/xml"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Pong struct {
	XMLName  xml.Name `xml:"http://evd.vd.ch/1000 response"`
	XmlnsMx  string   `xml:"mx,attr,omitempty"`
	XmlnsX14 string   `xml:"x14,attr,omitempty"`
	Result   string   `xml:"result"`
	Message  string   `xml:"message"`
}

func pong(ctx *gin.Context) {

	name := ctx.Param("name")

	if name != "" {

		response := Pong{
			XmlnsMx:  "http://mx",
			XmlnsX14: "http://x14",
			Result:   "pong " + name,
			Message:  "all ol",
		}
		ctx.XML(http.StatusOK, response)

		log.Println("pong with param " + name)

	} else {

		log.Println(ctx.Request.Header)

		autorisation := ctx.Request.Header.Get("Authorization")
		b64, found := strings.CutPrefix(autorisation, "Basic ")
		if found {
			plain, err := base64.RawStdEncoding.DecodeString(b64)
			if err != nil {
				log.Fatal(err)
			}

			userpass := strings.Split(string(plain), ":")
			user := userpass[0]
			pass := userpass[1]

			log.Println("username " + user + " password " + pass)

		}

		ctx.XML(http.StatusOK, gin.H{
			"message": "pong",
		})
	}

}

func setupRouter() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/ping", pong)
	r.GET("/ping/:name", pong)

	return r
}

func main() {

	r := setupRouter()
	r.Run()

}
