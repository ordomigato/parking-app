export enum LocalSettings {
    PermitTableFilter = "permit_table_filter"
}

export function saveLocal(settingName: string, data: string) {
    localStorage.setItem(settingName, data)
}

export function loadLocal(settingName: string): string | null {
    return localStorage.getItem(settingName)
}