package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const CHUNK_SIZE int64 = 1024 * 1024

func StreamVideo(context *gin.Context) {
	rangeHeader := context.Request.Header["Range"][0]
	ranges := strings.Split(strings.Replace(rangeHeader, "bytes=", "", 1), "-")
	start, err := strconv.ParseInt(ranges[0], 10, 64)
	if err != nil {
		context.Status(http.StatusNotAcceptable)
		context.Abort()
		return
	}
	// var end int64
	// if len(ranges) == 2 && len(ranges[1]) > 0 {
	// 	end, err = strconv.ParseInt(ranges[1], 10, 64)
	// 	if err != nil {
	// 		context.Status(http.StatusNotAcceptable)
	// 		context.Abort()
	// 		return
	// 	}
	// } else {
	// 	end = start + CHUNK_SIZE
	// }

	// binary read
	name := context.Param("name")
	file, err := os.Open("static/" + name + ".mp4")

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var totalFileSize int64 = stats.Size()
	end := start + CHUNK_SIZE
	if end > totalFileSize-1 {
		end = totalFileSize - 1
	}

	contentLength := end - start + 1
	data := make([]byte, contentLength)
	file.Seek(start, 0)
	bytesRead, _ := file.Read(data)
	fmt.Println("Bytes read: " + strconv.Itoa(bytesRead))

	context.Writer.Header().Add("Content-Range", "bytes "+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10)+"/"+strconv.FormatInt(totalFileSize, 10))
	context.Writer.Header().Add("Accept-Ranges", "bytes")
	context.Writer.Header().Add("Content-Length", strconv.FormatInt(contentLength, 10))
	context.Data(206, "video/mp4", data)
}

func TestClaims(context *gin.Context) {

}
