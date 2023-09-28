import instance from "../axios";
import router from "../router";


function handleError(error) {
  if (error.response.status === 401) {
    localStorage.removeItem('token');
    router.push('/login');
  }
  throw error;
}

export async function login(data) {
  try {
    return await instance.post('auth/login', data);
  } catch (error) {
    handleError(error);
  }
}
export async function register(data) {
  try {
    return await instance.post('auth/register', data);
  }
  catch (error) {
    handleError(error);
  }
}

export async function monthlyTaxes(dni) {
  try {
    return await instance.get(`users/monthly-expenses/${dni}`);

  } catch (error) {
    handleError(error);
  }
}

export async function logOut() {
  try {
    return await instance.post('auth/logout');
  } catch (error) {
    handleError(error);
  }
}