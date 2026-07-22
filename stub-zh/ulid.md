## 用法

来源：

- [官方扩展 README](https://github.com/Creatif/creatif-backend/blob/7b0466a2b0f505c824d98c2276107531ff1af2c3/pgx_ulid/README.md)
- [官方扩展源码](https://github.com/Creatif/creatif-backend/blob/7b0466a2b0f505c824d98c2276107531ff1af2c3/pgx_ulid/src/lib.rs)
- [官方变更日志](https://github.com/Creatif/creatif-backend/blob/7b0466a2b0f505c824d98c2276107531ff1af2c3/pgx_ulid/CHANGELOG.md)

`ulid` 提供二进制 PostgreSQL ULID 类型、普通与单调生成器、排序和哈希能力，以及到 UUID 与时间戳的隐式转换。它适合使用渲染为 26 字符 Crockford Base32 字符串的可排序 128 位标识符。

### 核心流程

普通生成不需要预加载设置。

```sql
CREATE EXTENSION ulid;

CREATE TABLE users (
  id ulid PRIMARY KEY DEFAULT gen_ulid(),
  name text NOT NULL,
  created_at timestamp
    GENERATED ALWAYS AS (id::timestamp) STORED
);

INSERT INTO users(name) VALUES ('Alice') RETURNING id, created_at;

SELECT *
FROM users
WHERE id = '01ARZ3NDEKTSV4RRFFQ69G5FAV'::ulid;
```

若要在共享生成器状态内保证单调生成，调用 `gen_monotonic_ulid` 前必须预加载动态库并重启 PostgreSQL。

```conf
shared_preload_libraries = 'ulid'
```

```sql
SELECT gen_monotonic_ulid();
```

### 主要对象

- `ulid` 是带文本输入输出、比较、哈希与排序支持的 128 位类型。
- `gen_ulid` 创建带随机部分的新时间型标识符。
- `gen_monotonic_ulid` 在同一毫秒内生成多个 ID 时递增前一个共享值。
- `ulid_from_uuid` 与 `ulid_to_uuid` 支持 `uuid` 和 `ulid` 间的隐式转换。
- `ulid_to_timestamp` 提取毫秒时间戳；`timestamp_to_ulid` 生成随机部分为零的范围下界。

```sql
SELECT *
FROM users
WHERE id BETWEEN '2026-07-01'::timestamp::ulid
             AND '2026-07-02'::timestamp::ulid;
```

### 运维说明

只有单调生成器依赖共享内存并获取排他 LWLock；扩展其余部分无需预加载即可工作。单调值溢出会报错。版本 `0.1.3` 增加了时间戳到 ULID 的转换；本次复核的清单提供 PostgreSQL `14`、`15` 与 `16` 构建特性。应测试混合 `uuid`、`ulid` 与时间戳表达式中的隐式转换行为，避免便利转换掩盖意外类型变化。
