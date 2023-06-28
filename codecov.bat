@echo off
:: 需要切換到go.mod所在的目錄
cd v2
go tool cover -html="../coverage.txt"
