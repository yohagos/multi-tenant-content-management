import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Amenity, API_BASE_URL } from 'shared-data-access';

@Injectable({
  providedIn: 'root',
})
export class AmenityService {
  private readonly http = inject(HttpClient)
  private readonly baseUrl = inject(API_BASE_URL)

  create(request: Amenity): Observable<void> {
    return this.http.post<void>(
      `${this.baseUrl}/amenity`,
      request,
    )
  }

  getByID(id: string): Observable<Amenity> {
    return this.http.get<Amenity>(
      `${this.baseUrl}/amenity/${id}`,
    )
  }

  list(): Observable<Amenity[]> {
    return this.http.get<Amenity[]>(
      `${this.baseUrl}/amenity`
    )
  }

  update(id: string, amenity: Amenity): Observable<void> {
    return this.http.put<void>(
      `${this.baseUrl}/amenity/${id}`,
      amenity,
    )
  }

  delete(id: string): Observable<void> {
    return this.http.delete<void>(
      `${this.baseUrl}/amenity/${id}`,
    )
  }
}
