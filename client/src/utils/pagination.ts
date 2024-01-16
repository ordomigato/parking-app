import type { IPagination, ISortBy } from "@/types";

export function getDefaultPagination(): IPagination {
    return {
        limit: 10,
        page: 1,
        sort: 'created_at'
    }
}

export class PaginationQuery implements IPagination {
    limit: number
    page: number
    sort: ISortBy
    constructor(
        limit: number = 10,
        page: number = 1,
        sort: ISortBy = 'created_at'
    ) {
        this.limit = limit
        this.page = page
        this.sort = sort
    }
}
