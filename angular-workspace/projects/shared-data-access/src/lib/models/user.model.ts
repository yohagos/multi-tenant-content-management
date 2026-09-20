export type UserRole = 'admin' | 'tenant' | 'editor' | 'public';

export interface User {
  id: string
  tenantId: string
  firstName: string
  lastName: string
  role: UserRole
  active: boolean
  lastLoginAt?: string
  createdAt: string
  updatedAt: string
}
