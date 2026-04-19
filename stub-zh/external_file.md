## 用法

来源：[README](https://github.com/darold/external_file/blob/master/README.md)，[Release v1.2](https://github.com/darold/external_file/releases/tag/v1.2)

`external_file` 将文件定位符存储为 `(directory, filename)` 对，并通过 PostgreSQL `lo_*` 辅助函数访问服务器端文件，而不是直接读取文件。

### 基本流程

```sql
CREATE EXTENSION external_file;

INSERT INTO directories(directory_name, directory_path)
VALUES ('temporary', '/tmp/');

INSERT INTO directory_roles(directory_name, directory_role, directory_read, directory_write)
VALUES ('temporary', 'app_reader', true, false);

SELECT writeEfile('\x48656c6c6f0a', ('temporary', 'hello.txt'));
SELECT readEfile(efilename('temporary', 'hello.txt'));
SELECT copyEfile(('temporary', 'hello.txt'), ('temporary', 'hello-copy.txt'));
```

### 核心对象

- `directories`：将别名映射到服务器上的目录路径。
- `directory_roles`：为角色授予该别名的读写权限。
- `efilename(directory, filename)`：构造一个 `efile` 定位符。
- `readEfile(efile)`：将目标文件读取为 `bytea`。
- `writeEfile(bytea, efile)`：将 `bytea` 写入目标文件。
- `copyEfile(src, dest)`：将一个外部文件复制到另一个外部文件。
- `getEfilePath(efile, need_read, need_write)`：解析完整路径并检查访问权限。

### 注意事项

- 创建扩展需要 PostgreSQL superuser。
- 上游默认会在 `external_file` schema 中创建全部对象。
- PostgreSQL 的 OS user 仍然需要对目标目录具备文件系统读写权限。
- 文件名不能包含 `/` 或 `\`；访问必须通过目录表进行受控中介。
