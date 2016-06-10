package main

import (
	"net/http"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"fmt"
	"strings"
	"net/url"
	//"reflect"
)

func httpDo(code string,start,end int, path string) {
    client := &http.Client{}

	
    req1, err := http.NewRequest("GET", "http://www.gdwater.gov.cn:9001/Report/RainReport.aspx",nil)
    if err != nil {
        // handle error
    }
	resp1, err := client.Do(req1)
	defer resp1.Body.Close()
	body1, err := ioutil.ReadAll(resp1.Body)
    if err != nil {
        // handle error

    }
	month := []int{1,31,28,31,30,31,30,31,31,30,31,30,31}
	str := string(body1)
	r,_ := regexp.Compile("name=\"__VIEWSTATE\" id=\"__VIEWSTATE\" value=\"(.*?)\"")

	byteViewState := r.FindAllString(str,1)
	viewState := strings.Replace(byteViewState[0],"name=\"__VIEWSTATE\" id=\"__VIEWSTATE\" value=\"","",1)
	viewState = strings.TrimSuffix(viewState,"\"")
	r,_ = regexp.Compile("name=\"__EVENTVALIDATION\" id=\"__EVENTVALIDATION\" value=\"(.*?)\"")
	byteEnv := r.FindAllString(str,1)
	env := strings.Replace(byteEnv[0],"name=\"__EVENTVALIDATION\" id=\"__EVENTVALIDATION\" value=\"","",1)
	env = strings.TrimSuffix(env,"\"")
	fmt.Println("get viewState and env done!")
	for m:=start;m<end;m++{
		for d:=1;d<=month[m];d++{
			for h:=0;h<=23;h++{
		str1,str2,p := getDay(m,d,h)
		fmt.Println(str1,str2,p)
		req, err := http.NewRequest("POST", "http://www.gdwater.gov.cn:9001/Report/RainReport.aspx",nil)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		form := url.Values{}
	
	
		form.Set("__VIEWSTATE", viewState)
    	form.Set("__EVENTVALIDATION", env)
		form.Set("ddl_addvcd", code)
		form.Set("txt_search", "")
		form.Set("txt_time1",str1)
 		form.Set("txt_time2",str2)
		form.Set("btn_query","查询")
		form.Set("hidsearch","")
		req.Body = ioutil.NopCloser(strings.NewReader(form.Encode()))
	
    	resp, err := client.Do(req)
 	
    
 
    	body, err := ioutil.ReadAll(resp.Body)

    	if err != nil {
        	// handle error
    	}
		dir :=filepath.Dir(path+"1.html")
		if err:=os.MkdirAll(dir,0777); err!=nil{
			return 
		}
		file, err := os.Create(path+p+".html")
	 	defer file.Close()
		if err!=nil{
			return 
		}
		file.WriteString(string(body))
		resp.Body.Close()
		
		}
		}
		return 
	}
}

func getDay(m,d,h int)(str1,str2,path string){


	num := []string{
		"00","01","02","03","04","05","06","07","08","09",
		"10","11","12","13","14","15","16","17","18","19",
		"20","21","22","23","24","25","26","27","28","29",
		"30","31","32","33"}
	month := []int{1,31,28,31,30,31,30,31,31,30,31,30,31}
	str1 ="2015-"+num[m]+"-"+num[d]+" "+num[h]+":00"
	path ="2015_"+num[m]+"_"+num[d]+"_"+num[h]
	if h == 23{
		h = 0
		d = d+1
	}else{
		h = h+1
	}
	if d == month[m]+1{
		d = 1
		m = m+1
	}
	if m == 13{
		str2 = "2016-"
		m=1
	}else{
		str2 = "2015-"
	}
	str2 =str2+ num[m]+"-"+num[d]+" "+num[h]+":00"
	
	return 
}

func main() {
	//fmt.Println(getDay(11,29,23))
	var s,e int
	var code,path string
	fmt.Println("输入开始月份，结束按回车键:")
	fmt.Scanln(&s)
	fmt.Printf("输入的月份是：%d\n",s)
	fmt.Println("输入结束月份(不含，想爬取12月请输入13)，结束按回车键:")
	fmt.Scanln(&e)
	fmt.Printf("输入的月份是：%d\n",s)
	fmt.Println("输入区域编码(东莞：441900惠州：441300 深圳：440300)，结束按回车键:")
	fmt.Scanln(&code)
	fmt.Printf("输入的区域编码是：%s\n",code)
	fmt.Println("输入保存地址，以'/'结尾，结束按回车键:")
	fmt.Scanln(&path)
	fmt.Printf("输入的地址是：%s\n",path)
	for i:=s;i<e;i++{
		httpDo(code,i,i+1,path)
	}
}
