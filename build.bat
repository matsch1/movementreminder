@echo off

set scriptPath=%~dp0
set outputfile=%scriptPath%movementreminder.exe
echo binary: %outputfile%

if exist %outputfile% (
    echo delete old binary
    del /f %outputfile%
)

go build -ldflags "-H windowsgui" -o  %outputfile% main.go

if %ERRORLEVEL% NEQ 0 (
    echo Build failed
    exit /b %ERRORLEVEL%
) else (
    echo Build successful
)