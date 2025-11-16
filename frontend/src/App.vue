<template>
    <div id="app" class="app-container">
        <div class="home-page">
            <header class="app-header">
                <h1>网站管理工具</h1>
                <button @click="showAddForm = true" class="add-btn">+ 添加网站</button>
            </header>
            
            <main class="website-list">
                <!-- 加载状态 -->
                <div v-if="isLoading" class="loading-state">
                    <div class="loading-spinner"></div>
                    <p>正在加载网站列表...</p>
                </div>
                
                <!-- 空状态 -->
                <div v-else-if="websites.length === 0" class="empty-state">
                    <div class="empty-icon">📂</div>
                    <p>暂无网站，点击上方按钮添加</p>
                </div>
                
                <!-- 网站列表 -->
                <div v-else class="website-grid">
                    <div v-for="website in websites" :key="website.id" class="website-card">
                        <div class="website-info">
                            <h3 class="website-name">{{ website.name }}</h3>
                            <p class="website-url">{{ website.url }}</p>
                            <span class="website-date">{{ formatDate(website.createdAt) }}</span>
                        </div>
                        <div class="website-actions">
                            <button @click="openWebsite(website.url, website.name)" class="visit-btn">访问</button>
                            <button @click="deleteWebsite(website.id)" class="delete-btn">删除</button>
                        </div>
                    </div>
                </div>
            </main>
            
            <!-- 添加网站表单 -->
            <div v-if="showAddForm" class="modal-overlay">
                <div class="add-form">
                    <h3>添加新网站</h3>
                    <form @submit.prevent="handleAddWebsite">
                        <div class="form-group">
                            <label for="name">网站名称:</label>
                            <input 
                                id="name" 
                                v-model="newWebsite.name" 
                                type="text" 
                                required 
                                placeholder="输入网站名称"
                            />
                        </div>
                        <div class="form-group">
                            <label for="url">网站地址:</label>
                            <input 
                                id="url" 
                                v-model="newWebsite.url" 
                                type="url" 
                                required 
                                placeholder="https://example.com"
                            />
                        </div>
                        <div class="form-actions">
                            <button type="button" @click="showAddForm = false" class="cancel-btn">取消</button>
                            <button type="submit" class="submit-btn">添加</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'WebsiteManager',
    data() {
        return {
            websites: [],
            showAddForm: false,
            isLoading: false,
            newWebsite: {
                name: '',
                url: ''
            }
        }
    },
    async mounted() {
        await this.loadWebsites();
        
        // 添加快捷键监听
        this.setupShortcuts();
    },
    beforeDestroy() {
        // 清理快捷键监听
        if (this.keydownHandler) {
            document.removeEventListener('keydown', this.keydownHandler);
        }
    },
    methods: {
        // 加载网站列表
        async loadWebsites() {
            try {
                this.isLoading = true;
                if (window.go && window.go.main && window.go.main.App) {
                    this.websites = await window.go.main.App.GetAllWebsites();
                }
            } catch (error) {
                console.error('加载网站列表失败:', error);
            } finally {
                this.isLoading = false;
            }
        },
        
        // 添加网站
        async handleAddWebsite() {
            try {
                if (window.go && window.go.main && window.go.main.App) {
                    await window.go.main.App.AddWebsite(this.newWebsite.name, this.newWebsite.url);
                    this.showAddForm = false;
                    this.newWebsite = { name: '', url: '' };
                    await this.loadWebsites();
                }
            } catch (error) {
                console.error('添加网站失败:', error);
                alert('添加网站失败，请检查输入格式');
            }
        },
        
        // 删除网站
        async deleteWebsite(id) {
            if (confirm('确定要删除这个网站吗？')) {
                try {
                    if (window.go && window.go.main && window.go.main.App) {
                        await window.go.main.App.DeleteWebsite(id);
                        await this.loadWebsites();
                    }
                } catch (error) {
                    console.error('删除网站失败:', error);
                }
            }
        },
        
        // 打开网站（使用Wails的内嵌浏览功能）
        async openWebsite(url, name) {
            try {
                // 使用Wails的WebView功能在当前窗口中打开网站
                // 这需要在Wails配置中允许跨域访问
                if (window.go && window.go.main && window.go.main.App) {
                    await window.go.main.App.NavigateToWebsite(url);
                } else {
                    // 如果后端方法不可用，使用iframe方式
                    this.openInIframe(url, name);
                }
            } catch (error) {
                console.error('打开网站失败:', error);
                alert('打开网站失败');
            }
        },
        
        // 使用iframe方式打开网站
        openInIframe(url, name) {
            // 创建iframe容器
            const iframeContainer = document.createElement('div');
            iframeContainer.className = 'website-iframe-container';
            iframeContainer.innerHTML = `
                <div class="website-header">
                    <button onclick="document.querySelector('.website-iframe-container').remove()" class="back-btn">← 返回</button>
                    <h2 class="website-title">${name}</h2>
                    <span class="website-url">${url}</span>
                </div>
                <iframe src="${url}" class="website-iframe" frameborder="0"></iframe>
            `;
            
            // 添加样式
            const style = document.createElement('style');
            style.textContent = `
                .website-iframe-container {
                    position: fixed;
                    top: 0;
                    left: 0;
                    width: 100%;
                    height: 100%;
                    background: white;
                    z-index: 1000;
                }
                .website-header {
                    display: flex;
                    align-items: center;
                    gap: 1rem;
                    padding: 1rem 2rem;
                    background: white;
                    border-bottom: 1px solid #e0e0e0;
                    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
                }
                .back-btn {
                    background: #6c757d;
                    color: white;
                    border: none;
                    padding: 0.5rem 1rem;
                    border-radius: 5px;
                    cursor: pointer;
                }
                .website-title {
                    margin: 0;
                    font-size: 1.4rem;
                    color: #333;
                }
                .website-url {
                    color: #666;
                    font-size: 0.9rem;
                }
                .website-iframe {
                    width: 100%;
                    height: calc(100% - 80px);
                    border: none;
                }
            `;
            
            document.head.appendChild(style);
            document.body.appendChild(iframeContainer);
        },
        
        // 设置快捷键
        setupShortcuts() {
            this.keydownHandler = (event) => {
                // Ctrl+D 快捷键添加网站
                if (event.ctrlKey && event.key === 'd') {
                    event.preventDefault();
                    this.showAddForm = true;
                }
            };
            
            document.addEventListener('keydown', this.keydownHandler);
        },
        
        // 格式化日期
        formatDate(dateString) {
            const date = new Date(dateString);
            return date.toLocaleDateString('zh-CN');
        }
    }
};
</script>

<style scoped>
.app-container {
    width: 100%;
    height: 100vh;
    margin: 0;
    padding: 0;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

/* 主页样式 */
.home-page {
    height: 100%;
    display: flex;
    flex-direction: column;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
}

.app-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 2rem;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
}

.app-header h1 {
    margin: 0;
    font-size: 1.8rem;
}

.add-btn {
    background: #4CAF50;
    color: white;
    border: none;
    padding: 0.75rem 1.5rem;
    border-radius: 5px;
    cursor: pointer;
    font-size: 1rem;
    transition: background 0.3s;
}

.add-btn:hover {
    background: #45a049;
}

.website-list {
    flex: 1;
    padding: 2rem;
    overflow-y: auto;
}

/* 加载状态样式 */
.loading-state {
    text-align: center;
    margin-top: 4rem;
    font-size: 1.2rem;
}

.loading-spinner {
    width: 40px;
    height: 40px;
    border: 4px solid rgba(255, 255, 255, 0.3);
    border-left: 4px solid white;
    border-radius: 50%;
    animation: spin 1s linear infinite;
    margin: 0 auto 1rem;
}

@keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
}

/* 空状态样式 */
/* 加载状态样式 */
.loading-state {
    text-align: center;
    margin-top: 4rem;
    font-size: 1.2rem;
}

.loading-spinner {
    width: 40px;
    height: 40px;
    border: 4px solid rgba(255, 255, 255, 0.3);
    border-left: 4px solid white;
    border-radius: 50%;
    animation: spin 1s linear infinite;
    margin: 0 auto 1rem;
}

@keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
}

/* 空状态样式 */
.empty-state {
    text-align: center;
    margin-top: 4rem;
    font-size: 1.2rem;
}

.empty-icon {
    font-size: 4rem;
    margin-bottom: 1rem;
    opacity: 0.7;
}

.empty-icon {
    font-size: 4rem;
    margin-bottom: 1rem;
    opacity: 0.7;
}

.website-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 1.5rem;
    max-width: 1200px;
    margin: 0 auto;
}

.website-card {
    background: rgba(255, 255, 255, 0.1);
    border-radius: 10px;
    padding: 1.5rem;
    backdrop-filter: blur(10px);
    border: 1px solid rgba(255, 255, 255, 0.2);
    transition: transform 0.3s;
}

.website-card:hover {
    transform: translateY(-5px);
}

.website-info h3 {
    margin: 0 0 0.5rem 0;
    font-size: 1.2rem;
}

.website-url {
    margin: 0 0 0.5rem 0;
    color: rgba(255, 255, 255, 0.8);
    font-size: 0.9rem;
    word-break: break-all;
}

.website-date {
    color: rgba(255, 255, 255, 0.6);
    font-size: 0.8rem;
}

.website-actions {
    display: flex;
    gap: 0.5rem;
    margin-top: 1rem;
}

.visit-btn, .delete-btn {
    flex: 1;
    padding: 0.5rem;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 0.9rem;
    transition: background 0.3s;
}

.visit-btn {
    background: #2196F3;
    color: white;
}

.visit-btn:hover {
    background: #1976D2;
}

.delete-btn {
    background: #f44336;
    color: white;
}

.delete-btn:hover {
    background: #d32f2f;
}

/* 添加快捷键提示样式 */
.shortcut-hint {
    position: fixed;
    bottom: 1rem;
    right: 1rem;
    background: rgba(255, 255, 255, 0.2);
    padding: 0.5rem 1rem;
    border-radius: 5px;
    backdrop-filter: blur(10px);
    font-size: 0.9rem;
}

.shortcut-hint kbd {
    background: rgba(255, 255, 255, 0.3);
    padding: 0.2rem 0.5rem;
    border-radius: 3px;
    border: 1px solid rgba(255, 255, 255, 0.5);
    font-family: monospace;
}

/* 添加表单样式 */
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.add-form {
    background: white;
    color: #333;
    padding: 2rem;
    border-radius: 10px;
    min-width: 400px;
    max-width: 500px;
}

.add-form h3 {
    margin: 0 0 1.5rem 0;
    text-align: center;
}

.form-group {
    margin-bottom: 1rem;
}

.form-group label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: 500;
}

.form-group input {
    width: 100%;
    padding: 0.75rem;
    border: 1px solid #ddd;
    border-radius: 5px;
    font-size: 1rem;
}

.form-actions {
    display: flex;
    gap: 1rem;
    margin-top: 1.5rem;
}

.cancel-btn, .submit-btn {
    flex: 1;
    padding: 0.75rem;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 1rem;
}

.cancel-btn {
    background: #f5f5f5;
    color: #333;
}

.cancel-btn:hover {
    background: #e0e0e0;
}

.submit-btn {
    background: #4CAF50;
    color: white;
}

    .submit-btn:hover {
        background: #45a049;
    }
</style>

<style>
html, body {
    margin: 0;
    padding: 0;
    width: 100%;
    height: 100%;
    overflow: hidden;
}
</style>