import http from "./http.service"

const BASE_URL = `${import.meta.env.VITE_BASE_API_URL}/api`

export async function registerUser(username: string, password: string, password_confirm: string) {
    const payload = {
        username,
        password,
        password_confirm,
    }
    const data = await http.post(`${BASE_URL}/register`, payload)
    return data;
}