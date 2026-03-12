

## 用法

> [plprofiler: 用于 PL/pgSQL 函数性能分析的服务端支持](https://github.com/bigsql/plprofiler)

`plprofiler` 是一个 PL/pgSQL 函数性能分析工具，能够识别性能瓶颈并生成交互式火焰图报告。

```sql
CREATE EXTENSION plprofiler;
```

### 配置

在 `postgresql.conf` 中添加：

```
shared_preload_libraries = 'plprofiler'
```

### 命令行工具

`plprofiler` 命令行工具用于控制性能分析并生成报告：

```bash
# 分析特定 SQL 命令
plprofiler run -d mydb --command "SELECT my_function()" --output report.html

# 实时监控性能分析
plprofiler monitor -d mydb

# 保存性能分析数据以便后续分析
plprofiler save -d mydb --name "my_profile"

# 生成包含火焰图的 HTML 报告
plprofiler report -d mydb --from-data "my_profile" --output report.html

# 列出已保存的性能分析数据集
plprofiler list -d mydb

# 重置性能分析数据
plprofiler reset-data -d mydb

# 导出/导入性能分析数据
plprofiler export -d mydb --from-data "my_profile" > profile.json
plprofiler import -d mydb --into-data "imported" < profile.json
```

### SQL 接口

```sql
-- 为当前会话启用性能分析
SELECT pl_profiler_set_enabled_local(true);

-- 执行待分析的函数
SELECT my_function();

-- 将性能分析数据收集到共享哈希表中
SELECT pl_profiler_collect_data();

-- 禁用性能分析
SELECT pl_profiler_set_enabled_local(false);

-- 全局启用性能分析（适用于所有会话）
SELECT pl_profiler_set_enabled_global(true);

-- 重置本地/共享性能分析数据
SELECT pl_profiler_reset_local();
SELECT pl_profiler_reset_shared();
```

### 报告输出

生成的 HTML 报告包括：

- **交互式火焰图**：展示 PL/pgSQL 代码中花费的时钟时间
- **每函数统计信息**：包含自身时间（总时间减去子函数时间）
- **排名靠前的函数**：按耗时排序（默认：前 10 个）
- 自包含的 HTML，无需外部依赖

### 分析方法

- **直接分析**：在收集数据时运行特定 SQL
- **定时采集**：基于时间间隔的统计收集
- **按用户分析**：通过 `ALTER USER` 为特定数据库用户启用性能分析
- **生产监控**：在生产系统上进行低开销的性能分析
