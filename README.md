# Caps Lock Prevent Screen Lock

一个简单的工具，通过定时按下和释放Caps Lock键来防止电脑锁屏。

## 功能特性

- 防止电脑自动锁屏
- 可自定义时间间隔（秒）
- 简单易用的图形界面
- 支持Windows平台

## 安装说明

### 直接使用

1. 从`build/bin`目录下载最新的`toolkit.exe`文件
2. 双击运行即可

### 编译安装

1. 确保已安装Go和Wails
2. 克隆仓库：`git clone <repository-url>`
3. 进入项目目录：`cd capslock`
4. 编译项目：`wails build`
5. 在`build/bin`目录找到生成的可执行文件

## 使用方法

1. 运行程序
2. 设置防止锁屏的时间间隔（默认为60秒）
3. 点击"开始"按钮启动防止锁屏功能
4. 点击"停止"按钮停止防止锁屏功能

## 技术栈

- Go语言
- Wails框架
- HTML/CSS/JavaScript

## 项目结构

```
capslock/
├── build/              # 编译输出目录
├── frontend/           # 前端代码
├── .gitignore         # Git忽略文件
├── README.md          # 项目说明文件
├── app.go             # 应用程序逻辑
├── go.mod             # Go模块依赖
├── go.sum             # Go模块校验和
├── main.go            # 主程序入口
└── wails.json         # Wails配置文件
```

## 许可证

MIT License
