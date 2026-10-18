export type MediaEntityType = 'building' | 'unit'

export interface Media {
  id?: string
  tenant_id?: string
  entity_type: MediaEntityType
  entity_id: string
  file_name: string
  file_path: string
  file_size: number
  mime_type: string
  alt_text: string
  sort_order: number
  is_primary: boolean
  created_at?: Date
}
