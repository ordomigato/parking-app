export function convertDate(date: Date): string {
    if (!date) {
        return ''
    }
    const d = new Date(date).toLocaleDateString()
    return d
}

export function convertTime(date: Date): string {
    if (!date) {
        return ''
    }
    const t = new Date(date).toLocaleTimeString()
    return t
}