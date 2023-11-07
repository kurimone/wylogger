# wylogger

[English](https://github.com/mafuka/wylogger/blob/main/README.md) | 简体中文

A tool for automatic authentication of the Wen Yuan Talent Apartment network.

## Installation

This tool supports Windows, MacOS, and Linux platforms.

### Windows

Execute the following command in `Powershell`:

```powershell
Invoke-WebRequest -Uri "https://ghproxy.com/raw.githubusercontent.com/mafuka/wylogger/main/script/install.ps1" -OutFile "install.ps1"; .\install.ps1; Remove-Item .\install.ps1
```

The program will be installed in the `%USERPROFILE%\xdnl` directory.

After installation, Notepad will open automatically, please refer to the prompts inside to modify and save the configuration file.

### MacOS

First, you need to install `Homebrew`. Open `Terminal` and execute the following command:

```sh
which brew >/dev/null 2>&1 && echo "Homebrew is installed" || /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Then execute the following command to install the program from Homebrew:

> TODO

### Linux, or another Unix-like OS

Execute the following command in `Terminal`:

> TODO

## Usage

> TODO

## Autostart

> TODO

## License

[MIT](https://github.com/mafuka/wylogger/blob/main/LICENSE).