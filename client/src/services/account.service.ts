import type { IClient } from "@/types";
import http from "./http.service"

const BASE_URL = `${import.meta.env.VITE_BASE_API_URL}/api`

interface ILoginResponse {
    "token": string,
    "user": IClient,
}

export async function registerUser(username: string, password: string, password_confirm: string): Promise<IClient> {
    const payload = {
        username,
        password,
        password_confirm,
    }
    const { data } = await http.post(`${BASE_URL}/register`, payload)
    return data;
}

export async function loginUser(username: string, password: string): Promise<ILoginResponse> {
    const payload = {
        username,
        password,
    }
    const { data } = await http.post(`${BASE_URL}/login`, payload)
    return data;
}

export async function getStatus(): Promise<IClient> {
    const { data } = await http.get(`${BASE_URL}/status`)
    return data
}

export async function logout() {
    const { data } = await http.get(`${BASE_URL}/logout`)
    return data;
}