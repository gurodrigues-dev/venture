package controllers

import (
	"fmt"
	"gin/config"
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func Health(c *gin.Context) {

	_, err := config.LoadEnvironmentVariables()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error loading variables",
			"error":  err.Error(),
		})

		return
	}

	cpu, err := cpu.Percent(0, false)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "loading error cpu metrics",
			"error":  err.Error(),
		})

		return
	}

	mem, err := mem.VirtualMemory()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "loading error memory metrics",
			"error":  err.Error(),
		})

		return
	}

	uptime, err := utils.GetUptime()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "loading error host metrics",
			"error":  err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"cpu":     fmt.Sprintf("%.1f", cpu[0]),
		"mem":     fmt.Sprintf("%.1f", mem.UsedPercent),
		"uptime":  uptime,
		"envs":    "load environments ok!",
	})

}

func PingPong(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

	return

}
