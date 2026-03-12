

## 用法

> [faker: Python Faker 库的包装器](https://github.com/anpandu/postgresql_faker)

`faker` 是一个 PostgreSQL 扩展，封装了 Python 的 Faker 库，提供在 SQL 查询中直接生成逼真假数据的函数。它需要 `plpython3u`。

```sql
CREATE EXTENSION faker;
```

### 生成假数据

```sql
SELECT faker.name();           -- 'John Smith'
SELECT faker.first_name();     -- 'Jane'
SELECT faker.last_name();      -- 'Doe'
SELECT faker.email();          -- 'jane.doe@example.com'
SELECT faker.address();        -- '123 Main St, Anytown, US 12345'
SELECT faker.company();        -- 'Smith LLC'
SELECT faker.phone_number();   -- '(555) 123-4567'
SELECT faker.text();           -- 随机文本段落
SELECT faker.city();           -- 'Portland'
SELECT faker.country();        -- 'United States'
```

注意：`faker.date()` 和 `faker.time()` **不可用**，因为 `date` 和 `time` 是 PostgreSQL 保留关键字。请改用 `faker.date_between()` 或 `faker.date_this_century()`。

### 用假数据填充表

```sql
INSERT INTO users (name, email, address, created_at)
SELECT
  faker.name(),
  faker.email(),
  faker.address(),
  faker.date_this_century()::timestamp
FROM generate_series(1, 1000);
```

### 本地化假数据

语言环境按会话设置，而非按函数调用：

```sql
SELECT faker.faker('de_DE');   -- 为当前会话设置语言环境
SELECT faker.name();           -- 返回德语名字
SELECT faker.address();        -- 返回德语地址
```

### 唯一值

使用 `unique_` 前缀确保会话内值唯一：

```sql
SELECT faker.unique_name();
SELECT faker.unique_email();
```

### 查看所有函数

```sql
SELECT faker._functions();     -- 列出所有 500+ 个可用函数
```

所有 faker 函数返回 `text` 类型。如需其他类型请显式转换。

### 常用 Faker 提供器

| 函数 | 说明 |
|------|------|
| `faker.name()` | 全名 |
| `faker.first_name()` | 名 |
| `faker.last_name()` | 姓 |
| `faker.email()` | 电子邮件地址 |
| `faker.company_email()` | 公司邮箱 |
| `faker.phone_number()` | 电话号码 |
| `faker.address()` | 完整地址 |
| `faker.city()` | 城市名 |
| `faker.country()` | 国家名 |
| `faker.company()` | 公司名 |
| `faker.text()` | 随机文本 |
| `faker.date_this_century()` | 随机日期 |
| `faker.ssn()` | 社会安全号码 |
| `faker.ean()` | EAN 条形码 |
