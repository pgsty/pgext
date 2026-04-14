## 用法

- 来源: [GitHub 仓库](https://github.com/darold/external_file), [README](https://github.com/darold/external_file/blob/master/README.md)
- `external_file` 提供通过 PostgreSQL 扩展访问外部文件的能力，类似 Oracle BFILE 风格的定位器。

```sql
CREATE EXTENSION external_file;
```

README 说明对象默认创建在 `external_file` 模式中，且创建扩展需要 PostgreSQL 超级用户权限。

## 核心流程

该扩展使用“目录别名 + 文件名”来标识外部文件。上游 README 展示的流程如下：

```sql
INSERT INTO directories(directory_name, directory_path)
VALUES ('temporary', '/tmp/');

INSERT INTO directory_roles(directory_name, directory_role, directory_read, directory_write)
VALUES ('temporary', 'a_role', true, false);

SELECT writeEfile('\x48656c6c6f2c0a0a596f75206172652072656164696e67206120746578742066696c652e0a0a526567617264732c0a',
                  ('temporary', 'blahblah.txt'));
SELECT readefile(the_file) FROM efile_test;
SELECT copyefile(('temporary', 'blahblah.txt'), ('temporary', 'copy_blahblah.txt'));
```

主要导出辅助函数包括 `efilename`、`readEfile`、`writeEfile`、`copyEfile` 和 `getEfilePath`。

## 备注

该扩展不会直接读取服务器文件系统上的文件。它通过服务器端 `lo_*` 系列函数访问文件，并通过目录表和角色表控制访问权限。
