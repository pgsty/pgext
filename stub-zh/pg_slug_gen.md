
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION pg_slug_gen;
> SELECT gen_random_slug();
> SELECT gen_random_slug(13);
> ```
>
> 来源：[PGXN 发布 README](https://pgxn.org/dist/pg_slug_gen/1.0.0/)

`pg_slug_gen` 使用密码学随机性生成基于时间戳的 slug。PGXN 发布页将其描述为一个 PostgreSQL 扩展，会把时间戳中的数字映射到字母桶，并在中间插入连字符，从而生成适合 URL 使用的 slug。

## 函数

上游文档给出的 SQL 函数是：

```sql
gen_random_slug(slug_length int DEFAULT 16) RETURNS text
```

README 展示了以下调用方式：

```sql
gen_random_slug()      -- 默认：16（微秒）
gen_random_slug(10)    -- 秒
gen_random_slug(13)    -- 毫秒
gen_random_slug(16)    -- 微秒
gen_random_slug(19)    -- 纳秒
```

## 精度与长度

发布页把精度与时间戳位数、以及可无碰撞吞吐量对应起来：

- `10` 位表示秒级，最多每秒 1 条
- `13` 位表示毫秒级，最多每秒 1,000 条
- `16` 位表示微秒级，最多每秒 1,000,000 条
- `19` 位表示纳秒级，最多每秒 10 亿条

slug 中间会插入一个连字符：

- 秒级：`5-5` 模式，总长度 11
- 毫秒级：`6-7` 模式，总长度 14
- 微秒级：`8-8` 模式，总长度 17
- 纳秒级：`9-10` 模式，总长度 20

## 示例

```sql
SELECT gen_random_slug();
SELECT gen_random_slug(10);
SELECT gen_random_slug(13);
SELECT gen_random_slug(16);
SELECT gen_random_slug(19);
```

作为表默认值：

```sql
CREATE TABLE products (
    id serial PRIMARY KEY,
    name text NOT NULL,
    slug text DEFAULT gen_random_slug() UNIQUE
);
```

## 工作原理

发布页将算法描述为：

1. 取出当前时间戳并按指定精度截断
2. 将每一位数字映射到基于 QWERTY 的字符桶
3. 使用 `pg_strong_random()` 从对应字符桶中随机取一个字符
4. 在中间位置插入连字符

README 还指出，同一时间戳下仍可能出现碰撞，但在微秒级精度下，碰撞概率约为一千万分之一。
