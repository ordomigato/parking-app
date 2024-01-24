import type { IPaginatedResult, IPagination, IPermit, IPermitCreateRequest } from "@/types";
import http from "./http.service"
import { PaginationQuery } from "@/utils/pagination";

const BASE_URL = `${import.meta.env.VITE_BASE_API_URL}/api`

export async function createPermit(formId: string, payload: IPermitCreateRequest): Promise<IPermit> {
    const { data } = await http.post(`${BASE_URL}/form/${formId}/permit`, payload)
    return data;
}

export async function getPermits(
    workspaceId: string,
    formId: string,
    p: IPagination = new PaginationQuery()
): Promise<IPaginatedResult<IPermit[]>> {
    const { data } = await http.get(`${BASE_URL}/workspace/${workspaceId}/form/${formId}/permit?limit=${p.limit}&page=${p.page}&sort=${p.sort}`)
    return data;
}

export async function deletePermit(workspaceId: string, formId: string, permitId: string): Promise<void> {
    await http.delete(`${BASE_URL}/workspace/${workspaceId}/form/${formId}/permit/${permitId}`)
}