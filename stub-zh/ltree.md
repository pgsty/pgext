

## 用法

> [ltree: 层次树状标签数据类型](https://www.postgresql.org/docs/current/ltree.html)

`ltree` 扩展提供了用于层次树状结构数据的数据类型，具有丰富的搜索功能。

```sql
CREATE EXTENSION ltree;
```

### 数据类型

- **`ltree`**：标签路径（例如 `Top.Science.Astronomy`）
- **`lquery`**：用于匹配 ltree 值的正则式模式
- **`ltxtquery`**：类似全文搜索的模式

### 基本用法

```sql
CREATE TABLE categories (path ltree);
INSERT INTO categories VALUES
    ('Top'), ('Top.Science'), ('Top.Science.Astronomy'),
    ('Top.Hobbies'), ('Top.Collections.Pictures');

-- 查找后代
SELECT path FROM categories WHERE path <@ 'Top.Science';

-- 模式匹配
SELECT path FROM categories WHERE path ~ '*.Astronomy.*';

-- 全文搜索
SELECT path FROM categories WHERE path @ 'Science & !Pictures';
```

### 运算符

| 运算符 | 描述 |
|--------|------|
| `@>` | 是祖先（或相等） |
| `<@` | 是后代（或相等） |
| `~` | 匹配 lquery 模式 |
| `?` | 匹配数组中的任一 lquery |
| `@` | 匹配 ltxtquery |
| `\|\|` | 连接路径 |

### lquery 模式

```sql
'*.Science.*'           -- 包含 Science 的任意路径
'Top.*{1,2}.Astronomy'  -- Top 和 Astronomy 之间有 1-2 个标签
'*.astro*'              -- 前缀匹配
'*.Astro*@'             -- 不区分大小写的前缀匹配
```

### 函数

```sql
SELECT nlevel('Top.Science.Astronomy');                     -- 3
SELECT subltree('Top.Science.Astronomy', 1, 2);            -- Science
SELECT subpath('Top.Science.Astronomy', 1);                 -- Science.Astronomy
SELECT index('a.b.c.d', 'b.c');                             -- 1
SELECT lca('1.2.3', '1.2.3.4.5');                          -- 1.2
SELECT lca(ARRAY['1.2.3'::ltree, '1.2.4'::ltree]);        -- 1.2
```

### 索引支持

```sql
-- GiST 索引（支持 @>、<@、~、?、@）
CREATE INDEX path_gist_idx ON categories USING gist (path);

-- B-tree 索引（支持 <、<=、=、>=、>）
CREATE INDEX path_idx ON categories USING btree (path);
```
