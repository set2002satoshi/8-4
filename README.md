# このプロジェクトについて
CleanArchitecture元の私なりに解釈して作成してます。
勉強したばかりなので、間違っているところやアドバイスがありましたら,教えていただけると幸いです。




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


