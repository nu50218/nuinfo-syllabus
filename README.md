# nuinfo-syllabus

名大情報のシラバスのAPIサーバーのDockerイメージ(キャッシュしてくれるよ)

## 使い方

`$ docker run -p 8080:8080 nu50218/nuinfo-syllabus`

`http://localhost:8080/subjects`で確認できます。

| 環境変数     | 意味                    | デフォルト                                |
| -- | -- | -- |
| ENDPOINT | クロールするエンドポイント         | https://syllabus.i.nagoya-u.ac.jp/i/ |
| EXPIRES  | キャッシュを保持する時間          | 1h                                   |
| INTERVAL | クロールするときにリクエストごとに待つ時間 | 500ms                                     |

時間はGoの[time.ParseDuration](https://golang.org/pkg/time/#ParseDuration)でパースできる文字列を指定してください。

ちなみにキャッシュはオンメモリなので、終了すると揮発します。

## API

### GET /subjects

こんなかんじ

```json
[
    {
        "timetable_code": "1000010",
        "course_title": "インフォマティックス1",
        "semester": "春2期",
        "day_and_period": "金曜 3限目",
        "grade": "1年",
        "credits": "必修1",
        "update": "2020-02-18"
    },
    {
        "timetable_code": "1000011",
        "course_title": "情報セキュリティとリテラシー1",
        "semester": "春1期",
        "day_and_period": "月曜 1限目",
        "grade": "1年",
        "credits": "必修1",
        "update": "2020-02-27"
    },
]
```

### GET /subjects/{時間割コード}

先生の名前とかでてるので載せるのアレかなって思ったので自分で確認してね
