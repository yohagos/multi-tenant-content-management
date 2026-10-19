import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { API_BASE_URL, Unit } from 'shared-data-access';

@Injectable({
  providedIn: 'root',
})
export class UnitService {
  private readonly http = inject(HttpClient)
  private readonly baseUrl = inject(API_BASE_URL)

  create(request: Unit): Observable<void> {
      return this.http.post<void>(
        `${this.baseUrl}/unit`,
        request,
      )
    }

    getByID(id: string): Observable<Unit> {
      return this.http.get<Unit>(
        `${this.baseUrl}/unit/${id}`,
      )
    }

    list(): Observable<Unit[]> {
      return this.http.get<Unit[]>(
        `${this.baseUrl}/unit`
      )
    }

    update(id: string, unit: Unit): Observable<void> {
      return this.http.put<void>(
        `${this.baseUrl}/unit/${id}`,
        unit,
      )
    }

    delete(id: string): Observable<void> {
      return this.http.delete<void>(
        `${this.baseUrl}/unit/${id}`,
      )
    }
}
