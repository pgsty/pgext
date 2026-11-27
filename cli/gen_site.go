package cli

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
)

// ExportSiteCSVs produces the inputs consumed by Hugo's resources.Get.
// Keep their order and PostgreSQL CSV representation aligned with make dump.
func ExportSiteCSVs(ctx context.Context, dir string) error {
	tx, err := DB.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead, AccessMode: pgx.ReadOnly})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	for _, table := range []struct{ name, order string }{
		{"universe", "id"},
		{"pg", "pg DESC"},
		{"os", "os_major, CASE WHEN os_code IN ('el10','d11','u22') THEN CASE os_arch WHEN 'aarch64' THEN 0 ELSE 1 END ELSE CASE os_arch WHEN 'x86_64' THEN 0 ELSE 1 END END"},
	} {
		var csv bytes.Buffer
		query := fmt.Sprintf("COPY (SELECT * FROM pgext.%s ORDER BY %s) TO STDOUT WITH (FORMAT CSV, HEADER)", table.name, table.order)
		if _, err := tx.Conn().PgConn().CopyTo(ctx, &csv, query); err != nil {
			return fmt.Errorf("export pgext.%s: %w", table.name, err)
		}
		if err := os.WriteFile(filepath.Join(dir, table.name+".csv"), csv.Bytes(), 0644); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}
