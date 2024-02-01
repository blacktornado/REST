package model

type Customer struct {
	ID         uint64 `json:"id"`
	State      string `json:"state"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Contactnum uint64 `json:"contactnum"`
	Gender     string `json:"gender"`
	Age        uint64 `json:"age"`
	Status     uint64 `json:"status"`
}

func GetAllCustomer() ([]Customer, error) {
	var customers []Customer
	query := "select id, state, name, email, contactnum, gender, age, status from `cook-accounts`"
	rows, err := DB.Query(query)
	if err != nil {
		return customers, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, age, contactnum, status uint64
		var name, state, email, gender string
		err := rows.Scan(&id, &state, &name, &email, &contactnum, &gender, &age, &status)
		if err != nil {
			return customers, err
		}
		customer := Customer{
			ID:         id,
			Name:       name,
			State:      state,
			Contactnum: contactnum,
			Email:      email,
			Gender:     gender,
			Age:        age,
			Status:     status,
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
