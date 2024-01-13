export function validatePath(path: string): boolean {
    const regexp = /^[a-zA-Z0-9-_/]+$/;
    return regexp.test(path)
}

// TODO: check if fully valid. Should return url safe path string
export function formatPath(str: string): string {
    const val = str.replace(/[^a-z0-9_]+/gi, "-").replace(/^-|-$/g, '').toLowerCase()
    return `/${val}`
}