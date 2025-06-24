@echo off
set ConvoName=%1
mkdir convo\%ConvoName% >nul 2>&1
copy config\debug.txt convo\%ConvoName%\ >nul
copy config\files.txt convo\%ConvoName%\ >nul
copy config\log.txt convo\%ConvoName%\ >nul
copy config\sum.txt convo\%ConvoName%\ >nul
echo Saved to convo\%ConvoName%
