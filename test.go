package main

import (
	"fmt"
	lichv "github.com/lichv/go"
)

func main() {
	db := lichv.InitPostgreDriver("127.0.0.1",5432,"postgre","123456","aliyun")
	query, err := db.Query("SELECT * FROM \"aliyun_config\"  WHERE \"aliyun_config\".\"deleted_at\" IS NULL AND ((code= $1)) ORDER BY \"aliyun_config\".\"id\" ASC LIMIT 1", "dg54df3g")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(query)

}
