package main

import (
	"edulx/hax_woiden_jk/config"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

func Hax() string {
	url := "https://hax.co.id/create-vps/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return ""
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.62")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Host", "hax.co.id")
	//req.Header.Add("Cookie", "PHPSESSID=c82418d070a8c8e767e0ce303a71159a; _ga=GA1.3.793762279.1679893292; __cf_bm=n1O7yJJxTW9hu.qhe6QRsCcE_kAZmAH1JHs6tEvmS.U-1680529193-0-Aemg00EIOnWChqBCRsyVnwuWHdxRP+evClmW8EKRQJQgiAb+BHppJXgtZc3FOuxR655TQ1GfpXRoYi62SQbA/5fTF3uyKoWeCOxr2tUojSIQVF8JcHQQAQcRo+HMqoOY8A; FCCDCF=%5Bnull%2Cnull%2Cnull%2C%5B%22CPpSJ0APpSJ0AEsABBENC9CoAP_AAG_AABAYINJB7D7FbSFCyP57aLsAMAhXRkCAQqQCAASBAmABQAKQIAQCkkAYFESgBAACAAAgIAJBIQIMCAgACUABQAAAAAEEAAAABAAIIAAAgAEAAAAIAAACAIAAEAAIAAAAEAAAmQhAAIIACAAAhAAAIAAAAAAAAAAAAgCAAAAAAAAAAAAAAAAAAQQaQD2F2K2kKEkfjWUWYAQBCujIEAhUAEAAECBIAAAAUgQAgFIIAwAIlACAAAAABAQAQCQgAQABAAAoACgAAAAAAAAAAAAAAQQAABAAIAAAAAAAAEAQAAIAAQAAAAAAABEhCAAQQAEAAAAAAAQAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAgAA%22%2C%221~2072.70.89.93.108.122.149.196.2253.2299.259.2357.311.317.323.2373.338.358.2415.415.2506.2526.482.486.494.495.2568.2571.2575.540.574.2624.624.2677.827.864.981.1048.1051.1095.1097.1171.1201.1205.1276.1301.1365.1415.1449.1570.1577.1651.1716.1753.1765.1870.1878.1889.1958.2012%22%2C%22ED667BE9-C3D9-464E-AE09-66952825F61F%22%5D%2Cnull%2Cnull%2C%5B%5D%5D; FCNEC=%5B%5B%22AKsRol8G5fkNHyoHbXzmXFTlKM7KH0iPAJbKDS4iyZ0XLicl7tvXtuH4sNjs4RPCJC30SL7NwBJC_E11HzjQmvk5cFN7VXN5qp94lNdxbHqZn86SmCcA1PlrgOgBm-1nKp4VBR9_8qwTZf8i3Q8BXnpdn-k3ZNSu0g%3D%3D%22%5D%2Cnull%2C%5B%5D%5D; PHPSESSID=c82418d070a8c8e767e0ce303a71159a")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("\rHAX status code error: %d %s                         ", res.StatusCode, res.Status)
		time.Sleep(5 * time.Second)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var R string
	// Find the review items
	doc.Find("#datacenter option").Each(func(i int, s *goquery.Selection) {
		if s.Text() != "-select-" && s.Text() != "" {
			R = s.Text()
			fmt.Println("HAX:" + s.Text() + time.Now().Format("2006-01-02 15:04:05"))
		} else {
			R = ""
			fmt.Print("\rHAX:" + s.Text() + time.Now().Format("2006-01-02 15:04:05") + "             ")
			time.Sleep(time.Second)
		}
	})
	return R
}
func Woiden() string {
	url := "https://woiden.id/create-vps/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return ""
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.62")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Host", "hax.co.id")
	//req.Header.Add("Cookie", "PHPSESSID=c82418d070a8c8e767e0ce303a71159a; _ga=GA1.3.793762279.1679893292; __cf_bm=n1O7yJJxTW9hu.qhe6QRsCcE_kAZmAH1JHs6tEvmS.U-1680529193-0-Aemg00EIOnWChqBCRsyVnwuWHdxRP+evClmW8EKRQJQgiAb+BHppJXgtZc3FOuxR655TQ1GfpXRoYi62SQbA/5fTF3uyKoWeCOxr2tUojSIQVF8JcHQQAQcRo+HMqoOY8A; FCCDCF=%5Bnull%2Cnull%2Cnull%2C%5B%22CPpSJ0APpSJ0AEsABBENC9CoAP_AAG_AABAYINJB7D7FbSFCyP57aLsAMAhXRkCAQqQCAASBAmABQAKQIAQCkkAYFESgBAACAAAgIAJBIQIMCAgACUABQAAAAAEEAAAABAAIIAAAgAEAAAAIAAACAIAAEAAIAAAAEAAAmQhAAIIACAAAhAAAIAAAAAAAAAAAAgCAAAAAAAAAAAAAAAAAAQQaQD2F2K2kKEkfjWUWYAQBCujIEAhUAEAAECBIAAAAUgQAgFIIAwAIlACAAAAABAQAQCQgAQABAAAoACgAAAAAAAAAAAAAAQQAABAAIAAAAAAAAEAQAAIAAQAAAAAAABEhCAAQQAEAAAAAAAQAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAgAA%22%2C%221~2072.70.89.93.108.122.149.196.2253.2299.259.2357.311.317.323.2373.338.358.2415.415.2506.2526.482.486.494.495.2568.2571.2575.540.574.2624.624.2677.827.864.981.1048.1051.1095.1097.1171.1201.1205.1276.1301.1365.1415.1449.1570.1577.1651.1716.1753.1765.1870.1878.1889.1958.2012%22%2C%22ED667BE9-C3D9-464E-AE09-66952825F61F%22%5D%2Cnull%2Cnull%2C%5B%5D%5D; FCNEC=%5B%5B%22AKsRol8G5fkNHyoHbXzmXFTlKM7KH0iPAJbKDS4iyZ0XLicl7tvXtuH4sNjs4RPCJC30SL7NwBJC_E11HzjQmvk5cFN7VXN5qp94lNdxbHqZn86SmCcA1PlrgOgBm-1nKp4VBR9_8qwTZf8i3Q8BXnpdn-k3ZNSu0g%3D%3D%22%5D%2Cnull%2C%5B%5D%5D; PHPSESSID=c82418d070a8c8e767e0ce303a71159a")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("\rwoiden status code error: %d %s                            ", res.StatusCode, res.Status)
		time.Sleep(5 * time.Second)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var R string
	// Find the review items
	doc.Find("#datacenter option").Each(func(i int, s *goquery.Selection) {
		if s.Text() != "-select-" && s.Text() != "" {
			R = s.Text()
			fmt.Println(s.Text() + time.Now().Format("2006-01-02 15:04:05"))
		} else {
			R = ""
			fmt.Print("\rWoiden:" + s.Text() + time.Now().Format("2006-01-02 15:04:05") + "       ")
			time.Sleep(time.Second)
		}
	})
	return R
}
func Pushplus(c string, ch string) {
	url := "http://www.pushplus.plus/send"
	method := "POST"
	var P struct {
		Title    string `json:"title"`
		Token    string `json:"token"`
		Content  string `json:"content"`
		Channel  string `json:"channel"`
		Template string `json:"template"`
	}
	P.Token = Con.PushPlus.Token
	//随机数设置种子
	rand.Seed(time.Now().UnixNano())
	P.Content = c + "<br>" + "<p style=\"display:none;\">防止提示重复通知:"
	for i := 0; i < rand.Intn(5); i++ {
		P.Content = P.Content + "<br>" + strconv.Itoa(rand.Intn(100000000))
	}
	P.Content = P.Content + "</p>"
	P.Channel = ch
	P.Title = c
	P.Template = "html"
	// P 转json
	v, err := json.Marshal(P)
	if err != nil {
		fmt.Println(err)
		return
	}
	payload := strings.NewReader(string(v))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "www.pushplus.plus")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func Email(to string, title string, content string) bool {
	subject := fmt.Sprintf("Subject: %s\r\n", title)
	send := fmt.Sprintf("From: %s 库存监控\r\n", Con.Email.From)
	receiver := fmt.Sprintf("To: %s\r\n", to)
	contentType := "Content-Type: text/html" + "; charset=UTF-8\r\n\r\n"
	msg := []byte(subject + send + receiver + contentType + content)
	addr := Con.Email.Smtp + ":" + Con.Email.Port
	auth := smtp.PlainAuth("", Con.Email.From, Con.Email.Key, Con.Email.Smtp)
	from := Con.Email.From
	var tos = []string{to}
	err := smtp.SendMail(addr, auth, from, tos, msg)
	if err != nil {
		return false
	}
	return true
}

var Con config.Config
var InitErr error
var TZtime map[string]int64

func RunHAX() {
	h := Hax()
	if h != "" {
		fmt.Println("HAX " + h)
		for _, v1 := range Con.Other.Concern {
			if i := strings.Index(h, v1); i != -1 {
				for _, v := range Con.PushPlus.Channel {
					Pushplus("!HAX"+v1+h+time.Now().String()[0:19], v)
				}
				for _, v := range Con.Email.To {
					if Email(v, "!HAX"+v1+h, "HAX"+h) {
						fmt.Println(v + "发信成功")

					} else {
						fmt.Println(v + "发信失败")
					}
				}
			}
		}
		if time.Now().Unix()-TZtime[h] > Con.Other.Time {
			for _, v := range Con.PushPlus.Channel {
				Pushplus("HAX"+h+time.Now().String()[0:19], v)
			}
			for _, v := range Con.Email.To {
				if Email(v, "HAX"+h, "HAX"+h) {
					fmt.Println(v + "发信成功")

				} else {
					fmt.Println(v + "发信失败")
				}
			}
			TZtime[h] = time.Now().Unix()
		}
	}
}
func RunWoiden() {
	w := Woiden()
	if w != "" {
		fmt.Println("WOIDEN " + w)
		for _, v1 := range Con.Other.Concern {
			if i := strings.Index(w, v1); i != -1 {
				for _, v := range Con.PushPlus.Channel {
					Pushplus("!WOIDEN"+v1+w+time.Now().String()[0:19], v)
				}
				for _, v := range Con.Email.To {
					if Email(v, "!WOIDEN"+v1+w, "WOIDEN"+w) {
						fmt.Println(v + "发信成功")
					} else {
						fmt.Println(v + "发信失败")
					}
				}
			}
		}
		if time.Now().Unix()-TZtime[w] > Con.Other.Time {
			for _, v := range Con.PushPlus.Channel {
				Pushplus("WOIDEN"+w+time.Now().String()[0:19], v)
			}
			for _, v := range Con.Email.To {
				if Email(v, "WOIDEN"+w, "WOIDEN"+w) {
					fmt.Println(v + "发信成功")

				} else {
					fmt.Println(v + "发信失败")
				}
			}
			TZtime[w] = time.Now().Unix()
		}

	}
}
func main() {
	if InitErr = Con.ReadYaml(); InitErr != nil {
		fmt.Println("config err")
		return
	}
	for _, v := range Con.Email.To {
		if Email(v, "监控启动成功", "监控启动成功") {
			fmt.Println(v + "发信成功")
		} else {
			fmt.Println(v + "发信失败")
		}
	}

	for _, v := range Con.PushPlus.Channel {
		Pushplus("监控启动成功"+v, v)
	}
	fmt.Println("请检查是否收到邮件和微信推送")
	for {
		go RunWoiden()
		go RunHAX()
		time.Sleep(time.Second)
	}
}
