## Usage

Sources:

- [Upstream installation and basic usage](https://github.com/no0p/plotpg/blob/20b7b3515c0de5c214d1ebb32c9063d9724386b0/README.md)
- [Query plotting implementation](https://github.com/no0p/plotpg/blob/20b7b3515c0de5c214d1ebb32c9063d9724386b0/plotpg.c)
- [Direct gnuplot implementation](https://github.com/no0p/plotpg/blob/20b7b3515c0de5c214d1ebb32c9063d9724386b0/gnuplot.c)

`plotpg` renders PostgreSQL query results through the server host's `gnuplot` executable. The `plot` function executes a supplied query through SPI, writes temporary data and scripts, invokes the external program, and returns its output.

```sql
CREATE EXTENSION plotpg;
SELECT plot('SELECT price, weight FROM widgets');
```

The implementation executes caller-supplied SQL and gnuplot commands and uses predictable process-based names under `/tmp`. Treat its functions as operating-system command capabilities: revoke execution from `PUBLIC`, grant only to trusted administrators, constrain database and OS permissions, and never pass user input. The project is an old experiment; prefer rendering query results in an application or maintained visualization service.
