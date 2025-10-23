---
title: 卸载
description: 如何卸载 PostgreSQL 扩展
icon: Trash2
---

## 移除扩展

要卸载扩展，通常需要执行 [**`DROP EXTENSION`**](/zh/usage/remove) SQL 语句：

```sql
DROP EXTENSION "<extname>";
```

如果有其他扩展或数据库对象依赖于该扩展，需先移除这些依赖，或使用 `CASCADE` 选项一并删除：

```sql
DROP EXTENSION "<extname>" CASCADE;
```


<Callout title="" type="warning">

    `CASCADE` 选项会删除所有依赖于该扩展的对象，
    包括数据库对象、函数、视图等。请谨慎操作！

</Callout>

部分扩展不包含 DDL，这类扩展无需执行 `DROP EXTENSION` 语句卸载。
只需从 `shared_preload_libraries`（如有配置）中移除，并卸载对应软件包即可。
详见 [**无 DDL 扩展**](/zh/usage/config) 部分。

------

## 移除动态加载

若使用了需[**动态加载**](/zh/usage/config)的扩展（即需修改 `shared_preload_libraries` 参数），需先[**重新配置**](/zh/docs/pgsql/admin#config-cluster) `shared_preload_libraries` 参数。

从 `shared_preload_libraries` 中移除扩展名，并重启数据库集群以使配置生效。

需动态加载的扩展详见 [**需加载扩展列表**](/zh/list/attr#need-loading)。

------

## 卸载软件包

在**所有数据库**中移除扩展（逻辑对象）后，可安全卸载扩展软件包。可通过 Ansible 命令便捷操作：

```bash
ansible <cls> -m package -a "name=<extname> state=absent"
```

也可直接使用 [**`pig`**](/zh/pig/)，或 `apt`/`yum` 命令卸载。

如不清楚扩展包名，可参考 [**扩展列表**](/zh/list/)，或查阅 [**`roles/node_id/vars`**](https://github.com/pgsty/pigsty/blob/main/roles/node_id/vars/) 中的扩展包名映射。