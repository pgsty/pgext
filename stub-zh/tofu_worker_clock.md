## 用法

来源：

- [Official source](https://github.com/NiuTrans/ToFu/blob/27c05b79b32b0ffd78398d8a98cdfdd1ff1afad0/docs/STORAGE_OPERATIONS.md)
- [Official source](https://github.com/NiuTrans/ToFu/blob/27c05b79b32b0ffd78398d8a98cdfdd1ff1afad0/lib/storage_authority_sql/native_worker_clock/tofu_worker_clock.control)
- [Official source](https://github.com/NiuTrans/ToFu/blob/27c05b79b32b0ffd78398d8a98cdfdd1ff1afad0/lib/storage_authority_sql/native_worker_clock/tofu_worker_clock--1.0.sql)
- [Official source](https://github.com/NiuTrans/ToFu/blob/27c05b79b32b0ffd78398d8a98cdfdd1ff1afad0/lib/storage_authority_sql/native_worker_clock/worker_clock.c)
- [Official source](https://github.com/NiuTrans/ToFu/blob/27c05b79b32b0ffd78398d8a98cdfdd1ff1afad0/lib/storage_authority_sql/native_worker_clock/generated_clock_contract.h)

`tofu_worker_clock` 1.0 为 ToFu 的 SQL 工作任务提供由服务端维护的租约时钟与纪元。它要求 PostgreSQL 17 或更新版本，是 ToFu 应用仓库内的独立原生模块。应用发行号不等于扩展的 control 版本。

### 基本流程

由部署运维人员安装匹配的扩展文件，再以超级用户在参与的每个数据库中创建扩展；应用启动不会自动安装。固定模式为 `public`。

```sql
CREATE EXTENSION tofu_worker_clock;
SELECT public.tofu_worker_lease_clock_v1();
```

函数返回 JSONB，包含 `format`、`epoch`、`nowMs`、`utcMs` 和 `stateBytes`。`nowMs` 是共享的单调高水位租约时间，`utcMs` 是采样得到的墙上时钟时间。绝对 UTC 调度时间应与租约到期判断分开。`epoch` 标识服务进程生命周期；ToFu 还会结合数据库 OID 限定该身份。

### 生命周期与故障边界

首次调用会按需创建大小受限的共享动态内存状态，其他会话共享该状态。关闭客户端或替换连接池不会重置它。启用函数不需要预加载、时钟表、后台工作进程或重启。

PostgreSQL 重启会产生新纪元，使之前的 ToFu 租约凭据失效。该机制不是持久化序列、跨服务器时钟，也不保证墙上时钟不会回退。上游 SQL 工作任务层遇到缺失或不兼容的输出会拒绝操作，而不会退回客户端时间。集成工作任务认领时，应限制安装权限并保持返回契约不变。
