## 用法

来源：

- [官方 README](https://github.com/pgbigm/ludia_funcs/blob/d220429554fbc1f04573f444a58bc5ab8919d4a6/README.md)
- [官方扩展 SQL](https://github.com/pgbigm/ludia_funcs/blob/d220429554fbc1f04573f444a58bc5ab8919d4a6/ludia_funcs--1.0.sql)
- [官方 C 实现](https://github.com/pgbigm/ludia_funcs/blob/d220429554fbc1f04573f444a58bc5ab8919d4a6/ludia_funcs.c)

`ludia_funcs` 为 PostgreSQL 提供兼容 Ludia 的文本规范化与摘要辅助函数。它主要用于迁移依赖 Ludia/Senna 行为的应用，或与 `pg_bigm` 等全文检索扩展组合使用。

### 核心流程

```sql
CREATE EXTENSION ludia_funcs;

SELECT pgs2norm('ｱﾌﾟﾘｹｰｼｮﾝ');
SELECT pgs2snippet1(1, 32, 1, '★', '★', 0,
                    'ｱﾌﾟﾘｹｰｼｮﾝ', '...');
SELECT * FROM pgs2seninfo();
```

安装后的 SQL 接口较为精简：

- `pgs2norm(text)` 使用 Senna 规范化文本。
- `pgs2snippet1(...)` 按调用方给出的宽度、结果数、标签、关键词数、源文本与关键词参数生成高亮摘要。
- `pgs2seninfo()` 返回所链接的 Senna 版本与配置选项。
- `pgs2textporter1(text)` 在构建时启用 TextPorter 后调用该可选集成。

### 配置与边界

`ludia_funcs` 在库首次加载时初始化 Senna；控制文件并不要求 `shared_preload_libraries`。可选 GUC 包括 `ludia_funcs.norm_cache_limit`（负数表示采用 `work_mem`，零表示无限制）与 `ludia_funcs.escape_snippet_keyword`。启用 TextPorter 的构建还可能暴露其他设置。

核心函数要求 UTF-8 数据库。`pgs2textporter1()` 会操作服务器端应用文件和临时文件，因此跨越文件系统与权限边界；不要传入不可信路径，也不要假定任意软件包都编译了 TextPorter。该扩展没有声明对 `pg_bigm` 的 SQL 依赖，此集成是可选的。
