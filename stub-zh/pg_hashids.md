

## 用法

> [pg_hashids: 从整数生成短唯一 ID](https://github.com/iCyberon/pg_hashids)

将类似 347 的数字转换为类似 "yr8" 的字符串，适用于混淆数据库主键。

```sql
CREATE EXTENSION pg_hashids;
```

### 函数

| 函数 | 描述 |
|---|---|
| `id_encode(number [, salt [, min_length [, alphabet]]])` | 将整数编码为 hashid 字符串 |
| `id_decode(hash, salt, min_length [, alphabet])` | 将 hashid 字符串解码为整数数组 |
| `id_decode_once(hash [, salt [, min_length [, alphabet]]])` | 将 hashid 字符串解码为单个整数 |

### 示例

```sql
-- 基本编码
SELECT id_encode(1001);                                    -- 'jNl'
SELECT id_encode(1234567, 'This is my salt');              -- 'Pdzxp'
SELECT id_encode(1234567, 'This is my salt', 10);          -- 'PlRPdzxpR7'

-- 自定义字母表
SELECT id_encode(1234567, 'This is my salt', 10, 'abcdefghijABCDxFGHIJ1234567890');
-- '3GJ956J9B9'

-- 解码（必须使用相同的盐值！）
SELECT id_decode('PlRPdzxpR7', 'This is my salt', 10);     -- 1234567
SELECT id_decode_once('jNl');                               -- 1001
SELECT id_decode_once('Pdzxp', 'This is my salt');          -- 1234567
```
