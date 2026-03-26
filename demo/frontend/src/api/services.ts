import axios from 'axios';
import { Service, CreateServiceRequest, UpdateServiceRequest } from '../types/service';

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080',
});

export const getServices = async (): Promise<Service[]> => {
  const response = await apiClient.get<Service[]>('/api/v1/services');
  return response.data;
};

export const getService = async (id: string): Promise<Service> => {
  const response = await apiClient.get<Service>(`/api/v1/services/${id}`);
  return response.data;
};

export const createService = async (req: CreateServiceRequest): Promise<Service> => {
  const response = await apiClient.post<Service>('/api/v1/services', req);
  return response.data;
};

export const updateService = async (id: string, req: UpdateServiceRequest): Promise<Service> => {
  const response = await apiClient.put<Service>(`/api/v1/services/${id}`, req);
  return response.data;
};

export const deleteService = async (id: string): Promise<void> => {
  await apiClient.delete(`/api/v1/services/${id}`);
};
