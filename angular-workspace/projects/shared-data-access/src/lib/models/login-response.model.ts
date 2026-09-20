import { User } from "./user.model"

export interface LoginResponse {
  token: string
  refreshToken: string
  expiresAt: string
  user: User
}
