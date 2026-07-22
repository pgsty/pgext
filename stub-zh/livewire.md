## 用法

来源：

- [官方 README](https://github.com/b4oshany/livewire/blob/547e59a951fe5d088d89039a2d078186e73c7862/README.md)
- [0.2.0 版 SQL 实现](https://github.com/b4oshany/livewire/blob/547e59a951fe5d088d89039a2d078186e73c7862/livewire--0.2.0.sql)
- [扩展控制文件](https://github.com/b4oshany/livewire/blob/547e59a951fe5d088d89039a2d078186e73c7862/livewire.control)

`livewire` 为配电 GIS 数据构建并缓存可路由的影子网络。它把 PostGIS 几何与 pgRouting 风格的图遍历结合起来，使工程人员能够跨已配置的线路和设备图层追踪上游、下游、电源、阻断点与末端节点。

### 核心流程

先安装 `postgis`、`postgis_topology` 与 pgRouting 支持。用 SRID 初始化一个模式，通过 JSON 配置注册边和节点参与者，生成影子网络，再缓存电源追踪结果。

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_topology;
CREATE EXTENSION livewire;

SELECT lw_initialise('powerflow', 3448);
SELECT lw_generate('powerflow');
SELECT lw_traceall('powerflow');
```

`lw_addedgeparticipant()` 注册线性图层。`lw_addnodeparticipant()` 还接受电源与阻断谓词。`lw_tracesource()`、`lw_traceupstream()` 与 `lw_tracednstream()` 在准备好的网络上工作；`lw_sourcenodes()`、`lw_blocknodes()`、`lw_endnodes()` 与 `lw_nodesnearnode()` 提供配套的拓扑检查。

### 数据与维护边界

参与者 JSON 必须标明模式、表、主键、几何、标签、相位和相位映射字段。节点参与者还可定义 `sourcequery` 与 `blockquery`；它们是 SQL 条件，只能来自可信配置。当源数据位于 LiveWire 模式外时，初始化可能按类表语义复制数据，因此要确认所有权、权限、索引与刷新行为。

`lw_generate()` 会重建组合网络，而 `lw_traceall()` 会追踪每个电源，成本可能很高。文档模型虽然能保存相位数据，但把开关设备视为所有相位同时开启或同时关闭。所审阅上游面向 PostgreSQL 10 时代的 PostGIS/pgRouting；用于当前生产 GIS 数据前，应在一次性副本上验证函数兼容性并执行生成流程。
