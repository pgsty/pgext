## 用法

来源：

- [mysqlcompat 0.0.7 README on PGXN](https://pgxn.org/dist/mysqlcompat/0.0.7/README.html)
- [Official upstream README](https://github.com/2ndQuadrant/mysqlcompat/blob/3b5319eed968ca32504d15cdd64ad0837b18929e/README.md)
- [Version 0.0.7 build manifest](https://github.com/2ndQuadrant/mysqlcompat/blob/3b5319eed968ca32504d15cdd64ad0837b18929e/Makefile)
- [Aggregate SQL entry point](https://github.com/2ndQuadrant/mysqlcompat/blob/3b5319eed968ca32504d15cdd64ad0837b18929e/all.sql)

`mysqlcompat` 0.0.7 是一组纯 SQL/PLpgSQL 的 MySQL 风格函数、聚合、操作符和类型转换，用于减少应用移植工作。它只提供部分兼容辅助对象，不会改变 PostgreSQL 语法，也不会让 PostgreSQL 完全等同于 MySQL。

### 核心用法

```sql
CREATE EXTENSION mysqlcompat;

SELECT true && false;
SELECT format(1234.432, 4);
SELECT ADDTIME(
  '01:00:00.999999'::interval,
  '02:00:00.999998'::interval
);
```

PGXS 构建会汇总 `sql_bits` 下的主题文件。主要分组包括 aggregate、bit、cast、control-flow、date/time、information、mathematical、miscellaneous、operator 和 string。以 `_` 开头的名称是内部辅助函数，不应直接调用。

### 兼容性边界

- 上游仅报告测试过 PostgreSQL 8.2 至 9.5。每个现代目标大版本都必须验证安装、重载解析、类型转换和结果。
- MySQL 命名操作符 `XOR`、`DIV`、`MOD` 无法通过普通扩展重现语法；上游建议分别使用 PostgreSQL `#`、整数除法和 `%` 等替代。
- PostgreSQL 字符串默认仍区分大小写，不同于许多 MySQL 排序规则。
- 一些 MySQL 时间类值会以 PostgreSQL `interval` 表示，MySQL interval 语法也必须改写为带引号的 PostgreSQL interval 字面量。
- 重载函数可能需要显式类型转换，如 `ADDTIME` 示例所示。

该扩展会在目标模式中创建大量通用名称对象。应尽可能安装到隔离模式、控制 `search_path`，并在暴露前盘点与核心或应用函数的冲突。移植代码仍需针对 NULL、溢出、排序规则、日期运算、聚合行为和隐式类型转换进行语义测试；对象创建成功并不能证明 MySQL 兼容性。
