export interface ServerError {
    error_message: string
}

export function isServerError(e: any): e is ServerError {
    return !!e.error_message
}

export class ServerErrorResponse implements ServerError {
    error_message = ''

    construct(e: ServerError) {
        if (e) {
            this.error_message = e.error_message
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
        return new Error(e.error_message)
    } else {
        console.error(e)
        return new Error('unknown error')
    }
}