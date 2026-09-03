## 用法

来源：

- [官方文档](https://github.com/bonesmoses/noddl/blob/7c9f1c6558bb8379ce176ce599542d123129f841/README.md)
- [扩展控制文件](https://github.com/bonesmoses/noddl/blob/7c9f1c6558bb8379ce176ce599542d123129f841/noddl.control)
- [官方仓库](https://github.com/bonesmoses/noddl)

`noddl` 支持角色与数据库例外规则的集群级 DDL 阻断器。

### 启用

将 `noddl` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `noddl`：

```ini
shared_preload_libraries = 'noddl'
```

```sql
CREATE EXTENSION noddl;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
ALTER SYSTEM SET noddl.enable TO true;
SELECT pg_reload_conf();
```

### 主要对象

官方来源通过动态方式或供应商工具定义扩展接口；授权前应检查实际安装版本。

### 运维与边界

- 预加载 `noddl` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
