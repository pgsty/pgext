


## 用法

> [omni_manifest: 扩展安装清单](https://docs.omnigres.org/omni_manifest/usage/)

`omni_manifest` 扩展管理扩展安装清单，支持依赖解析。

### 类型

**`omni_manifest.requirement`** -- 名称-版本对。可从文本（`'myext=1.0'`）或 JSON（`'{"myext": "1.0"}'`）转换。

**`omni_manifest.artifact`** -- 包含 `self` 需求和依赖需求数组。

```sql
SELECT (omni_manifest.artifact('ext=3.0'::text, 'myext=1.0,myotherext=2.0'::text)).*;
```

### 安装计划

```sql
SELECT * FROM unnest(omni_manifest.install_plan(
    array[
        omni_manifest.artifact('ext=3.0'::text, 'dep1=1.0,dep2=2.0'::text),
        omni_manifest.artifact('ext2=3.0'::text, 'dep1=1.0'::text)
    ]
)) WITH ORDINALITY t(name, version, position);
```

### 安装扩展

```sql
SELECT * FROM omni_manifest.install(artifacts);
```

返回 `omni_manifest.install_report` 行，状态包括：`installed`、`missing`、`updated` 或 `kept`。
