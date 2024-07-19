import axios, { AxiosError, type AxiosRequestConfig } from "axios";
import { API_URL } from "../config";

export const axiosInstance = axios.create({
  withCredentials: true,
  baseURL: API_URL,
});

export const post = async <TRequest, TResponse>(
  url: string,
  data: TRequest,
  config?: AxiosRequestConfig,
) => {
  try {
    const response = await axiosInstance.post(url, data, config);
    return response.data as { data: TResponse; error: null };
  } catch (error) {
    const message = (error as AxiosError<{ error: string }>).response?.data
      ?.error;
    console.log(message);
    return { data: null, error: message || "An error occurred." };
  }
};
