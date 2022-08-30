docker run -d --rm -p 8080:8080 \
-e GOOGLE_APPLICATION_CREDENTIALS="global-entry-c8373-fe72e1ae9c11.json" \
-e GLOBAL_ENTRY_NOTIFY_API_CONFIG="config.yml" \
-v /var/logs/global-entry-notify-api:/var/logs/global-entry-notify-api \
global-entry-notify-api 