## 用法

来源：

- [已复核 commit 的 pg_scws README](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/README.md)
- [已复核 commit 的 pg_scws.control](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/pg_scws.control)
- [版本 1.0 的安装 SQL](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/pg_scws--1.0.sql)
- [解析器与配置实现](https://github.com/jaiminpan/pg_scws/blob/338d1ebe911372165acad331264bbf48afa56a9e/pg_scws.c)
- [内置 SCWS 源码与词典](https://github.com/jaiminpan/pg_scws/tree/338d1ebe911372165acad331264bbf48afa56a9e/libscws)

`pg_scws` 把 SCWS 中文分词器嵌入为 PostgreSQL 的 `scws` 文本搜索解析器。安装会创建 `scwscfg` 文本搜索配置，并通过 `simple` 词典映射名词、动词、形容词、习语、感叹词和临时词等 token 类型。

### 中文全文解析

```sql
CREATE EXTENSION pg_scws;

SELECT to_tsvector(
  'scwscfg',
  '小明硕士毕业于中国科学院计算所，后在日本京都大学深造'
);
```

会话设置包括 `scws.charset`、`scws.rules`、`scws.extra_dicts`、`scws.punctuation_ignore`、`scws.seg_with_duality` 和 `scws.multi_mode`。规则和词典名称会在 PostgreSQL 共享 `tsearch_data` 目录下解析。

### 注意事项

- 版本 `1.0` 上游只在 PostgreSQL 9.4 上测试过，已复核源码则停留在 2016 年。必须针对确切目标版本移植并回归测试解析器。
- 内置默认词典和规则决定分词质量与词汇范围。修改这些文件属于服务器级部署变更，必须在副本间保持一致。
- `scws.extra_dicts` 接受服务器端词典文件名，而不是客户端文件。上游用户自定义词典文档为空；应自行建立受控的构建、分发与回滚流程。
