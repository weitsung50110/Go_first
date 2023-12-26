### Gin 專案文件結構
通常，一個典型的 Gin 專案包含以下檔案和資料夾：

main.go：主應用程式檔案。 <br/>
go.mod 和 go.sum：Go 模組文件，用於管理專案的依賴。 <br/>
controllers/：存放控制器的資料夾。 <br/>
models/：存放模型或資料結構的資料夾。 <br/>
middlewares/：存放中間件的資料夾。 <br/>
static/：存放靜態檔案（如 CSS、JavaScript 等）的資料夾。 <br/>
templates/：存放 HTML 範本文件的資料夾。

### 建立新項目：
使用 go mod init 指令初始化新的 Go 項目，並安裝 Gin 框架：

- 如果您的公司網域為 example.com，且專案名稱為 myproject，您可以使用 example.com/myproject 作為模組名稱。

- 如果您正在開發個人項目，您可以使用您自己的 GitHub 使用者名稱作為模組名稱，例如 github.com/yourusername/myproject。
專案名稱也可以基於專案的功能或用途來命名，但要確保它是獨一無二的。

    go mod init example.com/myproject

- go get 是 Go 語言提供的指令，用於從遠端倉庫（例如 GitHub、Bitbucket 等）下載並安裝指定的 Go 套件或模組。


- -u 是 go get 指令的選項，表示更新已安裝的套件到最新版本。如果已經安裝了 Gin，使用 -u 選項會將其升級到最新的版本（如果有可用的更新）。


- github.com/gin-gonic/gin 是 Gin 框架在 GitHub 上的倉庫位址。這條指令告訴 Go 工具鏈從 GitHub 下載 Gin 框架的原始碼，並將其安裝到 Go 工作空間。

    go get -u github.com/gin-gonic/gin

### 執行 Gin 專案：
在命令列中執行以下命令來執行您的 Gin 專案：

    go run main.go

### 建置可執行檔：
使用 go build 指令建置可執行檔。這個過程會產生一個名為 myapp 的可執行文件，其中包含了您的 Gin 應用程式。<br/>
這個可執行檔可以在相同平台上直接運行，無需原始程式碼。

適應不同的作業系統：如果您的專案需要跨平台，例如在Windows 上使用.exe 副檔名，而在類別Unix 系統上沒有副檔名，那麼使用-o 標誌可以確保在建置時正確設定輸出檔案的名稱。

    go build -o mymain.exe main.go
