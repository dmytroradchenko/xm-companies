//go:generate go install github.com/google/wire/cmd/wire@latest
//go:generate go run github.com/google/wire/cmd/wire
//go:generate mockery --dir internal/xm-companies/store --name CompaniesRepository --output internal/xm-companies/store/mocks --case underscore
//go:generate mockery --dir internal/xm-companies/store --name UsersRepository --output internal/xm-companies/store/mocks --case underscore
//go:generate mockery --dir internal/xm-companies/store --name Store --output internal/xm-companies/store/mocks --case underscore

package main

import (
	"fmt"
)

func main() {
	s, err := BuildServerCompileTime()
	if err != nil {
		fmt.Println("Server initialization has failed: ", err)
	} else {
		err = s.StartServer()
		if err != nil {
			fmt.Println("Server start has failed: ", err)
		}
	}
}
