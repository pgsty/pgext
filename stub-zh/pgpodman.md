## 用法

来源：

- [Crunchy Bridge Container Apps 文档](https://docs.crunchybridge.com/container-apps)
- [Crunchy Bridge 产品页](https://www.crunchydata.com/products/crunchy-bridge)

`pgpodman` 是 Crunchy Bridge Container Apps 背后的服务商扩展。它从 SQL 在托管数据库主机上启动和管理容器镜像。该功能目前处于 public beta，扩展版本由服务控制，而不是作为可移植软件包发布。

### 安装与检查

在数据库旁运行容器属于特权操作。Crunchy Bridge 要求由超级用户在 `postgres` 数据库中安装：

```sql
CREATE EXTENSION pgpodman;

SELECT pgpodman_version();
SELECT * FROM list_containers(true);
SELECT * FROM list_images();
```

API 包括 `run_container()`、`start_container()`、`stop_container()`、`container_status()`、`container_logs()`、`remove_container()` 和镜像管理函数。容器选项以命令字符串传入，因此执行前必须审查每个镜像摘要、端口、挂载、环境变量和权限。

### 安全与服务限制

`run_container()` 可以在 PostgreSQL 旁执行任意公开容器镜像。应把它视为位于数据库网络边界内的主机级代码执行，而不是普通 SQL 辅助函数。固定受信镜像摘要，最小化容器权限和出站访问，应用服务商防火墙控制，并且绝不能把长期凭据放入 SQL 文本、审计轨迹、语句日志、容器参数或镜像层。容器日志也可能包含数据库或应用数据。

beta 期间只开放公网端口 `5433` 至 `5442`，容器文件系统用量限制为 2 GB。维护、故障转移或调整规格后，服务大约每 10 分钟检查 `pgpodman_status`，并重启应处于运行状态的容器；这并非即时可用性保证。停止容器不会删除它，而且删除前必须先停止。运行方案中应包含镜像清理、端口归属、备份范围、故障转移恢复时间、监控和服务商支持升级流程。
