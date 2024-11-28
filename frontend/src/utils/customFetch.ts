import axios, { AxiosError } from 'axios';
import { toast } from 'react-toastify'

export default interface IBaseHttp {
  url: string;
  data?: object;
  headers?: object | undefined;
  method: 'get' | 'post' | 'put' | 'delete';
}

class CustomFetch {
  private readonly customFecth = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL
  })

  async post<T>(
    url: string,
    data: object,
    headers?: object | undefined
  ): Promise<[undefined, T] | [Error]> {
    return await this.baseHttp<T>({
      url,
      data,
      headers,
      method: 'post',
    });
  }

  async get<T>(
    url: string,
    headers?: object | undefined
  ): Promise<[undefined, T] | [Error]> {
    const response = await this.baseHttp<T>({
      url,
      headers,
      method: 'get',
    });
    return response;
  }

  async delete<T>(
    url: string,
    headers?: object | undefined
  ): Promise<[undefined, T] | [Error]> {
    const response = await this.baseHttp<T>({
      url,
      headers,
      method: 'delete',
    });
    return response;
  }

  async put<T>(
    url: string,
    data: object = {},
    headers?: object | undefined
  ): Promise<[undefined, T] | [Error]> {
    const response = await this.baseHttp<T>({
      url,
      data,
      headers,
      method: 'put',
    });
    return response;
  }

  async baseHttp<T>(baseHttp: IBaseHttp): Promise<[undefined, T] | [Error]> {
    try {
      const token = localStorage.getItem('token');
      baseHttp.headers = { ...baseHttp.headers, token };
      const result = await this.customFecth.request(baseHttp);
      if (result.data.messageAlert) {
        toast.success(result.data.messageAlert);
      }
      return [undefined, result.data as T];
    } catch (error) {
      if (error instanceof AxiosError) {
        const errorMsg: string = error.response?.data?.messageAlert
          ? (error.response?.data.messageAlert as string)
          : 'Internal error try later';
        toast.error(errorMsg);
      } else {
        toast.error('Internal error try later');
      }
      return [error as Error];
    }
  }
}

export const customFetch = new CustomFetch();
;
