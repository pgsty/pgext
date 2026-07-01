


## 用法

> 来源：[pg_pinyin upstream README](https://github.com/aiyou178/pg_pinyin)、[Chinese README](https://github.com/aiyou178/pg_pinyin/blob/main/README.zh-CN.md)、[local metadata](../db/extension.csv)。

`pg_pinyin` 将中文文本转换为拼音，可以按字符转换，也可以按词转换。它适用于生成搜索列、trigram 搜索，以及需要拼音输入的 `pg_search` BM25 查询。

```sql
CREATE EXTENSION pg_pinyin;
```

### 函数

| 函数 | 说明 |
|------|------|
| `pinyin_char_romanize(text)` | 字符级拼音罗马化 |
| `pinyin_char_romanize(text, suffix text)` | 使用自定义词典后缀进行字符级罗马化 |
| `pinyin_word_romanize(text)` | 词级拼音罗马化 |
| `pinyin_word_romanize(text, suffix text)` | 使用自定义词典后缀进行词级罗马化 |
| `pinyin_word_romanize(tokenizer_input anyelement)` | 从 `pg_search` tokenizer 输入进行词级罗马化，例如 `name::pdb.icu::text[]` |
| `pinyin_word_romanize(tokenizer_input anyelement, suffix text)` | 使用自定义词典后缀处理 tokenizer 输入 |
| `pinyin_regex_phrase(text, slope integer DEFAULT NULL, max_expansions integer DEFAULT NULL, generated_pinyin boolean DEFAULT false)` | 返回 `pdb.query` 的 `pg_search` 查询辅助函数；仅在 `CREATE EXTENSION pg_pinyin` 前已启用 `pg_search` 时可用 |
| `pinyin_regex_phrase_patterns(text, generated_pinyin boolean DEFAULT false)` | 内部辅助函数，以 `text[]` 返回正则短语 token |

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

### 分词 + pg_search

面向词的搜索使用 `pinyin_word_romanize`。当 `pg_search` 可用时，它可以消费 `pdb.icu::text[]` 这类 tokenizer 输入。

```sql
CREATE EXTENSION IF NOT EXISTS pg_search;
CREATE EXTENSION IF NOT EXISTS pg_pinyin;

CREATE TABLE voice (
  id bigserial PRIMARY KEY,
  description text NOT NULL,
  pinyin text GENERATED ALWAYS AS (public.pinyin_word_romanize(description)) STORED
);

CREATE INDEX voice_pinyin_bm25_idx
ON voice
USING bm25 (id, pinyin)
WITH (key_field='id');

SELECT *
FROM voice
WHERE pinyin @@@ public.pinyin_regex_phrase('zhengshuang');

SELECT public.pinyin_word_romanize('郑爽ABC'::pdb.icu::text[]);
```

`pinyin_regex_phrase` 的返回类型是 `pdb.query`，因此必须先在数据库中启用 `pg_search`，再创建 `pg_pinyin`。如果先创建 `pg_pinyin`，上游文档说明罗马化函数仍会安装，但 `pinyin_regex_phrase` 会以带有清晰异常信息的错误占位函数安装。

### 字典表

执行 `CREATE EXTENSION pg_pinyin` 时，扩展会在 `pinyin` schema 下初始化内置字典表；正常使用扩展时不需要额外加载数据。内置数据覆盖字符映射、词 token 和词映射。

可以在 `pinyin` schema 中提供带后缀的自定义词典表。调用时指定该后缀会把基础字典和后缀表合并，且后缀表条目优先。

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

INSERT INTO pinyin.pinyin_words_suffix1 (word, pinyin)
VALUES ('郑爽', '|zhengx| |shuangx|')
ON CONFLICT (word) DO UPDATE SET pinyin = EXCLUDED.pinyin;

SELECT public.pinyin_char_romanize('郑爽ABC', '_suffix1');
SELECT public.pinyin_word_romanize('郑爽ABC'::pdb.icu::text[], '_suffix1');
```
