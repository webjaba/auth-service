# Template сервиса авторизации

## Файловая структура

```
api/
├── api.proto               # grpc контракт сервиса
cmd/
├── configurations/         # статическая конфигурация сервиса
├── main.go
e2e/                        # make команды которые дергают grpc ручки
internal/
├── app/                    # Основной пакет с хэндлерами
├── pkg/
│   ├── configurations/     # Шаблоны конфигураций
│   ├── domain/             # Модели данных
│   ├── mapper/             # Маппер-функции для моделей данных
│   ├── pb/                 # Сгенерированных из .proto код
│   └── store/              # Слой бд
├── migrations/
├── README.md
├── .gitignore
├── makefile
├── go.sum
└── go.mod
```
