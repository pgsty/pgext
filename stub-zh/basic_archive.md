

## 用法

> [basic_archive: 归档模块示例](https://www.postgresql.org/docs/current/basic-archive.html)

`basic_archive` 模块是一个 WAL 归档模块，将已完成的 WAL 段文件复制到指定目录。它作为自定义归档模块的参考实现。

### 配置

添加到 `postgresql.conf`：

```ini
archive_mode = 'on'
archive_library = 'basic_archive'
basic_archive.archive_directory = '/path/to/archive/directory'
```

### 参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `basic_archive.archive_directory` | 字符串 | 复制 WAL 文件的目标目录（必须已存在） |

如果启用了 `archive_mode` 但 `basic_archive.archive_directory` 为空（默认值），服务器将持续累积 WAL 文件，直到配置目录路径。

### 注意事项

- 目标目录必须在使用前创建；模块不会自动创建
- 服务器崩溃后，归档目录中可能会留下前缀为 `archtemp` 的临时文件，应在重启前删除
- 这些临时文件也可以在服务器运行时安全删除，前提是它们与正在进行的归档操作无关
- 此模块主要用作开发自定义归档模块的简单示例和起点
