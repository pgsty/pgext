## 用法

来源：

- [官方示例](https://github.com/lameferret/pg_nanoid/blob/d216e0b481d8ffb8a2c93b1dc3420bdafe31b486/Examples.md)
- [官方控制文件](https://github.com/lameferret/pg_nanoid/blob/d216e0b481d8ffb8a2c93b1dc3420bdafe31b486/pg_nanoid.control)
- [官方构建清单](https://github.com/lameferret/pg_nanoid/blob/d216e0b481d8ffb8a2c93b1dc3420bdafe31b486/Cargo.toml)

`pg_nanoid` 增加 `nanoid` 标识符类型，以及生成单个、指定长度与批量 Nano ID 的函数。它适合紧凑的应用标识符，但唯一性仍是概率性的，必须继续用主键或唯一约束保证。

### 核心流程

```sql
CREATE EXTENSION pg_nanoid;

CREATE TABLE users (
    id nanoid PRIMARY KEY,
    name text NOT NULL
);

INSERT INTO users (id, name)
SELECT id, 'User #' || row_number() OVER ()
FROM nanoid_batch(5);

CREATE TABLE sessions (
    user_id nanoid REFERENCES users(id),
    token text UNIQUE NOT NULL
);

INSERT INTO sessions (user_id, token)
SELECT id, nanoid_len(64)
FROM users
LIMIT 1;
```

批量写入时使用 `nanoid_batch` 分摊生成开销；应用要求明确长度时使用 `nanoid_len`。

### 重要对象

- `nanoid` 是扩展的标识符类型。
- `nanoid_batch(count)` 返回包含生成 ID 的多行结果。
- `nanoid_len(length)` 返回指定长度的生成文本 ID。

使用固定示例没有展示的 API 前，应检查已安装函数签名；本次核验的软件包声明为版本 `0.0.0`，并非稳定的语义版本。

### 限制

本次核验的清单只启用了 PostgreSQL 18 pgrx 特性，并要求超级用户安装；其他服务器大版本需要独立构建并测试制品。控制文件将扩展标为不可信且不可迁移。

应根据预期数量与风险容忍度选择标识长度和字符表，保留数据库唯一约束，并在罕见冲突时重试。不要自动把生成 ID 当作认证秘密：安全令牌需要另行验证熵源、字符表、暴露期限、比较方式、日志记录与撤销要求。
