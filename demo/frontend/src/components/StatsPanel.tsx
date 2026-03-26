import { ServiceStats } from '../types/service';

interface StatsPanelProps {
  stats: ServiceStats | null;
  loading?: boolean;
}

export function StatsPanel({ stats, loading }: StatsPanelProps) {
  if (loading) {
    return <div className="stats-panel stats-panel--loading">Загрузка статистики...</div>;
  }

  if (!stats) {
    return null;
  }

  return (
    <div className="stats-panel">
      <div className="stats-card stats-card--total">
        <div className="stats-card__label">Всего сервисов</div>
        <div className="stats-card__value">{stats.total}</div>
      </div>

      <div className="stats-card">
        <div className="stats-card__label">По статусу</div>
        <div className="stats-card__items">
          {stats.by_status.map((item) => (
            <span
              key={item.key}
              className={`stats-badge stats-badge--${item.key}`}
            >
              {item.key}: {item.count}
            </span>
          ))}
        </div>
      </div>

      <div className="stats-card">
        <div className="stats-card__label">Категории</div>
        <div className="stats-card__items">
          {stats.by_category.map((item) => (
            <div key={item.key} className="stats-category-item">
              <span className="stats-category-item__name">{item.key || '—'}</span>
              <span className="stats-category-item__count">{item.count}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
