## 用法

来源：

- [官方资料](https://github.com/matthias-reissner/postrcs/blob/0cf6140c7665cdbf4fd9d7f14856ec349b026f9b/USAGE.md)
- [官方资料](https://github.com/matthias-reissner/postrcs/blob/0cf6140c7665cdbf4fd9d7f14856ec349b026f9b/postrcs.control)
- [官方资料](https://github.com/matthias-reissner/postrcs/blob/0cf6140c7665cdbf4fd9d7f14856ec349b026f9b/postrcs--1.2.1.sql)
- [官方资料](https://github.com/matthias-reissner/postrcs/blob/0cf6140c7665cdbf4fd9d7f14856ec349b026f9b/README.md)

postrcs 使用 rcs_archive 域存储 RCS 文本归档，不会自动跟踪其他列的变化：调用方必须保存每次返回的归档。已审核源码演示 PostgreSQL 18 Windows 环境，未确立更广的兼容范围或许可证。

### 保存与读取修订

```sql
CREATE EXTENSION postrcs;
CREATE TABLE documents (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  history rcs_archive NOT NULL DEFAULT ''::rcs_archive
);
INSERT INTO documents(history)
VALUES (rcs_checkin(''::rcs_archive, E'First line\n', 'initial'));
UPDATE documents
SET history = rcs_checkin(history, E'First line\nSecond line\n', 'append')
WHERE id = 1;
SELECT rcs_head(history), rcs_checkout(history, '1.1'),
       rcs_diff(history, '1.1', '1.2')
FROM documents WHERE id = 1;
```

rcs_revisions 列出元数据；rcs_create_branch、rcs_branch_checkin 和 rcs_tag 返回替换用的归档。rcs_merge 返回 content 与 has_conflicts，不会持久化合并。检查冲突标记并提交解决后的文本。rcs_import 验证归档文本，rcs_export 返回归档文本。

### 限制与升级

非空内容必须以换行符结尾，不支持二进制内容或嵌套分支。限制为每份归档 16 MiB、每份内容 8 MiB、每条日志 1 MiB，以及每次 LCS 差异/合并四百万行对。无需预加载。已有 1.0、1.1 或 1.2 安装应换上匹配的 1.2.1 文件，再执行 ALTER EXTENSION postrcs UPDATE。
