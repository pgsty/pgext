## 用法

来源：

- [上游 README](https://github.com/jasonmp85/pg_fiu/blob/f2e0afc6963e49f6fac31d4bb1d169f8a4eaf312/README.md)
- [SQL 故障点 API](https://github.com/jasonmp85/pg_fiu/blob/f2e0afc6963e49f6fac31d4bb1d169f8a4eaf312/sql/pg_fiu.sql)
- [钩子与 libfiu 实现](https://github.com/jasonmp85/pg_fiu/blob/f2e0afc6963e49f6fac31d4bb1d169f8a4eaf312/src/pg_fiu.c)
- [构建配置](https://github.com/jasonmp85/pg_fiu/blob/f2e0afc6963e49f6fac31d4bb1d169f8a4eaf312/Makefile)

pg_fiu 0.0.1 是一个未完成的 2018 年故障注入原型，基于 libfiu 以及 PostgreSQL 解析器、执行器、工具命令和事务钩子。它只能用于可丢弃测试集群。

### 隔离设置

预加载动态库并重启可丢弃服务器：

```conf
shared_preload_libraries = 'pg_fiu'
```

安装到专用模式，并在检查前移除公共函数权限：

```sql
CREATE SCHEMA fiu_lab;
CREATE EXTENSION pg_fiu WITH SCHEMA fiu_lab;
REVOKE EXECUTE ON ALL FUNCTIONS IN SCHEMA fiu_lab FROM PUBLIC;
SELECT * FROM fiu_lab.failure_points();
```

在审查精确链接二进制与恢复流程前，不要启用任何注入点。

### 注意事项

- 故障注入会故意更改后端内部的 C 返回值与辅助指针。错误的名称、类型、回调、栈目标或返回值可能破坏内存、使进程崩溃或损坏数据。
- 钩子只有通过 shared_preload_libraries 加载才会激活，并影响每个新后端。下一事务开关仅限超级用户，但创建定义的 SQL 函数若未撤权则向 public 开放。
- 修改后端全局故障状态的 SQL 函数被错误声明为不可变。定义会持续到后端生命周期结束，而且没有 SQL 删除或清空函数。
- 目标代码必须包含匹配的 libfiu 插桩或兼容栈符号。上游没有可用使用示例；其回归文件调用了不存在的函数。
- Makefile 与一个平台分支包含开发者绝对路径。2018 年的钩子签名和后端内部接口需要为当前 PostgreSQL 主版本移植。
- 绝不能在生产环境、共享测试服务或包含不可替代数据的集群上预加载 pg_fiu。
