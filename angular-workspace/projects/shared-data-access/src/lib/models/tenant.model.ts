export interface Tenant {
  id?: string
  name: string
  slug: string
  config?: TenantConfig
  domain: string
  active: string
  created_at?: Date
  updated_at?: Date
  deleted_at?: Date
}

export interface TenantConfig {
  theme: string
  features: string[]
  settings: Record<string, string>
}

export interface TenantFilter {
  limit: number
  offset: number
  search?: string
  active?: boolean
}
