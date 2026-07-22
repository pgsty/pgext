## 用法

来源：

- [官方 README](https://github.com/vibhorkum/efm_extension/blob/ecb305cc185b05eaf0e766d6a9c5d23b3072e2fb/README.md)
- [扩展控制文件](https://github.com/vibhorkum/efm_extension/blob/ecb305cc185b05eaf0e766d6a9c5d23b3072e2fb/efm_extension.control)
- [1.1 版安装 SQL](https://github.com/vibhorkum/efm_extension/blob/ecb305cc185b05eaf0e766d6a9c5d23b3072e2fb/efm_extension--1.1.sql)

`efm_extension` 通过 SQL 暴露 EDB Failover Manager 状态与管理命令。它通过受限制的 `sudo` 边界执行已安装的 EFM 命令行程序，增加结构化状态、缓存、历史和监控视图，并将只读监控与仅超级用户可用的拓扑变更分离。

### 配置与监控

先安装并配置 EFM，为 PostgreSQL 操作系统账号创建严格限定的 sudoers 策略，再设置集群、二进制、属性、Java 与 sudo 路径。可选的缓存/历史工作进程需要预加载并重启。

```conf
shared_preload_libraries = 'efm_extension'
efm.cluster_name = 'efm'
efm.command_path = '/usr/edb/efm-4.9/bin/efm'
efm.properties_location = '/etc/edb/efm-4.9'
efm.java_home = '/usr/lib/jvm/java-11-openjdk'
efm.bgw_enabled = true
```

```sql
CREATE EXTENSION dblink;
CREATE EXTENSION pgcrypto;
CREATE EXTENSION efm_extension;

SELECT efm_extension.grant_access_to_user('monitoring_user');
SELECT efm_extension.efm_cluster_status_json();
SELECT * FROM efm_extension.efm_nodes_details;
SELECT * FROM efm_extension.efm_metrics;
```

`efm_cluster_status()`、`efm_cluster_status_json()`、`efm_get_nodes()`、`efm_list_properties()` 以及缓存/指标视图组成监控表面。`efm_allow_node()`、`efm_disallow_node()`、`efm_set_priority()`、`efm_failover()`、`efm_switchover()` 与 `efm_resume_monitoring()` 仍是超级用户管理操作。

### 故障切换与主机安全边界

管理调用会改变外部 HA 集群，外围 PostgreSQL 事务中止时不会回滚。故障切换或主备切换还可能断开调用它的会话。应要求明确的运维授权、串行化变更、确认候选节点健康度与延迟，并独立验证最终 EFM 和 PostgreSQL 拓扑。

数据库后端能够以 EFM 账号执行已配置的主机命令。应固定绝对路径、在 sudoers 中只允许准确的 EFM 子命令、保护 GUC 与属性文件，且不向不可信角色授予配置或函数权限。缓存状态可能按 `efm.cache_ttl` 产生陈旧，而后台历史持久化会增加本地写入和保留需求。应将 1.1 版扩展与准确的 EFM CLI/输出版本配套，并在生产采用前演练退出码映射、超时、stderr、工作进程重启以及 EFM 连接丢失。
