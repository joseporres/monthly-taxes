package repository

import (
	"backend/encryption"
	"backend/global"
	"backend/internal/entity"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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
	qryInsertLog = `
		insert into LOGS ( method, endpoint, status_code, execution_time, timestamp)	
		values (?, ?, ?);`
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

func insertLog(ctx context.Context, method string, endpoint string, status_code int, execution_time float64, timestamp string, r *repo) error {
	_, err := r.db.ExecContext(ctx, qryInsertLog, method, endpoint, status_code, execution_time, timestamp)
	return err
}

func (r *repo) GetMonthlyTaxes(ctx context.Context, dni string) (string, error) {
	// Log data
	method := "GET"
	endpoint := "/monthly-taxes"
	status_code := 200
	start_time := time.Now()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	// Construct the API request URL.
	url := fmt.Sprintf("http://api.aciertaperu.com/app4/v2/reniec?dni=%s&token=%s", dni, global.TOKEN)

	resp, err := http.Get(url)
	if err != nil {
		status_code = 500
		end_time := time.Now()
		execution_time := end_time.Sub(start_time).Seconds()
		err = insertLog(ctx, method, endpoint, status_code, execution_time, timestamp, r)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		status_code = resp.StatusCode
		end_time := time.Now()
		execution_time := end_time.Sub(start_time).Seconds()
		err = insertLog(ctx, method, endpoint, status_code, execution_time, timestamp, r)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		status_code = 500
		end_time := time.Now()
		execution_time := end_time.Sub(start_time).Seconds()
		err = insertLog(ctx, method, endpoint, status_code, execution_time, timestamp, r)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	status_code = resp.StatusCode
	end_time := time.Now()
	execution_time := end_time.Sub(start_time).Seconds()
	err = insertLog(ctx, method, endpoint, status_code, execution_time, timestamp, r)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (r *repo) Logout(ctx context.Context, token string) (string, error) {
	encryption.RevokeToken(token)
	return "Logout", nil
}
