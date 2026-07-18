## 用法

来源：

- [扩展 control 文件](https://github.com/danolivo/gtables_fdw/blob/c7b3727b574e2c9d59e4c6d20d590740bf0a25c4/gtables_fdw.control)
- [0.1 版安装 SQL](https://github.com/danolivo/gtables_fdw/blob/c7b3727b574e2c9d59e4c6d20d590740bf0a25c4/gtables_fdw--0.1.sql)
- [原生模块](https://github.com/danolivo/gtables_fdw/blob/c7b3727b574e2c9d59e4c6d20d590740bf0a25c4/gtables_fdw.c)
- [唯一的集成测试](https://github.com/danolivo/gtables_fdw/blob/c7b3727b574e2c9d59e4c6d20d590740bf0a25c4/t/001_basic.pl)

`gtables_fdw` 0.1 是未完成的全局表原型，并非可用的外部数据包装器。其原生模块只有空的 `_PG_init()` 函数。安装 SQL 仅创建元数据表 `gt.GlobalTables` 和 `gt.GlobalTableSpace`；设想中的函数和触发器均被注释。

### 检查但不要假定功能存在

如果扩展已经存在于一次性测试数据库，可用以下查询检查其全部有效 SQL 表面：

```sql
SELECT extversion
FROM pg_extension
WHERE extname = 'gtables_fdw';

SELECT * FROM gt.GlobalTables;
SELECT * FROM gt.GlobalTableSpace;
```

向这些表添加行不会创建、同步、路由或复制全局表。仓库没有可工作的管理 API、查询规划器、写入路径、故障处理或一致性协议。

### 破坏性安装警告

0.1 版安装脚本先执行 `DROP SCHEMA IF EXISTS gt CASCADE`，然后重新创建 `gt`。因此安装可能删除无关的已有 `gt` 模式及其所有依赖对象。若数据库中该模式名下存在有价值对象，绝不能执行 `CREATE EXTENSION gtables_fdw`。需要研究源码时，应先审查脚本，并仅使用隔离的一次性数据库。

唯一测试会预加载 `postgres_fdw, gtables_fdw`，但该库没有执行任何预加载工作，测试也不能证明表同步。最后一次源码变更发生在 2020 年，且没有兼容矩阵或 release。实际全局表需求应使用受维护的逻辑复制、外部表或分布式 PostgreSQL 设计。
