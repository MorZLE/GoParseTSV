# GoParseTSV

## Сервис сканирования директории на наличие .tsv файлов и их перезаписи в формат PDF

### Поддерживается только UTF-8 для tsv файлов
### При запуске укажите свой config файл  -с=config.json 

# Config example
```json
 {
 "directory_in": "D:\\tsv",
 "directory_out": "D:\\pdf\\",
 "dsn": "postgres://postgres:postgres@localhost:5434/postgres",
 "refresh_interval": 1
 }
```


### Request example
```http
POST http://localhost:8080/ HTTP/1.1
Content-Type: application/json
{
    "unitguid": "01749246-95f6-57db-b7c3-2ae0e8be6715",
    "page": 2
    "limit": 5
}
```
### Response example
```json
{
    "ID": "14d013b1-3de3-4dda-8ee6-42474a53e56f",
    "Number": 1,
    "MQTT": "",
    "InventoryID": "G-044322",
    "UnitGUID": "01749246-95f6-57db-b7c3-2ae0e8be6715",
    "MessageID": "cold7_Defrost_status",
    "MessageText": "Разморозка",
    "Context": "",
    "MessageClass": "waiting",
    "Level": 100,
    "Area": "LOCAL",
    "Address": "cold7_status.Defrost_status",
    "Block": false,
    "Type": "",
    "Bit": 0,
    "InvertBit": 0
}
```
## Architecture

![image](https://github.com/MorZLE/GoParseTSV/assets/122459662/f6320f1f-6690-4052-8ce2-13306c972c75)



## Framework

- Web : fiber
- Database : Postgres/gorm
