package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDataOfRequest(c *gin.Context) models.RequestData {

	start := time.Now()

	requestData := models.RequestData{}
	requestData.ClientIP = c.ClientIP()
	requestData.StatusCode = c.Writer.Status()
	requestData.UserAgent = c.Request.UserAgent()
	requestData.Endpoint = c.Request.RequestURI
	requestData.Host = c.Request.Host
	requestData.Method = c.Request.Method
	requestData.BytesSent = int64(c.Writer.Size())
	requestData.BytesReceived = c.Request.ContentLength

	end := time.Now()

	requestData.Timestamp = end
	latencyMs := time.Duration(end.Sub(start).Milliseconds())
	latencyFormatted := fmt.Sprintf("%.2f", float64(latencyMs)/1000.0)
	requestData.Latency = latencyFormatted

	return requestData
}

func LoggingDataOfRequest(data models.RequestData) (bool, error) {

	jsonData, err := json.Marshal(data)

	if err != nil {
		return false, err
	}

	url := "http://localhost:9832/v1/log"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	return true, nil

}
