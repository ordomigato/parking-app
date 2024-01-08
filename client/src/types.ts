export interface IClient {
    user_uuid: string,
    username: string,
    created_at: Date,
    updated_at: Date,
    last_login: Date,
}

export interface IWorkspace {
    workspace_id: string,
    updated_at: Date,
    created_at: Date,
}