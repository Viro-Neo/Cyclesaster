import axios, { type AxiosResponse } from 'axios';
import { API_URL } from "../config";

const axiosInstance = axios.create({
    baseURL: API_URL,
});

export interface ApiResponse {
    filters: string[],
    filter_values: string[],
    data: Object,
}

export const fetchData = async (path: string): Promise<ApiResponse> => {
    try {
        const response: AxiosResponse<ApiResponse> = await axiosInstance.get(path);
        return response.data;
    } catch (error) {
        throw new Error(`API Error: ${error}`);
    }
};

const fetchFilters = async (): Promise<ApiResponse> => {
    return await fetchData('/get_filters');
}

export const fetchFiltersValues = async (filter: string): Promise<ApiResponse> => {
    return await fetchData(`/get_filters_values?filter_name=${filter}`);
}

export const fetchGraphData = async (filter: string, filter2: string, filter3: string): Promise<ApiResponse> => {
    return await fetchData(`/graph?filter1_name=${filter}&filter1_value=${filter2}&filter2=${filter3}`);
}

export async function handleFilter(filtersName: string[]) {
    try {
        const filtersApi: ApiResponse = await fetchFilters();
        filtersName = filtersApi.filters;
    } catch (error) {
        console.log(error);
    }
    return filtersName;
}
