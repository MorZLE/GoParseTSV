# ParseTSVBiocad

## Сервис сканирования директории на наличие .tsv файлов

### Поддерживается только UTF-8 для tsv файлов
### При запуске укажите свой config файл  -с=config.json 

# Пример
config.json
#### {
#### "directory_in": "D:\\tsv",
#### "directory_out": "D:\\pdf\\",
#### "dsn": "postgres://postgres:postgres@localhost:5434/postgres",
#### "refresh_interval": 1
#### }

## Architecture

![image](https://github.com/MorZLE/GoParseTSV/assets/122459662/f6320f1f-6690-4052-8ce2-13306c972c75)



## Framework

- Web : fiber
- Database : Postgres/gorm
