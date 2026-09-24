## 用法

来源：

- [官方版本 v0.1.1](https://github.com/pgsty/pgs3/releases/tag/v0.1.1)
- [官方 README v0.1.1](https://github.com/pgsty/pgs3/blob/v0.1.1/README.md)
- [扩展控制文件](https://github.com/pgsty/pgs3/blob/v0.1.1/pgs3.control)
- [配置参数参考](https://github.com/pgsty/pgs3/blob/v0.1.1/docs/guc.md)
- [运维指南](https://github.com/pgsty/pgs3/blob/v0.1.1/docs/operations.md)
- [已知限制](https://github.com/pgsty/pgs3/blob/v0.1.1/docs/known-limitations.md)

`pgs3` 0.1.1 将一个 PostgreSQL 数据库变成仅支持 path-style 的 S3 兼容端点。PostgreSQL 后台 worker 验证 SigV4 请求，并在普通 SQL 表上执行带版本的对象操作，因此对象 metadata、payload、授权、WAL、物理备份与恢复都留在 PostgreSQL 内部。它面向 PostgreSQL 17 与 18，目前仍属于早期 alpha 软件。

### 核心流程

先在承载对象存储的数据库中安装扩展，创建受限租户角色与凭据，再启动 worker 池：

```sql
CREATE EXTENSION pgs3;

CREATE ROLE tenant_app
  NOLOGIN NOINHERIT NOSUPERUSER NOCREATEDB NOCREATEROLE
  NOREPLICATION NOBYPASSRLS;

SELECT pgs3.create_credential(
  'tenant-access-key', 'replace-with-a-secret', 'tenant_app'::name, true
);

SELECT pgs3.start();
TABLE pgs3.worker_state;
TABLE pgs3.stats;
```

如需自动启动，应先预加载动态库、配置目标数据库，再重启 PostgreSQL：

```conf
shared_preload_libraries = 'pgs3'
pgs3.enabled = on
pgs3.target_database = 'artifacts'
pgs3.listen_addr = '127.0.0.1'
pgs3.port = 9000
pgs3.workers = 4
```

增删 `shared_preload_libraries` 或修改 `pgs3.target_database` 都需要重启 PostgreSQL。其他已记录参数使用 SIGHUP 语义，但运维就绪状态必须通过 `pgs3.worker_state` 与日志判断，不能只看 TCP 端口是否开放。手工启动的 worker 池可以用 `pgs3.stop()` 停止。

### 客户端与存储行为

客户端必须使用 path-style addressing，并显式指定 endpoint：

```bash
export PGS3_ENDPOINT='https://s3.example.com'
export AWS_ACCESS_KEY_ID='<access-key>'
export AWS_SECRET_ACCESS_KEY='<secret-key>'
export AWS_DEFAULT_REGION='us-east-1'

aws --endpoint-url "$PGS3_ENDPOINT" s3api list-buckets
```

务必传入 `--endpoint-url`，否则 AWS 客户端可能把请求静默发送到 AWS。已验证的客户端路径包括 AWS CLI、boto3、rclone、s3fs 和 DuckDB `httpfs`。Bucket 与 object 操作涵盖 range 和条件读写、ListObjectsV2、永久版本历史、delete marker、CopyObject 与 multipart upload。

规范 payload 保存在 `pgs3.blob`。Object version、CopyObject、SQL Restore 与仅复制 metadata 的 Fork 操作可以共享同一 blob，而不重复复制字节。一个 endpoint 只服务一个已配置数据库。

### 运维与安全

Credential access key 映射到 PostgreSQL 租户角色。角色应保持 `NOLOGIN`、`NOINHERIT` 和 `NOBYPASSRLS`；不要把 `pgs3.server_role` 用作应用身份。Row-level security 是租户隔离边界，credential 管理与 worker 控制函数仍只应交给管理员。

SigV4 需要可逆存储 secret，因此数据库备份与副本包含 credential 材料，必须加密并限制访问。`pgs3` 只提供明文 HTTP；生产环境需要外部 TLS 代理，并保持签名 path、host 与 header 不变。对象状态应使用物理备份：扩展自有对象数据目前不支持逻辑 dump/restore。

### 兼容性与限制

当前打包和上游支持路径为 PostgreSQL 17、18 与 pgrx 0.19.2。版本 0.1.1 包含已验证的 `0.1.0 -> 0.1.1` 扩展升级边。

`pgs3` 不是通用生产级 S3 替代品。它尚未实现 virtual-host addressing、内置 TLS、IAM 或 bucket policy 语言、ACL、lifecycle rule 与跨数据库路由；小对象 GET/PUT 目标和十万对象 Fork 目标也未达成，完整对象 GET 当前还会在内存中物化响应。部署时应设置适合自身场景的对象大小上限，并在暴露 endpoint 前阅读上游限制。
