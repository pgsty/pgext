## 用法

来源：

- [官方上游文档](https://github.com/mpihlak/pg_stat_usage/blob/8e8173b3a8fd61a36491e86b4c9d8c021ad77a30/README.md)
- [所需 PostgreSQL 核心补丁](https://github.com/mpihlak/pg_stat_usage/blob/8e8173b3a8fd61a36491e86b4c9d8c021ad77a30/patches/enable-hooks.patch)
- [官方扩展 SQL](https://github.com/mpihlak/pg_stat_usage/blob/8e8173b3a8fd61a36491e86b4c9d8c021ad77a30/pg_stat_usage--1.0.sql)

`pg_stat_usage` 1.0 是已放弃的原型，用于在单个 PostgreSQL 后端内关联存储函数调用与关系或索引活动。它无法在标准服务器上运行：扩展要求应用随附的旧 PostgreSQL 核心补丁，以添加私有统计钩子和关系关闭回调。

### 补丁服务器流程

只有使用准确匹配的补丁构建 PostgreSQL 服务器后，才能将模块安装到本地预加载所需路径，并创建 SQL 对象：

```conf
local_preload_libraries = 'pg_stat_usage'
```

```sql
CREATE EXTENSION pg_stat_usage;

SELECT * FROM pg_stat_usage;
SELECT pg_stat_usage_reset();
```

`local_preload_libraries` 影响新启动的会话。应确认共享库位置：控制文件引用 `$libdir/plugins/pg_stat_usage`，但 README 和默认 PGXS 构建使用较短的加载名。

### 报告列

`pg_stat_usage` 视图包装 `pg_stat_usage()`，报告对象 OID、父函数 OID、对象类型、模式与名称、调用与扫描次数、总时间与自身时间、元组读取或变更计数，以及块读取或命中计数。父 OID 可以近似表示访问表或索引时哪个函数处于活动状态。`pg_stat_usage_reset()` 清除当前后端收集的计数器。

### 原型限制

收集状态位于进程本地；不存在聚合所有后端的共享或外部收集器。源码明确说明，部分关系计数器的事务提交或回滚归因只有偶尔正确，而且始终假定提交。计时使用墙上时钟，关系上下文也可能只是近似值。虽然 SQL 视图授予 `PUBLIC`，但它只暴露调用后端自身的状态。

所需补丁面向已被替代的统计子系统，不能假定可应用到现代 PostgreSQL。不得随意将其移植到生产服务器。若用于研究，应隔离补丁构建、协调模块安装路径、限制预加载与重置权限，并测试嵌套函数、错误、回滚、并行查询和连接更替。生产调用与对象统计应使用受支持的可观测性扩展。
