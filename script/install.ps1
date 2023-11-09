param (
    [bool]$UseProxy = $false,
    [string]$Action = "Install",
    [string]$PackageVersion = "latest"
)

$projectName = "wylogger"

function Install-Package {
    # Check if installed.
    try {
        $commandInfo = Get-Command $projectName -ErrorAction Stop
        $isInstalled = $true
    } catch {
        $isInstalled = $false
    }

    # Exit if installed.
    if ($isInstalled -eq $true) {
        Write-Host "$projectName is already in $($commandInfo.Source), no changes have been made."
        Exit-Script
    }

    Write-Host "$projectName is not installed, proceeding with installation."

    # Initialize Environment.
    Write-Host "============ INITIALISING ENV ============"

    # Check system architecture.
    $arch = $env:PROCESSOR_ARCHITECTURE
    $archWow = $env:PROCESSOR_ARCHITEW6432

    if ($arch -eq "AMD64" -or $archWow -eq "AMD64") {
        $architecture = "amd64"
    } elseif ($arch -eq "x86" -and -not $archWow) {
        $architecture = "386"
    } elseif ($arch -eq "ARM64") {
        $architecture = "arm64"
    } else {
        $architecture = "unknown"
        Write-Host "The system architecture could not be identified."
        Exit-Script
    }

    Write-Host "- Architecture: $architecture"

    # Get package url.
    Write-Host "- Use Proxy: $UseProxy"
    Write-Host "- Package Version: $PackageVersion"

    $repoURL = if ($UseProxy) {
        # "https://github.hscsec.cn/mafuka/$projectName"
        "https://mirror.ghproxy.com/https://github.com/mafuka/$projectName"
    } else {
        "https://github.com/mafuka/$projectName"
    }

    $os = "windows"
    $packageName = $projectName + "_" + $os + "_" + $arch + ".zip"
    $packageURL = $repoURL + "/releases/$PackageVersion/download/" + $packageName

    Write-Host "- Package URL: $packageURL"

    # Get temp dir.
    $tempDir = [System.IO.Path]::GetTempPath()
    Write-Host "- Temp directory: $tempDir"

    # Get download dir.
    $downloadPath = Join-Path -Path $tempDir -ChildPath "$projectName.zip"
    Write-Host "- Download package as: $downloadPath"

    # Get install dir.
    $installDir = Join-Path -Path $home -ChildPath "$projectName"
    Write-Host "- Install package to: $installDir"

    Write-Host "============ ENV INITIALISED ============="


    # Download package.
    Write-Host "Start downloading the package..."
    try {
        Invoke-WebRequest -Uri $packageURL -OutFile $downloadPath
    } catch {
        Write-Host "Package download failed."
        Write-Host $_
        Exit-Script
    }

    # Unzip package.
    try {
        Expand-Archive -Path $downloadPath -DestinationPath $installDir -Force
    } catch {
        Write-Host "Failed to unpack the package."
        Write-Host $_
        Exit-Script
    }

    # Copy configuration file.
    $confExamplePath = Join-Path -Path $installDir -ChildPath "config.example.yml"
    $confPath = Join-Path -Path $installDir -ChildPath "config.yml"
    Copy-Item -Path $confExamplePath -Destination $confPath -Force

    # Set user env PATH.
    $currentPATH = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::User)
    $newPATH = $currentPATH + ";" + $installDir
    [Environment]::SetEnvironmentVariable("PATH", $newPATH, [EnvironmentVariableTarget]::User)

    # Check if install successful.
    Write-Host "Run $projectName -v..."
    try {
        $commandInfo = Get-Command $projectName -ErrorAction Stop
    }
    catch {
        Write-Host "$projectName not found, installation failed."
        Exit-Script
    }
    Invoke-Expression "& $($commandInfo.Source) `-v` "

    # Installation complete.
    Write-Host "$projectName has been installed to $installDir." 

    # Open configuration file.
    Write-Host "The configuration file is open for modification."
    Start-Process "notepad.exe" -ArgumentList $confPath

    Exit-Script
}

function Remove-Package {
    try {
        $commandInfo = Get-Command $projectName -ErrorAction Stop
    }
    catch {
        Write-Host "$projectName not found, no changes have been made."
        Exit-Script
    }
    $installDir = Split-Path $($commandInfo.Source)

    Write-Host "Removing $projectName ($($commandInfo.Source))..."

    $currentPATH = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::User)
    $currentPATHArr = $currentPath -split [System.IO.Path]::PathSeparator
    $newPATHArr = $currentPATHArr.Where({ $_ -ne $installDir })
    $newPATH = $newPATHArr -join [System.IO.Path]::PathSeparator
    [Environment]::SetEnvironmentVariable("PATH", $newPATH, [EnvironmentVariableTarget]::User)

    try {
        Remove-Item -Path $installDir -Recurse -Force
    } catch {
        Write-Host "An error occurred while removing the package directory $installDir."
        Exit-Script
    }
    Write-Host "Removed $projectName ($($commandInfo.Source))."
    Exit-Script
}

# Main
function Start-Script {
    Write-Host "$projectName Installation Script"
    Write-Host "=========================================="
    Write-Host "- Author: Mafuka<i@mafuka.com>"
    Write-Host "- Version: 202311091411"
    Write-Host "- Action: $Action"
    Write-Host "=========================================="

    if ($Action -eq "Install") {
        Install-Package
    } elseif ($Action -eq "Remove") {
        Remove-Package
    } else {
        Write-Host "Unknown action $Action, available Action values are: Install, Remove."
        Exit-Script
    }
}

function Exit-Script {
    Read-Host -Prompt "Press Enter to exit..."
    Exit
}

Start-Script
