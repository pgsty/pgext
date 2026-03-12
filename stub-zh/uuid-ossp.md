

## 用法

> [uuid-ossp: UUID 生成函数](https://www.postgresql.org/docs/current/uuid-ossp.html)

提供使用多种标准算法生成 UUID 的函数。注意：如果只需要随机 UUID，可以考虑使用内置的 `gen_random_uuid()` 函数。

```sql
CREATE EXTENSION "uuid-ossp";
```

### UUID 生成函数

| 函数 | 说明 |
|---|---|
| `uuid_generate_v1()` | 版本 1：MAC 地址 + 时间戳 |
| `uuid_generate_v1mc()` | 版本 1，使用随机多播 MAC |
| `uuid_generate_v3(namespace uuid, name text)` | 版本 3：命名空间 + 名称的 MD5 哈希 |
| `uuid_generate_v4()` | 版本 4：完全随机 |
| `uuid_generate_v5(namespace uuid, name text)` | 版本 5：命名空间 + 名称的 SHA-1 哈希（优先于 v3） |

### 命名空间常量

| 函数 | 说明 |
|---|---|
| `uuid_nil()` | 空 UUID（全零） |
| `uuid_ns_dns()` | DNS 命名空间 |
| `uuid_ns_url()` | URL 命名空间 |
| `uuid_ns_oid()` | ISO OID 命名空间 |
| `uuid_ns_x500()` | X.500 DN 命名空间 |

### 示例

```sql
-- 随机 UUID（v4）
SELECT uuid_generate_v4();

-- 基于时间戳的 UUID（v1）
SELECT uuid_generate_v1();

-- 基于名称的确定性 UUID（v5，优先于 v3）
SELECT uuid_generate_v5(uuid_ns_url(), 'http://www.postgresql.org');

-- 用作默认主键
CREATE TABLE items (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  name text
);
```
