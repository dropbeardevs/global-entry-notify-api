#/bin/bash

if [ $# -eq 0 ]
then
    echo "Usage: --daemon, --shell, --interactive"

elif [ $1 = "--daemon" ]
then
    docker run -d --rm -p 8080:8080 \
    -e GOOGLE_APPLICATION_CREDENTIALS="fbauth.json" \
    -e GLOBAL_ENTRY_NOTIFY_API_CONFIG="config.yml" \
    -v /var/logs/global-entry-notify-api:/var/logs/global-entry-notify-api \
    global-entry-notify-api 

elif [ $1 = "--shell" ]
then
    docker run --rm -it -p 8080:8080 \
    -e GOOGLE_APPLICATION_CREDENTIALS="fbauth.json" \
    -e GLOBAL_ENTRY_NOTIFY_API_CONFIG="config.yml" \
    -v /var/logs/global-entry-notify-api:/var/logs/global-entry-notify-api \
    global-entry-notify-api sh

elif [ $1 = "--interactive" ]
then
    docker run --rm -p 8080:8080 \
    -e GOOGLE_APPLICATION_CREDENTIALS="fbauth.json" \
    -e GLOBAL_ENTRY_NOTIFY_API_CONFIG="config.yml" \
    -v /var/logs/global-entry-notify-api:/var/logs/global-entry-notify-api \
    global-entry-notify-api
fi