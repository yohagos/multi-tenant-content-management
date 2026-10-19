import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { API_BASE_URL, Tenant } from 'shared-data-access';

@Injectable({
  providedIn: 'root',
})
export class TenantService {
  private readonly http = inject(HttpClient)
  private readonly baseUrl = inject(API_BASE_URL)

  create(request: Tenant): Observable<void> {
      return this.http.post<void>(
        `${this.baseUrl}/tenant`,
        request,
      )
    }

    getByID(id: string): Observable<Tenant> {
      return this.http.get<Tenant>(
        `${this.baseUrl}/tenant/${id}`,
      )
    }

    list(): Observable<Tenant[]> {
      return this.http.get<Tenant[]>(
        `${this.baseUrl}/tenant`
      )
    }

    update(id: string, tenant: Tenant): Observable<void> {
      return this.http.put<void>(
        `${this.baseUrl}/tenant/${id}`,
        tenant,
      )
    }

    delete(id: string): Observable<void> {
      return this.http.delete<void>(
        `${this.baseUrl}/tenant/${id}`,
      )
    }
}
