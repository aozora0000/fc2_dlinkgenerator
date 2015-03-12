# FC2動画用ダウンロードURLジェネレーター

FC2動画のダウンロードURLをjson形式で取得します。

## Usage

### Server
```
go run ./server.go -p 8080

or

go build ./server.go
./server -p 8080
```

#### params

- -p or --port : ポート番号(defailt:8080)

### Client

```
% curl -sS http://localhost/?mid=20150309euLnarHF
{"Iserror":false,"Path":"http://vip.cvideocache21.fc2.com/videocache/up/flv/201503/09/e/20150309euLnarHF.flv?mid=be372a8ee9f343a4843b2ae084d04efd","Title":"(無)ハメてハメてこの世界で愛することも出来（ｒｙ","Isadult":true,"Payment":false,"Length":4167}
```

#### params

- mid : 動画ID(16桁の英数字)

## Sample

```
% curl -sS http://fc2.aozora0000.biz/?mid=20150309euLnarHF
```
