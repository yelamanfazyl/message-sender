# Kolesa Upgrade Bot

<img src="https://i.ibb.co/gJzCG73/upgrade-logo.jpg" alt="Upgrade Logo" width="200"/>

Репозиторий чат-бота для финального проекта Kolesa Upgrade.

Ссылка на чат-бота: [Announcements Bot]

Как все запустить?
1. Создать .toml файл и положить API ключ бота
2. Введите в терминале ``` go get ``` чтобы установить все зависимости
2. Введите в терминале ``` go run main.go --config="путь к .toml файлу" ``` чтобы запустить сервер и бота

Как все работает?

Сервер:
1. По роуту ``` /health ``` получите ответ json ```{"status" : "ok"}```, если сервер правильно запущен
2. По роуту ``` /message ``` при отправке POST запроса с json ```{"message : "Сообщение"}``` сообщение отправиться всем пользователям зарегистрованным у бота

Бот:
1. По команде ```/start``` и ```/hello``` Бот отвечает "Привет, имя" и сохраняет информацию о пользователе в базе
2. По команде ```/info``` пользователь получает информацию о том, сколько сообщений он получил и дату регистрацию у бота
2. По команде ```/delete``` бот удаляет информацию о пользователе и пользователь перестает получать сообщения от бота

Контакты:

Еламан Фазыл - <a href = "https://github.com/yelamanfazyl">GitHub</a> <a href="https://www.linkedin.com/in/yelamanfazyl/">Linkedin</a> <a href="https://t.me/elfazyl">Telegram</a>

Роман Сидаш - <a href = "https://github.com/rsidash">GitHub</a> <a href="https://www.linkedin.com/in/roman-sidash-3b29a91a3/">Linkedin</a> <a href="https://t.me/rsidash">Telegram</a>

[Announcements Bot]: <http://t.me/announcements_kolesa_bot>
