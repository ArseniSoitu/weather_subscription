package weatherSubscription

import "log"

type MockDbApi struct {
}

func (dbApi *MockDbApi) CreateUser(user *User) error {
	log.Println("User with", user.TgId, "telegram id created")

	return nil
}

func (dbApi *MockDbApi) DeleteUser(user *User) error {
	log.Println("User with", user.TgId, "telegram id deleted")

	return nil
}

func (dbApi *MockDbApi) UpdateUser(user *User) error {
	log.Println("User with", user.TgId, "telegram id updated")

	return nil
}

func (dbApi *MockDbApi) getUserByTgId(id int) (*User, error) {
	user := &User{
		TgId: id,
		Lat:  123,
		Lon:  456,
		City: "SomeCity",
		Hour: 12,
	}
	return user, nil
}

func (dbApi *MockDbApi) getUserByCoords(lat, lon int) (*User, error) {
	user := &User{
		TgId: 0xDEADBEEF,
		Lat:  lat,
		Lon:  lon,
		City: "SomeCity",
		Hour: 12,
	}

	return user, nil
}

func (dbApi *MockDbApi) getUserByCity(city string) (*User, error) {
	user := &User{
		TgId: 0xDEADBEEF,
		Lat:  123,
		Lon:  456,
		City: city,
		Hour: 12,
	}

	return user, nil
}

func (dbApi *MockDbApi) getUserByHour(hour int) (*User, error) {
	user := &User{
		TgId: 0xDEADBEEF,
		Lat:  123,
		Lon:  456,
		City: "SomeCity",
		Hour: hour,
	}

	return user, nil
}
