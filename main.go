package main

import(
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
   // "github.com/PuerkitoBio/goquery"
    "regexp"
    "sync"
    "strings"
)

type Crawler struct{
    indexUrl string
    scheme string
    host string
    schemeAndHost string
    hadDoneUrl map[string] bool
    path string
    ch chan bool
    w sync.WaitGroup
    lock *sync.Mutex
    gtNum int
}

func NewCrawler() *Crawler {
    c := &Crawler{
        indexUrl: "url",
        path : "d:\\t.txt",
        gtNum:0,
        lock:&sync.Mutex{}}
    c.hadDoneUrl = make(map[string]bool,1000)
    c.ch = make(chan bool, 1000)
    
    return c
}

//func (this * Crawler) write(ctx string, p string){
    
//}

func (this *Crawler)downloadUrl(url string)(string){
    resp,err:=http.Get(url)
    if err!=nil{
        fmt.Println(err)
        log.Fatal(err)
        return ""
    }
    defer resp.Body.Close()
    body, err1 :=ioutil.ReadAll(resp.Body)
    if(err1 !=nil){
        fmt.Println(err1)
            return ""
    }
   // fmt.Println(string(body))
    return string(body)
}
func (this * Crawler) GoCraw(url string){
    this.w.Add(1)
    
    this.ch <-true
    this.lock.Lock()
    this.gtNum ++
    fmt.Println("Now run",this.gtNum,"goroutines.")
    this.lock.Unlock()
    go func(){
        defer func(){
            this.w.Done()
        }()
        ctx := this.downloadUrl(url)
        children:=this.pharseHTML(ctx)
        this.lock.Lock()
        this.gtNum --
        fmt.Println("Now run",this.gtNum,"goroutines.")
        this.lock.Unlock()
        
        <-this.ch
        for _,cUrl :=range children{
            fmt.Println("Get",cUrl)
            this.GoCraw(cUrl)
        }
    }()
}


func (this * Crawler) RunCrawler(url string){
      
      ctx:=this.downloadUrl(url)
      children:=this.pharseHTML(ctx)
      
      for _,cUrl :=range children {
            fmt.Println("Get",cUrl)
            
            this.RunCrawler(cUrl)
   

      }
}

func (this * Crawler)pharseHTML(content string)(children []string){
    host := "ithome.com"
    //this.hasDone :=make(map[string]bool,1000)
    regular :="(?i)(src=|href=)[\"']([^#].*?)[\"']"
    reg := regexp.MustCompile(regular)
    re := reg.FindAllStringSubmatch(content, -1)
    for _,each :=range re{
       // fmt.Println("%d:%s",index,each[0])
        
     //   rawFullUrl :=each[0]
     //   rawFullUrlPrefix :=each[1]
        rawUrl :=each[2]
        cUrl := rawUrl

        if strings.HasPrefix(cUrl,"//"){
            cUrl = "http:" + cUrl
        }else if strings.HasPrefix(cUrl,"/"){
            cUrl = "http://" + host + cUrl
        }
        if strings.Contains(cUrl,host) == false{
            continue
        }
        if strings.HasSuffix(cUrl,"/"){
            cUrl = cUrl + "index.html"
        }else if (strings.HasSuffix(cUrl,"html") == false && strings.HasSuffix(cUrl,"htm")== false) {
            continue
        }
        if _,ok :=this.hadDoneUrl[cUrl];ok{
            continue
        }  else {
            this.hadDoneUrl[cUrl] = true
        }
        children = append(children,cUrl)
    }
    return 
}

func login()bool{
    return false
}
func main() {
    cr := NewCrawler()
    cr.GoCraw("http://www.ithome.com")
    cr.w.Wait()
}