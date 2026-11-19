## 用法

来源：

- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/README.md)
- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/pgext/pgtier.control)
- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/pgext/pgtier--1.0.sql)
- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/pgext/pgtier.c)
- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/docs/BUILD.md)
- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/LICENSE)

`pgtier` `1.0` 将 PostgreSQL 后端移入不同 Linux cgroup，并为外部 scx_pgtier 调度器切换调度策略。这是早期实验项目，目前仅有虚拟机中的 PostgreSQL 16 测试证据。

### 工作流程

管理员需先准备 cgroup v2 目录，向 PostgreSQL 操作系统用户授予所需写权限，并在 `/sys/fs/cgroup/ufs-pg/main` 内启动主进程。独立调度器要求 Linux 内核支持 sched_ext 和 BTF，并以 root 权限启动。准备 cgroup 目录树可能更改根目录的 CPU 控制器设置。

在新会话中加载模块，并配置已有应用角色：

```sql
ALTER SYSTEM SET session_preload_libraries = 'pgtier';
SELECT pg_reload_conf();
ALTER ROLE oltp_app SET pgtier.tier = 'ts';
ALTER ROLE report_user SET pgtier.tier = 'bg';
```

配置 `session_preload_libraries` 时保留已有条目，重载后重新连接。层级分配通过会话加载生效；`CREATE EXTENSION` 用于在数据库中提供可选诊断函数：

```sql
CREATE EXTENSION pgtier;
SELECT pgtier_hello();
```

### 配置与边界

`pgtier.tier` 接受 `none`、`ts`、`bg`；`pgtier.cgroup_root` 默认为 `/sys/fs/cgroup/ufs-pg`。两项设置均仅限超级用户，超级用户可通过 `GRANT SET ON PARAMETER pgtier.tier` 委派层级修改权。安装诊断函数同样需要超级用户。

诊断结果包含实际层级、调度策略和 cgroup。移动 cgroup 或切换策略失败时只记录警告，不会中断查询；因此配置的层级不能证明操作系统已执行隔离，后端可能停留在部分完成的中间状态。

时间敏感任务以牺牲后台任务资源为代价获得优先权。调度器可选的锁持有者轮询需要具备监控权限的独立数据库连接，不覆盖自旋锁或 LWLocks。目前没有通用延迟或生产安全保证。使用前应在可丢弃虚拟机中演练 cgroup 权限、调度器关闭和数据库重启。
