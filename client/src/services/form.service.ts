import type { IForm, IFormCreateRequest } from "@/types";
import http from "./http.service"

const BASE_URL = `${import.meta.env.VITE_BASE_API_URL}/api`

export async function createForm(workspaceId: string, payload: IFormCreateRequest): Promise<IForm> {
    const { data } = await http.post(`${BASE_URL}/workspace/${workspaceId}/form`, payload)
    return data;
}

export async function getForms(workspaceId: string): Promise<IForm[]> {
    const { data } = await http.get(`${BASE_URL}/workspace/${workspaceId}/form`)
    return data;
}

export async function getForm(workspaceId: string, formId: string): Promise<IForm> {
    const { data } = await http.get(`${BASE_URL}/workspace/${workspaceId}/form/${formId}`)
    return data;
}

export async function updateForm(workspaceId: string, payload: IForm): Promise<void> {
    await http.put(`${BASE_URL}/workspace/${workspaceId}/form/${payload.form_id}`, payload)
}

export async function deleteWorkspace(formId: string): Promise<void> {
    await http.delete(`${BASE_URL}/form/${formId}`)
}