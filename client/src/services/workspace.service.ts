import type { IWorkspace, IWorkspaceCreateRequest } from "@/types";
import http from "./http.service"

const BASE_URL = `${import.meta.env.VITE_BASE_API_URL}/api`

export async function createWorkspace(payload: IWorkspaceCreateRequest): Promise<IWorkspace> {
    const { data } = await http.post(`${BASE_URL}/workspace`, payload)
    return data;
}

export async function getWorkspaces(): Promise<IWorkspace[]> {
    const { data } = await http.get(`${BASE_URL}/workspace`)
    return data;
}

export async function updateWorkspace(workspaceId: string, name: string): Promise<void> {
    const payload = {
        name
    }
    await http.put(`${BASE_URL}/workspace/${workspaceId}`, payload)
}

export async function deleteWorkspace(workspaceId: string): Promise<void> {
    await http.delete(`${BASE_URL}/workspace/${workspaceId}`)
}