# wylogger

English | [简体中文](https://github.com/mafuka/wylogger/blob/main/README.zh-CN.md)

A tool for automatically authenticate the network of Wen Yuan Talent Apartment.

## Installation

This tool supports Windows, macOS, and Linux platforms.

### Windows

Run Powershell in **administrator mode** and execute the following command:

```powershell
powershell -NoExit -Command "& {Set-ExecutionPolicy -ExecutionPolicy Bypass; Invoke-WebRequest -Uri 'https://gh.api.99988866.xyz/https://raw.githubusercontent.com/mafuka/wylogger/main/script/install.ps1' -OutFile 'install.ps1'; .\install.ps1; Remove-Item .\install.ps1}"
```

*Note: Users within the Chinese Mainland might not be able to access GitHub. In this case, please use the following command:*

```powershell
powershell -NoExit -Command "& {Set-ExecutionPolicy -ExecutionPolicy Bypass; Invoke-WebRequest -Uri 'https://gh.api.99988866.xyz/https://raw.githubusercontent.com/mafuka/wylogger/main/script/install.ps1' -OutFile 'install.ps1'; .\install.ps1 -UseProxy; Remove-Item .\install.ps1}"
```

> **When prompted with "Do you want to change the execution policy?", be sure to answer "Yes to All" (A).** The default execution policy will prevent you from running scripts, and the installation will not be able to proceed.

If all goes well, wylogger will be installed in the `%USERPROFILE%\wylogger` directory.

After the installation is complete, Notepad will open automatically. Please refer to the prompts in it to modify and save the configuration file.

### macOS

First, you need to install **Homebrew**. Open the **Terminal** and execute the following command:

```sh
which brew >/dev/null 2>&1 && echo "Homebrew is installed" || /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Then install wylogger from Homebrew:

```sh
brew tap mafuka/wylogger
brew install wylogger
```

The configuration file will be here: `/etc/wylogger/config.yml`.

### Linux, or another Unix-like OS

Execute the following command in the terminal:

> TODO

## Usage

>TODO

## Autostart

> TODO

## License

[MIT](https://github.com/mafuka/wylogger/blob/main/LICENSE).
