## 用法

来源：

- [README.md](https://github.com/FranckPachot/pgwm/blob/03d43e81136fceecd210823cfc171345f591275e/README.md)
- [pgwm.control](https://github.com/FranckPachot/pgwm/blob/03d43e81136fceecd210823cfc171345f591275e/pgwm.control)
- [sql/pgwm--0.1.0.sql](https://github.com/FranckPachot/pgwm/blob/03d43e81136fceecd210823cfc171345f591275e/sql/pgwm--0.1.0.sql)
- [compose.yaml](https://github.com/FranckPachot/pgwm/blob/03d43e81136fceecd210823cfc171345f591275e/compose.yaml)

`pgwm` 0.1.0 是面向 PostgreSQL 表数据的实验性 SQL 工作区管理器，上游测试环境使用 PostgreSQL 17。应使用可丢弃的评估数据：项目明确不适用于生产数据或租户隔离。

### 基本用法

```sql
CREATE EXTENSION pgwm;
CREATE TABLE public.workspace_demo (id integer PRIMARY KEY, value text);
INSERT INTO public.workspace_demo VALUES (1, 'original');
SELECT pgwm.enable_versioning('public.workspace_demo');
SELECT pgwm.create_workspace('trial');
SELECT pgwm.goto_workspace('trial');
UPDATE public.workspace_demo SET value = 'proposal' WHERE id = 1;
SELECT pgwm.goto_workspace('LIVE');
SELECT * FROM pgwm.conflicts('trial');
```

### 版本管理与维护

`pgwm.enable_versioning` 要求表有主键，将实体表重命名为带 `_lt` 后缀的历史表，并以原名提供有类型的视图。工作区写入会形成历史行。`pgwm.refresh_workspace`、`pgwm.conflicts` 与 `pgwm.merge_workspace` 用于父子工作区的审核和合并；合并前应检查冲突。

工作区选择只影响当前会话。连接池归还连接前须重置为 `LIVE`，也不能将底层历史表当作隔离边界。主键修改、破坏性模式变更、用户触发器重放及部分清理操作均有限制。安装需要超级用户及 `plpgsql`，无需原生库或共享预加载。上游未声明许可证。
