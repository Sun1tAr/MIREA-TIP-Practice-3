# Тестовые запросы в формате JSON
```json
{
	"info": {
		"_postman_id": "92118d31-b46d-4dbf-b3ab-1899c31461b1",
		"name": "MIREA-TIP-Practice-3",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "42992055"
	},
	"item": [
		{
			"name": "GET /health",
			"request": {
				"auth": {
					"type": "noauth"
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "http://localhost:8080/health",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"health"
					]
				}
			},
			"response": []
		},
		{
			"name": "POST /tasks + rawData",
			"request": {
				"auth": {
					"type": "noauth"
				},
				"method": "POST",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "default"
					},
					{
						"key": "Access-Control-Allow-Origin",
						"value": "123",
						"type": "default"
					}
				],
				"body": {
					"mode": "raw",
					"raw": "{\"title\":\"Молоко купить\"}"
				},
				"url": {
					"raw": "http://localhost:8080/tasks",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"tasks"
					]
				}
			},
			"response": []
		},
		{
			"name": "GET /tasks",
			"request": {
				"auth": {
					"type": "noauth"
				},
				"method": "GET",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "default"
					},
					{
						"key": "Access-Control-Allow-Origin",
						"value": "123",
						"type": "default"
					}
				],
				"url": {
					"raw": "http://localhost:8080/tasks",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"tasks"
					]
				}
			},
			"response": []
		},
		{
			"name": "GET /tasks/{id}",
			"request": {
				"auth": {
					"type": "noauth"
				},
				"method": "GET",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "default"
					},
					{
						"key": "Access-Control-Allow-Origin",
						"value": "123",
						"type": "default"
					}
				],
				"url": {
					"raw": "http://localhost:8080/tasks/1",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"tasks",
						"1"
					]
				}
			},
			"response": []
		},
		{
			"name": "PATCH /tasks/{id}",
			"request": {
				"auth": {
					"type": "noauth"
				},
				"method": "PATCH",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "default"
					},
					{
						"key": "Access-Control-Allow-Origin",
						"value": "123",
						"type": "default"
					}
				],
				"url": {
					"raw": "http://localhost:8080/tasks/1",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"tasks",
						"1"
					]
				}
			},
			"response": []
		},
		{
			"name": "DELETE /tasks/{id}",
			"request": {
				"auth": {
					"type": "noauth"
				},
				"method": "DELETE",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "default"
					},
					{
						"key": "Access-Control-Allow-Origin",
						"value": "123",
						"type": "default"
					}
				],
				"url": {
					"raw": "http://localhost:8080/tasks/1",
					"protocol": "http",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"tasks",
						"1"
					]
				}
			},
			"response": []
		}
	]
}

```