


## 用法

来源：[PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md)，[pgh_consistency SQL](https://github.com/pghydro/pghydro/blob/master/pgh_consistency--6.6.sql)

`pgh_consistency` 是 PgHydro 的一致性检查扩展，用于检查和清理水文网络数据。它创建 `pgh_consistency` schema，提供排水线、排水区几何、拓扑以及 Pfafstetter 编码相关的质量检查表和函数。

### 启用

在 PostGIS 和核心 `pghydro` 扩展之后安装：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_consistency;
```

### 查看检查表

扩展会按问题类型创建结果表：

```sql
SELECT tablename
FROM pg_tables
WHERE schemaname = 'pgh_consistency'
ORDER BY tablename;
```

常见表包括 `pghft_drainagelineisdisconnected`、`pghft_drainagelineisnotvalid`、`pghft_drainageareaisnotvalid`、`pghft_drainageareaoverlapdrainagearea`、`pghft_pointvalencevalue2`、`pghft_pointdivergent`。

### 主要流程

在运行 PgHydro 分析函数之前，可以先用 `pgh_consistency` 中的函数准备和校验排水网络。SQL 对象包括拆分多点排水线、网格吸附、移除重复点、创建排水线顶点交点、打断排水线，以及识别无效、非简单、断连、重叠、成环或分类异常排水要素的辅助函数。

```sql
SELECT nspname, proname
FROM pg_proc p
JOIN pg_namespace n ON n.oid = p.pronamespace
WHERE nspname = 'pgh_consistency'
ORDER BY proname;
```

### 注意事项

该扩展假定 PgHydro 基础排水表已经加载完成。导出或计算更高层水文产品之前，先运行一致性检查可以更早暴露无效几何和断裂网络连接。
