


## Usage

> [omni_manifest: Package installation manifests](https://docs.omnigres.org/omni_manifest/usage/)

The `omni_manifest` extension manages extension installation manifests with dependency resolution.

### Types

**`omni_manifest.requirement`** -- A name-version pair. Can be cast from text (`'myext=1.0'`) or JSON (`'{"myext": "1.0"}'`).

**`omni_manifest.artifact`** -- Contains a `self` requirement and an array of dependency requirements.

```sql
SELECT (omni_manifest.artifact('ext=3.0'::text, 'myext=1.0,myotherext=2.0'::text)).*;
```

### Installation Plan

```sql
SELECT * FROM unnest(omni_manifest.install_plan(
    array[
        omni_manifest.artifact('ext=3.0'::text, 'dep1=1.0,dep2=2.0'::text),
        omni_manifest.artifact('ext2=3.0'::text, 'dep1=1.0'::text)
    ]
)) WITH ORDINALITY t(name, version, position);
```

### Install Extensions

```sql
SELECT * FROM omni_manifest.install(artifacts);
```

Returns `omni_manifest.install_report` rows with status: `installed`, `missing`, `updated`, or `kept`.
