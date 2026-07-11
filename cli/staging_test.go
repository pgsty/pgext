package cli

import (
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestStageTableNamesAreUniqueAndQuoted(t *testing.T) {
	first := nextStageTable("apt")
	second := nextStageTable("apt")
	if first == second {
		t.Fatalf("stage table names collided: %s", first)
	}
	for _, name := range []string{first, second} {
		if !strings.HasPrefix(name, `"pgext"."_stage_apt_`) {
			t.Fatalf("unexpected stage table name: %s", name)
		}
	}
}

func TestPipelineAndPublishLocksUseIndependentKeys(t *testing.T) {
	if packagePipelineLockStatement == packageCatalogLockStatement {
		t.Fatal("parse pipeline and catalog publish locks must use independent advisory keys")
	}
	if !strings.Contains(packagePipelineLockStatement, "package_catalog_pipeline") {
		t.Fatalf("unexpected parse pipeline lock statement: %s", packagePipelineLockStatement)
	}
	if strings.Contains(packagePipelineLockStatement, "xact") {
		t.Fatalf("parse pipeline must not hold an idle transaction: %s", packagePipelineLockStatement)
	}
	if !strings.Contains(packageCatalogLockStatement, "package_catalog_publish") {
		t.Fatalf("unexpected catalog publish lock statement: %s", packageCatalogLockStatement)
	}
}

func TestPackageCatalogTransactionsUseReadCommitted(t *testing.T) {
	// The catalog lock protocol relies on statements after a lock wait seeing
	// the latest committed catalog, regardless of the database-wide default.
	if got := packageCatalogTxOptions().IsoLevel; got != pgx.ReadCommitted {
		t.Fatalf("catalog transaction isolation = %q, want %q", got, pgx.ReadCommitted)
	}
}

func TestStagedReloadSQLPreservesPkgAlias(t *testing.T) {
	source := `
TRUNCATE pgext.pkg;
INSERT INTO pgext.pkg (pg) SELECT 17;
UPDATE pgext.pkg SET count = (SELECT count(*) FROM pgext.bin b WHERE b.pg = pkg.pg);
UPDATE pgext.status SET recap_time = now();`

	got, err := stagedReloadSQL(source, `"pgext"."stage_pkg"`, `"pgext"."stage_bin"`)
	if err != nil {
		t.Fatalf("stagedReloadSQL: %v", err)
	}
	for _, want := range []string{
		`TRUNCATE "pgext"."stage_pkg"`,
		`INSERT INTO "pgext"."stage_pkg"`,
		`UPDATE "pgext"."stage_pkg" AS pkg SET`,
		`FROM "pgext"."stage_bin" b`,
		`b.pg = pkg.pg`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("staged SQL missing %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, recapStatusStatement) {
		t.Fatalf("staged SQL updates live recap timestamp:\n%s", got)
	}
}
