## 用法

来源：

- [用法与运行边界](https://github.com/roelofrossouw/simply-cpp-pglib/blob/5b437c79bf5cda145e85c336a5c156f9b6b1939a/README.md)
- [扩展控制文件](https://github.com/roelofrossouw/simply-cpp-pglib/blob/5b437c79bf5cda145e85c336a5c156f9b6b1939a/sql/sc_pglib.control)
- [初始 SQL](https://github.com/roelofrossouw/simply-cpp-pglib/blob/5b437c79bf5cda145e85c336a5c156f9b6b1939a/sql/sc_pglib--1.0.sql)
- [升级到 SQL 版本 1.18](https://github.com/roelofrossouw/simply-cpp-pglib/blob/5b437c79bf5cda145e85c336a5c156f9b6b1939a/sql/sc_pglib--1.17--1.18.sql)

`sc_pglib` 将 simply-cpp 库的功能提供给 PostgreSQL。SQL/control 版本 `1.18` 与项目的软件包版本分开管理。API 包括 Base64 转换、电子表格和 CSV 操作、服务器文件、Redis、Firebase 及特定应用集成。文档所列 Ubuntu 软件包面向 PostgreSQL 16 和 18；每份库必须与服务器主版本匹配。

### 启用与使用

安装匹配的库和 SQL 文件后，以超级用户启用：

```sql
CREATE EXTENSION sc_pglib;

SELECT sc_base64_encode('Hello World!');
SELECT convert_from(sc_base64_decode('SGVsbG8gV29ybGQh'), 'UTF8');
```

这些 Base64 函数为不可变、严格且并行安全的函数，无效编码会报错。已有数据库安装对应文件后执行 `ALTER EXTENSION sc_pglib UPDATE`；源码提供初始脚本及连续的升级脚本。扩展可迁移模式，不要求共享预加载。

### 文件与外部服务

`sc_xlsx` 将查询结果写入数据库服务器上的工作簿：

```sql
SELECT sc_xlsx('/var/www/data/output/result.xlsx',
               'SELECT 1 AS sample', 'Results');
```

文件操作使用数据库服务器的操作系统身份。普通调用者只能访问上游配置或编译时指定的目录；`pg_read_server_files` 与 `pg_write_server_files` 允许更广泛的访问。应核对覆盖文件的行为，仅授权应用所需函数。数据库事务回滚不会撤销文件及外部服务的副作用。

Redis 使用 `sc_redis` 包装器的外部服务器及可选用户映射。Firebase 使用 `sc_firebase` 外部服务器和服务账户用户映射。这些 API 执行外部操作，不能替代具有事务保证的 PostgreSQL 表。凭据应存放在受保护的用户映射或文档指定的服务器配置中，不应写入 SQL 示例或公开日志。

### 敏感操作

`sc_read_text`、`sc_write_text`、`sc_read_dsn` 和 `sc_http_get` 已撤销 `PUBLIC` 执行权限。HTTP 访问受超级用户配置的 `sc_pglib.http_urls` 允许列表限制，且不跟随重定向。`sc_oneweb` 集成使用服务器端配置文件，由 `sc_pglib.oneweb_config` 指定；其特权操作同样撤销了公开执行权限。将这些函数授权给应用角色前，应检查对应版本的上游文档。
