Простой сервер на Go для вычисления математических выражений.

1. *Запуск*

Чтобы запустить сервер, введите в командой строке: 
```
go run cmd/calc_service/start.go
```

2. *Отправка запросов*

Запросы на сервер отправляются с помощью curl в командой строке.

Например
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type:application/json" -d "{\"expression\":\"2+2*2\"}"
```
Вернется ответ:
```
{"result": 6.0000}
```
3. *Возможные ошибки*

При отправке запроса с неподдерживающимся знаком (например, %):
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type:application/json" -d "{\"expression\":\"1%1\"}"
```
Сервер вернет ошибку:
```
{"error": "incorrect input"}
```

