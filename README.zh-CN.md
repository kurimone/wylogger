# xjtlu-dorm-net-logger

[English](https://github.com/mafuka/xjtlu-dorm-net-logger/blob/main/README.md) | 简体中文

一个自动认证西交利物浦大学宿舍网络的工具。

## 安装

本工具支持 Windows、MacOS 和 Linux 平台。

### Windows

在 `Powershell` 中执行以下命令：

```powershell
Invoke-WebRequest -Uri "https://ghproxy.com/raw.githubusercontent.com/mafuka/xjtlu-dorm-net-logger/main/script/install.ps1" -OutFile "install.ps1"; .\install.ps1; Remove-Item .\install.ps1
```

程序将会被安装在 `%USERPROFILE%\xdnl` 目录下。

### MacOS

首先需要先安装 `Homebrew`，打开 `终端`，执行以下命令：

```sh
which brew >/dev/null 2>&1 && echo "Homebrew is installed" || /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

然后执行以下命令从 Homebrew 安装程序：

```sh
```

### Linux, or another Unix-like OS

在 `终端` 中执行以下命令：

```
```

## 许可证

[MIT](https://github.com/mafuka/xjtlu-dorm-net-logger/blob/main/LICENSE).