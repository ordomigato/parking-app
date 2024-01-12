export function validatePath(path: string): boolean {
    const regexp = /^[a-zA-Z0-9-_]+$/;
    return regexp.test(path)
    // return path.match(regexp) === -1
}