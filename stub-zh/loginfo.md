## 用法

来源：

- [Official extension control file](https://github.com/sroeschus/loginfo/blob/e78fb0312856ff0c4fc4000fa49f75be24ad3d67/loginfo.control)
- [Official upstream documentation](https://github.com/sroeschus/loginfo/blob/e78fb0312856ff0c4fc4000fa49f75be24ad3d67/README.md)

`loginfo` 1.0 通过 SQL 视图暴露 PostgreSQL 日志配置、日志文件清单与日志内容。它用于数据库内诊断和监控，但也可能暴露敏感语句与服务器路径。

### 核心流程

```sql
CREATE EXTENSION loginfo;

SELECT * FROM log_info;
SELECT * FROM ls_log_files;
SELECT * FROM cat_log_file;
```

`log_info` 报告当前日志相关参数和日志目标。`ls_log_files` 列出 `log_directory` 下的 `.log` 文件，`cat_log_file` 读取当前文本日志。启用 CSV 日志后，`ls_csv_files` 列出 CSV 日志，`cat_csv_file` 把其中 23 个字段作为文本列暴露。

### 配置与注意事项

要使用 CSV 视图，应在 `log_destination` 中加入 `csvlog`；上游 README 说明此变更需要重启。文件清单视图按后缀选择文件，不会展开 `log_filename` 模式。日志扫描可能很大，并可能暴露凭据、个人数据、查询文本、文件系统路径或错误详情，因此只应授权给可信运维人员，并施加严格过滤与保留策略。
