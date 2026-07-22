## 用法

来源：

- [官方 README](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/README.md)
- [官方控制文件](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/plmruby.control)
- [官方语言定义](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/plmruby--0.0.1.sql)
- [官方标量示例](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/sql/plmruby.sql)
- [官方运行时构建配置](https://github.com/choplin/plmruby/blob/5fd693cfa3fd44eb1f1178bbd94d184c399fd4a6/build_config.rb)

`plmruby` 将 mruby 嵌入为 PostgreSQL 过程语言，可编写标量、集合返回、内联与触发器函数。已复核的 `0.0.1` 源码树是 2015 年未完成的实现；README 中的快速入门与集合返回文档仍未写完。

### 核心流程

只有超级用户能安装该扩展。具名 PostgreSQL 参数在 mruby 代码中使用相同名称。

```sql
CREATE EXTENSION plmruby;
REVOKE USAGE ON LANGUAGE plmruby FROM PUBLIC;

CREATE FUNCTION ruby_add(a integer, b integer)
RETURNS integer
LANGUAGE plmruby
IMMUTABLE STRICT
AS $$
  a + b
$$;

SELECT ruby_add(2, 3);
```

未具名参数依次使用 `_1`、`_2` 等名称。可变参数可以按普通具名参数使用。`DO` 块可以使用该语言的内联处理器。

### 数据与触发器模型

该实现会把 PostgreSQL 数值、布尔值、字符串、日期与时间戳、JSON、数组及记录转换成对应的 mruby 值。不受支持的类型通过 PostgreSQL 文本输入/输出函数传递，可能丢失类型特有语义。

行触发器代码会收到 `new`、`old`、`tg_name`、`tg_when`、`tg_level`、`tg_op`、`tg_relid`、`tg_table_name`、`tg_table_schema` 与 `tg_argv`。对于行级前置触发器，返回 `:skip` 会取消该行，返回类似元组的哈希会替换该行；语句级与后置触发器会忽略返回值。

### 信任与兼容性边界

安装 SQL 将 `plmruby` 声明为受信任语言，但嵌入式运行时构建包含文件 I/O、目录、环境、进程和系统信息 mruby gem。应把它视作能够在数据库后端内执行主机级操作。撤销公开语言使用权，只允许在操作系统边界上可信的角色创建函数。

该语言没有验证器，上游源码也早于当前仍受支持的 PostgreSQL 版本。必须在目标服务器构建上测试编译、类型转换、错误、触发器与集合返回函数。不要假设嵌入式 mruby 运行时提供现代 Ruby 库或 CRuby 行为。
