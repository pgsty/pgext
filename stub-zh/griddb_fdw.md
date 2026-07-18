## 用法

来源：

- [上游 README](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/README.md)
- [扩展 control 文件](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/griddb_fdw.control)
- [扩展 SQL](https://github.com/pgspider/griddb_fdw/blob/a05fbf78a57d1941538aa11a0038037b08ff0e79/griddb_fdw--1.2.sql)
- [项目 v2.4.0 发行版](https://github.com/pgspider/griddb_fdw/releases/tag/v2.4.0)

`griddb_fdw` 用于把 PostgreSQL 连接到 GridDB 容器。构建和运行时安装依赖外部 GridDB C 客户端，其中包括 `gridstore.h` 和 `libgridstore.so`。

创建服务器和用户映射后，导入远端模式：

```sql
CREATE EXTENSION griddb_fdw;

CREATE SERVER griddb_remote
  FOREIGN DATA WRAPPER griddb_fdw
  OPTIONS (host '239.0.0.1', port '31999', clustername 'exampleCluster', database 'public');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER griddb_remote
  OPTIONS (username 'appuser', password 'change-me');

IMPORT FOREIGN SCHEMA public
  FROM SERVER griddb_remote
  INTO public;
```

### 版本与限制

项目发行版 `v2.4.0` 包含扩展 SQL 版本 `1.2`，两者属于不同版本体系。该发行版记录支持 PostgreSQL 13 至 17，并以 GridDB 5.5.0 完成测试。它支持读写，但远端容器执行 `UPDATE` 或 `DELETE` 时必须具有行键。它不支持 `ON CONFLICT`、`RETURNING`、`TRUNCATE` 或 `SAVEPOINT`，且 PostgreSQL 列定义必须与 GridDB 匹配，因此上游建议使用 `IMPORT FOREIGN SCHEMA`，而非手写外部表定义。
