

## 简介

- GitHub 仓库: [`reorg/pg_repack`](https://github.com/reorg/pg_repack)
- [**项目文档**](https://reorg.github.io/pg_repack/)

pg_repack 是一个 PostgreSQL 扩展，用于消除表和索引中的膨胀（bloat），并可选择性地恢复聚簇索引的物理排列顺序。
与 [CLUSTER](https://www.postgresql.org/docs/current/static/sql-cluster.html) 和 [VACUUM FULL](https://www.postgresql.org/docs/current/static/sql-vacuum.html) 不同，pg_repack 可以**在线工作**，在处理过程中不会对目标表持有排他锁。pg_repack 性能出色，效果与直接使用 CLUSTER 相当。

pg_repack 是早期 [pg_reorg](https://github.com/reorg/pg_reorg) 项目的分支。

你可以选择以下任一方式来重组表：

- 在线 CLUSTER（按聚簇索引排序）
- 按指定列排序
- 在线 VACUUM FULL（仅压缩行存储）
- 仅重建或迁移表的索引

注意事项：

- 只有超级用户或表和索引的所有者才能使用此工具。若要以表所有者身份运行 pg_repack，需使用 `--no-superuser-check` 选项。
- 目标表必须拥有 PRIMARY KEY，或者至少有一个定义在 NOT NULL 列上的 UNIQUE 索引。

### 关于 pg_reorg

pg_repack 是 [pg_reorg](https://github.com/reorg/pg_reorg) 项目的分支，pg_reorg 曾经非常成功，但不幸的是，自 2011 年底以来，pg_reorg 的新功能开发已经放缓或停止。

pg_repack 最初作为 pg_reorg 的直接替代品发布，解决了 pg_reorg 最后一个版本的一些不足之处（例如对 PostgreSQL 9.2 的支持以及 EXTENSION 打包方式），并修复了已知的 Bug。

pg_repack 1.2 引入了更多新特性（并行索引构建、仅重建索引的功能）和 Bug 修复。在某些情况下，其行为可能与 1.1.x 版本有所不同，因此不应视为直接替代品：建议在从旧版本升级前仔细阅读文档。


## 环境要求

PostgreSQL 版本
: 支持 PostgreSQL 9.5、9.6、10、11、12、13、14、15、16、17、18。不支持 PostgreSQL 9.4 及更早版本。

磁盘空间
: 执行全表 repack 所需的可用磁盘空间大约为目标表及其索引大小的两倍。例如，如果要重组的表和索引总大小为 1GB，则需要额外 2GB 的磁盘空间。


## 下载与安装

你可以从 PGXN 网站[下载 pg_repack](http://pgxn.org/dist/pg_repack/)，解压后按照下面的安装说明操作。

也可以使用 [PGXN Client](https://pgxn.github.io/pgxnclient/) 来下载、编译和安装：

```bash
$ pgxn install pg_repack
```

pg_repack 可以在 UNIX 或 Linux 上使用 `make` 构建，会自动使用 PGXS 构建框架。在构建之前，你可能需要安装 PostgreSQL 开发包（如 `postgresql-devel` 等），并将包含 `pg_config` 的目录添加到 `$PATH` 中。然后执行：

```bash
$ cd pg_repack
$ make
$ sudo make install
```

也可以在 Windows 上使用 Microsoft Visual C++ 2010 构建。项目文件位于 `msvc` 目录中。

安装完成后，在你要处理的数据库中加载 pg_repack 扩展。pg_repack 以标准扩展形式打包，可以这样加载：

```bash
$ psql -c "CREATE EXTENSION pg_repack" -d your_database
```

可以使用 `DROP EXTENSION pg_repack` 或直接删除 `repack` 模式来卸载 pg_repack。

如果你从旧版本的 pg_repack 或 pg_reorg 升级，只需按上述方法删除旧版本，然后安装新版本即可。


## 用法

```
pg_repack [OPTION]... [DBNAME]
```

可以在 `OPTIONS` 中指定以下选项：

```
选项：
  -a, --all                          重组所有数据库
  -t, --table=TABLE                  仅重组指定的表
  -I, --parent-table=TABLE           重组指定的父表及其继承子表
  -c, --schema=SCHEMA                仅重组指定模式中的表
  -s, --tablespace=TBLSPC            将重组后的表移动到新的表空间
  -S, --moveidx                      同时将索引移动到 TBLSPC 指定的表空间
  -o, --order-by=COLUMNS             按指定列排序（替代聚簇键）
  -n, --no-order                     执行 VACUUM FULL 而非 CLUSTER
  -N, --dry-run                      仅打印将要重组的内容，不实际执行
  -j, --jobs=NUM                     为每张表使用指定数量的并行任务
  -i, --index=INDEX                  仅移动指定的索引
  -x, --only-indexes                 仅移动指定表的索引
  -T, --wait-timeout=SECS            锁冲突时取消其他后端进程的超时时间
  -D, --no-kill-backend              超时后不终止其他后端进程
  -Z, --no-analyze                   完成后不执行 ANALYZE
  -k, --no-superuser-check           跳过客户端的超级用户检查
  -C, --exclude-extension            不重组属于指定扩展的表
      --no-error-on-invalid-index    即使发现无效索引也继续重组
      --error-on-invalid-index       发现无效索引时不执行重组（默认行为）
      --apply-count                  回放期间每个事务中应用的元组数量
      --switch-threshold             当日志表中剩余指定数量的元组时执行表切换

连接选项：
  -d, --dbname=DBNAME                要连接的数据库
  -h, --host=HOSTNAME                数据库服务器主机名或 socket 目录
  -p, --port=PORT                    数据库服务器端口
  -U, --username=USERNAME            连接使用的用户名
  -w, --no-password                  从不提示输入密码
  -W, --password                     强制提示输入密码

通用选项：
  -e, --echo                         回显发送到服务器的查询
  -E, --elevel=LEVEL                 设置输出消息级别
  --help                             显示帮助信息后退出
  --version                          显示版本信息后退出
```


## 选项详解

### 重组选项

`-a`, `--all`
: 尝试重组集群中的所有数据库。未安装 `pg_repack` 扩展的数据库将被跳过。

`-t TABLE`, `--table=TABLE`
: 仅重组指定的表。可以通过多次使用 `-t` 来指定多张表。默认情况下，目标数据库中所有符合条件的表都会被重组。

`-I TABLE`, `--parent-table=TABLE`
: 重组指定的表及其所有继承子表。可以通过多次使用 `-I` 来指定多个表继承层次。

`-c`, `--schema`
: 仅重组指定模式（schema）中的表。可以通过多次使用 `-c` 来指定多个模式。可与 `--tablespace` 配合使用，将表移动到不同的表空间。

`-o COLUMNS [,...]`, `--order-by=COLUMNS [,...]`
: 按指定列执行在线 CLUSTER。

`-n`, `--no-order`
: 执行在线 VACUUM FULL。从 1.2 版本起，这是非聚簇表的默认行为。

`-N`, `--dry-run`
: 列出将要重组的对象，然后退出，不实际执行。

`-j`, `--jobs`
: 创建指定数量的额外 PostgreSQL 连接，利用这些连接并行重建每张表的索引。并行索引构建仅支持全表重组，不适用于 `--index` 或 `--only-indexes` 选项。如果你的 PostgreSQL 服务器有额外的 CPU 核心和磁盘 I/O 资源，这是加速 pg_repack 的有效方法。

`-s TBLSPC`, `--tablespace=TBLSPC`
: 将重组后的表移动到指定的表空间，本质上是 `ALTER TABLE ... SET TABLESPACE` 的在线版本。除非同时指定了 `--moveidx`，否则表的索引将保留在原始表空间中。

`-S`, `--moveidx`
: 同时将重组表的索引移动到 `--tablespace` 选项指定的表空间中。

`-i`, `--index`
: 仅重组指定的索引。可以通过多次使用 `-i` 来指定多个索引。可与 `--tablespace` 配合使用，将索引移动到不同的表空间。

`-x`, `--only-indexes`
: 仅重组指定表的索引，目标表必须通过 `--table` 或 `--parent-table` 选项指定。

`-T SECS`, `--wait-timeout=SECS`
: pg_repack 需要在重组过程的开始和结束阶段各获取一次排他锁。此设置控制 pg_repack 等待获取该锁的秒数。如果在超时后仍未能获取锁，且未指定 `--no-kill-backend` 选项，pg_repack 将强制取消冲突的查询。对于 PostgreSQL 8.4 及更新版本，如果超过两倍超时时间后仍有后端进程阻塞，pg_repack 会回退使用 `pg_terminate_backend()` 来断开这些连接。默认值为 60 秒。

`-D`, `--no-kill-backend`
: 如果在 `--wait-timeout` 指定的时间内无法获取锁，则跳过该表的重组，而非取消冲突的查询。默认值为 false。

`-Z`, `--no-analyze`
: 在全表重组后禁用 ANALYZE。如未指定此选项，则在重组后自动执行 ANALYZE。

`-k`, `--no-superuser-check`
: 跳过客户端的超级用户检查。此设置适用于支持非超级用户运行 pg_repack 的平台。

`-C`, `--exclude-extension`
: 跳过属于指定扩展的表。某些扩展可能在查询规划等阶段高度依赖这些表。

`--switch-threshold`
: 当日志表中剩余的元组数量达到此阈值时执行表切换。此设置可用于避免在写入密集型表上无法追赶变更的问题。


### 连接选项

连接服务器的选项。不能同时使用 `--all` 和 `--dbname` 或 `--table` 或 `--parent-table`。

`-d DBNAME`, `--dbname=DBNAME`
: 指定要重组的数据库名称。如果未指定此选项且未使用 `-a`（或 `--all`），则从环境变量 `PGDATABASE` 读取数据库名。如果该变量也未设置，则使用连接时指定的用户名。

`-h HOSTNAME`, `--host=HOSTNAME`
: 指定服务器运行所在机器的主机名。如果值以斜杠开头，则用作 Unix 域套接字的目录。

`-p PORT`, `--port=PORT`
: 指定服务器监听连接的 TCP 端口或本地 Unix 域套接字文件扩展名。

`-U USERNAME`, `--username=USERNAME`
: 连接使用的用户名。

`-w`, `--no-password`
: 从不提示输入密码。如果服务器要求密码认证且无法通过其他方式（如 `.pgpass` 文件）提供密码，则连接将失败。此选项在无人值守的批处理作业和脚本中非常有用。

`-W`, `--password`
: 在连接数据库之前强制提示输入密码。此选项并非必需，因为如果服务器要求密码认证，程序会自动提示。但 pg_repack 会浪费一次连接尝试来探测服务器是否需要密码。在某些情况下，使用 `-W` 可以避免这次额外的连接尝试。


### 通用选项

`-e`, `--echo`
: 回显发送到服务器的命令。

`-E LEVEL`, `--elevel=LEVEL`
: 选择输出消息级别，可选值包括 `DEBUG`、`INFO`、`NOTICE`、`WARNING`、`ERROR`、`LOG`、`FATAL` 和 `PANIC`。默认值为 `INFO`。

`--help`
: 显示程序用法。

`--version`
: 显示程序版本号。


## 环境变量

`PGDATABASE`、`PGHOST`、`PGPORT`、`PGUSER`
: 默认连接参数。

与大多数其他 PostgreSQL 工具一样，此工具也使用 libpq 支持的环境变量（参见[环境变量](http://www.postgresql.org/docs/current/static/libpq-envars.html)）。


## 使用示例

对数据库 `test` 中所有已聚簇的表执行在线 CLUSTER，对所有未聚簇的表执行在线 VACUUM FULL：

```bash
$ pg_repack test
```

对数据库 `test` 中的表 `foo` 和 `bar` 执行在线 VACUUM FULL（忽略已有的聚簇索引）：

```bash
$ pg_repack --no-order --table foo --table bar test
```

将表 `foo` 的所有索引移动到表空间 `tbs`：

```bash
$ pg_repack -d test --table foo --only-indexes --tablespace tbs
```

将指定索引移动到表空间 `tbs`：

```bash
$ pg_repack -d test --index idx --tablespace tbs
```


## 诊断信息

pg_repack 失败时会报告错误消息。以下列表列出了各类错误的原因。

发生致命错误后需要手动清理。清理方法是从数据库中删除 pg_repack 然后重新安装：在出错的数据库中执行 `DROP EXTENSION pg_repack CASCADE`，然后执行 `CREATE EXTENSION pg_repack`。

**INFO: database "db" skipped: pg_repack VER is not installed in the database**
: 使用 `--all` 选项时，目标数据库中未安装 pg_repack。请在该数据库中创建 pg_repack 扩展。

**ERROR: pg_repack VER is not installed in the database**
: 通过 `--dbname` 指定的数据库中未安装 pg_repack。请在该数据库中创建 pg_repack 扩展。

**ERROR: program 'pg_repack V1' does not match database library 'pg_repack V2'**
: `pg_repack` 可执行文件与数据库中的库文件（`.so` 或 `.dll`）版本不匹配。可能是 `$PATH` 中指向了错误的二进制文件，或者连接到了错误的数据库。请检查程序目录和数据库；如果确认无误，可能需要重新安装 pg_repack。

**ERROR: extension 'pg_repack V1' required, found 'pg_repack V2'**
: 数据库中的 SQL 扩展版本与 pg_repack 程序要求的版本不一致。应该从数据库中删除该扩展并重新加载。

**ERROR: relation "table" must have a primary key or not-null unique keys**
: 目标表没有定义 PRIMARY KEY 或任何 UNIQUE 约束。请为该表定义 PRIMARY KEY 或 UNIQUE 约束。

**ERROR: query failed: ERROR: column "col" does not exist**
: 目标表中不存在 `--order-by` 选项指定的列。请指定已存在的列。

**WARNING: the table "tbl" already has a trigger called repack_trigger**
: 该触发器可能是之前在该表上运行 pg_repack 时安装的，但由于中断或其他原因未能清理临时对象。可以通过删除并重新创建扩展来移除所有临时对象。

**ERROR: Another pg_repack command may be running on the table. Please try again later.**
: 当两个 pg_repack 命令并发操作同一张表时，存在死锁风险。请稍后重试。

**WARNING: Cannot create index "schema"."index_xxxxx", already exists**
: 详情：可能是之前中断的 pg_repack 操作留下的无效索引。请使用 `DROP INDEX "schema"."index_xxxxx"` 删除该索引后重试。这是一个 pg_repack 创建的临时索引残留，程序不会冒险自动删除它。如果该索引确实是旧的 pg_repack 任务遗留的，你可以直接 DROP INDEX 然后重新执行 repack。


## 使用限制

### 临时表

pg_repack 无法重组临时表。

### GiST 索引

pg_repack 无法按 GiST 索引对表进行聚簇。

### DDL 命令

在 pg_repack 运行期间，无法对目标表执行 DDL 命令。pg_repack 会在全表重组期间持有 SHARE UPDATE EXCLUSIVE 锁来保证这一点。

如果你使用的是 1.1.8 或更早的版本，绝对不要在 pg_repack 运行期间对目标表执行任何 DDL 命令。在大多数情况下 pg_repack 会正确回滚并失败，但在这些早期版本中存在某些可能导致数据损坏的情况。


## 工作原理

### 全表重组

执行全表重组时，pg_repack 将：

1. 创建一张日志表，用于记录对原始表的变更
2. 在原始表上添加触发器，将 INSERT、UPDATE 和 DELETE 操作记录到日志表中
3. 创建一张包含旧表所有行的新表
4. 在新表上构建索引
5. 将日志表中积累的所有变更应用到新表
6. 通过系统目录（system catalogs）交换新旧表，包括索引和 TOAST 表
7. 删除原始表

pg_repack 仅在初始设置阶段（步骤 1 和 2）以及最终的交换与删除阶段（步骤 6 和 7）短暂持有 ACCESS EXCLUSIVE 锁。在其余时间，pg_repack 只需要持有 SHARE UPDATE EXCLUSIVE 锁，这意味着 INSERT、UPDATE 和 DELETE 操作可以照常进行。


### 仅索引重组

执行仅索引重组时，pg_repack 将：

1. 使用 CONCURRENTLY 方式在表上创建新索引，其定义与旧索引一致
2. 在系统目录中将旧索引替换为新索引
3. 删除旧索引

并发创建索引有一些注意事项，详情请参阅 [PostgreSQL 文档](http://www.postgresql.org/docs/current/static/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY)。


## 版本发布记录

- **pg_repack 1.5.3**
  - 使用 `--only-indexes` 时出错会清理临时索引，避免失败后残留无效对象
  - 全表重组期间在目标表上获取 SHARE UPDATE EXCLUSIVE 锁（替代之前的 ACCESS SHARE）
  - 修复子事务缓存溢出问题

- **pg_repack 1.5.2**
  - 允许非超级用户运行 pg_repack
  - 将 `--error-on-invalid-index` 设为默认行为

- **pg_repack 1.5.1**
  - 新增 PostgreSQL 17 支持
  - 修复 `repack_trigger` 中 OID 格式类型错误
  - 修复 `repack.primary_keys` 的 NOT NULL 检查
  - 修复需要引号标识符的表空间名称处理
  - 将 `PQconnectdb()` 替换为 `PQconnectdbParams()`
  - 新增 `--apply-count` 选项
  - 使用 `--only-indexes` 时不再包含声明式分区表
  - 修复两个 VACUUM 进程并发处理同一 relfilenode 的问题
  - 重试获取 AccessShareLock 时使用保存点（savepoint）
  - 修复 relfrozenxid 和 relallvisible 的交换问题

- **pg_repack 1.5.0**
  - 新增 PostgreSQL 16 支持
  - 修复潜在的 SQL 注入漏洞
  - 支持更长的密码长度
  - 修复空密码导致的无限循环
  - 新增 `--switch-threshold` 选项
  - 修复使用无效关系时 `get_order_by()` 崩溃的问题
  - 支持之前经 VACUUM FULL 重写且所有列使用 `storage=plain` 的表
  - 更谨慎的锁获取策略

- **pg_repack 1.4.8**
  - 新增 PostgreSQL 15 支持
  - 修复声明式分区表上 `--parent-table` 的问题
  - 从错误日志中移除连接信息

- **pg_repack 1.4.7**
  - 新增 PostgreSQL 14 支持

- **pg_repack 1.4.6**
  - 新增 PostgreSQL 13 支持
  - 移除对 PostgreSQL 9.4 之前版本的支持

- **pg_repack 1.4.5**
  - 新增 PostgreSQL 12 支持
  - 修复使用 public 模式中运算符的索引并行处理问题

- **pg_repack 1.4.4**
  - 新增 PostgreSQL 11 支持
  - 移除重复的密码提示

- **pg_repack 1.4.3**
  - 修复可能的 CVE-2018-1058 攻击路径
  - 修复 PostgreSQL 中 CVE-2018-1058 变更后出现的"unexpected index definition"错误
  - 修复在较新 Ubuntu 包下的构建问题

- **pg_repack 1.4.2**
  - 新增 PostgreSQL 10 支持
  - 修复 DROP INDEX CONCURRENTLY 无法在事务块内运行的错误

- **pg_repack 1.4.1**
  - 修复 `--order-by` 选项失效的问题

- **pg_repack 1.4**
  - 新增 PostgreSQL 9.6 支持，移除对 9.1 之前版本的支持
  - 使用 `AFTER` 触发器解决 `INSERT ... ON CONFLICT` 的并发问题
  - 新增 `--no-kill-backend` 选项
  - 新增 `--no-superuser-check` 选项
  - 新增 `--exclude-extension` 选项
  - 新增 `--parent-table` 选项
  - 恢复重组表上的 TOAST 存储参数
  - 恢复重组表中列的存储类型

- **pg_repack 1.3.4**
  - 在删除原始表之前先获取排他锁
  - 不再尝试重组非日志表（unlogged table）

- **pg_repack 1.3.3**
  - 新增 PostgreSQL 9.5 支持
  - 修复 pg_repack 命令中断时可能出现的死锁
  - 修复使用 `--help` 和 `--version` 时的退出码
  - 新增日文用户手册

- **pg_repack 1.3.2**
  - 修复 pg_repack 命令中断时清理临时对象的问题
  - 修复 pg_repack 共享库与 pg_statsinfo 同时加载时可能出现的崩溃

- **pg_repack 1.3.1**
  - 新增 PostgreSQL 9.4 支持

- **pg_repack 1.3**
  - 新增 `--schema` 选项，仅重组指定模式
  - 新增 `--dry-run` 选项，执行空运行
  - 修复 OID 值超过 2B 时的 Advisory Lock 问题
  - 避免其他会话锁定待重组表时出现的死锁
  - 优化 sql_pop 批量 DELETE 的性能
  - 尝试避免在持续高写入负载的表上 pg_repack 永远无法完成的问题

- **pg_repack 1.2**
  - 支持 PostgreSQL 9.3
  - 新增 `--tablespace` 和 `--moveidx` 选项，支持在线 SET TABLESPACE
  - 新增 `--index` 选项，仅重组指定索引
  - 新增 `--only-indexes` 选项，仅重组表的索引
  - 新增 `--jobs` 选项，支持并行操作
  - 非聚簇表不再需要指定 `--no-order` 即可执行 VACUUM FULL
  - 不再等待其他数据库持有的锁
  - Bug 修复：正确处理带有 DESC、NULL FIRST/LAST、COLLATE 等选项的键索引
  - 修复删除操作导致的数据损坏 Bug
  - 更友好的程序输出和错误提示

- **pg_repack 1.1.8**
  - 新增 PostgreSQL 9.2 支持
  - 新增 PostgreSQL 9.1 及以上版本的 CREATE EXTENSION 支持
  - 等待事务完成时提供用户反馈
  - Bug 修复：允许在新提升的流复制从库上运行
  - Bug 修复：修复 pg_repack 与 Slony 2.0/2.1 的交互问题
  - Bug 修复：正确转义列名
  - Bug 修复：避免重建无效索引或将其选为键
  - Bug 修复：绝不选择部分索引作为主键

- **pg_reorg 1.1.7**（2011-08-07）
  - Bug 修复：使用了含已删除列的重组表的视图和函数可能被损坏
  - 支持 PostgreSQL 9.1 和 9.2dev（但尚未支持 EXTENSION 方式）


## 参见

- [clusterdb](http://www.postgresql.org/docs/current/static/app-clusterdb.html)
- [vacuumdb](http://www.postgresql.org/docs/current/static/app-vacuumdb.html)
