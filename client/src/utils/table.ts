import type { IColumn } from "@/types"

// returns an array of data in the object sorted by the cols array passed in
export function filterTableData(data: any, cols: string[]): string[] {
    const result = cols.map(c => {
        return data[c]
    })

    return result
}

export function filterTableCols(cols: IColumn[]): IColumn[] {
    return cols.filter(c => c.visible)
}