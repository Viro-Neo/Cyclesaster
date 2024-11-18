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

export const fetchGraphData = async (filters: { name: string; value: string}[], perFilter: string): Promise<ApiResponse> => {
    const filterParams = filters.map((filter, index) => `filter${index + 1}_name=${filter.name}&filter${index + 1}_value=${filter.value}`).join('&');
    return await fetchData(`/graph?${filterParams}&perFilter=${perFilter}`);
}

export const fetchMapData = async (filters: { name: string; value: string}[]): Promise<ApiResponse> => {
    const filterParams = filters.map((filter, index) => `filter${index + 1}_name=${filter.name}&filter${index + 1}_value=${filter.value}`).join('&');
    return await fetchData(`/map?${filterParams}`);
}

export const fetchAccidentById = async (id: number): Promise<ApiResponse> => {
    return await fetchData(`/get_accident?accident_id=${id}`);
}

export async function handleFilter() {
    let filtersName: string[] = [];
    try {
        const filtersApi: ApiResponse = await fetchFilters();
        filtersName = filtersApi.filters;
    } catch (error) {
        console.log(error);
    }
    return filtersName;
}
