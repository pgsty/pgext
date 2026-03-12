


## 用法

> [omni_vfs: 虚拟文件系统](https://docs.omnigres.org/omni_vfs/reference/)

`omni_vfs` 扩展为 PostgreSQL 提供统一的虚拟文件系统 API，支持多种后端。

### 核心函数

**`omni_vfs.list(fs, path)`** -- 列出目录内容，返回 `(name, kind)` 集合。

**`omni_vfs.list_recursively(fs, path, max)`** -- 递归列出。使用 `max` 限制结果数量。

**`omni_vfs.file_info(fs, path)`** -- 返回文件元数据：`size`、`created_at`、`accessed_at`、`modified_at`、`kind`。

**`omni_vfs.read(fs, path, file_offset, chunk_size)`** -- 读取文件内容为 `bytea`。

**`omni_vfs.write(fs, path, content, create_file, append)`** -- 写入内容，返回写入字节数。

### 后端

**本地文件系统：**

```sql
SELECT omni_vfs.local_fs('/mount/path');
-- 不允许访问该目录之外的内容
```

**数据库后端（table_fs）：**

```sql
SELECT omni_vfs.table_fs('filesystem_name');
```

**远程文件系统：**

```sql
SELECT omni_vfs.remote_fs('dbname=otherdb host=127.0.0.1', $$omni_vfs.local_fs('/path')$$);
```

### 示例

```sql
WITH fs AS (SELECT omni_vfs.local_fs('/data') AS fs),
     entries AS (SELECT fs, (omni_vfs.list(fs, '')).* FROM fs)
SELECT entries.name, entries.kind,
       (omni_vfs.file_info(fs, '/' || entries.name)).*
FROM entries;
```
