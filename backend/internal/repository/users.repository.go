package repository

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"backend/internal/entity"
)

const (
	qryInsertUser = `
		insert into USERS (email, name, password)
		values (?, ?, ?);`

	qryGetUserByEmail = `
		select
			id,
			email,
			name,
			password
		from USERS
		where email = ?;`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, qryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repo) GetMonthlyTaxes(ctx context.Context, dni string) (string, error) {
	token := "ASPJGDSP4SD783H375S"
	// Construct the URL with the provided DNI and token.
	url := fmt.Sprintf("http://api.aciertaperu.com/app4/v2/reniec?dni=%s&token=%s", dni, token)

	// Perform an HTTP GET request to the API.
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Assuming the API returns a string response, you can return it here.
	return string(body), nil
}
