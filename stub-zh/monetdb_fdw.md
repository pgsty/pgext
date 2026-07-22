## 用法

来源：

- [官方扩展 SQL](https://github.com/snaga/monetdb_fdw/blob/fbff386be673bb6f1d43cddbceb0081a20e9b459/monetdb_fdw--0.0.sql)
- [官方实现](https://github.com/snaga/monetdb_fdw/blob/fbff386be673bb6f1d43cddbceb0081a20e9b459/monetdb_fdw.c)
- [官方上游 README](https://github.com/snaga/monetdb_fdw/blob/fbff386be673bb6f1d43cddbceb0081a20e9b459/README.md)

`monetdb_fdw` 是通过 MonetDB MAPI 客户端库读取 MonetDB 数据的 0.0 版原型。它只实现了外部扫描：没有 DML、连接或聚合下推、远端代价估算，也没有生产级选项和凭据模型。

### 核心流程

连接凭据和远端关系或查询直接配置在每个外部表上：

```sql
CREATE EXTENSION monetdb_fdw;
CREATE SERVER monet_remote FOREIGN DATA WRAPPER monetdb_fdw;

CREATE FOREIGN TABLE public.remote_lineitem (
  orderkey bigint,
  quantity numeric
)
SERVER monet_remote
OPTIONS (
  host 'monetdb.example.com',
  port '50000',
  user 'monet_user',
  passwd 'secret',
  dbname 'demo',
  table 'sys.lineitem'
);

SELECT * FROM public.remote_lineitem;
```

只能使用 `table` 或 `query` 之一。设置 `table` 时，包装器构造简单的 `SELECT * FROM ...`；设置 `query` 时，则原样执行所给文本。

### 选项与构建边界

外部表选项包括 `host`、`port`、`user`、`passwd`、`dbname`、`table` 和 `query`。项目必须针对 MonetDB MAPI 头文件与 `libmapi` 编译。它没有用户映射凭据层；结果值以字符串到达，再由 PostgreSQL 转换成在本地声明的列类型。

### 安全与限制

验证器只允许超级用户修改这些外部表选项，但凭据仍保存在外部表目录元数据中，具备足够权限的角色可以看到。应限制目录和 DDL 访问，且不要把生产密钥放入这个原型。表名路径只进行最少的字符串构造，查询路径则是任意远端 SQL；绝不能传入不可信标识符或文本。应在隔离环境中测试准确的 PostgreSQL ABI 与 MonetDB MAPI 版本、类型与空值转换、长标识符和查询、编码、取消、部分读取以及连接中断。生产负载应选择成熟替代方案。
