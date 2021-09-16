В данном разделе описана только установка проекта.

С условием задачи и комментариями к решению можно ознакомиться здесь: https://github.com/plutonio00/books-api/wiki

## Установка

Для установки необходимо наличие docker, docker-compose и root-прав.

Для развёртывания проекта в UNIX-подобных ОС в проекте предусмотрен скрипт на bash.

Чтобы запустить проект, необходимо выполнить в консоли следующие действия:

1. `cd /path/to/project`
2. `sudo chmod +x init_project.sh`
3. `./init_project.sh`

В результате:
1. Будут созданы файлы env для docker и проекта
2. Будут собраны и запущены docker-контейнеры
3. Будет произведена сборка go кода
4. Накатятся миграции и фикстуры
5. Будет запущен сервер

Кроме того, при сборке docker-контейнера для mysql будут автоматически созданы база данных для проекта, пользователь с правами на нее, изменен root пароль для mysql.

После выполнения после проект будет доступен по url `http://127.0.0.1:8080/`

Swagger документация будет доступна по адресу `http://127.0.0.1:8080/swagger/index.html`