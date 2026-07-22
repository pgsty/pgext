## 用法

来源：

- [官方 0.0.6 Rust 软件包清单](https://github.com/supabase/supamonitor/blob/497f1ecc66e297f04b32954efdfabf0ac21b52ff/Cargo.toml)
- [官方 0.0.6 钩子实现](https://github.com/supabase/supamonitor/blob/497f1ecc66e297f04b32954efdfabf0ac21b52ff/src/lib.rs)
- [官方当前开发清单](https://github.com/supabase/supamonitor/blob/3f8c4b8ec72dc60132513b22c4bbd65eb2a2c4a5/Cargo.toml)

`supamonitor` 0.0.6 是 Supabase 的 pgrx 预加载模块，将逐语句规划和执行计时以 JSON 记录写入 PostgreSQL 服务器日志。它是日志生产者，不是 SQL 统计视图或持久聚合器。已复核的 0.0.6 发布提交支持 PostgreSQL 13 至 18 构建特性。

### 预加载与安装

模块拒绝普通会话加载。将它加入预加载配置，重启服务器，然后登记扩展：

```conf
shared_preload_libraries = 'supamonitor'
```

```sql
CREATE EXTENSION supamonitor;
```

没有有文档的 SQL 函数或 GUC。移除或升级模块同样需要受控的服务器重启。

### 日志记录形状

计划器、执行器与工具语句钩子生成的消息，以带版本前缀开头，后接 JSON：

```text
supamonitor_0.0.6_log:{"query_id":123,"query":"SELECT 1","calls":1,"total_plan_time":0.0,"total_exec_time":0.1}
```

规划记录的 `calls` 为零并填充 `total_plan_time`；执行完毕或工具语句记录的 `calls` 为一并填充 `total_exec_time`。尽管字段名称如此，这些是单个事件，必须由外部日志管道收集并聚合。

### 安全与开销

记录包含经过清理但未规范化的查询文本，因此 SQL 字面量可能把凭据、个人数据、token 或租户标识符暴露到日志中。应实施日志访问控制、加密、下游脱敏、有界保留和摄取背压。查询没有计时状态时，执行器钩子会请求完整 instrumentation；所有主要钩子都会增加计时和 JSON 序列化工作，因此应以代表性查询速率与大量工具语句的工作负载测试开销。

项目仍在活动开发，但 API 不稳定。目录中的发布提交报告 0.0.6，而已复核的当前默认分支清单在一次实质不同的源码重写后报告 0.0.0。应固定准确的发布源码与 PostgreSQL 大版本，而不是构建未限定的分支。投入生产前，应测试与其他预加载扩展的钩子链接、预备语句、嵌套查询、错误、日志轮转和收集器故障。
