import type { IFormattedPermit, IPermit } from "@/types";
import { convertDate, convertTime } from "./date";

export function formatPermits(permits: IPermit[]): IFormattedPermit[] {
    return permits.map(p => ({
        ...p,
        start_date: p.start_date ? `${ convertDate(p.start_date) }, ${ convertTime(p.start_date) }` : 'None',
        end_date: p.end_date ? `${ convertDate(p.end_date) }, ${ convertTime(p.end_date) }` : 'None',
        created_at: `${ convertDate(p.created_at) }, ${ convertTime(p.created_at) }`,
        updated_at: `${ convertDate(p.updated_at) }, ${ convertTime(p.updated_at) }`,
    }))
}

export function formatTime(date: Date | null): string {
    if (!date) {
        return '00:00'
    }
    const hour = ("0" + date.getHours()).slice(-2)
    const min = ("0" + date.getMinutes()).slice(-2)
    return `${hour}:${min}`
}
