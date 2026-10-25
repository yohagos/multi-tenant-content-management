export type UnitStatus = 'draft' | 'available' | 'reserved' | 'rented' | 'archived'

export interface Unit {
  id?: string
  tenant_id?: string
  building_id?: string
  unit_number: string
  floor: number
  rooms: number
  bedrooms: number
  bathrooms: number
  living_area_sqm: number
  total_area_sqm: number
  balcony_area_sqm: number
  ceiling_height_m: number
  cold_rent: number
  warm_rent: number
  deposit: number
  available_from: Date
  status: UnitStatus
  published: boolean
  published_at: Date
  title: string
  description: string
  created_at?: Date
  updated_at?: Date
  deleted_at?: Date
}

export interface UnitList {
  data: Unit[],
  total: number,
  limit: number,
  offset: number,
}

export interface UnitFilter {
  tenant_id?: string
  building_id?: string
  limit: number
  offset: number
  search?: string
  status?: UnitStatus
  min_rooms?: number
  max_rooms?: number
  min_rent?: number
  max_rent?: number
  published?: boolean
}
