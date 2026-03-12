

## 用法

> [pg_bigm 文档](https://pgbigm.osdn.jp/pg_bigm_en-1-2.html) | [GitHub: pgbigm/pg_bigm](https://github.com/pgbigm/pg_bigm)

`pg_bigm` 模块为 PostgreSQL 提供全文搜索能力。该模块允许用户创建 **2-gram**（bigram）索引以加速全文搜索。

## 功能特性

- **Bigram 索引**：为文本列创建 2-gram（bigram）GIN 索引
- **更快的 LIKE 搜索**：加速 `LIKE` 查询，包括前缀、后缀和子串搜索
- **全语言支持**：无需额外配置即可支持所有语言，包括中日韩（CJK）
- **简单 API**：提供相似度搜索函数和运算符

## 函数与运算符

### 函数

| 函数 | 返回类型 | 说明 |
|------|----------|------|
| `likequery(text)` | `text` | 根据关键词生成全文搜索查询 |
| `show_bigm(text)` | `text[]` | 显示给定字符串中的所有 2-gram |
| `pg_gin_pending_stats(regclass)` | `record` | 返回 GIN 索引待处理列表中的页面数和元组数 |

### 运算符

| 运算符 | 说明 |
|--------|------|
| `text =% text` | 当左右操作数的相似度大于等于 `pg_bigm.similarity_limit` 时返回 true |

## GUC 参数

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `pg_bigm.last_update` | `text` | - | 显示模块的最后更新日期（只读） |
| `pg_bigm.enable_recheck` | `bool` | `on` | 控制索引扫描时是否执行复查 |
| `pg_bigm.gin_key_limit` | `int` | `0` | 限制全文搜索使用的最大 bigram 数量。0 表示无限制 |
| `pg_bigm.similarity_limit` | `real` | `0.3` | 设置 `=%` 运算符的最小相似度阈值 |

## 示例

### 基本全文搜索

```sql
-- 创建扩展
CREATE EXTENSION pg_bigm;

-- 创建含文本数据的表
CREATE TABLE documents (id serial PRIMARY KEY, content text);
INSERT INTO documents (content) VALUES
  ('PostgreSQL is a powerful database'),
  ('Full text search with bigram indexing'),
  ('Japanese text: 日本語テキスト検索');

-- 创建 bigram 索引
CREATE INDEX docs_bigm_idx ON documents USING gin (content gin_bigm_ops);

-- 使用 LIKE 搜索
SELECT * FROM documents WHERE content LIKE '%search%';

-- 使用 likequery 函数搜索
SELECT * FROM documents WHERE content LIKE likequery('database');
```

### 相似度搜索

```sql
-- 显示字符串的 bigram
SELECT show_bigm('PostgreSQL');

-- 相似度搜索
SET pg_bigm.similarity_limit = 0.2;
SELECT * FROM documents WHERE content =% 'database search';
```

### 与 pg_trgm 的比较

pg_bigm 相比内置的 `pg_trgm` 有以下优势：

| 特性 | pg_bigm | pg_trgm |
|------|---------|---------|
| N-gram 类型 | 2-gram（bigram） | 3-gram（trigram） |
| 最小搜索字符串 | 1 个字符 | 3 个字符 |
| 非字母语言 | 完全支持 | 有限支持 |
| LIKE 搜索类型 | 前缀、后缀和子串 | 前缀、后缀和子串 |

详细文档包括高级用法和性能调优，请参见 [pg_bigm 官方文档](https://pgbigm.osdn.jp/pg_bigm_en-1-2.html)。
