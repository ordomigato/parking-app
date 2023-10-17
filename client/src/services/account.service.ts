import http from "./http.service"

const BASE_URL = '/api'

export async function registerUser(username: string, password: string, confirm_password: string) {
    const payload = {
        username,
        password,
        confirm_password,
    }
    const data = await http.post(`${BASE_URL}/register`, payload)
    return data;
}