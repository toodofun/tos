import axios from 'axios'
import { Message } from '@arco-design/web-vue'
import type { AxiosInstance, AxiosError, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig } from 'axios'

export declare interface Pager<T> {
  current: number
  size: number
  total: number
  data?: T[]
}

export enum RequestEnums {
  TIMEOUT = 60000,
  SUCCESS = '00000', // 请求成功
  NotLogin = 'B1001',  // 未登录
  BASEURL = '/api/v1'
}

const config: AxiosRequestConfig = {
  timeout: RequestEnums.TIMEOUT as number,
  withCredentials: false,
  baseURL: RequestEnums.BASEURL as string,
}

class RequestHttp {
  private service: AxiosInstance;

  public constructor(config: AxiosRequestConfig) {
    this.service = axios.create(config);

    this.service.interceptors.request.use(
      (config: InternalAxiosRequestConfig) => {
        // const token = localStorage.getItem('token') || '';
        // if (token) {
        //   config.headers['x-access-token'] = token;
        // }
        return config;
      },
      (error: AxiosError) => {
        return Promise.reject(error);
      }
    );

    this.service.interceptors.response.use(
      (response: AxiosResponse) => {
        const { data } = response;
        if (data.code && data.code !== RequestEnums.SUCCESS) {
          Message.error(data.message);
          return Promise.reject(data);
        }
        return data.code ? data.data : data;
      },
      (error: AxiosError) => {
        if (error.response) {
          this.handleCode(error.response.status);
        } else {
          Message.error('请求超时，服务器异常');
        }
        return Promise.reject(error);
      }
    );
  }

  private handleCode(code: number): void {
    switch(code) {
      case 401:
        Message.error('登录失效，请重新登录');
        break;
      default:
        Message.error('数据请求失败');
        break;
    }
  }

  public get<T>(url: string, params?: object): Promise<T> {
    return this.service.get(url, { params });
  }

  public post<T>(url: string, params?: object, config?: AxiosRequestConfig): Promise<T> {
    return this.service.post(url, params, config);
  }

  public put<T>(url: string, params?: object): Promise<T> {
    return this.service.put(url, params);
  }

  public delete<T>(url: string, params?: object): Promise<T> {
    return this.service.delete(url, { params });
  }

  public sse(url: string, onMessage?: (event: MessageEvent) => void, onError?: (event: Event) => void, onOpen?: () => void): EventSource {
    const eventSource = new EventSource(`${location.protocol}//${window.location.hostname}:${window.location.port}${RequestEnums.BASEURL}${url.startsWith("/") ? url : `/${url}`}`);

    if (onMessage) eventSource.onmessage = onMessage;
    if (onError) eventSource.onerror = onError;
    if (onOpen) eventSource.onopen = onOpen;

    return eventSource;
  }
}

export default new RequestHttp(config);
