## 用法

来源：

- [Official README](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/README.md)
- [Extension control file](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_fdw/pgspider_fdw.control)
- [Installation SQL](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_fdw/pgspider_fdw--1.4.sql)
- [Build configuration](https://github.com/pgspider/pgspider/blob/d2ad2403600b3ff35fd35acf32864fabf958cb9b/contrib/pgspider_fdw/Makefile)

`pgspider_fdw` 连接 PGSpider 节点，并支持该分支通过云函数实现的可选压缩传输。应使用匹配的 PGSpider 发行版；它与 `pgspider_core_fdw`、`pgspider_ext` 是不同扩展。

### 远端节点用法

在兼容的 PGSpider 节点上，管理员可为已有远端表创建服务器与映射。需替换实际主机、数据库和凭据；PGSpider 默认端口为 4813。

```sql
CREATE EXTENSION pgspider_fdw;
CREATE SERVER remote_spider FOREIGN DATA WRAPPER pgspider_fdw
  OPTIONS (host '127.0.0.1', port '4813', dbname 'pgspider');
CREATE USER MAPPING FOR CURRENT_USER SERVER remote_spider
  OPTIONS (user 'app_user', password 'replace-me');
CREATE FOREIGN TABLE remote_t1(i integer, t text, __spd_url text)
  SERVER remote_spider OPTIONS (table_name 't1');
SELECT * FROM remote_t1;
```

### 连接与传输

`pgspider_fdw_get_connections()` 列出缓存的远端连接，`pgspider_fdw_disconnect(text)` 和 `pgspider_fdw_disconnect_all()` 释放连接。可选传输流程使用 `endpoint`、`proxy`、`batch_size` 及内核专用迁移命令，需另外部署兼容的云函数。使用前须遵循官方部署流程；上面的 SQL 示例不会配置该功能。

### 依赖与限制

模块链接 libpq、libcurl 和 LZ4。SQL/control 版本 `1.4` 与外围 PGSpider 内核版本分别维护。仅授予必要的服务器与表权限，并保护用户映射凭据。操作取决于远端可用性和 FDW 事务行为，不表示具有分布式原子提交保证。源码未提供当前原生 PostgreSQL 兼容矩阵，也未要求预加载。

### 安装对象与名称冲突

1.4 的安装脚本还创建 `pgspider_create_or_replace_stub(text, text, regtype)` 过程、`mysql_string_type`、`time_unit`、`path_value` 类型及数百个下推占位函数和聚合。这些例程需要 `plpgsql`，在本地执行时会抛出错误，并不提供对应远端函数的本地实现。

脚本使用 `CREATE OR REPLACE FUNCTION` 和 `CREATE OR REPLACE AGGREGATE`，且部分类型存在性判断仅按类型名、未限定模式。安装前须审查安装模式中的同名函数及整个数据库中的同名类型，避免与应用对象冲突。
