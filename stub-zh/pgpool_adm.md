## 用法

来源：

- [Pgpool-II 4.7.3 control](https://github.com/pgpool/pgpool2/blob/V4_7_3/src/sql/pgpool_adm/pgpool_adm.control)
- [SQL API 1.6](https://github.com/pgpool/pgpool2/blob/V4_7_3/src/sql/pgpool_adm/pgpool_adm--1.6.sql)
- [PCP node reference](https://github.com/pgpool/pgpool2/blob/V4_7_3/doc/src/sgml/ref/pgpool_adm_pcp_node_info.sgml)
- [Foreign-server setup](https://github.com/pgpool/pgpool2/blob/V4_7_3/src/sql/pgpool_adm/pgpool_adm.c)

`pgpool_adm` 提供 Pgpool-II PCP 管理命令的 SQL 包装函数。它需要可访问的 Pgpool-II PCP 服务和适当凭证；安装扩展不会部署 Pgpool-II。

### 安装和检查节点

```sql
CREATE EXTENSION pgpool_adm;

SELECT * FROM pcp_node_info(
  node_id => 0, host => 'localhost', port => 9898,
  username => 'pcp_admin', password => 'example-password');
SELECT pcp_node_count(
  host => 'localhost', port => 9898,
  username => 'pcp_admin', password => 'example-password');
```

实际 SQL 名称是 `pcp_node_info`、`pcp_health_check_stats`、`pcp_pool_status`、`pcp_node_count`、`pcp_attach_node`、`pcp_detach_node`、`pcp_proc_info`。文档页面名称可能带有安装包前缀；这些 SQL 函数名称并不带该前缀。

### 凭证和服务器引用

直接调用重载接收主机、端口、用户名和密码，需要节点 ID 的函数将其放在首位。其他重载通过 `pcp_server` 参数引用外部服务器。应按源码实现配置服务器及用户映射；不要假设数据库后端会从客户端的 `.pcppass` 文件获取凭证。

```sql
SELECT * FROM pcp_node_info(node_id => 0, pcp_server => 'pgpool_server');
SELECT pcp_node_count(pcp_server => 'pgpool_server');
```

常用 PCP 端口是 9898。凭证处理和网络访问发生在 PostgreSQL 服务器上。密码字面量可能进入 SQL 日志和活动视图；适当时应优先使用已配置的服务器引用。

### 管理节点

```sql
SELECT pcp_detach_node(
  node_id => 1, gracefully => true,
  host => 'localhost', port => 9898,
  username => 'pcp_admin', password => 'example-password');
SELECT pcp_attach_node(
  node_id => 1, host => 'localhost', port => 9898,
  username => 'pcp_admin', password => 'example-password');
```

摘除或接入节点会改变 Pgpool-II 的路由，可能影响应用流量。应将函数访问和 PCP 凭证限制给管理员。数据库事务回滚不会撤销已经执行的外部 PCP 操作。

### 版本和加载

由超级用户安装扩展，无需预加载或重启 PostgreSQL。Pgpool-II 安装包版本 4.7.3 携带 SQL 扩展版本 1.6；安装对应文件后，在已有数据库中执行 `ALTER EXTENSION pgpool_adm UPDATE`，不要把 4.7.3 当作 SQL 版本。
