/**
 * @author sd_sase
 * @version 1.0.0
 * @copyright www.sd-sase.com
 * @e-mail sd_sase@163.com
 * @2021/11/14
 */

package idoictos_mfa_service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	authentication = "http://0.0.0.0:21515"
	authorization  = "http://0.0.0.0:21118"
	daas           = "http://0.0.0.0:21522"
)

func Keepalived(c *gin.Context) {
	url := authorization + c.Request.RequestURI
	log.Println("Keepalived url:", url)
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	token := c.GetHeader("IDaaS_Token")
	request.Header.Set("IDaaS_Token", token)
	request.Header.Set("Content-Type", c.ContentType())
	request.Header.Set("User-Agent", c.Request.UserAgent())
	response, err := client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	// output, err := ioutil.ReadAll(reader)
	// var outputString = string(output)
	defer response.Body.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

}

func Get(c *gin.Context) {
	url := authorization + c.Request.RequestURI
	log.Println("request url:", url)
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	token := c.GetHeader("IDaaS_Token")
	request.Header.Set("IDaaS_Token", token)
	request.Header.Set("Content-Type", c.ContentType())
	request.Header.Set("User-Agent", c.Request.UserAgent())
	response, err := client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	// outputResponse, err := ioutil.ReadAll(reader)
	// var outputString = string(outputResponse)
	// log.Println("output:", outputString)
	defer response.Body.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func Post(c *gin.Context) {
	url := authorization + c.Request.RequestURI
	body := c.Request.Body
	client := http.Client{}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return
	}
	token := c.GetHeader("IDaaS_Token")
	request.Header.Set("IDaaS_Token", token)
	request.Header.Set("Content-Type", c.ContentType())
	request.Header.Set("User-Agent", c.Request.UserAgent())
	response, err := client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	defer response.Body.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func Put(c *gin.Context) {
	url := authorization + c.Request.RequestURI
	body := c.Request.Body
	client := http.Client{}
	request, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return
	}
	token := c.GetHeader("IDaaS_Token")
	request.Header.Set("IDaaS_Token", token)
	request.Header.Set("Content-Type", c.ContentType())
	request.Header.Set("User-Agent", c.Request.UserAgent())
	response, err := client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	defer response.Body.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func Patch(c *gin.Context) {
	url := authorization + c.Request.RequestURI
	body := c.Request.Body
	client := http.Client{}
	log.Println(url)
	request, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		return
	}
	token := c.GetHeader("IDaaS_Token")
	request.Header.Set("IDaaS_Token", token)
	request.Header.Set("Content-Type", c.ContentType())
	request.Header.Set("User-Agent", c.Request.UserAgent())
	response, err := client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	defer response.Body.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func Delete(c *gin.Context) {
	url := authorization + c.Request.RequestURI
	log.Println("request url:", url)
	client := http.Client{}
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}
	token := c.GetHeader("IDaaS_Token")
	request.Header.Set("IDaaS_Token", token)
	request.Header.Set("Content-Type", c.ContentType())
	request.Header.Set("User-Agent", c.Request.UserAgent())
	response, err := client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	reader := response.Body
	// outputResponse, err := ioutil.ReadAll(reader)
	// var outputString = string(outputResponse)
	// log.Println("output:", outputString)
	defer response.Body.Close()
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")
	extraHeaders := map[string]string{}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
