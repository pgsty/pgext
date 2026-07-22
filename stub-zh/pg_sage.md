## 用法

来源：

- [Current project README and architecture boundary](https://github.com/jasonmassie01/pg_sage/blob/492cd7b3ad840be3d8fd6ad1005a029f2396ecd7/README.md)
- [Version 0.5.0 extension SQL](https://github.com/jasonmassie01/pg_sage/blob/492cd7b3ad840be3d8fd6ad1005a029f2396ecd7/sql/pg_sage--0.5.0.sql)
- [Legacy extension entry point](https://github.com/jasonmassie01/pg_sage/blob/492cd7b3ad840be3d8fd6ad1005a029f2396ecd7/src/pg_sage.c)
- [Current reverse specification](https://github.com/jasonmassie01/pg_sage/blob/492cd7b3ad840be3d8fd6ad1005a029f2396ecd7/docs/REVERSE_SPEC.md)

pg_sage 0.5.0 是 pg_sage 仓库中被目录收录、基于预加载的 C 扩展。它收集快照、分析数据库健康状况、记录发现与动作，并在固定的 `sage` 模式中暴露控制及 JSON 辅助函数。当前上游产品已经转为独立 Go sidecar，并明确说明无需 PostgreSQL 扩展；不能混用这两套架构。

### 核心流程

对于已有的 0.5.0 扩展部署，预加载它和 `pg_stat_statements`，重启 PostgreSQL，再在目标数据库中创建扩展：

```ini
shared_preload_libraries = 'pg_stat_statements,pg_sage'
pg_sage.database = 'postgres'
```

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION pg_sage;

SELECT sage.status();
SELECT id, severity, title, recommendation
FROM sage.findings
WHERE status = 'open'
ORDER BY last_seen DESC;
```

应先使用观察或建议行为，审查每条拟执行 SQL 动作，并在启用任何自治模式前验证紧急停止机制。

### 重要对象

- `sage.snapshots`、`sage.findings` 和 `sage.explain_cache` 存储收集证据与分析结果。
- `sage.action_log` 记录已执行 SQL 及回滚元数据。
- `sage.config` 保存运行时配置，包括端点和自治模式。
- `sage.status`、`sage.emergency_stop` 和 `sage.resume` 暴露运维控制。
- `sage.explain`、`sage.suppress`、`sage.diagnose` 和 `sage.briefing` 暴露分析流程。
- `sage.health_json`、`sage.findings_json`、`sage.schema_json` 及相关辅助函数支持面向 sidecar 的 JSON 接口。

### 安全与生命周期说明

旧版库会安装执行器钩子、共享内存、libcurl 和四个后台工作进程，并要求超级用户安装。它可以收集查询文本和计划、连接 LLM 端点、保留与 API 密钥有关的配置，并执行建议 SQL，因此必须限制模式访问，并把配置与日志作为敏感信息保护。当前 README 和逆向规格描述的是走 wire protocol 的 sidecar，而不是这个扩展。应把 0.5.0 视为遗留部署边界，谨慎迁移，不能把当前 sidecar 指令直接用于 C 模块。
