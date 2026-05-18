## 用法

来源：[repo README](https://github.com/nandoolle/pg_slug_gen), [official PGXN release page](https://pgxn.org/dist/pg_slug_gen/), [official release README](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/README.md), [official release SQL](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/sql/pg_slug_gen--1.0.sql), [official release metadata](https://api.pgxn.org/src/pg_slug_gen/pg_slug_gen-1.0.0/META.json)

`pg_slug_gen` 使用加密随机性生成基于时间戳的 slug。官方 1.0.0 release 将它描述为安全、URL-friendly 的短 ID 生成器，请求长度决定时间戳精度。

```sql
CREATE EXTENSION pg_slug_gen;

SELECT gen_random_slug();
SELECT gen_random_slug(13);
```

### 函数

- `gen_random_slug(slug_length int DEFAULT 16) returns text`

release SQL comment 和 README 记录了这些支持值：

- `10`：秒
- `13`：毫秒
- `16`：微秒，也是默认值
- `19`：纳秒

### 精度与格式

每种精度对应一种时间戳宽度和固定 slug 形状：

- `10` digits：`5-5` 格式，总长 11 个字符
- `13` digits：`6-7` 格式，总长 14 个字符
- `16` digits：`8-8` 格式，总长 17 个字符
- `19` digits：`9-10` 格式，总长 20 个字符

README 说明无冲突窗口受时间戳精度限制：每秒、每毫秒、每微秒或每纳秒最多 1 次插入。

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

### 工作方式

官方 README 描述的算法如下：

- 以所选精度获取当前时间戳
- 将每个数字映射到一个基于 QWERTY 的字符桶
- 使用 `pg_strong_random()` 从该桶中随机选择一个字符
- 在中点插入连字符

### 注意事项

- 这是安全短 ID 生成器，不是文本音译或标题转 URL slugifier。
- 相同时间戳下仍可能冲突；上游只声称在插入量不超过所选时间单位每单位一次时保持唯一。
- catalog URL 是当前 `https://github.com/nandoolle/pg_slug_gen` 仓库；PGXN release metadata 仍指向较旧的 `fernandoolle` GitHub URL。
- curated package matrix 面向 PostgreSQL 15 到 18，而 PGXN metadata 只声明最低 PostgreSQL 版本。包可用性请以 catalog matrix 为准。
