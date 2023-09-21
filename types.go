package weatherSubscription

type User struct {
	TgId int
	Lat  int
	Lon  int
	City string
	Hour int
}

type WeatherApi interface {
	GetForecast(lat, lon int) (string, error)
}

type DbApi interface {
	CreateUser(user *User) error
	DeleteUser(user *User) error
	UpdateUser(user *User) error
	getUserByTgId(id int) (*User, error)
	getUserByCoords(lat, lon int) (*User, error)
	getUserByCity(city string) (*User, error)
	getUserByHour(hour int) (*User, error)
}
