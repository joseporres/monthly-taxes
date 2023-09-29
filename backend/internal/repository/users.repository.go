package repository

import (
	"backend/encryption"
	"backend/global"
	"backend/internal/entity"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
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

	// Construct the API request URL.
	url := fmt.Sprintf("http://api.aciertaperu.com/app4/v2/reniec?dni=%s&token=%s", dni, global.TOKEN)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	return string(body), nil
}

func (r *repo) Logout(ctx context.Context, token string) (string, error) {
	encryption.RevokeToken(token)
	return "Logout", nil
}
