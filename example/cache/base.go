/*
 * Created on Tue Mar 04 2025
 *
 * © 2025 Nevilsoft Ltd., Part. All Rights Reserved.
 *
 * * ข้อมูลลับและสงวนสิทธิ์ *
 * ไฟล์นี้เป็นทรัพย์สินของ Nevilsoft Ltd., Part. และมีข้อมูลที่เป็นความลับทางธุรกิจ
 * อนุญาตให้เฉพาะพนักงานที่ได้รับสิทธิ์เข้าถึงเท่านั้น
 * ห้ามเผยแพร่ คัดลอก ดัดแปลง หรือใช้งานโดยไม่ได้รับอนุญาตจากฝ่ายบริหาร
 *
 * การละเมิดข้อตกลงนี้ อาจมีผลให้ถูกลงโทษทางวินัย รวมถึงการดำเนินคดีตามกฎหมาย
 * ตามพระราชบัญญัติว่าด้วยการกระทำความผิดเกี่ยวกับคอมพิวเตอร์ พ.ศ. 2560 (มาตรา 7, 9, 10)
 * และกฎหมายอื่นที่เกี่ยวข้อง
 */

package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/burapha44/example/config"
)

var (
	client *redis.Client
)

func Init() error {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Conf.RedisHost, config.Conf.RedisPort),
		Password:     config.Conf.RedisPassword,
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ทดสอบการเชื่อมต่อ
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
		return err
	}

	log.Println("💎 Redis connected successfully")

	return nil
}

func GetCacheClient() *redis.Client {
	return client
}

func Close() {
	if client != nil {
		_ = client.Close()
	}
}
