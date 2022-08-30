docker run -d --rm -p 8080:8080 \
-e GOOGLE_APPLICATION_CREDENTIALS="fbauth.json" \
-e GLOBAL_ENTRY_NOTIFY_API_CONFIG="config.yml" \
-v /var/logs/global-entry-notify-api:/var/logs/global-entry-notify-api \
global-entry-notify-api 