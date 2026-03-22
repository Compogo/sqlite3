# Compogo SQLite3

**SQLite3** — это готовый драйвер для интеграции SQLite3 с экосистемой Compogo. Он предоставляет клиент для работы с базой данных, поддерживает миграции, генерацию SQL-запросов и унифицированный интерфейс через `db-client`. Всё настраивается через флаги и подключается одной строкой.

## 🚀 Установка

```bash
go get github.com/Compogo/sqlite3
```

### 📦 Быстрый старт

#### Всё сразу (клиент + миграции + генератор SQL)

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/sqlite3"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithOsSignalCloser(),
        sqlite3.AllComponent,  // всё в одном компоненте
        compogo.WithComponents(
            userRepositoryComponent,
        ),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}
```

#### Только клиент

```go
app := compogo.NewApp("myapp",
    compogo.WithOsSignalCloser(),
    sqlite3.Component,  // только клиент
)
```

### ⚙️ Конфигурация

```shell
./myapp --db.sqlite.dsn="file:/data/app.db?cache=shared&_journal=WAL"
```
