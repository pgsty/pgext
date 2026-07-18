## 用法

来源：

- [固定版本 Rust 实现](https://github.com/boringSQL/pg_sysload/blob/7c4db3cef4df26dd12718d43c83a75f1f0b5508b/src/lib.rs)
- [固定版本 Cargo 元数据](https://github.com/boringSQL/pg_sysload/blob/7c4db3cef4df26dd12718d43c83a75f1f0b5508b/Cargo.toml)
- [固定版本扩展控制文件](https://github.com/boringSQL/pg_sysload/blob/7c4db3cef4df26dd12718d43c83a75f1f0b5508b/pg_sysload.control)

pg_sysload 0.0.0 把 Linux procfs 负载均值的前三个字段暴露为浮点数组，分别代表大约一、五和十五分钟内的主机可运行负载平均值。

### 查询主机负载

```sql
CREATE EXTENSION pg_sysload;

SELECT sys_loadavg();
```

无需预加载；如果后端找不到 /proc/loadavg，加载共享库就会失败。

### 解释与可移植性

负载均值反映主机或容器命名空间状态，而不是 PostgreSQL 专属工作量。它包含其他进程中可运行与不可中断的任务，也没有除以 CPU 数量。跨服务器比较前必须记录 CPU 配额、cgroup 边界、容器 procfs 行为和采样时间。归因负载时应结合 PostgreSQL 与操作系统监控。

每次 SQL 调用都会读取并解析 /proc/loadavg。某个数字无法解析时，实现会静默丢弃该元素，因此调用方应验证返回数组恰有三个有限值。procfs 读取失败会在调用后端抛出错误。

项目没有 README、用户文档、SQL 测试或兼容策略。Cargo 特性基于 pgrx 0.11.4，只覆盖 PostgreSQL 11 至 16，不含 17 与 18；仓库最后更新于 2024 年。如果主机遥测敏感，应限制执行权并限制采集频率；生产监控优先使用受维护的 node exporter 或平台代理。
