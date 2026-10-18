import { User } from "./user.model"

export interface LoginResponse {
  token: string
  refresh_token: string
  expires_at: string
  user: User
}
