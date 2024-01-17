export function convertDate(date: Date) {
    const d = new Date(date).toLocaleDateString()
    return d
}

export function convertTime(date: Date) {
    const d = new Date(date).toLocaleTimeString()
    return d
}