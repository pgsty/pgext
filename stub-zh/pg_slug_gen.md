## 用法

来源: [official PGXN release page](https://pgxn.org/dist/pg_slug_gen/), [official release README](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/README.md), [official release SQL](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/sql/pg_slug_gen--1.0.sql), [official release metadata](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/META.json)

`pg_slug_gen` 使用密码学随机性生成基于时间戳的 slug。官方 1.0.0 版本将其描述为安全、适合 URL 的短 ID 生成器，且请求的长度决定时间戳精度。

```sql
CREATE EXTENSION pg_slug_gen;

SELECT gen_random_slug();
SELECT gen_random_slug(13);
```

### 函数

- `gen_random_slug(slug_length int DEFAULT 16) returns text`

发布版 SQL 注释和 README 记录了这些受支持的取值：

- `10`: 秒
- `13`: 毫秒
- `16`: 微秒，也是默认值
- `19`: 纳秒

### 精度与格式

每种精度都对应一个时间戳宽度和固定的 slug 形状：

- `10` 位：`5-5` 格式，总长度 11 个字符
- `13` 位：`6-7` 格式，总长度 14 个字符
- `16` 位：`8-8` 格式，总长度 17 个字符
- `19` 位：`9-10` 格式，总长度 20 个字符

README 表示，无碰撞窗口受时间戳精度约束：对应秒、毫秒、微秒、纳秒四种精度时，最多分别每个时间单位插入 1 次。

### 示例

```sql
SELECT gen_random_slug();
SELECT gen_random_slug(10);
SELECT gen_random_slug(16);

CREATE TABLE products (
  id serial PRIMARY KEY,
  name text NOT NULL,
  slug text DEFAULT gen_random_slug() UNIQUE
);
```

### 工作原理

官方 README 将算法描述为：

- 取当前时间戳，并使用所选精度
- 将每一位数字映射到基于 QWERTY 的字符桶
- 使用 `pg_strong_random()` 从对应字符桶中随机选择一个字符
- 在中点插入连字符

### 注意事项

- 这是一个安全短 ID 生成器，不是文本转写器，也不是把标题转换为 URL 的 slugifier。
- 同一时间戳下仍然可能发生碰撞；上游只在插入速率不超过所选时间单位每单位 1 次时声称可保持唯一。
- 官方发布元数据仍指向 `https://github.com/fernandoolle/pg_slug_gen`，但该仓库 URL 当前返回 404。
