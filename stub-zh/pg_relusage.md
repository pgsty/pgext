## 用法

来源：

- [README.md](https://github.com/adept/pg_relusage/blob/017ca90bc08aa131d074017511ea9ed6ab89d6d3/README.md)
- [Control file](https://github.com/adept/pg_relusage/blob/017ca90bc08aa131d074017511ea9ed6ab89d6d3/pg_relusage.control)
- [Implementation](https://github.com/adept/pg_relusage/blob/017ca90bc08aa131d074017511ea9ed6ab89d6d3/pg_relusage.c)

`pg_relusage` 1.0.2 记录执行计划简化后语句使用的关系，支持 PostgreSQL 10–18，不定义 SQL 对象。加载共享库才会启用功能；可选的 CREATE EXTENSION 仅登记安装状态。

### 核心用法

```sql
LOAD 'pg_relusage';
SET pg_relusage.nested_statements = 'fold';
SET pg_relusage.log_level = 'notice';
CREATE TABLE relusage_demo (id integer);
SELECT * FROM relusage_demo;
```

### 运行边界

全部设置都需要超级用户：`pg_relusage.enabled` 默认 on，`pg_relusage.log_level` 默认 log，`pg_relusage.rel_kinds` 默认 riSvmfp。`pg_relusage.nested_statements` 可选 `all`（默认，每条语句单独记录）、`top`（忽略嵌套语句）或 `fold`（将嵌套关系合并到外层日志）。1.0.1 修复子事务失败后 fold 模式的处理。

用 `session_preload_libraries` 可在新会话中加载，无需重启；用 `shared_preload_libraries` 覆盖全部会话则需要重启，应保留已有预加载项。LOAD 通常要求超级用户。日志不包含 SQL 文本，可配置 `log_line_prefix` 标识客户端。

系统目录不记录。规划或执行器启动时被排除的关系不记录，但执行过程中跳过的分区仍可能出现。向分区表 INSERT 时只记录父表。COMMIT 时的延迟触发器、COPY FROM 触发器及规划器调用的函数按顶层语句处理。DDL 与维护命令不记录。日志中暂未出现某对象，并不足以证明可以删除它。

### 1.0.2 与分区表

1.0.2 降低了涉及大量分区时的关系集合检查与分区匹配开销，日志格式和配置不变。替换共享库后，使用共享预加载的实例需要重启，使用会话预加载或 LOAD 的连接需要重新连接。若曾用 CREATE EXTENSION 登记安装，还应执行 `ALTER EXTENSION pg_relusage UPDATE`。
