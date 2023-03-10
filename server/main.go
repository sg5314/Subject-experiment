package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	// "strconv"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func save_data(index []string, contents []string, subject_name string) {

	f, err := os.Create(fmt.Sprintf("%s.csv", subject_name)) // 書き込む先のファイル   subject_name
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(f)

	for ind, con := range contents {
		fmt.Println(ind, index[ind], con)
		if con == "モダンと思う" {
			writer.Write([]string{strconv.Itoa(ind), index[ind], "modern"})
		} else if con == "モダンでないと思う" {
			writer.Write([]string{strconv.Itoa(ind), index[ind], "not_modern"})
		}

	}
	writer.Flush()

}

func subject_set(subject_id string) (int, []string, string) {

	const (
		subject_living = 100
		subject_office = 97
	)
	var image_num int
	var subject_imgs []string
	var subject_name string

	num, _ := strconv.Atoi(subject_id) //　subject_id　半角数字
	fmt.Println(num)

	switch num {
	case 1:
		image_num = subject_living
		subject_imgs = dirwalk("./img/Subject1_living/")
		subject_name = "Subject1_living"
	case 2:
		image_num = subject_living
		subject_imgs = dirwalk("./img/Subject2_living/")
		subject_name = "Subject2_living"
	case 3:
		image_num = subject_living
		subject_imgs = dirwalk("./img/Subject3_living/")
		subject_name = "Subject3_living"
	case 4:
		image_num = subject_living
		subject_imgs = dirwalk("./img/Subject4_living/")
		subject_name = "Subject4_living"
	case 5:
		image_num = subject_living
		subject_imgs = dirwalk("./img/Subject5_living/")
		subject_name = "Subject5_living"
	case 6:
		image_num = subject_office
		subject_imgs = dirwalk("./img/Subject6_office/")
		subject_name = "Subject6_office"
	case 7:
		image_num = subject_office
		subject_imgs = dirwalk("./img/Subject7_office/")
		subject_name = "Subject7_office"
	case 8:
		image_num = subject_office
		subject_imgs = dirwalk("./img/Subject8_office/")
		subject_name = "Subject8_office"
	case 9:
		image_num = subject_office
		subject_imgs = dirwalk("./img/Subject9_office/")
		subject_name = "Subject9_office"
	case 10:
		image_num = subject_office
		subject_imgs = dirwalk("./img/Subject10_office/")
		subject_name = "Subject10_office"
	default:
		//　どれにも当てはまらない
		image_num = -1
	}

	return image_num, subject_imgs, subject_name
}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/img", "./img/")
	router.Static("/css", "./assets/css")
	router.Static("/js", "./assets/js")

	var subjecter_ans []string
	var page_count int
	var subject_imgs []string
	var image_num int
	var subject_name string

	router.GET("/", func(c *gin.Context) {
		// 最初の画面
		c.HTML(http.StatusOK, "entry.html", gin.H{})
		page_count = 0
		subjecter_ans = nil
	})

	router.POST("/", func(c *gin.Context) {
		ans := c.PostForm("str")
		fmt.Println(ans)

		if ans != "スタート" {
			// 随時、保存
			subjecter_ans = append(subjecter_ans, ans)
		} else {
			// スタートボタンを押した時
			id := c.PostForm("id")
			image_num, subject_imgs, subject_name = subject_set(id)
			if image_num == -1 {
				c.HTML(http.StatusOK, "entry.html", gin.H{})
				return
			}
		}

		if page_count == image_num {
			save_data(subject_imgs, subjecter_ans, subject_name) //データの保存
			c.HTML(http.StatusOK, "finish.html", gin.H{})
		} else {
			c.HTML(http.StatusOK, "research.html", gin.H{
				"route":        subject_imgs[page_count],
				"progress":     page_count,
				"max_progress": image_num,
			})
			page_count += 1
		}

	})
	router.Run(":3000")

}
