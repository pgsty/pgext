## 用法

来源：

- [官方项目页面](https://pgxn.org/dist/plparrot/)
- [官方上游README](https://github.com/leto/plparrot/blob/5adaf4be2d00d8ca0aee06a6dbd1cec21ceff12a/README.md)
- [官方语言定义SQL](https://github.com/leto/plparrot/blob/5adaf4be2d00d8ca0aee06a6dbd1cec21ceff12a/plparrot.sql.in)

`plparrot` 0.4.0 将 Parrot 虚拟机嵌入到 PostgreSQL 中。其遗留安装脚本会创建受信任和不受信任的 Parrot/PIR 语言名称，并且当可用时，还会创建 Perl 6 语言处理器。

### 核心工作流

针对目标 PostgreSQL 和 Parrot 安装构建模块，然后在目标数据库中运行其安装的 SQL 脚本：

```sh
psql -d appdb -f "$(pg_config --sharedir)/contrib/plparrot.sql"
```

该分发包早于现代扩展打包；请勿使用 `CREATE EXTENSION plparrot`。

### 已安装语言

- `plparrot` 及其别名 `plpir` 被创建为受信任语言。
- `plparrotu` 和 `plpiru` 是不受信任的变体。
- 当相应的 Rakudo 支持被构建时，会创建 `plperl6` 和 `plperl6u`。

### 要求与注意事项

- 审查过的控制、注册表或目录证据标识版本 `0.4.0`。
- 上游测试工作流期望 `plpgsql`；运行时语言支持需要兼容的 Parrot 和可选的 Rakudo 库。
- 上游分发包使用了遗留或非控制安装布局；请勿假设现代 `ALTER EXTENSION UPDATE` 行为。
- 0.4.0 版本发布于 2011 年，审查过的仓库的最新提交是历史记录。在尝试进行现代构建之前，请验证其虚拟机、Perl 6 命名、C API 和安全假设是否可行。
