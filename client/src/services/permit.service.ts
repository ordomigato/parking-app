import type { IDownloadPermitRequest, IPaginatedResult, IPagination, IPermit, IPermitCreateRequest } from "@/types";
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
    p: IPagination = new PaginationQuery(),
    search?: string,
): Promise<IPaginatedResult<IPermit[]>> {
    let q = `limit=${p.limit}&page=${p.page}&sort=${p.sort}`
    if (search) {
        q = q + `&search=${search}`
    }
    const { data } = await http.get(`${BASE_URL}/workspace/${workspaceId}/form/${formId}/permit?${q}`)
    return data;
}

export async function downloadPermits(
    workspaceId: string,
    formId: string,
    payload: IDownloadPermitRequest
): Promise<IPermit[]> {
    const { data } = await http.get(`${BASE_URL}/workspace/${workspaceId}/form/${formId}/permit/download?from=${payload.from}&to=${payload.to}`)
    return data;
}

export async function deletePermit(workspaceId: string, formId: string, permitId: string): Promise<void> {
    await http.delete(`${BASE_URL}/workspace/${workspaceId}/form/${formId}/permit/${permitId}`)
}