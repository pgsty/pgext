

## 用法

> [unaccent: 用于去除重音的全文搜索词典](https://www.postgresql.org/docs/current/unaccent.html)

去除文本中的重音符号（变音符号）。可用作全文搜索词典或独立函数。

```sql
CREATE EXTENSION unaccent;
```

### 函数

| 函数 | 说明 |
|---|---|
| `unaccent(text)` | 使用默认词典去除重音 |
| `unaccent(dictionary regdictionary, text)` | 使用指定词典去除重音 |

### 全文搜索用法

```sql
-- 测试词典
SELECT ts_lexize('unaccent', 'Hôtel');
-- {Hotel}

-- 创建不区分重音的法语文本搜索配置
CREATE TEXT SEARCH CONFIGURATION fr (COPY = french);
ALTER TEXT SEARCH CONFIGURATION fr
  ALTER MAPPING FOR hword, hword_part, word
  WITH unaccent, french_stem;

SELECT to_tsvector('fr', 'Hôtels de la Mer');
-- 'hotel':1 'mer':4

-- 不区分重音的搜索
SELECT to_tsvector('fr', 'Hôtel de la Mer') @@ to_tsquery('fr', 'Hotels');
-- true
```

### 独立函数用法

```sql
SELECT unaccent('Hôtel');           -- Hotel
SELECT unaccent('café');            -- cafe
SELECT unaccent('résumé');          -- resume
SELECT unaccent('niño');            -- nino
```
