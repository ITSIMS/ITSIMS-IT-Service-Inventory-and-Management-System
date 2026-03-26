import { useState, useEffect } from 'react';
import { Service, CreateServiceRequest } from '../types/service';

interface ServiceFormProps {
  service?: Service | null;
  allServices: Service[];
  initialDependencies?: string[];
  onSave: (data: CreateServiceRequest, depIds: string[]) => void;
  onCancel: () => void;
}

export function ServiceForm({
  service,
  allServices,
  initialDependencies = [],
  onSave,
  onCancel,
}: ServiceFormProps) {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [category, setCategory] = useState('');
  const [status, setStatus] = useState<'active' | 'inactive'>('active');
  const [selectedDeps, setSelectedDeps] = useState<string[]>([]);

  useEffect(() => {
    if (service) {
      setName(service.name);
      setDescription(service.description);
      setCategory(service.category);
      setStatus(service.status);
    } else {
      setName('');
      setDescription('');
      setCategory('');
      setStatus('active');
    }
    setSelectedDeps(initialDependencies);
  }, [service, initialDependencies]);

  const handleToggleDep = (id: string) => {
    setSelectedDeps((prev) =>
      prev.includes(id) ? prev.filter((d) => d !== id) : [...prev, id]
    );
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSave({ name, description, category, status }, selectedDeps);
  };

  const inputStyle: React.CSSProperties = {
    width: '100%',
    padding: '8px 12px',
    border: '1px solid #ced4da',
    borderRadius: '4px',
    fontSize: '14px',
    boxSizing: 'border-box',
    marginTop: '4px',
  };

  const labelStyle: React.CSSProperties = {
    display: 'block',
    marginBottom: '12px',
    fontSize: '14px',
    fontWeight: 500,
  };

  // Exclude the service being edited from the available deps list
  const availableServices = allServices.filter((s) => s.id !== service?.id);

  return (
    <div
      style={{
        backgroundColor: '#fff',
        padding: '24px',
        borderRadius: '8px',
        boxShadow: '0 2px 8px rgba(0,0,0,0.1)',
        marginBottom: '24px',
        maxWidth: '600px',
      }}
    >
      <h3 style={{ marginTop: 0, marginBottom: '20px' }}>
        {service ? 'Редактировать сервис' : 'Добавить сервис'}
      </h3>
      <form onSubmit={handleSubmit}>
        <label style={labelStyle}>
          Название *
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            minLength={1}
            maxLength={200}
            style={inputStyle}
            placeholder="Введите название сервиса"
          />
        </label>
        <label style={labelStyle}>
          Описание
          <textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            style={{ ...inputStyle, minHeight: '80px', resize: 'vertical' }}
            placeholder="Введите описание сервиса"
          />
        </label>
        <label style={labelStyle}>
          Категория
          <input
            type="text"
            value={category}
            onChange={(e) => setCategory(e.target.value)}
            style={inputStyle}
            placeholder="Например: DevOps, Monitoring, Security"
          />
        </label>
        <label style={labelStyle}>
          Статус
          <select
            value={status}
            onChange={(e) => setStatus(e.target.value as 'active' | 'inactive')}
            style={inputStyle}
          >
            <option value="active">Активен</option>
            <option value="inactive">Неактивен</option>
          </select>
        </label>

        {availableServices.length > 0 && (
          <div style={{ marginBottom: '16px' }}>
            <div style={{ fontSize: '14px', fontWeight: 500, marginBottom: '8px' }}>
              Зависимости
              <span style={{ fontWeight: 400, color: '#6c757d', marginLeft: '6px', fontSize: '12px' }}>
                (от каких сервисов зависит этот)
              </span>
            </div>
            <div
              style={{
                border: '1px solid #ced4da',
                borderRadius: '4px',
                maxHeight: '180px',
                overflowY: 'auto',
                padding: '4px 0',
              }}
            >
              {availableServices.map((s) => (
                <label
                  key={s.id}
                  style={{
                    display: 'flex',
                    alignItems: 'center',
                    gap: '8px',
                    padding: '6px 12px',
                    cursor: 'pointer',
                    fontSize: '13px',
                    backgroundColor: selectedDeps.includes(s.id) ? '#e8f0fe' : 'transparent',
                    transition: 'background-color 0.1s',
                  }}
                >
                  <input
                    type="checkbox"
                    checked={selectedDeps.includes(s.id)}
                    onChange={() => handleToggleDep(s.id)}
                    style={{ cursor: 'pointer' }}
                  />
                  <span style={{ flex: 1, fontWeight: 500 }}>{s.name}</span>
                  {s.category && (
                    <span
                      style={{
                        fontSize: '11px',
                        color: '#6c757d',
                        background: '#f1f3f5',
                        padding: '1px 6px',
                        borderRadius: '10px',
                      }}
                    >
                      {s.category}
                    </span>
                  )}
                  <span
                    style={{
                      fontSize: '11px',
                      padding: '1px 6px',
                      borderRadius: '10px',
                      background: s.status === 'active' ? '#d4edda' : '#e2e3e5',
                      color: s.status === 'active' ? '#155724' : '#383d41',
                    }}
                  >
                    {s.status === 'active' ? 'активен' : 'неактивен'}
                  </span>
                </label>
              ))}
            </div>
            {selectedDeps.length > 0 && (
              <div style={{ fontSize: '12px', color: '#0d6efd', marginTop: '4px' }}>
                Выбрано: {selectedDeps.length}
              </div>
            )}
          </div>
        )}

        <div style={{ display: 'flex', gap: '12px', marginTop: '8px' }}>
          <button
            type="submit"
            style={{
              padding: '8px 20px',
              backgroundColor: '#0d6efd',
              color: '#fff',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
              fontSize: '14px',
              fontWeight: 500,
            }}
          >
            Сохранить
          </button>
          <button
            type="button"
            onClick={onCancel}
            style={{
              padding: '8px 20px',
              backgroundColor: '#6c757d',
              color: '#fff',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
              fontSize: '14px',
            }}
          >
            Отмена
          </button>
        </div>
      </form>
    </div>
  );
}
