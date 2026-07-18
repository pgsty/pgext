## 用法

来源：

- [已复核 commit 的扩展 control 文件](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest.control)
- [1.0 版本 SQL 定义](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest--1.0.sql)
- [词法分析器实现](https://github.com/RhodiumToad/pg_lextest/blob/11325ee8d16459caa13f57bca1042ede61f256d4/lextest.c)

`lextest` 是调用 PostgreSQL 核心 SQL 词法分析器的轻量开发与基准测试扩展。`lexer(cstring)` 返回每个标记的 `pos`、数字 `token` 和已解码 `value`；`lexer_test(cstring)` 扫描输入但不返回标记流。

```sql
CREATE EXTENSION lextest;

SELECT *
FROM lexer('SELECT 1 + 2');

SELECT lexer_test('SELECT 1 + 2');
```

### 注意事项

这是解析器开发辅助工具，不是应用层文本分析 API。它包含 PostgreSQL 内部解析器头文件并依赖私有 token 定义，因此换用其他大版本时可能需要修改源码。上游没有发布 README、发行标签、兼容性矩阵或许可证文件。
