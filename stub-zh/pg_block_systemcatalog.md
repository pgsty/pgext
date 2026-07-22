## 用法

来源：

- [官方 README](https://github.com/nuko-yokohama/pg_block_systemcatalog/blob/0849fa12fd6049e422b37c2912ae56f10f6f1f22/README.md)
- [扩展控制文件](https://github.com/nuko-yokohama/pg_block_systemcatalog/blob/0849fa12fd6049e422b37c2912ae56f10f6f1f22/pg_block_systemcatalog.control)

`pg_block_systemcatalog` 会拒绝普通用户引用 PostgreSQL 系统目录、统计视图或 `information_schema` 的查询。超级用户不受限制，并可配置一个角色作为监控账号的允许列表角色。

### 核心流程

在 `postgresql.conf` 中预加载库，并可选指定一个组角色，然后重启 PostgreSQL 并创建扩展。

```conf
shared_preload_libraries = 'pg_block_systemcatalog'
pg_block_systemcatalog.allow_role = 'monitor_catalog'
```

```sql
CREATE EXTENSION pg_block_systemcatalog;

CREATE ROLE monitor_catalog;
CREATE ROLE monitor LOGIN;
GRANT monitor_catalog, pg_monitor TO monitor;
```

`pg_block_systemcatalog.allow_role` 接受一个角色。把该角色授予每个需要目录访问权的账号；如果它为空，则仅允许特权用户访问。当需要保持职责分离时，还应另外撤销监控账号对应用对象的权限。

### 安全边界

这是查询引用防护，并非完整的元数据沙箱。上游演示表明，即使在 `FROM` 中选择 C 统计函数 `pg_stat_get_activity()` 会被阻止，标量调用仍可能返回信息。其他 C 函数也可能在内部读取目录，而不把目录关系暴露到查询树中。

文档仅验证了 CentOS 7 上的 PostgreSQL 10。该模块检查计划或解析后的查询结构，可能影响管理工具、驱动、ORM、监控、迁移以及查询目录的扩展代码。应在准确的服务端大版本上测试每个必需流程和每个使用钩子的扩展；允许列表角色成员资格不能替代 PostgreSQL 的正常权限控制。
