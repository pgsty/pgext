## Usage

Sources:

- [Official README](https://github.com/hjort/pgafis/blob/98096d694a33414ace34ca51ff194890717ef8da/README.md)
- [Version 1.2 extension SQL](https://github.com/hjort/pgafis/blob/98096d694a33414ace34ca51ff194890717ef8da/src/pgafis--1.2.sql)
- [Official PGXS Makefile](https://github.com/hjort/pgafis/blob/98096d694a33414ace34ca51ff194890717ef8da/src/Makefile)

`pgafis` adds Automated Fingerprint Identification System primitives to PostgreSQL. It compresses fingerprint images, extracts and inspects minutiae templates, measures image quality, and scores one-to-one or one-to-many template matches. The extension wraps the legacy NIST Biometric Image Software (NBIS) libraries; it does not provide an indexed search method or decide an application's identity threshold.

### Core Workflow

Install the extension, then store the original image and derived artifacts separately. The upstream examples use raw PGM images, WSQ-compressed images, binary MDT templates, and text XYT minutiae:

```sql
CREATE EXTENSION pgafis;

CREATE TABLE fingerprints (
    id  text PRIMARY KEY,
    pgm bytea,
    wsq bytea,
    mdt bytea,
    xyt text
);

UPDATE fingerprints
SET wsq = cwsq(pgm, 2.25, 300, 300, 8, NULL)
WHERE wsq IS NULL;

UPDATE fingerprints
SET mdt = mindt(wsq, true)
WHERE mdt IS NULL;
```

Compare two templates with `bz_match`. Its integer result is a similarity score; the application must choose and validate an appropriate threshold for its own data:

```sql
SELECT bz_match(probe.mdt, candidate.mdt) AS score
FROM fingerprints AS probe
CROSS JOIN fingerprints AS candidate
WHERE probe.id = '101_1'
  AND candidate.id = '101_6';
```

For one-to-many identification, score candidate rows and order or filter them explicitly. The upstream example performs a sequential scan, so benchmark the candidate population before placing this operation on a request path.

### Function Index

- `cwsq(bytea, real, int, int, int, int) -> bytea` encodes a raw fingerprint image as WSQ.
- `nfiq(bytea) -> int` computes an NFIQ image-quality score.
- `mindt(bytea)` and `mindt(bytea, boolean) -> bytea` extract a binary minutiae template.
- `mdt2text(bytea) -> text` renders a template as text, and `mdt_mins(bytea) -> int` counts its minutiae.
- `bz_match(bytea, bytea) -> int` and `bz_match(text, text) -> int` calculate Bozorth3 match scores for binary or text templates.

### Operational Notes

Version `1.2` is relocatable and does not declare a preload requirement. Its pinned build files, however, hard-code NBIS `5.0.0` under `/opt/nbis-5.0.0`, statically link the NBIS archives, pass `-m32`, and retain a PostgreSQL `9.3` in-tree build default. Treat those as substantial portability constraints: rebuild and test against the exact PostgreSQL, compiler, architecture, NBIS, and image-library stack you intend to run.

Fingerprint templates and thresholds are biometric security decisions, not generic constants. Validate false-accept and false-reject behavior with representative data, restrict access to raw images and templates, and avoid presenting a match score as an identity guarantee.
