// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Door represents a row from 'public.doors'.
type Door struct {
	ID             int           `json:"id"`               // id
	CardDoorID     int           `json:"card_door_id"`     // card_door_id
	FullName       string        `json:"full_name"`        // full_name
	DsxDoorNum     sql.NullInt64 `json:"dsx_door_num"`     // dsx_door_num
	DsxPanelNum    sql.NullInt64 `json:"dsx_panel_num"`    // dsx_panel_num
	DsxDoorName    string        `json:"dsx_door_name"`    // dsx_door_name
	BuildingID     int           `json:"building_id"`      // building_id
	FloorNum       int16         `json:"floor_num"`        // floor_num
	ActiveConfigID sql.NullInt64 `json:"active_config_id"` // active_config_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Door exists in the database.
func (d *Door) Exists() bool {
	return d._exists
}

// Deleted provides information if the Door has been deleted from the database.
func (d *Door) Deleted() bool {
	return d._deleted
}

// Insert inserts the Door to the database.
func (d *Door) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if d._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.doors (` +
		`card_door_id, full_name, dsx_door_num, dsx_panel_num, dsx_door_name, building_id, floor_num, active_config_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, d.CardDoorID, d.FullName, d.DsxDoorNum, d.DsxPanelNum, d.DsxDoorName, d.BuildingID, d.FloorNum, d.ActiveConfigID)
	err = db.QueryRow(sqlstr, d.CardDoorID, d.FullName, d.DsxDoorNum, d.DsxPanelNum, d.DsxDoorName, d.BuildingID, d.FloorNum, d.ActiveConfigID).Scan(&d.ID)
	if err != nil {
		return err
	}

	// set existence
	d._exists = true

	return nil
}

// Update updates the Door in the database.
func (d *Door) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !d._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if d._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.doors SET (` +
		`card_door_id, full_name, dsx_door_num, dsx_panel_num, dsx_door_name, building_id, floor_num, active_config_id` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) WHERE id = $9`

	// run query
	XOLog(sqlstr, d.CardDoorID, d.FullName, d.DsxDoorNum, d.DsxPanelNum, d.DsxDoorName, d.BuildingID, d.FloorNum, d.ActiveConfigID, d.ID)
	_, err = db.Exec(sqlstr, d.CardDoorID, d.FullName, d.DsxDoorNum, d.DsxPanelNum, d.DsxDoorName, d.BuildingID, d.FloorNum, d.ActiveConfigID, d.ID)
	return err
}

// Save saves the Door to the database.
func (d *Door) Save(db XODB) error {
	if d.Exists() {
		return d.Update(db)
	}

	return d.Insert(db)
}

// Upsert performs an upsert for Door.
//
// NOTE: PostgreSQL 9.5+ only
func (d *Door) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if d._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.doors (` +
		`id, card_door_id, full_name, dsx_door_num, dsx_panel_num, dsx_door_name, building_id, floor_num, active_config_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, card_door_id, full_name, dsx_door_num, dsx_panel_num, dsx_door_name, building_id, floor_num, active_config_id` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.card_door_id, EXCLUDED.full_name, EXCLUDED.dsx_door_num, EXCLUDED.dsx_panel_num, EXCLUDED.dsx_door_name, EXCLUDED.building_id, EXCLUDED.floor_num, EXCLUDED.active_config_id` +
		`)`

	// run query
	XOLog(sqlstr, d.ID, d.CardDoorID, d.FullName, d.DsxDoorNum, d.DsxPanelNum, d.DsxDoorName, d.BuildingID, d.FloorNum, d.ActiveConfigID)
	_, err = db.Exec(sqlstr, d.ID, d.CardDoorID, d.FullName, d.DsxDoorNum, d.DsxPanelNum, d.DsxDoorName, d.BuildingID, d.FloorNum, d.ActiveConfigID)
	if err != nil {
		return err
	}

	// set existence
	d._exists = true

	return nil
}

// Delete deletes the Door from the database.
func (d *Door) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !d._exists {
		return nil
	}

	// if deleted, bail
	if d._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.doors WHERE id = $1`

	// run query
	XOLog(sqlstr, d.ID)
	_, err = db.Exec(sqlstr, d.ID)
	if err != nil {
		return err
	}

	// set deleted
	d._deleted = true

	return nil
}

// DoorCompiledContent returns the DoorCompiledContent associated with the Door's ActiveConfigID (active_config_id).
//
// Generated from foreign key 'active_config_fk'.
func (d *Door) DoorCompiledContent(db XODB) (*DoorCompiledContent, error) {
	return DoorCompiledContentByID(db, int(d.ActiveConfigID.Int64))
}

// Building returns the Building associated with the Door's BuildingID (building_id).
//
// Generated from foreign key 'doors_building_id_fkey'.
func (d *Door) Building(db XODB) (*Building, error) {
	return BuildingByID(db, d.BuildingID)
}

// DoorByCardDoorID retrieves a row from 'public.doors' as a Door.
//
// Generated from index 'doors_card_door_id_key'.
func DoorByCardDoorID(db XODB, cardDoorID int) (*Door, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, card_door_id, full_name, dsx_door_num, dsx_panel_num, dsx_door_name, building_id, floor_num, active_config_id ` +
		`FROM public.doors ` +
		`WHERE card_door_id = $1`

	// run query
	XOLog(sqlstr, cardDoorID)
	d := Door{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, cardDoorID).Scan(&d.ID, &d.CardDoorID, &d.FullName, &d.DsxDoorNum, &d.DsxPanelNum, &d.DsxDoorName, &d.BuildingID, &d.FloorNum, &d.ActiveConfigID)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// DoorByFullName retrieves a row from 'public.doors' as a Door.
//
// Generated from index 'doors_full_name_key'.
func DoorByFullName(db XODB, fullName string) (*Door, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, card_door_id, full_name, dsx_door_num, dsx_panel_num, dsx_door_name, building_id, floor_num, active_config_id ` +
		`FROM public.doors ` +
		`WHERE full_name = $1`

	// run query
	XOLog(sqlstr, fullName)
	d := Door{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, fullName).Scan(&d.ID, &d.CardDoorID, &d.FullName, &d.DsxDoorNum, &d.DsxPanelNum, &d.DsxDoorName, &d.BuildingID, &d.FloorNum, &d.ActiveConfigID)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// DoorByID retrieves a row from 'public.doors' as a Door.
//
// Generated from index 'doors_pkey'.
func DoorByID(db XODB, id int) (*Door, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, card_door_id, full_name, dsx_door_num, dsx_panel_num, dsx_door_name, building_id, floor_num, active_config_id ` +
		`FROM public.doors ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	d := Door{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&d.ID, &d.CardDoorID, &d.FullName, &d.DsxDoorNum, &d.DsxPanelNum, &d.DsxDoorName, &d.BuildingID, &d.FloorNum, &d.ActiveConfigID)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
