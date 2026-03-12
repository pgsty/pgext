

## 用法

> [base62: PostgreSQL 的 base62 编解码数据类型](https://github.com/adjust/pg-base62)

提供使用 base62 方案（0-9、A-Z、a-z）进行编码和解码的数据类型。

```sql
CREATE EXTENSION base62;
```

### 类型

| 类型 | 存储 | 最大字符串长度 | 最大数值 |
|---|---|---|---|
| `base62` | 4 字节（int） | 6 字符 | 2,147,483,647 |
| `bigbase62` | 8 字节（bigint） | 11 字符 | 9,223,372,036,854,775,807 |
| `hugebase62` | 16 字节 | 20 字符 | （bytea 转换） |

### 示例

```sql
-- 编码/解码 base62
SELECT 2147483647::base62;          -- '2LKcb1'
SELECT '2LKcb1'::base62::int;      -- 2147483647

-- 用于更大数值的 bigbase62
SELECT 9223372036854775807::bigbase62;           -- 'AzL8n0Y58m7'
SELECT 'AzL8n0Y58m7'::bigbase62::bigint;        -- 9223372036854775807

-- hugebase62 与 bytea 转换
SELECT 'AzL8n0Y58m7AzL8n0Y58'::hugebase62;
SELECT 'AzL8n0Y58m7AzL8n0Y58'::hugebase62::bytea;
SELECT '\x960c06065a6ed8ffff1e7149f40b1800'::bytea::hugebase62;

-- 注意：base62 区分大小写
SELECT '2lkcb'::base62::int;   -- 40933305
SELECT '2LKCB'::base62::int;   -- 34635195
```
