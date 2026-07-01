


## 用法

来源：[README](https://github.com/credativ/db2fce/blob/master/README.md)，[SQL objects](https://github.com/credativ/db2fce/blob/master/db2fce.sql)，[control file](https://github.com/credativ/db2fce/blob/master/db2fce.control)

`db2fce` 为 PostgreSQL 提供 DB2 兼容环境。它创建 DB2 风格的函数、类型、操作符，以及 `SYSIBM.SYSDUMMY1` 兼容视图，便于迁移或适配原本面向 IBM Db2 的 SQL。

### 启用

```sql
CREATE EXTENSION db2fce;

SET search_path = db2, sysibm, public;
```

大多数兼容对象位于 `db2` schema，`sysibm.sysdummy1` 用于兼容需要 dummy 单行表的 DB2 查询。

```sql
SELECT * FROM sysibm.sysdummy1;
```

### 兼容函数

`db2` schema 包含日期时间函数，例如 `microsecond`、`second`、`minute`、`hour`、`day`、`month`、`year`、`days`、`months_between`、`date`、`time`、`timestamp_format`。

字符串和转换函数包括 `locate`、`translate`、`lcase`、`upper`、`lower`、`strip`、`char`、`integer`、`int`、`double`、`decimal`、`dec`、`hex`、`round`、`digits`、`value`。

### 操作符

SQL 层还定义了 DB2 风格操作符，例如用于不等比较的 `^=`，以及多种类型上的拼接操作符 `!!`。

```sql
SELECT db2.int('42');
SELECT db2.days(current_date);
SELECT 'db' !! '2';
```

### 注意事项

把 `db2` 加入 `search_path` 后，很多 DB2 函数调用可以不带 schema 前缀。部分与 PostgreSQL 语法或内置行为冲突的名称，仍然需要显式写成 `db2.` 前缀。
