## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/postgrespro/snowball_ext/blob/4ed178a9e11ec7c1fbd20e4c708e162ef6b5d4c0/README.md)
- [1.0 版本 SQL 对象](https://github.com/postgrespro/snowball_ext/blob/4ed178a9e11ec7c1fbd20e4c708e162ef6b5d4c0/snowball_ext--1.0.sql)
- [尼泊尔语词干源码](https://github.com/postgrespro/snowball_ext/blob/4ed178a9e11ec7c1fbd20e4c708e162ef6b5d4c0/libstemmer/stem_UTF_8_nepali.c)
- [扩展 control 文件](https://github.com/postgrespro/snowball_ext/blob/4ed178a9e11ec7c1fbd20e4c708e162ef6b5d4c0/snowball_ext.control)

`snowball_ext` 复制 PostgreSQL 不可扩展的 Snowball 字典模板，从而注册额外词干算法。1.0 提供尼泊尔语词干器和 `nepali` 文本搜索配置。

```sql
CREATE EXTENSION snowball_ext;
SELECT to_tsvector('nepali', 'अँगअँकाउछन्');
```

其用法与内置文本搜索配置相同；构建生产索引前，应使用经审查的尼泊尔语语料测试词干处理。分词与词干质量会直接影响召回率、精确率与索引内容。

目录将该扩展标记为 abandoned。它复制 PostgreSQL 内部 Snowball 模板代码，且没有当前兼容矩阵，因此较新的后端变化可能破坏原生构建。文本搜索索引输出依赖算法，不能在备库或升级过程中静默更换版本；只有在比较结果并批准语言学变化后，才可重建受影响索引。
