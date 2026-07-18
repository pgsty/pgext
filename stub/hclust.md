## Usage

Sources:

- [Pinned extension control file](https://github.com/IrNesterova/hierarchicalClusteringForDiploma/blob/0e9890797682dced54eca9a857d9eb1c11664f6a/hclust.control)
- [Version 0.0.1 installation SQL](https://github.com/IrNesterova/hierarchicalClusteringForDiploma/blob/0e9890797682dced54eca9a857d9eb1c11664f6a/hclust--0.0.1.sql)
- [Pinned clustering implementation](https://github.com/IrNesterova/hierarchicalClusteringForDiploma/blob/0e9890797682dced54eca9a857d9eb1c11664f6a/hclust.c)

hclust version 0.0.1 is a diploma-project C function for hierarchical clustering. It executes a supplied SELECT through SPI, treats the first selected column as a numeric row identifier and the remaining columns as features, adds a cluster_id column to a supplied table, and updates each selected row with its assigned cluster.

### Experimental invocation

The six arguments are query text, target cluster count, target table name, linkage selector, distance selector, and hierarchy selector. A minimal numeric experiment is:

```sql
CREATE EXTENSION hclust;

CREATE TABLE hclust_points (
    id integer PRIMARY KEY,
    x double precision,
    y double precision
);

INSERT INTO hclust_points VALUES
    (1, 0.0, 0.0),
    (2, 0.1, 0.1),
    (3, 9.0, 9.0),
    (4, 9.1, 9.1);

SELECT hclust(
    'SELECT id, x, y FROM hclust_points ORDER BY id',
    2,
    'hclust_points',
    1,
    1,
    'a'
);

SELECT id, cluster_id FROM hclust_points ORDER BY id;
```

Selectors are undocumented SQL constants encoded in C. The reviewed source maps linkage 1 through 4 to average, centroid, complete, and single linkage; distance 1 through 6 selects Euclidean, Chebyshev, Manhattan, cosine, Jaccard, or Sorensen logic; a or A selects agglomerative clustering.

### Severe implementation limits

The project has no README, tests, release notes, or compatibility matrix. It uses fixed 100-feature and short label buffers, converts values from text with minimal validation, emits extensive server messages, and contains unfinished or incorrect mathematical paths. Table names and row labels are interpolated into ALTER and UPDATE commands with fixed-size sprintf buffers rather than identifier or literal quoting. The supplied SELECT is also executed verbatim.

Only a superuser in an isolated disposable database should test hclust, using hard-coded trusted identifiers and input. Do not expose the function to users, do not run it against valuable tables, and do not treat its cluster assignments as validated analytical output.
