export function convertDate(date: Date) {
    const d = new Date(date)
    return d.toLocaleString()
}