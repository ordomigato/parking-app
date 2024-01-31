import type { IPaginatedResult, IPagination } from "@/types";
import http from "./http.service"
import { PaginationQuery } from "@/utils/pagination";

const BASE_URL = `${import.meta.env.VITE_BASE_API_URL}/api`

export async function getBlacklist(workspaceId: string, formId: string, p: IPagination = new PaginationQuery()): Promise<IPaginatedResult<string[]>> {
    const { data } = await http.get(`${BASE_URL}/workspace/${workspaceId}/form/${formId}/blacklist?limit=${p.limit}&page=${p.page}&sort=${p.sort}`)
    return data;
}

export async function addToBlacklist(workspaceId: string, formId: string, plate: string): Promise<void> {
    await http.post(`${BASE_URL}/workspace/${workspaceId}/form/${formId}/blacklist`, {
        v_plate: plate,
    })
}

export async function deleteFromBlacklist(workspaceId: string, formId: string, plate: string): Promise<void> {
    await http.delete(`${BASE_URL}/workspace/${workspaceId}/form/${formId}/blacklist/${plate}`)
}