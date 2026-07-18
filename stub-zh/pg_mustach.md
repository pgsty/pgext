## 用法

来源：

- [上游 README](https://github.com/RekGRpth/pg_mustach/blob/fc36a2203d2007fd93a4d1d6a49e8c24f02fd8e2/README.md)
- [扩展 control 文件](https://github.com/RekGRpth/pg_mustach/blob/fc36a2203d2007fd93a4d1d6a49e8c24f02fd8e2/pg_mustach.control)
- [SQL API](https://github.com/RekGRpth/pg_mustach/blob/fc36a2203d2007fd93a4d1d6a49e8c24f02fd8e2/pg_mustach--1.0.sql)
- [C 实现](https://github.com/RekGRpth/pg_mustach/blob/fc36a2203d2007fd93a4d1d6a49e8c24f02fd8e2/pg_mustach.c)
- [发行元数据](https://github.com/RekGRpth/pg_mustach/blob/fc36a2203d2007fd93a4d1d6a49e8c24f02fd8e2/META.json)

`pg_mustach` 通过 C `mustach` 库，用 PostgreSQL `json` 值渲染 Mustache 模板。简短的双参数形式会返回渲染后的文本：

### 渲染模板

```sql
CREATE EXTENSION pg_mustach;

SELECT mustach(
  '{"name":"PostgreSQL"}'::json,
  'Hello {{name}}!'
);
```

扩展还提供 `mustach_cjson()`、`mustach_jansson()` 和 `mustach_json_c()` 适配器。`mustach_with_*()` 函数会修改当前会话进程内的标志；`mustach_with_noextensions()` 会重置标志，其他开关则启用各个 Mustach 扩展。这些标志变更不受事务边界约束。

每个渲染器都有一个三参数重载，它接受服务端文件名，并以 PostgreSQL 操作系统账号打开文件写入。安装 SQL 没有撤销默认 `EXECUTE`，因此在允许非管理员使用扩展前，必须撤销或严格授予所有文件名重载的权限。JSON 或模板不可信时，还应限制模板输出与错误处理的资源消耗。

目录/control 对象版本为 `1.0`，而同一提交的 `META.json` 报告发行版本 `1.0.16`。打包或升级时应区分这两个版本通道。
