import { Service } from '../types/service';

interface ServiceCardProps {
  service: Service;
  onEdit: (service: Service) => void;
  onDelete: (id: string) => void;
}

const statusBadgeStyle = (status: 'active' | 'inactive'): React.CSSProperties => ({
  display: 'inline-block',
  padding: '2px 10px',
  borderRadius: '12px',
  fontSize: '12px',
  fontWeight: 600,
  backgroundColor: status === 'active' ? '#d4edda' : '#e2e3e5',
  color: status === 'active' ? '#155724' : '#383d41',
});

export function ServiceCard({ service, onEdit, onDelete }: ServiceCardProps) {
  return (
    <tr style={{ borderBottom: '1px solid #dee2e6' }}>
      <td style={{ padding: '12px 16px' }}>{service.name}</td>
      <td style={{ padding: '12px 16px', color: '#6c757d' }}>{service.description || '—'}</td>
      <td style={{ padding: '12px 16px' }}>{service.category || '—'}</td>
      <td style={{ padding: '12px 16px' }}>
        <span style={statusBadgeStyle(service.status)}>
          {service.status === 'active' ? 'Активен' : 'Неактивен'}
        </span>
      </td>
      <td style={{ padding: '12px 16px' }}>
        <button
          onClick={() => onEdit(service)}
          style={{
            marginRight: '8px',
            padding: '6px 14px',
            backgroundColor: '#fd7e14',
            color: '#fff',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '13px',
          }}
        >
          Редактировать
        </button>
        <button
          onClick={() => onDelete(service.id)}
          style={{
            padding: '6px 14px',
            backgroundColor: '#dc3545',
            color: '#fff',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '13px',
          }}
        >
          Удалить
        </button>
      </td>
    </tr>
  );
}
