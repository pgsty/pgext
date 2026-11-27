## 用法

来源：

- [Official source](https://github.com/RekGRpth/gp/blob/a52d096479a185f1ad362c1588e80efa466a5583/contrib/gp/README.md)
- [Official source](https://github.com/RekGRpth/gp/blob/a52d096479a185f1ad362c1588e80efa466a5583/contrib/gp/gp.control)
- [Official source](https://github.com/RekGRpth/gp/blob/a52d096479a185f1ad362c1588e80efa466a5583/contrib/gp/gp--1.0.sql)
- [Official source](https://github.com/RekGRpth/gp/blob/a52d096479a185f1ad362c1588e80efa466a5583/contrib/gp/gp.c)
- [Official source](https://github.com/RekGRpth/gp/blob/a52d096479a185f1ad362c1588e80efa466a5583/contrib/gp/LICENSE)

`gp` 1.0 是与项目修改版 PostgreSQL 20 开发源码配套的实验性协调节点/数据节点扩展。它要求 `new_oid_hook` 补丁，已发布的原版服务器不能直接替代。仓库明确标为开发中，不适合生产使用。

### 建立测试集群

所有节点使用相同的修改版服务器和扩展二进制，在各节点预加载库并重启。协调节点的最小配置如下：

```ini
shared_preload_libraries = 'gp'
gp.role = 'coordinator'
gp.database = 'postgres'
```

```sql
CREATE EXTENSION gp;
SELECT * FROM gp.cluster_segments();
```

数据节点使用 `gp.role = 'segment'`、正整数且唯一的 `gp.segment_id`，以及可访问的 `gp.coordinator` 连接串。数据节点的 `max_prepared_transactions` 必须大于零。注册工作进程创建本地 SQL 对象，并在协调节点指定的数据库中注册。安装需要超级用户，固定模式为 `gp`，SQL 使用 `plpgsql`。

### 范围与权限

协调节点保存数据节点元数据，并维护共享内存副本。DDL 传播会保持各节点目录 OID 一致，因此参与节点的扩展版本和角色必须匹配。`gp.cluster_segments()` 用于查看已注册节点。管理函数 `gp.exec_on_all()` 可在数据节点执行命令；SQL 契约明确提醒，协调节点回滚通常不会撤销远端独立提交的命令。

分布式表放置、查询分派、预备事务和恢复都是当前源码中的实验功能。README 开头的摘要落后于详细路线图和 SQL 接口，必须验证具体修订与工作负载，不能假定完全兼容 Greenplum。已核对文档没有证明镜像副本及自动重分布能力。

### 连接与恢复边界

节点间会话需明确配置认证和访问限制。协调节点元数据、数据节点数据、预备事务状态、互联秘密及配置必须配套保留。公共模式访问权不等于管理函数执行权，安装 SQL 从 PUBLIC 撤销了特权控制函数权限。当前修订没有版本化 SQL 升级链；修改内核补丁或模块需要协调迁移和重启集群。
