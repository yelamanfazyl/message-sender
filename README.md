# Message Sender

Link to Bot: [Announcements Bot]

How to start?
1. Create .toml file and put API key
2. Run ```go get``` to install all dependencies
2. Run ``` go run main.go --config="path to .toml file" ``` to start server and bot

How it works?

Server:
1. Request on route ``` /health ``` receive json ```{"status" : "ok"}```, if everything works
2. Post request on ``` /message ``` with json ```{"message : "Message"}```, message would be sent to all users of bot

Bot:
1. Command ```/start``` and ```/hello```: Bot responds "Привет, name" and saves user in the database
2. Command ```/info```: user receives an information about, how many messages he got and when he subscribed to the bot
2. Command ```/delete```: bot deletes user's information and users stops receiving messages

Contacts:

Yelaman Fazyl - <a href = "https://github.com/yelamanfazyl">GitHub</a> <a href="https://www.linkedin.com/in/yelamanfazyl/">Linkedin</a> <a href="https://t.me/elfazyl">Telegram</a>

[Announcements Bot]: <http://t.me/announcements_kolesa_bot>
