

## 用法

> [prefix: 用于电话号码路由的前缀范围类型](https://github.com/dimitri/prefix)

`prefix` 扩展提供了 `prefix_range` 数据类型，用于高效的前缀匹配，尤其适用于电话呼叫路由场景。

### 数据类型

使用构造函数或文本转换创建 `prefix_range` 值：

```sql
CREATE EXTENSION prefix;

SELECT prefix_range('123');           -- 123
SELECT prefix_range('123', '4', '5'); -- 123[4-5]
SELECT '123'::prefix_range;           -- 123
```

### 运算符

| 运算符 | 说明 |
|----------|-------------|
| `@>` | 包含（前缀包含值） |
| `<@` | 被包含于 |
| `&&` | 重叠 |
| `\|` | 并集 |
| `&` | 交集 |
| `=`, `<>`, `<`, `>`, `<=`, `>=` | 比较 |

### 示例

```sql
-- 查找电话号码的最长匹配前缀
SELECT * FROM prefixes
WHERE prefix @> '0123456789'
ORDER BY length(prefix) DESC
LIMIT 1;

-- 包含检查
SELECT '123'::prefix_range @> '123456';     -- true

-- 交集
SELECT '123[4-5]' & '123[2-7]';            -- 123[4-5]

-- 并集
SELECT '123' | '124';                       -- 12[3-4]
```

### 索引支持

创建 GiST 索引以实现高效的前缀查找：

```sql
CREATE INDEX idx_prefix ON prefixes USING gist(prefix);
```

该索引可加速 `@>`、`<@`、`&&` 和 `=` 运算符。
