
## 用法

> [pg_textsearch: 使用 BM25 的 PostgreSQL 现代排序文本搜索](https://github.com/timescale/pg_textsearch)

`pg_textsearch` 提供基于 BM25 评分并结合 Block-Max WAND 优化的现代排序文本搜索。它语法简洁，支持快速 top-k 查询、并行索引构建以及分区表。

添加到 `shared_preload_libraries`：

```
shared_preload_libraries = 'pg_textsearch'
```

```sql
CREATE EXTENSION pg_textsearch;
```

### 快速开始

```sql
CREATE TABLE documents (id bigserial PRIMARY KEY, content text);
INSERT INTO documents (content) VALUES
    ('PostgreSQL is a powerful database system'),
    ('BM25 is an effective ranking function'),
    ('Full text search with custom scoring');

-- 创建 BM25 索引
CREATE INDEX docs_idx ON documents USING bm25(content) WITH (text_config='english');

-- 使用 <@> 运算符查询（返回负 BM25 分数，分数越低匹配越好）
SELECT * FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;
```

### 查询

```sql
-- 从列自动检测索引
SELECT * FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;

-- 显式指定索引
SELECT * FROM documents
WHERE content <@> to_bm25query('database system', 'docs_idx') < -1.0;
```

### 过滤

**前置过滤**会在评分前缩减行数，适合选择性强的过滤条件：

```sql
CREATE INDEX ON documents (category_id);
SELECT * FROM documents
WHERE category_id = 123
ORDER BY content <@> 'search terms'
LIMIT 10;
```

**后置过滤**会先执行 BM25 扫描，再应用过滤条件：

```sql
SELECT * FROM documents
WHERE content <@> to_bm25query('search terms', 'docs_idx') < -5.0
ORDER BY content <@> 'search terms'
LIMIT 10;
```

### 索引选项

| 选项 | 默认值 | 说明 |
|------|--------|------|
| `text_config` | （必需） | PostgreSQL 文本搜索配置 |
| `k1` | 1.2 | 词频饱和参数 |
| `b` | 0.75 | 长度归一化参数 |

```sql
CREATE INDEX ON documents USING bm25(content)
  WITH (text_config='english', k1=1.5, b=0.8);

-- 语言专用配置
CREATE INDEX ON french_docs USING bm25(content) WITH (text_config='french');
CREATE INDEX ON german_docs USING bm25(content) WITH (text_config='german');
```

### 数据类型

**bm25query** - 表示 BM25 评分查询：

```sql
SELECT to_bm25query('search query text', 'docs_idx');
-- docs_idx:search query text
```
