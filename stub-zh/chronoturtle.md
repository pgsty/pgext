## 用法

来源：

- [README](https://github.com/geoinfo-applications/chrono-turtle/blob/a2a7a5a9b73a3f5abf126cfd5ed88cd9142e57b6/README.md)
- [Control file / 控制文件](https://github.com/geoinfo-applications/chrono-turtle/blob/a2a7a5a9b73a3f5abf126cfd5ed88cd9142e57b6/chronoturtle.control)
- [Makefile](https://github.com/geoinfo-applications/chrono-turtle/blob/a2a7a5a9b73a3f5abf126cfd5ed88cd9142e57b6/Makefile)
- [sql/functions/chronoturtle.sql](https://github.com/geoinfo-applications/chrono-turtle/blob/a2a7a5a9b73a3f5abf126cfd5ed88cd9142e57b6/sql/functions/chronoturtle.sql)

`chronoturtle` 将普通表转换为具有分支语义的视图，底层以追加方式保存行版本。上游明确将 0.1.0 定义为已停止维护的可行性研究，仅测试了 PostgreSQL 16。

### 核心工作流

```sql
CREATE EXTENSION chronoturtle;
SET search_path TO chronoturtle, public;
CREATE TABLE public.contacts (id serial PRIMARY KEY, email text NOT NULL);
INSERT INTO public.contacts(email) VALUES ('alice@example.com');
SELECT migrate_to_vie('public.contacts');
SELECT set_commit_author('editor@example.com');
SELECT set_current_branch('feature/contact', 'main');
UPDATE public.contacts SET email = 'alice@example.org' WHERE id = 1;
SELECT set_current_branch('main');
SELECT * FROM public.contacts;
```

### 对象与维护

`migrate_to_vie` 转换指定表，`migrate_schemas_to_vie` 处理整个模式并按外键依赖排序。`set_current_branch` 选择分支，`set_commit_author` 记录后续修改的作者。公共函数位于 `chronoturtle`，元数据位于 `_vie`，`_backing` 保存全部行版本。`add_unique_constraint` 和 `drop_unique_constraint` 管理支持的视图约束。

历史不会自动清理。`drop_vie_table` 默认保留全部底层行，包括历史版本，并非只导出当前可见分支的快照；第二个参数可要求删除数据。

### 限制

这是纯 SQL 扩展，无共享库或预加载要求。应先由具有相应权限的表所有者在副本上测试转换。尚未实现合并、变基、检出和差异比较；不支持结构变更、带引号的大小写混合标识符，以及对转换后视图执行 `TRUNCATE`，也没有升级脚本。
