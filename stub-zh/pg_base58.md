

## 用法

> [pg_base58: PostgreSQL 的 Base58 编解码函数](https://github.com/Fell-x27/pg_base58)

提供使用 Base58 编码对数据进行编码和解码的函数。

```sql
CREATE EXTENSION pg_base58;
```

### 函数

| 函数 | 说明 |
|---|---|
| `base58_encode(bytea)` | 将 bytea 数据编码为 Base58 文本 |
| `base58_decode(text)` | 将 Base58 文本解码为 bytea |

### 示例

```sql
-- 将字符串编码为 Base58
SELECT base58_encode('hello'::bytea);
-- 'Cn8eVZg'

-- 将 Base58 字符串解码
SELECT base58_decode('Cn8eVZg');
-- '\x68656c6c6f'  (即 'hello')

-- 往返转换
SELECT convert_from(base58_decode(base58_encode('hello'::bytea)), 'UTF8');
-- 'hello'
```
