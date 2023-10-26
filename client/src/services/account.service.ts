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

export async function loginUser(username: string, password: string) {
    const payload = {
        username,
        password,
    }
    const data = await http.post(`${BASE_URL}/login`, payload)
    return data;
}

export async function getStatus() {
    const data = await http.get(`${BASE_URL}/status`)
    return data
}

export async function logout() {
    const data = await http.get(`${BASE_URL}/logout`)
    return data;
}