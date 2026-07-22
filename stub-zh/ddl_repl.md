## 用法

来源：

- [官方 README](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/README.md)
- [扩展控制文件](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/ddl_repl.control)
- [扩展 SQL](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/ddl_repl--1.0.sql)
- [钩子实现](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/ddl_repl.c)

`ddl_repl` 是一个实验性的 ProcessUtility 钩子，会把工具语句转发到已配置的 PostgreSQL 节点。它试图让多个实例的 DDL 与 DCL 保持相似，但本质是同步重放 SQL 文本，不是逻辑复制，也不是原子的分布式事务。

### 核心流程

需要在每个可能发起重放的后端中加载动态库；实践中应进行服务器级预加载。修改 `shared_preload_libraries` 后重启，安装目录对象，再用 PostgreSQL 服务器可用的 DSN 注册单向目标。

```conf
shared_preload_libraries = 'ddl_repl'
```

```sql
CREATE EXTENSION ddl_repl;

SELECT ddl_repl.create_node(
  'reporting',
  'host=reporting-db dbname=app user=ddl_repl password=secret',
  'one-way DDL destination'
);

SELECT ddl_repl.set_node_active('reporting', false);
SELECT ddl_repl.drop_node('reporting');
```

`ddl_repl.nodes` 保存 `node_name`、`dsn`、`description`、`active` 与 `creation_date`。`ddl_repl.enabled` 控制当前后端是否重放；源码中还提供默认关闭的 `ddl_repl.only_repl_users`。管理接口包括 `ddl_repl.create_node`、`ddl_repl.set_node_active` 与 `ddl_repl.drop_node`。

### 运行边界

钩子先在本地执行，再建立 libpq 连接、开启远端事务、重放原始查询文本，并逐个提交活动节点。后续连接或提交失败时，较早的节点可能已经提交，因此需要人工恢复且各节点可能分叉。不要配置互相指回的节点，否则已加载的远端钩子可能把同一语句再次重放回来。

过滤器只排除 `SET`、`SHOW`、`PREPARE`、`EXECUTE`、`DEALLOCATE` 等少量工具语句，并不是严格的 DDL 白名单。DSN 以普通目录文本保存，扩展 SQL 也没有收紧默认权限。应撤销公众对模式、表和管理函数的访问权，使用最小权限远端角色，妥善保护凭据，并针对每个目标 PostgreSQL 大版本单独测试。固定版本的上游源码停留在 2018 年并依赖旧服务器内部接口，因此只应把 `ddl_repl` 当作实验室原型，而不是生产一致性机制。
