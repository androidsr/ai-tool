package main

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type Website struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB

func initDatabase() error {
	var err error
	db, err = sql.Open("sqlite", "websites.db")
	if err != nil {
		return err
	}

	// 创建网站表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS websites (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		url TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	// 插入示例数据（如果表为空）
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM websites").Scan(&count)
	if err == nil && count == 0 {
		//写入一条默认数据
		_, err = db.Exec("INSERT INTO websites (name, url) VALUES (?, ?)", "示例网站", "https://example.com")
		if err != nil {
			return err
		}
	}
	return nil
}

// CloseDatabase 关闭数据库连接
func CloseDatabase() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// GetAllWebsites 获取所有网站
func GetAllWebsites() ([]Website, error) {
	rows, err := db.Query("SELECT id, name, url, created_at FROM websites ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var websites []Website
	for rows.Next() {
		var website Website
		err := rows.Scan(&website.ID, &website.Name, &website.URL, &website.CreatedAt)
		if err != nil {
			return nil, err
		}
		websites = append(websites, website)
	}

	return websites, nil
}

// AddWebsite 添加新网站
func AddWebsite(name, url string) error {
	_, err := db.Exec("INSERT INTO websites (name, url) VALUES (?, ?)", name, url)
	return err
}

// DeleteWebsite 删除网站
func DeleteWebsite(id int) error {
	_, err := db.Exec("DELETE FROM websites WHERE id = ?", id)
	return err
}

// GetWebsiteByID 根据ID获取网站
func GetWebsiteByID(id int) (*Website, error) {
	var website Website
	err := db.QueryRow("SELECT id, name, url, created_at FROM websites WHERE id = ?", id).
		Scan(&website.ID, &website.Name, &website.URL, &website.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &website, nil
}
