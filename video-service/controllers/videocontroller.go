package controllers

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"video-service/database"
	"video-service/models"
	"video-service/utils"

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

func ReportVideo(context *gin.Context) {
	videoId := context.Param("id")
	var video models.Video

	if err := database.Instance.First(&video, videoId).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	video.Reported = true

	database.Instance.Save(&video)

	context.Status(http.StatusOK)
}

func GetAllReportedVideos(context *gin.Context) {
	_, claims := utils.GetTokenClaims(context)
	if claims.Role != "Administrator" {
		context.JSON(401, gin.H{"error": "unauthorized role"})
		context.Abort()
		return
	}

	var videos []models.Video
	database.Instance.Where("reported = ?", true).Find(&videos)

	context.JSON(http.StatusOK, videos)
}

func UploadVideo(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	fmt.Println(file.Filename)

	// generate random number for filename
	rand.Seed(time.Now().UnixNano())
	rndNum := rand.Intn(math.MaxInt32-0) + 0
	file.Filename = strconv.Itoa(rndNum) + ".mp4"

	// Upload the file to specific dst.
	// wd, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// parent := filepath.Dir(wd)
	err := c.SaveUploadedFile(file, "static/"+file.Filename)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
