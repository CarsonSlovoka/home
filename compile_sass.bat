@echo off
Set root_dir=%~dp0
Set target_dir=%root_dir%\docs\static\sass
Set output_dir=%root_dir%\docs\static\css
cd %target_dir%
:: echo %cd%
sass main.sass:%output_dir%\styles.css --no-source-map --style compressed
:: sass main.sass:%output_dir%\styles.css --no-source-map
:: start %output_dir%
echo all done & pause > nul
