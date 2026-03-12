

## 用法

> [xxhash: PostgreSQL 的 xxHash 哈希函数](https://github.com/hatarist/pg_xxhash)

提供 xxHash 非加密哈希函数，返回十六进制字符串或 bytea。

### 文本输出函数

```sql
SELECT xxh32('https://example.com');     -- 'ba15a4a8'
SELECT xxh64('https://example.com');     -- 'b131752760b48654'
SELECT xxh3_64('https://example.com');   -- '9398cc7c078760e6'
SELECT xxh128('https://example.com');    -- '4879d6aa9d88e9c7a169c008892d4829'
```

### 二进制输出函数

```sql
SELECT xxh32b('https://example.com');    -- \xba15a4a8
SELECT xxh64b('https://example.com');    -- \xb131752760b48654
SELECT xxh3_64b('https://example.com');  -- \x9398cc7c078760e6
SELECT xxh128b('https://example.com');   -- \x4879d6aa9d88e9c7a169c008892d4829
```

### 可用函数

| 函数 | 位数 | 输出 |
|------|------|------|
| `xxh32(text)` | 32 | 十六进制文本 |
| `xxh64(text)` | 64 | 十六进制文本 |
| `xxh3_64(text)` | 64 | 十六进制文本 |
| `xxh128(text)` | 128 | 十六进制文本 |
| `xxh32b(text)` | 32 | bytea |
| `xxh64b(text)` | 64 | bytea |
| `xxh3_64b(text)` | 64 | bytea |
| `xxh128b(text)` | 128 | bytea |
