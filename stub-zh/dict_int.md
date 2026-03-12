

## 用法

> [dict_int: 全文搜索的整数词典](https://www.postgresql.org/docs/current/dict-int.html)

控制全文搜索中整数的索引方式，防止因唯一词过多导致搜索性能下降。

```sql
CREATE EXTENSION dict_int;
```

### 配置参数

| 参数 | 说明 | 默认值 |
|---|---|---|
| `maxlen` | 允许的最大位数 | 6 |
| `rejectlong` | 为 `true` 时拒绝超长整数（作为停用词），为 `false` 时截断 | `false` |
| `absval` | 为 `true` 时在应用 maxlen 之前去除前导 `+`/`-` | `false` |

### 示例

```sql
-- 测试默认词典
SELECT ts_lexize('intdict', '12345678');
-- {123456}  （默认截断为 6 位）

-- 配置为拒绝过长整数
ALTER TEXT SEARCH DICTIONARY intdict (MAXLEN = 4, REJECTLONG = true);
SELECT ts_lexize('intdict', '12345');
-- {}  （作为停用词被拒绝）

SELECT ts_lexize('intdict', '1234');
-- {1234}  （接受）

-- 应用到文本搜索配置
ALTER TEXT SEARCH CONFIGURATION english
  ALTER MAPPING FOR int, uint WITH intdict;
```
