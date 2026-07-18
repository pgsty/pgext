## 用法

来源：

- [上游 README](https://github.com/papercompute/pg_acce/blob/4ba51aa93e7d4d568fbf0ec524051f176fd03c80/README.md)
- [1.0 版安装 SQL](https://github.com/papercompute/pg_acce/blob/4ba51aa93e7d4d568fbf0ec524051f176fd03c80/pg_acce--1.0.sql)
- [后台工作进程实现](https://github.com/papercompute/pg_acce/blob/4ba51aa93e7d4d568fbf0ec524051f176fd03c80/pg_acce.c)
- [内部 API 声明](https://github.com/papercompute/pg_acce/blob/4ba51aa93e7d4d568fbf0ec524051f176fd03c80/pg_acce.h)

pg_acce 1.0 是一个未完成的 2015 年后台工作进程实验。尽管描述为 Accelerated Engine，已发布 SQL 只公开一个日志函数，以及一个启动工作进程执行所提供 SQL 文本的函数。

### 锁定安装

不要调用 acce_setup。如果必须检查该扩展，应在安装事务中撤销其默认公共权限：

```sql
BEGIN;
CREATE EXTENSION pg_acce;
REVOKE EXECUTE ON FUNCTION acce_setup(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION acce_info() FROM PUBLIC;
COMMIT;
```

即使是隔离检查构建，也应将其视为高风险实验代码。

### 注意事项

- acce_setup 接受任意 SQL 文本、启动动态后台工作进程，而且默认允许 public 执行。工作进程会复制调用者名称，却在初始化数据库连接时忽略它，造成不安全且可能提升权限的语义。
- 工作进程以只读 SPI 模式执行所提供文本，但会把完整查询与每个返回字段写入服务器标准错误。敏感查询文本和数据可能泄露到日志。
- 结果处理路径假定至少存在一行一列，并无条件解引用元组表。空结果或不符合预期的结果可能使工作进程崩溃。
- 声明的工作进程上限没有执行。重复调用可能耗尽动态工作进程槽位、共享内存与服务器资源。
- 代码通过旧版 PostgreSQL 私有 API 手工处理事务、快照、资源所有者、动态共享内存与后台工作进程。没有实质测试或当前兼容证据。
- README 将项目称为模板，测试脚本调用了不存在的函数，仓库许可证也不完整。不得部署该代码或暴露 acce_setup。
