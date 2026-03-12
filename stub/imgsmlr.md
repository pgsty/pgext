


## Usage

> [imgsmlr: similar images search for PostgreSQL using Haar wavelet transform](https://github.com/postgrespro/imgsmlr)

The `imgsmlr` extension implements similar image searching functionality based on Haar wavelet transforms. It provides two data types and functions for converting images into searchable signatures.

```sql
CREATE EXTENSION imgsmlr;
```

### Data Types

| Datatype    | Storage Length | Description |
|-------------|---------------:|-------------|
| `pattern`   | 16388 bytes    | Result of Haar wavelet transform on the image |
| `signature` | 64 bytes       | Short representation of pattern for fast GiST index searches |

### Functions

| Function | Return Type | Description |
|----------|-------------|-------------|
| `jpeg2pattern(bytea)` | pattern | Convert JPEG image data into pattern |
| `png2pattern(bytea)` | pattern | Convert PNG image data into pattern |
| `gif2pattern(bytea)` | pattern | Convert GIF image data into pattern |
| `pattern2signature(pattern)` | signature | Create signature from pattern |
| `shuffle_pattern(pattern)` | pattern | Shuffle pattern for less sensitivity to image shift |

### Operators

| Operator | Left | Right | Return | Description |
|----------|------|-------|--------|-------------|
| `<->` | pattern | pattern | float8 | Euclidean distance between two patterns |
| `<->` | signature | signature | float8 | Euclidean distance between two signatures |

The `signature` type supports GiST indexing with KNN on the `<->` operator.

### Example

Create a table of patterns and signatures from JPEG images:

```sql
CREATE TABLE pat AS (
    SELECT
        id,
        shuffle_pattern(pattern) AS pattern,
        pattern2signature(pattern) AS signature
    FROM (
        SELECT id, jpeg2pattern(data) AS pattern
        FROM image
    ) x
);

ALTER TABLE pat ADD PRIMARY KEY (id);
CREATE INDEX pat_signature_idx ON pat USING gist (signature);
```

Search for the top 10 similar images to a given image:

```sql
SELECT id, smlr
FROM (
    SELECT
        id,
        pattern <-> (SELECT pattern FROM pat WHERE id = :id) AS smlr
    FROM pat
    WHERE id <> :id
    ORDER BY signature <-> (SELECT signature FROM pat WHERE id = :id)
    LIMIT 100
) x
ORDER BY x.smlr ASC
LIMIT 10;
```

The inner query selects the top 100 candidates by signature using the GiST index. The outer query refines to the top 10 by pattern distance.
