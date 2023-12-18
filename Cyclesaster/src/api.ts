import axios, { type AxiosResponse, type AxiosRequestConfig } from 'axios';
import { API_URL } from "../config";

const axiosInstance = axios.create({
    baseURL: API_URL,
    timeout: 5000,
});

export interface ApiResponse {
    filters: string[],
}

export const fetchData = async (path: string): Promise<ApiResponse> => {
    try {
        const response: AxiosResponse<ApiResponse> = await axiosInstance.get(path);
        return response.data;
    } catch (error) {
        throw new Error(`API Error: ${error}`);
    }
};
