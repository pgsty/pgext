## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.7.0/README.md)
- [pg_pure.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.7.0/pg_pure.control)
- [pg_pureu.control](https://api.pgxn.org/src/pg_pure/pg_pure-0.7.0/pg_pureu.control)
- [pg_pure--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.7.0/pg_pure--0.1.sql)
- [pg_pureu--0.1.sql](https://api.pgxn.org/src/pg_pure/pg_pure-0.7.0/pg_pureu--0.1.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_pure/pg_pure-0.7.0/CHANGELOG.md)

`pg_pure` 提供受信任的 Pure 过程语言与 SPI 数据库访问。分发版本 0.7.0 仍使用 SQL 扩展版本 0.1。从 0.2.0 起，受信任函数已能在 libpure 安全模式解释器中执行，早期主动拒绝执行的限制不再适用。

### 核心流程

```sql
CREATE EXTENSION pg_pure;
CREATE FUNCTION pure_factorial(n int) RETURNS numeric
LANGUAGE pg_pure AS $$ foldl (*) 1L (1..n) $$;
SELECT pure_factorial(30);
```

扩展受信任且不可迁移。具备相应数据库与模式权限的普通角色可以安装扩展并定义受信任语言函数。需要 PostgreSQL 17–19，以及上游适合嵌入的 libpure 分支 ab12de13 或更新修订。

### 类型与数据库访问

标量类型包括整数、浮点、numeric、布尔与文本。SQL NULL 用 `sql::null` 表示。精确 `numeric` 值使用 Pure 大整数或有理数；中间运算可能超出机器整数范围时应使用大整数。

`float8[]` 映射为稠密 Pure 双精度矩阵：一维数组成为行向量，二维数组成为矩阵。NULL 元素和超过二维的数组会被拒绝。`sql::query`、`sql::exec`、`sql::notice` 与 `sql::warning` 提供数据库访问与消息，数据库错误会转为 `sql::error`。已支持参数化 SPI 和集合返回函数；触发器与复合类型仍不支持。不要把不可信值直接拼入 SQL 字符串。

### 沙箱与模块

安全模式拒绝 `extern`、内联 C、`eval`/`evalcmd`、任意原生或未获准的 `using`，以及危险的前导库操作。受信任与非受信任语言使用相互独立的后端本地解释器。非受信任的配套语言 `pg_pureu` 仍可执行原生代码，必须限于超级用户。

DBA 通过 `pg_pure.using` 选择模块，默认值为 `math`。使用 0.6.0 的已编译 GSL Cholesky 求解器时，应安装匹配的 Pure 模块，并通过该设置预载和允许模块：

```ini
pg_pure.using = 'math gsl::linalg'
```

完成配置后，受信任函数体可直接调用 `gsl::linalg::cholesky_solve`。受信任函数体中的命名空间导入 `using gsl::linalg` 仍会被 libpure 的模块路径检查拒绝，应使用 DBA 管理的设置。

### 运行与升级

两个语言均加载 `pg_pure` 库。可选的 `shared_preload_libraries` 加载需要重启，但上游测量表明主进程预热不能消除派生后端首次调用的初始化成本；连接池可避免反复支付这项成本。要求使用的 libpure 分支隐藏了 LLVM 符号，早期要求禁用 PostgreSQL JIT 的建议已过时。应部署匹配的扩展与 libpure 文件，并替换已加载旧库的进程。分发版本 0.6.0 不改变 SQL 版本 0.1，因此不能单凭该版本判断实际加载的库。

### 0.6.0 发行版

0.6.0 为 `sql::query` 与 `sql::exec` 增加参数数组，通过 SPI 绑定 Pure 值，不再把值拼接进 SQL 文本。它还支持从 Pure 列表或在函数体内有界的惰性流返回 `RETURNS SETOF` 标量或数组结果。函数外层的 `LIMIT` 不会限制无界流，因此应在 Pure 函数体内设置边界。受信任语言仍不可迁移，并继续使用 DBA 管理的模块白名单。


### 0.7.0 发行版

`bytea` 映射为 Pure 整数行矩阵，可保留包括 NUL 字节在内的任意二进制数据；元素超出 0..255 时会报错。域类型参数与结果按基础类型转换，同时执行域约束。此前 0.6.1 修复 numeric 精度、符号和显示精度转换；0.6.2 拒绝受信任函数体及 DO 块逃出单表达式包装、向同一后端的其他函数泄露规则。应部署修正后的动态库并替换旧会话，SQL 版本仍为 0.1。
