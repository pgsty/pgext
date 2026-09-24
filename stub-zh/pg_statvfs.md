## 用法

来源：

- [Official documentation](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/README)
- [Control file](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/pg_statvfs.control)
- [Version 1.0 SQL](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/pg_statvfs--1.0.sql)
- [Privilege and path checks](https://github.com/michaelpq/pg_plugins/blob/ae57c1f3df697fb942f4576f873a655187193ace/pg_statvfs/pg_statvfs.c)

`pg_statvfs` 1.0 通过服务器的 statvfs 调用提供文件系统容量、inode 可用量和挂载标志。查询对象是服务器路径所在的文件系统，不是客户端目录，也不是单个关系的大小。

### 核心流程

安装扩展文件后，使用超级用户会话：

```sql
CREATE EXTENSION pg_statvfs;
SELECT * FROM pg_statvfs(current_setting('data_directory'));
SELECT f_frsize * f_blocks AS total_bytes,
       f_frsize * f_bavail AS available_bytes
FROM pg_statvfs(current_setting('data_directory'));
```

### 返回结果与权限

`pg_statvfs(path text)` 返回一条记录。`f_bsize` 是首选块大小，`f_frsize` 是块计数的单位；`f_blocks`、`f_bfree` 和 `f_bavail` 分别表示总块数、空闲块数及非特权用户可用块数。`f_files`、`f_ffree` 和 `f_favail` 描述 inode 数量；`f_fsid`、`f_namemax` 和 `flags` 提供文件系统标识、名称长度限制及支持的挂载标志。

C 函数显式要求超级用户权限，并在请求操作系统前检查路径。路径不存在或不可访问时会报错。扩展自身不需要预加载或重启；控制文件允许迁移模式，但未声明为受信任扩展。上游子目录没有给出完整的 PostgreSQL 主版本兼容矩阵，操作系统支持和可返回的挂载标志也存在差异。
