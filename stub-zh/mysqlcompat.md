## 用法

来源：

- [PGXN 上的 mysqlcompat 0.0.7 README](https://pgxn.org/dist/mysqlcompat/0.0.7/README.html)
- [已复核 commit 的 mysqlcompat 构建规则](https://github.com/2ndQuadrant/mysqlcompat/blob/3b5319eed968ca32504d15cdd64ad0837b18929e/Makefile)
- [已复核 commit 的 mysqlcompat 汇总安装脚本](https://github.com/2ndQuadrant/mysqlcompat/blob/3b5319eed968ca32504d15cdd64ad0837b18929e/all.sql)

`mysqlcompat` 0.0.7 是一组纯 SQL/PLpgSQL 的 MySQL 风格函数、聚合、运算符和类型转换，用于减少移植工作。其 PGXS 构建会从 `sql_bits` 下的文件生成带版本号的扩展 SQL。

```sql
CREATE EXTENSION mysqlcompat;

SELECT true && false;
SELECT format(1234.432, 4);
SELECT ADDTIME('01:00:00.999999'::interval, '02:00:00.999998');
```

对象安装在扩展的目标模式中；上游示例假设当前搜索路径可以找到这些对象。

### 注意事项

- 上游只报告测试过 PostgreSQL 8.2 至 9.5。每个受支持的现代大版本都应验证对象创建、重载解析和行为。
- 这只是部分兼容，并不模拟 MySQL 语法。命名形式的 XOR、DIV、MOD 运算符未实现，MySQL 的区间语法也必须改写。
- PostgreSQL 字符串仍区分大小写，而且部分 MySQL 时间类结果使用 PostgreSQL interval 表示。
- 名称以下划线开头的函数是内部辅助函数，不应直接调用。
