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
    updated_at: Date,
    created_at: Date,
}

export interface IWorkspaceCreateRequest {
    name: string,
}

export interface IForm {
    form_id: string,
    workspace_id: string,
    name: string,
    submission_constraint_type: string,
    submission_constraint_limit: string,
    created_at: Date,
    updated_at: Date,
}

export interface IFormCreateRequest {
    name: string,
    submission_constraint_type: string,
    submission_constraint_limit: number,
}

export interface IDropdownItem<S extends string, T extends string> {
    value: S
    name: T
}

export enum IFormSubmissionConstraintTypes {
    minutes = 'minutes',
    hours = 'hours',
    days = 'days',
    months = 'months'
}