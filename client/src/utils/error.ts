export interface ServerError {
    error: string
}

export function isServerError(e: any): e is ServerError {
    return !!e.error
}

export class ServerErrorResponse implements ServerError {
    error = ''

    construct(e: ServerError) {
        if (e) {
            this.error = e.error
        }
    }
}

export function handleError(e: Error | ServerError | string | null | unknown): Error | null {
    if (!e) {
        return null;
    } else if (e instanceof Error) {
        return e;
    } else if (typeof e === 'string') {
        return new Error(e)
    } else if (isServerError(e)) {
        return new Error(e.error)
    } else {
        console.error(e)
        return new Error('unknown error')
    }
}