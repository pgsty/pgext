
## 用法

> [!WARNING] 该扩展缺乏维护

在所有集群节点上安装 `pgml` 扩展及 Python 依赖后，即可在 PostgreSQL 集群中启用 `pgml`。

使用 `patronictl` 命令[配置](/pgsql/admin/#config-cluster)集群，将 `pgml` 添加到 `shared_preload_libraries` 中，并在 `pgml.venv` 中指定 `venv` 目录：

```yaml
shared_preload_libraries: pgml, timescaledb, pg_stat_statements, auto_explain
pgml.venv: '/data/pgml'
```

完成后，重启数据库集群，然后使用 SQL 命令创建扩展：

```sql
CREATE EXTENSION vector;        -- 建议同时安装 pgvector！
CREATE EXTENSION pgml;          -- 在当前数据库中创建 PostgresML
SELECT pgml.version();          -- 打印 PostgresML 版本信息
```

如果一切正常，应当看到类似如下的输出：

```bash
# create extension pgml;
INFO:  Python version: 3.11.2 (main, Oct  5 2023, 16:06:03) [GCC 8.5.0 20210514 (Red Hat 8.5.0-18)]
INFO:  Scikit-learn 1.3.0, XGBoost 2.0.0, LightGBM 4.1.0, NumPy 1.26.1
CREATE EXTENSION

# SELECT pgml.version(); -- 打印 PostgresML 版本信息
 version
---------
 2.7.8
```

一切就绪！更多详情请参阅 PostgresML 官方文档：https://postgresml.org/docs/guides/use-cases/

