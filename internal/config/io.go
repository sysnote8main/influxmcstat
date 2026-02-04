package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func LoadConfig(path string) *Config {
	// ファイルの存在確認
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// ファイルがないので作成する
		f, err := os.Create(path)
		if err != nil {
			log.Fatalf("設定ファイルの作成に失敗しました: %v", err)
		}
		defer f.Close()

		// デフォルト値をTOML形式で書き込む
		conf := NewConfig()
		if err := toml.NewEncoder(f).Encode(conf); err != nil {
			log.Fatalf("デフォルト値の書き込みに失敗しました: %v", err)
		}

		// メッセージを出して終了
		log.Printf("設定ファイル '%s' が見つからなかったため、デフォルト設定で作成しました。", path)
		log.Fatal("設定内容を確認・編集してから、再度アプリを起動してください。")
	}

	var conf Config
	// ファイルが存在する場合は読み込む
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		log.Fatalf("設定ファイルの読み込みに失敗しました: %v", err)
	}

	return &conf
}
