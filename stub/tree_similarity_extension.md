## Usage

Sources:

- [Extension control file](https://github.com/LukMRVC/tree_similarity_extension/blob/4b5b0753dd00504d7df3cbca8e22b67b55017fa6/tree_similarity_extension.control)
- [Rust package metadata](https://github.com/LukMRVC/tree_similarity_extension/blob/4b5b0753dd00504d7df3cbca8e22b67b55017fa6/Cargo.toml)
- [Extension implementation and test](https://github.com/LukMRVC/tree_similarity_extension/blob/4b5b0753dd00504d7df3cbca8e22b67b55017fa6/src/lib.rs)

Despite its name, `tree_similarity_extension` 0.0.0 does not implement tree similarity. The reviewed source is a minimal pgrx demonstration that installs one greeting function and a matching unit test.

### Call the implemented function

```sql
CREATE EXTENSION tree_similarity_extension;

SELECT hello_tree_similarity_extension();
```

The function returns the fixed text `Hello, from my own extension 2!`. The repository defines no tree type, parser, distance or similarity function, operator, index, table, or configuration parameter.

### Prototype boundary

Do not select this extension based on its catalog name when an application needs a tree-similarity algorithm. Treat version 0.0.0 as a scaffold and verify the actual installed object list before allowing dependencies to form.

The package pins pgrx 0.11.4, defaults to a PostgreSQL 13 build feature, declares features for PostgreSQL 11 through 16, and provides no compatibility claim beyond those build declarations. Its control file marks installation as superuser-required and non-relocatable. The repository has no README, release, SQL upgrade path, or usage contract; use it only as reviewed demonstration source, not as a production similarity service.
