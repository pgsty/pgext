## 用法

来源：

- [官方文档](https://github.com/rogerwelin/pg_fk_indexer/blob/5ab97eaff8e456d58bda8d300efd64083247b534/README.md)
- [扩展控制文件](https://github.com/rogerwelin/pg_fk_indexer/blob/5ab97eaff8e456d58bda8d300efd64083247b534/pg_fk_indexer.control)
- [官方仓库](https://github.com/rogerwelin/pg_fk_indexer)

`pg_fk_indexer` 为新定义的外键列自动创建索引。

### 启用

将 `pg_fk_indexer` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_fk_indexer`：

```ini
shared_preload_libraries = 'pg_fk_indexer'
```

```sql
CREATE EXTENSION pg_fk_indexer;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Foreign keys are now auto-indexed
CREATE TABLE users (id int PRIMARY KEY, username text);
CREATE TABLE orders (user_id int REFERENCES users(id));
--  index on orders(user_id) is created automatically
```

### 主要对象

官方来源通过动态方式或供应商工具定义扩展接口；授权前应检查实际安装版本。

### 运维与边界

- 预加载 `pg_fk_indexer` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
