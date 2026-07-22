## 用法

来源：

- [官方 README](https://github.com/obartunov/dict_roman/blob/22c2cd1acd8eb077a6d76e543e4ae6a4a0071d61/README.dict_roman)
- [官方扩展控制文件](https://github.com/obartunov/dict_roman/blob/22c2cd1acd8eb077a6d76e543e4ae6a4a0071d61/dict_roman.control)

`dict_roman` 提供 `roman` 全文搜索词典，可把罗马数字词元规范化为阿拉伯整数的文本形式。当 `XXIV` 这类数字应当以 `24` 而不是普通字母文本进入搜索向量时，可以使用它。

### 核心流程

安装扩展后，先用 `ts_lexize` 直接检查词典：

```sql
CREATE EXTENSION dict_roman;

SELECT ts_lexize('roman', 'XXIV');  -- {24}
SELECT ts_lexize('roman', 'XLII');  -- {42}
```

若要在建立索引和查询时应用这种规范化，请在文本搜索配置中映射合适的词元类型：

```sql
CREATE TEXT SEARCH CONFIGURATION english_with_roman (COPY = english);
ALTER TEXT SEARCH CONFIGURATION english_with_roman
  ALTER MAPPING FOR asciiword
  WITH roman, english_stem;

SELECT to_tsvector('english_with_roman', 'Chapter XXIV');
```

映射中的后备词典很重要，它会处理 `roman` 无法识别的普通字母词元。

### 运维说明

- 已公开的 SQL 接口是 `roman` 词典；上游 README 演示了 `ts_lexize`，但没有给出完整语法或可接受数值范围。
- 在依赖规范化结果前，应针对实际打包版本测试零、减法记数法、大小写、非法输入和大数值。
- 修改文本搜索配置只影响之后的解析。若现有行也要采用新映射，需要重建已存储的 `tsvector` 值及其相关索引。
