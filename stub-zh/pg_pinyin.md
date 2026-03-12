

## 用法

> [pg_pinyin: PostgreSQL 的拼音转换与搜索辅助](https://github.com/aiyou178/pg_pinyin)

将中文字符转换为拼音罗马化，用于搜索和索引。可与 `pg_trgm` 配合实现模糊拼音搜索，或与 `pg_search` 配合实现基于词的搜索。

```sql
CREATE EXTENSION pg_pinyin;
```

### 函数

| 函数 | 说明 |
|------|------|
| `pinyin_char_romanize(text)` | 字符级拼音罗马化 |
| `pinyin_char_romanize(text, suffix)` | 使用自定义词典后缀 |
| `pinyin_word_romanize(text)` | 词级拼音罗马化 |
| `pinyin_word_romanize(text, suffix)` | 使用自定义词典后缀 |

### 生成列 + Trigram 搜索

```sql
CREATE EXTENSION IF NOT EXISTS pg_pinyin;
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE voice (
  id bigserial PRIMARY KEY,
  description text NOT NULL,
  pinyin text GENERATED ALWAYS AS (public.pinyin_char_romanize(description)) STORED
);

CREATE INDEX voice_pinyin_trgm_idx ON voice USING gin (pinyin gin_trgm_ops);

INSERT INTO voice (description) VALUES ('郑爽ABC');
SELECT id, description, pinyin FROM voice;
```

### 自定义词典

在 `pinyin` schema 中提供带后缀的自定义词典表：

```sql
CREATE TABLE IF NOT EXISTS pinyin.pinyin_mapping_suffix1 (
  character text PRIMARY KEY,
  pinyin text NOT NULL
);

CREATE TABLE IF NOT EXISTS pinyin.pinyin_words_suffix1 (
  word text PRIMARY KEY,
  pinyin text NOT NULL
);

INSERT INTO pinyin.pinyin_mapping_suffix1 (character, pinyin)
VALUES ('郑', '|zhengx|')
ON CONFLICT (character) DO UPDATE SET pinyin = EXCLUDED.pinyin;

-- 使用自定义词典
SELECT public.pinyin_char_romanize('郑爽ABC', '_suffix1');
```
