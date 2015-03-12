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

### Response
| param |  type  | context  |
|:-----:|:------:|:--------:|
| Iserror | bool | エラーかどうか |
| Path | string | ダウンロードパス |
| Title | string | 動画タイトル |
| Isadult | bool | エロかどうか |
| Payment | bool | 会員限定かどうか |
| Length | integer | 動画の長さ(second) |

## Sample

Crossdomain制約を外しているので、ブラウザからでも動作する・・・ハズ

```
% curl -sS http://fc2.aozora0000.biz/?mid=20150309euLnarHF
```

```
<html>
    <head>
        <meta charset="UTF-8">
    </head>
    <body>
        <a></a>
    </body>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.3/jquery.min.js"></script>
    <script>
        $(document).ready(function() {
            $.ajax({
                type: "GET",
                url: "http://fc2.aozora0000.biz/",
                data: {
                    mid: "20150309euLnarHF"
                }
            }).done(function(data){
                $("body > a").attr("href",data.Path);
                $("body > a").html(data.Title);
            }).fail(function(data){
                alert('error!!!');
            });
        });
    </script>
</html>
```

## License
MIT License
自由に使って頂いても結構です。
ご利用の場合、いかなる責任も当方は負いません。
