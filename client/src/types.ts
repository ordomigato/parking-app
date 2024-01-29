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

export interface IDurationUnitMeasurement {
    unit: IFormDurationMeasurementUnits,
    value: number
}

export interface IResetInterval extends IDurationUnitMeasurement {
    reference_date: Date
}

export interface ICycleData {
    duration_limit: IDurationUnitMeasurement,
    reset_interval: IResetInterval,
}

export interface IForm {
    form_id: string,
    workspace_id: string,
    name: string,
    path: {
        path: string
    },
    cycle_data: ICycleData | null,
    created_at: Date,
    updated_at: Date,
}

export interface IFormCreateRequest {
    name: string,
    path: string,
    cycle_data: ICycleData | null,
}

export interface IFormUpdateRequest {
    name: string,
    cycle_data: ICycleData | null,
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
    duration: number,
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

export interface IDownloadPermitRequest {
    from: string,
    to: string,
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

export type ToastStatus = "success" | "warning" | "error";

export interface IToast {
  text: string;
  status: ToastStatus;
  id: number;
}
