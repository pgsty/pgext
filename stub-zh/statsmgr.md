## 用法

来源：

- [官方 README](https://codeberg.org/Data-Bene/StatsMgr/src/commit/4e934fb1f74178c8300160b456531ae6850068e7/README.statsmgr.md)
- [扩展 SQL](https://codeberg.org/Data-Bene/StatsMgr/src/commit/4e934fb1f74178c8300160b456531ae6850068e7/statsmgr--0.1-alpha.sql)
- [C 实现](https://codeberg.org/Data-Bene/StatsMgr/src/commit/4e934fb1f74178c8300160b456531ae6850068e7/statsmgr.c)

`statsmgr` 0.1-alpha 是预览扩展，把 PostgreSQL 累计的归档、后台写、检查点、I/O、SLRU 和 WAL 统计定期快照到动态共享内存。它面向 PostgreSQL 17 及以上版本，只保留内存中的滚动历史，并非持久监控存储。

### 设置模式

临时评估时，创建扩展并手动启动主动态后台工作进程：

```sql
CREATE EXTENSION statsmgr;
SELECT statsmgr_start_main_worker();
SELECT statsmgr_snapshot();
SELECT * FROM statsmgr_wal();
```

若要自动启动，应预加载模块并重启 PostgreSQL：

```conf
shared_preload_libraries = 'statsmgr'
```

每个结果行会在对应核心统计字段之外增加 `tick` 与 `tick_tz`。

### 函数索引

- 管理函数：`statsmgr_start_main_worker()`、`statsmgr_stop_main_worker()`、`statsmgr_snapshot()` 和 `statsmgr_reset()`。
- 历史读取函数：`statsmgr_archiver()`、`statsmgr_bgwriter()`、`statsmgr_checkpointer()`、`statsmgr_io()`、`statsmgr_slru()` 和 `statsmgr_wal()`。

当前实现使用 512 槽环，并固定每 60 秒执行一次周期快照。源码中虽然存在间隔与各类开关变量，但没有注册为 PostgreSQL GUC，README 的配置章节也被注释，因此该版本不能在运行时配置这些值。

### 安全与运维说明

SQL 安装脚本没有撤销 PUBLIC 对工作进程启动停止、重置和手动快照函数的执行权限，已检查的 C 包装器也没有显式超级用户检查。应立即撤销 PUBLIC 对管理函数的权限，只授予管理角色；评估数据暴露后，再把读取函数单独授予监控角色。状态位于集群级共享内存并会在重启时丢失，而扩展对象属于单个数据库，因此应协调一个控制数据库。后台工作进程会占用工作进程槽位。该 alpha 实现依赖 PostgreSQL 17 内部接口；离开测试集群前，应验证准确小版本构建、并发管理调用、环回绕、重启行为与资源限制。
