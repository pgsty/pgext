

## 用法

> [base36: PostgreSQL 的 base36 编解码数据类型](https://github.com/adjust/pg-base36)

提供使用 base36 方案进行编码和解码的数据类型，支持 base36 与整数类型之间的转换。

```sql
CREATE EXTENSION base36;
```

### 类型

| 类型 | 存储 | 最大字符串长度 | 最大数值 |
|---|---|---|---|
| `base36` | 4 字节（int） | 6 字符 | 2,147,483,647 |
| `bigbase36` | 8 字节（bigint） | 13 字符 | 9,223,372,036,854,775,807 |

### 示例

```sql
-- 将整数编码为 base36
SELECT 1234567::base36;          -- 'qglj'

-- 将 base36 解码为整数
SELECT 'qglj'::base36::int;     -- 1234567

-- 用于更大数值的 bigbase36
SELECT 9223372036854775807::bigbase36;           -- 'i1y004og0svr'
SELECT 'i1y004og0svr'::bigbase36::bigint;        -- 9223372036854775807
```
