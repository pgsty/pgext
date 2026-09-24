## Usage

Sources:

- [README](https://github.com/RustedBytes/pg-fx/blob/46d76607ae9a35d4742990d75f48b332132f14fb/README.md)
- [Control file](https://github.com/RustedBytes/pg-fx/blob/46d76607ae9a35d4742990d75f48b332132f14fb/pg_fx.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-fx/blob/46d76607ae9a35d4742990d75f48b332132f14fb/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-fx/blob/46d76607ae9a35d4742990d75f48b332132f14fb/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-fx/blob/46d76607ae9a35d4742990d75f48b332132f14fb/docs/SECURITY.md)

`pg_fx` stores externally supplied FX rates, calculates exact prices, and creates expiring quotes on PostgreSQL 14–18. Applications collect the market data; the extension does not access market feeds or move balances.

### Core Workflow

```sql
CREATE EXTENSION pg_fx;
SELECT fx_source_upsert('primary_bank', source_priority => 10,
                        source_max_age => interval '30 seconds');
SELECT fx_rate_insert(source => 'primary_bank', rate_pair => 'USD/EUR',
  rate_bid => 0.8510, rate_ask => 0.8520,
  rate_observed_at => clock_timestamp(), rate_volume => 100000);
SELECT fx_bid('USD/EUR'), fx_ask('USD/EUR'), fx_mid('USD/EUR');
SELECT fx_vwap('USD/EUR'), fx_weighted_median('USD/EUR');
```

### Pricing and Quotes

`fx_rule_create` configures segment and amount-dependent markup and fees. `fx_create_quote` records a time-limited quote; `fx_execute_quote` changes its state but does not post money to accounts. Keep any application ledger posting and quote transition in the same transaction. Source priority, staleness thresholds, asset precision, and rounding affect the returned price.

`fx_enable_pg_money` and `fx_enable_pg_cryptocurrency` install optional typed adapters after those extensions are present. The separate ledger integration targets the RustedBytes project and must not be confused with another project sharing its extension name.

### Privileges and Maintenance

No preload is required. The install schema may be chosen at creation, but the extension is not relocatable afterward. Core use has no hard companion-extension dependency. Review the security document and grant only required source, rate, rule, and quote operations; do not let untrusted callers publish authoritative rates.
