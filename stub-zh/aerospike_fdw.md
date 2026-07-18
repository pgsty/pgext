## 用法

来源：

- [已复核 commit 的 README](https://github.com/sanglijunfengDev/aerospike_fdw/blob/61f0e5143d73dc1d5ce15249daab98ca8e7a46ea/README)
- [扩展 control 文件](https://github.com/sanglijunfengDev/aerospike_fdw/blob/61f0e5143d73dc1d5ce15249daab98ca8e7a46ea/aerospike_fdw.control)
- [扩展安装 SQL](https://github.com/sanglijunfengDev/aerospike_fdw/blob/61f0e5143d73dc1d5ce15249daab98ca8e7a46ea/aerospike_fdw--1.0.sql)
- [FDW 实现](https://github.com/sanglijunfengDev/aerospike_fdw/blob/61f0e5143d73dc1d5ce15249daab98ca8e7a46ea/aerospike_fdw.c)

`aerospike_fdw` 是一个实验性外部数据包装器，可通过 PostgreSQL 外部表读取 Aerospike namespace 与 set 中的记录，并向其中插入记录。服务器连接是进程全局的：`aerospike.as_server_ip` 与 `aerospike.as_server_port` 都是 postmaster 启动参数，因此必须预加载该库并重启服务器。

### 配置与外部表

上游构建依赖 Aerospike C 客户端及其 Lua 支持文件。其实现还要求 Lua 文件位于硬编码的 `/home/pg/lua` 目录树下。

```conf
shared_preload_libraries = 'aerospike_fdw'
aerospike.as_server_ip = '127.0.0.1'
aerospike.as_server_port = 3000
```

重启 PostgreSQL 后，创建包装器并映射 Aerospike set。`key` 选项指定用作 Aerospike 记录键的外部表列名。

```sql
CREATE EXTENSION aerospike_fdw;

CREATE SERVER aerospike_server
  FOREIGN DATA WRAPPER aerospike_fdw;

CREATE FOREIGN TABLE aerospike_demo (
  id bigint,
  value text
)
SERVER aerospike_server
OPTIONS (namespace 'test', set 'demo', key 'id');
```

### 注意事项

- 已复核的 FDW routine 支持扫描与插入，但不支持更新或删除。其 validator 只接受外部表选项 `namespace`、`set` 和 `key`，却不检查三者是否齐全。
- 该库在 PostgreSQL 启动期间连接 Aerospike，并会在某些连接失败路径中终止进程。请在远离生产环境的地方测试启动及失败处理。
- 扫描执行会生成并上传 Aerospike Lua UDF 文件，而且使用硬编码的服务器端路径。这些路径、文件权限和 UDF 权限属于部署要求，并非普通 SQL 配置。
- 上游源码停留在 2016 年，且没有发布当前 PostgreSQL 或 Aerospike 兼容矩阵。应把版本 `1.0` 视为遗留实验代码，并在使用前针对确切的服务器与客户端版本进行验证。
