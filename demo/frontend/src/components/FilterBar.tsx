import { useState, useEffect } from 'react';
import { Service } from '../types/service';

interface FilterParams {
  category?: string;
  status?: string;
  search?: string;
}

interface FilterBarProps {
  services: Service[];
  onFilter: (params: FilterParams) => void;
}

export function FilterBar({ services, onFilter }: FilterBarProps) {
  const [search, setSearch] = useState('');
  const [status, setStatus] = useState('');
  const [category, setCategory] = useState('');

  // Collect unique categories from loaded services
  const categories = Array.from(new Set(services.map((s) => s.category).filter(Boolean))).sort();

  // Debounce search
  useEffect(() => {
    const timer = setTimeout(() => {
      onFilter({ search: search || undefined, status: status || undefined, category: category || undefined });
    }, 300);
    return () => clearTimeout(timer);
  }, [search, status, category, onFilter]);

  const handleReset = () => {
    setSearch('');
    setStatus('');
    setCategory('');
  };

  return (
    <div className="filter-bar">
      <input
        type="text"
        className="filter-bar__search"
        placeholder="Поиск по названию..."
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />

      <select
        className="filter-bar__select"
        value={status}
        onChange={(e) => setStatus(e.target.value)}
      >
        <option value="">Все статусы</option>
        <option value="active">Active</option>
        <option value="inactive">Inactive</option>
      </select>

      <select
        className="filter-bar__select"
        value={category}
        onChange={(e) => setCategory(e.target.value)}
      >
        <option value="">Все категории</option>
        {categories.map((cat) => (
          <option key={cat} value={cat}>
            {cat}
          </option>
        ))}
      </select>

      <button className="filter-bar__reset" onClick={handleReset}>
        Сбросить
      </button>
    </div>
  );
}
