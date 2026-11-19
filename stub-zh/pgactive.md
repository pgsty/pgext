## 用法

来源：

- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/README.md)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/docs/quickstart.md)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/docs/README.md)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/docs/referenceguide.md)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/pgactive.control.in)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/include/pgactive_version.h.in)
- [Official source](https://github.com/aws/pgactive/blob/v2.1.10/pgactive--2.1.10.sql)
- [Official source](https://github.com/aws/pgactive/releases/tag/v2.1.10)

`pgactive` 2.1.10 在数据库之间以全互联拓扑提供多主逻辑复制。写入先本地提交再复制到其他节点；应用必须处理延迟、冲突和节点本地锁，不能假定具备全局可串行化能力。

### 配置复制组

此版本要求 PostgreSQL 18.7、17.12、16.16、15.20 或 14.25，以及对应分支的更新次版本；发布说明还列出即将到来的 PostgreSQL 19。预加载库，配置逻辑 WAL 和足够的复制槽及工作进程，再重启。安装和管理函数需要超级用户权限。

```ini
shared_preload_libraries = 'pgactive'
wal_level = logical
```

```sql
CREATE EXTENSION pgactive;
SELECT pgactive.pgactive_create_group(
  node_name := 'node1-app',
  node_dsn := 'user_mapping=postgres pgactive_foreign_server=node1'
);
SELECT pgactive.pgactive_wait_for_node_ready();
```

示例假定已经存在使用 pgactive_fdw 的 node1 外部服务器、受保护的用户映射，以及连接自身数据库和所有其他节点的通路。实际凭据应通过受保护的映射或密码文件配置。对新加入的空数据库，先创建扩展，再用自身 node_dsn 和对端 join_using_dsn 调用 pgactive.pgactive_join_group，并等待就绪。内置加入流程负责协调初始化，独立制作的普通转储不能替代它。

### 运维与升级

加入节点时所有对端都必须可达。提前规划唯一键、序列分配、冲突处理和复制槽保留。DDL 复制默认关闭，启用时要求发送端和接收端采用一致策略。应用工作进程默认以表所有者身份执行 DML，但复制连接及工作进程管理仍具备高权限。

2.1.10 加固复制名称和元组处理，并修复冲突处理器 ABI 与共享内存锁偏移。2.1.10 节点可以加入旧版本组，但 2.1.9 及更早节点必须先升级，才能加入 2.1.10 节点。安装匹配文件，在每个数据库执行随附 SQL 升级，并协调组内的重启、重连和加入顺序。
