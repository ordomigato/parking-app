import axios, { AxiosError, type AxiosResponse } from "axios";
import type { App } from "vue";

const instance = axios.create();

const responseHandler = (resp: AxiosResponse): AxiosResponse => {
    return resp;
}

const responseErrorHandler = (error: AxiosError): AxiosError => {
    if (error.response && error.response.status === 401) {
        // logout user
    }

    return error;
};

instance.interceptors.response.use(responseHandler, responseErrorHandler)

export const HttpPlugin = {
    install: (app: App) => {
        app.config.globalProperties.$http = instance;
    }
};

export default instance;