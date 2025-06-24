@echo off
set ConvoName=%1
copy convo\%ConvoName%\debug.txt config\ >nul
copy convo\%ConvoName%\files.txt config\ >nul
copy convo\%ConvoName%\log.txt config\ >nul
copy convo\%ConvoName%\sum.txt config\ >nul
echo Loaded from convo\%ConvoName%
