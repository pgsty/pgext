## 用法

来源：

- [官方扩展 SQL](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/sql/postfga--1.0.0.sql)
- [官方预加载与初始化源码](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/src/init.c)
- [官方运行时配置源码](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/src/guc.c)
- [官方测试指南](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/TESTING.md)

`postfga` 版本 `1.0.0` 是一个只能预加载的预览扩展，通过 gRPC 将 PostgreSQL 中的细粒度授权检查与元组变更发送到 OpenFGA 服务器。其后台工作进程、共享内存队列、网络调用与外部副作用使它成为基础设施组件，而不是事务内局部 SQL 辅助函数。

### 前置条件与配置

将库和源码确认的 `fga.*` 设置加入 PostgreSQL 配置，然后重启服务器：

```conf
shared_preload_libraries = 'postfga'
fga.endpoint = 'dns:///openfga:8081'
fga.store_id = '01HXXXXXXXXXXXXXXX'
fga.model_id = ''
fga.cache_enabled = off
```

```sql
CREATE EXTENSION postfga;
```

当前 C 源码定义了 `fga.endpoint`、`fga.store_id`、`fga.model_id`、`fga.cache_enabled`、`fga.cache_size` 与 `fga.cache_ttl_ms`。测试指南部分位置仍使用过时的 `postfga.endpoint` 名称；应以复核源码实际注册的名称为准。

### 核心流程

检查关系，然后通过扩展 API 添加或删除元组：

```sql
SELECT fga_check(
  'document', 'budget',
  'user', 'alice',
  'reader'
);

SELECT fga_write_tuple(
  'document', 'budget',
  'user', 'alice',
  'reader'
);

SELECT fga_delete_tuple(
  'document', 'budget',
  'user', 'alice',
  'reader'
);
```

### 重要对象

- `fga_check(...)` 返回 OpenFGA 是否允许该关系。
- `fga_write_tuple(...)` 与 `fga_delete_tuple(...)` 修改远端存储中的元组。
- `fga_create_store(name)` 与 `fga_delete_store(id)` 管理 OpenFGA 存储。
- `fga_stats()` 报告扩展队列/缓存统计。
- 扩展注册了 `fga_fdw`，但复核的 SQL 没有公开已文档化的外部表流程。

### 运维注意事项

`postfga` 拒绝在 `shared_preload_libraries` 之外加载，因此启用或移除它都需要重启。OpenFGA 元组写入是外部网络副作用，外围 PostgreSQL 事务回滚时不会撤销。应限制变更与存储管理函数的 EXECUTE 权限，保护 gRPC 通道和端点，并测试超时、工作进程故障、队列耗尽、OpenFGA 中断与缓存过期时的拒绝行为。源码与测试指南存在命名和路线图不一致，包括测试专用函数，以及被称为未来工作的 FDW 流程。应将该构建视为预览代码，固定提交，并在任何敏感用途前验证授权能否安全地失败关闭。
