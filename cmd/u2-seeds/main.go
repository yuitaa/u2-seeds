package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/yuitaa/u2-seeds/internal/capture"
	"github.com/yuitaa/u2-seeds/internal/mapping"
	"github.com/yuitaa/u2-seeds/internal/move"
	"github.com/yuitaa/u2-seeds/internal/seed"
)

func main() {
	var engineFlag string
	var versionFlag string
	flag.StringVar(&engineFlag, "engine", "", "Engine name (required)")
	flag.StringVar(&versionFlag, "version", "", "Version string (required)")
	flag.Parse()

	if engineFlag == "" || versionFlag == "" {
		log.Fatal("エラー: -engine と -version フラグは必須です。例: u2-seeds -engine <engine_name> -version <version>")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	var shutdownRequested bool
	go func() {
		<-sigChan
		shutdownRequested = true

		<-sigChan
		log.Println("強制終了します。")
		os.Exit(1)
	}()

	// マッピング用レジストリの初期化
	registry, err := mapping.NewRegistry("mapping.json")
	if err != nil {
		log.Fatalf("マッピングレジストリの初期化に失敗しました: %v", err)
	}

	// 起動時のフラグ情報をマッピングIDに変換
	engineID, err := registry.GetID("engines", engineFlag)
	if err != nil {
		log.Fatalf("engineのマッピングに失敗しました: %v", err)
	}
	versionID, err := registry.GetID("versions", versionFlag)
	if err != nil {
		log.Fatalf("versionのマッピングに失敗しました: %v", err)
	}

	seedGenerator := seed.NewSeedGenerator()
	move.PressAltTab()

	fileName := "seeds.csv"
	fileExist := true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fileExist = false
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("ファイルのオープンに失敗しました: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if !fileExist {
		header := []string{"Seed", "UpperTask", "UpperExtension", "LowerTask", "LowerExtension", "Wagons", "Engine", "Version"}
		if err := writer.Write(header); err != nil {
			log.Fatalf("ヘッダーの書き込みに失敗しました: %v", err)
		}
		writer.Flush()
	}

	for {
		if shutdownRequested {
			return
		}

		s := seedGenerator.Generate()
		log.Printf("Seed: %s", s)

		move.RestartMove(s)

		info, ok := capture.GetInformation()
		log.Println("\n" + info.String())

		if !ok {
			continue
		}

		upperTaskID, err := registry.GetID("tasks", info.UpperTask)
		if err != nil {
			continue
		}
		upperExtensionID, err := registry.GetID("extensions", info.UpperExtension)
		if err != nil {
			continue
		}
		lowerTaskID, err := registry.GetID("tasks", info.LowerTask)
		if err != nil {
			continue
		}
		lowerExtensionID, err := registry.GetID("extensions", info.LowerExtension)
		if err != nil {
			continue
		}

		var wagonIDs []string
		var wagonMappingError bool
		for _, w := range info.Wagons {
			id, err := registry.GetID("wagons", w)
			if err != nil {
				wagonMappingError = true
				break
			}
			wagonIDs = append(wagonIDs, strconv.Itoa(id))
		}
		if wagonMappingError {
			continue
		}

		// CSVに書き込むレコードを作成
		record := []string{
			string(s),
			strconv.Itoa(upperTaskID),
			strconv.Itoa(upperExtensionID),
			strconv.Itoa(lowerTaskID),
			strconv.Itoa(lowerExtensionID),
			strings.Join(wagonIDs, ";"),
			strconv.Itoa(engineID),
			strconv.Itoa(versionID),
		}

		if err := writer.Write(record); err != nil {
		} else {
			writer.Flush()
			if err := writer.Error(); err != nil {
			}
		}
	}
}
