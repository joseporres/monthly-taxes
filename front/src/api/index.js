
import axios from 'axios';

export async function login(data) {
  return await axios.post('http://localhost:8080/auth/login', data);
}
export async function register(data) {
  return await axios.post('http://localhost:8080auth/register', data);
}

export async function monthlyTaxes(dni) {
  return await axios.get(`http://localhost:8080/users/monthly-expenses/${dni}`);
}