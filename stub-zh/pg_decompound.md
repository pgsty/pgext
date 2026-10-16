## 用法

来源：

- [README.md](https://github.com/fgehrlicher/pg_decompound/blob/07487d222cd730985f289440a8032f31137e19b5/README.md)
- [sql/pg_decompound--1.0.sql](https://github.com/fgehrlicher/pg_decompound/blob/07487d222cd730985f289440a8032f31137e19b5/sql/pg_decompound--1.0.sql)
- [flake.nix](https://github.com/fgehrlicher/pg_decompound/blob/07487d222cd730985f289440a8032f31137e19b5/flake.nix)
- [pg_decompound.control](https://github.com/fgehrlicher/pg_decompound/blob/07487d222cd730985f289440a8032f31137e19b5/pg_decompound.control)

`pg_decompound` 提供 `decompound` 全文检索词典模板，在建立索引和查询时将复合词拆为词表中的组成词。该项目针对德语开发，实际处理的语言与输出词元取决于提供的词表。

### 核心用法

```sql
CREATE EXTENSION pg_decompound;
CREATE TEXT SEARCH DICTIONARY german_compounds (
    TEMPLATE = decompound,
    WordList = german_compounds,
    MinPiece = 4,
    MinWord = 8,
    MinCoverage = 70
);
CREATE TEXT SEARCH CONFIGURATION german_dc (COPY = german);
ALTER TEXT SEARCH CONFIGURATION german_dc
    ALTER MAPPING FOR asciiword, word, hword_part, hword_asciipart
    WITH german_compounds, german_stem;
SELECT to_tsvector('german_dc', 'Konfigurationsdatei');
```

### 运行边界

此原生扩展需要管理员安装。上游开发环境以 PostgreSQL 16 和 17 为目标。创建词典前需准备词表：`WordList` 指定 `$SHAREDIR/tsearch_data` 下的 `.dict` 文件。每行包含一个表面词，也可在制表符后附加已提取词干的词元。项目不附带词表；词表的许可与扩展本身的 MIT 许可分开。

复合词词典应排在词干词典之前。不能拆分的词会交给后续词典，但已输出的片段不会继续提取词干，因此词表中的词元必须预先处理。`tools/build-wordlist.sh` 可调用现有词干词典生成词元。

`MinPiece` 默认为 4 个字符，`MinWord` 为 8，`MaxWord` 为 100。`MinCoverage` 默认为 0，会输出所有子串匹配；设为正百分比后，会选择不重叠的覆盖方案并拒绝覆盖不足的结果。`OnlyLongest` 在每个偏移位置保留最长匹配，`KeepOriginal` 额外保留未经词干处理的原复合词，两项默认均为 false。

大型词表容易产生无关的子串匹配，`MaxWord` 用于限制平方级匹配开销。查询词产生的每个词元都必须匹配，因此词形变化导致的不同拆分可能漏查，启用 `KeepOriginal` 时尤其需要注意。修改词典行为后，应重新计算已存储的搜索向量。
