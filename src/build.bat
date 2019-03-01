@echo off
@pushd .
@cd %~dp0

SETLOCAL
set GOPATH=%GOPATH%;%~dp0
echo %GOPATH%

go build casbind
if errorlevel 1 (
    echo build failed
    @popd
    exit /b %errorlevel%
)
echo build success
@popd
