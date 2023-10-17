export function handleError(e: Error | string | null | unknown): Error | null {
    if (!e) {
        return null;
    } else if (e instanceof Error) {
        return e;
    } else if (typeof e === 'string') {
        return new Error(e)
    } else {
        console.error(e)
        return new Error('unknown error')
    }
}