## 用法

来源：

- [README.md](https://github.com/postgrespro/pg_uprobe/blob/336336c087070c8cc35a697216a6a51bd34905b4/README.md)
- [pg_uprobe--0.3.sql](https://github.com/postgrespro/pg_uprobe/blob/336336c087070c8cc35a697216a6a51bd34905b4/pg_uprobe--0.3.sql)
- [doc/trace_session.md](https://github.com/postgrespro/pg_uprobe/blob/336336c087070c8cc35a697216a6a51bd34905b4/doc/trace_session.md)
- [pg_uprobe.control](https://github.com/postgrespro/pg_uprobe/blob/336336c087070c8cc35a697216a6a51bd34905b4/pg_uprobe.control)

`pg_uprobe` 捕获查询文本、计划、逐节点执行、等待与锁，并通过 Frida Gum 分析原生函数。上游列明 PostgreSQL 15–17、Linux 和 FreeBSD；ARM64 尚未测试。

### 核心用法

```ini
shared_preload_libraries = 'pg_uprobe'
```

```sql
CREATE EXTENSION pg_uprobe;
SELECT start_session_trace();
SELECT 1;
SELECT stop_session_trace();
```

### 运行边界

先预加载库并重启，再以超级用户创建扩展。`start_session_trace()` 与 `stop_session_trace()` 界定跟踪范围。`set_uprobe()` 安装 TIME、HIST 或 MEM 探针，`list_uprobes()` 和 `delete_uprobe()` 管理探针。原生目标必须是已导出、非静态、非内联的函数。`pg_uprobe.data_dir` 默认为服务器数据目录下的子目录，`pg_uprobe.trace_file_limit` 限制跟踪文件大小。文件包含查询文本与计划，应限制访问。跟踪增加开销，适合有针对性的排查。
