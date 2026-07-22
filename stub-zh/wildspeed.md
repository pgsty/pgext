## 用法

来源：

- [官方扩展控制文件](https://github.com/postgrespro/wildspeed/blob/a1b4e02f56e81596993848cd6692b0cd5b2e38d5/wildspeed.control)
- [官方 README](https://github.com/postgrespro/wildspeed/blob/a1b4e02f56e81596993848cd6692b0cd5b2e38d5/README.md)
- [官方扩展 SQL](https://github.com/postgrespro/wildspeed/blob/a1b4e02f56e81596993848cd6692b0cd5b2e38d5/wildspeed--1.0.sql)

`wildspeed` 1.0 提供 GIN 运算符类 `wildcard_ops`，用于加速带前导或内部通配符的 `LIKE` 搜索。它使用 permuterm 风格索引，使 `%suffix` 和 `%middle%` 等模式可以转化为可索引的前缀搜索；它不支持 `ILIKE` 或正则表达式运算符。

### 核心流程

使用 `wildcard_ops` 创建 GIN 索引，然后在被索引的文本表达式上编写普通 `LIKE` 条件。

```sql
CREATE EXTENSION wildspeed;

CREATE TABLE words (
    word text NOT NULL
);

CREATE INDEX words_wildcard_idx
    ON words USING gin (word wildcard_ops);

ANALYZE words;

SELECT word
FROM words
WHERE word LIKE '%graph%';
```

应使用 `EXPLAIN (ANALYZE, BUFFERS)` 与有代表性的数据和模式，确认规划器确实选择该索引，并且总成本合算。

### 对象与模式行为

- `wildcard_ops` 是面向 `text` 和 PostgreSQL `LIKE` 运算符 `~~` 的 GIN 运算符类。
- `permute(text)` 用于测试时查看生成的旋转项；应用通常无需调用它。
- 内部支持函数负责提取 permuterm 键，并把通配项转换成 GIN partial match。

Permuterm 索引会围绕结束标记旋转每个被索引单词，因此前导和内部通配符可能远快于顺序扫描。普通前缀模式仍可能更适合 B-tree `text_pattern_ops` 索引，尤其是短前缀。

### 空间、写入与兼容性边界

每个旋转项都会成为索引条目，因此 `wildcard_ops` 索引可能比表或相当的 B-tree 索引大很多倍。官方示例展示了显著的构建时间与空间膨胀；应把它视为警示，而不是性能保证。需要为 GIN 构建、插入/更新、vacuum、备份和复制成本预留资源。

像 `%` 这样几乎匹配全部数据的模式可能需要完整扫描 GIN 索引，几乎没有选择性收益。必须在准确部署环境中测试短词、重复通配符、非 ASCII 数据、数据库编码和排序规则。不要假定 `ILIKE`、正则或区域相关的大小写折叠会使用该运算符类。

该项目在 2016 年被现代化为扩展，近期没有兼容性维护。采用前应针对目标 PostgreSQL 版本构建并运行回归测试，同时与当前内置索引或应用搜索方案比较。
