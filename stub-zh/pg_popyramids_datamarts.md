## 用法

来源：

- [上游 README](https://github.com/torrespri/pg_popyramids_datamarts/blob/24a3360122d8c3ba9503b36cc41b50d48a248552/README.md)
- [上游文档](https://github.com/torrespri/pg_popyramids_datamarts/blob/24a3360122d8c3ba9503b36cc41b50d48a248552/doc/pg_popyramids_datamarts.md)
- [扩展 control 文件](https://github.com/torrespri/pg_popyramids_datamarts/blob/24a3360122d8c3ba9503b36cc41b50d48a248552/pg_popyramids_datamarts.control)
- [扩展 SQL](https://github.com/torrespri/pg_popyramids_datamarts/blob/24a3360122d8c3ba9503b36cc41b50d48a248552/sql/pg_popyramids_datamarts--1.0.0.sql)

`pg_popyramids_datamarts` 是 popyramids 人口金字塔应用的数据集市层，并非独立的通用扩展。版本 `1.0.0` 依赖 PostGIS，并假定应用数据库中已经存在 `dms` 与 `ods` 模式、`ods.pyrint` 类型以及 `ods.main` 源表。

创建前先检查应用前置条件；只有匹配的 popyramids 数据库才应继续执行启用和查询步骤：

```sql
SELECT
  to_regnamespace('dms') AS dms_schema,
  to_regnamespace('ods') AS ods_schema,
  to_regtype('ods.pyrint') AS ods_pyrint,
  to_regclass('ods.main') AS ods_main;

CREATE EXTENSION postgis;
CREATE EXTENSION pg_popyramids_datamarts;

SELECT dms.chibo_give_me_pyramids(ARRAY[1, 2]);
```

### 部署约束

安装脚本会创建应用专用的类型、函数、类型转换和物化视图，并直接读取 `ods.main`。其中还包含针对角色 `postgres` 与 `chichinabo` 的所有权及授权语句。如果预期的模式、类型、表或角色缺失，扩展会在普通数据库中创建失败；部署前必须审查并修正这些环境假设。
