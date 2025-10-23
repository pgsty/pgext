---
title: Remove
description: How to remove PostgreSQL extensions
icon: Trash2
---

## Remove Extension

To uninstall an extension, you typically need to run the [**`DROP EXTENSION`**](https://www.postgresql.org/docs/current/sql-dropextension.html) SQL statement:

```sql
DROP EXTENSION "<extname>";
```

If other extensions or database objects depend on this extension, you’ll need to remove those dependencies first before uninstalling the extension.
Or remove all of them with `CASCADE` option:

```sql
DROP EXTENSION "<extname>" CASCADE;
```


<Callout title="" type="warning">

    The `CASCADE` option will delete all objects that depend on this extension, \
    including database objects, functions, views, etc. Use with caution!

</Callout>

Some extensions don't have `DDL`, these extensions do not require the `DROP EXTENSION` statement to uninstall.
Instead, you can simply remove the extension from the `shared_preload_libraries` (if configured) and uninstall the package.
Refer to the [**Extensions Without DDL**](/usage/config) section for more details.


------

## Remove Loading

If you’re using an extension that requires [**dynamic loading**](/usage/config) (which modifies the `shared_preload_libraries` parameter), you need to first [**re-confnigure**](/docs/pgsql/admin#config-cluster) the `shared_preload_libraries` parameter.

Remove the extension name from `shared_preload_libraries` and restart the database cluster for the changes to take effect.

For extensions that need dynamic loading, refer to the [**Extensions that Need Loading**](/list/attr#need-loading) list.



------

## Uninstall Package

After removing the extension (logical object) from **all databases** in the cluster, you can safely uninstall the extension’s software package. Ansible commands can help you do this conveniently:

```bash
ansible <cls> -m package -a "name=<extname> state=absent"
```

You can also use [**`pig`**](/pig/), or `apt`/`yum` commands directly to uninstall.

If you don’t know the extension package name, you can refer to the [**Extension List**](/list/) or check the extension package name mapping defined in [**`roles/node_id/vars`**](https://github.com/pgsty/pigsty/blob/main/roles/node_id/vars/).