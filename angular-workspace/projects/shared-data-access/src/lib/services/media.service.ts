import { HttpClient, HttpParams } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { API_BASE_URL, Media, MediaEntityType } from 'shared-data-access';

export interface MediaListResponse {
  data: Media[]
  total: number
}

export interface UploadMediaRequest {
  tenant_id?: string
  entity_type: MediaEntityType
  entity_id: string
  file: File
  alt_text?: string
  sort_order?: number
  is_primary?: boolean
}

@Injectable({
  providedIn: 'root',
})
export class MediaService {
  private readonly http = inject(HttpClient)
  private readonly baseUrl = inject(API_BASE_URL)

  upload(request: UploadMediaRequest): Observable<Media> {
    const formData = new FormData()

    if (request.tenant_id) {
      formData.append('tenant_id', request.tenant_id)
    }
    if (request.alt_text) {
      formData.append('alt_text', request.alt_text)
    }
    formData.append("entity_type", request.entity_type)
    formData.append("entity_id", request.entity_id)
    formData.append("file", request.file)
    if (request.sort_order) {
      formData.append("sort_order", String(request.sort_order ?? 0))
    }
    if (request.is_primary !== undefined) {
      formData.append("is_primary", String(request.is_primary ?? false))
    }

    return this.http.post<Media>(
      `${this.baseUrl}/media`,
      formData,
    )
  }

  listByEntity(
    entityType: MediaEntityType,
    entity_id: string,
  ): Observable<MediaListResponse> {
    const params = new HttpParams()
      .set('entity_type', entityType)
      .set('entity_id', entity_id)

    return this.http.get<MediaListResponse>(
      `${this.baseUrl}/media`,
      {params},
    )
  }

  getImageUrl(id: string): string {
    return `${this.baseUrl}/media/${id}`
  }

  delete(id: string): Observable<void> {
    return this.http.delete<void>(
      `${this.baseUrl}/media/${id}`,
    )
  }
}
