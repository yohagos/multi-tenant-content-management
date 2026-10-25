import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { API_BASE_URL, Building, BuildingList } from 'shared-data-access';

@Injectable({
  providedIn: 'root',
})
export class BuildingService {
  private readonly http = inject(HttpClient)
  private readonly baseUrl = inject(API_BASE_URL)

  create(request: Building): Observable<void> {
      return this.http.post<void>(
        `${this.baseUrl}/building`,
        request,
      )
    }

    getByID(id: string): Observable<Building> {
      return this.http.get<Building>(
        `${this.baseUrl}/building/${id}`,
      )
    }

    list(): Observable<BuildingList> {
      return this.http.get<BuildingList>(
        `${this.baseUrl}/building`
      )
    }

    update(id: string, building: Building): Observable<void> {
      return this.http.put<void>(
        `${this.baseUrl}/building/${id}`,
        building,
      )
    }

    delete(id: string): Observable<void> {
      return this.http.delete<void>(
        `${this.baseUrl}/building/${id}`,
      )
    }
}
