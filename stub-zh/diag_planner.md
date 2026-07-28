## 用法

来源：

- [官方扩展控制文件（diag_planner.control）](https://github.com/masahikosawada/incubator/blob/ae553ea4cc728b2d7a742a7ecf721996b3ada8b5/diag_planner/diag_planner.control)
- [官方实现源代码](https://github.com/masahikosawada/incubator/blob/ae553ea4cc728b2d7a742a7ecf721996b3ada8b5/diag_planner/diag_planner.c)

`diag_planner` 是一个轻量级的规划诊断模块，来自孵化库。加载该库会安装关系和连接路径挂钩，这些挂钩会发出 `NOTICE` 消息，描述候选扫描或连接路径及其估算成本。

### 核心工作流

使用匹配的 PostgreSQL 服务器头文件构建并安装 `diag_planner` 库，然后仅在隔离的诊断会话中通过服务器支持的库加载机制加载该库。运行一个代表性查询并检查发出的扫描和连接路径通知。

审查的源代码不安装 SQL 对象，并未记录独立的 `CREATE EXTENSION` 命令。

### 诊断输出

- 扫描通知区分顺序扫描、采样扫描、索引扫描、索引仅扫描、位图索引扫描和位图堆扫描路径。
- 连接通知区分哈希连接、合并连接和嵌套循环连接候选方案。
- 每个报告的路径包括规划器的启动成本和总成本估算。

### 要求与注意事项

- 审查的控制文件、注册表或目录证据标识版本 `1.0`。
- 控制文件将该扩展标记为不可重定位。
- 该模块挂钩 PostgreSQL 规划器内部，来自孵化树；兼容性与其构建所针对的具体服务器源代码相关。
- 它在规划期间发出输出，仅用于诊断，而非常规生产工作负载。
