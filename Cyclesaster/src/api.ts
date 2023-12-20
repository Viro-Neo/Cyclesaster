import axios, { type AxiosResponse } from 'axios';
import { API_URL } from "../config";

const axiosInstance = axios.create({
    baseURL: API_URL,
});

export interface ApiResponse {
    filters: string[],
    filter_values: string[],
}

export const fetchData = async (path: string): Promise<ApiResponse> => {
    try {
        const response: AxiosResponse<ApiResponse> = await axiosInstance.get(path);
        return response.data;
    } catch (error) {
        throw new Error(`API Error: ${error}`);
    }
};

export const fetchFilters = async (): Promise<ApiResponse> => {
    return await fetchData('/get_filters');
}

export const fetchFiltersValues = async (filter: string): Promise<ApiResponse> => {
    return await fetchData(`/get_filters_values?filter_name=${filter}`);
}
