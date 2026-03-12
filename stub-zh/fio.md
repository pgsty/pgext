

## 用法

> [fio: PostgreSQL 文件 I/O 函数](https://github.com/csimsek/pgsql-fio)

`fio` 扩展提供可从 SQL 访问的文件系统 I/O 函数，允许直接在 PostgreSQL 中读取、写入和管理文件与目录。

### 文件操作

```sql
-- 读取文件内容（返回 bytea）
SELECT fio_readfile('/etc/hostname');

-- 将内容写入文件
SELECT fio_writefile('/tmp/output.txt', 'Hello World'::bytea);

-- 自动创建目录并覆盖写入
SELECT fio_writefile('/tmp/newdir/output.txt', 'data'::bytea, true, true);

-- 删除文件
SELECT fio_removefile('/tmp/output.txt');

-- 重命名/移动文件
SELECT fio_renamefile('/tmp/old.txt', '/tmp/new.txt');
```

### 目录操作

```sql
-- 列出目录内容
SELECT fio_readdir('/usr/', '*');

-- 使用模式过滤列出
SELECT fio_readdir('/var/log/', '*.log');

-- 创建目录并设置权限
SELECT fio_mkdir('/tmp/mydir', '0755');

-- 递归创建嵌套目录
SELECT fio_mkdir('/tmp/a/b/c', '0755', true);

-- 更改文件/目录权限
SELECT fio_chmod('/tmp/mydir', '0700');
```

### 函数参考

| 函数 | 描述 |
|----------|-------------|
| `fio_readfile(path)` | 以 bytea 形式读取文件内容 |
| `fio_writefile(path, content, mkdir, overwrite)` | 将 bytea 内容写入文件 |
| `fio_removefile(path)` | 删除文件 |
| `fio_renamefile(old, new)` | 重命名或移动文件 |
| `fio_readdir(path, pattern)` | 列出匹配模式的目录条目 |
| `fio_mkdir(path, mode, recursive)` | 创建目录并设置权限 |
| `fio_chmod(path, mode)` | 更改文件/目录权限 |
