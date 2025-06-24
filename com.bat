@echo off
go build -o ai.exe ai.go
if %ERRORLEVEL% equ 0 (
    echo Built.
) else (
    echo Failed.
)
pause
