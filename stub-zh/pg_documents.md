## 用法

来源：

- [官方上游 README](https://gitlab.com/pierre_forstmann/pg_documents/-/blob/main/README.md)
- [官方扩展控制文件](https://gitlab.com/pierre_forstmann/pg_documents/-/blob/main/pg_documents.control)
- [官方项目页面](https://gitlab.com/pierre_forstmann/pg_documents)

`pg_documents` 将一个 HTTP 服务器嵌入 PostgreSQL，并将配置数据库中的表以 JSON 格式暴露出来。它是一个仅预加载的服务器组件，而不是一个经过确认的 `CREATE EXTENSION` 工作流程的 SQL 扩展。

### 核心工作流

在构建并安装库及其 `json-c` 依赖项后，配置目标数据库和监听端口：

```ini
shared_preload_libraries = 'pg_documents'
pg_documents.database = 'documents'
pg_documents.port = '8000'
```

重启 PostgreSQL，在配置的数据库中创建测试表，并查询该 daemon：

```sh
curl http://127.0.0.1:8000/_all_dbs
```

上游示例返回一个包含可见表名的 JSON 数组。

### 重要设置

- `pg_documents.database` 选择由工作进程打开的单个数据库。
- `pg_documents.port` 选择 HTTP 监听端口。

### 要求与注意事项

- 审查过的控制、注册或目录证据标识版本 `1.0.0`。
- 控制文件将扩展标记为不可重定位。
- README 文档记录了 PostgreSQL 15.1 和一个本地修改的 `json-c` 0.16 构建。它没有确定与其他组合的兼容性。
- 更改预加载配置需要重启服务器。
- 审查过的 README 没有记录 HTTP 端点的认证、授权或 TLS。在将其暴露到隔离测试主机之外之前，将其绑定到受保护的接口并验证其安全行为。
- 在依赖该 daemon 之前，请先测试错误请求、大量结果、事务行为、数据库重启和工作进程失败。
