package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	url2 "net/url"
	"os"
	"strconv"
	"strings"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Value  string
	Course []string
}

func https() {
	filePtr, err := os.Open("./config.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	var info []Website
	// 创建json解码器
	var Cook1, Cook2, Cook3, Cook4, t, s, m, traineeId string
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	var rj1 []string
	var rj2 []string
	var rj3 []string
	var rj [31]string
	for i := 0; i < len(info); i++ {
		switch info[i].Name {
		case "开场头":
			rj1 = info[i].Course
		case "天气头":
			rj2 = info[i].Course
		case "一堆废话":
			rj3 = info[i].Course
		}
	}
	for i := 0; i < len(info); i++ {
		switch info[i].Name {
		case "Hm_lvt_5b943524066f14e8c8dc6a3c3a69d9ca":
			Cook1 = info[i].Value
			//fmt.Println(Cook1)
		case "acw_tc":
			Cook2 = info[i].Value
			//fmt.Println(Cook2)
		case "JSESSIONID":
			Cook3 = info[i].Value
			//fmt.Println(Cook3)
		case "Hm_lpvt_5b943524066f14e8c8dc6a3c3a69d9ca":
			Cook4 = info[i].Value
			//fmt.Println(Cook4)
		case "t":
			t = info[i].Value
		case "m":
			m = info[i].Value
		case "s":
			s = info[i].Value
		case "traineeId":
			traineeId = info[i].Value
		}
	}
	//fmt.Println(rj[28])
	//fmt.Println(traineeId)
	client := http.Client{}
	//Post请求示例
	url := "https://www.xybsyw.com/practice/student/blogs/save.action"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"

	//contentType := "application/json"
	//data := `{"schoolTermId":"4875","type":d}`
	cookie1 := &http.Cookie{
		Name:  "Hm_lvt_5b943524066f14e8c8dc6a3c3a69d9ca",
		Value: Cook1,
	}
	cookie2 := &http.Cookie{
		Name:  "acw_tc",
		Value: Cook2,
	}
	cookie3 := &http.Cookie{
		Name:  "JSESSIONID",
		Value: Cook3,
	}
	cookie4 := &http.Cookie{
		Name:  "Hm_lpvt_5b943524066f14e8c8dc6a3c3a69d9ca",
		Value: Cook4,
	}
	for i := 0; i < len(info); i++ {
		switch info[i].Name {
		case "日记时间":
			a, _ := strconv.Atoi(info[i].Course[3])
			b, _ := strconv.Atoi(info[i].Course[2])
			for j := b; j <= a; j++ {

				rj[j] = fmt.Sprintf("%v%v%v", rj1[rand.Intn(len(rj1))], rj2[rand.Intn(len(rj2))], rj3[rand.Intn(len(rj3))])
				data := "traineeId=" + traineeId + "&title=%E5%AE%9E%E4%B9%A0%E6%97%A5%E5%BF%97&content=" + url2.QueryEscape(rj[j]) + "&status=1&visicty=2&type=d&startDate=" + info[i].Course[0] + "." + info[i].Course[1] + "." + strconv.Itoa(j)
				resp, err := http.NewRequest("POST", url, strings.NewReader(data))
				resp.AddCookie(cookie1)
				resp.AddCookie(cookie2)
				resp.AddCookie(cookie3)
				resp.AddCookie(cookie4)
				resp.Header.Set("Connection", "close")
				resp.Header.Set("Content-Length", "24")
				resp.Header.Set("n", "content,practicePurpose,practiceContent,practiceRequirement,otherRequirement,practiceDescript,securityBook,responsibilities,selfAppraisal,file")
				resp.Header.Set("t", t)
				resp.Header.Set("sec-ch-ua-mobile", "?0")
				resp.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36")
				resp.Header.Set("m", m)
				resp.Header.Set("Accept", "application/json, text/plain, */*")
				resp.Header.Set("s", s)
				resp.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				resp.Header.Set("sec-ch-ua-platform", "macOS")
				resp.Header.Set("Origin", "https:www.xybsyw.com")
				resp.Header.Set("Sec-Fetch-Site", "same-origin")
				resp.Header.Set("Sec-Fetch-Mode", "cors")
				resp.Header.Set("Sec-Fetch-Dest", "empty")
				resp.Header.Set("Referer", "https:www.xybsyw.com/personal/")
				resp.Header.Set("Accept-Encoding", "gzip, deflate")
				resp.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
				//fmt.Println(resp)
				res, err := client.Do(resp)
				if err != nil {
					log.Println("err")
				}
				defer res.Body.Close()
				b, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Println("err")
				}
				fmt.Printf("已成功提交%v份,如果后面没有出现操作成功，那就是出问题了请联系我", j)
				fmt.Println(string(b))
			}

		}
	}

}

func main() {
	https()

}
