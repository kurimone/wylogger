# wylogger

[English](https://github.com/mafuka/wylogger/blob/main/README.md) | 简体中文

一个自动认证文缘人才公寓网络的工具。

## 安装

本工具支持 Windows、macOS 和 Linux 平台。

### Windows

以**管理员模式**运行 Powershell ，执行以下命令：

```powershell
powershell -NoExit -Command "& {Set-ExecutionPolicy -ExecutionPolicy Bypass; Invoke-WebRequest -Uri 'https://raw.githubusercontent.com/mafuka/wylogger/main/script/install.ps1' -OutFile 'install.ps1'; .\install.ps1; Remove-Item .\install.ps1}"
```

*注：国内网络用户可能无法访问 GitHub，此时请使用以下命令：*

```powershell
powershell -NoExit -Command "& {Set-ExecutionPolicy -ExecutionPolicy Bypass; Invoke-WebRequest -Uri 'https://gh.api.99988866.xyz/https://raw.githubusercontent.com/mafuka/wylogger/main/script/install.ps1' -OutFile 'install.ps1'; .\install.ps1 -UseProxy 1; Remove-Item .\install.ps1}"
```

> **当遇到 “是否要更改执行策略?” 的询问时，请务必回答 “全是”（A）。** 默认的执行策略会阻止您运行脚本，安装将无法进行。

如无意外，wylogger 将会被安装在 `%USERPROFILE%\wylogger` 目录下。

安装完成后，记事本会自动打开，请参考其中的提示修改并保存配置文件。

### macOS

首先需要安装 **Homebrew**，打开**终端**，执行以下命令：

```sh
which brew >/dev/null 2>&1 && echo "Homebrew is installed" || /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

然后从 Homebrew 安装 wylogger：

```sh
brew tap mafuka/tap
brew install wylogger
```

配置文件将会在这里： `/etc/wylogger/config.yml`。

> TODO

### Linux, or another Unix-like OS

在终端中执行以下命令：

> TODO


## 使用

> TODO

## 自启动

> TODO

## 许可证

[MIT](https://github.com/mafuka/wylogger/blob/main/LICENSE).
