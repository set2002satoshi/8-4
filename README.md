## このプロジェクトについて
CleanArchitecture元の私なりに解釈して作成してます。
勉強したばかりなので、間違っているところやアドバイスがありましたら,教えていただけると幸いです。

このプロジェクトを見ていただけるとわかりますが命名するのが苦手です。


## ファイル構成

ファイル構成を以下のようになっています。
```
root
├── Dockerfile
├── README.md
├── docker-compose.yml
├── go.mod
├── go.sum
├── infrastructure
│   ├── Config.go
│   ├── DB.go
│   └── Routing.go
├── interfaces
│   ├── controllers
│   │   ├── blog
│   │   │   ├── blog_controller.go
│   │   │   ├── blog_create_contoller.go
│   │   │   ├── blog_delete_contoller.go
│   │   │   ├── blog_find_all_contoller.go
│   │   │   ├── blog_findbyid_contoller.go
│   │   │   └── blog_update_contoller.go
│   │   ├── context.go
│   │   ├── h.go
│   │   └── user
│   │       ├── user_controller.go
│   │       ├── user_create_controller.go
│   │       ├── user_delete_controller.go
│   │       ├── user_findall_controller.go
│   │       ├── user_findbyid_controller.go
│   │       ├── user_login_controller.go
│   │       └── user_update_controller.go
│   └── database
│       ├── blog
│       │   └── blog_repository.go
│       ├── db.go
│       ├── db_repository.go
│       └── user
│           └── user_repository.go
├── main.go
├── models
│   ├── active_blog.go
│   ├── active_user.go
│   ├── history_blog.go
│   └── history_user.go
├── pkg
│   ├── aws
│   └── module
│       ├── db
│       ├── dto
│       │   ├── request
│       │   │   ├── blog_request.go
│       │   │   └── user_request.go
│       │   └── response
│       │       ├── blog_response.go
│       │       ├── response.go
│       │       └── user_response.go
│       ├── service
│       │   ├── authentication
│       │   │   ├── auth
│       │   │   │   ├── issue_token_service.go
│       │   │   │   └── password_and_hash_compare_service.go
│       │   │   ├── aws
│       │   │   └── dpo
│       │   │       └── convert_to_model_dpo.go
│       │   └── image
│       └── temporary
│           └── data_type.go
└── usecase
    ├── ResultStatus.go
    ├── blog
    │   ├── blog_interactor.go
    │   └── blog_repository.go
    ├── db_repository.go
    └── user
        ├── user_interactor.go
        └── user_repository.go
```


## pkg/module/dto/response構成、解説

controller層にてapi利用者に返すデータ構成を以下の用にする
レスポンスの構造は三層に分けて定義する。

### 1層目
1層目は作成したAPIの機能に応じて１層目を作成する、なので、apiの機能量に応じて定義するstructの量も依存する
```
type (
    XXXResponse struct {
        Result 　xxxResult

        CodeErr  error `json:"code_err"`
		MsgErr   string `json:"msg"`
    }
)
```
### 2層目
2層目は作るentityの量に依存し、作られるケースが多い
例えば、単数形と複数形を作成を行う
```
type (
    ActiveXXXResult struct {
        XXX *ActiveXXXEntity
        // 単数形　or 複数形
        XXX []*ActiveXXXEntity
    }
)
```
### 3層目
３層目は機能に応じて返したいデータ型を定義する。
```
type (
    ActiveXXXEntity struct {
        XXX int     `json:"xxx"`
        XXX string  `json:"xxx"
        History HistoryOOOEntity
    }
    HistoryOOOEntity struct {
        OOO int     `json:"ooo"`
    }
)
```


