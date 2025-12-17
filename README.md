# Caps Lock Prevent Screen Lock

一个轻量级工具，通过定时模拟按下和释放 Caps Lock 键来防止电脑自动锁屏，适用于需要长时间运行但不希望被系统自动锁定的场景。

## 功能特性

- 🛡️ **防止自动锁屏**：通过模拟键盘操作保持系统活跃状态
- ⏱️ **自定义时间间隔**：可根据需要设置任意秒数的操作间隔（默认 60 秒）
- 🎨 **简洁图形界面**：直观的用户界面，易于操作
- 🖥️ **Windows 支持**：专为 Windows 平台开发
- 📦 **轻量级**：体积小，资源占用低


## 安装说明

### 直接使用

1. 从`build/bin`目录下载最新的`capslock.exe`文件
2. 双击运行即可，无需安装

### 编译安装

如果你想自己编译项目，请按照以下步骤操作：

1. **安装依赖**

   - 安装 Go 语言（1.18+）
   - 安装 Wails 框架：`go install github.com/wailsapp/wails/v2/cmd/wails@latest`
   - 安装 Node.js（用于前端构建）

2. **克隆仓库**

   ```bash
   git clone <repository-url>
   cd capslock
   ```

3. **编译项目**

   ```bash
   wails build
   ```

4. **运行程序**
   在`build/bin`目录找到生成的`capslock.exe`文件，双击运行

## 使用方法

1. **运行程序**：双击`capslock.exe`启动应用
2. **设置间隔**：在输入框中设置防止锁屏的时间间隔（单位：秒）
3. **开始防护**：点击"开始"按钮启动防止锁屏功能
4. **停止防护**：点击"停止"按钮停止防止锁屏功能

## 工作原理

该工具通过`github.com/go-vgo/robotgo`库模拟键盘操作，定时按下并释放 Caps Lock 键，这种操作对系统和其他程序几乎没有影响，但可以有效防止系统进入锁屏状态。

## 技术栈

- **后端**：Go 语言
- **框架**：Wails v2
- **前端**：HTML/CSS/JavaScript
- **UI 组件**：原生 HTML/CSS

## 项目结构

```
capslock/
├── build/              # 编译输出目录
│   ├── bin/           # 可执行文件
│   └── windows/       # Windows特定资源
├── frontend/           # 前端代码
│   ├── src/           # 源代码
│   └── dist/          # 构建输出
├── .gitignore         # Git忽略文件
├── LICENSE            # MIT许可证
├── README.md          # 项目说明文件
├── app.go             # 应用程序逻辑
├── go.mod             # Go模块依赖
├── go.sum             # Go模块校验和
├── main.go            # 主程序入口
└── wails.json         # Wails配置文件
```

## 贡献指南

欢迎贡献代码或提出建议！

1. Fork 本仓库
2. 创建特性分支：`git checkout -b feature/AmazingFeature`
3. 提交更改：`git commit -m 'Add some AmazingFeature'`
4. 推送到分支：`git push origin feature/AmazingFeature`
5. 提交 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看[LICENSE](LICENSE)文件了解详情

## 联系方式

如有问题或建议，欢迎通过 GitHub Issues 提出。
