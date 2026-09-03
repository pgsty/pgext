## 用法

来源：

- [官方文档](https://github.com/styk-tv/pgCK/blob/49638f0a0c9ef48c051cc7c0ccc2e48d105d29a1/README.md)
- [扩展控制文件](https://github.com/styk-tv/pgCK/blob/49638f0a0c9ef48c051cc7c0ccc2e48d105d29a1/pgck.control)
- [构建清单](https://github.com/styk-tv/pgCK/blob/49638f0a0c9ef48c051cc7c0ccc2e48d105d29a1/Cargo.toml)

`pgck` 集成 NATS 传输、SHACL 校验与 RDF 物化的概念内核运行时。

### 启用

将 `pgck` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pgck`：

```ini
shared_preload_libraries = 'pgck'
```

```sql
CREATE EXTENSION pgck CASCADE;
```

经审查的控制文件或官方流程要求 `pgrdf`, `pgcrypto`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT ckp.dispatch(
  'instance.create',
  '{"type":"urn:ckp:demo/type/Ship","name":"Aurora"}'::jsonb
);
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `ckp.dispatch` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 18；不要推断未列出的主版本。
- 预加载 `pgck` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
