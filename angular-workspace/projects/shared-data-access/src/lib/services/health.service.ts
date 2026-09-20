import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { API_BASE_URL } from '../tokens/api-base-url.token';

@Injectable({
  providedIn: 'root',
})
export class HealthService {
  private readonly httpClient = inject(HttpClient)
  private readonly baseUrl = inject(API_BASE_URL)

  getHealthStats() {
    console.log('HealthService => backend call triggered')
    return this.httpClient.get(
      `${this.baseUrl}/health`
    )
  }
}
