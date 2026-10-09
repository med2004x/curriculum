$ErrorActionPreference = 'Stop'

$toolchainBin = 'C:\msys64\ucrt64\bin'
$makeExecutable = Join-Path $toolchainBin 'mingw32-make.exe'
$gccExecutable = Join-Path $toolchainBin 'gcc.exe'
$shimDirectory = Join-Path $env:LOCALAPPDATA 'Programs\DART\bin'
$makeShim = Join-Path $shimDirectory 'make.cmd'

if (-not (Test-Path -LiteralPath $makeExecutable)) {
    throw "GNU Make was not found at $makeExecutable. Install the MSYS2 UCRT64 make package first."
}
if (-not (Test-Path -LiteralPath $gccExecutable)) {
    throw "GCC was not found at $gccExecutable. Install the MSYS2 UCRT64 GCC package first."
}

New-Item -ItemType Directory -Path $shimDirectory -Force | Out-Null
@'
@echo off
"C:\msys64\ucrt64\bin\mingw32-make.exe" %*
exit /b %ERRORLEVEL%
'@ | Set-Content -LiteralPath $makeShim -Encoding Ascii

$userPath = [Environment]::GetEnvironmentVariable('Path', 'User')
$pathEntries = [System.Collections.Generic.List[string]]::new()
if ($userPath) {
    foreach ($entry in $userPath -split ';') {
        if ($entry -and -not $pathEntries.Contains($entry)) {
            $pathEntries.Add($entry)
        }
    }
}
foreach ($entry in @($toolchainBin, $shimDirectory)) {
    if (-not ($pathEntries | Where-Object { $_.TrimEnd('\') -ieq $entry.TrimEnd('\') })) {
        $pathEntries.Add($entry)
    }
}
[Environment]::SetEnvironmentVariable('Path', ($pathEntries -join ';'), 'User')

Write-Host "Added GCC and make to your user PATH. Open a new terminal, then run 'gcc --version' and 'make --version'."
