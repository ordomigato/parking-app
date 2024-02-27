import type { IResetInterval } from "@/types"

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

export function convertDateTime(date: Date): string {
    return `${convertDate(date)} ${convertTime(date)}`
}

export function getResetDate(resetInterval: IResetInterval): string {
    const time = convertTime(new Date(resetInterval.reference_date))

    // month related logic
    const refDate = new Date(resetInterval.reference_date)
    const now = new Date(Date.now())
    let date = new Date(refDate.setMonth(now.getMonth() + resetInterval.value))
    // if current day of month is less than the reset day of month -> set date to 1 month less
    if (date.getDate() >= now.getDate()) {
        if (date.getHours() >= now.getHours()) {
            if (date.getMinutes() > now.getMinutes()) {
                date = new Date(refDate.setMonth(now.getMonth() + resetInterval.value - 1))
            }
        }
    }
    
    switch (resetInterval.unit) {
    case 'days':
        return `${time}`
    case 'months':
        return `${convertDate(date)} @ ${time}`
    default:
        return ''
    }
}