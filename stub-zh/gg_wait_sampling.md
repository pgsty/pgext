## 用法

来源：

- [官方文档](https://github.com/GreengageDB/greengage/blob/3253a8120ecdbafb29019740b90af5bef07790c2/gpcontrib/gg_wait_sampling/README.md)
- [扩展控制文件](https://github.com/GreengageDB/greengage/blob/3253a8120ecdbafb29019740b90af5bef07790c2/gpcontrib/gg_wait_sampling/gg_wait_sampling.control)
- [官方仓库](https://github.com/GreengageDB/greengage)

`gg_wait_sampling` 为 Greengage 提供带历史记录和查询归因的等待事件采样。

### 启用

将 `gg_wait_sampling` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `gg_wait_sampling`：

```ini
shared_preload_libraries = 'gg_wait_sampling'
```

```sql
CREATE EXTENSION gg_wait_sampling;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT * FROM gg_wait_sampling.gg_wait_sampling_reset_profile;
```

### 主要对象

官方来源通过动态方式或供应商工具定义扩展接口；授权前应检查实际安装版本。

### 运维与边界

- 预加载 `gg_wait_sampling` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
