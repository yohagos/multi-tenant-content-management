export interface Content {
  id?: string
  tenant_id?: string
  title: string
  slug: string
  body: string
  status: string
  published: boolean
  published_at?: Date
  created_at?: Date
  updated_at?: Date
  deleted_at?: Date
}

export interface ContentFilter {
  tenant_id?: string
  limit: number
  offset: number
  search?: string
  status?: string
  published?: boolean
}
