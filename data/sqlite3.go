package data

import (
	"fmt"

	"github.com/bulwark-crypto/bwk-dash/model"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	client *sqlx.DB
}

// Close will close the client.
func (db *DB) Close() {
	db.client.Close()
}

// GetAll will return the last 60 of the information in the database.
func (db *DB) GetAll() (info *[]model.InfoResponse, err error) {
	db.client.Select(info, "SELECT * FROM entries ORDER BY id DESC LIMIT 60")
	return
}

// Last will return the latest entry into the database.
func (db *DB) Last() (info *model.InfoResponse, err error) {
	row := db.client.QueryRowx("SELECT * FROM entries ORDER BY id DESC LIMIT 1")
	info = new(model.InfoResponse)
	err = row.StructScan(info)
	fmt.Println(info)
	return
}

// Save will add the info into the database.
func (db *DB) Save(info *model.InfoResponse) (err error) {
	stmt := `
        INSERT INTO entries (
            blocks, connections, country, difficulty, donationAddress,
            ip, incomingData, maxFee, maxMemory, midFee, minFee, network,
            networkHashPS, outgoingData, protocol, rank, status, stakingStatus,
            subversion, totalData, transactions, usedMemory, version
        )
        VALUES (
            :blocks, :connections, :country, :difficulty, :donationAddress,
            :ip, :incomingData, :maxFee, :maxMemory, :midFee, :minFee, :network,
            :networkHashPS, :outgoingData, :protocol, :rank, :status, :stakingStatus,
            :subversion, :totalData, :transactions, :usedMemory, :version
        )
    `
	_, err = db.client.NamedExec(stmt, info)
	return
}

// Setup will try to create the table structure.
func (db *DB) Setup() (err error) {
	stmt := `
        CREATE TABLE IF NOT EXISTS entries (
            id INTEGER NOT NULL PRIMARY KEY,
            blocks INTEGER NOT NULL,
            connections INTEGER NOT NULL,
            country TEXT NOT NULL,
            difficulty REAL NOT NULL,
            donationAddress TEXT NOT NULL,
            ip TEXT NOT NULL,
            incomingData REAL NOT NULL,
            maxFee REAL NOT NULL,
            maxMemory REAL NOT NULL,
            midFee REAL NOT NULL,
            minFee REAL NOT NULL,
            network TEXT NOT NULL,
            networkHashPS REAL NOT NULL,
            outgoingData REAL NOT NULL,
            protocol INTEGER NOT NULL,
            rank INTEGER,
            status TEXT NOT NULL,
            stakingStatus TEXT NOT NULL,
            subversion TEXT NOT NULL,
            totalData REAL NOT NULL,
            transactions INTEGER NOT NULL,
            usedMemory REAL NOT NULL,
            version INTEGER NOT NULL
        )
    `
	_, err = db.client.Exec(stmt)
	return
}

// NewSQL will setup a new database connection.
func NewSQL(dsn string) (db *DB, err error) {
	db = new(DB)
	db.client, err = sqlx.Connect("sqlite3", dsn)
	return
}
