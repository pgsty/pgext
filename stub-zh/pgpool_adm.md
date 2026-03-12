

## 用法

> [pgpool_adm: Pgpool 管理函数](https://pgpool.net/)

`pgpool_adm` 扩展为 Pgpool-II PCP（Pgpool 控制协议）命令提供可从 SQL 调用的包装函数，允许在 PostgreSQL 内管理 Pgpool-II。

### 可用函数

| 函数 | 描述 |
|----------|-------------|
| `pgpool_adm_pcp_node_info` | 显示给定后端节点的信息 |
| `pgpool_adm_pcp_health_check_stats` | 显示节点的健康检查统计 |
| `pgpool_adm_pcp_pool_status` | 从 pgpool.conf 获取参数 |
| `pgpool_adm_pcp_node_count` | 获取后端节点数量 |
| `pgpool_adm_pcp_attach_node` | 附加后端节点 |
| `pgpool_adm_pcp_detach_node` | 分离后端节点 |
| `pgpool_adm_pcp_proc_info` | 显示 Pgpool-II 子进程信息 |

### 调用方式

函数支持两种调用约定：

**直接参数**（主机名、端口、用户名、密码，加上函数特定参数）：

```sql
SELECT * FROM pgpool_adm_pcp_node_info('localhost', 9898, 'admin', 'password', 0);
SELECT * FROM pgpool_adm_pcp_node_count('localhost', 9898, 'admin', 'password');
SELECT * FROM pgpool_adm_pcp_pool_status('localhost', 9898, 'admin', 'password');
```

**外部服务器引用**（使用端口 9898 和 `~/.pcppass` 中的凭据）：

```sql
SELECT * FROM pgpool_adm_pcp_node_info(server_name := 'pgpool_server', node_id := 0);
SELECT * FROM pgpool_adm_pcp_node_count(server_name := 'pgpool_server');
```

### 节点管理

```sql
-- 分离后端节点
SELECT pgpool_adm_pcp_detach_node('localhost', 9898, 'admin', 'password', 1);

-- 重新附加后端节点
SELECT pgpool_adm_pcp_attach_node('localhost', 9898, 'admin', 'password', 1);
```

默认 PCP 通信端口为 9898。凭据可通过用户主目录中的 `.pcppass` 文件管理。
