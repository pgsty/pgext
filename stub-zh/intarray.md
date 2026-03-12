

## 用法

> [intarray: 带索引支持的整数数组函数和运算符](https://www.postgresql.org/docs/current/intarray.html)

提供操作无空值整数数组的函数和运算符，支持 GiST 和 GIN 索引以实现快速数组搜索。

```sql
CREATE EXTENSION intarray;
```

### 函数

| 函数 | 说明 | 示例 |
|---|---|---|
| `icount(int[])` | 元素数量 | `icount('{1,2,3}')` -- 3 |
| `sort(int[], dir)` | 排序数组（`'asc'` 或 `'desc'`） | `sort('{3,1,2}','asc')` -- `{1,2,3}` |
| `sort_asc(int[])` | 升序排序 | `sort_asc('{3,1,2}')` -- `{1,2,3}` |
| `sort_desc(int[])` | 降序排序 | `sort_desc('{3,1,2}')` -- `{3,2,1}` |
| `uniq(int[])` | 去除相邻重复元素 | `uniq(sort('{1,2,3,2,1}'))` -- `{1,2,3}` |
| `idx(int[], item)` | 第一次匹配的索引 | `idx('{11,22,33}', 22)` -- 2 |
| `subarray(int[], start, len)` | 提取子数组 | `subarray('{1,2,3,4}', 2, 2)` -- `{2,3}` |
| `intset(int)` | 创建单元素数组 | `intset(42)` -- `{42}` |

### 运算符

| 运算符 | 说明 |
|---|---|
| `&&` | 数组重叠（有共同元素） |
| `@>` | 左数组包含右数组 |
| `<@` | 左数组被右数组包含 |
| `#` | 元素数量 |
| `+` | 数组连接 / 追加元素 |
| `-` | 移除元素 |
| `\|` | 数组并集 |
| `&` | 数组交集 |
| `@@` | 数组匹配查询表达式 |
| `~~` | 查询表达式匹配数组 |

### 索引支持

```sql
-- 用于数组包含/重叠查询的 GiST 索引
CREATE INDEX idx ON messages USING GIST (tags gist__intbig_ops);

-- GIN 索引（备选方案）
CREATE INDEX idx ON messages USING GIN (tags gin__int_ops);

-- 带索引支持的查询
SELECT * FROM messages WHERE tags && '{1,2}';     -- 重叠
SELECT * FROM messages WHERE tags @> '{1,2}';     -- 包含
SELECT * FROM messages WHERE tags @@ '1&(2|3)';  -- 查询表达式
```
