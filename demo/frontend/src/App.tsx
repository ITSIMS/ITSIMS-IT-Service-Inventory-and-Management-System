import { useState, useEffect, useCallback } from 'react';
import { Service, CreateServiceRequest, ServiceStats } from './types/service';
import {
  getServicesFiltered,
  createService,
  updateService,
  deleteService,
  getStats,
  getDependencies,
  addDependency,
  removeDependency,
} from './api/services';
import { ServiceList } from './components/ServiceList';
import { ServiceForm } from './components/ServiceForm';
import { StatsPanel } from './components/StatsPanel';
import { FilterBar } from './components/FilterBar';
import { DependencyPanel } from './components/DependencyPanel';
import { GraphView } from './components/GraphView';
import './App.css';

type Tab = 'services' | 'graph';

interface FilterParams {
  category?: string;
  status?: string;
  search?: string;
}

function App() {
  const [services, setServices] = useState<Service[]>([]);
  const [allServices, setAllServices] = useState<Service[]>([]);
  const [stats, setStats] = useState<ServiceStats | null>(null);
  const [loading, setLoading] = useState(true);
  const [statsLoading, setStatsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showForm, setShowForm] = useState(false);
  const [editingService, setEditingService] = useState<Service | null>(null);
  const [editingDeps, setEditingDeps] = useState<string[]>([]);
  const [selectedService, setSelectedService] = useState<Service | null>(null);
  const [activeTab, setActiveTab] = useState<Tab>('services');
  const [filterParams, setFilterParams] = useState<FilterParams>({});

  const loadAllServices = async () => {
    try {
      const data = await getServicesFiltered({});
      setAllServices(data);
    } catch {
      // silent
    }
  };

  const loadServices = useCallback(async (params: FilterParams = {}) => {
    try {
      setLoading(true);
      setError(null);
      const data = await getServicesFiltered(params);
      setServices(data);
    } catch {
      setError('Не удалось загрузить список сервисов');
    } finally {
      setLoading(false);
    }
  }, []);

  const loadStats = async () => {
    try {
      setStatsLoading(true);
      const data = await getStats();
      setStats(data);
    } catch {
      // stats are non-critical
    } finally {
      setStatsLoading(false);
    }
  };

  useEffect(() => {
    loadServices(filterParams);
  }, [filterParams, loadServices]);

  useEffect(() => {
    loadStats();
    loadAllServices();
  }, []);

  const handleFilter = useCallback((params: FilterParams) => {
    setFilterParams(params);
  }, []);

  const handleAddClick = () => {
    setEditingService(null);
    setEditingDeps([]);
    setShowForm(true);
  };

  const handleEdit = async (service: Service) => {
    setEditingService(service);
    setShowForm(true);
    try {
      const deps = await getDependencies(service.id);
      setEditingDeps(deps.depends_on.map((d) => d.id));
    } catch {
      setEditingDeps([]);
    }
  };

  const handleCancel = () => {
    setShowForm(false);
    setEditingService(null);
    setEditingDeps([]);
  };

  const handleSave = async (data: CreateServiceRequest, depIds: string[]) => {
    try {
      let serviceId: string;
      if (editingService) {
        await updateService(editingService.id, data);
        serviceId = editingService.id;

        // Sync dependencies: add new, remove deleted
        const toAdd = depIds.filter((id) => !editingDeps.includes(id));
        const toRemove = editingDeps.filter((id) => !depIds.includes(id));
        await Promise.all(toAdd.map((id) => addDependency(serviceId, id)));
        await Promise.all(toRemove.map((id) => removeDependency(serviceId, id)));
      } else {
        const created = await createService(data);
        serviceId = created.id;
        await Promise.all(depIds.map((id) => addDependency(serviceId, id)));
      }
      setShowForm(false);
      setEditingService(null);
      setEditingDeps([]);
      await loadServices(filterParams);
      await loadStats();
      await loadAllServices();
    } catch {
      setError('Не удалось сохранить сервис');
    }
  };

  const handleDelete = async (id: string) => {
    if (!window.confirm('Вы уверены, что хотите удалить этот сервис?')) {
      return;
    }
    try {
      await deleteService(id);
      await loadServices(filterParams);
      await loadStats();
      await loadAllServices();
      if (selectedService?.id === id) {
        setSelectedService(null);
      }
    } catch {
      setError('Не удалось удалить сервис');
    }
  };

  const handleRowClick = (service: Service) => {
    setSelectedService((prev) => (prev?.id === service.id ? null : service));
  };

  return (
    <div className="app">
      <header className="app-header">
        <h1>ITSIMS — Каталог IT-сервисов</h1>
      </header>

      <main className="app-main">
        <StatsPanel stats={stats} loading={statsLoading} />

        {error && (
          <div className="error-banner">
            {error}
            <button onClick={() => setError(null)} className="error-close">
              &times;
            </button>
          </div>
        )}

        <div className="tabs">
          <button
            className={`tab-btn${activeTab === 'services' ? ' tab-btn--active' : ''}`}
            onClick={() => setActiveTab('services')}
          >
            Сервисы
          </button>
          <button
            className={`tab-btn${activeTab === 'graph' ? ' tab-btn--active' : ''}`}
            onClick={() => setActiveTab('graph')}
          >
            Граф зависимостей
          </button>
        </div>

        {activeTab === 'services' && (
          <>
            <FilterBar services={allServices} onFilter={handleFilter} />

            {!showForm && (
              <div style={{ marginBottom: '20px' }}>
                <button onClick={handleAddClick} className="btn-primary">
                  + Добавить сервис
                </button>
              </div>
            )}

            {showForm && (
              <ServiceForm
                service={editingService}
                allServices={allServices}
                initialDependencies={editingDeps}
                onSave={handleSave}
                onCancel={handleCancel}
              />
            )}

            <div className="services-with-panel">
              <div className="services-table-wrap">
                {loading ? (
                  <p style={{ color: '#6c757d', textAlign: 'center' }}>Загрузка...</p>
                ) : (
                  <ServiceList
                    services={services}
                    onEdit={handleEdit}
                    onDelete={handleDelete}
                    onRowClick={handleRowClick}
                    selectedId={selectedService?.id}
                  />
                )}
              </div>

              {selectedService && (
                <DependencyPanel
                  service={selectedService}
                  allServices={allServices}
                  onClose={() => setSelectedService(null)}
                />
              )}
            </div>
          </>
        )}

        {activeTab === 'graph' && <GraphView />}
      </main>
    </div>
  );
}

export default App;
