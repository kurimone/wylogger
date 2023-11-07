$repoURL = "https://github.com/mafuka/wylogger"
$arch = if ([System.Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }
$releaseName = "wylogger_windows_" + $arch + ".zip"
$releaseURL = $repoURL + "/releases/latest/download/" + $releaseName

$tempPath = [System.IO.Path]::GetTempPath()
$downloadFile = Join-Path -Path $tempPath -ChildPath "wylogger.zip"
$userHome = [System.Environment]::GetFolderPath('MyDocuments')
$xdnlPath = Join-Path -Path $userHome -ChildPath "wylogger"

Invoke-WebRequest -Uri $releaseURL -OutFile $downloadFile

Expand-Archive -Path $downloadFile -DestinationPath $xdnlPath -Force

$configExamplePath = Join-Path -Path $xdnlPath -ChildPath "config.example.yml"
$configPath = Join-Path -Path $xdnlPath -ChildPath "config.yml"
Copy-Item -Path $configExamplePath -Destination $configPath -Force

Write-Host "Please edit the configuration file at $configPath with your preferred text editor."

Start-Process "notepad.exe" -ArgumentList $configPath