import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { API_BASE_URL, Content } from 'shared-data-access';

@Injectable({
  providedIn: 'root',
})
export class ContentService {
  private readonly http = inject(HttpClient)
  private readonly baseUrl = inject(API_BASE_URL)

  create(request: Content): Observable<void> {
      return this.http.post<void>(
        `${this.baseUrl}/content`,
        request,
      )
    }

    getByID(id: string): Observable<Content> {
      return this.http.get<Content>(
        `${this.baseUrl}/content/${id}`,
      )
    }

    list(): Observable<Content[]> {
      return this.http.get<Content[]>(
        `${this.baseUrl}/content`
      )
    }

    update(id: string, content: Content): Observable<void> {
      return this.http.put<void>(
        `${this.baseUrl}/content/${id}`,
        content,
      )
    }

    delete(id: string): Observable<void> {
      return this.http.delete<void>(
        `${this.baseUrl}/content/${id}`,
      )
    }
}
