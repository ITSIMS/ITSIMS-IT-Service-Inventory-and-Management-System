import { useState, useEffect } from 'react';
import { Service, CreateServiceRequest } from './types/service';
import {
  getServices,
  createService,
  updateService,
  deleteService,
} from './api/services';
import { ServiceList } from './components/ServiceList';
import { ServiceForm } from './components/ServiceForm';
import './App.css';

function App() {
  const [services, setServices] = useState<Service[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showForm, setShowForm] = useState(false);
  const [editingService, setEditingService] = useState<Service | null>(null);

  const loadServices = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await getServices();
      setServices(data);
    } catch {
      setError('Не удалось загрузить список сервисов');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadServices();
  }, []);

  const handleAddClick = () => {
    setEditingService(null);
    setShowForm(true);
  };

  const handleEdit = (service: Service) => {
    setEditingService(service);
    setShowForm(true);
  };

  const handleCancel = () => {
    setShowForm(false);
    setEditingService(null);
  };

  const handleSave = async (data: CreateServiceRequest) => {
    try {
      if (editingService) {
        await updateService(editingService.id, data);
      } else {
        await createService(data);
      }
      setShowForm(false);
      setEditingService(null);
      await loadServices();
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
      await loadServices();
    } catch {
      setError('Не удалось удалить сервис');
    }
  };

  return (
    <div className="app">
      <header className="app-header">
        <h1>ITSIMS — Каталог IT-сервисов</h1>
      </header>

      <main className="app-main">
        {error && (
          <div className="error-banner">
            {error}
            <button onClick={() => setError(null)} className="error-close">
              &times;
            </button>
          </div>
        )}

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
            onSave={handleSave}
            onCancel={handleCancel}
          />
        )}

        {loading ? (
          <p style={{ color: '#6c757d', textAlign: 'center' }}>Загрузка...</p>
        ) : (
          <ServiceList
            services={services}
            onEdit={handleEdit}
            onDelete={handleDelete}
          />
        )}
      </main>
    </div>
  );
}

export default App;
