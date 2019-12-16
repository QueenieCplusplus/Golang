# Go_GoMod_No_GoPath_Anymore

專案結構中取代 GoPath 的套件包管理工具！

# GO Package Structure

因為 Golang 本身推崇專案理念，所以它為開發週期每個環節都提供了完整的工具和支援。從表現出結構與城市體制的細節處能窺見 Go 強調代碼和專案的一致性。

透過 go get，可以從公共程式庫 load 下開源程式碼並且使用它們。

# Go Work Dir

關於 Go 的工作區（即一對應專案的工作目錄），它包含了三個子目錄：

* src

* pkg

儲存經由 go install 指令建置安裝後的程式套件的 “.a” 壓縮檔。該目錄與 GoRoot 目錄下的 pkg 功能類似。區別在 pkg 目錄專門儲存開發者程式碼的壓縮檔，而 build 和 install 的程式壓縮檔則會以程式套件為單位，例如 <套件名稱>.a 的壓縮檔案，並且儲存於目前工作區域的 pkg 目錄下目錄中。

* pkg/../modules

與 Nodejs 的 pakage.json 作用相同。

* bin

透過 go install 指令完成 load modules or load dependencies 後，儲存由 go src file 產生的可執行檔。（在 Linux 系統下，此執行檔案是一個與原始程式檔案名稱相同的檔案，而 Microsoft 則產生原程式碼名稱額外加上副檔名 .exe 的檔案。）

# Src file & srcfile in Package

兩者有區別性：

*src, 指令原始程式檔案

宣告為屬於 main 程式套件，包含無參數宣告和結果宣告的 src file （但須先為套件初始化，利用 https://github.com/poupougo/Go_init_dependency/blob/master/README.md ），此種原始檔案可以利用 go run 指令獨立執行，也可以被 go build 或是 go install 指令轉為可執行檔案。

*srcInPackage, 函式庫原始程式檔案

存在某套件中的普通原始程式檔案。

# to verify Go Version

    ~$ go version
    go version go1.13.4 darwin/amd64
    
# print out Go enviroment

    ~$ go env
    
>>>

        GO111MODULE=""
        GOARCH="amd64"
        GOBIN=""
        GOCACHE="/Users/pintred/Library/Caches/go-build"
        GOENV="/Users/pintred/Library/Application Support/go/env"
        GOEXE=""
        GOFLAGS=""
        GOHOSTARCH="amd64"
        GOHOSTOS="darwin"
        GONOPROXY=""
        GONOSUMDB=""
        GOOS="darwin"
        GOPATH="/Users/pintred/go"
        GOPRIVATE=""
        GOPROXY="https://proxy.golang.org,direct"
        GOROOT="/usr/local/go"
        GOSUMDB="sum.golang.org"
        GOTMPDIR=""
        GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
        GCCGO="gccgo"
        AR="ar"
        CC="clang"
        CXX="clang++"
        CGO_ENABLED="1"
        GOMOD=""
        CGO_CFLAGS="-g -O2"
        CGO_CPPFLAGS=""
        CGO_CXXFLAGS="-g -O2"
        CGO_FFLAGS="-g -O2"
        CGO_LDFLAGS="-g -O2"
        PKG_CONFIG="pkg-config"
        GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/81/ncqxl7t95_363jgkl65d573m0000gn/T/go-build919953390=/tmp/go-build -gno-record-gcc-switches -fno-common"
  
# to build & run go file, and reformat Go src file

    ~$ go build <gofile>
    ~$ go run <gofile>
    
# add & install Dependencies to Package

    ~$ go get 
    ~$ go install

# Go CLI called "go help < topic >"
  
  topics includes :
  
    buildmode   build modes
    c           calling between Go and C
    cache       build and test caching
    testflag    testing flags
    testfunc    testing functions
    
    environment environment variables
    filetype    file types
    go.mod      the go.mod file
    goproxy     module proxy protocol
    
    
    packages    package lists and patterns
    importpath  import path syntax
    modules     modules, module versions, and more
    module-get  module-aware go get
    module-auth module authentication using go.sum
    module-private module configuration for non-public modules
   
   （ 版本1.13後，建議使用 GoMod，不建議繼續使用相互依賴的 GoPath ）
      https://github.com/poupougo/Go_init_dependency
   
    gopath      GOPATH environment variable
    gopath-get  legacy GOPATH go get
    
# Go CLI called "go < command > [arguments]"

    bug         start a bug report
    
    
    build       compile packages and dependencies
    
    
    clean       remove object files and cached files
    doc         show documentation for package or symbol
    
    
    env         print Go environment information
    fix         update packages to use new APIs
    
    
    fmt         gofmt (reformat) package sources
    generate    generate Go files by processing source
    
    
    get         add dependencies to current module and install them
    install     compile and install packages and dependencies
    
    
    list        list packages or modules
    mod         module maintenance
    run         compile and run Go program
    
    test        test packages
    tool        run specified go tool
    version     print Go version
    vet         report likely mistakes in packages
    

    
    
