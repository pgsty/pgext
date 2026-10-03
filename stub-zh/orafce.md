## 用法

来源：

- [README.asciidoc](https://github.com/orafce/orafce/blob/d905cb474fb8e2e31589f3c75940a3b9e7feb014/README.asciidoc)
- [orafce.control](https://github.com/orafce/orafce/blob/d905cb474fb8e2e31589f3c75940a3b9e7feb014/orafce.control)
- [orafce--4.16.sql](https://github.com/orafce/orafce/blob/d905cb474fb8e2e31589f3c75940a3b9e7feb014/orafce--4.16.sql)
- [4.16.12 发布说明](https://github.com/orafce/orafce/releases/tag/VERSION_4_16_12)
- [文件访问实现](https://github.com/orafce/orafce/blob/d905cb474fb8e2e31589f3c75940a3b9e7feb014/file.c)

`orafce` 提供兼容 Oracle 的函数、类型与工具包。发行版 4.16.12 仍使用控制文件和 SQL 扩展版本 4.16；两个版本号分别描述不同层次。

### 核心工作流

```sql
CREATE EXTENSION orafce;
SELECT oracle.add_months(date '2026-01-31', 1);
SELECT oracle.nvl(NULL::text, 'fallback');
SELECT oracle.decode(1, 1, 'one', 2, 'two', 'other');
SELECT dbms_output.enable();
SELECT dbms_output.put_line('Hello');
SELECT * FROM dbms_output.get_line();
```

### 类型与工具包

需要在 Oracle 风格日期中保留时分秒时，应使用 `oracle.date`。日期函数包括 `oracle.add_months`、`oracle.last_day`、`oracle.next_day`、`oracle.months_between` 及舍入、截断操作。应显式限定 `oracle.decode`、`oracle.greatest` 和 `oracle.least` 的模式，否则 PostgreSQL 解析器可能采用内置语义。

`dbms_output` 管理缓冲输出；`dbms_pipe` 与 `dbms_alert` 支持会话通信。`dbms_sql` 提供动态游标和带类型的列值读取。`dbms_utility`、`dbms_assert`、`plvstr`、`plvchr` 与 `plvsubst` 提供诊断、校验及字符串辅助功能。这些兼容函数不会将 PostgreSQL 变成 Oracle，也不提供 Oracle 过程语言运行时。

### 文件访问与 4.16.12 变更

`utl_file` 在管理员配置的允许目录内访问服务端文件；应限制授权与文件系统权限。实现会拒绝路径规范化后仍然存在的父目录引用，文件路径应避免使用父目录引用。4.16.12 发布说明明确列出的修复是 `dbms_sql` 可能发生的崩溃。

安装需要超级用户，并创建固定模式；无需预加载。修改 `search_path` 前应遵循上游配置说明。SQL 版本仍为 4.16，因此已经安装的 4.16 扩展不会仅因二进制发行版修补就获得新的 SQL 版本。应按需重新连接以使用更新后的共享库，并对照实际安装的发行版验证行为。
