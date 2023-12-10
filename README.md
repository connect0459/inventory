# inventory_control

Echo(v4) + GORM + MySQL で構築された RESTful API サーバーです。DBeaver があればデータ管理が楽なので併せてセッティングしましょう。

## ローカル環境へのインストール

コマンド実行においては、特段断りがない限り、プロジェクトのルートディレクトリで実行するものとします。

### 1. Git からプロジェクトをクローン

```bash
git clone path/to/inventory_control.git
```

### 2. コマンドを実行して依存関係をインストール

```bash
RUN go mod download && go mod verify
```

### 3. `config.json` ファイルを作成

プロジェクトのルートに `config.json` ファイルを作成し、必要な設定を記述します。以下は `config.json` ファイルの例です。

```json
{
  "DevMode": true,
  "DSN": {
    "Dev": "user:password@tcp(host:port)/dbname",
    "Production": ""
  }
}
```

## CI/CD 設定(任意)
### GitLab
GitLabでCI/CDを行う場合、 `.gitlab-ci.yml` をプロジェクトのルートに設置してください。以下は `.gitlab-ci.yml` の例です。
```yml:.gitlab-ci.yml
---
stages:
- build
- test
- deploy
include:
- remote: https://gitlab.com/gitlab-org/incubation-engineering/five-minute-production/library/-/raw/main/gcp/cloud-run.gitlab-ci.yml
```

### Dockerfile・docker-compose.ymlの作成
CI/CD で使用する本番環境用の Dockerfile はプロジェクトのルートに設置し、開発環境用のものは `build/package/go` などに設置してください。

以下は `docker-compose.yml` の例です。Go と MySQL にそれぞれ Dockerfile を作成し、 `app` `db` の2つのコンテナを連立させています。

```yml:docker-compose.yml
version: "3.8" # docker-composeのバージョン
services: # services配下に各コンテナの情報を記載する
  app: # golangのコンテナ
    container_name: app
    build: # Dockerfileのあるディレクトリのパスを指定する
      context: . # 起点となるディレクトリ
      dockerfile: ./build/package/go/Dockerfile
    volumes: # マウント元と先を指定
      - .:/go/src/app
    environment:
      - TZ=Asia/Tokyo
      - OPENAPI_SPEC=./api/api-docs.yaml
    tty: true # コンテナの永続化
    ports: # 外部に公開するポートを指定する
      - 7000:7000
  db:
    container_name: db
    # image: mysql:8.0
    build: # Dockerfileのあるディレクトリのパスを指定する
      context: . # 起点となるディレクトリ
      dockerfile: ./build/package/mysql/Dockerfile
    environment:
      MYSQL_ROOT_PASSWORD: fortwo42omu
      MYSQL_DATABASE: inventory_control
    volumes:
      - ./build/package/mysql/data:/var/lib/mysql
      - ./build/package/mysql/init:/docker-entrypoint-initdb.d
      - ./build/package/mysql/log:/var/log/mysql
      - ./build/package/mysql/conf.d:/etc/mysql/conf.d
    ports: # 外部に公開するポートを指定する
      - 3309:3309
    tty: true
    networks:
      - default
```

## API エンドポイント（URI）の追加
### Model・Migration・Controller の作成

作成する Model 名を `ModelName` とすると、 `model/modelname.go` としてファイルを作成する。ファイルはスネークケースで命名すること。以下は `bookinfo.go` の例である。

```go:bookinfo.go
package model

import "time"

type BookInfo struct {
    ID            int       `json:"id"`
    Title         string    `json:"title"`
    Author        string    `json:"author"`
    Publisher     string    `json:"publisher"`
    Description   string    `json:"description"`
    TypeGenreID   int       `json:"type_genre_id"`
    ListPrice     int       `json:"list_price"`
    SalePrice     int       `json:"sale_price"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}
func (b *BookInfo) TableName() string {
	return "books_info"
}
```

Migration（マイグレーション）ファイルとは、コマンドラインからデータベースの作成を行う際に参照される定義ファイルです。このプロジェクトでは `database/database.go` をマイグレーションファイルとする。AutoMigrate に構造体を渡すことでマイグレーションが設定できる。なお、GORM のマイグレーションは一度作成したテーブルに対する変更はできない。再構築したい場合は、データベースを削除する必要がある。

```go:database.go
// database/database.go
package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"inventory_control/config"
	"inventory_control/model"
	"log"
	"strings"
)

var DB *gorm.DB
var err error

func init() {
	var dsn string
	if config.AppConfig.DevMode {
		dsn = config.AppConfig.DSN.Dev
	} else {
		dsn = config.AppConfig.DSN.Production
	}
	if strings.HasPrefix(dsn, "DSN=") {
		log.Fatalln("DSN are not expected to contain the 'DSN=' prefix. Please correct the value in config.json")
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + " database can't connect")
	}

	err = DB.AutoMigrate(
		&model.TypeBranch{},
		&model.TypeGenre{},
		&model.BookInfo{},
		&model.BookStock{})
	if err != nil {
		log.Fatalln("Failed to auto-migrate tables")
	}
}
```

Controller を作成する場合は、 `controller/modelname_controller.go` として作成してください。

```go:bookinfo_controller.go
package controller

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/labstack/echo/v4"
	"inventory_control/database"
	"inventory_control/model"
)

func GetBookInfoAll(c echo.Context) error {
	bookInfo := []model.BookInfo{}
	database.DB.Find(&bookInfo)
	return c.JSON(http.StatusOK, bookInfo)
}

func GetBookInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookInfo := model.BookInfo{}
	database.DB.First(&bookInfo, id)
	return c.JSON(http.StatusOK, bookInfo)
}

func CreateBookInfo(c echo.Context) error {
	bookInfo := model.BookInfo{}
	c.Bind(&bookInfo)
	database.DB.Save(&bookInfo)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/books_info/%d", bookInfo.ID))
	return c.NoContent(http.StatusCreated)
}

func UpdateBookInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookInfo := model.BookInfo{}
	database.DB.First(&bookInfo, id)
	c.Bind(&bookInfo)
	database.DB.Save(&bookInfo)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/books_info/%d", bookInfo.ID))
	return c.NoContent(http.StatusNoContent)
}

func DeleteBookInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookInfo := model.BookInfo{}
	database.DB.Delete(&bookInfo, id)
	return c.NoContent(http.StatusNoContent)
}
```

### Migration と Route 設定

Migration は DB 関数が初期化(init)されたとき、ここでは `main.go` で呼び出したときに実行されます。 `main.go` はサーバーの起動を担当しています。

```bash
go run main.go
```

次に、ルーティング設定の確認です。 `route/route.go` にコントローラーファイルと URI の結び付けを記述してエンドポイントを作成します。echo-swagger の設定はこの後に記述します。

```go:route.go
package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"inventory_control/config"
	"inventory_control/controller"
	_ "inventory_control/docs/v1"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	v1 := e.Group("/api/v1")
	{
		// 開発モード時のみSwaggerドキュメントを表示
		if config.AppConfig.DevMode {
			v1.GET("/swagger/*", echoSwagger.WrapHandler)
			// ドキュメントの生成は"swag init -o ./docs/v1"
			v1.File("/swagger/doc.json", "./docs/v1/swagger.json")
			v1.File("/swagger/doc.yaml", "./docs/v1/swagger.yaml")
		}

		// books_info
		v1.GET("/books_info", controller.GetBookInfos)
		v1.GET("/books_info/:id", controller.GetBookInfo)
		v1.POST("/books_info", controller.CreateBookInfo)
		v1.PUT("/books_info/:id", controller.UpdateBookInfo)
		v1.DELETE("/books_info/:id", controller.DeleteBookInfo)
	}

	return e
}
```

設定が確認出来たら、`curl`コマンドでエンドポイントが機能するか確認しましょう。

### curl でサーバーにリクエストを送信

#### 1. ローカルサーバーの起動

次のコマンドを実行して、Echo の開発サーバーを立ち上げましょう。ポートは`7000`です。

```bash
go run main.go
```

## API 仕様書の生成

### swaggo、echo-swagger のインストール

次のコマンドで swaggo と echo-swagger をインストールします。

```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/echo-swagger
```

### 共通スキーマの定義

`main.go` に、共通の設定を記述しましょう。

```go:main.go
//	@title			inventory_control API
//	@version		1.0
//	@description	在庫管理を行うAPIサーバー
//	@host			localhost:7000
//	@BasePath		/api/v1
func main() {
	// ...
}
```

Model を Schema として登録します。以下の例だと、`@OA\Schema(ref="#/components/schemas/BookInfo")`で Schema が読み込めるようになります。

```go:bookinfo.go
// model/BookInfo.go
package model

import "time"

// @swagger:response BookInfo
type BookInfo struct {
	// ID
	// in: body
	// required: true
	ID int `json:"id" example:"1"`
	// ISBN
	// in: body
	// required: true
	ISBN string `json:"isbn" example:"9784908434266"`
	// 書籍のタイトル
	// in: body
	// required: true
	Title string `json:"title" example:"これだけは知っておこう!情報リテラシー"`
	// 書籍の著者
	// in: body
	// required: true
	Author string `json:"author" example:"noa出版"`
	// 書籍の出版社
	// in: body
	// required: true
	Publisher string `json:"publisher" example:"noa出版"`
	// 書籍の説明
	// in: body
	// required: false
	Description string `json:"description" example:"ハードウェア、ソフトウェア、ネットワーク、セキュリティ、情報モラルまで知っておきたい基礎知識が1冊に分かりやすくまとまっているテキストです。"`
	// ジャンルID
	// in: body
	// required: true
	TypeGenreId int `json:"type_genre_id" example:"1"`
	// 書籍の定価
	// in: body
	// required: true
	// 
	ListPrice int `json:"list_price" example:"1000"`
	// レコードの作成日時
	// in: body
	// required: true
	CreatedAt time.Time `json:"created_at" example:"2023-06-20T21:20:14.000000Z"`
	// レコードの更新日時
	// in: body
	// required: true
	UpdatedAt time.Time `json:"updated_at" example:"2023-06-20T21:20:14.000000Z"`
}

// TableName sets the insert table name for this struct type
func (b *BookInfo) TableName() string {
	return "books_info"
}
```

そして、Controller ファイルにメソッドの説明を記述します。

```go:bookinfo_controller.go
// controller/bookinfo_controller.go
package controller

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/labstack/echo/v4"
	"inventory_control/database"
	"inventory_control/model"
)

//	@Summary		全ての書籍情報を取得
//	@Description	全ての書籍情報を取得する
//	@Tags			books_info
//	@Produce		json
//	@Success		200	{array}	model.BookInfo
//	@Router			/books [get]
func GetBookInfoAll(c echo.Context) error {
	bookInfo := []model.BookInfo{}
	database.DB.Find(&bookInfo)
	return c.JSON(http.StatusOK, bookInfo)
}

//	@Summary		書籍情報を取得
//	@Description	IDによる書籍情報の取得
//	@Tags			books_info
//	@Produce		json
//	@Param			id	path		int	true	"Book ID"
//	@Success		200	{object}	model.BookInfo
//	@Router			/books/{id} [get]
func GetBookInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookInfo := model.BookInfo{}
	database.DB.First(&bookInfo, id)
	return c.JSON(http.StatusOK, bookInfo)
}

//	@Summary		書籍情報を作成
//	@Description	新しい書籍情報の作成
//	@Tags			books_info
//	@Accept			json
//	@Produce		json
//	@Param			bookInfo	body		model.BookInfo	true	"Book Info"
//	@Success		200			{object}	model.BookInfo
//	@Router			/books [post]
func CreateBookInfo(c echo.Context) error {
	bookInfo := model.BookInfo{}
	c.Bind(&bookInfo)
	database.DB.Save(&bookInfo)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/books_info/%d", bookInfo.ID))
	return c.NoContent(http.StatusCreated)
}

//	@Summary		書籍情報を更新
//	@Description	IDによる書籍情報の更新
//	@Tags			books_info
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int				true	"Book ID"
//	@Param			bookInfo	body		model.BookInfo	true	"Book Info"
//	@Success		200			{object}	model.BookInfo
//	@Router			/books/{id} [put]
func UpdateBookInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookInfo := model.BookInfo{}
	database.DB.First(&bookInfo, id)
	c.Bind(&bookInfo)
	database.DB.Save(&bookInfo)
	c.Response().Header().Set(echo.HeaderLocation, fmt.Sprintf("/books_info/%d", bookInfo.ID))
	return c.NoContent(http.StatusNoContent)
}

//	@Summary		書籍情報を削除
//	@Description	IDによる書籍情報の削除
//	@Tags			books_info
//	@Produce		json
//	@Param			id	path		int		true	"Book ID"
//	@Success		200	{string}	string	"BookInfo is deleted"
//	@Router			/books/{id} [delete]
func DeleteBookInfo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookInfo := model.BookInfo{}
	database.DB.Delete(&bookInfo, id)
	return c.NoContent(http.StatusNoContent)
}

```

### ドキュメントの生成

次のコマンドでドキュメントを生成します。

```bash
swag init -o ./docs/v1
```

指定したディレクトリ、ここでは `./docs/v1` に json と yaml が生成されています。 `http://localhost:7000/api/v1/swagger/index.html` から生成されたドキュメントを確認できます。

## 著者

[connect0459 (Akira Nakaoka)](https://github.com/connect0459)

