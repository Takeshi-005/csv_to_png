package main

import (
	"dividend-visualizer/csv"
	"dividend-visualizer/graph"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("使用方法: %s path/to/your/data.csv", os.Args[0])
	}

	// CSVファイルのパス
	csvFilePath := os.Args[1]
	outputImagePath := "output/graph.png"

	// CSVファイルを開く
	records, err := csv.ReadCSV(csvFilePath)
	if err != nil {
		log.Fatalf("CSV読み込みエラー: %v", err)
	}

	fmt.Printf("CSVファイルを読み込みました: %s\n", records)
	// CSVデータを解析
	dataSet, err := csv.ParseCSV(records)
	if err != nil {
		log.Fatalf("CSV解析エラー: %v", err)
	}

	// 出力フォルダを作成（存在しない場合）
	if err := os.MkdirAll("output", os.ModePerm); err != nil {
		log.Fatalf("出力フォルダ作成エラー: %v", err)
	}

	fmt.Println("データセットを読み込みました:", dataSet)
	// PNG画像を生成
	if err := graph.PlotGraph(dataSet, outputImagePath); err != nil {
		log.Fatalf("グラフ生成エラー: %v", err)
	}

	fmt.Println("グラフ画像を生成しました:", outputImagePath)
}
