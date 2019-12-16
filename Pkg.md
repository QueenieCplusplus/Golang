# Go_Packages
Golang 語言常用套件，可勤加練習。

# go std 標準

  * types
  
  * import
  
  * constant
  
  * printer
  
  * parser
  
  * ast (represent syntax trees for Go packages)
  
# errors 報錯

 https://golang.org/pkg/errors/

# fmt 輸出

  https://golang.org/pkg/fmt/

# time 計時器

  https://golang.org/pkg/time/
  
    start := time.Now()
    //... operation that takes 20 milliseconds ...
    t := time.Now()
    elapsed := t.Sub(start)
    
# sync/atomic 

  https://golang.org/pkg/sync/atomic/
  
# signal 訊號

  https://golang.org/pkg/os/signal/
  
    func Notify(c chan<- os.Signal, sig ...os.Signal)

# PRC 遠程呼叫

  https://golang.org/pkg/net/rpc/jsonrpc/
  
     func Dial(network, address string) (*rpc.Client, error)
     func NewClient(conn io.ReadWriteCloser) *rpc.Client
     func ServeConn(conn io.ReadWriteCloser)

# http 網路傳輸協定

  https://golang.org/pkg/net/http/
  
  //Server
  
    s := &http.Server{
    
      Addr:           ":8080",
      Handler:        myHandler,
      ReadTimeout:    10 * time.Second,
      WriteTimeout:   10 * time.Second,
      MaxHeaderBytes: 1 << 20,
      
    } 
    
    log.Fatal(s.ListenAndServe())
    
  //Router    
  
    http.Handle("/foo", fooHandler)
    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
    
  //RWD
  
    resp, err := http.Get("http://example.com/")
    if err != nil {
      // handle error...
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    
  // Post within UrlID
  
    resp, err := http.PostForm("http://example.com/form",
    url.Values{"key": {"Value"}, "id": {"123"}})
    
# ioutil 輸出入
    
   https://golang.org/pkg/io/ioutil/
   
# html escaping text 

    func EscapeString(s string) string
    
    //EscapeString escapes special characters like "<" to become "&lt;". It escapes only five such characters: <, >, &, ' and ". 
    
 # token 權杖
 
  https://golang.org/pkg/go/token/
  
# hash 亂碼

  https://golang.org/pkg/hash/
  
# crpto 加密

  https://golang.org/pkg/crypto/

  many others etc,. =>
  
  (1) cypher
  
  (2) hmac
  
  (3) sha
  
  (4) dsa
  
  (5) rsa
  
  (6) x509
  
# compress 壓縮器

  https://golang.org/pkg/compress/
  
  (1) bzip2
  
  (2) gzip
  
  (3) flat
  
  (4) zlib
  
# encode 編碼器

  (1) csv

  (2) json
  
  (3) byte
  
  (4) gob
  
  (5) hex
  
  (6) xml
  
  (7) bin
  
  (8) asci
  
  (9) base64
  
  (10) asn1 x690

# tar | zip 檔案

  https://golang.org/pkg/archive/
  
# bufferio 緩衝器

  https://golang.org/pkg/bufio/
  
# SQL Driver 資料庫驅動中介軟體

  https://golang.org/pkg/database/sql/driver/






