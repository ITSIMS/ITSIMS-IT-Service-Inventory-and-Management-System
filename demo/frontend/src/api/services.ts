import axios from 'axios';
import {
  Service,
  CreateServiceRequest,
  UpdateServiceRequest,
  ServiceDependency,
  ServiceDependencies,
  DependencyGraph,
  ServiceStats,
} from '../types/service';

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080',
});

export const getServicesFiltered = async (params: {
  category?: string;
  status?: string;
  search?: string;
}): Promise<Service[]> => {
  const response = await apiClient.get<Service[]>('/api/v1/services', { params });
  return response.data;
};

export const getServices = async (): Promise<Service[]> => {
  return getServicesFiltered({});
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

export const getDependencies = async (id: string): Promise<ServiceDependencies> => {
  const response = await apiClient.get<ServiceDependencies>(`/api/v1/services/${id}/dependencies`);
  return response.data;
};

export const addDependency = async (
  serviceId: string,
  dependsOnId: string
): Promise<ServiceDependency> => {
  const response = await apiClient.post<ServiceDependency>(
    `/api/v1/services/${serviceId}/dependencies`,
    { depends_on_id: dependsOnId }
  );
  return response.data;
};

export const removeDependency = async (serviceId: string, depId: string): Promise<void> => {
  await apiClient.delete(`/api/v1/services/${serviceId}/dependencies/${depId}`);
};

export const getDependencyGraph = async (): Promise<DependencyGraph> => {
  const response = await apiClient.get<DependencyGraph>('/api/v1/dependencies/graph');
  return response.data;
};

export const getStats = async (): Promise<ServiceStats> => {
  const response = await apiClient.get<ServiceStats>('/api/v1/stats');
  return response.data;
};
