import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { API_BASE_URL } from 'shared-data-access';

@Injectable({
  providedIn: 'root',
})
export class ContentService {
  private readonly http = inject(HttpClient)
  private readonly baseUrl = inject(API_BASE_URL)
}
