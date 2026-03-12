


## Usage

> [pgqr: QR code generation for PostgreSQL](https://github.com/AbdulYadi/pgqr)

### Function

```sql
pgqr(t text, correction_level integer, model_number integer, scale integer) RETURNS bytea
```

Parameters:
- `t` -- text to encode into the QR code
- `correction_level` -- error correction level: 0 to 3
- `model_number` -- QR model number: 0 to 2
- `scale` -- pixels per dot (scaling factor)

### Example

Generate a QR code as a monochrome bitmap:

```sql
SELECT pgqr('QR Code with PostgreSQL', 0, 0, 4);
```

The output is a monochrome bitmap image (BMP format) returned as `bytea`, ready for display or storage.
