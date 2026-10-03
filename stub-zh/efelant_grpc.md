## 用法

来源：

- [database/extensions/efelant_grpc/efelant_grpc.control](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/extensions/efelant_grpc/efelant_grpc.control)
- [docs/core.md](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/docs/core.md)
- [docs/api.md](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/docs/api.md)
- [database/extensions/efelant_grpc/efelant_grpc.c](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/extensions/efelant_grpc/efelant_grpc.c)
- [database/extensions/efelant_grpc/efelant_grpc--1.0.sql](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/extensions/efelant_grpc/efelant_grpc--1.0.sql)
- [database/migrations/015_api.sql](https://github.com/tafaust/efelant/blob/36457aad0cbe81fb07c46623efdf5b7203a3433f/database/migrations/015_api.sql)

`efelant_grpc` 是 Efelant 的 Connect/gRPC-JSON 传输工作进程，将请求交给应用数据库中的 `api.handle_grpc`；扩展本身不安装业务 API 或其授权模型。

### 核心用法

```conf
shared_preload_libraries = 'efelant_grpc'
efelant_grpc.database = 'efelant'
efelant_grpc.listen_addresses = '127.0.0.1'
efelant_grpc.port = 18081
```

```sql
CREATE EXTENSION efelant_grpc;
```

### 运行边界

启用监听器前，应部署匹配的 PostgreSQL 17 Efelant 模式与迁移。需要超级用户安装、预加载并重启。`efelant_grpc.database` 默认指向应用数据库，`efelant_grpc.port` 默认 18081，`efelant_grpc.listen_addresses` 默认监听所有地址。应限制绑定地址和网络访问，并使用应用部署提供的 TLS 代理。身份认证与授权由 SQL 处理函数负责。仅安装控制文件与扩展 SQL 不能得到独立的业务 API。
