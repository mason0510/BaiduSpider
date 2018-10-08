package main

import (
	"fmt"
	"strconv"
	"net/http"
	"os"
)

// 明确目标 http://tieba.baidu.com/f?kw=%E7%BA%AA%E5%BD%95%E7%89%87&ie=utf-8&pn=0
func DoWork(start, end int) {
	fmt.Printf("正在爬取\n", start, end)

	page:=make(chan int)
	for i := start; i <= end; i++ {
	go	SpiderPage(i,page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n",<-page)
	}

}
func SpiderPage(i int,page chan<- int) {
	url := "http://tieba.baidu.com/f?kw=%E7%BA%AA%E5%BD%95%E7%89%87&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Println("正在爬取网页 %s\n",i,url)
	//fmt.Printf("url=", url);
	result, err := HttpGet(url);
	if err != nil {
		fmt.Println("HTTP ERROR", err);
		return
	}
	//创建文件
	fileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(fileName)
	if err1 != nil {
		fmt.Println("os.Create", err1)
		return
	}
	//将结果写入文件中
	f.WriteString(result)
	//关闭流
	f.Close()

	page<-i
}
func HttpGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HttpGet", err)
		return
	}
	defer resp.Body.Close()
	//读取网页内容
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err", err)
			break
		}
		result += string(buf[:n])

	}
	return
}
func main() {
	var start, end int
	fmt.Printf("输入开始页面")
	fmt.Scan(&start)
	fmt.Printf("输入终止页面")
	fmt.Scan(&end)
	DoWork(start, end)
}

//爬取数据

//处理数据
