export type BuildingStatus = 'draft' | 'published' | 'archived'

export interface Building {
  id?: string
  tenant_id?: string
  name: string
  slug?: string
  street: string
  house_number: string
  postal_code: string
  city: string
  country: string
  latitude?: number
  longitude?: number
  year_build: number
  total_floors: number
  total_units: number
  description: string
  status: BuildingStatus
  published: boolean
  published_at?: Date
  created_at?: Date
  updated_at?: Date
  deleted_at?: Date
}

export interface BuildingList {
  data: Building[],
  total: number,
  limit: number,
  offset: number,
}

export interface BuildingFilter {
  tenant_id?: string
  limit: number
  offset: number
  search?: string
  city?: string
  status?: string
  published: boolean
}
