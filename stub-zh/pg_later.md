

## 用法

> [pg_later: 立即执行 SQL，稍后获取结果](https://github.com/tembo-io/pg_later)

一个用于异步执行查询的 PostgreSQL 扩展。基于 [pgmq](https://github.com/pgmq/pgmq) 构建。

### 快速开始

初始化扩展后端：

```sql
CREATE EXTENSION pg_later CASCADE;
SELECT pglater.init();
```

立即执行一条 SQL 查询：

```sql
SELECT pglater.exec(
  'SELECT * FROM pg_available_extensions ORDER BY name LIMIT 2'
) AS job_id;
```

```text
 job_id
--------
     1
```

稍后通过任务 ID 获取结果：

```sql
SELECT pglater.fetch_results(1);
```

```json
{
  "query": "select * from pg_available_extensions order by name limit 2",
  "job_id": 1,
  "result": [
    {
      "name": "adminpack",
      "comment": "administrative functions for PostgreSQL",
      "default_version": "2.1",
      "installed_version": null
    },
    {
      "name": "amcheck",
      "comment": "functions for verifying relation integrity",
      "default_version": "1.3",
      "installed_version": null
    }
  ],
  "status": "success"
}
```
