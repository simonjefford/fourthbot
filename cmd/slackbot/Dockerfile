FROM ubuntu
ADD slackbot /slackbot
ADD config.json /config.json
ENTRYPOINT ["/slackbot", "-log.format", "logger:stdout?json=true"]
