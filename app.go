package main

import (
	"context"
	"os"
	"log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	
	// 初始化数据库
	err := initDatabase()
	if err != nil {
		log.Printf("数据库初始化失败: %v", err)
	}
}

func (a *App) Close() {
	os.Exit(0)
}

// OpenURL 在前端调用的方法，用于打开网页
func (a *App) OpenURL(url string) {
	runtime.BrowserOpenURL(a.ctx, url)
}

// GetAllWebsites 获取所有网站列表
func (a *App) GetAllWebsites() ([]Website, error) {
	return GetAllWebsites()
}

// AddWebsite 添加新网站
func (a *App) AddWebsite(name, url string) error {
	return AddWebsite(name, url)
}

// DeleteWebsite 删除网站
func (a *App) DeleteWebsite(id int) error {
	return DeleteWebsite(id)
}

// NavigateToWebsite 导航到指定网站
func (a *App) NavigateToWebsite(url string) {
	// 使用Wails的WebView导航功能，在当前窗口中打开网站
	runtime.WindowExecJS(a.ctx, "window.location.href = '"+url+"'")
}

// GoBackHome 返回主页
func (a *App) GoBackHome() {
	// 重新加载主页（实际上在前端处理）
	// 这里可以添加额外的逻辑，如果需要的话
}
