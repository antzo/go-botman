# How to build the bot

1. Run `docker-compose up -d`
2. Visit http://localhost:4040 and get the ngrok links
3. Visit https://api.slack.com/apps/AKTMBJSM8/event-subscriptions? and set the Request URL to the ngrok link
4. ...

Use `docker-compose build && docker-compose up -d` to rebuild the go service and the containers.