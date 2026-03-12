

## 用法

> [pg_similarity](https://github.com/eulerto/pg_similarity)：PostgreSQL 相似度查询支持。
> 来源：[README.md](https://raw.githubusercontent.com/eulerto/pg_similarity/master/README.md)

**pg_similarity** 是支持 PostgreSQL 相似度查询的扩展。实现与 RDBMS 紧密集成，定义了运算符，你可以使用 `~~~` 和 `~!~` 等代替传统运算符（`=` 和 `<>`）。

**pg_similarity** 有三个主要组件：

- **函数**：实现文献中可用的相似度算法的函数集合。这些函数可作为 UDF 使用，并作为实现相似度运算符的基础；
- **运算符**：基于相似度函数定义的运算符集合。它们使用相似度函数获取相似度阈值，并与用户定义的阈值比较以决定是否匹配；
- **会话变量**：存储相似度函数参数的变量集合。这些变量可在运行时定义。


--------

## 函数与运算符

该扩展支持一系列相似度算法。最知名的算法均已覆盖。请注意每种算法适用于特定领域。提供以下算法：

- L1 距离（又称城市街区距离或曼哈顿距离）
- 余弦距离
- Dice 系数
- 欧几里得距离
- 汉明距离
- Jaccard 系数
- Jaro 距离
- Jaro-Winkler 距离
- Levenshtein 距离
- 匹配系数
- Monge-Elkan 系数
- Needleman-Wunsch 系数
- 重叠系数
- Q-Gram 距离
- Smith-Waterman 系数
- Smith-Waterman-Gotoh 系数
- Soundex 距离

| 算法 | 函数 | 运算符 | 支持索引 | 参数 |
|------|------|--------|----------|------|
| L1 距离 | `block(text, text) returns float8` | `~++` | 是 | `pg_similarity.block_tokenizer`, `pg_similarity.block_threshold`, `pg_similarity.block_is_normalized` |
| 余弦距离 | `cosine(text, text) returns float8` | `~##` | 是 | `pg_similarity.cosine_tokenizer`, `pg_similarity.cosine_threshold`, `pg_similarity.cosine_is_normalized` |
| Dice 系数 | `dice(text, text) returns float8` | `~-~` | 是 | `pg_similarity.dice_tokenizer`, `pg_similarity.dice_threshold`, `pg_similarity.dice_is_normalized` |
| 欧几里得距离 | `euclidean(text, text) returns float8` | `~!!` | 是 | `pg_similarity.euclidean_tokenizer`, `pg_similarity.euclidean_threshold`, `pg_similarity.euclidean_is_normalized` |
| 汉明距离 | `hamming(bit varying, bit varying) returns float8` / `hamming_text(text, text) returns float8` | `~@~` | 否 | `pg_similarity.hamming_threshold`, `pg_similarity.hamming_is_normalized` |
| Jaccard 系数 | `jaccard(text, text) returns float8` | `~??` | 是 | `pg_similarity.jaccard_tokenizer`, `pg_similarity.jaccard_threshold`, `pg_similarity.jaccard_is_normalized` |
| Jaro 距离 | `jaro(text, text) returns float8` | `~%%` | 否 | `pg_similarity.jaro_threshold`, `pg_similarity.jaro_is_normalized` |
| Jaro-Winkler 距离 | `jarowinkler(text, text) returns float8` | `~@@` | 否 | `pg_similarity.jarowinkler_threshold`, `pg_similarity.jarowinkler_is_normalized` |
| Levenshtein 距离 | `lev(text, text) returns float8` | `~==` | 否 | `pg_similarity.levenshtein_threshold`, `pg_similarity.levenshtein_is_normalized` |
| 匹配系数 | `matchingcoefficient(text, text) returns float8` | `~^^` | 是 | `pg_similarity.matching_tokenizer`, `pg_similarity.matching_threshold`, `pg_similarity.matching_is_normalized` |
| Monge-Elkan 系数 | `mongeelkan(text, text) returns float8` | `~\|\|` | 否 | `pg_similarity.mongeelkan_tokenizer`, `pg_similarity.mongeelkan_threshold`, `pg_similarity.mongeelkan_is_normalized` |
| Needleman-Wunsch 系数 | `needlemanwunsch(text, text) returns float8` | `~#~` | 否 | `pg_similarity.nw_threshold`, `pg_similarity.nw_is_normalized` |
| 重叠系数 | `overlapcoefficient(text, text) returns float8` | `~**` | 是 | `pg_similarity.overlap_tokenizer`, `pg_similarity.overlap_threshold`, `pg_similarity.overlap_is_normalized` |
| Q-Gram 距离 | `qgram(text, text) returns float8` | `~~~` | 是 | `pg_similarity.qgram_threshold`, `pg_similarity.qgram_is_normalized` |
| Smith-Waterman 系数 | `smithwaterman(text, text) returns float8` | `~=~` | 否 | `pg_similarity.sw_threshold`, `pg_similarity.sw_is_normalized` |
| Smith-Waterman-Gotoh 系数 | `smithwatermangotoh(text, text) returns float8` | `~!~` | 否 | `pg_similarity.swg_threshold`, `pg_similarity.swg_is_normalized` |
| Soundex 距离 | `soundex(text, text) returns float8` | `~*~` | 否 | |


--------

## 参数

多个参数控制 pg_similarity 函数和运算符的行为。可分为三类：

- **tokenizer**：控制字符串的分词方式。有效值为 **alnum**、**gram**、**word** 和 **camelcase**。所有 token 均为小写。默认为 **alnum**。
  - **alnum**：分隔符为任何非字母数字字符。
  - **gram**：n-gram 是使用逐一滑动技术提取的长度为 n 的子序列。
  - **word**：分隔符为空白字符。
  - **camelcase**：分隔符为大写字母，但大写字母也作为 token 的首字符。
- **threshold**：控制结果集的灵活程度。取值范围 **0.0** 到 **1.0**。默认为 **0.7**。
- **normalized**：控制相似度系数/距离是否归一化（在 0.0 和 1.0 之间）。默认为 **true**。


--------

## 示例

运行时设置参数：

```sql
SHOW pg_similarity.levenshtein_threshold;
-- 0.7

SET pg_similarity.levenshtein_threshold TO 0.5;

SET pg_similarity.cosine_tokenizer TO camelcase;

SET pg_similarity.euclidean_is_normalized TO false;
```

示例用表：

```sql
CREATE TABLE foo (a text);
INSERT INTO foo VALUES('Euler'),('Oiler'),('Euler Taveira de Oliveira'),('Maria Taveira dos Santos'),('Carlos Santos Silva');

CREATE TABLE bar (b text);
INSERT INTO bar VALUES('Euler T. de Oliveira'),('Euller'),('Oliveira, Euler Taveira'),('Sr. Oliveira');
```

### 使用相似度函数

```sql
SELECT a, b, cosine(a,b), jaro(a, b), euclidean(a, b) FROM foo, bar;
```

### 使用 levenshtein 运算符 (~==)

```sql
SHOW pg_similarity.levenshtein_threshold;
-- 0.7

SELECT a, b, lev(a,b) FROM foo, bar WHERE a ~== b;
--              a             |          b           |   lev
-- ---------------------------+----------------------+----------
--  Euler                     | Euller               | 0.833333
--  Euler Taveira de Oliveira | Euler T. de Oliveira |     0.76

SET pg_similarity.levenshtein_threshold TO 0.5;

SELECT a, b, lev(a,b) FROM foo, bar WHERE a ~== b;
--              a             |          b           |   lev
-- ---------------------------+----------------------+----------
--  Euler                     | Euller               | 0.833333
--  Oiler                     | Euller               |      0.5
--  Euler Taveira de Oliveira | Euler T. de Oliveira |     0.76
```

### 使用 qgram 运算符 (~~~)

```sql
SET pg_similarity.qgram_threshold TO 0.7;

SELECT a, b, qgram(a, b) FROM foo, bar WHERE a ~~~ b;
--              a             |            b            |  qgram
-- ---------------------------+-------------------------+----------
--  Euler                     | Euller                  |      0.8
--  Euler Taveira de Oliveira | Euler T. de Oliveira    |  0.77551
--  Euler Taveira de Oliveira | Oliveira, Euler Taveira | 0.807692
```

### 比较不同运算符

```sql
SELECT * FROM bar WHERE b ~@@ 'euler'; -- jaro-winkler 运算符
SELECT * FROM bar WHERE b ~~~ 'euler'; -- qgram 运算符
SELECT * FROM bar WHERE b ~== 'euler'; -- levenshtein 运算符
SELECT * FROM bar WHERE b ~## 'euler'; -- cosine 运算符
```
