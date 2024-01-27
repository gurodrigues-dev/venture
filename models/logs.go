package models

import "time"

type RequestData struct {
	Timestamp     time.Time `json:"timestamp"`
	ClientIP      string    `json:"client_ip"`
	Method        string    `json:"method"`
	StatusCode    int       `json:"status_code"`
	Endpoint      string    `json:"endpoint"`
	Host          string    `json:"host"`
	Latency       string    `json:"latency"`
	UserAgent     string    `json:"user_agent"`
	BytesSent     int64     `json:"bytes_sent"`
	BytesReceived int64     `json:"bytes_received"`
}
