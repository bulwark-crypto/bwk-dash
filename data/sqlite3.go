package data

import (
	"github.com/dustinengle/bwk-dash/handler"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sqlx.DB
}

// GetAll will return all of the information in the database.
func (db *DB) GetAll() (info []*handler.InfoResponse, err error) {
	db.Select(&info, "SELECT * FROM entries ORDER BY id DESC")
	return
}

// Last will return the latest entry into the database.
func (db *DB) Last() (info *handler.InfoResponse, err error) {
	db.Get(&info, "SELECT * FROM entries ORDER BY id DESC")
	return
}

// Save will add the info into the database.
func (db *DB) Save(info *handler.InfoResponse) (err error) {
	stmt := `
        INSERT INTO entries (
            blocks, connections, country, difficulty, donationAddress,
            ip, incomingData, maxFee, maxMemory, midFee, minFee, network,
            networkHashPS, outgoingData, protocol, rank, status, stakingStatus,
            subversion, totalData, transactions, usedMemory, version
        )
        VALUES (
            blocks, :connections, :country, :difficulty, :donationAddress,
            :ip, :incomingData, :maxFee, :maxMemory, :midFee, :minFee, :network,
            :networkHashPS, :outgoingData, :protocol, :rank, :status, :stakingStatus,
            :subversion, :totalData, :transactions, :usedMemory, :version
        )
    `
	_, err = db.NamedQuery(stmt, info)
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
            difficulty FLOAT NOT NULL,
            donationAddress TEXT NOT NULL,
            ip TEXT NOT NULL,
            incomingData FLOAT NOT NULL,
            maxFee FLOAT NOT NULL,
            maxMemory FLOAT NOT NULL,
            midFee FLOAT NOT NULL,
            minFee FLOAT NOT NULL,
            network TEXT NOT NULL,
            networkHashPS FLOAT NOT NULL,
            outgoingData FLOAT NOT NULL,
            protocol INTEGER NOT NULL,
            rank INTEGER,
            status TEXT NOT NULL,
            stakingStatus TEXT NOT NULL,
            totalData FLOAT NOT NULL,
            transactions INTEGER NOT NULL,
            usedMemory FLOAT NOT NULL,
            version INTEGER NOT NULL
        )
    `
	_, err = db.Exec(stmt)
	return
}

// NewSQL will setup a new database connection.
func NewSQL(dsn string) (db *DB, err error) {
	db = new(DB)
	db.DB, err = sqlx.Connect("sqlite3", dsn)
	return
}
