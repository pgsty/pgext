## Usage

Sources:

- [Upstream README](https://github.com/drotiro/postpic/blob/98307c665f4f83aa0cd89812e9640f737f1df0ea/README.md)
- [Extension control file](https://github.com/drotiro/postpic/blob/98307c665f4f83aa0cd89812e9640f737f1df0ea/server/postpic.control)
- [SQL installation script](https://github.com/drotiro/postpic/blob/98307c665f4f83aa0cd89812e9640f737f1df0ea/server/sql/postpic.sql)
- [C implementation](https://github.com/drotiro/postpic/blob/98307c665f4f83aa0cd89812e9640f737f1df0ea/server/src/postpic.c)

`postpic` version `0.9.1` adds an `image` type backed by GraphicsMagick. SQL functions can load images from `bytea` or a large object, read width, height, date and EXIF-style attributes, and produce thumbnails, crops, rotations, drawings, and contact sheets.

### Example

```sql
CREATE EXTENSION postpic;
CREATE TABLE pictures (id bigint GENERATED ALWAYS AS IDENTITY, body image);
INSERT INTO pictures(body) VALUES (image_from_bytea(pg_read_binary_file('/srv/import/photo.jpg')));
SELECT id, width(body), height(body), colorspace(body) FROM pictures;
SELECT thumbnail(body, 320) FROM pictures WHERE id = 1;
```

The example reads a server-side path and therefore requires privileged file access; applications should instead pass approved bytes through a controlled ingestion path. Image decoding runs native GraphicsMagick code inside a PostgreSQL backend, so malformed media can consume memory/CPU or hit decoder vulnerabilities. This is old code with no current PostgreSQL or GraphicsMagick compatibility statement. Restrict execution, cap input size, keep codecs patched, and evaluate only in an isolated environment.
