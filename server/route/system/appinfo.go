package route

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

// GetAppInfo returns metrics about the current running go process
func GetAppInfo(c *gin.Context) {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)

	c.JSON(http.StatusOK, gin.H{
		"num_goroutines": runtime.NumGoroutine(),

		"num_gc":        mem.NumGC,
		"alloc":         mem.Alloc,
		"total_alloc":   mem.TotalAlloc,
		"sys":           mem.Sys,
		"lookups":       mem.Lookups,
		"mallocs":       mem.Mallocs,
		"frees":         mem.Frees,
		"heap_alloc":    mem.HeapAlloc,
		"heap_sys":      mem.HeapSys,
		"heap_idle":     mem.HeapIdle,
		"heap_inuse":    mem.HeapInuse,
		"heap_released": mem.HeapReleased,
		"heap_objects":  mem.HeapObjects,
	})
}
