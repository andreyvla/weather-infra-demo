# Weather Infra Demo

Мини-проект: погодный HTTP-сервис + Docker + базовый мониторинг (Prometheus + Blackbox Exporter + Grafana).

## Цели

- Простое веб-приложение: показывает текущую температуру
- Упаковка в Docker
- Мониторинг доступности (HTTP-check)
- Метрики приложения и дашборд

## Запуск

Будет добавлено позже.

## Service design

### HTTP API

#### GET /weather

Возвращает текущую температуру для заданной локации.

Пример ответа:

```
{
  "location": "Colombo",
  "temperature_c": 28.4,
  "source": "open-meteo",
  "updated_at": "2025-12-12T07:15:00Z"
}
```

Назначение:

- основной бизнес-эндпоинт
- обращается к внешнему погодному API
- может быть медленным или временно недоступным

#### GET /health

Health-check эндпоинт.

- всегда возвращает HTTP 200
- не зависит от внешних API или сети
- используется для мониторинга доступности сервиса

Пример ответа:

```
{ "status": "ok" }
```

#### GET /metrics

Эндпоинт для Prometheus.

- используется только системой мониторинга
- отдаёт метрики в формате Prometheus

### Выбор погодного API

В проекте используется Open-Meteo API.

Причины выбора:

- не требует API-ключей
- бесплатный и публичный
- простой JSON-ответ

Пример запроса:
https://api.open-meteo.com/v1/forecast?latitude={LAT}&longitude={LON}&current_weather=true

### Конфигурация через переменные окружения

| Variable          | Description          | Example |
| ----------------- | -------------------- | ------- |
| PORT              | HTTP server port     | 8080    |
| LATITUDE          | Location latitude    | 6.93    |
| LONGITUDE         | Location longitude   | 79.85   |
| WEATHER_CACHE_TTL | Cache TTL in seconds | 60      |
| WEATHER_TIMEOUT   | External API timeout | 2s      |

### Weather caching strategy

Для уменьшения нагрузки на внешний API и повышения стабильности сервиса используется
простой in-memory кэш.

Принцип работы:

- при запросе /weather сервис проверяет, есть ли актуальные данные в кэше
- если данные моложе WEATHER_CACHE_TTL — они возвращаются сразу
- если данные устарели — выполняется запрос к внешнему API и кэш обновляется
