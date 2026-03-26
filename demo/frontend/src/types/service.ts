export interface Service {
  id: string;
  name: string;
  description: string;
  category: string;
  status: 'active' | 'inactive';
  created_at: string;
  updated_at: string;
}

export interface CreateServiceRequest {
  name: string;
  description: string;
  category: string;
  status: 'active' | 'inactive';
}

export type UpdateServiceRequest = CreateServiceRequest;
