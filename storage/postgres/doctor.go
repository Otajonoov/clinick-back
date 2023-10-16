package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/storage/repo"
)

type artMedicalRepo struct {
	db *sqlx.DB
}

func NewArtMedical(db *sqlx.DB) repo.MedicalStorageI {
	return &artMedicalRepo{
		db: db,
	}
}

// Doctor...
func (dr *artMedicalRepo) CreateDoctor(ctx context.Context, req *repo.Doctor) (*repo.Doctor, error) {
	var result repo.Doctor
	query := `
		INSERT INTO doctors(
			fullname,
			type,
			about,
			img_url
		) VALUES($1, $2, $3, $4)
		RETURNING 
			id,
			fullname,
			type,
			about,
			img_url
		`
	if err := dr.db.DB.QueryRow(query,
		req.Fullname,
		req.Type,
		req.About,
		req.ImageUrl,
	).Scan(
		&result.ID,
		&result.Fullname,
		&result.Type,
		&result.About,
		&result.ImageUrl,
	); err != nil {
		return &repo.Doctor{}, err
	}

	return &result, nil
}

func (u *artMedicalRepo) UpdateDoctor(ctx context.Context, doctor *repo.Doctor) error {
	query := `
		UPDATE doctors SET
			fullname = $1,
			type = $2, 
			about = $3, 
			img_url = $4
		WHERE id = $5
	`

	res, err := u.db.Exec(query,
		doctor.Fullname,
		doctor.Type,
		doctor.About,
		doctor.ImageUrl,
		doctor.ID)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (u *artMedicalRepo) GetAllDoctors(ctx context.Context) (*repo.AllDoctors, error) {
	query := `
		SELECT 
			id,
			fullname,
			type,
			about,
			img_url
		FROM doctors 
	`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	var doctors repo.AllDoctors
	doctors.Doctors = make([]*repo.Doctor, 0)
	for rows.Next() {
		var doctor repo.Doctor
		err = rows.Scan(
			&doctor.ID,
			&doctor.Fullname,
			&doctor.Type,
			&doctor.About,
			&doctor.ImageUrl,
		)
		if err != nil {
			return nil, err
		}
		doctors.Doctors = append(doctors.Doctors, &doctor)
	}

	return &doctors, nil
}

func (u *artMedicalRepo) DeleteDoctor(ctx context.Context, id int64) error {
	query := `
		DELETE FROM doctors WHERE id = $1
	`
	res, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (dr *artMedicalRepo) GetDoctorById(ctx context.Context, id int64) (*repo.Doctor, error) {
	var result repo.Doctor

	query := `
		SELECT
			id,
			fullname,
			type,
			about,
			img_url
		FROM doctors
		WHERE id = $1
	`
	err := dr.db.DB.QueryRow(query, id).Scan(
		&result.ID,
		&result.Fullname,
		&result.Type,
		&result.About,
		&result.ImageUrl,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &repo.Doctor{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &repo.Doctor{}, err
	}

	return &result, nil
}

// Services...
func (dr *artMedicalRepo) CreateService(ctx context.Context, req *repo.Service) (*repo.Service, error) {
	var result repo.Service

	query := `
		INSERT INTO services(
			servicename,
			about,
			img_url
		) VALUES($1, $2, $3)
		RETURNING 
			id,
			serviceName,
			about,
			img_url
		`
	if err := dr.db.DB.QueryRow(query,
		req.ServiceName,
		req.About,
		req.ImageUrl,
	).Scan(
		&result.ID,
		&result.ServiceName,
		&result.About,
		&result.ImageUrl,
	); err != nil {
		return &repo.Service{}, err
	}

	return &result, nil
}

func (u *artMedicalRepo) UpdateService(ctx context.Context, req *repo.Service) error {
	query := `
		UPDATE services SET
			servicename = $1,
			about = $2, 
			img_url = $3
		WHERE id = $4
	`
	res, err := u.db.Exec(query,
		req.ServiceName,
		req.About,
		req.ImageUrl,
		req.ID,
	)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (u *artMedicalRepo) GetAllServices(ctx context.Context) (*repo.AllServices, error) {
	query := `
		SELECT 
			id,
			servicename,
			about,
			img_url
		FROM services 
	`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	var services repo.AllServices
	services.Services = make([]*repo.Service, 0)
	for rows.Next() {
		var service repo.Service
		err = rows.Scan(
			&service.ID,
			&service.ServiceName,
			&service.About,
			&service.ImageUrl,
		)
		if err != nil {
			return nil, err
		}
		services.Services = append(services.Services, &service)
	}

	return &services, nil
}

func (u *artMedicalRepo) DeleteService(ctx context.Context, id int64) error {
	query := `
		DELETE FROM services WHERE id = $1
	`
	res, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (dr *artMedicalRepo) GetServiceById(ctx context.Context, id int64) (*repo.Service, error) {
	var result repo.Service

	query := `
		SELECT
			id,
			servicename,
			about,
			img_url
		FROM services
		WHERE id = $1
	`
	err := dr.db.DB.QueryRow(query, id).Scan(
		&result.ID,
		&result.ServiceName,
		&result.About,
		&result.ImageUrl,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &repo.Service{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &repo.Service{}, err
	}

	return &result, nil
}

// Customer
func (dr *artMedicalRepo) CreateCustomer(ctx context.Context, req *repo.Customer) (*repo.Customer, error) {
	var result repo.Customer
	query := `
		INSERT INTO customers(
			fullname,
			stars,
			about,
			img_url,
			created_at
		) VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP)
		RETURNING 
			id,
			fullname,
			stars,
			about,
			img_url,
			created_at
		`
	if err := dr.db.DB.QueryRow(query,
		req.Fullname,
		req.Stars,
		req.About,
		req.ImageUrl,
	).Scan(
		&result.ID,
		&result.Fullname,
		&result.Stars,
		&result.About,
		&result.ImageUrl,
		&result.CreatedAt,
	); err != nil {
		return &repo.Customer{}, err
	}

	return &result, nil
}

func (u *artMedicalRepo) UpdateCustomer(ctx context.Context, customer *repo.Customer) error {
	query := `
		UPDATE customers SET
			fullname = $1,
			stars = $2, 
			about = $3, 
			img_url = $4
		WHERE id = $5
	`

	res, err := u.db.Exec(query,
		customer.Fullname,
		customer.Stars,
		customer.About,
		customer.ImageUrl,
		customer.ID)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (u *artMedicalRepo) GetAllCustomers(req repo.CustomersFindReq) (*repo.AllCustomers, error) {
	result := repo.AllCustomers{
		Customers: make([]*repo.Customer, 0),
	}

	offset := (req.Page - 1) * req.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", req.Limit, offset)

	query := `
	SELECT 
		id,
    	fullname,
    	stars,
    	about,
    	img_url		
	FROM customers ORDER BY created_at DESC
	` + limit

	rows, err := u.db.Query(query)
	if err != nil {
		return &repo.AllCustomers{}, err
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		count++
		temp := repo.Customer{}
		err = rows.Scan(
			&temp.ID,
			&temp.Fullname,
			&temp.Stars,
			&temp.About,
			&temp.ImageUrl,
		)
		if err != nil {
			return &repo.AllCustomers{}, err
		}
		result.Customers = append(result.Customers, &temp)
	}
	result.Count = count

	return &result, nil
}

func (u *artMedicalRepo) DeleteCustomer(ctx context.Context, id int64) error {
	query := `
		DELETE FROM customers WHERE id = $1
	`
	res, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return err
}

func (dr *artMedicalRepo) GetCustomerById(ctx context.Context, id int64) (*repo.Customer, error) {
	var result repo.Customer

	query := `
		SELECT
			id,
    		fullname,
    		stars,
    		about,
    		img_url		
		FROM customers
		WHERE id = $1
	`
	err := dr.db.DB.QueryRow(query, id).Scan(
		&result.ID,
		&result.Fullname,
		&result.Stars,
		&result.About,
		&result.ImageUrl,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &repo.Customer{}, err
	} else if err != nil {
		log.Fatalln(err.Error())
		return &repo.Customer{}, err
	}

	return &result, nil
}

// Users
func (u *artMedicalRepo) CreateNewUser(ctx context.Context, user *repo.UserReq) (*repo.User, error) {
	var res repo.User
	query := `
		INSERT INTO users(
			username,
			password
		) VALUES ($1, $2)
		RETURNING id, username
	`
	if err := u.db.QueryRow(query, user.Username, user.Password).Scan(&res.ID, &res.Username); err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *artMedicalRepo) GetUserByUsername(ctx context.Context, username string) (*repo.User, error) {
	query := `
		SELECT 
			id,
			username,
			password
		FROM users WHERE username = $1
	`
	var user repo.User
	if err := u.db.QueryRowx(query, username).Scan(&user.ID, &user.Username, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}
