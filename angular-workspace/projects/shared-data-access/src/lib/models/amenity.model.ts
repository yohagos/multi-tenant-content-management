export type AmenityCategory = 'accessibility' | 'bathroom' | 'building' | 'interior' | 'kitchen' | 'parking' | 'other' | 'outdoor' | 'technology'

export interface Amenity {
  id?: string
  tenant_id?: string
  code: string
  name: string
  category: AmenityCategory
  icon: string
  is_global: boolean
  created_at?: Date
}
