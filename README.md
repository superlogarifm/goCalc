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

При отправке запроса с неподдерживающимся¹ знаком (например, %) или другими символами:
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type:application/json" -d "{\"expression\":\"1%1\"}"
```
Сервер вернет ошибку:
```
{"error": "incorrect input"}
```

При делении на ноль вернется
```
{"error": "division by zero"}
```
При отправке запроса с ошибками в синтаксисе
```
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type:application/json" -d "{\"expression\":\"1*(1++\"}"
```
Вернется ошибка:
```
{"error": "expression is not valid"}
```

Если HTTP-метод не POST, то вернется ошибка со статусом 405:
```{"error": "method not allowed", "status": 405}```


1 - Поддерживаются выражения со знаками +, -, * (умножить), / (деление)


