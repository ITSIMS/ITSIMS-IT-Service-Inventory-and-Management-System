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

export interface ServiceDependency {
  id: string;
  service_id: string;
  depends_on_id: string;
  created_at: string;
}

export interface ServiceDependencies {
  depends_on: Service[];
  used_by: Service[];
}

export interface GraphNode {
  id: string;
  name: string;
  category: string;
  status: 'active' | 'inactive';
}

export interface GraphEdge {
  id: string;
  service_id: string;
  depends_on_id: string;
}

export interface DependencyGraph {
  nodes: GraphNode[];
  edges: GraphEdge[];
}

export interface StatsItem {
  key: string;
  count: number;
}

export interface ServiceStats {
  total: number;
  by_status: StatsItem[];
  by_category: StatsItem[];
}
