

## 用法

> [dict_xsyn: 全文搜索的扩展同义词词典](https://www.postgresql.org/docs/current/dict-xsyn.html)

提供用于全文搜索的扩展同义词词典模板，将词语替换为同义词组。

```sql
CREATE EXTENSION dict_xsyn;
```

### 配置参数

| 参数 | 说明 | 默认值 |
|---|---|---|
| `matchorig` | 接受原始词语 | `true` |
| `matchsynonyms` | 接受同义词作为输入 | `false` |
| `keeporig` | 在输出中包含原始词语 | `true` |
| `keepsynonyms` | 在输出中包含同义词 | `true` |
| `rules` | `$SHAREDIR/tsearch_data/` 目录下同义词文件的基本名称（`.rules` 扩展名） | -- |

### 规则文件格式

```
word syn1 syn2 syn3
```

以 `#` 开头的行为注释。

### 示例

```sql
-- 配置词典
ALTER TEXT SEARCH DICTIONARY xsyn (RULES='my_rules', KEEPORIG=true);

-- 测试词典
SELECT ts_lexize('xsyn', 'word');
-- {word,syn1,syn2,syn3}

-- 同时将同义词作为输入进行匹配
ALTER TEXT SEARCH DICTIONARY xsyn (RULES='my_rules', MATCHSYNONYMS=true);
SELECT ts_lexize('xsyn', 'syn1');
-- {syn1,syn2,syn3}

-- 在文本搜索配置中使用
ALTER TEXT SEARCH CONFIGURATION english
  ALTER MAPPING FOR word, asciiword WITH xsyn, english_stem;
```
