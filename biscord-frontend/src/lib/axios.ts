import axios, { AxiosError, type AxiosRequestConfig } from "axios";
import { API_URL } from "../config";

export const axiosInstance = axios.create({
  withCredentials: true,
  baseURL: API_URL,
});

/**
 * Sends a POST request to the specified URL with the given data and configuration.
 * @template TRequest The type of the request data.
 * @template TResponse The type of the response data.
 * @param {string} url The URL to send the request to.
 * @param {TRequest} data The data to send with the request.
 * @returns {Promise<TResponse>} A promise that resolves with the response data.
 * @throws {AxiosError} If the request fails.
 */
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
