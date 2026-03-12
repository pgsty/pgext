

## 用法

> [pg_checksums: 在离线 Postgres 集群中激活/停用/验证校验和](https://github.com/credativ/pg_checksums)

`pg_checksums_ext` 是一个命令行工具（基于 PostgreSQL 内置的 `pg_checksums`），可以验证、激活或停用 PostgreSQL 集群的页级校验和。它扩展了内置工具，增加了在线验证、`SIGUSR1` 进度切换、精细进度报告和 I/O 速率限制功能。

### 验证校验和（可在线执行）

```bash
pg_checksums_ext -D /path/to/data --check
```

### 启用校验和（需要干净关闭）

```bash
pg_checksums_ext -D /path/to/data --enable
```

### 禁用校验和（需要干净关闭）

```bash
pg_checksums_ext -D /path/to/data --disable
```

### 其他选项

- `-D, --pgdata` -- 数据目录路径
- `--check` / `--enable` / `--disable` -- 操作模式
- `--progress` -- 显示进度报告
- `--filenode` -- 仅检查特定 filenode
- `--no-sync` -- 跳过 fsync
- `--verbose` -- 详细输出
- `--debug` -- 调试输出
- 发送 `SIGUSR1` 可在操作期间切换进度报告
