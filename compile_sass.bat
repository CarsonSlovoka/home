@echo off
Set root_dir=%~dp0
Set target_dir=%root_dir%\src\url\static\sass
:: 這是讓go server可以找到該css
Set output_src_dir=%root_dir%\src\url\static\css
Set output_docs_dir=%root_dir%\docs\static\css

cd %target_dir%
:: echo %cd%
sass main.sass:%output_src_dir%\styles.css --no-source-map --style compressed
sass main.sass:%output_docs_dir%\styles.css --no-source-map
start %output_src_dir%
:: start %output_docs_dir%
echo all done & pause > nul
