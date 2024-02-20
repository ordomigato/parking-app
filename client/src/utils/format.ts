import type { IFormattedPermit, IPermit } from "@/types";
import { convertDate, convertTime } from "./date";

export function formatPermits(permits: IPermit[]): IFormattedPermit[] {
    return permits.map(p => ({
        ...p,
        expiry: p.expiry ? `${convertDate(p.expiry)}, ${ convertTime(p.expiry)}` : '',
        created_at: `${ convertDate(p.created_at) }, ${ convertTime(p.created_at) }`,
        updated_at: `${ convertDate(p.updated_at) }, ${ convertTime(p.updated_at) }`,
    }))
}
