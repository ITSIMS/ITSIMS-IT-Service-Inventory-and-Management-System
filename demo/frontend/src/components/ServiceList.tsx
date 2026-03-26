import { Service } from '../types/service';
import { ServiceCard } from './ServiceCard';

interface ServiceListProps {
  services: Service[];
  onEdit: (service: Service) => void;
  onDelete: (id: string) => void;
  onRowClick?: (service: Service) => void;
  selectedId?: string;
}

export function ServiceList({ services, onEdit, onDelete, onRowClick, selectedId }: ServiceListProps) {
  if (services.length === 0) {
    return (
      <p style={{ color: '#6c757d', textAlign: 'center', marginTop: '40px' }}>
        Нет сервисов. Добавьте первый сервис.
      </p>
    );
  }

  return (
    <div style={{ overflowX: 'auto' }}>
      <table
        style={{
          width: '100%',
          borderCollapse: 'collapse',
          backgroundColor: '#fff',
          boxShadow: '0 1px 4px rgba(0,0,0,0.08)',
          borderRadius: '6px',
          overflow: 'hidden',
        }}
      >
        <thead>
          <tr style={{ backgroundColor: '#f8f9fa', borderBottom: '2px solid #dee2e6' }}>
            <th style={{ padding: '12px 16px', textAlign: 'left', fontWeight: 600 }}>Название</th>
            <th style={{ padding: '12px 16px', textAlign: 'left', fontWeight: 600 }}>Описание</th>
            <th style={{ padding: '12px 16px', textAlign: 'left', fontWeight: 600 }}>Категория</th>
            <th style={{ padding: '12px 16px', textAlign: 'left', fontWeight: 600 }}>Статус</th>
            <th style={{ padding: '12px 16px', textAlign: 'left', fontWeight: 600 }}>Действия</th>
          </tr>
        </thead>
        <tbody>
          {services.map((service) => (
            <ServiceCard
              key={service.id}
              service={service}
              onEdit={onEdit}
              onDelete={onDelete}
              onRowClick={onRowClick}
              isSelected={selectedId === service.id}
            />
          ))}
        </tbody>
      </table>
    </div>
  );
}
