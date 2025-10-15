/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// APTParser handles parsing of APT/DEB repository metadata
type APTParser struct {
	ctx context.Context
	tx  pgx.Tx
}

// NewAPTParser creates a new APT parser
func NewAPTParser(ctx context.Context) *APTParser {
	return &APTParser{ctx: ctx}
}

// ParseRepository parses a single APT repository
func (p *APTParser) ParseRepository(repoID string, data []byte) (int, error) {
	// Parse Packages file content
	packages := p.parsePackages(string(data))
	if len(packages) == 0 {
		return 0, nil
	}

	// Begin transaction
	tx, err := Begin(p.ctx)
	if err != nil {
		return 0, fmt.Errorf("begin transaction: %w", err)
	}

	// Store transaction for use in insertPackages
	p.tx = tx

	// Insert packages in batches
	count := p.insertPackages(repoID, packages)

	// Commit transaction
	if err := tx.Commit(p.ctx); err != nil {
		// Transaction failed, rollback is automatic
		p.tx = nil
		return count, fmt.Errorf("commit transaction: %w", err)
	}

	p.tx = nil
	return count, nil
}

// parsePackages parses APT Packages file content into structured data
func (p *APTParser) parsePackages(content string) []map[string]interface{} {
	var packages []map[string]interface{}

	// Split into individual package records
	records := strings.Split(content, "\n\n")

	for _, record := range records {
		record = strings.TrimSpace(record)
		if record == "" {
			continue
		}

		pkg := p.parsePackageRecord(record)
		if pkg != nil && pkg["Package"] != nil {
			packages = append(packages, pkg)
		}
	}

	return packages
}

// parsePackageRecord parses a single package record from APT format
func (p *APTParser) parsePackageRecord(record string) map[string]interface{} {
	pkg := make(map[string]interface{})
	var currentKey string
	var currentValue []string

	// Process each line
	for _, line := range strings.Split(record, "\n") {
		if line == "" {
			continue
		}

		// Check if this is a continuation line (starts with space)
		if strings.HasPrefix(line, " ") {
			// Continuation of previous field
			if currentKey != "" {
				currentValue = append(currentValue, strings.TrimSpace(line))
			}
		} else {
			// Save previous field if exists
			if currentKey != "" {
				p.saveField(pkg, currentKey, currentValue)
			}

			// Parse new field
			colonIdx := strings.Index(line, ":")
			if colonIdx > 0 {
				currentKey = line[:colonIdx]
				value := strings.TrimSpace(line[colonIdx+1:])
				currentValue = []string{value}
			}
		}
	}

	// Save last field
	if currentKey != "" {
		p.saveField(pkg, currentKey, currentValue)
	}

	// Process special fields
	p.processSpecialFields(pkg)

	return pkg
}

// saveField saves a field value, handling multi-line values
func (p *APTParser) saveField(pkg map[string]interface{}, key string, values []string) {
	if len(values) == 0 {
		pkg[key] = nil
		return
	}

	// Join multi-line values with newlines (especially for Description)
	if key == "Description" || len(values) > 1 {
		pkg[key] = strings.Join(values, "\n")
	} else {
		pkg[key] = values[0]
	}
}

// processSpecialFields handles type conversion for special fields
func (p *APTParser) processSpecialFields(pkg map[string]interface{}) {
	// Convert Size to integer
	if sizeStr, ok := pkg["Size"].(string); ok && sizeStr != "" {
		if size, err := strconv.Atoi(sizeStr); err == nil {
			pkg["Size"] = size
		}
	}

	// Convert Installed-Size from KB to bytes
	if sizeStr, ok := pkg["Installed-Size"].(string); ok && sizeStr != "" {
		if size, err := strconv.Atoi(sizeStr); err == nil {
			pkg["Installed-Size"] = size * 1024
		}
	}
}

// insertPackages inserts parsed packages into the database
func (p *APTParser) insertPackages(repoID string, packages []map[string]interface{}) int {
	if len(packages) == 0 {
		return 0
	}

	batch := &pgx.Batch{}
	count := 0

	// Fixed fields that map to table columns
	fixedFields := []string{
		"Package", "Version", "Architecture", "Size", "Installed-Size",
		"Priority", "Section", "Filename", "SHA256", "SHA1", "MD5sum",
		"Maintainer", "Homepage", "Depends", "Source", "Provides",
		"Recommends", "Suggests", "Conflicts", "Breaks", "Replaces",
		"Enhances", "Pre-Depends", "Build-Ids", "Package-Type",
		"Auto-Built-Package", "Multi-Arch", "Description",
	}

	for _, pkg := range packages {
		// Build values array for fixed fields
		values := make([]interface{}, 0, 30) // Exactly 30 parameters expected
		values = append(values, repoID) // First column is repo

		// Add fixed field values (28 fields)
		for _, field := range fixedFields {
			if val, ok := pkg[field]; ok {
				values = append(values, val)
			} else {
				values = append(values, nil)
			}
		}

		// Collect extra fields as JSON
		extraFields := make(map[string]interface{})
		for key, value := range pkg {
			isFixed := false
			for _, field := range fixedFields {
				if key == field {
					isFixed = true
					break
				}
			}
			if !isFixed && value != nil {
				extraFields[key] = value
			}
		}

		// Add extra fields as JSON
		if len(extraFields) > 0 {
			if jsonData, err := json.Marshal(extraFields); err == nil {
				values = append(values, string(jsonData))
			} else {
				values = append(values, nil)
			}
		} else {
			values = append(values, nil)
		}

		// Queue insert - 30 columns total: 1 repo + 28 fields + 1 extra
		batch.Queue(`
			INSERT INTO pgext.apt (
				repo, package, version, architecture, size, size_install,
				priority, section, filename, sha256, sha1, md5sum,
				maintainer, homepage, depends, source, provides,
				recommends, suggests, conflicts, breaks, replaces,
				enhances, pre_depends, build_ids, package_type,
				auto_built, multi_arch, description, extra
			) VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
				$11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
				$21, $22, $23, $24, $25, $26, $27, $28, $29, $30
			)
		`, values...)

		count++

		// Send batch every 1000 records
		if batch.Len() >= 1000 {
			if p.tx != nil {
				br := p.tx.SendBatch(p.ctx, batch)
				if err := br.Close(); err != nil {
					logrus.Warnf("failed to send batch: %v", err)
				}
			}
			batch = &pgx.Batch{}
		}
	}

	// Send remaining batch
	if batch.Len() > 0 && p.tx != nil {
		br := p.tx.SendBatch(p.ctx, batch)
		if err := br.Close(); err != nil {
			logrus.Warnf("failed to send final batch: %v", err)
		}
	}

	return count
}