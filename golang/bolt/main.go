package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("test.db", 0600, nil)
	defer db.Close()

	if err != nil {
		log.Panic("打开数据库失败")
	}

	db.Update(func(tx *bolt.Tx) error {
		// 打开一个bucket
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			// 创建bucket
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建b1失败")
			}
		}

		bucket.Put([]byte("1111"), []byte("hello"))
		bucket.Put([]byte("2222"), []byte("world"))

		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		// 选择bucket
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("bucket b1 不应该为空，请检查！！！")
		}
		// 直接读取数据
		v1 := bucket.Get([]byte("1111"))
		v2 := bucket.Get([]byte("2222"))

		fmt.Println(string(v1))
		fmt.Println(string(v2))

		return nil
	})
}
