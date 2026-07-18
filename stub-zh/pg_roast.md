## 用法

来源：

- [上游 README](https://github.com/samirketema/pg_roast/blob/ccbf012d01ebbb8edcb13b02add981705dab2308/README.md)
- [1.0 版安装 SQL](https://github.com/samirketema/pg_roast/blob/ccbf012d01ebbb8edcb13b02add981705dab2308/pg_roast--1.0.sql)
- [后台工作进程实现](https://github.com/samirketema/pg_roast/blob/ccbf012d01ebbb8edcb13b02add981705dab2308/pg_roast_bgw.c)

pg_roast 运行一组带有明确观点的 PostgreSQL 健康检查，并将发现存入其固定的 roast 模式。它检查配置、模式设计、索引、清理与膨胀指标、安全状态、复制、连接和负载信号。1.0 版面向 PostgreSQL 14 及更高版本。

### 手动审计

手动模式不需要预加载：

```sql
CREATE EXTENSION pg_roast;

SELECT * FROM roast.run();
SELECT severity, check_name, object_name, roast
FROM roast.latest
ORDER BY severity, check_name;

SELECT * FROM roast.summary;
```

每次运行都会持久化审计历史与发现。使用 latest 视图查看最新一次运行，使用 summary 视图查看分组结果。

### 定时审计

可选的后台工作进程需要预加载配置与重启：

```ini
shared_preload_libraries = 'pg_roast'
pg_roast.database = 'mydb'
pg_roast.interval = 3600
```

数据库设置在服务器启动时固定。在生产负载上启用自动审计前，应审查上游设置。

### 注意事项

- 发现是启发式建议，并不自动证明存在缺陷。应在应用任何建议前，审查工作负载背景、维护窗口和 PostgreSQL 文档。
- 审计会执行系统目录与统计查询，并写入历史表。应在大型系统目录上测量开销，并防止不可信用户访问 roast 模式。
- 结果可能暴露对象名、配置、安全发现和查询相关运维细节。应实施最小权限与明确的保留策略。
- 手动运行无需预加载，但后台工作进程必须预加载；修改仅启动时设置需要重启。
