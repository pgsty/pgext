## 用法

来源：

- [官方 README](https://github.com/arrxy/pg_flux/blob/794ca3e531a3aea2c4cce112ceece18ec84fd342/README.md)
- [控制文件](https://github.com/arrxy/pg_flux/blob/794ca3e531a3aea2c4cce112ceece18ec84fd342/pg_gftr.control)
- [扩展 SQL](https://github.com/arrxy/pg_flux/blob/794ca3e531a3aea2c4cce112ceece18ec84fd342/sql/pg_gftr--0.1.0.sql)
- [上游验证记录](https://github.com/arrxy/pg_flux/blob/794ca3e531a3aea2c4cce112ceece18ec84fd342/VALIDATION.md)

`pg_gftr` 0.1.0 是面向 PostgreSQL 18 的实验性 CUDA 排序分组聚合扩展。GPU 执行需要 NVIDIA GPU、兼容驱动，以及使用 CUDA Toolkit 12.2 或更高版本构建的 CUDA 版本。独立的 CPU 后端用于正确性参照，无需修改 PostgreSQL 内核。

### 核心工作流

由超级用户安装并加载，然后显式启用规划器路径：

```sql
CREATE EXTENSION pg_gftr;
LOAD 'pg_gftr';
SELECT pg_gftr_version();
SET pg_gftr.enabled = on;
SET pg_gftr.backend = 'cuda';
SET pg_gftr.mode = 'gftr';

CREATE TEMP TABLE measurements (k integer, a integer, b integer);
INSERT INTO measurements
SELECT g % 100, g % 1000, g % 50 FROM generate_series(1, 200000) g;
ANALYZE measurements;
EXPLAIN (ANALYZE, BUFFERS)
SELECT k, count(*), count(a), sum(a), min(b), max(b)
FROM measurements GROUP BY k;
```

符合条件的计划可以使用 `Custom Scan (GftrAggregate)`。PostgreSQL 负责子计划扫描、过滤及 CPU 连接，扩展对其输出行做聚合。当前支持单个普通整数分组列和示例中的内置整数聚合，包括空分组键和空值。

### 控制参数

| 设置 | 默认值 | 用途 |
| --- | --- | --- |
| `pg_gftr.enabled` | off | 提供自定义聚合路径 |
| `pg_gftr.force` | off | 为实验人为降低符合条件路径的成本 |
| `pg_gftr.backend` | cuda | 选择 CUDA 或显式 CPU 参照实现 |
| `pg_gftr.mode` | gftr | 选择 GFTR 或 GFUR CUDA 聚合方式 |
| `pg_gftr.device` | 0 | 可见 GPU 集合内的设备编号 |
| `pg_gftr.min_rows` | 100000 | 非强制模式下的估算输入行数阈值 |
| `pg_gftr.max_input_mb` | 256 | 打包输入上限，单位 MiB |
| `pg_gftr.max_device_mb` | 4096 | 单个算子的设备内存上限，单位 MiB |
| `pg_gftr.setup_cost` | 1000 | 暂定的启动成本 |
| `pg_gftr.tuple_cost` | 0.02 | 暂定的逐行成本 |

成本模型尚未针对所有 GPU 校准。修改规划器参数后应重新准备缓存语句。`CUDA_VISIBLE_DEVICES` 限制可用设备集合。CUDA 执行失败会报查询错误，不会自动退回 CPU。

### 会话与资源边界

安装脚本会在创建扩展的会话中加载钩子。每个新会话仍须加载动态库、调用版本函数，或使用管理员配置的预加载。会话加载不需要重启服务器。仅安装 SQL 对象不会激活其他会话，而且规划器路径在显式启用前保持关闭。

输入会被完整缓冲。主机端结果、排序工作区与 PostgreSQL 子计划内存不计入输入上限。设备临时空间计入设备上限，并进一步受当前空闲 GPU 内存 90% 的约束；并发会话之间不协调内存预留。不支持磁盘溢写、部分/最终聚合或多 GPU 执行。超出容量会使语句失败，不返回部分结果。正在进行的 GPU 操作无法立即中断。

当前版本不支持多个分组键、表达式键、非整数输入列、无分组键聚合、去重或带过滤条件的聚合、分组集和窗口查询；这些查询保留原生路径。上游记录了 RTX A6000 上的 GPU 正确性检查及 macOS 上的 CPU 检查，该记录中的 A100 执行仍未验证。应将其视为研究实现，测量包含数据传输在内的完整 SQL 查询。
