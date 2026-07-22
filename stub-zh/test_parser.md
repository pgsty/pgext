## 用法

来源：

- [官方仓库](https://github.com/r888888888/test_parser)
- [目录源码版本的扩展控制文件](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser.control)
- [1.0 版扩展 SQL](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser--1.0.sql)
- [1.0 版 C 实现](https://github.com/r888888888/test_parser/blob/356c676d16a987f9226b3ea419b4e18612c61fef/test_parser.c)
- [PostgreSQL 文本搜索解析器测试文档](https://www.postgresql.org/docs/current/textsearch-debugging.html)

test_parser 是一个极简的 PostgreSQL 自定义全文搜索解析器示例。扩展名包含下划线，而安装的解析器名为 testparser；其实现只把连续的字面 ASCII 空格与所有其他字节序列分开。

### 检查解析器

先在开发数据库中安装扩展，检查两个标记类型和生成的标记流，再创建文本搜索配置。

```sql
CREATE EXTENSION test_parser;

SELECT * FROM ts_token_type('testparser');
SELECT * FROM ts_parse('testparser', 'alpha  beta,123');
```

非空格序列会报告为单词标记类型，连续空格会报告为空白标记类型。标点、数字、制表符、换行、Unicode 空白和语言学词界都不会得到特殊分类。

### 最小配置

把单词标记映射到词典；空白标记通常保持未映射，但解析器仍会保留它们，使内置标题回调可以维持间距。

```sql
CREATE TEXT SEARCH CONFIGURATION testcfg (PARSER = testparser);
ALTER TEXT SEARCH CONFIGURATION testcfg
    ADD MAPPING FOR word WITH simple;

SELECT to_tsvector('testcfg', 'alpha beta alpha');
```

### 预期范围与兼容性

这段代码是教学解析器，而不是生产分词器。它只有两种标记类型，没有配置选项，也不进行语言感知分词。可用它学习解析器回调接口，或作为单独维护实现的起点，但不能据此声称具有实用搜索语义。

C 代码使用 PostgreSQL 的解析器回调 ABI，并复用内置标题函数，因此必须针对准确的 PostgreSQL 主版本构建和测试。验收测试应覆盖多字节文本、嵌入空字节假设、长输入、所有空白形式、标记偏移、配置映射和标题输出。该扩展可重定位，也没有声明预加载或重启要求。
