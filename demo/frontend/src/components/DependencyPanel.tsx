import { useState, useEffect } from 'react';
import { Service, ServiceDependencies } from '../types/service';
import { getDependencies, addDependency, removeDependency } from '../api/services';

interface DependencyPanelProps {
  service: Service;
  allServices: Service[];
  onClose: () => void;
}

export function DependencyPanel({ service, allServices, onClose }: DependencyPanelProps) {
  const [deps, setDeps] = useState<ServiceDependencies | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [selectedDependsOn, setSelectedDependsOn] = useState('');

  const loadDeps = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await getDependencies(service.id);
      setDeps(data);
    } catch {
      setError('Не удалось загрузить зависимости');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadDeps();
  }, [service.id]);

  const handleAdd = async () => {
    if (!selectedDependsOn) return;
    try {
      setError(null);
      await addDependency(service.id, selectedDependsOn);
      setSelectedDependsOn('');
      await loadDeps();
    } catch (err: unknown) {
      const message =
        err && typeof err === 'object' && 'response' in err
          ? (err as { response?: { data?: { error?: string } } }).response?.data?.error
          : undefined;
      setError(message || 'Не удалось добавить зависимость');
    }
  };

  const handleRemove = async (depId: string) => {
    try {
      setError(null);
      await removeDependency(service.id, depId);
      await loadDeps();
    } catch {
      setError('Не удалось удалить зависимость');
    }
  };

  // Services that are already in depends_on or are the current service
  const alreadyDependsOnIds = new Set(deps?.depends_on.map((s) => s.id) || []);
  const availableServices = allServices.filter(
    (s) => s.id !== service.id && !alreadyDependsOnIds.has(s.id)
  );

  return (
    <div className="dep-panel">
      <div className="dep-panel__header">
        <h3>Зависимости: {service.name}</h3>
        <button className="dep-panel__close" onClick={onClose}>
          &times;
        </button>
      </div>

      {error && <div className="dep-panel__error">{error}</div>}

      {loading ? (
        <p>Загрузка...</p>
      ) : (
        <>
          <div className="dep-panel__section">
            <h4>Зависит от:</h4>
            {deps?.depends_on.length === 0 ? (
              <p className="dep-panel__empty">Нет зависимостей</p>
            ) : (
              <ul className="dep-panel__list">
                {deps?.depends_on.map((s) => (
                  <li key={s.id} className="dep-panel__item">
                    <span
                      className={`dep-panel__status dep-panel__status--${s.status}`}
                    >
                      {s.status}
                    </span>
                    <span className="dep-panel__name">{s.name}</span>
                    <span className="dep-panel__category">{s.category}</span>
                    <button
                      className="dep-panel__remove"
                      onClick={() => handleRemove(s.id)}
                      title="Удалить зависимость"
                    >
                      &times;
                    </button>
                  </li>
                ))}
              </ul>
            )}
          </div>

          <div className="dep-panel__section">
            <h4>Используется в:</h4>
            {deps?.used_by.length === 0 ? (
              <p className="dep-panel__empty">Никем не используется</p>
            ) : (
              <ul className="dep-panel__list">
                {deps?.used_by.map((s) => (
                  <li key={s.id} className="dep-panel__item">
                    <span
                      className={`dep-panel__status dep-panel__status--${s.status}`}
                    >
                      {s.status}
                    </span>
                    <span className="dep-panel__name">{s.name}</span>
                    <span className="dep-panel__category">{s.category}</span>
                  </li>
                ))}
              </ul>
            )}
          </div>

          <div className="dep-panel__add">
            <h4>Добавить зависимость:</h4>
            <div className="dep-panel__add-form">
              <select
                className="dep-panel__select"
                value={selectedDependsOn}
                onChange={(e) => setSelectedDependsOn(e.target.value)}
              >
                <option value="">Выберите сервис...</option>
                {availableServices.map((s) => (
                  <option key={s.id} value={s.id}>
                    {s.name} ({s.category})
                  </option>
                ))}
              </select>
              <button
                className="btn-primary"
                onClick={handleAdd}
                disabled={!selectedDependsOn}
              >
                Добавить зависимость
              </button>
            </div>
          </div>
        </>
      )}
    </div>
  );
}
