package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/types"
)

type DriverRepositoryInterface interface {
	CreateDriver(ctx context.Context, driver *types.Driver) error
	ReadDriver(ctx context.Context, cnh *string) (*types.Driver, error)
	UpdateDriver(ctx context.Context) error
	DeleteDriver(ctx context.Context, cnh *string) error
	AuthDriver(ctx context.Context, driver *types.Driver) (*types.Driver, error)
	ReadInvite(ctx context.Context, invite_id *int) (*types.Invite, error)
	ReadAllInvites(ctx context.Context, cnh *string) ([]types.Invite, error)
	UpdateInvite(ctx context.Context, invite_id *int) error
	DeleteInvite(ctx context.Context, invite_id *int) error
	GetWorkplaces(ctx context.Context, cnh *string) ([]types.School, error)
	CreateEmployee(ctx context.Context, invite *types.Invite) error
}

type DriverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
}

func (d *DriverRepository) CreateDriver(ctx context.Context, driver *types.Driver) error {
	sqlQuery := `INSERT INTO drivers (name, cpf, email, password, cnh, qrcode, street, number, complement, zip) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := d.db.Exec(sqlQuery, driver.Name, driver.CPF, driver.Email, driver.Password, driver.CNH, driver.QrCode, driver.Street, driver.Number, driver.Complement, driver.ZIP)
	return err
}

func (d *DriverRepository) ReadDriver(ctx context.Context, cnh *string) (*types.Driver, error) {
	sqlQuery := `SELECT id, name, cpf, cnh, qrcode, email, street, number, zip, complement FROM drivers WHERE cnh = $1 LIMIT 1`
	var driver types.Driver
	err := d.db.QueryRow(sqlQuery, *cnh).Scan(
		&driver.ID,
		&driver.Name,
		&driver.CPF,
		&driver.CNH,
		&driver.QrCode,
		&driver.Email,
		&driver.Street,
		&driver.Number,
		&driver.ZIP,
		&driver.Complement,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &driver, nil
}

func (d *DriverRepository) UpdateDriver(ctx context.Context) error {
	return nil
}

func (d *DriverRepository) DeleteDriver(ctx context.Context, cnh *string) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec("DELETE FROM drivers WHERE cnh = $1", *cnh)
	return err
}

func (d *DriverRepository) AuthDriver(ctx context.Context, driver *types.Driver) (*types.Driver, error) {
	sqlQuery := `SELECT id, name, cpf, cnh, email, qrcode, password FROM drivers WHERE email = $1 LIMIT 1`
	var driverData types.Driver
	err := d.db.QueryRow(sqlQuery, driver.Email).Scan(
		&driverData.ID,
		&driverData.Name,
		&driverData.CPF,
		&driverData.CNH,
		&driverData.Email,
		&driverData.QrCode,
		&driverData.Password,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	match := driverData.Password == driver.Password
	if !match {
		return nil, fmt.Errorf("email or password wrong")
	}
	driverData.Password = ""
	return &driverData, nil
}

func (d *DriverRepository) ReadInvite(ctx context.Context, invite_id *int) (*types.Invite, error) {
	sqlQuery := `SELECT invite_id, requester, school, email_school, guest, driver, email_driver, status FROM invites WHERE invite_id = $1 LIMIT 1`
	var invite types.Invite
	err := d.db.QueryRow(sqlQuery, *invite_id).Scan(
		&invite.ID,
		&invite.School.CNPJ,
		&invite.School.Name,
		&invite.School.Email,
		&invite.Driver.CNH,
		&invite.Driver.Name,
		&invite.Driver.Email,
		&invite.Status,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &invite, nil
}

func (d *DriverRepository) ReadAllInvites(ctx context.Context, cnh *string) ([]types.Invite, error) {
	sqlQuery := `SELECT invite_id, school, requester, email_school, driver, guest, email_driver, status FROM invites WHERE status = 'pending' AND guest = $1`

	rows, err := d.db.Query(sqlQuery, *cnh)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invites []types.Invite

	for rows.Next() {
		var invite types.Invite
		err := rows.Scan(&invite.ID, &invite.School.Name, &invite.School.CNPJ, &invite.School.Email, &invite.Driver.Name, &invite.Driver.CNH, &invite.Driver.Email, &invite.Status)
		if err != nil {
			return nil, err
		}
		invites = append(invites, invite)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return invites, nil

}

func (d *DriverRepository) UpdateInvite(ctx context.Context, invite_id *int) error {
	sqlQuery := `UPDATE invites SET status = 'accepted' WHERE invite_id = $1`
	_, err := d.db.Exec(sqlQuery, invite_id)

	return err
}

func (d *DriverRepository) DeleteInvite(ctx context.Context, invite_id *int) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec("DELETE FROM invites WHERE invite_id = $1", invite_id)
	return err
}

func (d *DriverRepository) CreateEmployee(ctx context.Context, invite *types.Invite) error {
	sqlQuery := `INSERT INTO schools_drivers (name_school, school, email_school, name_driver, driver, email_driver) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := d.db.Exec(sqlQuery, invite.School.Name, invite.School.CNPJ, invite.School.Email, invite.Driver.Name, invite.Driver.CNH, invite.Driver.Email)
	return err
}

func (d *DriverRepository) GetWorkplaces(ctx context.Context, cnh *string) ([]types.School, error) {
	sqlQuery := `SELECT school_name, school, school_email FROM schools_drivers WHERE driver = $1`

	rows, err := d.db.Query(sqlQuery, *cnh)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schools []types.School

	for rows.Next() {
		var school types.School
		err := rows.Scan(&school.CNPJ)
		if err != nil {
			return nil, err
		}
		schools = append(schools, school)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schools, nil

}
