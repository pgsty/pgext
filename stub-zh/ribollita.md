## 用法

来源：

- [官方上游 README](https://github.com/rvianello/ribollita/blob/b7b48c0b29d22166d7e42859e7b5b132a3e22b73/README.md)
- [官方扩展控制文件 (ribollita.control)](https://github.com/rvianello/ribollita/blob/b7b48c0b29d22166d7e42859e7b5b132a3e22b73/src/ribollita.control)
- [官方扩展 SQL (ribollita--0.0.sql)](https://github.com/rvianello/ribollita/blob/b7b48c0b29d22166d7e42859e7b5b132a3e22b73/src/ribollita--0.0.sql)

`ribollita` — 如果不是汤，那就是面包蘸酱。当应用程序需要此特定数据库功能时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION ribollita;
```

在目标数据库中安装扩展，如果有可用的上游最小示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `boost_version()` 是一个扩展函数，返回 `cstring`。
- `mol_add_hs(mol, explicit_only bool DEFAULT false, add_coords bool DEFAULT false, add_residue_info bool DEFAULT false)` 是一个扩展函数，返回 `mol`。
- `mol_amw(mol)` 是一个扩展函数，返回 `float8`。
- `mol_formal_charge(mol)` 是一个扩展函数，返回 `integer`。
- `mol_fragments(mol, sanitize_fragments bool DEFAULT true, copy_conformers bool DEFAULT true)` 是一个扩展函数，返回 `TABLE`。
- `mol_from_molblock(molblock cstring, sanitize bool DEFAULT true, remove_hs bool DEFAULT true, strict_parsing bool DEFAULT true)` 是一个扩展函数，返回 `mol`。
- `mol_from_smarts(smiles cstring, allow_cxsmiles bool DEFAULT true, strict_cxsmiles bool DEFAULT true, parse_name bool DEFAULT false, merge_hs bool DEFAULT true)` 是一个扩展函数，返回 `mol`。
- `mol_from_smiles(smiles cstring, sanitize bool DEFAULT true, allow_cxsmiles bool DEFAULT true, strict_cxsmiles bool DEFAULT true, parse_name bool DEFAULT false, remove_hs bool DEFAULT true, use_legacy_stereo bool DEFAULT true)` 是一个扩展函数，返回 `mol`。
- `mol_in(cstring)` 是一个扩展函数，返回 `mol`。
- `mol_kekulize(mol)` 是一个扩展函数，返回 `mol`。
- `mol_out(mol)` 是一个扩展函数，返回 `cstring`。
- `mol_recv(internal)` 是一个扩展函数，返回 `mol`。
- `mol_remove_all_hs(mol, sanitize bool DEFAULT true)` 是一个扩展函数，返回 `mol`。
- `mol_remove_hs(mol, implicit_only bool DEFAULT false, update_explicit_count bool DEFAULT false, sanitize bool DEFAULT true)` 是一个扩展函数，返回 `mol`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
