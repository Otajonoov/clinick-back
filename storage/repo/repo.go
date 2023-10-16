package repo

import (
	"context"
	"time"
)

type MedicalStorageI interface {
	// doctor
	GetAllDoctors(ctx context.Context) (*AllDoctors, error)
	CreateDoctor(ctx context.Context, doctor *Doctor) (*Doctor, error)
	DeleteDoctor(ctx context.Context, id int64) error
	UpdateDoctor(ctx context.Context, doctor *Doctor) error
	GetDoctorById(ctx context.Context, id int64) (*Doctor, error)

	// Service
	GetAllServices(ctx context.Context) (*AllServices, error)
	CreateService(ctx context.Context, service *Service) (*Service, error)
	DeleteService(ctx context.Context, id int64) error
	UpdateService(ctx context.Context, service *Service) error
	GetServiceById(ctx context.Context, id int64) (*Service, error)

	// Customer
	GetAllCustomers(ctx CustomersFindReq) (*AllCustomers, error)
	CreateCustomer(ctx context.Context, doctor *Customer) (*Customer, error)
	DeleteCustomer(ctx context.Context, id int64) error
	UpdateCustomer(ctx context.Context, customer *Customer) error
	GetCustomerById(ctx context.Context, id int64) (*Customer, error)

	// users
	GetUserByUsername(ctx context.Context, username string) (user *User, err error)
	CreateNewUser(ctx context.Context, user *UserReq) (*User, error)
}

type Doctor struct {
	ID       int64
	Fullname string
	Type     string
	About    string
	ImageUrl string
}

type AllDoctors struct {
	Doctors []*Doctor
}

type User struct {
	ID        int64
	Username  string
	Password  string
	CreatedAt time.Time
}

type UserReq struct {
	Username string
	Password string
}

type Service struct {
	ID          int64
	ServiceName string
	About       string
	ImageUrl    string
}

type AllServices struct {
	Services []*Service
}

type Customer struct {
	ID        int64
	Fullname  string
	Stars     string
	About     string
	ImageUrl  string
	CreatedAt string
}

type AllCustomers struct {
	Customers []*Customer
	Count     int64
}

type CustomersFindReq struct {
	Limit int64
	Page  int64
}

type DoctorId struct {
	Id int64
}
