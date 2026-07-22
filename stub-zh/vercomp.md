## 用法

来源：

- [官方上游文档](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/README.md)
- [官方扩展 SQL](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/vercomp--1.0.0.sql)
- [官方范围操作符回归测试](https://github.com/eendroroy/pg_vercomp/blob/767244c2ab544bbe660e85a4d61cae774ab61441/test/sql/vercomp_test_caret.sql)

`vercomp` 1.0.0 添加 `version` 基础类型、比较、B-tree 与 hash 操作符类，以及类似 semver 的范围操作符。上游明确将项目标为 obsolete，并建议用户改用 `pg_semver`。只有维护已经依赖其准确解析与排序语义的数据时才应使用 `vercomp`。

### 存储并比较版本

定义该类型的列之前先创建扩展：

```sql
CREATE EXTENSION vercomp;

CREATE TABLE releases (
  package text NOT NULL,
  release version NOT NULL
);

INSERT INTO releases VALUES
  ('demo', '1.2.3'),
  ('demo', '2.0.0-rc1'),
  ('demo', '2.0.0');

SELECT * FROM releases ORDER BY release;
SELECT version_cmp('1.2.3', '2.0.0');
SELECT version_bump('1.2.3', 2);
```

`version_cmp(version, version)` 返回负数、零或正数。`version_bump(version, integer)` 增加第 0、1 或 2 个组件。普通 `=`、`<>`、`<`、`<=`、`>` 和 `>=` 操作符支持索引比较。

### 范围操作符与转换

`~` 与 `!~` 实现项目的 tilde 风格匹配；`^` 与 `!^` 实现 caret 风格匹配。它们是两个 `version` 值之间的二元操作符，不是字符串前缀或正则表达式操作符。扩展还安装 text 与 `version` 之间的隐式转换、从 `version` 到整数的隐式转换，以及 B-tree 与 hash 操作符类。

不要假定 npm、Cargo 或其他 semver 库会产生相同范围。已复核的零主版本回归测试相对 README 公式会产生意外匹配，因此应为应用使用的每个范围建立样例。整数转换是自定义的有损编码，不是面向外部系统的稳定单调表示。

### 迁移边界

README 包含旧式预加载说明，但面向用户的类型与操作符由扩展 SQL 脚本登记；应在准确服务器上验证软件包加载流程，而不是盲目复制旧配置。替换为 `pg_semver` 前，应比较接受的语法、预发布排序、类型转换、范围操作符、索引以及转储与恢复输出。类型变更需要显式数据迁移和索引重建，不能只替换共享库。
