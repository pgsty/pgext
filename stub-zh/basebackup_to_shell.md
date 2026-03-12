

## 用法

> [basebackup_to_shell: 添加名为 shell 的自定义基础备份目标](https://www.postgresql.org/docs/current/basebackup-to-shell.html)

`basebackup_to_shell` 模块为 `pg_basebackup` 添加一个自定义的 `shell` 备份目标，允许管理员通过任意 Shell 命令传输备份归档。

### 配置

添加到 `postgresql.conf`：

```ini
shared_preload_libraries = 'basebackup_to_shell'

# 每个归档执行的命令；在 stdin 上接收归档数据
basebackup_to_shell.command = 'gzip > /backup/%f.gz'

# 可选：限制仅特定角色可使用
basebackup_to_shell.required_role = 'backup_admin'
```

### 参数

| 参数 | 类型 | 描述 |
|-----------|------|-------------|
| `basebackup_to_shell.command` | 字符串 | 要执行的 Shell 命令；通过 stdin 接收归档 |
| `basebackup_to_shell.required_role` | 字符串 | 使用 shell 目标所需的角色（空 = 任何复制用户） |

### 命令占位符

| 占位符 | 替换为 |
|-------------|---------------|
| `%f` | 归档文件名（例如 `base.tar`） |
| `%d` | 用户提供的目标详情字符串 |
| `%%` | 字面量 `%` |

### 示例

```bash
# 将备份压缩到本地目录
# postgresql.conf: basebackup_to_shell.command = 'gzip > /backup/%f.gz'
pg_basebackup --target=shell

# 使用详情参数上传到 S3
# postgresql.conf: basebackup_to_shell.command = 'aws s3 cp - s3://bucket/%d/%f'
pg_basebackup --target=shell:myprefix

# 自定义处理管道
# postgresql.conf: basebackup_to_shell.command = 'zstd | ssh backup-host "cat > /backup/%f.zst"'
pg_basebackup --target=shell
```

`%d` 占位符需要通过 `--target=shell:DETAIL` 提供目标详情。如果命令中没有 `%d`，则禁止提供详情。详情字符串只能包含字母数字字符。
