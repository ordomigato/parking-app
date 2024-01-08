import { isServerError, ServerErrorResponse, type ServerError } from "@/utils/error";
import axios, { AxiosError, type AxiosResponse } from "axios";
import type { App } from "vue";
import { logout } from "./account.service";

const instance = axios.create({
    headers: {
        'Content-Type': 'application/json'
    },
    withCredentials: true,
});

const responseHandler = (resp: AxiosResponse): AxiosResponse => {
    return resp;
}

const responseErrorHandler = (error: AxiosError): ServerError => {
    console.log(error)
    if (
        error.response
        && error.response.status === 401
        && error.request.responseText !== "{\"error_message\":\"You are not logged in\"}"
    ) {
        logout()
    }

    if (error.response && isServerError(error.response.data)) {
        throw error.response.data
    }

    throw new ServerErrorResponse();
};

instance.interceptors.response.use(responseHandler, responseErrorHandler)

export const HttpPlugin = {
    install: (app: App) => {
        app.config.globalProperties.$http = instance;
    }
};

export default instance;