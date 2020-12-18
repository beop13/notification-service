# Notification-service

## Config example
 
```yaml
service_settings:
    host: localhost
    port: 8001
email_settings:
    login: 12345@test.test
    password: _________
    host: smtp.test.test
    port: 587
telegram_settings:
    bot_token: 1111111111:AAGzTTlCab3omaUyFMULxVeWYHln4fW5Ufs
    chat_id: -1111111111
```

## Usage

go run /path/to/package/notification-service ./config.yml

## Postman

```json
{
	"info": {
		"_postman_id": "99868202-0221-4861-b75c-1d92aac55747",
		"name": "notification-service",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "notification",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": ""
				},
				"url": {
					"raw": "{{host}}:{{port}}/notification?message_text=Произошло что-то стоящее уведомления LAJ&subject=СРОЧНО URGENT&send_via=telegram&to=mkaren.online@gmail.com",
					"host": [
						"{{host}}"
					],
					"port": "{{port}}",
					"path": [
						"notification"
					],
					"query": [
						{
							"key": "message_text",
							"value": "Произошло что-то стоящее уведомления LAJ"
						},
						{
							"key": "subject",
							"value": "СРОЧНО URGENT"
						},
						{
							"key": "send_via",
							"value": "telegram"
						},
						{
							"key": "to",
							"value": "mkaren.online@gmail.com"
						}
					]
				}
			},
			"response": []
		}
	]
}