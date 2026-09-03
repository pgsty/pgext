## 用法

来源：

- [官方文档](https://github.com/nminoru/pg_plan_tree_dot/blob/13eff984a59ccfeb332685e870195885db02a8e3/README.md)
- [扩展控制文件](https://github.com/nminoru/pg_plan_tree_dot/blob/13eff984a59ccfeb332685e870195885db02a8e3/pg_plan_tree_dot.control)
- [官方仓库](https://github.com/nminoru/pg_plan_tree_dot)

`pg_plan_tree_dot` 将 PostgreSQL 执行计划树渲染为 Graphviz DOT。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_plan_tree_dot`：

```sql
CREATE EXTENSION pg_plan_tree_dot;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE EXTENSION pg_plan_tree_dot;
SELECT generate_plan_tree_dot('sql', 'output.dot');
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `public.generate_plan_tree_dot` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 9.0, 9.1, 9.2, 9.3, 9.4, 9.5, 9.6, 10, 11, 12；不要推断未列出的主版本。
- 目录生命周期为 abandoned；生产使用前应测试升级、备份恢复与服务器兼容性。
