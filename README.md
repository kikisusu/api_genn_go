# api_genn_go

## 機能概要

APIを複数実行する際のオーケストレートをFactory method パターンを活用し簡易化したもの。  
APIを増やす際に、APIの基本的な処理は実装不要。  

以下はサンプルで用意しているAPI

### サンプル①

http://localhost:3001/marvelbuzz

APIをコールすると、裏で2APIを実行し、両方のレスポンスデータから一部のデータを取り出しレスポンス
* https://whenisthenextmcufilm.com/api
* https://corporatebs-generator.sameerkumar.website/

### サンプル②

http://localhost:3001/activeuuid

APIをコールすると、裏で2APIを実行し、両方のレスポンスデータから一部のデータを取り出しレスポンス
  * https://www.uuidtools.com/api/generate/v1
  * https://www.boredapi.com/api/activity/