package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"strconv"
	"time"
)

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 获取时间戳
func GetUnix() int64 {

	return time.Now().Unix()
}

// 获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// md5加密
func MD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 表示把string转换成int
func Int(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// 表示把int转换成string
func String(n int) string {
	str := strconv.Itoa(n)
	return str
}

func UploadImg(c *gin.Context, pictureName string) (string, error) {
	//获取图像文件
	file, err1 := c.FormFile(pictureName)
	if err1 != nil {
		return "", err1
	}
	//获取后缀名，判断类型
	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件后缀不合法")
	}
	day := GetDay()
	dir := "./static/upload/" + day

	err2 := os.MkdirAll(dir, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
		return "", err2
	}

	fileName := strconv.FormatInt(GetUnix(), 10) + extName
	dst := path.Join(dir, fileName)
	err3 := c.SaveUploadedFile(file, dst)
	if err3 != nil {
		fmt.Println(err3)
	}
	return dst, nil
}
