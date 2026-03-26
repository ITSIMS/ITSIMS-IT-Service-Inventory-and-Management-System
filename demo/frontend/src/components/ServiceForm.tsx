import { useState, useEffect } from 'react';
import { Service, CreateServiceRequest } from '../types/service';

interface ServiceFormProps {
  service?: Service | null;
  onSave: (data: CreateServiceRequest) => void;
  onCancel: () => void;
}

export function ServiceForm({ service, onSave, onCancel }: ServiceFormProps) {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [category, setCategory] = useState('');
  const [status, setStatus] = useState<'active' | 'inactive'>('active');

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
  }, [service]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSave({ name, description, category, status });
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
