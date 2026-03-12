


## 用法

> [jsquery: 用于 jsonb 检索的数据类型](https://github.com/postgrespro/jsquery)

JsQuery 提供了一种 JSONB 数据查询语言，类似于 tsquery 之于全文搜索的作用。它提供了一种简洁的方式来搜索嵌套对象和数组，并支持通过 GIN 建立索引。

### 运算符

| 运算符 | 描述 |
|--------|------|
| `@@` | 匹配运算符：测试 jsonb 值是否匹配 jsquery 表达式 |

### 查询语法

表达式遵循 `路径 运算符 值` 的模式：

**二元运算符：**
- `=`（相等）、`>`、`>=`、`<`、`<=`（比较）
- `IN`（列表成员）
- `&&`（重叠）、`@>`（包含）、`<@`（被包含）

**一元运算符：**
- `= *`（存在性检查）
- `IS ARRAY`、`IS NUMERIC`、`IS OBJECT`、`IS STRING`、`IS BOOLEAN`（类型检查）

### 路径表达式

| 符号 | 含义 |
|------|------|
| `#` | 任意数组索引 |
| `#N` | 特定数组索引 N |
| `%` | 任意对象键 |
| `*` | 任意键/索引序列 |
| `@#` | 数组/对象长度 |
| `$` | 整个文档 |

"全部"语义（所有元素必须匹配）：
- `#:` -- 所有数组元素
- `%:` -- 所有对象键
- `*:` -- 所有嵌套路径

### 示例

简单值匹配：

```sql
SELECT * FROM jsonb_table WHERE data @@ 'name = "Alice"';
SELECT * FROM jsonb_table WHERE data @@ 'age > 21';
SELECT * FROM jsonb_table WHERE data @@ 'tags.#: IS STRING';
```

逻辑组合：

```sql
SELECT * FROM jsonb_table WHERE data @@ 'a = 1 AND (b = 2 OR c = 3)';
```

数组元素匹配（查找同时满足两个条件的数组元素）：

```sql
SELECT * FROM jsonb_table WHERE data @@ '#(a = 1 AND b = 2)';
```

对象键范围匹配：

```sql
SELECT * FROM jsonb_table WHERE data @@ '%($ >= 10 AND $ <= 20)';
```

### GIN 索引

两种运算符类适用于不同的查询模式：

```sql
-- 最适合已知完整路径时的范围和精确搜索
CREATE INDEX ON jsonb_table USING gin (data jsonb_path_value_ops);

-- 最适合精确值搜索；支持路径中使用 % 和 *
CREATE INDEX ON jsonb_table USING gin (data jsonb_value_path_ops);
```

索引使用的优化提示：

```sql
SELECT * FROM jsonb_table WHERE data @@ 'x = 1 /*-- index */ AND y = 2';
SELECT * FROM jsonb_table WHERE data @@ 'x = 1 /*-- noindex */ AND y = 2';
```

### 通过 CHECK 约束进行模式校验

```sql
CREATE TABLE documents (
    id serial PRIMARY KEY,
    data jsonb CHECK (data @@ 'name IS STRING AND similar_ids.#: IS NUMERIC'::jsquery)
);
```
