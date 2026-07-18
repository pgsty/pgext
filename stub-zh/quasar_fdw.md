## 用法

来源：

- [上游 README 与选项参考](https://github.com/slamdata/quasar-fdw/blob/23ece25661b2071644179e079134f4041e099026/README.md)
- [扩展 control 文件](https://github.com/slamdata/quasar-fdw/blob/23ece25661b2071644179e079134f4041e099026/quasar_fdw.control)
- [选项校验器](https://github.com/slamdata/quasar-fdw/blob/23ece25661b2071644179e079134f4041e099026/src/quasar_options.c)
- [已归档的上游仓库](https://github.com/slamdata/quasar-fdw)

`quasar_fdw` 是面向历史 SlamData Quasar 查询引擎的只读外部数据包装器。它会转发 `SELECT` 查询，并可把受支持的 `WHERE`、`ORDER BY` 和连接子句下推到 Quasar。

### 定义遗留 Quasar 表

```sql
CREATE EXTENSION quasar_fdw;

CREATE SERVER quasar
  FOREIGN DATA WRAPPER quasar_fdw
  OPTIONS (
    server 'http://localhost:8080',
    path '/local/quasar',
    timeout_ms '1000',
    use_remote_estimate 'true'
  );

CREATE FOREIGN TABLE public.zips (
  city text,
  population integer OPTIONS (map 'pop'),
  state char(2)
)
SERVER quasar
OPTIONS (table 'zips');

SELECT city, population
FROM public.zips
WHERE population > 100000
ORDER BY population DESC;
```

服务端选项还包括 `fdw_startup_cost` 和 `fdw_tuple_cost`。表选项为 `table` 与 `use_remote_estimate`；列选项为 `map`、`nopushdown` 和 `join_rowcount_estimate`。远端字符串在本地转换为其他类型、且谓词下推会改变语义时，应使用 `nopushdown`。

上游仓库已归档，README 仅支持 PostgreSQL 9.4/9.5 与 Quasar 9 至 13。control/目录版本是 `1.4.1`，而 README 仍称 `1.4.0` 为最新版本。上游没有记录写操作或当前兼容路径。该扩展只应用于维持经过验证的遗留部署；新系统应选择仍在维护的数据源和 FDW。
