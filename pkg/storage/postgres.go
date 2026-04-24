package storage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresStorage struct {
	db *sql.DB
}

func connectToDB(conn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", conn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db, nil
}

func NewPostgresStorage() *PostgresStorage {

	db, err := connectToDB(os.Getenv("DB_URI"))
	if err != nil {
		panic(err)
	}

	
	return &PostgresStorage{
		db: db,
	}
}

func (pg *PostgresStorage) GetSecretValue(tenantID int64, secretKey string, version *int) (*SecretEntry, error) {
	query := "SELECT id, tenant_id, secret_key, secret_value, created_at, version, dek_version FROM SECRETS WHERE tenant_id = $1 AND secret_key = $2"
	args := []any{tenantID, secretKey}

	if version != nil {
		query += " AND version = $3"
		args = append(args, *version)
	} else {
		query += " ORDER BY version DESC LIMIT 1"
	}

	secret := SecretEntry{}
	err := pg.db.QueryRow(query, args...).Scan(&secret.ID, &secret.TenantID, &secret.SecretKey, &secret.SecretValue, &secret.CreatedAt, &secret.Version, &secret.DEKVersion)
	if err != nil {
		return nil, err
	}

	return &secret, nil
}

// Set DEKVersion to nil for latest DEK entry
func (pg *PostgresStorage) GetDEK(tenantID int64, DEKVersion *int) (*DEKDTO, error) {
	query := "SELECT dek, version FROM DEKS WHERE tenant_id = $1"
	args := []any{tenantID}

	if DEKVersion != nil {
		query += " AND version = $2"
		args = append(args, *DEKVersion)
	} else {
		query += " ORDER BY version DESC LIMIT 1"
	}

	dek := DEKDTO{TenantID: tenantID}
	err := pg.db.QueryRow(query, args...).Scan(&dek.DEK, &dek.Version)
	if err != nil {
		return nil, err
	}

	return &dek, nil
}
func (pg *PostgresStorage) ValidateAuth(tenantID int64, apiKey string) (bool, error) {
	var count int
	err := pg.db.QueryRow("SELECT COUNT(id) FROM AUTH WHERE tenant_id = $1 AND api_key = $2", tenantID, apiKey).Scan(&count)
	if err != nil {
		return false, err
	}
	log.Printf("api-key validation count: %d",count)
	if count != 1 {
		return false, nil
	}
	return true, nil
}

func (pg *PostgresStorage) AddSecret(tenantID int64, secretKey string, secretValue []byte, DEKVersion int) (int, error) {
	tx, txErr := pg.db.Begin()

	if txErr != nil {
		return 0, txErr
	}

	var version int

	err := tx.QueryRow(`
    SELECT COALESCE(MAX(version), 0) + 1
    FROM secrets
    WHERE tenant_id = $1 AND secret_key = $2
    FOR UPDATE
	`, tenantID, secretKey).Scan(&version)

	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(`
    INSERT INTO secrets (tenant_id, secret_key, secret_value, dek_version, version)
    VALUES ($1, $2, $3, $4, $5)
	`, tenantID, secretKey, secretValue, DEKVersion, version)

	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return version, nil
}
