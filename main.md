# Go_init_mod
為套件初始化的方法

# Package Structure, 專案架構

    $ tree .
    
>>>

    .
    ├── LICENSE
    ├── README.md
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── v2
        ├── go.mod
        ├── go.sum
        └── main.go

# Check Go Version, 確認版本

Go 版本在 1.12 起，GoMod 的套件管理工具功能已經穩定，1.13 開始設為默認工具。

    $go version
    
>>>

    go version go1.13.4 darwin/amd64
    
# init Go Module, 初始化一Go套件

    $mkdir -p ~/gopher/gomodQueen
    $cd ~/gopher/gomodQueen
    $go mod init gomodQueen
    go: creating new go.mod: module called gomodQueen
    $vim go.mod
    也可以使用 cat 直接在終端機上輸出
    
# check go.mod, 進入模組內容

此源檔案類似 Node.js 的 Package.json 中的 package name 和 

    //當前模組名稱
    module gomodQueen

    //此模組版本號
    go 1.13
    ~                                                                               
    ~                                                                               
    ~                                                                               
    ~                                                                               
    ~                                                                               
    ~                                                                               
    ~                                                                               

    "go.mod" 3L, 27C
    
 # call Go Get to add Packages, 下載依賴套件
 
    $cd ~/gopher/gomodQueen
 
>>>
 
    $go get golang.org/x/crypto/sha3@master
>>>

    go: finding golang.org/x/crypto/sha3 master
    go: finding golang.org master
    go: finding golang.org/x/crypto master
    go: finding golang.org/x master
    go: downloading golang.org/x/crypto v0.0.0-20191119213627-4f8c1d86b1ba
    go: extracting golang.org/x/crypto v0.0.0-20191119213627-4f8c1d86b1ba
    go: downloading golang.org/x/sys v0.0.0-20190412213103-97732733099d
    go: extracting golang.org/x/sys v0.0.0-20190412213103-97732733099d
    
>>>

    $vim go.mod
    
>>>

    module gomodQueen

    go 1.13

    require golang.org/x/crypto v0.0.0-20191119213627-4f8c1d86b1ba // indirect
    ~                                                                               
    ~                                                                               
    ~                                                                               
    ~                                                                                                                                                                                                                                      
    "go.mod" 5L, 103C
    
 # go sum
 
 
    $cd ~/gopher/gomodQueen
    $ ls
    
>>>

    go.mod  go.sum
 
 作用類似 Nodejs 的 package-lock.json
    
 # to code another go src file with installing modules, 繕打程式碼建立代碼頁
 
     package main

    import(
        "fmt"
    )

    func main() {
        fmt.Println("Queenie plays Golang in 2019.")
    }

>>>
IDE | Terminal will suggest you the modules (dependencies) required shall be install!

    go.toolsGopath setting is not set. Using GOPATH /Users/pintred/go
    Installing 17 tools at /Users/pintred/go/bin in module mode.
      gocode
      gopkgs
      go-outline
      go-symbols
      guru
      gorename
      gotests
      gomodifytags
      impl
      fillstruct
      goplay
      godoctor
      dlv
      gocode-gomod
      godef
      goreturns
      golint

>>>
Install Proccess

    Installing github.com/mdempsky/gocode SUCCEEDED
    Installing github.com/uudashr/gopkgs/cmd/gopkgs SUCCEEDED
    Installing github.com/ramya-rao-a/go-outline SUCCEEDED
    Installing github.com/acroca/go-symbols SUCCEEDED
    Installing golang.org/x/tools/cmd/guru SUCCEEDED
    Installing golang.org/x/tools/cmd/gorename SUCCEEDED
    Installing github.com/cweill/gotests/... SUCCEEDED
    Installing github.com/fatih/gomodifytags SUCCEEDED
    Installing github.com/josharian/impl SUCCEEDED
    Installing github.com/davidrjenni/reftools/cmd/fillstruct SUCCEEDED
    Installing github.com/haya14busa/goplay/cmd/goplay SUCCEEDED
    Installing github.com/godoctor/godoctor SUCCEEDED
    Installing github.com/go-delve/delve/cmd/dlv SUCCEEDED
    Installing github.com/stamblerre/gocode SUCCEEDED
    Installing github.com/rogpeppe/godef SUCCEEDED
    Installing github.com/sqs/goreturns SUCCEEDED
    Installing golang.org/x/lint/golint SUCCEEDED

    All tools successfully installed. You're ready to Go :).

# Go Run, 執行代碼並且觀看輸出結果

    $cd /Users/pintred/Desktop        
    $go run main.go
>>>
    Queenie plays Golang in 2019.
    
# add main.go to working dir, 專案加入源碼

    //move path/text to newPath/newTextName
    $mv /Users/pintred/Desktop/main.go ~/gopher/gomodQueen/main.go


# Update Versions to Dependencies of the Working Package, 升級依賴套件

利用 vim 打開專案下的 go.mod，注入為 require，其後跟著依賴套件的 URL 路徑，並且操作指令如下：

    $go mod download
    
查看需要版本升級的依賴套件

    $go list -u -m all
    
升級特定版本號

    $go get foo@'<v1.6.2'
    
移除依賴套件

    $go mod tidy
    
轉換 json 格式輸出依賴套件工具名稱

    $go list -m -json all 
    
>>>
    
    {
        "Path": "golang.org/x/text",
        "Version": "v0.3.0",
        "Time": "2017-12-14T13:08:43Z",
        "Indirect": true,
        "Dir": "/Users/lishude/go/pkg/mod/golang.org/x/text@v0.3.0",
        "GoMod": 
        "/Users/lishude/go/pkg/mod/cache/download/golang.org/x/text/@v/v0.3.0.mod"
    }
    
    {
        "Path": "rsc.io/quote",
        "Version": "v1.5.2",
        "Time": "2018-02-14T15:44:20Z",
        "Dir": "/Users/lishude/go/pkg/mod/rsc.io/quote@v1.5.2",
        "GoMod":           
        "/Users/lishude/go/pkg/mod/cache/download/rsc.io/quote/@v/v1.5.2.mod"
    }
    
 https://zhuanlan.zhihu.com/p/59687626?fbclid=IwAR3OrFFsUYuIi5A1ITKi10JPjg2dPNnXWI3MgkQsTHMoiw30NPaPnRul9C4
