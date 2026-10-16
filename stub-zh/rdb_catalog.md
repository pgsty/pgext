## 用法

来源：

- [docker/raodb/README](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/README)
- [docker/raodb/ext/rdb_catalog/c.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/c.sql)
- [docker/raodb/ext/rdb_catalog/Makefile](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/Makefile)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.0.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.0.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.0--1.1.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.0--1.1.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.1--1.2.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.1--1.2.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.2--1.3.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.2--1.3.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.3--1.4.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.3--1.4.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.4--1.5.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.4--1.5.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.5--1.6.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.5--1.6.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog--1.6--1.7.sql](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog--1.6--1.7.sql)
- [docker/raodb/ext/rdb_catalog/rdb_catalog.control](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/docker/raodb/ext/rdb_catalog/rdb_catalog.control)

`rdb_catalog` 1.7 是 RAODB 开发树中的 SQL 扩展，在 `public` 中安装系统目录封装与诊断视图；它不提供监控采集器或自动调优。

### 核心用法

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_qualstats;
CREATE EXTENSION rdb_catalog;
SELECT * FROM public.v$databases;
SELECT * FROM public.ses$blocked;
SELECT * FROM public.vac$needed;
```

### 运行边界

由管理员安装。control 文件的依赖列表为空，但 1.4→1.6 的 SQL 脚本直接查询 `pg_qualstats` 与 `pg_stat_statements`。加载完整的 1.7 升级链前，需在同一数据库准备这些扩展，并配置其采集和预加载条件。此扩展自身没有共享库。所核对源码未声明 PostgreSQL 大版本兼容范围或再分发许可。

`v$` 视图涵盖对象、权限、分区和统计；`ses$` 查看会话、阻塞和锁；`stx$` 检查统计信息状态；`ind$` 汇总索引与谓词观测；`vac$` 查看 vacuum 活动和冻结风险。部分底层目录需要特权访问。应按需授予视图权限，结果可能包含数据库结构、会话查询和用户映射选项。

诊断名称只代表启发式筛选。例如，`ind$safe_to_drop` 使用当前使用计数并排除主键命名，但不能证明索引可安全删除。仍需检查计数重置、尚未观测到的工作负载及约束。脚本不会自动删除索引。

默认版本由 1.0 安装脚本和七个连续升级脚本组成，必须提供完整 SQL 文件集。使用前需按目标服务器及依赖版本核对脚本。备份中应保留对内置等待事件字典表的修改；本次仅收录源码，不代表通过了 Pigsty 软件包验证。
