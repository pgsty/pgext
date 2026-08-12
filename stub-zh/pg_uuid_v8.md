## 用法

来源：

- [PGXN 上的 pg_uuid_v8 1.1.0](https://pgxn.org/dist/pg_uuid_v8/1.1.0/)
- [pg_uuid_v8 1.1.0 README](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/README.md)
- [pg_uuid_v8 1.1.0 控制文件](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/pg_uuid_v8.control)
- [pg_uuid_v8 1.0 基础 SQL](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/pg_uuid_v8--1.0.sql)
- [pg_uuid_v8 1.0 至 1.1 升级 SQL](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/pg_uuid_v8--1.0--1.1.sql)
- [Pigsty pg_uuid_v8 软件包矩阵](https://pgext.cloud/ext/pg_uuid_v8)

`pg_uuid_v8` 1.1.0 生成带有 UUID-v4 版本位与变体位的 UUID，同时在随机载荷中嵌入经过混淆的创建时间。它的 `uuid_v8_*` 便捷函数与底层 `uuid_stego_*` API 对应。适合需要提取隐藏时间并建立时间范围索引的场景，但不要把嵌入值当成认证令牌，也不要用它替代独立、可信的创建时间列。

### 生成值

```sql
CREATE EXTENSION pg_uuid_v8;

SELECT uuid_v8_set_seed('replace-with-a-unique-secret');
SELECT uuid_v8_set_encryption_mode('AES128');

CREATE TABLE events (
  id uuid PRIMARY KEY DEFAULT uuid_v8_generate(),
  data jsonb,
  created_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO events(data) VALUES ('{"type":"login"}');
```

上游实现默认使用公开的内置 seed 与 `XOR` 模式。生成值之前，应设置当前部署独有的秘密。也可选择 `AES128` 和 `AES256`，但提取值时必须选用相同的 seed 与模式。

### 提取隐藏时间并建立索引

```sql
SELECT
  uuid_v8_extract_timestamp(id) AS epoch_microseconds,
  stego_time_to_timestamp(uuid_v8_extract_timestamp(id)) AS created_time
FROM events;

CREATE INDEX events_uuid_time_idx
ON events USING btree (uuid_v8_extract_timestamp(id));

SELECT *
FROM events
WHERE uuid_v8_extract_timestamp(id)
      BETWEEN timestamp_to_stego_time('2026-01-01'::timestamptz)
          AND timestamp_to_stego_time(now())
ORDER BY uuid_v8_extract_timestamp(id);
```

`uuid_v8_extract_timestamp(uuid)` 返回按微秒缩放的 `bigint`，从而继续兼容 `timestamp_to_stego_time()` 与 `stego_time_to_timestamp()`。在 1.1 版本中，内部 48 位字段存储毫秒，因此返回值只有毫秒分辨率，最后三位十进制数字始终为零。

`uuid_stego_in_range()` 提供布尔型时间戳范围辅助函数。对提取函数建立 B-tree 函数索引，是时间谓词走索引时明确且可预期的路径。

### 比较隐藏时间

`uuid_v8_compare(uuid, uuid)` 与 `uuid_stego_compare(uuid, uuid)` 按提取出的隐藏时间返回顺序。扩展还为 UUID 参数定义了 `<`、`<=`、`>` 与 `>=` 操作符。

Pigsty 软件包把这些新增操作符安装到 `public`，并限定其 commutator 与 negator 引用，以兼容 PostgreSQL 17 和 18。PostgreSQL 已有内置 UUID 排序操作符；必须明确使用隐藏时间语义时，应使用比较函数或带模式限定的 `OPERATOR(public.<)` 表达式。

### Seed 与模式控制

```sql
SELECT uuid_v8_set_seed('replace-with-a-unique-secret');
SELECT uuid_v8_get_seed();

SELECT uuid_v8_set_encryption_mode('XOR');
SELECT uuid_v8_set_encryption_mode('AES128');
SELECT uuid_v8_set_encryption_mode('AES256');
SELECT uuid_v8_get_encryption_mode();

ALTER SYSTEM SET uuid_v8.encryption_mode = 'AES128';
SELECT pg_reload_conf();
```

seed 由 `uuid_v8.stego_seed` 暴露，模式由 `uuid_v8.encryption_mode` 暴露。设置函数改变当前会话，配置参数可以为后续会话建立默认值。`uuid_v8_get_seed()` 会返回当前 seed，因此应相应限制数据库访问，并且绝不能记录其返回值。

### 升级与兼容性边界

```sql
ALTER EXTENSION pg_uuid_v8 UPDATE TO '1.1';
```

1.1 版本把时间戳存储从微秒改为毫秒。旧的 48 位微秒字段大约每 8.9 年回绕一次，无法可靠恢复当前绝对日期；48 位毫秒字段约可持续 8,925 年。1.1 之前的值，其相对顺序不受影响；但升级不会重写既有编码，因此这些旧值的绝对时间提取与范围谓词仍不可靠。

PGXN 元数据面向 PostgreSQL 12 或以上版本；当前 Pigsty 软件包覆盖 PostgreSQL 14–18。Pigsty 软件包把扩展固定在 `public` 并设为不可重定位，以便新增操作符一致解析。当数据来源审计、亚毫秒精度或跨 seed、跨模式迁移很重要时，应保留普通的 `created_at` 列。
