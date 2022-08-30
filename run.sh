docker run --rm -p 8080:8080 \
-e GOOGLE_APPLICATION_CREDENTIALS="global-entry-c8373-fe72e1ae9c11.json" \
-e GLOBAL_ENTRY_NOTIFY_API_CONFIG="config.yml" \
-v /var/log:/usr/local/bin/logs
global-entry-notify-api 