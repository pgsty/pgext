## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ratelimit/pg_ratelimit.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ratelimit/pg_ratelimit--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ratelimit/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ratelimit/pg_ratelimit.c)

`pg_ratelimit` `1.0` 提供共享内存限流与会话值存储。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_ratelimit;
SELECT ratelimit_check('demo', 100, 1, '1 second');
SELECT ratelimit_remaining('demo');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

执行限流操作前需预加载 pg_ratelimit 并重启。pg_ratelimit.max_entries 限制共享条目，令牌桶和滑动窗口函数跟踪额度，session_set/get/touch 管理可过期会话记录。共享限流状态会在重启时丢失，也不是事务型应用数据。应隔离键名并限制重置、会话操作，共享键本身不构成租户安全边界。

此模块对应上游旧名 `alohadb_ratelimit`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
