import type { IWorkspace } from "@/types";
import http from "./http.service"

const BASE_URL = `${import.meta.env.VITE_BASE_API_URL}/api`

export async function createWorkspace(name: string): Promise<IWorkspace> {
    const payload = {
        name
    }
    const { data } = await http.post(`${BASE_URL}/workspace`, payload)
    return data;
}