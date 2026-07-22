## 用法

来源：

- [扩展控制文件](https://github.com/goldenhelix/tsf-fdw/blob/bf8d839bf447d3300cd371b9e059b612b9330923/tsf_fdw.control)
- [版本 1.0 SQL API](https://github.com/goldenhelix/tsf-fdw/blob/bf8d839bf447d3300cd371b9e059b612b9330923/tsf_fdw--1.0.sql)
- [官方外部表示例](https://github.com/goldenhelix/tsf-fdw/blob/bf8d839bf447d3300cd371b9e059b612b9330923/example.sql)
- [FDW 实现](https://github.com/goldenhelix/tsf-fdw/blob/bf8d839bf447d3300cd371b9e059b612b9330923/tsf_fdw.c)

`tsf_fdw` 版本 `1.0` 是 Golden Helix 的外部数据包装器，用于把本地 TSF 基因组数据文件中的字段作为 PostgreSQL 外部表读取。它只实现扫描，不能向 TSF 文件写入。

### 核心流程

```sql
CREATE EXTENSION tsf_fdw;
CREATE SERVER tsf_server FOREIGN DATA WRAPPER tsf_fdw;

CREATE FOREIGN TABLE low_level (
  _id integer,
  Chr text,
  Start integer,
  Stop integer,
  FloatField real,
  StringField text
) SERVER tsf_server
  OPTIONS (filename '/data/Low_level.tsf', sourceid '1');

SELECT Chr, Start, Stop FROM low_level LIMIT 20;
```

PostgreSQL 服务器进程必须能够读取该文件路径。外部表选项包括 `filename`、可选的基础 `path`、`sourceid` 和 `fieldtype`；逐列选项可以覆盖文件/来源，并选择 `fieldidx` 或映射来源。列名与类型必须和 TSF 来源精确匹配。

为了发现模式，`tsf_generate_schemas(prefix, tsf_path, source_id)` 可以返回单个或全部来源的生成 DDL。执行前应审核输出，特别是生成的服务器名与绝对路径。

### 查询与运行边界

该 FDW 会跳过未引用字段，并对 TSF 迭代器执行本地限制条件处理，但没有写入回调。文件内容、来源 ID、字段种类和映射必须始终与外部表定义一致。

该项目没有实质性的上游 README，并依赖内嵌 TSF C 库。应使用有代表性的 Golden Helix 数据验证文件格式兼容性、权限、并发访问、数组/映射字段、查询计划与损坏文件错误。需要单独备份源文件；PostgreSQL 备份不包含外部 TSF 内容。
