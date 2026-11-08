## 用法

来源：

- [README.md](https://github.com/obartunov/pg_agent/blob/6acd7cdce9ccb9a2a560492cc44394bf430f61f6/README.md)
- [pg_agent.control](https://github.com/obartunov/pg_agent/blob/6acd7cdce9ccb9a2a560492cc44394bf430f61f6/pg_agent.control)
- [pg_agent--1.0.sql](https://github.com/obartunov/pg_agent/blob/6acd7cdce9ccb9a2a560492cc44394bf430f61f6/pg_agent--1.0.sql)
- [Makefile](https://github.com/obartunov/pg_agent/blob/6acd7cdce9ccb9a2a560492cc44394bf430f61f6/Makefile)

`pg_agent` 是纯 SQL 原型。1.0 版本实现了感知权限的目录摘要，并创建策略、标注、缓存和日志表。语句检查、执行策略强制约束和审计钩子仍属于规划功能。

### 核心用法

安装扩展后，以需要检查可见性的角色查询目录：

```sql
CREATE EXTENSION pg_agent;
SELECT pg_agent.catalog_json(ARRAY['public']::name[]);
SELECT pg_agent.check_statement('SELECT 1');
```

最后一次调用返回 `ok=false` 和 `reason=not_implemented`，并不验证语句。`relation_json` 只返回固定状态及关系名。不能把这些函数当作 SQL 沙箱或授权判断。

### 目录与权限

`catalog_json` 以调用者身份运行，通过 PostgreSQL 权限检查过滤模式、关系和列。可选参数控制注释、索引、外键和估算行数的输出。目录注释作为数据返回，并标记为不可作为可信指令。关系可见并不表示行级安全策略允许读取其中所有行。

固定的 `pg_agent` 模式包含 `object_annotation`、`role_policy`、`statement_log`、`cache_state` 和 `catalog_cache`。安装脚本撤销这些表的公共访问权限，管理员需要显式授予必要的模式和对象权限。仅写入策略行不会使策略生效。

### 运维边界

控制文件将扩展设为不可迁移，并设置 `superuser=false`；安装者仍需具备创建模式和对象的权限。纯 SQL 构建没有共享库，也不要求预加载或重启。上游未公布 PostgreSQL 主版本兼容矩阵。应继续使用常规角色授权、超时和只读事务控制，不能用此原型替代它们。
