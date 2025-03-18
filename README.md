# データ可視化ツール 

CSVファイルからデータを読み込み、見やすいグラフを生成します。


## ディレクトリ構造

```
│── main.go               // エントリーポイント
│── csv/
│   ├── reader.go         // CSV読み込みロジック
│   ├── parser.go         // CSVデータのパース
├── data
│   └── sample.csv        // サンプルCSVファイル
├── go.mod                // Goモジュールファイル
├── go.sum                // Go依存関係ファイル
│── graph/
│   ├── plotter.go        // PNG画像の生成
│── models/
│   ├── data.go           // CSVデータの構造体
│── output/               // 生成されたPNG画像を保存するディレクトリ

```

## 前提条件

- Go 1.23以上

## インストール

依存関係をインストール
```bash
go mod tidy
```

## 使い方

### 基本的な使用方法

```bash
go run . path/to/your/data.csv
```

または、ビルドして実行:

```bash
go build -o csv_to_png
./csv_to_png path/to/your/data.csv
```

### CSVファイルの形式

CSVファイルは以下の形式である必要があります:

```csv
日付, 金額
2024/01/15,71.53
2024/02/20,42.75
2024/03/10,23.40
```

- 1行目はヘッダー行です
- 2行目以降がデータ行です
- 日付は「YYYY/MM/DD」形式で指定してください
