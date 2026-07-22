## 用法

来源：

- [官方 ddl_versioning README](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/ddl_versioning/README.ddl_versioning)
- [1.0 版扩展 SQL](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/ddl_versioning/ddl_versioning--1.0.sql)
- [扩展控制文件](https://github.com/filiprem/pg-tools/blob/3ad4b21874bed6b629f984d046650de354e95d92/ddl_versioning/ddl_versioning.control)

`ddl_versioning` 会在受支持的 DDL 命令完成后，连续记录所选数据库对象的文本定义。它适合用作表、索引、函数和视图的轻量审计辅助工具，但不是完整的模式备份或回滚系统。

### 核心流程

该扩展是 trusted 扩展，固定安装到 `public` 模式且不可重定位：

```sql
CREATE EXTENSION ddl_versioning;

CREATE TABLE public.accounts (
    id bigint PRIMARY KEY,
    balance numeric NOT NULL
);

ALTER TABLE public.accounts ADD COLUMN updated_at timestamptz;
```

数据库级 `ddl_versioning_trigger` 事件触发器在 `ddl_command_end` 阶段运行。可按下面方式查看对象目录及其历史版本：

```sql
SELECT object_id, object_type, object_identity
FROM public.ddl_versioning_object
ORDER BY object_id;

SELECT o.object_identity, v.version_id, v.object_definition,
       v.created_at, v.created_by
FROM public.ddl_versioning_object AS o
JOIN public.ddl_versioning_version AS v USING (object_id)
ORDER BY o.object_identity, v.version_id;
```

### 记录的对象

- 表定义由 `ddl_versioning_get_tabledef` 重建；索引（包括约束产生或单独创建的索引）会另行记录。
- 索引使用 `pg_get_indexdef`，函数使用 `pg_get_functiondef`，视图使用 `pg_get_viewdef`。
- 其他 DDL 标签也可能进入事件触发器，但不受支持的对象类型只会写日志，不会生成可恢复定义。

每条受支持的 DDL 命令都会追加一行。扩展没有保留策略、差异界面、重放过程或自动回滚；如有需要，应另行实现历史清理与恢复工具。表重建结果并不能把所有模式属性都捕获成独立迁移脚本，用于恢复前必须测试生成的定义。全局事件触发器还会给 DDL 操作增加写入和日志开销，应在目标 PostgreSQL 版本上评估性能与权限。
