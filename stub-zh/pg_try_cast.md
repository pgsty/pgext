## 用法

来源：

- [README.md](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/README.md)
- [pg_try_cast--1.0.sql](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/pg_try_cast--1.0.sql)
- [src/pg_try_cast.c](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/src/pg_try_cast.c)
- [.github/workflows/ci.yml](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/.github/workflows/ci.yml)
- [pg_try_cast.control](https://github.com/Mazuco/pg_try_cast/blob/d8013ce29c39e53c7317725927f003379743292c/pg_try_cast.control)

`pg_try_cast` 提供 `try_cast` 与 `try_convert`，在受支持的输入转换出错时返回回退值，避免整个查询失败。发布版本 1.0.0 安装的 SQL 版本为 1.0。要求 PostgreSQL 16 或更新版本，上游 CI 覆盖 16–18。扩展可信、可迁移，无需预加载。

### 核心用法

```sql
CREATE EXTENSION pg_try_cast;
SELECT try_cast('42', NULL::integer);
SELECT try_cast('invalid', 0);
SELECT try_convert(NULL::date, '2026-02-30');
```

### 运行边界

`try_cast(value text, fallback anyelement)` 的第二个参数同时决定结果类型和出错回退值；`try_convert(fallback anyelement, value text)` 的参数顺序相反。NULL 输入仍返回 NULL，应为回退值指定明确类型。扩展会检查类型修饰符与域约束。它处理文本输入转换，不涵盖任意原生类型转换。尚未支持软错误机制的输入函数仍可能抛出错误，包括向 regtype 传入格式错误的类型名称。安装文件就绪后，拥有数据库 CREATE 权限的角色可安装此可信扩展。
