## 用法

来源：

- [已复核 commit 的扩展 control 文件](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest.control)
- [1.0 版本 SQL 对象](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest--1.0.sql)
- [词法器实现](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest.c)

`lextest` 暴露 PostgreSQL 内部的核心 SQL 扫描器，用于解析器开发与诊断。它不是通用的文本搜索分词器：数字 token 编号和载荷规则直接来自 PostgreSQL 语法内部实现，可能随大版本变化。

### 核心流程

使用 `lexer(cstring)` 检查 token 流，或用 `lexer_test(cstring)` 运行扫描器并丢弃结果。

```sql
CREATE EXTENSION lextest;

SELECT *
FROM lexer('SELECT 1 + $2 FROM public.demo');

SELECT lexer_test('UPDATE demo SET value = 42 WHERE id = 7');
```

`lexer(cstring)` 返回 `pos`、`token` 和 `value`。`pos` 是扫描器输入中的字节位置，不是字符下标。`token` 是 PostgreSQL 内部的数字语法 token。只有当此包装器会提取字符串或数字载荷时，`value` 才有值；其他情况为 `NULL`。函数会先物化所有行再返回。

### 兼容性边界

C 代码调用不稳定的后端扫描器函数，并包含解析器内部头文件。每个确切的 PostgreSQL 大版本和编译环境都必须单独构建；某个服务器上构建成功并不能证明兼容其他版本。数字 token 值属于实现细节，不应存成持久的应用协议。`lextest` 应仅限扩展开发或解析器调试环境。
