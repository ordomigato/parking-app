import type { IPermit, IPermitCreateRequest } from "@/types";
import http from "./http.service"

const BASE_URL = `${import.meta.env.VITE_BASE_API_URL}/api`

export async function createPermit(formId: string, payload: IPermitCreateRequest): Promise<IPermit> {
    const { data } = await http.post(`${BASE_URL}/form/${formId}/permit`, payload)
    return data;
}

export async function getPermits(formId: string): Promise<IPermit[]> {
    console.log('Getting permits')
    const { data } = await http.get(`${BASE_URL}/form/${formId}/permit`)
    console.log(data)
    return data;
}