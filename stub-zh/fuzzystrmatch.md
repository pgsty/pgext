

## 用法

> [fuzzystrmatch: 判断字符串之间的相似度和距离](https://www.postgresql.org/docs/current/fuzzystrmatch.html)

`fuzzystrmatch` 模块提供了判断字符串之间相似度和距离的函数。

```sql
CREATE EXTENSION fuzzystrmatch;
```

### Soundex

将字符串转换为 Soundex 编码（适用于匹配发音相似的名称）：

```sql
SELECT soundex('Anne'), soundex('Ann'), difference('Anne', 'Ann');
-- A500, A500, 4

SELECT soundex('Anne'), soundex('Andrew'), difference('Anne', 'Andrew');
-- A500, A536, 2

SELECT soundex('Anne'), soundex('Margaret'), difference('Anne', 'Margaret');
-- A500, M626, 0
```

`difference` 函数返回 0-4，其中 4 表示最相似。

### Daitch-Mokotoff Soundex

返回 Daitch-Mokotoff soundex 编码集合（对非英语名称效果更好）：

```sql
SELECT daitch_mokotoff('George');
-- {595000}

SELECT daitch_mokotoff('John');
-- {160000,460000}

-- 查找发音类似 'Schwartzenegger' 的名称
SELECT * FROM s WHERE daitch_mokotoff(nm) && daitch_mokotoff('Schwartzenegger');
```

支持 GIN 索引：

```sql
CREATE INDEX ON s USING gin (daitch_mokotoff(nm) gin__int_ops);
```

### Levenshtein 距离

计算两个字符串之间的编辑距离（插入、删除、替换）：

```sql
SELECT levenshtein('GUMBO', 'GAMBOL');
-- 2

SELECT levenshtein('GUMBO', 'GAMBOL', 2, 1, 1);
-- 3（自定义代价：插入=2，删除=1，替换=1）

-- 有界版本（更快，提前终止）
SELECT levenshtein_less_equal('extensive', 'exhaustive', 2);
-- 3（实际距离超过阈值，返回实际值）

SELECT levenshtein_less_equal('extensive', 'exhaustive', 4);
-- 4
```

### Metaphone

返回字符串的 metaphone 编码：

```sql
SELECT metaphone('GUMBO', 4);
-- KM
```

### Double Metaphone

返回主要和备选编码（处理更多名称变体）：

```sql
SELECT dmetaphone('gumbo');
-- KMP

SELECT dmetaphone_alt('gumbo');
-- KMP
```
