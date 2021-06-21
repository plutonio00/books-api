В данном разделе описана только установка проекта.

С условием задачи и комментариями к решению можно ознакомиться здесь: https://github.com/plutonio00/books-api/wiki

## Установка

Для установки необходимо наличие docker, docker-compose и root-прав.

Чтобы запустить проект, необходимо последовательно выполнить в консоли команды (далее все команды даны для UNIX-подобных ОС):

1. `cd /path/to/project`
2. `cp .env.example .env`
3. `cd docker`
4. `cp .env.example .env`
5. `docker-compose build --no-cache`
6. `docker-compose up -d`
7. `docker exec -ti go-books sh`
8. `make`

Следующая команда накатит миграции для БД. 
Для её выполнения следующей нужен параметр MYSQL_URI из файла .env в корне проекта

9. `goose -dir internal/migration mysql "<MYSQL_URI>"  up`

Следующая команда накатит фикстуры (10 авторов + по одной книге для каждого автора). 
Для её выполнения нужно взять параметры для подключения к базе данных из упоминавшегося на предыдущем шаге параметра MYSQL_URI, который имеет вид `<user>:<password>@tcp(<host>:<port>)/<database>?parseTime=true`

10. `charlatan load ./tests/fixtures -u=<user> -d=<database> -p=<password> --host=<host>`

11. `make exec`


После выполнения команд будет запущен сервер, а проект будет доступен по url `http://127.0.0.1:8080/`
При необходимости (например в случае конфликта портов) можно изменить порты в файле docker/.env, а также в файле .env (если порт изменился для mysql)