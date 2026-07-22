## 用法

来源：

- [官方 README](https://github.com/johto/call_graph/blob/9bcba8927be7374263db9d0dd55ec6f8015664d9/README)
- [官方 1.4 版 SQL](https://github.com/johto/call_graph/blob/9bcba8927be7374263db9d0dd55ec6f8015664d9/call_graph--1.4.sql)
- [官方 control 文件](https://github.com/johto/call_graph/blob/9bcba8927be7374263db9d0dd55ec6f8015664d9/call_graph.control)

`call_graph` 对 PostgreSQL 函数执行进行插桩，累计调用者/被调用者边，并可据此绘制调用图。它面向 PostgreSQL 9.1 时代服务器，主要用于诊断：跟踪会给每次被插桩的函数调用增加工作量，可选的表使用跟踪开销尤其明显。

### 核心流程

预加载模块并重启 PostgreSQL，然后创建扩展：

```conf
shared_preload_libraries = 'call_graph'
```

只在待测会话或事务中启用采集；整理缓冲记录前先关闭采集：

```sql
CREATE EXTENSION call_graph;

SET call_graph.enable = true;
SELECT application_function();
SET call_graph.enable = false;

SELECT *
FROM call_graph.processcallgraphbuffers(10000);
```

处理函数最多消费指定数量的缓冲记录，并更新持久化图与统计表。仓库中的 `utils/create_graphs.pl` 工具可把这些表转换成图形输出。

### 配置与数据

- `call_graph.enable`：启用函数调用采集，默认关闭。
- `call_graph.track_table_usage`：同时记录调用图触及的关系，性能代价显著。
- `CallGraphBuffer` 与 `TableAccessBuffer`：用于降低后端锁争用的 unlogged 追加缓冲表。
- `CallGraphs` 与 `Edges`：整理后的调用图标识及调用者/被调用者数据。
- `TableUsage`、`DailyStats`、`HourlyStats`：关系使用情况与聚合统计。
- `call_graph_version()`：返回已安装扩展 API 版本。

### 注意事项

版本 `1.4` 是旧式服务器钩子模块；加载到实例前必须针对确切 PostgreSQL 大版本验证其 C 钩子。Unlogged 缓冲不具备崩溃持久性；执行整理的工作进程应关闭采集，避免产生自身记录。

安装脚本向 `PUBLIC` 授予了两个缓冲表的插入权限。上游明确警告恶意用户可以注入伪造记录。在把调用图当作诊断证据前，应撤销这些授权或隔离可信角色；启用表使用跟踪或实例级采集前也应先测试开销。
