

## 用法

> [pg_trgm: 基于三元组的文本相似度度量和索引搜索](https://www.postgresql.org/docs/current/pgtrgm.html)

`pg_trgm` 模块提供了基于三元组匹配来确定字母数字文本相似度的函数和运算符，以及用于快速字符串相似度搜索的索引运算符类。

```sql
CREATE EXTENSION pg_trgm;
```

三元组是从字符串中提取的三个连续字符组。如果两个字符串共享许多三元组，则它们是相似的。

### 函数

| 函数 | 说明 |
|------|------|
| `similarity(text, text)` → `real` | 返回 0 到 1 之间的相似度 |
| `show_trgm(text)` → `text[]` | 返回字符串中所有三元组的数组 |
| `word_similarity(text, text)` → `real` | 第一个字符串与第二个中最相似词的相似度 |
| `strict_word_similarity(text, text)` → `real` | 类似但词边界匹配更严格 |
| `show_limit()` → `real` | *（已弃用）* 返回 `pg_trgm.similarity_threshold` |
| `set_limit(real)` → `real` | *（已弃用）* 设置 `pg_trgm.similarity_threshold` |

```sql
SELECT similarity('word', 'two words');
-- 0.36363637

SELECT show_trgm('word');
-- {"  w"," wo",ord,"rd ",wor}
```

### 运算符

| 运算符 | 说明 |
|--------|------|
| `text % text` → `boolean` | 相似度 > `pg_trgm.similarity_threshold` 时为 true |
| `text <% text` → `boolean` | 词相似度 > `pg_trgm.word_similarity_threshold` 时为 true |
| `text %> text` → `boolean` | `<%` 的交换子 |
| `text <<% text` → `boolean` | 严格词相似度 > 阈值时为 true |
| `text %>> text` → `boolean` | `<<%` 的交换子 |
| `text <-> text` → `real` | 距离（1 - 相似度） |
| `text <<-> text` → `real` | 词距离（1 - 词相似度） |
| `text <->> text` → `real` | `<<->` 的交换子 |
| `text <<<-> text` → `real` | 严格词距离 |
| `text <->>> text` → `real` | `<<<->` 的交换子 |

### GUC 参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `pg_trgm.similarity_threshold` | 0.3 | `%` 运算符的阈值 |
| `pg_trgm.word_similarity_threshold` | 0.6 | `<%` 和 `%>` 运算符的阈值 |
| `pg_trgm.strict_word_similarity_threshold` | 0.5 | `<<%` 和 `%>>` 运算符的阈值 |

### 索引支持

GiST 和 GIN 索引支持相似度运算符：

```sql
-- GIN 索引（查找更快，构建更慢）
CREATE INDEX trgm_idx ON test_trgm USING GIN (t gin_trgm_ops);

-- GiST 索引（支持 KNN 的距离运算符）
CREATE INDEX trgm_idx ON test_trgm USING GIST (t gist_trgm_ops);

-- 自定义签名长度的 GiST 索引
CREATE INDEX trgm_idx ON test_trgm USING GIST (t gist_trgm_ops(siglen=32));
```

### 文本搜索示例

使用三元组索引加速 `LIKE` / `ILIKE` / 正则查询：

```sql
SELECT t, similarity(t, 'word') AS sml
FROM test_trgm
WHERE t % 'word'
ORDER BY sml DESC, t;

-- 使用距离运算符的 KNN 搜索
SELECT t, t <-> 'word' AS dist
FROM test_trgm
ORDER BY dist
LIMIT 10;
```

GIN 和 GiST 三元组索引也能自动加速 `LIKE`、`ILIKE`、`~` 和 `~*` 查询。
