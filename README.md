# AI-TOOL - 企业网站书签管理工具

[![Go Version](https://img.shields.io/badge/Go-1.25.2-blue)](https://golang.org)
[![Wails](https://img.shields.io/badge/Wails-v2.11.0-green)](https://wails.io)
[![Vue](https://img.shields.io/badge/Vue-3.x-brightgreen)](https://vuejs.org)

## 项目概述

AI-TOOL 是一款基于 Wails + Vue.js + SQLite 开发的现代化企业网站书签管理工具。该软件提供直观的用户界面，帮助企业用户高效管理常用网站链接，支持快速访问、内嵌浏览和一键管理功能。

## ✨ 核心特性

### 🔖 智能书签管理
- **可视化网站管理**：清晰展示所有收藏网站
- **快速添加/删除**：支持添加新网站和删除不需要的网站
- **自动数据持久化**：使用 SQLite 数据库自动保存所有数据
- **创建时间记录**：自动记录每个网站的添加时间

### 🌐 内嵌浏览体验
- **集成化浏览**：网站直接在应用内部以内嵌方式打开
- **完整浏览器功能**：支持全屏内嵌浏览体验
- **便捷导航**：顶部导航栏显示网站信息，一键返回主页

### ⌨️ 便捷操作支持
- **快捷键支持**：Ctrl+D 快速添加新网站
- **响应式设计**：优雅的界面适配不同屏幕尺寸
- **一键访问**：点击即可快速打开目标网站

### 💾 数据安全可靠
- **本地数据库**：使用 SQLite 确保数据安全存储
- **自动初始化**：首次运行自动创建示例数据
- **数据完整性**：确保所有操作的数据一致性

## 🚀 快速开始

### 系统要求
- Windows 7/8/10/11（64位）
- 无需安装额外运行环境

### 安装使用
1. 下载最新版本的 `AI-TOOL.exe` 文件
2. 双击运行即可开始使用
3. 无需安装，绿色便携

### 基础操作
- **添加网站**：点击"添加网站"按钮或使用 Ctrl+D 快捷键
- **访问网站**：点击网站卡片上的"访问"按钮
- **删除网站**：点击网站卡片上的"删除"按钮
- **内嵌浏览**：网站会在应用内部打开，支持全屏浏览

## 🛠️ 技术架构

### 后端技术栈
- **Go 1.25.2**：高性能后端语言
- **Wails v2.11.0**：现代化桌面应用框架
- **SQLite**：轻量级数据库引擎
- **原生系统API**：文件系统操作和窗口管理

### 前端技术栈
- **Vue.js 3.x**：现代化前端框架
- **Vite**：快速构建工具
- **CSS3**：现代化样式设计

### 核心功能模块
```go
// 数据库操作
type Website struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    URL       string    `json:"url"`
    CreatedAt time.Time `json:"created_at"`
}

// 主要API方法
- GetAllWebsites()   // 获取所有网站
- AddWebsite()       // 添加新网站
- DeleteWebsite()    // 删除网站
- NavigateToWebsite() // 内嵌浏览网站
```

## 📁 项目结构

```
AI-TOOL/
├── main.go                 # 应用入口文件
├── app.go                 # 主应用逻辑和API接口
├── database.go           # 数据库操作模块
├── go.mod                # Go 模块配置
├── go.sum               # 依赖版本锁定
├── wails.json           # Wails 配置文件
├── frontend/            # 前端代码目录
│   ├── src/
│   │   ├── App.vue      # 主应用Vue组件
│   │   ├── main.js      # Vue 入口文件
│   │   └── assets/      # 静态资源
│   └── package.json     # 前端依赖配置
├── websites.db          # SQLite数据库文件（自动生成）
└── build/               # 构建输出目录
```

## 🔧 开发指南

### 环境搭建
```bash
# 克隆项目
git clone <repository-url>
cd AI-TOOL

# 安装Go依赖
go mod tidy

# 安装前端依赖
cd frontend && npm install

# 开发模式运行
wails dev

# 构建发布版本
wails build
```

### 数据库配置
应用首次运行时会自动创建 `websites.db` 数据库文件，并初始化示例数据。

### 功能扩展
- **添加分类功能**：可以在数据库中添加分类字段
- **搜索功能**：实现网站名称和URL的搜索
- **导入导出**：支持书签数据的导入导出
- **标签系统**：为网站添加标签进行分类

## 📊 功能演示

### 主界面
- 显示所有收藏的网站列表
- 每个网站显示名称、URL和创建时间
- 提供访问和删除操作按钮

### 添加网站
- 弹出表单输入网站名称和URL
- 数据验证确保URL格式正确
- 自动保存到数据库

### 内嵌浏览
- 全屏内嵌iframe显示目标网站
- 顶部导航栏显示网站信息
- 一键返回应用主页

## 🤝 贡献指南

欢迎提交 Issue 和 Pull Request 来完善这个项目！

## 📞 联系我们

- **项目维护者**：sirui
- **技术支持邮箱**：srandroid@163.com

## 💰支持与打赏
- 🙏** 感谢土豪大哥的打赏 **
- <img width="500" height="475" alt="image" src="https://github.com/user-attachments/assets/f13e8645-95cd-4e00-98c3-a232c552fbd6" />

## 📄 许可证

本项目采用 MIT 许可证，详情请参阅 LICENSE 文件。

---

感谢您使用 AI-TOOL！如有任何问题或建议，欢迎反馈。
