$projectName = "wylogger"

# Check if installed.
try {
    $commandInfo = Get-Command wylogger -ErrorAction Stop
    Write-Host "$projectName is already in $($commandInfo.Source), the installation will now terminate, press Enter to exit..."
    Read-Host
    exit 0
} catch {
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
        Write-Host "The system architecture could not be identified, the installation will now terminate, press Enter to exit..."
        Read-Host
        exit 1
    }


    Write-Host "Architecture: $architecture"

    # Get package url.
    $repoURL = "https://ghproxy.com/github.com/mafuka/wylogger"
    $os = "windows"
    $packageName = $projectName + "_" + $os + "_" + $arch + ".zip"
    $packageURL = $repoURL + "/releases/latest/download/" + $packageName

    Write-Host "Package URL: $packageURL"

    # Get temp dir.
    $tempDir = [System.IO.Path]::GetTempPath()
    Write-Host "Temp directory: $tempDir"

    # Get download dir.
    $downloadPath = Join-Path -Path $tempDir -ChildPath "wylogger.zip"
    Write-Host "Download package as: $downloadPath"

    # Get install dir.
    $installDir = Join-Path -Path $home -ChildPath "wylogger"
    Write-Host "Install package to: $installDir"

    # Download package.
    Write-Host "Start downloading the package..."
    Invoke-WebRequest -Uri $packageURL -OutFile $downloadPath

    # Unzip package.
    Expand-Archive -Path $downloadFile -DestinationPath $installDir -Force

    # Copy configuration file.
    $confExamplePath = Join-Path -Path $installDir -ChildPath "config.example.yml"
    $confPath = Join-Path -Path $installDir -ChildPath "config.yml"
    Copy-Item -Path $confExamplePath -Destination $confPath -Force

    # Set user env PATH.
    $currentPATH = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::User)
    $newPATH = currentPATH + ";" + $installDir
    [Environment]::SetEnvironmentVariable("PATH", $newPATH, [EnvironmentVariableTarget]::User)

    # Open configuration file.
    Start-Process "notepad.exe" -ArgumentList $configPath

    # End.
    Write-Output "$projectName has been installed to $installDir, the configuration file is open for modification, the installer will now conclude, press Enter to exit..."
    Read-Host
    exit 0
}
