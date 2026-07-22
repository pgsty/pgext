## 用法

来源：

- [官方仓库 README](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)
- [扩展控制文件](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_show_hooks/pgx_show_hooks.control)
- [Hook 检查实现](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgx_show_hooks/src/lib.rs)

`pgx_show_hooks` 是一个诊断扩展，用于报告 PostgreSQL hook 变量与 PL/pgSQL 插件回调中当前安装的进程内地址。它适合在测试扩展加载和 hook 冲突时比较某个后端的 hook 状态；它能识别已占用的回调槽位，但不能确定每个回调所属的扩展，也不会修改 hook 顺序。

### 核心流程

以超级用户创建扩展，然后在需要观察的同一个后端中检查 hook 状态：

```sql
CREATE EXTENSION pgx_show_hooks;

SELECT name, addr
FROM show_hooks.all()
WHERE addr IS NOT NULL
ORDER BY name;
```

`show_hooks.all()` 为每个 hook 或 PL/pgSQL 回调返回一行。存在回调时，`addr` 是十六进制进程地址；槽位为空时则为 `NULL`。可以先采集基线，再加载或预加载待测扩展，并对等价的后端会话进行结果比较。

### 对象索引

- `show_hooks.all() -> table(name text, addr text)` 列出 PostgreSQL hook 变量、`PLpgSQL_plugin` rendezvous 指针以及各 PL/pgSQL 回调指针。

### 运维说明

`pgx_show_hooks` 安装到固定的 `show_hooks` 模式，创建时需要超级用户权限。其控制文件没有声明预加载要求。结果是进程局部的观察值：PostgreSQL 后端是独立进程，hook 地址可能随会话或重启变化，仅在某个后端加载的扩展也可能不会出现在另一个后端中。

应将 `addr` 视为敏感诊断数据。它适合在受控测试中判断相等性和是否存在，不是跨进程或跨版本的稳定标识符。审阅的仓库是 PostgreSQL 15 时期的研究项目；在依据其结果判断 hook 覆盖范围前，应验证准确的二进制与服务器组合。
