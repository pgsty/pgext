


## Usage

> [fio: PostgreSQL File I/O Functions](https://github.com/csimsek/pgsql-fio)

The `fio` extension provides file system I/O functions accessible from SQL, enabling reading, writing, and managing files and directories directly from PostgreSQL.

### File Operations

```sql
-- Read file contents (returns bytea)
SELECT fio_readfile('/etc/hostname');

-- Write content to file
SELECT fio_writefile('/tmp/output.txt', 'Hello World'::bytea);

-- Write with auto-create directory and overwrite
SELECT fio_writefile('/tmp/newdir/output.txt', 'data'::bytea, true, true);

-- Remove a file
SELECT fio_removefile('/tmp/output.txt');

-- Rename / move a file
SELECT fio_renamefile('/tmp/old.txt', '/tmp/new.txt');
```

### Directory Operations

```sql
-- List directory contents
SELECT fio_readdir('/usr/', '*');

-- List with pattern filter
SELECT fio_readdir('/var/log/', '*.log');

-- Create a directory with permissions
SELECT fio_mkdir('/tmp/mydir', '0755');

-- Create nested directories recursively
SELECT fio_mkdir('/tmp/a/b/c', '0755', true);

-- Change file/directory permissions
SELECT fio_chmod('/tmp/mydir', '0700');
```

### Function Reference

| Function | Description |
|----------|-------------|
| `fio_readfile(path)` | Read file contents as bytea |
| `fio_writefile(path, content, mkdir, overwrite)` | Write bytea content to file |
| `fio_removefile(path)` | Delete a file |
| `fio_renamefile(old, new)` | Rename or move a file |
| `fio_readdir(path, pattern)` | List directory entries matching pattern |
| `fio_mkdir(path, mode, recursive)` | Create directory with permissions |
| `fio_chmod(path, mode)` | Change file/directory permissions |
