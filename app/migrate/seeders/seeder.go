package main

import (
	"fmt"
	"go/db"
	"io/fs"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// データベースに接続
	dbConn := db.Initialize()
	defer db.CloseDB(dbConn) // DBを閉じることを忘れないようにする

	// 'seeders' ディレクトリのファイルをリスト
	files, err := os.ReadDir("migrate/seeders")
	if err != nil {
		log.Fatalf("Error reading seeders directory: %v", err)
	}

	// .sqlファイルだけを取得してソート
	var sqlFiles []fs.DirEntry
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			sqlFiles = append(sqlFiles, file)
		}
	}
	sort.Slice(sqlFiles, func(i, j int) bool {
		return sqlFiles[i].Name() < sqlFiles[j].Name()
	})

	// 各SQLファイルを適用
	for _, file := range sqlFiles {
		data, err := os.ReadFile("migrate/seeders/" + file.Name())
		if err != nil {
			log.Fatalf("Error reading file %s: %v", file.Name(), err)
		}
	
		// SQLファイルの内容をセミコロンで分割
		statements := strings.Split(string(data), ";")
	
		// 各SQL文を個別に実行
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)  // 余分な空白や改行を削除
			if stmt != "" {
				_, err = dbConn.Exec(stmt)
				if err != nil {
					log.Fatalf("Error executing statement from file %s: %v", file.Name(), err)
				}
			}
		}
		fmt.Printf("Applied seeder from file %s successfully!\n", file.Name())
	}

	fmt.Println("seeder applied successfully!")
}
