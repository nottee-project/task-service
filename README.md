## POSTGRES

### 1) CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; - Для создания расширения для автоматического создания UUID v4 при создании объекта

### 2) sudo docker exec -it go-task-service psql -h postgres -U feitan -d post

### 3) Создание таблицы в postgres

```
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    title TEXT NOT NULL,
    body TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
```

