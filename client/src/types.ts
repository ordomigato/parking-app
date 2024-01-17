export interface IClient {
    client_id: string,
    username: string,
    created_at: Date,
    updated_at: Date,
    last_login: Date,
}

export interface IWorkspace {
    workspace_id: string,
    name: string,
    path: string,
    updated_at: Date,
    created_at: Date,
}

export interface IWorkspaceCreateRequest {
    name: string,
    path: string,
}

export interface IWorkspaceUpdateRequest {
    name: string,
    path: string,
}

export interface IForm {
    form_id: string,
    workspace_id: string,
    name: string,
    path: {
        path: string
    },
    duration_measurement_unit: IFormDurationMeasurementUnits,
    duration_limit: number,
    duration_reset_time: string
    created_at: Date,
    updated_at: Date,
}

export interface IFormCreateRequest {
    name: string,
    path: string,
    duration_measurement_unit: IFormDurationMeasurementUnits,
    duration_limit: number,
    duration_reset_time: string
}

export interface IFormUpdateRequest {
    name: string,
    duration_measurement_unit: IFormDurationMeasurementUnits,
    duration_limit: number,
    duration_reset_time: string
}

export interface IFormPathUpdateRequest {
    path: string,
}

export interface IDropdownItem<S extends string, T extends string> {
    value: S
    name: T
}

export enum IFormDurationMeasurementUnits {
    none = '',
    minutes = 'minutes',
    hours = 'hours',
    days = 'days',
    months = 'months'
}

export interface IPermitCreateRequest {
    first_name: string,
    last_name: string,
    email: string,
    primary_phone: string,
    v_plate: string,
    v_make: string,
    v_model: string,
    v_color: string,
}

export interface IPermit {
    permit_id: string,
    first_name: string,
    last_name: string,
    email: string,
    primary_phone: string,
    v_plate: string,
    v_make: string,
    v_model: string,
    v_color: string,
    expiry: Date,
    created_at: Date,
    updated_at: Date,
}

export interface IPaginatedResult<T> {
    data: T,
    count: number
}

export type ISortBy = 'created_at'
export interface IPagination {
    limit: number,
    page: number,
    sort: ISortBy
}