# Docker Compose

**Docker Compose** – это инструмент для определения и запуска 
мультиконтейнерных приложений.

**Docker Compose** использует конфигурацию на основе **YAML**-файла.

## Источники

- [Docker Compose Tutorial by *Programming with Mosh*](https://www.youtube.com/watch?v=HG6yIjZapSA)

## **YAML** VS **JSON**

Пример **JSON**-файла:

```json
{
  "name": "Alexander Konovalov",
  "age": 18,
  "is_developer": true,
  "programming_langs": ["CSharp", "JS"],
  "skills": {
    "CSharp": 60,
    "JS": 30
  }
}
```

Тот же пример но на **YAML**:

```yaml
---
name: Alexander Konovalov
age: 18
is_developer: true
programming_langs:
  - CSharp
  - JS
skills:
  CSharp: 60
  JS: 30
```

> **YAML** сложнее парсить, чем **JSON**. **YAML** проще для людей, **JSON** 
> проще для машин.

## `docker-compose.yml` (`docker-compose.yaml`)

В начале файла указывается 
[версия спецификации](https://docs.docker.com/compose/compose-file/compose-versioning/).


