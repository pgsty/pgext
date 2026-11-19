## 用法

来源：

- [3.0 changes](https://github.com/s0l0ist/ferroid/blob/3.0.0/CHANGELOG.md)
- [PostgreSQL 扩展 README](https://github.com/s0l0ist/ferroid/blob/3.0.0/crates/pg-ferroid/README.md)
- [扩展 control 文件](https://github.com/s0l0ist/ferroid/blob/3.0.0/crates/pg-ferroid/pg_ferroid.control)
- [类型、转换与函数实现](https://github.com/s0l0ist/ferroid/blob/3.0.0/crates/pg-ferroid/src/lib.rs)
- [包元数据](https://github.com/s0l0ist/ferroid/blob/3.0.0/crates/pg-ferroid/Cargo.toml)

`pg_ferroid` 3.0.0 是一个 pgrx 扩展，增加原生 16 字节 `ulid` 类型。ULID 编码一个毫秒时间戳和 80 位尾部，使用 26 字符 Base32 表示，并支持比较、B-tree 与哈希索引、验证和显式转换。

### 生成并存储 ULID

```sql
CREATE EXTENSION pg_ferroid;

CREATE TABLE users (
  id ulid PRIMARY KEY DEFAULT gen_ulid_mono(),
  name text NOT NULL
);

INSERT INTO users (name) VALUES ('alice'), ('bob');

SELECT id, id::timestamptz AS created_at, name
FROM users
ORDER BY id;

SELECT ulid_is_valid('01JEPY8K5V3XQZW6M9N7P8Q2RT');
```

`gen_ulid()` 生成带随机尾部、按时间排序的值。`gen_ulid_mono()` 会在同一毫秒内递增尾部，但其单调状态属于单个后端或线程，并非整个集群。不要把它当成全局序列。

### 排序与信息暴露

显式转换把 `ulid` 与 `text`、`bytea` 和 `timestamptz` 连接起来。把时间戳转换为 `ulid` 时，随机部分会置零，适合作为范围下界，但不会生成正常的应用标识符。

内嵌时间戳会以毫秒精度暴露大致创建时间。若标识符泄露不能暴露年龄或顺序，则不应使用 ULID。该类型使用原生扩展表示，因此跨 PostgreSQL 或扩展版本时应测试二进制升级、逻辑解码、转储恢复、驱动和索引重建。根据 control 文件，3.0.0 版仅限超级用户安装且不属于 trusted 扩展。

### 3.0 版本边界

要求 PostgreSQL 15–18，已移除 PostgreSQL 13/14 支持。`ulid` 与 `timestamp without time zone` 之间的直接转换已移除，应明确时区并经 `timestamptz` 转换：

```sql
SELECT (timestamp '2025-01-01 12:00:00' AT TIME ZONE 'UTC')::ulid;
```

时间戳必须有限、不早于 Unix epoch，且在 ULID 的 48 位时间范围内。时钟回拨会延迟单调生成，但取消请求和语句超时仍有效。匹配的 Rust/pgrx 构建使用 pgrx 0.20.0。
