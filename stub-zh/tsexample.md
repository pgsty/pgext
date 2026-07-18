## 用法

来源：

- [扩展 control 文件](https://github.com/postgrespro/tsexample/blob/9635a72ba8232f0a4216a4e3e2e4832adfc171a2/tsexample.control)
- [1.0 版安装 SQL](https://github.com/postgrespro/tsexample/blob/9635a72ba8232f0a4216a4e3e2e4832adfc171a2/tsexample--1.0.sql)
- [C 语言解析器与词典实现](https://github.com/postgrespro/tsexample/blob/9635a72ba8232f0a4216a4e3e2e4832adfc171a2/tsexample.c)
- [回归示例 SQL](https://github.com/postgrespro/tsexample/blob/9635a72ba8232f0a4216a4e3e2e4832adfc171a2/sql/tsexample.sql)

`tsexample` 1.0 是一个演示 PostgreSQL 可扩展全文检索接口的精简 C 示例。它安装 `sample_parser` 解析器、`cutdict` 词典模板、`cut3` 词典和 `sample` 文本检索配置。

### 查看解析结果和词元

```sql
CREATE EXTENSION tsexample;

SELECT * FROM ts_parse('sample_parser', 'abc def 123 1xx yy3');
SELECT to_tsvector('sample', 'abcdef 12345678 xyz');
SELECT plainto_tsquery('sample', 'abcdef 12345678 xyz');
```

`sample_parser` 把字母、数字和下划线归入同一 token：全数字 token 分类为 `number`，含任意非数字字符的 token 分类为 `word`。`sample` 配置让数字使用 PostgreSQL 的 `simple` 词典，让单词使用 `cut3`。`cut3` 会把单词转换为小写；对较长输入，则生成由前三个和后三个字符组成的词元。

### 示例定位

该项目是教学代码，不是通用语言分析器。固定对象名和三字符词典参数都应视为可改造示例，而不是完整检索方案。仓库没有 README、明确的许可证文件、发布历史或兼容矩阵。应针对实际 PostgreSQL 大版本构建并测试 C 模块；若把这种模式纳入其他扩展，还应使用不同的对象名。
