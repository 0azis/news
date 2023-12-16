# News Service

## Сборка и запуск приложения
```
git clone https://github.com/0azis/news.git
cd news/
docker-compose build
docker-compose up
```
При запуске докер автоматически выполнит все миграции (создаст таблицы и заполнит их тестовыми данными).

## Документация API
1. Получение всех новостей 
```
GET HTTP/1.1 http://localhost:5000/list
```

Заголовки: 
```
Authorization: Bearer <jwt_token>
```

Параметры: 
```
/?limit=5&page=0
```

``limit`` - лимит получаемых данных (default: 5) \
``page`` - номер страницы (default: 0)

Тело ответа:
```json
// HTTP 200 OK
{
	"Success": true,
	"News": [
      {
        "Id": 64,
        "Title": "Lorem ipsum",
        "Content": "Dolor sit amet <b>foo</b>",
        "Categories": [1,2,3]
      },
      {
        "Id": 1,
        "Title": "first",
        "Content": "tratata",
        "Categories": [1]
      }
    ]
}
```

В остальных случаях сервер вернет ошибку (400 Bad Request, 401 Unauthorized, 500 Internal)

2. Обновление данных новости
```
POST HTTP/1.1 http://localhost:5000/edit/:Id
```

Заголовки: 
```
Authorization: Bearer <jwt_token>
```

Параметры:
```
/edit/:Id  *Id - id обновляемой новости
```

Тело запроса:
```json
{
	"Id": 1,
	"Title": "Lorem",
	"Content": "Lorem ipsum doloras",
	"Categories": [1,2,3]
}
```

Тело ответа:
```json
// HTTP 200 OK
{
	"Success": true,
	"News": []
}
```

В остальных случаях сервер вернет ошибку (400 Bad Request, 401 Unauthorized, 500 Internal)

3. Регистрация
```
POST HTTP/1.1 http://localhost:5000/auth/signup
```

Тело запроса:
```json
{
	"login": "John",
	"password": "test123"
}
```

Тело ответа:
```json
// HTTP 200 OK
{
	"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```
В остальных случаях сервер вернет ошибку (400 Bad Request, 409 Conflict, 500 Internal)

3. Вход
```
POST HTTP/1.1 http://localhost:5000/auth/signin
```

Тело запроса:
```json
{
	"login": "John",
	"password": "test123"
}
```

Тело ответа:
```json
// HTTP 200 OK
{
	"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

В остальных случаях сервер вернет ошибку (400 Bad Request, 404 Not Found, 500 Internal)












 
