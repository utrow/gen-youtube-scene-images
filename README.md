# gen-youtube-scene-images

YouTube動画をシーン検出して画像に出力するツール

## 動作環境
- Go 1.12
- ffmpeg 4.0.2
    - brewでインストールできる(mac)

## 実行
```
$ go run main.go

...

YouTube Video URL: { ENTER YOUTUBE VIDEO URL } 
```

## 出力
- videos/
    - ダウンロードされたYouTube動画のMP4ファイルが保存される
- output/
    - {連番}.jpg
        - シーンのキャプチャ画像
        - ffout.log
            - 処理のログ
            - シーンの時間などが取得できる
