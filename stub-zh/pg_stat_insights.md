## 用法

来源：

- [官方文档](https://github.com/pgElephant/pg_stat_insights/blob/07d83c9e3c606f4ef1b6b7d571d48c63d4ad57c5/README.md)
- [扩展控制文件](https://github.com/pgElephant/pg_stat_insights/blob/07d83c9e3c606f4ef1b6b7d571d48c63d4ad57c5/pg_stat_insights.control)
- [官方仓库](https://github.com/pgElephant/pg_stat_insights)

`pg_stat_insights` 提供查询规划与执行统计，以及复制、索引、趋势和争用分析。

### 启用

将 `pg_stat_insights` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_stat_insights`：

```ini
shared_preload_libraries = 'pg_stat_insights'
```

```sql
CREATE EXTENSION pg_stat_insights;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- ステップ 1：PostgreSQL 設定で拡張機能を有効化
ALTER SYSTEM SET shared_preload_libraries = 'pg_stat_insights';
-- PostgreSQL サーバーの再起動が必要

-- ステップ 2：データベースに拡張機能を作成
CREATE EXTENSION pg_stat_insights;

-- ステップ 3：最も遅いクエリを即座に表示
SELECT
    query,
    calls,
    total_exec_time,
    mean_exec_time,
    rows
FROM pg_stat_insights_top_by_time
LIMIT 10;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_stat_insights_top_by_time` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_stat_insights_by_bucket` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_stat_insights_errors` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_stat_insights_histogram_summary` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_stat_insights_index_by_bucket` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_stat_insights_index_size_by_bucket` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_stat_insights_plan_errors` | VIEW | 扩展创建的检查或查询视图。 |
| `pg_stat_insights_replication_by_bucket` | VIEW | 扩展创建的检查或查询视图。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 16, 17, 18；不要推断未列出的主版本。
- 预加载 `pg_stat_insights` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
